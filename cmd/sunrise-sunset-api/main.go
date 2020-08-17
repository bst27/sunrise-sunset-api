package main

import (
	"fmt"
	"github.com/kelvins/sunrisesunset"
	"os"
	"time"
)

func main() {
	d, err := time.Parse("2006-01-02", "2020-04-27")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	p := sunrisesunset.Parameters{
		Latitude:  51.963691,
		Longitude: 7.610362,
		UtcOffset: 0,
		Date:      d,
	}

	sunrise, sunset, err := p.GetSunriseSunset()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(sunrise, sunset)
}
