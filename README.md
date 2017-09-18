# Imgsim [![GoDoc](https://godoc.org/github.com/Nr90/imgsim?status.svg)](https://godoc.org/github.com/Nr90/imgsim) [![Go Report Card](https://goreportcard.com/badge/github.com/Nr90/imgsim)](https://goreportcard.com/report/github.com/Nr90/imgsim) [![unstable](http://badges.github.io/stability-badges/dist/unstable.svg)](http://github.com/badges/stability-badges)
#


Imgsim is a library allows you to compute a fast average hash of an image. These hashes can then be used to compare images for similarity.

# Average hash #
An average hash is an example of a perceptual hash. Similar looking images will get similar perceptual hashes. Unlike cryptographic hashes
that would be very different for images with slight differences.
This makes them suitable to compare how similar images are.

# Installation #

The package is go-gettable: `go get -u github.com/Nr90/imgsim`. 

# Example #
```
package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/Nr90/imgsim
)

func main() {
	imgfile, err := os.Open("assets/gopher.png")
	defer imgfile.Close()
	if err != nil {
		panic(err)
	}
	img, err := png.Decode(imgfile)
	if err != nil {
		panic(err)
	}
	hash := imgsim.AverageHash(img)
	fmt.Println(hash)
}
```

# Notes #
This library uses the math/bits package introduced in go 1.9. Since 1.9 is so new, the math/bits package is also vendored in this package.
