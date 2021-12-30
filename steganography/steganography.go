package steganography

import (
	"errors"
	"image"
	"image/color"
	"strings"
)

func newBitConverter(text string) func() byte {
	characterIndex := 0
	currentCharacterBitIndex := 7
	return func() byte {
		if characterIndex == len(text) {
			return 0
		}

		bit := (text[characterIndex] & (1 << currentCharacterBitIndex)) >> byte(currentCharacterBitIndex)

		if currentCharacterBitIndex == 0 {
			currentCharacterBitIndex = 7
			characterIndex++
		} else {
			currentCharacterBitIndex--
		}

		return bit
	}
}

const (
	NoLastBitMask uint8 = ^uint8(1)
)

func setLastColorBit(color32 uint32, bit uint8) uint8 {
	color8 := uint8(color32)
	color8 &= NoLastBitMask
	color8 |= bit
	return color8
}

func LsbCypher(img image.Image, text string) (image.Image, error) {
	bounds := img.Bounds()

	if bounds.Dx()*bounds.Dy()/8 < len(text) {
		return nil, errors.New("the image is too small to encode given text")
	}

	nextBitFromText := newBitConverter(text)
	cyImageRGBA := image.NewRGBA(image.Rect(img.Bounds().Min.X, img.Bounds().Min.Y, img.Bounds().Dx(), img.Bounds().Dy()))
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			newR := setLastColorBit(r, nextBitFromText())
			newG := setLastColorBit(g, nextBitFromText())
			newB := setLastColorBit(b, nextBitFromText())

			cyColor := color.RGBA{newR, newG, newB, uint8(a)}
			cyImageRGBA.Set(x, y, cyColor)
		}
	}

	return cyImageRGBA, nil
}

func newCumulativeStringConverter() func(c byte) *strings.Builder {
	var sb strings.Builder

	i := 0
	bi := 7

	var currChar byte
	return func(c byte) *strings.Builder {
		currChar |= (c & 1) << bi

		if bi == 0 {
			sb.WriteByte(currChar)
			currChar = 0
			bi = 7
			i++
		} else {
			bi--
		}

		return &sb
	}
}

func LsbDecypher(img image.Image) string {
	bounds := img.Bounds()
	appendChar := newCumulativeStringConverter()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			appendChar(byte(r))
			appendChar(byte(g))
			appendChar(byte(b))
		}
	}
	return appendChar(0).String()
}
