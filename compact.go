/*
Package compact implements a simple library for optimizing and
compressing images.
*/
package compact

import (
	"image"

	"github.com/guillonapa/compact/compressor"
	"github.com/guillonapa/compact/internal"
	"github.com/guillonapa/compact/optimizer"
)

// current version of the compact library
const version = "0.0.1"

// Version returns the current version of the package.
func Version() string {
	return version
}

// Optimize returns a copy of the image optimized for
// the web.
func Optimize(i image.Image) (image.Image, error) {
	return OptimizeAs(i, optimizer.WEB)
}

// OptimizeAs returns a copy of the image optimized for
// the specified type of media.
func OptimizeAs(i image.Image, optimizerType string) (image.Image, error) {
	res, err := optimizer.Optimize(internal.Read(i), optimizerType)
	if err != nil {
		return nil, err
	}
	return internal.Draw(res), nil
}

// Compress compresses a file using the default compression
// method and returns the compressed file.
func Compress(i image.Image, path string) error {
	return CompressAs(i, compressor.SIMPLE, path)
}

// CompressAs compresses a file using the specified compression
// method and returns the compressed file.
func CompressAs(i image.Image, compressorType string, path string) error {
	res, err := compressor.Compress(internal.Read(i), compressorType)
	if err != nil {
		return err
	}
	err = internal.WriteCompactImage(res, path)
	if err != nil {
		return err
	}
	return nil
}

// Decompress reads a compressed file, as compressed by 'compact',
// and returns the original image, while writing the uncompressed
// image to the file system if the path is provided.
func Decompress(compressedPath string, path string) (image.Image, error) {
	cimg, err := internal.ReadCompactImage(compressedPath)
	if err != nil {
		return nil, err
	}
	res, err := compressor.Decompress(cimg)
	if err != nil {
		return nil, err
	}
	i := internal.Draw(res)
	if path != "" {
		internal.WriteImage(i, path)
	}
	return i, nil
}

// OptimizeAndCompress optimizes the image for the web and compresses
// the resulting image. The compressed file is then returned.
func OptimizeAndCompress(i image.Image, path string) error {
	return OptimizeAndCompressAs(i, optimizer.WEB, compressor.SIMPLE, path)
}

// OptimizeAndCompressAs optimizes the image for the specified type of media
// and compresses the resulting image using the specified method. The
// compressed file is then returned.
func OptimizeAndCompressAs(i image.Image, optimizerType string, compressorType string, path string) error {
	res, err := OptimizeAs(i, optimizerType)
	if err != nil {
		return err
	}
	err = CompressAs(res, compressorType, path)
	if err != nil {
		return err
	}
	return nil
}
