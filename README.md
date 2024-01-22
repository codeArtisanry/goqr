# GoQR


GoQR is a simple Go package that provides an easy way to generate QR codes and save them as PNG files. It utilizes the popular github.com/skip2/go-qrcode library to generate QR codes with customizable options.

### Features
- Generate QR codes for URLs.
- Save generated QR codes as PNG files.
- Adjustable error correction level.

### Installation
To use GoQR in your Go project, you can run:
```
go get -u github.com/codeArtisanry/goqr
```

### Usage
```
package main

import (
	"fmt"
	"log"

	"github.com/codeArtisanry/goqr"
)

func main() {
	url := "https://vatsalchauhan.site"
	outputPath := "qrcode.png"

	result, err := goqr.GenerateAndSave(url, outputPath, true)
	if err != nil {
		log.Fatal("Error:", result, err)
	}

	fmt.Printf("QR code saved as %s\n", outputPath)
}
```

### License
This project is licensed under the MIT License - see the LICENSE file for details.

### Contributing
Feel free to open issues, submit pull requests, or provide suggestions to enhance the functionality of GoQR.

