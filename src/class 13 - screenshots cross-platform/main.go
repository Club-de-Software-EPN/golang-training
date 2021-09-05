package main

import (
 	"github.com/kbinani/screenshot"
 	"image/png"
 	"os"
 	"fmt"
	"time"
)

func captureEntireScreenshot(imageName string){
	const imgExt = "png"
	numberOfDisplays := screenshot.NumActiveDisplays()
	for i := 0; i < numberOfDisplays; i++ {
		bounds := screenshot.GetDisplayBounds(i) // for example: (0,0) - (1920,1080)
		capture, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		imageName := fmt.Sprintf("%d_%s.%s", i, imageName, imgExt)
		file, _ := os.Create(imageName)
		defer file.Close()
		png.Encode(file, capture)
		fmt.Println("A screenshot has been generated")
	}
}

func getActualDate(dateFormat string) string{
	// time package uses this date formate as a reference: Mon Jan 2 15:04:05 MST 2006
	currentTime := time.Now()
	return  currentTime.Format(dateFormat)
}

func main() {
	const dateFormat = "(Mon) 02-01-2006 15h04m05s"
	const seconds = 60
	for {
		captureEntireScreenshot(getActualDate(dateFormat))
		time.Sleep(seconds*time.Second)
	}
}