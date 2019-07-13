package helpers

import (
	"fmt"
	"github.com/ttacon/chalk"
	"os"
	"time"
)

// Log logs in green.
func Log(s string) {
	service := os.Getenv("GO_SERVICE")
	t := time.Now()
	fmt.Println(chalk.Green.Color(t.Format("2006-01-02 15:04:05") + ":" + service + ":" + s))
}

// LogErr logs in red.
func LogErr(s string) {
	service := os.Getenv("GO_SERVICE")
	t := time.Now()
	fmt.Println(chalk.Red.Color(t.Format("2006-01-02 15:04:05") + ":" + service + ":" + s))
}

// LogInfo logs in magenta.
func LogInfo(s string) {
	service := os.Getenv("GO_SERVICE")
	t := time.Now()
	fmt.Println(chalk.Magenta.Color(t.Format("2006-01-02 15:04:05") + ":" + service + ":" + s))
}

// LogWarn logs in yellow.
func LogWarn(s string) {
	service := os.Getenv("GO_SERVICE")
	t := time.Now()
	fmt.Println(chalk.Yellow.Color(t.Format("2006-01-02 15:04:05") + ":" + service + ":" + s))
}
