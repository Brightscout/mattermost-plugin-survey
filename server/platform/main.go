package platform

import (
	serverModel "github.com/mattermost/mattermost-server/v5/model"
	"github.com/pkg/errors"
	"github.com/rifflearning/mattermost-plugin-survey/server/config"
	"github.com/rifflearning/mattermost-plugin-survey/server/model"
	"github.com/rifflearning/mattermost-plugin-survey/server/platform/reminders"
)

// GetSurvey returns the survey with a given id and version.
// Returns the survey if found and nil if not.
func GetSurvey(id string, version int) *model.Survey {
	survey, err := config.Store.GetSurvey(id, version)
	if err != nil {
		config.Mattermost.LogError("Unable to get existing survey.", "Error", err.Error())
		return nil
	}
	return survey
}

// SaveSurvey creates a new survey in the DB.
// If this is the first survey for a given ID, the version is set to 1.
// Otherwise, the version for an existing survey is incremented.
func SaveSurvey(survey *model.Survey) error {
	currentSurveyVersion := 0

	if latestSurveyInfo := GetLatestSurveyInfo(survey.ID); latestSurveyInfo != nil {
		currentSurveyVersion = latestSurveyInfo.SurveyVersion

		// Check for existing survey in DB
		if s := GetSurvey(survey.ID, currentSurveyVersion); s != nil && s.Equals(survey) {
			config.Mattermost.LogInfo("Survey already exists and is the same as the current survey. New version not created.", "SurveyID", s.ID, "SurveyVersion", s.Version)
			return nil
		}
	}

	// Create the first survey or new version of an existing survey
	survey = survey.PreSave(currentSurveyVersion + 1)
	if err := config.Store.SaveSurvey(survey); err != nil {
		config.Mattermost.LogError("Failed to save survey.", "Error", err.Error())
		return err
	}

	if err := SaveLatestSurveyInfo(survey.ID, survey.Version); err != nil {
		config.Mattermost.LogError("Survey saved successfully but latest survey information not updated.", "Error", err.Error())
		return err
	}

	return nil
}

// GetSurvey returns the latest survey information for the survey with a given id.
// Returns the info if found and nil if not.
func GetLatestSurveyInfo(id string) *model.LatestSurveyInfo {
	latestSurveyInfo, err := config.Store.GetLatestSurveyInfo(id)
	if err != nil {
		config.Mattermost.LogError("Unable to get latest survey information.", "Error", err.Error())
		return nil
	}
	return latestSurveyInfo
}

// SaveLatestSurveyInfo saves the latest survey information for a survey with a given id and version.
func SaveLatestSurveyInfo(id string, version int) error {
	info := &model.LatestSurveyInfo{
		SurveyID:      id,
		SurveyVersion: version,
	}
	info = info.PreSave()
	if err := config.Store.SaveLatestSurveyInfo(info); err != nil {
		config.Mattermost.LogError("Failed to save latest survey information.", "Error", err.Error())
	}
	return nil
}

// SubmitSurveyResponse saves the survey response to the DB.
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

	if userMeetingMetadata.RespondedAt != 0 {
		config.Mattermost.LogError("User has already responded to this survey. New response not recorded.", "UserID", response.UserID, "MeetingID", response.MeetingID, "Response", string(response.EncodeToByte()))
		return errors.New("unable to record user response: response already exists")
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

// GetSurveyInfoForMeeting is called to select the survey for a meeting
func GetSurveyInfoForMeeting(meetingID string) (string, int, error) {
	if meetingMetadata := GetMeetingMetadata(meetingID); meetingMetadata != nil {
		return meetingMetadata.SurveyID, meetingMetadata.SurveyVersion, nil
	}

	// TODO: for v2, select surveyID based on some defined criteria
	surveyID := config.HardcodedSurveyID
	info := GetLatestSurveyInfo(surveyID)
	if info == nil {
		return "", 0, errors.New("survey does not exist")
	}

	if err := SaveMeetingMetadata(meetingID, info.SurveyID, info.SurveyVersion); err != nil {
		return "", 0, errors.Wrap(err, "failed to save meeting metadata")
	}

	return info.SurveyID, info.SurveyVersion, nil
}

// SendSurveyPost creates the survey post for a user who participated in a meeting.
func SendSurveyPost(userID, meetingID string) error {
	if userMeetingMetadata := GetUserMeetingMetadata(userID, meetingID); userMeetingMetadata != nil && userMeetingMetadata.SurveySentAt != 0 {
		config.Mattermost.LogInfo("Survey already sent to the user. New survey post not created.", "UserID", userID, "MeetingID", meetingID)
		return nil
	}

	conf := config.GetConfig()
	channel, appErr := config.Mattermost.GetDirectChannel(conf.BotUserID, userID)
	if appErr != nil {
		return errors.Wrap(appErr, "Unable to create DM Channel.")
	}

	surveyID, surveyVersion, surveyInfoErr := GetSurveyInfoForMeeting(meetingID)
	if surveyInfoErr != nil {
		config.Mattermost.LogError("Failed to send survey to the user. Unable to get survey info for the meeting.", "UserID", userID, "MeetingID", meetingID, "Error", surveyInfoErr.Error())
		return surveyInfoErr
	}

	post := &serverModel.Post{
		UserId:    conf.BotUserID,
		ChannelId: channel.Id,
		Type:      "custom_survey",
		Message:   "Survey",
		Props: serverModel.StringInterface{
			"from_webhook":      "true",
			"override_username": config.BotDisplayName,
			"override_icon_url": config.BotIconURL,
			"meeting_id":        meetingID,
			"survey_id":         surveyID,
			"survey_version":    surveyVersion,
		},
	}

	post, createPostErr := config.Mattermost.CreatePost(post)
	if createPostErr != nil {
		return errors.Wrap(createPostErr, "failed to create survey post for the channel: "+channel.Id)
	}
	go reminders.AddNew(post.Id, channel.Id, userID, meetingID, post.CreateAt)

	userMeetingMetadata := &model.UserMeetingMetadata{
		MeetingID:    meetingID,
		UserID:       userID,
		SurveySentAt: post.CreateAt,
	}
	if err := SaveUserMeetingMetadata(userMeetingMetadata); err != nil {
		return err
	}

	return nil
}
