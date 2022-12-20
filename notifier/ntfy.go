package notifier

import (
	"bytes"
	"net/http"
)

type NtfyNotifier struct {
	NotifierOptions NotifierOptions
}

func (ntfyNotifier NtfyNotifier) Notify(message string) {
	http.Post(ntfyNotifier.NotifierOptions.Endpoint, "text/plain", bytes.NewBuffer([]byte(message)))
}
