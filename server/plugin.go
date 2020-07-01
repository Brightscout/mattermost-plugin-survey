package main

import (
	"net/http"

	"github.com/mattermost/mattermost-server/v5/plugin"

	"github.com/rifflearning/mattermost-plugin-survey/server/config"
	"github.com/rifflearning/mattermost-plugin-survey/server/controller"
	"github.com/rifflearning/mattermost-plugin-survey/server/platform"
	"github.com/rifflearning/mattermost-plugin-survey/server/platform/reminders"
	"github.com/rifflearning/mattermost-plugin-survey/server/store/kvstore"
)

type Plugin struct {
	plugin.MattermostPlugin
}

func (p *Plugin) OnActivate() error {
	config.Mattermost = p.API
	config.Helpers = p.Helpers
	config.Store = kvstore.NewStore()
	reminders.InitReminders()

	if err := p.OnConfigurationChange(); err != nil {
		return err
	}

	return nil
}

func (p *Plugin) OnDeactivate() error {
	reminders.StopReminders()
	return nil
}

func (p *Plugin) OnConfigurationChange() error {
	// If the plugin is not activated
	if config.Mattermost == nil {
		return nil
	}

	var configuration config.Configuration

	if err := config.Mattermost.LoadPluginConfiguration(&configuration); err != nil {
		config.Mattermost.LogError("Error in LoadPluginConfiguration: " + err.Error())
		return err
	}

	if err := configuration.ProcessConfiguration(); err != nil {
		config.Mattermost.LogError("Error in ProcessConfiguration: " + err.Error())
		return err
	}

	if err := configuration.IsValid(); err != nil {
		config.Mattermost.LogError("Error in Validating Configuration: " + err.Error())
		return err
	}

	config.SetConfig(&configuration)

	if err := initSurvey(); err != nil {
		config.Mattermost.LogError("Error in Initialising Survey: " + err.Error())
		return err
	}

	return nil
}

func initSurvey() error {
	survey := config.GetConfig().ParsedSurvey

	// TODO: for v2, determine ID for survey
	survey.ID = config.HardcodedSurveyID

	if err := platform.SaveSurvey(survey); err != nil {
		return err
	}
	return nil
}

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	p.API.LogDebug("New request:", "Host", r.Host, "RequestURI", r.RequestURI, "Method", r.Method, "URL", r.URL.String())

	conf := config.GetConfig()
	if err := conf.IsValid(); err != nil {
		p.API.LogError("This plugin is not configured.", "Error", err.Error())
		http.Error(w, "This plugin is not configured.", http.StatusNotImplemented)
		return
	}

	controller.InitAPI().ServeHTTP(w, r)
}

func main() {
	plugin.ClientMain(&Plugin{})
}
