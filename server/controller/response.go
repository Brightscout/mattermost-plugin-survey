package controller

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rifflearning/mattermost-plugin-survey/server/config"
	"github.com/rifflearning/mattermost-plugin-survey/server/platform"
)

var getSurveyResponse = &Endpoint{
	Path:         "/meetings/{meetingID:[A-Za-z0-9]+}/response",
	Method:       http.MethodGet,
	Execute:      executeGetResponse,
	RequiresAuth: true,
}

// Get survey by meetingID or by surveyID and surveyVersion
func executeGetResponse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	meetingID := params["meetingID"]
	userID := r.Header.Get(config.HeaderMattermostUserID)

	userMeetingMetadata := platform.GetUserMeetingMetadata(userID, meetingID)
	if userMeetingMetadata == nil || userMeetingMetadata.RespondedAt == 0 {
		config.Mattermost.LogError("User has not responded to the survey yet.", "userID", userID, "meetingID", meetingID)
		http.Error(w, "User has not responded to the survey yet.", http.StatusNotFound)
		return
	}

	response := platform.GetSurveyResponse(userID, meetingID)
	if response == nil {
		config.Mattermost.LogError("Failed to get survey response.", "userID", userID, "meetingID", meetingID)
		http.Error(w, "Failed to get survey response.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(response.EncodeToByte())
}
