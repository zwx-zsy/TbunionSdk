package top_test

import (
	"fmt"
	"github.com/yancyzhou/TbunionSdk/top"
	"testing"
)

func TestClient(t *testing.T) {
	//c := top.New()
	//c.AppKey = top.AppKey
	//c.AppSecret = top.AppSecret
	//
	//for i := 0; i < 150; i++ {
	//	req := top.TaobaoTbkOrderDetailsGetRequest()
	//	req.SetParam("start_time", "2020-05-27 11:00:00")
	//	req.SetParam("end_time", "2020-05-27 11:10:00")
	//	req.SetParam("order_scene", "2")
	//	body, err := c.Exec(req)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	result, err := req.Result(body)
	//	if err != nil {
	//		responseError, _ := top.ErrorResponse(body)
	//		t.Error(responseError)
	//	} else {
	//		fmt.Println(result)
	//	}
	//}

	//c := top.New()
	//c.AppKey = top.AppKey
	//c.AppSecret = top.AppSecret
	////c.Session = "Bearer 62010159c341165b1172fe5ZZab75623ec63569229eea132210884338310"
	//req := top.TaobaoTbkTpwdConvertRequest()
	//req.SetParam("password_content", "0/ej31XiOw87M/:/")
	//req.SetParam("adzone_id", "111470550243")
	//body, err := c.Exec(req)
	//if err != nil {
	//	t.Error(err)
	//}
	//
	//result, err := req.Result(body)
	//if err != nil {
	//	responseError, _ := top.ErrorResponse(body)
	//	t.Error(responseError.ErrorResponse.Code, responseError.ErrorResponse.Msg)
	//} else {
	//	fmt.Println(result.Data)
	//}

	//c := top.New()
	//c.AppKey = top.AppKey
	//c.AppSecret = top.AppSecret
	//
	//req := top.TaobaoTopAuthTokenCreateRequest()
	//req.SetParam("code", "MqdsvzXg9OADRFwzPU78E3H79952272")
	//body, err := c.Exec(req)
	//if err != nil {
	//	t.Error(err)
	//}
	//
	//result, err := req.Result(body)
	//if err != nil {
	//	responseError, _ := top.ErrorResponse(body)
	//	t.Error(responseError.ErrorResponse.Code, responseError.ErrorResponse.Msg)
	//} else {
	//	fmt.Println(result.TokenResult)
	//
	//}
	//搜索商品
	c := top.New()
	c.AppKey = top.AppKey
	c.AppSecret = top.AppSecret

	req := top.TaobaoTbkDgMaterialOptionalRequest()
	req.SetParam("q", "路亚竿套装水滴轮全套碳素钓鱼竿抛竿海竿远投海杆马口路亚杆单竿\n超轻超硬路亚套装，超远抛投，免费配节")
	req.SetParam("material_id", "17004")
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
		for _, v := range result.ResultList {
			fmt.Println(v.ItemID, v)
		}
		//fmt.Println(result.ResultList)
		t.Fatalf("数组长度为 %d", len(result.ResultList))
		t.Fatal(result.TotalResults)
	}
}
