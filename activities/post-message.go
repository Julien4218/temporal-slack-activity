package activities

import (
	"context"
)

func PostMessageActivity(ctx context.Context, param string) (*string, error) {
	result := "pass"
	return &result, nil
}
