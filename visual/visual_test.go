package visual

import (
	"fmt"
	"image"
	"os"
	"testing"

	"image/png"
	_ "image/png"
)

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

func TestGenerateAndCombineShares(t *testing.T) {
	os.Remove("assets_test/s1.png")
	os.Remove("assets_test/s2.png")
	os.Remove("assets_test/combined.png")

	img := readImage("assets_test/original.png")
	s1, s2 := GetShares(img)

	err := savePNG("assets_test/s1.png", s1)
	if err != nil {
		t.Fatalf("error was not expected: %v", err)
	}

	err = savePNG("assets_test/s2.png", s2)
	if err != nil {
		t.Fatalf("error was not expected: %v", err)
	}

	s1 = readImage("assets_test/s1.png")
	s2 = readImage("assets_test/s2.png")

	img, err = GetImage(s1, s2)
	if err != nil {
		t.Fatalf("error was not expected: %v", err)
	}

	err = savePNG("assets_test/combined.png", img)
	if err != nil {
		t.Fatalf("error was not expected: %v", err)
	}
}
