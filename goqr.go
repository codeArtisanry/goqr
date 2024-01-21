package goqr

import (
	"image"
	"image/png"
	"os"

	"github.com/skip2/go-qrcode"
)

// GenerateAndSave generates a QR code for the given URL and saves it as a PNG file.
func GenerateAndSave(url string, outputPath string, savePng bool) (string, error) {
	// Generate QR code
	qrCode, err := qrcode.New(url, qrcode.High)
	if err != nil {
		return "Found error in generating QR code", err
	}

	// Print QR code to the console
	QR := qrCode.ToSmallString(false)

	if savePng {
		// Save QR code as PNG file
		err = savePNG(outputPath, qrCode.Image(256))
		if err != nil {
			return "Found error in saving PNG", err
		}
	}

	return QR, err
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
