package sdk

type ActivityListRequest struct {
	BaseRequest
	Status int `json:"status,omitempty"`
}

func (request *ActivityListRequest) Method() string {
	return "get"
}

func (request *ActivityListRequest) Path() string {
	return "/api/v2/activity/list"
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
		}
	}
}

// ActivityList
func (client *Client) ActivityList(request *ActivityListRequest) (response *ActivityListResponse, err error) {
	response = &ActivityListResponse{}
	err = client.Do(request, response)
	return
}
