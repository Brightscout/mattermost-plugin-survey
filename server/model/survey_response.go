package model

import (
	"encoding/json"

	serverModel "github.com/mattermost/mattermost-server/v5/model"
)

// SurveyResponse records the users responses to the survey questions
type SurveyResponse struct {
	Type          string            `json:"type"`
	UserID        string            `json:"user_id"`
	MeetingID     string            `json:"meeting_id"`
	SurveyID      string            `json:"survey_id"`
	SurveyVersion int               `json:"survey_version"`
	CreatedAt     int64             `json:"created_at"`
	Responses     map[string]string `json:"responses"`
}

func (r *SurveyResponse) PreSave() *SurveyResponse {
	r.Type = TypeSurveyResponse
	r.CreatedAt = serverModel.GetMillis()
	return r
}

// EncodeToByte returns a survey response as a byte array
func (r *SurveyResponse) EncodeToByte() []byte {
	b, _ := json.Marshal(r)
	return b
}

// DecodeSurveyResponseFromByte tries to create a survey response from a byte array
func DecodeSurveyResponseFromByte(b []byte) *SurveyResponse {
	r := SurveyResponse{}
	err := json.Unmarshal(b, &r)
	if err != nil {
		return nil
	}
	return &r
}
