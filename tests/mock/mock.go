package mock

import (
	"github.com/jarcoal/httpmock"

	"ascenda_assessment/apis/client"
)

func MockAPI(client client.Client, method string, url string, resBody string, resStatus int) {
	httpmock.RegisterResponder(method, url, httpmock.NewStringResponder(resStatus, resBody))
	httpmock.ActivateNonDefault(client.GetHTTPClient())
}
