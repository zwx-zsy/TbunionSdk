package requests_test

import (
	"fmt"
	"testing"

	"github.com/yancyzhou/TbunionSdk/top"
)

func TestTaobaoTopAuthTokenCreate(t *testing.T) {
	c := top.New()
	c.AppKey = top.AppKey
	c.AppSecret = top.AppSecret

	req := top.TaobaoTopAuthTokenCreateRequest()
	req.SetParam("code", "RHc7jlrZsm9OKFoDIts84Oa27388846")
	body, err := c.Exec(req)
	if err != nil {
		t.Error(err)
	}

	result, err := req.Result(body)
	if err != nil {
		responseError, _ := top.ErrorResponse(body)
		t.Error(responseError.ErrorResponse.Code, responseError.ErrorResponse.Msg)
	} else {
		fmt.Println(result)
	}
}
