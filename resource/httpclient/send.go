package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"tomato/resource"
)

const (
	InputMethod = "method"
	InputURL    = "url"
	InputBody   = "body"

	OutputStatusCode     = "status_code"
	OutputResponseBody   = "response_body"
	OutputResponseHeader = "response_headers"
)

func (h *HTTPClient) Send(ctx context.Context) error {
	method, ok := h.storage.Get(resource.GetExecID(ctx), InputMethod)
	if !ok {
		return fmt.Errorf("missing argument method: %+v", h.storage)
	}

	url, ok := h.storage.Get(resource.GetExecID(ctx), InputURL)
	if !ok {
		return fmt.Errorf("missing argument url")
	} else {
		url = h.BaseURL + url
	}

	body, ok := h.storage.Get(resource.GetExecID(ctx), InputBody)
	if !ok {
		log.WithField("resource", "httpclient").Debug("body is not being passed")
		fmt.Printf("[DEBUG] body is not being passed\n")
	}

	req, err := http.NewRequestWithContext(ctx, strings.ToUpper(method), url, bytes.NewBufferString(body))
	if err != nil {
		return err
	}

	resp, err := h.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	respHeader, err := json.Marshal(resp.Header)
	if err != nil {
		return fmt.Errorf("failed to json marshal response header: %w", err)
	}

	h.storage.Set(resource.GetExecID(ctx), OutputStatusCode, fmt.Sprintf("%d", resp.StatusCode))
	h.storage.Set(resource.GetExecID(ctx), OutputResponseBody, string(b))
	h.storage.Set(resource.GetExecID(ctx), OutputResponseHeader, string(respHeader))

	return nil
}
