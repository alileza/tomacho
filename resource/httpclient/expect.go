package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"tomato/resource"
)

const (
	InputCode = "code"
)

func (h *HTTPClient) Expect(ctx context.Context) error {
	code, ok := h.storage.Get(resource.GetExecID(ctx), InputCode)
	if ok {
		actualStatusCode, ok := h.storage.Get(resource.GetExecID(ctx), OutputStatusCode)
		if !ok {
			return fmt.Errorf("can't find status code, expecting to send request before expect")
		}
		if code != actualStatusCode {
			return fmt.Errorf("unexpected status code of %s, expecting %s", actualStatusCode, code)
		}

	} else {
		fmt.Printf("[DEBUG] skipping checking code\n")
	}

	body, ok := h.storage.Get(resource.GetExecID(ctx), InputBody)
	if ok {
		responseHeader, ok := h.storage.Get(resource.GetExecID(ctx), OutputResponseHeader)
		if !ok {
			return fmt.Errorf("response header not found")
		}

		headers := make(http.Header)
		if err := json.Unmarshal([]byte(responseHeader), &headers); err != nil {
			return fmt.Errorf("failed to unmarshal response header body: %w", err)
		}

		actualResponseBody, ok := h.storage.Get(resource.GetExecID(ctx), OutputResponseBody)
		if !ok {
			return fmt.Errorf("can't find response body, expecting to send request before expect")
		}

		if !reflect.DeepEqual(strings.TrimSpace(body), strings.TrimSpace(actualResponseBody)) {
			return fmt.Errorf("unexpected response body\nactual: %sexpecting: %s", actualResponseBody, body)
		}

	} else {
		fmt.Printf("[DEBUG] skipping checking body\n")
	}

	return nil
}
