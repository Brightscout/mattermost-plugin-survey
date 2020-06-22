package controller

import (
	"net/http"

	serverModel "github.com/mattermost/mattermost-server/v5/model"
	"github.com/pkg/errors"

	"github.com/rifflearning/mattermost-plugin-survey/server/config"
)

var getDashboardPath = &Endpoint{
	Path:         "/dashboard",
	Method:       http.MethodGet,
	Execute:      executeGetDashboardPath,
	RequiresAuth: true,
}

func executeGetDashboardPath(w http.ResponseWriter, r *http.Request) error {
	conf := config.GetConfig()
	response := []byte(serverModel.MapToJson(map[string]string{"path": conf.DashboardPath}))

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(response); err != nil {
		return errors.Wrap(err, "failed to write data to HTTP response")
	}

	return nil
}
