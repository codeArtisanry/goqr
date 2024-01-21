package goqr

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/skip2/go-qrcode"
)

// GenerateAndSave generates a QR code for the given URL and saves it as a PNG file.
func GenerateAndSave(url string, outputPath string, savePng bool) error {
	// Generate QR code
	qrCode, err := qrcode.New(url, qrcode.High)
	if err != nil {
		return err
	}

	// Print QR code to the console
	fmt.Println("Generated QR Code:")
	fmt.Println(qrCode.ToSmallString(false))

	if savePng {
		// Save QR code as PNG file
		err = savePNG(outputPath, qrCode.Image(256))
		if err != nil {
			return err
		}
	}

	return nil
}

// savePNG saves an image as a PNG file
func savePNG(filename string, img image.Image) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return err
	}

	return nil
}
