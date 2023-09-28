package controller

import (
	"github.com/yukiHaga/web_server/src/pkg/henagin/http"
	"github.com/yukiHaga/web_server/src/pkg/henagin/view"
)

type BasicAuthentication struct{}

func NewBasicAuthentication() *BasicAuthentication {
	return &BasicAuthentication{}
}

func (c *BasicAuthentication) Action(request *http.Request) *http.Response {
	// ミドルウェアにおけるbasic認証を通過したので、認証が必要なページを見れる
	body := view.Render("basic_authentication.html")
	return http.NewResponse(
		http.VersionsFor11,
		http.StatusSuccessCode,
		http.StatusReasonOk,
		request.TargetPath,
		body,
	)
}
