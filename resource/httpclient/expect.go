package httpclient

import (
	"context"
	"fmt"
	"tomato/resource"
)

func (h *HTTPClient) Expect(ctx context.Context) error {
	code, ok := h.storage.Get(resource.GetExecID(ctx), "code")
	if ok {
		actualStatusCode, ok := h.storage.Get(resource.GetExecID(ctx), OutputStatusCode)
		if !ok {
			return fmt.Errorf("can't find status code, expecting to send request before expect.")
		}
		if code != actualStatusCode {
			return fmt.Errorf("Unexpected status code of %s, expecting %s", actualStatusCode, code)
		}

	} else {
		fmt.Printf("[DEBUG] skipping checking code\n")
	}

	return nil
}
