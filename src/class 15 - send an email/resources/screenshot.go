package resources

import(
	"github.com/kbinani/screenshot"
	"image/png"
	"os"
	"fmt"
)

func CaptureEntireScreenshot(imageName string){
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
