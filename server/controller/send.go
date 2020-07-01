package controller

import (
	"net/http"

	"github.com/rifflearning/mattermost-plugin-survey/server/config"
	"github.com/rifflearning/mattermost-plugin-survey/server/platform"
)

var sendSurvey = &Endpoint{
	Path:         "/send",
	Method:       http.MethodGet,
	Execute:      executeSendSurvey,
	RequiresAuth: true,
}

func executeSendSurvey(w http.ResponseWriter, r *http.Request) {
	meetingID := r.URL.Query().Get("meeting_id")
	userID := r.Header.Get(config.HeaderMattermostUserID) // can only request a survey for self

	config.Mattermost.LogDebug("Send survey executed.", "userID", userID, "meetingID", meetingID)

	if err := platform.SendSurveyPost(userID, meetingID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		config.Mattermost.LogError("Failed to send survey to the user.", "userID", userID, "meetingID", meetingID, "Error", err.Error())
		return
	}
	returnStatusOK(w)
}
