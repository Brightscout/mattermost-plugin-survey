package controller

import (
	"net/http"

	"github.com/rifflearning/mattermost-plugin-survey/server/config"
	"github.com/rifflearning/mattermost-plugin-survey/server/util"
)

type Endpoint struct {
	Path         string
	Method       string
	Execute      func(w http.ResponseWriter, r *http.Request) error
	RequiresAuth bool
}

// Endpoints is a map of endpoint key to endpoint object
// Usage: getEndpointKey(GetMetadata): GetMetadata
var Endpoints = map[string]*Endpoint{
	getEndpointKey(sendSurvey):           sendSurvey,
	getEndpointKey(getSurvey):            getSurvey,
	getEndpointKey(submitSurveyResponse): submitSurveyResponse,
}

func getEndpointKey(endpoint *Endpoint) string {
	return util.GetKeyHash(endpoint.Path + "-" + endpoint.Method)
}

// GetEndpoint returns an endpoint for an http request
func GetEndpoint(r *http.Request) *Endpoint {
	return Endpoints[util.GetKeyHash(r.URL.Path+"-"+r.Method)]
}

// Authenticated verifies if provided request is performed by a logged-in Mattermost user.
func Authenticated(w http.ResponseWriter, r *http.Request) bool {
	userID := r.Header.Get(config.HeaderMattermostUserID)
	if userID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return false
	}

	return true
}
