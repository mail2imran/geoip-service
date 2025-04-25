package handler

import (
	"geoip-service/internal/geoip"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CheckRequest struct {
	IP               string   `json:"ip"`
	AllowedCountries []string `json:"allowed_countries"`
}

type CheckResponse struct {
	Allowed bool   `json:"allowed"`
	Country string `json:"country"`
}

func CheckIPHandler(resolver *geoip.Resolver) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CheckRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		country, err := resolver.Country(req.IP)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Geo lookup failed"})
			return
		}

		allowed := false
		for _, ac := range req.AllowedCountries {
			if ac == country {
				allowed = true
				break
			}
		}

		c.JSON(http.StatusOK, CheckResponse{Allowed: allowed, Country: country})
	}
}
