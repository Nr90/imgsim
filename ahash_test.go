package imgsim

import (
	"image"
	"image/png"
	"os"
	"testing"
)

func TestAverageHash(t *testing.T) {
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
		{"gopher image", args{img}, 9331034559709847552},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AverageHash(tt.args.img); got != tt.want {
				t.Errorf("AverageHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mean(t *testing.T) {
	imgfile, err := os.Open("assets/gopher.png")
	defer imgfile.Close()
	if err != nil {
		t.Errorf("failed to open image: %s", err)
	}
	img, err := png.Decode(imgfile)
	if err != nil {
		t.Errorf("failed to decode image: %s", err)
	}

	emptyimg := image.NewGray(image.Rectangle{})

	type args struct {
		img image.Image
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{"gopher image", args{img}, 9872},
		{"0x0 image", args{emptyimg}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mean(tt.args.img); got != tt.want {
				t.Errorf("mean() = %v, want %v", got, tt.want)
			}
		})
	}
}
