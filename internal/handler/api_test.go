package handler

import (
	"geoip-service/internal/geoip"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCheckIPHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	resolver, err := geoip.NewResolver("../geoip/testdata/GeoLite2-Country-Test.mmdb")
	if err != nil {
		t.Fatalf("failed to load test mmdb: %v", err)
	}

	r := gin.New()
	r.POST("/check", CheckIPHandler(resolver))

	reqBody := `{"ip":"8.8.8.8","allowed_countries":["US"]}`
	req := httptest.NewRequest(http.MethodPost, "/check", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.Code)
	}
}
