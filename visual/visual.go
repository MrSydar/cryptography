package visual

import (
	"errors"
	"image"
	"math/rand"
	"time"

	"image/color"
)

type pattern int

const (
	BW pattern = 0b01
	WB pattern = 0b10
)

type PatternedImage struct {
	pImg [][]pattern
}

func (pImg *PatternedImage) toImage() image.Image {
	height, width := len(pImg.pImg), len(pImg.pImg[0])*2

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width/2; x++ {
			if pImg.pImg[y][x] == BW {
				img.Set(2*x, y, color.Black)
				img.Set(2*x+1, y, color.White)
			} else {
				img.Set(2*x, y, color.White)
				img.Set(2*x+1, y, color.Black)
			}
		}
	}

	return img
}

func GetShares(img image.Image) (s1 image.Image, s2 image.Image) {
	rand.Seed(time.Now().UnixNano())

	bounds := img.Bounds()

	ps1 := PatternedImage{make([][]pattern, bounds.Dy())}
	ps2 := PatternedImage{make([][]pattern, bounds.Dy())}

	for y, iy := 0, bounds.Min.Y; iy < bounds.Max.Y; iy++ {
		ps1.pImg[iy] = make([]pattern, bounds.Dx())
		ps2.pImg[iy] = make([]pattern, bounds.Dx())

		for x, ix := 0, bounds.Min.X; ix < bounds.Max.X; ix++ {
			r, g, b, _ := img.At(ix, iy).RGBA()

			if (r+g+b)/3 < 65535/2 {
				if rand.Int()%2 == 0 {
					ps1.pImg[y][x], ps2.pImg[y][x] = BW, WB
				} else {
					ps1.pImg[y][x], ps2.pImg[y][x] = WB, BW
				}
			} else {
				if rand.Int()%2 == 0 {
					ps1.pImg[y][x], ps2.pImg[y][x] = BW, BW
				} else {
					ps1.pImg[y][x], ps2.pImg[y][x] = WB, WB
				}
			}

			x++
		}

		y++
	}

	return ps1.toImage(), ps2.toImage()
}

func GetImage(shares ...image.Image) (img image.Image, err error) {
	if len(shares) == 0 {
		return nil, errors.New("no shares provided")
	}

	sharesBounds := shares[0].Bounds()
	for _, s := range shares {
		if s.Bounds() != sharesBounds {
			return nil, errors.New("shares bounds are not the same")
		}
	}

	type changeableImage interface {
		Set(x, y int, c color.Color)
	}

	chImg, ok := shares[0].(changeableImage)
	if !ok {
		return nil, errors.New("image has immutable type")
	}

	for _, s := range shares[1:] {
		for y := sharesBounds.Min.Y; y < sharesBounds.Max.Y; y++ {
			for x := sharesBounds.Min.X; x < sharesBounds.Max.X; x++ {
				r, g, b, _ := s.At(x, y).RGBA()
				if (r+g+b)/3 <= 0 {
					chImg.Set(x, y, color.Black)
				}
			}
		}
	}

	img, ok = chImg.(image.Image)
	if !ok {
		return nil, errors.New("type assertion error")
	}

	return img, nil
}
