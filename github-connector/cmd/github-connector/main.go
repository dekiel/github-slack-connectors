package main

import (
	"net/http"

	"github.com/dekiel/github-slack-connectors/github-connector/pkg/events"
	"github.com/dekiel/github-slack-connectors/github-connector/pkg/github"
	"github.com/dekiel/github-slack-connectors/github-connector/pkg/handlers"

	log "github.com/sirupsen/logrus"
	"github.com/vrischmann/envconfig"
)

//Config containing all program configs
type Config struct {
	GitHubConnectorName string `envconfig:"GITHUB_CONNECTOR_NAME"`
	GitHubSecret        string `envconfig:"GITHUB_SECRET"`
	KymaEventsService   string `envconfig:"EVENTING_SERVICE"` //http://test-gh-connector-app-event-service.kyma-integration:8081/test-gh-connector-app/v1/events
	Port                string `envconfig:"EVENTING_PORT"`
}

func main() {
	var conf Config
	err := envconfig.Init(&conf)
	if err != nil {
		log.Fatal("Env error: ", err.Error())
	}
	log.Infof("Github secret: %s", conf.GitHubSecret)
	log.Infof("Eventing service URL: %s", conf.KymaEventsService)
	log.Infof("Port: %s", conf.Port)

	kyma := events.NewSender(&http.Client{}, events.NewValidator(), conf.KymaEventsService)
	webhook := handlers.NewWebHookHandler(
		github.NewReceivingEventsWrapper(conf.GitHubSecret),
		kyma,
	)

	http.HandleFunc("/webhook", webhook.HandleWebhook)
	log.Info(http.ListenAndServe(":"+conf.Port, nil))

	log.Info("Happy GitHub-Connecting!")

}
