package platform

import (
	serverModel "github.com/mattermost/mattermost-server/v5/model"
	"github.com/pkg/errors"
	"github.com/rifflearning/mattermost-plugin-survey/server/config"
	"github.com/rifflearning/mattermost-plugin-survey/server/model"
)

// GetUserResponses returns the user response to the survey for a meeting with a given userID and meetingID.
// Returns the SurveyResponse if found and nil if not.
func GetSurveyResponse(userID, meetingID string) *model.SurveyResponse {
	surveyID, surveyVersion, err := GetSurveyInfoForMeeting(meetingID)
	if err != nil {
		config.Mattermost.LogError("Unable to get meeting metadata.", "Error", err.Error())
		return nil
	}

	response, err := config.Store.GetSurveyResponse(userID, meetingID, surveyID, surveyVersion)
	if err != nil {
		config.Mattermost.LogError("Unable to get existing survey.", "Error", err.Error())
		return nil
	}
	return response
}

// SubmitSurveyResponse saves the survey response to the DB.
// Updating the survey responses is allowed.
func SubmitSurveyResponse(surveyPostID string, response *model.SurveyResponse) error {
	userMeetingMetadata := GetUserMeetingMetadata(response.UserID, response.MeetingID)
	// create user-meeting-metadata if survey was not sent to user
	if userMeetingMetadata == nil {
		userMeetingMetadata = &model.UserMeetingMetadata{
			MeetingID:    response.MeetingID,
			UserID:       response.UserID,
			SurveySentAt: serverModel.GetMillis(),
		}
	}

	response = response.PreSave()
	if err := config.Store.SaveSurveyResponse(response); err != nil {
		config.Mattermost.LogError("Failed to save the survey response.", "Error", err.Error())
		return err
	}

	if surveyPostID != "" {
		surveyPost, appErr := config.Mattermost.GetPost(surveyPostID)
		if appErr != nil {
			config.Mattermost.LogError("Failed to get the survey post.", "Error", appErr.Error())
			return errors.New(appErr.Error())
		}

		surveyPost.AddProp(config.PropSurveySubmitted, true)
		if _, appErr := config.Mattermost.UpdatePost(surveyPost); appErr != nil {
			config.Mattermost.LogError("Failed to update the survey post.", "Error", appErr.Error())
			return errors.New(appErr.Error())
		}
	}

	userMeetingMetadata.RespondedAt = response.CreatedAt
	if err := SaveUserMeetingMetadata(userMeetingMetadata); err != nil {
		return err
	}

	return nil
}
