package optimizer

import (
	"github.com/guillonapa/compact/internal"
)

/*
Media type constants.

	WEB: web pages.
	PRINT: printed publications.
	MOBILE: mobile apps.
*/
const (
	WEB    = "web"
	PRINT  = "print"
	MOBILE = "mobile"
)

// Optimize TODO
func Optimize(i internal.ImageIr, optimizerType string) (internal.ImageIr, error) {
	o, err := shrink(i, scalingFactor(i, optimizerType))
	if err != nil {
		return i, err
	}
	return o, nil
}

// shrink TODO
func shrink(i internal.ImageIr, scale float64) (internal.ImageIr, error) {
	return i, nil
}

// scalingFactor TODO
func scalingFactor(i internal.ImageIr, optimizerType string) float64 {
	return 1
}
