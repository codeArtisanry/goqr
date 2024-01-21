package main

import (
	"flag"
	"fmt"
	"log"

	goqr "github.com/codeArtisanry/goqr" 
)

func main() {

	url := flag.String("url", "", "URL for generating QR code")
	output := flag.String("output", "qrcode.png", "Output file for QR code PNG")
	savePng := flag.Bool("savePng", false, "Save QR code as PNG file")

	// Parse command-line flags
	flag.Parse()

	// Check if the URL is provided
	if *url == "" {
		fmt.Println("Please provide a URL using the -url flag.")
		return
	}

	fmt.Println(" Generating QR code for URL:\n", *url)

	qr, err := goqr.GenerateAndSave(*url, *output, *savePng)
	if err != nil {
		log.Fatal("Error generating and saving QR code:", err)
	}

	fmt.Println("Generated QR Code:")
	fmt.Print(qr)

}
