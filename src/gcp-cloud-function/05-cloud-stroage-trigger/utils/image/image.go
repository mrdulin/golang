package image

import (
	"fmt"
	"image/png"
	"io"
)

func IsPNG(r io.Reader) bool {
	_, err := png.Decode(r)
	if err != nil {
		fmt.Printf("png.Decode %v", err)
		return false
	}
	return true
}
