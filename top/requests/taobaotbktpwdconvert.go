package requests

import (
	"encoding/json"
)

// 淘宝客-推广者-淘口令解析&转链
// Link: https://alimama.open.taobao.com/doc.htm?docId=32932&docType=2
// Author: vincent
// Email: zhoubiao@xcess.cn

type TaobaoTbkTpwdConvert struct {
	params Parameter // 参数
}

func (o TaobaoTbkTpwdConvert) New() *TaobaoTbkTpwdConvert {
	r := new(TaobaoTbkTpwdConvert)
	r.params = make(Parameter)
	return r
}

func (o TaobaoTbkTpwdConvert) Result(data []byte) (TaobaoTbkTpwdConvertResponse, error) {
	var result TaobaoTbkTpwdConvertResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return TaobaoTbkTpwdConvertResponse{}, err
	}
	if result.RequestID == "" {
		return TaobaoTbkTpwdConvertResponse{}, API_RESPONSE_ERROR
	}

	return result, nil
}

func (o *TaobaoTbkTpwdConvert) SetParam(key, value string) {
	o.params[key] = value
}

func (o *TaobaoTbkTpwdConvert) GetMethod() string {
	return "taobao.tbk.tpwd.convert"
}

func (o *TaobaoTbkTpwdConvert) GetApiParams() Parameter {
	return o.params
}
