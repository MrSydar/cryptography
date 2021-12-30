package steganography

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"strings"
	"testing"
)

func savePNG(path string, img image.Image) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		return err
	}

	return nil
}

func readImage(fName string) image.Image {
	f, err := os.Open(fName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return img
}
func TestDecypher(t *testing.T) {
	img := readImage("assets_test/puppy2.png")
	expectedText := "Hello, world!"

	cypheredImage, err := LsbCypher(img, expectedText)

	if err != nil {
		t.Fatalf("error was not expected: %v", err)
	}

	actualText := LsbDecypher(cypheredImage)

	if !strings.HasPrefix(actualText, expectedText) {
		t.Fatalf("expected decyphered text to have %s prefix, but the actual text is %s", expectedText, actualText)
	}
}

func TestCypher(t *testing.T) {
	img := readImage("assets_test/puppy2.png")
	cypheredText := "Hello, world!"

	cypheredImage, err := LsbCypher(img, cypheredText)

	if err != nil {
		t.Fatalf("error was not expected: %v", err)
	}

	lastBit := func(value uint32) string {
		if value&1 == 1 {
			return "1"
		} else {
			return "0"
		}
	}

	var sb strings.Builder
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := cypheredImage.At(x, y).RGBA()
			sb.WriteString(lastBit(r))
			sb.WriteString(lastBit(g))
			sb.WriteString(lastBit(b))
		}
	}

	expectedPrefix := "01001000011001010110110001101100011011110010110000100000011101110110111101110010011011000110010000100001"
	if actual := sb.String(); !strings.HasPrefix(actual, expectedPrefix) {
		t.Fatalf("actual binary cypher prefix %s expected to be %s", actual[:len(expectedPrefix)], expectedPrefix)
	}

	// savePNG("assets_test/cyphered.png", cypheredImage)
}
