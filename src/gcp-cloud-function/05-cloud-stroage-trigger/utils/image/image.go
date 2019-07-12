package image

import (
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

func IsPNG(r io.Reader) bool {
	_, err := png.Decode(r)
	if err != nil {
		fmt.Printf("png.Decode %v\n", err)
		return false
	}
	return true
}

func IsJPG(r io.Reader) bool {
	_, err := jpeg.Decode(r)
	if err != nil {
		fmt.Printf("jpeg.Decode %v\n", err)
		return false
	}
	return true
}

func IsGIF(r io.Reader) bool {
	_, err := gif.Decode(r)
	if err != nil {
		fmt.Printf("gif.Decode %v\n", err)
		return false
	}
	return true
}
