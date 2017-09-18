package imgsim

import (
	"fmt"
	"image"
	"math/bits"

	"github.com/nfnt/resize"
)

// Hash is the type of the average hash. The average has is computed by setting the value of a bit
// to 1 if it's pixel value exceeds the image average pixel value, otherwise it is 0.
// These bits are stored in a 64 bit unsigned integer.
type Hash uint64

func (h Hash) String() string {
	return fmt.Sprintf("%b", h)
}

// AverageHash calculates the average hash of an image. First the image is resized to 8x8.
// Then it is converted to grayscale. Lastly the average hash is computed.
func AverageHash(img image.Image) Hash {
	img = resize.Resize(8, 8, img, resize.NearestNeighbor)
	img = rgbaToGray(img)
	mean := mean(img)
	return calcHash(img, mean)
}

// Distance calculates the number of different bits in the hash
func Distance(a, b Hash) int {
	return bits.OnesCount64(uint64(a ^ b))
}

// rbgaToGray converts an rgba image to a greyscale image
func rgbaToGray(img image.Image) *image.Gray {
	var (
		bounds = img.Bounds()
		gray   = image.NewGray(bounds)
	)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			var rgba = img.At(x, y)
			gray.Set(x, y, rgba)
		}
	}
	return gray
}

// mean computes the mean of all pixels.
func mean(img image.Image) uint32 {

	rect := img.Bounds()
	w := rect.Max.X - rect.Min.X
	h := rect.Max.Y - rect.Min.Y
	t := uint32(w * h)

	if t == 0 {
		return 0
	}

	var x, y int
	var r, sum uint32
	for x = rect.Min.X; x < rect.Max.X; x++ {
		for y = rect.Min.Y; y < rect.Max.Y; y++ {
			r, _, _, _ = img.At(x, y).RGBA()
			sum += r
		}
	}

	return sum / t
}

// calcHash computes the average hash for the given image and mean.
func calcHash(img image.Image, mean uint32) Hash {
	var x, y int
	var hash, p Hash
	p = 1
	var r uint32

	rect := img.Bounds()

	for y = rect.Min.Y; y < rect.Max.Y; y++ {
		for x = rect.Min.X; x < rect.Max.X; x++ {
			r, _, _, _ = img.At(x, y).RGBA()

			if r > mean {
				hash |= p
			}

			p = p << 1
		}
	}

	return hash
}
