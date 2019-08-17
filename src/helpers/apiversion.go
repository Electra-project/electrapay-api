package helpers

import (
	"fmt"
	"os"
	"strconv"
)

func GetVersion() (s string) {
	var version = os.Getenv("VERSION")
	v, err := strconv.ParseFloat(version, 32)
	if err != nil {
		v = 1
	}
	vdir := fmt.Sprint("v", v)
	return vdir
}

func GetVersionNumber() (i float64) {
	var version = os.Getenv("VERSION")
	v, err := strconv.ParseFloat(version, 32)
	if err != nil {
		v = 1
	}
	return v
}
