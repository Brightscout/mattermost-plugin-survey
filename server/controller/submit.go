package controller

import (
	"encoding/json"
	"net/http"

	"github.com/rifflearning/mattermost-plugin-survey/server/config"
	"github.com/rifflearning/mattermost-plugin-survey/server/model"
	"github.com/rifflearning/mattermost-plugin-survey/server/platform"
)

var submitSurveyResponse = &Endpoint{
	Path:         "/submit",
	Method:       http.MethodPost,
	Execute:      executeSubmitSurveyResponse,
	RequiresAuth: true,
}

func executeSubmitSurveyResponse(w http.ResponseWriter, r *http.Request) {
	surveyPostID := r.URL.Query().Get("survey_post_id")

	response := &model.SurveyResponse{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(response); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		config.Mattermost.LogError("Failed to decode request body into survey response object.", "Error", err.Error())
		return
	}

	response.UserID = r.Header.Get(config.HeaderMattermostUserID)

	if err := platform.SubmitSurveyResponse(surveyPostID, response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		config.Mattermost.LogError("Failed to save survey responses.", "Error", err.Error())
		return
	}

	returnStatusOK(w)
}
