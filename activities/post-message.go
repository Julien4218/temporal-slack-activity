package activities

import (
	"context"

	"github.com/Julien4218/temporal-slack-activity/instrumentation"
)

func PostMessageActivity(ctx context.Context, param string) (*string, error) {
	instrumentation.Log("PostMessageActivity")
	result := "pass"
	return &result, nil
}
