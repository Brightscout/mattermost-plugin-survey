package model

import (
	"encoding/json"

	serverModel "github.com/mattermost/mattermost-server/v5/model"
)

// Survey stores a survey question
type SurveyQuestion struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Text string `json:"text"`
}

// Survey stores all needed information for a survey
type Survey struct {
	Type        string            `json:"type"`
	ID          string            `json:"id"`
	Version     int               `json:"version"`
	CreatedAt   int64             `json:"created_at"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Questions   []*SurveyQuestion `json:"questions"`
}

func (s *Survey) PreSave(version int) *Survey {
	s.Type = TypeSurvey
	s.Version = version
	s.CreatedAt = serverModel.GetMillis()
	for _, question := range s.Questions {
		question.ID = serverModel.NewId()
	}
	return s
}

// Equals checks the equality of deterministic fields of two surveys.
// Checks for the equality of Title, Description, Number of Questions and
// Type and Text for the Questions
func (s *Survey) Equals(survey *Survey) bool {
	if s.Title != survey.Title {
		return false
	}

	if s.Description != survey.Description {
		return false
	}

	if len(s.Questions) != len(survey.Questions) {
		return false
	}

	for i := range s.Questions {
		if s.Questions[i].Text != survey.Questions[i].Text {
			return false
		}
		if s.Questions[i].Type != survey.Questions[i].Type {
			return false
		}
	}

	return true
}

// EncodeToByte returns a survey as a byte array
func (s *Survey) EncodeToByte() []byte {
	b, _ := json.Marshal(s)
	return b
}

// DecodeSurveyFromByte tries to create a survey from a byte array
func DecodeSurveyFromByte(b []byte) *Survey {
	s := Survey{}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return nil
	}
	return &s
}
