package aikucun

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

type ActivityListRequest struct {
	//BaseRequest
	Status int `json:"status,omitempty"`
}

func (request *ActivityListRequest) Method() string {
	return http.MethodGet
}

func (request *ActivityListRequest) Path() string {
	return "/api/v2/activity/list"
}

func (request *ActivityListRequest) Params() map[string]interface{} {
	m := make(map[string]interface{})
	m["status"] = strconv.Itoa(request.Status)
	return m
}

type ActivityListResponse struct {
	BaseResponse
	Data struct {
		Size int `json:"size"`
		List []struct {
			Id            string `json:"id"`
			No            string `json:"no"`
			Name          string `json:"name"`
			CorpId        string `json:"corpid"`
			CorpName      string `json:"corpname"`
			BrandId       string `json:"brandid"`
			BrandName     string `json:"brandname"`
			BeginTime     string `json:"begintime"`
			EndTime       string `json:"endtime"`
			AfterSaleTime string `json:"aftersaletime"`
		} `json:"list"`
	} `json:"data"`
}

// NewActivityListRequest
func NewActivityListRequest() *ActivityListRequest {
	return new(ActivityListRequest)
}

//
func (resp *ActivityListResponse) String() string {
	//buf, _ := json.Marshal(resp)
	//return string(buf)
	//buf := bytes.NewBuffer([]byte(""))
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(resp)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

// ActivityList
func (client *Client) ActivityList(request *ActivityListRequest) (response *ActivityListResponse, err error) {
	response = &ActivityListResponse{}
	err = client.Do(request, response)
	return
}
