package zeropsSdk

import (
	"net/http"

	"github.com/zeropsio/zerops-go/sdk"
	"github.com/zeropsio/zerops-go/sdkBase"
)

func New(config sdkBase.Config, client *http.Client) *Handler {
	return &Handler{
		zdk: sdk.New(config, client),
	}
}

type Handler struct {
	zdk sdk.Handler
}

func (h *Handler) Authorized(token string) sdk.Handler {
	authorizedSdk := sdk.AuthorizeSdk(h.zdk, token)
	return authorizedSdk
}
