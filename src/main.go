package main

import (
	"io/ioutil"
	"log"
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {

	file, err := os.Open("vcard.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	str := string(b)
	qrFname := "vcard-qr.png"

	err = qrcode.WriteFile(str, qrcode.Medium, 256, qrFname)
	if err != nil {
		log.Fatal(err)
	}

}
