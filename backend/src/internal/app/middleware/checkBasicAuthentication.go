package middleware

import (
	"github.com/yukiHaga/web_server/src/internal/app/controller"
	"github.com/yukiHaga/web_server/src/pkg/henagin/http"
)

type CheckBasicAuthenticationController struct {
	nextAction func(request *http.Request) *http.Response
}

func (c CheckBasicAuthenticationController) Action(request *http.Request) *http.Response {
	if request.CheckBasicAuthentication() {
		// 次に渡す
		return c.nextAction(request)
	} else {
		// Authorizationヘッダーがないか、あるけどvalueが間違っている場合
		// 再度Unauthorizedをを返す
		response := http.NewResponse(
			http.VersionsFor11,
			http.StatusUnauthorizedCode,
			http.StatusReasonUnauthorized,
			request.TargetPath,
			[]byte{},
		)

		// WWW-Authenticateヘッダーをセットする
		response.SetBasicAuthenticationHeader()
		return response
	}
}

// ミドルウェア
// / ダミーコントローラとダミーアクションを作って、ダミーアクションの中で元々のコントローラのアクションを呼び出して、最終的にレスポンスを返せばOK
// goではメソッドを書き換えるのはできなかった
func CheckBasicAuthentication(c controller.Controller) controller.Controller {
	return CheckBasicAuthenticationController{nextAction: c.Action}
}
