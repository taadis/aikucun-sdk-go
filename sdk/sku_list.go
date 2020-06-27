package sdk

// SkuListRequest
type SkuListRequest struct {
	ActivityId string `json:"activityid"`
	Page       int    `json:"page"`
	PageSize   int    `json:"pagesize"`
}

// SkuListResponse
type SkuListResponse struct {
}

// SkuList
func (client *Client) SkuList(request *SkuListRequest) (response *SkuListResponse, err error) {
	response = &SkuListResponse{}
	// err = client.Do(request, response)
	return
}
