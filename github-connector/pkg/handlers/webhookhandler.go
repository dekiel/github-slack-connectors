package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dekiel/github-slack-connectors/github-connector/pkg/apperrors"
	git "github.com/dekiel/github-slack-connectors/github-connector/pkg/github"
	"github.com/dekiel/github-slack-connectors/github-connector/pkg/httperrors"
	"github.com/google/go-github/github"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

//Sender is an interface used to allow mocking sending events to Kyma's event bus
type Sender interface {
	SendToKyma(eventType, sourceID, eventTypeVersion, eventID string, data json.RawMessage) apperrors.AppError
}

//WebHookHandler is a struct used to allow mocking the github library methods
type WebHookHandler struct {
	validator git.Validator
	sender    Sender
}

//NewWebHookHandler creates a new webhook handler with the passed interface
func NewWebHookHandler(v git.Validator, s Sender) *WebHookHandler {
	return &WebHookHandler{validator: v, sender: s}
}

//HandleWebhook is a function that handles the /webhook endpoint.
func (wh *WebHookHandler) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	payload, apperr := wh.validator.ValidatePayload(r, []byte(wh.validator.GetToken()))

	if apperr != nil {
		apperr = apperr.Append("While handling '/webhook' endpoint")

		log.Warn(apperr.Error())
		httperrors.SendErrorResponse(apperr, w)
		return
	}

	event, apperr := wh.validator.ParseWebHook(github.WebHookType(r), payload)
	if apperr != nil {
		apperr = apperr.Append("While handling '/webhook' endpoint")

		log.Warn(apperr.Error())
		httperrors.SendErrorResponse(apperr, w)
		return
	}

	var eventType string
	switch event := event.(type) {
	// Supported github events
	case *github.IssuesEvent:
		eventType = fmt.Sprintf("issuesevent.%s", *event.Action)
	}
	sourceID := os.Getenv("GITHUB_CONNECTOR_NAME")
	log.Info(fmt.Sprintf("Event type '%s' received.", eventType))
	apperr = wh.sender.SendToKyma(eventType, sourceID, "v1", "", payload)

	if apperr != nil {
		log.Info(apperrors.Internal("While handling the event: %s", apperr.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
