package health

import (
	"net/http"
	"qtkstorebe/pkg/router"
)

// GetIndex Function to Show API Information and Health
func GetIndex(w http.ResponseWriter, r *http.Request) {
	router.HealthCheck(w)
}
