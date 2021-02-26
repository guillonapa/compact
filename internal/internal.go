/*
Package internal deals with common operations that are
called from other packages.
*/
package internal

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// ImageIr is an internal representation for an image.
type ImageIr struct {
	width  int
	height int
	pixels [][][]uint32
}

// String returns the string representation for an internal
// image representation.
func (i ImageIr) String() string {
	return fmt.Sprintf("[%v x %v]:\n%v\n", i.width, i.height, i.pixels)
}

// CompactImage is an internal representation of a compressed
// image.
type CompactImage struct {
	Dummy ImageIr
}

// Copy creates a copy of the interal image representation.
func Copy(i ImageIr) ImageIr {
	c := make([][][]uint32, len(i.pixels))
	copy(c, i.pixels)
	return ImageIr{i.width, i.height, c}
}

// Read reads an image into an internal image representation.
func Read(i image.Image) ImageIr {
	b := i.Bounds()
	// get width and height
	h := b.Dy()
	w := b.Dx()
	// get all pixels
	minX := b.Min.X
	minY := b.Min.Y
	pixels := make([][][]uint32, w)
	for x := 0; x < w; x++ {
		pixels[x] = make([][]uint32, h)
		for y := 0; y < h; y++ {
			pixels[x][y] = make([]uint32, 4)
			c := i.At(minX+x, minY+y)
			p := pixels[x][y]
			r, g, b, a := c.RGBA()
			p[0] = r
			p[1] = g
			p[2] = b
			p[3] = a
		}
	}
	return ImageIr{w, h, pixels}
}

// Draw creates an image from an internal image representation.
func Draw(i ImageIr) image.Image {
	pixels := i.pixels
	// create the image of the needed size
	res := image.NewRGBA(image.Rect(0, 0, i.width, i.height))
	for x := 0; x < i.width; x++ {
		for y := 0; y < i.height; y++ {
			p := pixels[x][y]
			rgba := color.RGBA{uint8(p[0]), uint8(p[1]), uint8(p[2]), uint8(p[3])}
			res.Set(x, y, rgba)
		}
	}
	return res
}

// WriteCompactImage writes the compact image representation
// to the file system at the given path.
func WriteCompactImage(c CompactImage, path string) error {
	bytes := []byte(scompact(c.Dummy.pixels))
	err := ioutil.WriteFile(path, bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

// scompact creates a string representation of the pixel data.
// This is NOT a good implementation for medium to large images.
func scompact(pixels [][][]uint32) string {
	var s string
	for _, row := range pixels {
		for i, p := range row {
			if i != 0 {
				s += ","
			}
			s += fmt.Sprintf("%v|%v|%v|%v", p[0], p[1], p[2], p[3])
		}
		s += "\n"
	}
	return s
}

// WriteImage writes an image to the file system at the
// given path.
func WriteImage(i image.Image, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	// Encode to `PNG` with `DefaultCompression` level
	// then save to file
	err = png.Encode(f, i)
	if err != nil {
		return err
	}
	return nil
}

// ReadCompactImage reads a file and returns an internal representation
// of a compact image.
func ReadCompactImage(path string) (CompactImage, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return CompactImage{}, err
	}
	text := string(content)
	// construct the compact image
	pixels := make([][][]uint32, 0)
	for row, line := range strings.Split(text, "\n") {
		if line != "" {
			pixels = append(pixels, make([][]uint32, 0))
			for column, pixel := range strings.Split(line, ",") {
				pixels[row] = append(pixels[row], make([]uint32, 0))
				p := make([]uint32, 4)
				for i, v := range strings.Split(pixel, "|") {
					r, err := strconv.ParseUint(v, 10, 32)
					if err != nil {
						return CompactImage{}, err
					}
					p[i] = uint32(r)
				}
				pixels[row][column] = p
			}
		}
	}
	return CompactImage{ImageIr{width: len(pixels), height: len(pixels[0]), pixels: pixels}}, nil
}
