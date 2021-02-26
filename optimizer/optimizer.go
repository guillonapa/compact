/*
Package optimizer provides functionality to transform an
image into an optimized version of itself, depending on
the type of media for which the image might be used.
*/
package optimizer

import (
	"github.com/guillonapa/compact/internal"
)

// Media type constants. The options are WEB, PRINT, and MOBILE
const (
	WEB    = "web"
	PRINT  = "print"
	MOBILE = "mobile"
)

// Optimize creates a copy of an internal image representation
// and transforms it depending on the optimizerType. Currently,
// however, the optimizerType is ignored.
func Optimize(i internal.ImageIr, optimizerType string) (internal.ImageIr, error) {
	c := internal.Copy(i)
	o, err := shrink(c, scalingFactor(c, optimizerType))
	if err != nil {
		return c, err
	}
	return o, nil
}

// shrink reduces the size of the image using the given scale,
// a number between 0.0 and 1.0.
func shrink(i internal.ImageIr, scale float64) (internal.ImageIr, error) {
	return i, nil
}

// scalingFactor returns the scaling factor for the given
// image and optimizerType.
func scalingFactor(i internal.ImageIr, optimizerType string) float64 {
	return 1
}
