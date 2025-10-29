package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Failed to get NTP time: %s", err))
	}

	fmt.Printf("ntp: %s\n", ntpTime.Format(time.RFC1123))
	fmt.Printf("local: %s\n", time.Now().Format(time.RFC1123))
}
