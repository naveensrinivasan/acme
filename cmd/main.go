package main

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/naveensrinivasan/acme/internal/cache"
	"github.com/naveensrinivasan/acme/internal/certs"
)

const acme = "acme"

func main() {
	r := gin.Default()

	// local cache simulating something like redis
	cache := cache.New()

	// manages the certificates
	certs := certs.New(cache)

	// generate a cert for the acme server
	// * generates its own certificate and keeps its certificate up-to-date when it expires
	certs.Get(acme)

	r.GET("/cert/:domain", func(c *gin.Context) {
		domain := strings.ToLower(c.Param("domain"))
		// acme cannot be used because
		// it is assumed that it is the name for this to service that would use the subdomain
		// to address the below requirement
		// * generates its own certificate and keeps its certificate up-to-date when it expires
		if domain == acme {
			c.JSON(400, gin.H{
				"error": "domain cannot be acme.",
			})
			log.Println("acme domain was requested.")
			return
		}
		certificate := certs.Get(domain)
		c.JSON(200, gin.H{
			"domain":      domain,
			"certificate": certificate,
		})
	})
	r.Run()
}
