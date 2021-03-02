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
	"io"
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

// ImageIrReader is used to read pixels from the image.
type ImageIrReader struct {
	i ImageIr
	x int
	y int
}

// Read reads pixels into a slice. There's is no indication
// of where a row ends or starts.
func (r *ImageIrReader) Read(b [][]uint32) (int, error) {
	read := 0
	for read < len(b) {
		// check if there's nothing to read
		if r.x >= r.i.width {
			return read, io.EOF
		}
		// read the pixel at (x, y)
		p := r.i.pixels[r.x][r.y]
		b[read] = p
		read = read + 1
		// update current position
		if r.y == r.i.height-1 {
			r.x = r.x + 1
			r.y = 0
		} else {
			r.y = r.y + 1
		}
	}
	// check if we are done reading the image
	if r.x >= r.i.width {
		return read, io.EOF
	}
	return read, nil
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

// CompactImageReader is used to read pixels from the
// internal representation for a compact image.
type CompactImageReader struct {
	imageReader ImageIrReader
}

// compactImageReader creates an instance of the reader.
func compactImageReader(c CompactImage) CompactImageReader {
	return CompactImageReader{imageReader: ImageIrReader{i: c.Dummy, x: 0, y: 0}}
}

// Read reads pixels into a slice. There's is no indication
// of where a row ends or starts.
func (r *CompactImageReader) Read(b [][]uint32) (int, error) {
	return r.imageReader.Read(b)
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
	// open the file to write to
	f, err := os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// truncate to size 0
	err = f.Truncate(0)
	if err != nil {
		return err
	}

	// read rows of data as we attempt to write them
	r := compactImageReader(c)
	b := make([][]uint32, c.Dummy.height)
	count := 0
	for {
		n, err := r.Read(b)
		count += n
		if _, writeErr := f.WriteString(scompact(b)); writeErr != nil {
			return writeErr
		}
		if err == io.EOF {
			break
		}
	}
	fmt.Printf("[INFO]: Number of pixels read: %v\n", count)
	return nil
}

// scompact creates a string representation of the pixel data for a row.
func scompact(pixels [][]uint32) string {
	var sb strings.Builder
	sb.Reset()
	for i, p := range pixels {
		if i != 0 {
			sb.WriteString(",")
		}
		sb.WriteString(fmt.Sprintf("%v|%v|%v|%v", p[0], p[1], p[2], p[3]))
	}
	sb.WriteString("\n")
	return sb.String()
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
