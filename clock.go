package clock

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
	"time"
)

var seconds float64 = 0
var tick string = "tick"
var tock string = "tock"
var bong string = "bong"

// RunClock -
/*
Used ticker to run the method for each second in an till 3 hours to print tick, tock and bong values on the screen
After three hours
*/
func RunClock() {
	ticker := time.NewTicker(1 * time.Second) // Timer to for each second interval
	go func() {
		for range ticker.C {
			seconds = seconds + 1
			Update("input.txt")                 // File reading
			fmt.Println(GetPrintLabel(seconds)) // Prints the values for each second
		}
	}()

	// After 3 hours we'll stop the ticker
	time.Sleep(3 * time.Hour)
	ticker.Stop()
}

// GetPrintLabel -
/*
It takes each second as a value from 1 from the ticker start based on that second, It returns the label to print.
*/
func GetPrintLabel(secs float64) string {
	if secs >= 3600 && math.Mod(secs, float64(3600)) == 0 {
		return "bong"
	}
	if secs >= float64(60) && math.Mod(secs, float64(60)) == 0 {
		return tock
	}
	return tick
}

// Update -
/*
Reads the file and split the line seperated by comma
assigns the first three words to tick, tock and bong variables.
*/
func Update(fileName string) bool {
	line, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
		return false
	}
	str := string(line)
	input := strings.Split(str, ",")
	tick = input[0]
	tock = input[1]
	bong = input[2]
	return true
}
