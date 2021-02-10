package compact

import (
	"image"
	"os"
)

// Version TODO
func Version() string {
	return "0.0.1-dev"
}

// Optimize TODO
func Optimize(i image.Image) (image.Image, error) {
	return nil, nil
}

// OptimizeAs TODO
func OptimizeAs(i image.Image, optimizerType string) (image.Image, error) {
	return nil, nil
}

// Compress TODO
func Compress(i image.Image, path string) (os.File, error) {
	return os.File{}, nil
}

// CompressAs TODO
func CompressAs(i image.Image, compressorType string, path string) (os.File, error) {
	return os.File{}, nil
}

// Decompress TODO
func Decompress(f os.File, path string) (image.Image, error) {
	return nil, nil
}

// DecompressAs TODO
func DecompressAs(f os.File, compressorType string, path string) (image.Image, error) {
	return nil, nil
}

// OptimizeAndCompress TODO
func OptimizeAndCompress(i image.Image, path string) (os.File, error) {
	return os.File{}, nil
}

// OptimizeAndCompressAs TODO
func OptimizeAndCompressAs(i image.Image, optimizerType string, compressorType string, path string) (os.File, error) {
	return os.File{}, nil
}
