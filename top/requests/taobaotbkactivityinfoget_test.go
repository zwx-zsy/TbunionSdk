package requests_test

import (
	"fmt"
	"testing"

	"github.com/yancyzhou/TbunionSdk/top"
)

func TestTaobaoTbkActivityInfoGet(t *testing.T) {
	c := top.New()
	c.AppKey = top.AppKey
	c.AppSecret = top.AppSecret

	req := top.TaobaoTbkActivityInfoGetRequest()
	req.SetParam("adzone_id", "111563200399")
	req.SetParam("activity_material_id", "20150318020002597")
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
