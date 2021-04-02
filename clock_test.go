package clock

import (
	"log"
	"testing"
)

// TestUpdate - Test scenario for Postivie and Negetive checking for file reading
func TestUpdate(t *testing.T) {
	updated := Update("input.txt")
	if !updated {
		t.Errorf("Test Update Positive Case FAILED")

	}
	updated = Update("input1.txt")
	if updated {
		t.Errorf("Expected Test case to be Update Failed")
	}
	t.Log("PASSED")
}

func TestRunClock(t *testing.T) {
	// RunClock()
}

// TestGetPrintLabel - Test scenarios for printing the labels
func TestGetPrintLabel(t *testing.T) {
	if GetPrintLabel(1) != "tick" {
		t.Error("FAILED")
	}
	if GetPrintLabel(60) == "tick" {
		log.Println(GetPrintLabel(60))
		t.Error("FAILED")
	}
	if GetPrintLabel(60) != "tock" {
		t.Error("FAILED")
	}
	if GetPrintLabel(67) == "tock" {
		t.Error("FAILED")
	}
	log.Println(GetPrintLabel(3600), "&&&&&&")
	if GetPrintLabel(3600) != "bong" {
		t.Error("FAILED")
	}
	if GetPrintLabel(10) == "bong" {
		t.Error("FAILED")
	}
}
