package imgsim

import (
	"image"

	"github.com/nfnt/resize"
)

// AverageHash calculates the average hash of an image. First the image is resized to 8x8.
// Then it is converted to grayscale. Lastly the average hash is computed.
func AverageHash(img image.Image) Hash {
	img = resize.Resize(8, 8, img, resize.NearestNeighbor)
	img = rgbaToGray(img)
	mean := mean(img)
	return calcAvgHash(img, mean)
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

// calcAvgHash computes the average hash for the given image and mean.
func calcAvgHash(img image.Image, mean uint32) Hash {
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
