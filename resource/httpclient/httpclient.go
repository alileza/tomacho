package httpclient

import (
	"context"
	"fmt"
	"net/http"
	"tomato/resource"
	"tomato/storage"
)

type HTTPClient struct {
	// shared storages
	storage *storage.Storage

	BaseURL string
	Client  *http.Client

	headers http.Header
}

func NewHTTPClient(options resource.Options) *HTTPClient {
	var h HTTPClient

	h.Client = &http.Client{}
	h.headers = make(http.Header)
	h.storage = storage.New()

	if url, ok := options["base_url"]; ok {
		h.BaseURL = url
	}

	return &h
}

func (h *HTTPClient) Status() error {
	if h.BaseURL == "" {
		return nil
	}

	resp, err := h.Client.Get(h.BaseURL)
	if err != nil {
		return fmt.Errorf("Failed to send GET request to %s: %w", h.BaseURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("Unexpected status code from %s: status_code=%d", h.BaseURL, resp.StatusCode)
	}

	return nil
}

func (h *HTTPClient) Exec(ctx context.Context, command string, args resource.Arguments) error {
	for key, value := range args {
		h.storage.Set(resource.GetExecID(ctx), key, value)
	}

	switch command {
	case "send":
		return h.Send(ctx)
	case "set-header":
		return h.SetHeader(ctx)
	case "expect":
		return h.Expect(ctx)
	default:
		return fmt.Errorf("Unexpected command: %s", command)
	}
}
