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
