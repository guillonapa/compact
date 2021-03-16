package internal

import (
	"fmt"
	"testing"
)

func TestPixelsAsString(t *testing.T) {
	p := make([][]uint32, 4)
	for i := range p {
		p[i] = make([]uint32, 4)
		if i%2 == 0 {
			p[i][0] = 0
			p[i][1] = 1
			p[i][2] = 2
			p[i][3] = 3
		} else {
			p[i][0] = 256
			p[i][1] = 255
			p[i][2] = 254
			p[i][3] = 253
		}
	}
	odd := "256|255|254|253"
	even := "0|1|2|3"
	expected := fmt.Sprintf("%v,%v,%v,%v", even, odd, even, odd)
	result := scompact(p)
	if expected != result {
		t.Fatalf("Image pixel row stringification failed: got \"%v\", but expected \"%v\"", result, expected)
	}
}
