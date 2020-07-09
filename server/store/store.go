package store

import (
	"github.com/rifflearning/mattermost-plugin-survey/server/model"
)

// SurveyStore allows to access surveys with some kind of store.
type SurveyStore interface {
	GetLatestSurveyInfo(id string) (*model.LatestSurveyInfo, error)
	SaveLatestSurveyInfo(info *model.LatestSurveyInfo) error
	GetSurvey(id string, version int) (*model.Survey, error)
	SaveSurvey(survey *model.Survey) error
	GetMeetingMetadata(meetingID string) (*model.MeetingMetadata, error)
	SaveMeetingMetadata(data *model.MeetingMetadata) error
	GetUserMeetingMetadata(userID, meetingID string) (*model.UserMeetingMetadata, error)
	SaveUserMeetingMetadata(data *model.UserMeetingMetadata) error
	GetSurveyResponse(userID, meetingID, surveyID string, surveyVersion int) (*model.SurveyResponse, error)
	SaveSurveyResponse(response *model.SurveyResponse) error
	GetReminderMetadata(postID string) (*model.ReminderMetadata, error)
	SaveReminderMetadata(data *model.ReminderMetadata) error
	GetRemindersList() ([]string, error)
	SaveRemindersList(list []string) error
}
