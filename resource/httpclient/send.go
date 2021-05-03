package httpclient

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"tomato/resource"
)

const (
	OutputStatusCode   = "status_code"
	OutputResponseBody = "response_body"
)

func (h *HTTPClient) Send(ctx context.Context) error {
	method, ok := h.storage.Get(resource.GetExecID(ctx), "method")
	if !ok {
		return fmt.Errorf("missing argument method: %+v", h.storage)
	}

	url, ok := h.storage.Get(resource.GetExecID(ctx), "url")
	if !ok {
		return fmt.Errorf("missing argument url")
	} else {
		url = h.BaseURL + url
	}

	body, ok := h.storage.Get(resource.GetExecID(ctx), "body")
	if !ok {
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
	h.storage.Set(resource.GetExecID(ctx), "status_code", fmt.Sprintf("%d", resp.StatusCode))
	h.storage.Set(resource.GetExecID(ctx), "response_body", string(b))

	return nil
}
