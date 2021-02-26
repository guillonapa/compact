/*
Package compressor provides functionality translate between
ImageIr and CompactImage, as defined in the internal package.
*/
package compressor

import "github.com/guillonapa/compact/internal"

/*
Constants used to specify the type of compression.
The options are SIMPLE, and EXPERIMENTAL.
*/
const (
	SIMPLE       = "simple"
	EXPERIMENTAL = "experimental"
)

// Compress transforms an internal image representation into a compact image
// representation. Currently, the compressorType is ignored.
func Compress(i internal.ImageIr, compressorType string) (internal.CompactImage, error) {
	return internal.CompactImage{Dummy: i}, nil
}

// Decompress transforms a compact image into an intermal image representation.
// Currently, the compressorType is ignored.
func Decompress(i internal.CompactImage) (internal.ImageIr, error) {
	return i.Dummy, nil
}
