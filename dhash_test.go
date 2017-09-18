package imgsim

import (
	"image"
	"image/png"
	"os"
	"testing"
)

func TestDifferenceHash(t *testing.T) {
	imgfile, err := os.Open("assets/gopher.png")
	defer imgfile.Close()
	if err != nil {
		t.Errorf("failed to open image: %s", err)
	}
	img, err := png.Decode(imgfile)
	if err != nil {
		t.Errorf("failed to decode image: %s", err)
	}

	type args struct {
		img image.Image
	}
	tests := []struct {
		name string
		args args
		want Hash
	}{
		{"gopher image", args{img}, 10233023882142496138},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DifferenceHash(tt.args.img); got != tt.want {
				t.Errorf("DifferenceHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
