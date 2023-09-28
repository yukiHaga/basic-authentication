package controller

import (
	"github.com/yukiHaga/web_server/src/pkg/henagin/http"
	"github.com/yukiHaga/web_server/src/pkg/henagin/view"
)

type Sample struct{}

func NewSample() *Sample {
	return &Sample{}
}

func (c *Sample) Action(request *http.Request) *http.Response {
	// ミドルウェアにおけるbasic認証を通過したので、認証が必要なページを見れる
	// サブディレクトリなので、認証フォームは表示されないはず
	body := view.Render("sample.html")
	return http.NewResponse(
		http.VersionsFor11,
		http.StatusSuccessCode,
		http.StatusReasonOk,
		request.TargetPath,
		body,
	)
}
