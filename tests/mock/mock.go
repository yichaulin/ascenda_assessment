package mock

import (
	"github.com/jarcoal/httpmock"

	"ascenda_assessment/apis/resty"
)

func MockAPI(method string, url string, resBody string, resStatus int) {
	httpmock.RegisterResponder(method, url, httpmock.NewStringResponder(resStatus, resBody))
	httpmock.ActivateNonDefault(resty.GetHTTPClient())
}
