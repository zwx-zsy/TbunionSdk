package top

import (
	"github.com/yancyzhou/TbunionSdk/top/requests"
	"encoding/json"
)

func ErrorResponse(data []byte) (requests.ErrorResponse, error) {
	var result requests.ErrorResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return requests.ErrorResponse{}, err
	}

	return result, nil
}
