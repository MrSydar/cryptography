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

func TestGetImage(t *testing.T) {
	s1 := readImage("s1.png")
	s2 := readImage("s2.png")

	img, err := GetImage(s1, s2)
	if err != nil {
		t.Fatalf("error was not expected: %v", err)
	}

	err = savePNG("combined.png", img)
	if err != nil {
		t.Fatalf("error was not expected: %v", err)
	}
}

func TestGetShare(t *testing.T) {
	img := readImage("original.png")
	s1, s2 := GetShares(img)

	err := savePNG("s1.png", s1)
	if err != nil {
		t.Fatalf("error was not expected: %v", err)
	}

	err = savePNG("s2.png", s2)
	if err != nil {
		t.Fatalf("error was not expected: %v", err)
	}
}
