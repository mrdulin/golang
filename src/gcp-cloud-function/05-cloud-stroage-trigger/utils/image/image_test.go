package image

import (
	"fmt"
	"os"
	"testing"
)

func TestIsPNG(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     bool
	}{
		{
			name:     "should be a png file when detect a png file",
			filename: "15625760447371547012340909WX20190108-124331.png",
			want:     true,
		},
		{
			name:     "should not be a png file when detect a pdf file",
			filename: "clean_arch.pdf",
			want:     false,
		},
		{
			name:     "should not be a png file when detect a fake png file",
			filename: "clean_arch.png",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := os.Open(fmt.Sprintf("../../tmp/%s", tt.filename))
			if err != nil {
				t.Fatalf("os.Open %v", err)
			}
			defer file.Close()

			got := IsPNG(file)
			if got != tt.want {
				t.Errorf("file: %s is png file = %t, want %t", file.Name(), got, tt.want)
			}
		})
	}
}
