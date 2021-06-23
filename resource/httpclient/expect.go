package httpclient

import (
	"context"
	"fmt"
	"reflect"
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
		actualResponseBody, ok := h.storage.Get(resource.GetExecID(ctx), OutputResponseBody)
		if !ok {
			return fmt.Errorf("can't find response body, expecting to send request before expect")
		}

		if !reflect.DeepEqual(body, actualResponseBody) {
			return fmt.Errorf("unexpected response body\nactual: %sexpecting: %s", actualResponseBody, body)
		}

	} else {
		fmt.Printf("[DEBUG] skipping checking body\n")
	}

	return nil
}
