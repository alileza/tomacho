package httpclient

import (
	"context"
	"testing"
)

func TestSend(t *testing.T) {
	ctx := context.WithValue(context.Background(), InputURL, "something")
	ctx = context.WithValue(ctx, InputMethod, "GET")
	ctx = context.WithValue(ctx, InputBody, "")
	ctx = context.WithValue(ctx, "exec_id", "123")

	httpcli := &HTTPClient{}
	if err := httpcli.Send(ctx); err != nil {
		t.Error(err)
	}
}
