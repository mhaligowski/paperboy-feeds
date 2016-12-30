package feeds

import (
	"net/http"
	"golang.org/x/net/context"
)

type getContext func(*http.Request) context.Context
