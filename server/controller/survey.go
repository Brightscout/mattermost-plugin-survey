package controller

import (
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	"github.com/rifflearning/mattermost-plugin-survey/server/config"
	"github.com/rifflearning/mattermost-plugin-survey/server/platform"
)

var getSurvey = &Endpoint{
	Path:         "/survey",
	Method:       http.MethodGet,
	Execute:      executeGetSurvey,
	RequiresAuth: true,
}

// Get survey by meetingID or by surveyID and surveyVersion
func executeGetSurvey(w http.ResponseWriter, r *http.Request) error {
	meetingID := r.URL.Query().Get("meeting_id")
	surveyID := r.URL.Query().Get("survey_id")
	surveyVersion := r.URL.Query().Get("survey_version")

	if surveyID == "" && meetingID == "" {
		http.Error(w, "Please provide either the meetingID or the survey id and version.", http.StatusBadRequest)
		return errors.New("surveyID or meetingID not provided")
	}

	var surveyVersionInt int
	if meetingID != "" {
		// Get the survey ID and Version from the `meetingID`
		var surveyInfoErr error
		surveyID, surveyVersionInt, surveyInfoErr = platform.GetSurveyInfoForMeeting(meetingID)
		if surveyInfoErr != nil {
			config.Mattermost.LogError("Failed to get survey. Unable to get survey info for the meeting.", "MeetingID", meetingID, "Error", surveyInfoErr.Error())
			http.Error(w, "Unable to get survey info for the meeting.", http.StatusInternalServerError)
			return surveyInfoErr
		}
	} else {
		version, err := strconv.Atoi(surveyVersion)
		if err != nil {
			http.Error(w, "Invalid Survey Version.", http.StatusBadRequest)
			return errors.Wrap(err, "invalid survey version")
		}
		surveyVersionInt = version
	}

	survey := platform.GetSurvey(surveyID, surveyVersionInt)
	if survey == nil {
		http.Error(w, "Please check the meetingID or the survey id and version and try again later.", http.StatusBadRequest)
		return errors.New("unable to get the requested survey")
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(survey.EncodeToByte()); err != nil {
		return errors.Wrap(err, "failed to write data to HTTP response")
	}

	return nil
}
