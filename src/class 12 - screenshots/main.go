package main

import (
	"bytes"	// to create a buffer
	"fmt"     // format package
	"os/exec" // execute commands
	"time"
)

// First way using CombinedOutput
func captureEntireScreenshot(imageName string){
	// output, err
	_, err := exec.Command("import", "-window", "root", imageName).CombinedOutput()
	if err != nil {
		panic("Fatal error")
	}
}

// Second way using Buffers
func captureEntireScreenshotWithABuffer(imageName string){
	/*
		A buffer is a temporary holding area for data while\
		it's waiting to be transferred to another location.
		It is usually located in the RAM.
	*/
	var buffer bytes.Buffer
	outputCmd := exec.Command("import", "-window", "root", imageName)
	outputCmd.Stdout = &buffer
	outputCmd.Stderr = &buffer

	err := outputCmd.Run()
	if err != nil {
		panic("Fatal error")
	}
}

func getActualDate(dateFormat string) string{
	// time package uses this date formate as a reference: Mon Jan 2 15:04:05 MST 2006
	currentTime := time.Now()
	return  currentTime.Format(dateFormat)
}

func main(){
	const dateFormat = "(Mon) 02-01-2006 15h04m05s"
	const imgExt = "png"

	// Using the first way
	captureEntireScreenshot(fmt.Sprintf("%s.%s", getActualDate(dateFormat),imgExt))
	captureEntireScreenshot(fmt.Sprintf("%s.%s", getActualDate(dateFormat),imgExt))
	
	// Using the second way
	captureEntireScreenshotWithABuffer(fmt.Sprintf("(buffer) %s.%s", getActualDate(dateFormat),imgExt))
	captureEntireScreenshotWithABuffer(fmt.Sprintf("(buffer) %s.%s", getActualDate(dateFormat),imgExt))
}