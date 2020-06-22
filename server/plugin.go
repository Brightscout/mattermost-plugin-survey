package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mattermost/mattermost-server/v5/plugin"
	"github.com/pkg/errors"

	"github.com/rifflearning/mattermost-plugin-survey/server/config"
	"github.com/rifflearning/mattermost-plugin-survey/server/controller"
	"github.com/rifflearning/mattermost-plugin-survey/server/platform"
	"github.com/rifflearning/mattermost-plugin-survey/server/platform/reminders"
	"github.com/rifflearning/mattermost-plugin-survey/server/store/kvstore"
)

type Plugin struct {
	plugin.MattermostPlugin

	handler http.Handler
}

func (p *Plugin) OnActivate() error {
	config.Mattermost = p.API
	config.Store = kvstore.NewStore()
	reminders.InitReminders()

	if err := p.setupStaticFileServer(); err != nil {
		config.Mattermost.LogError("Unable to setup static file server.", "Error", err.Error())
		return err
	}

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

func (p *Plugin) setupStaticFileServer() error {
	exe, err := os.Executable()
	if err != nil {
		return errors.Wrap(err, "couldn't find the plugin executable path")
	}
	p.handler = http.FileServer(http.Dir(filepath.Dir(exe) + config.ServerExeToWebappRootPath))
	return nil
}

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	conf := config.GetConfig()

	if err := conf.IsValid(); err != nil {
		config.Mattermost.LogError("This plugin is not configured: " + err.Error())
		http.Error(w, "This plugin is not configured.", http.StatusNotImplemented)
		return
	}

	endpoint := controller.GetEndpoint(r)
	if endpoint == nil {
		p.handler.ServeHTTP(w, r)
		return
	}

	if endpoint.RequiresAuth && !controller.Authenticated(w, r) {
		config.Mattermost.LogError(fmt.Sprintf("Endpoint: %s '%s' requires Authentication.", endpoint.Method, endpoint.Path))
		http.Error(w, "This endpoint requires authentication.", http.StatusForbidden)
		return
	}

	if err := endpoint.Execute(w, r); err != nil {
		config.Mattermost.LogError(fmt.Sprintf("Processing: %s '%s'.", r.Method, r.URL.String()), "Error", err.Error())
	}
}

func main() {
	plugin.ClientMain(&Plugin{})
}
