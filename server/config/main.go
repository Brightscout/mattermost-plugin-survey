package config

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	serverModel "github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"github.com/pkg/errors"
	"go.uber.org/atomic"

	"github.com/rifflearning/mattermost-plugin-survey/server/model"
	"github.com/rifflearning/mattermost-plugin-survey/server/store"
)

const (
	URLPluginBase = "/plugins/" + PluginName
	URLStaticBase = URLPluginBase + "/static"

	HeaderMattermostUserID = "Mattermost-User-Id"

	BotUsername    = "riffbot"
	BotDisplayName = "Riff Bot"
	BotIconURL     = URLStaticBase + "/riffbot.png"
	botIconPath    = "assets/riffbot.png"

	HardcodedSurveyID = "f298903f8a80054ba09e342d0d9780635d3675a2"

	PropSurveySubmitted = "submitted"

	ReminderTickerDuration = 20 * time.Second
)

var (
	config     atomic.Value
	Mattermost plugin.API
	Helpers    plugin.Helpers
	Store      store.SurveyStore
)

type Configuration struct {
	Survey           string `json:"Survey"`
	ReminderText     string `json:"ReminderText"`
	MaxReminderCount string `json:"MaxReminderCount"`
	ReminderInterval string `json:"ReminderInterval"`

	// Derived Attributes
	BotUserID                string
	ParsedSurvey             *model.Survey
	MaxReminderCountInt      int
	ReminderIntervalDuration time.Duration
}

func GetConfig() *Configuration {
	return config.Load().(*Configuration)
}

func SetConfig(c *Configuration) {
	config.Store(c)
}

func (c *Configuration) ProcessConfiguration() error {
	// Derive BotUserID
	bot := &serverModel.Bot{
		Username:    BotUsername,
		DisplayName: BotDisplayName,
	}
	botID, ensureBotError := Helpers.EnsureBot(bot, plugin.ProfileImagePath(botIconPath))
	if ensureBotError != nil {
		return errors.Wrap(ensureBotError, "failed to ensure riff bot")
	}
	c.BotUserID = botID

	// Derive ParsedSurvey
	var parsedSurvey *model.Survey
	if err := json.Unmarshal([]byte(c.Survey), &parsedSurvey); err != nil {
		return errors.Wrap(err, "Unable to parse json for the Survey. Please make sure it is a valid JSON of the provided format. Error")
	}
	c.ParsedSurvey = parsedSurvey

	// Process ReminderText
	c.ReminderText = strings.TrimSpace(c.ReminderText)

	// Derive MaxReminderCountInt
	maxReminderCountInt, conversionErr := strconv.Atoi(c.MaxReminderCount)
	if conversionErr != nil {
		return errors.Wrap(conversionErr, "MaxReminderCount is not a valid number")
	}
	c.MaxReminderCountInt = maxReminderCountInt

	// Derive ReminderIntervalInt
	reminderIntervalInt, conversionErr := strconv.Atoi(c.ReminderInterval)
	if conversionErr != nil {
		return errors.Wrap(conversionErr, "ReminderInterval is not a valid number")
	}
	c.ReminderIntervalDuration = time.Duration(reminderIntervalInt) * time.Minute

	return nil
}

func (c *Configuration) IsValid() error {
	if c.ReminderText == "" {
		return errors.New("Reminder text cannot be empty")
	}

	if c.MaxReminderCountInt < 0 {
		return errors.New("Max reminder count cannot negative")
	}

	if c.ReminderIntervalDuration <= 0 {
		return errors.New("Reminder interval must be greater than zero")
	}

	return nil
}
