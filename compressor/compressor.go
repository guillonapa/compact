package compressor

import "github.com/guillonapa/compact/internal"

// SIMPLE:
// EXPERIMENTAL:
const (
	SIMPLE       = "simple"
	EXPERIMENTAL = "experimental"
)

// Compress TODO
func Compress(i internal.ImageIr, compressorType string) (internal.CompactImage, error) {
	return internal.CompactImage{Dummy: i}, nil
}

// Decompress TODO
func Decompress(i internal.CompactImage) (internal.ImageIr, error) {
	return i.Dummy, nil
}
