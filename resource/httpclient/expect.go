package httpclient

import (
	"context"
	"fmt"
	"tomato/resource"

	"tomato/log"
)

func (h *HTTPClient) Expect(ctx context.Context) error {
	code, ok := h.Storage.Get(resource.GetExecID(ctx), "code")
	if ok {
		actualStatusCode, ok := h.Storage.Get(resource.GetExecID(ctx), OutputStatusCode)
		if !ok {
			return fmt.Errorf("can't find status code, expecting to send request before expect.")
		}
		if code != actualStatusCode {
			return fmt.Errorf("Unexpected status code of %s, expecting %s", actualStatusCode, code)
		}

	} else {
		log.Debug("skipping checking code")
	}

	return nil
}
