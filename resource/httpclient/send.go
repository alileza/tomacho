package httpclient

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strings"
	"tomato/log"
	"tomato/resource"
)

const (
	OutputStatusCode   = "status_code"
	OutputResponseBody = "response_body"
)

func (h *HTTPClient) Send(ctx context.Context) error {
	method, ok := h.Storage.Get(resource.GetExecID(ctx), "method")
	if !ok {
		return fmt.Errorf("missing argument method: %+v", h.Storage)
	}

	url, ok := h.Storage.Get(resource.GetExecID(ctx), "url")
	if !ok {
		return fmt.Errorf("missing argument url")
	} else {
		url = h.BaseURL + url
	}

	body, ok := h.Storage.Get(resource.GetExecID(ctx), "body")
	if !ok {
		log.Info(ctx, "body is not being passed")
	}

	headers, ok := h.Storage.GetMap(resource.GetExecID(ctx), "headers")
	if !ok {
		log.Info(ctx, "body is not being passed")
	}

	req, err := http.NewRequestWithContext(ctx, strings.ToUpper(method), url, bytes.NewBufferString(body))
	if err != nil {
		return err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	dumpreq, err := httputil.DumpRequest(req, true)
	if err != nil {
		dumpreq = []byte(err.Error())
	}
	h.Storage.Set(resource.GetExecID(ctx), "request", string(dumpreq))

	resp, err := h.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	dumpresp, err := httputil.DumpResponse(resp, true)
	if err != nil {
		dumpresp = []byte(err.Error())
	}
	h.Storage.Set(resource.GetExecID(ctx), "response", string(dumpresp))

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	h.Storage.Set(resource.GetExecID(ctx), "status_code", fmt.Sprintf("%d", resp.StatusCode))
	h.Storage.Set(resource.GetExecID(ctx), "response_body", string(b))

	return nil
}
