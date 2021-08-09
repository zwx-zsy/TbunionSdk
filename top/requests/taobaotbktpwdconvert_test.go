package requests

import (
	"fmt"
	"github.com/yancyzhou/TbunionSdk/top"
	"testing"
)

func TestTaobaoTbkTpwdConvert(t *testing.T) {
	c := top.New()
	c.AppKey = top.AppKey
	c.AppSecret = top.AppSecret

	req := top.TaobaoTbkTpwdConvertRequest()
	req.SetParam("password_content", "0/ej31XiOw87M/:/")
	req.SetParam("adzone_id", "111470550243")
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
