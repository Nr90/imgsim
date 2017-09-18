package imgsim

import (
	"image"

	"github.com/nfnt/resize"
)

// DifferenceHash calculates the average hash of an image. First the image is converted to grayscale
// Then it is resized to 9x8. Lastly the average hash is computed.
func DifferenceHash(img image.Image) Hash {
	img = rgbaToGray(img)
	img = resize.Resize(9, 8, img, resize.NearestNeighbor)
	return calcDiffHash(img)
}

// calcDiff computes the average hash for the given image and mean.
func calcDiffHash(img image.Image) Hash {
	var x, y int
	var hash, p Hash
	p = 1
	var r, left uint32

	rect := img.Bounds()

	for y = rect.Min.Y; y < rect.Max.Y; y++ {
		left, _, _, _ = img.At(rect.Min.X, y).RGBA()
		for x = rect.Min.X + 1; x < rect.Max.X; x++ {
			r, _, _, _ = img.At(x, y).RGBA()

			if r > left {
				hash |= p
			}

			p = p << 1
			left = r
		}
	}

	return hash
}
