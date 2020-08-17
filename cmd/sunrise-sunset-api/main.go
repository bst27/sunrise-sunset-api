package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/kelvins/sunrisesunset"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	debug := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()

	if *debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Time": time.Now().Unix(),
		})
	})

	r.GET("/sunrise-sunset", func(c *gin.Context) {
		lat, err := strconv.ParseFloat(c.Query("latitude"), 64) // TODO: Simplify input validation and extract code
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Invalid latitude",
			})
			return
		}

		lon, err := strconv.ParseFloat(c.Query("longitude"), 64)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Invalid longitude",
			})
			return
		}

		utcOffset, err := strconv.ParseFloat(c.Query("utcOffset"), 64)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Invalid utcOffset",
			})
			return
		}

		date, err := time.Parse("2006-01-02", c.Query("date"))
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Invalid date",
			})
			return
		}

		p := sunrisesunset.Parameters{
			Latitude:  lat,
			Longitude: lon,
			UtcOffset: utcOffset,
			Date:      date,
		}

		sunrise, sunset, err := p.GetSunriseSunset()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Failed to determine sunrise/sunset",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Latitude":  lat,
			"Longitude": lon,
			"UtcOffset": utcOffset,
			"Date":      date.Format("2006-01-02"),
			"Sunrise":   sunrise.Format("15:04:05"),
			"Sunset":    sunset.Format("15:04:05"),
		})
	})

	log.Println(r.Run(":8080")) // TODO: Make configurable
}
