package external

import (
	"context"
	"net/http"
	"time"
)

type External struct {
	User interface {
		Profile(context.Context, string) (*Response, error)
	}
}

func NewExternal() External {
	return External{
		User: &UserExternal{
			httpClient: &http.Client{
				Timeout: time.Second * 10,
			},
		},
	}
}
