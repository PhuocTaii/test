package models

type ListPaging struct {
	Page    int64       `json:"page" form:"page"`
	Size    int64       `json:"size" form:"size"`
	Count   int64       `json:"count" form:"count"`
	Records interface{} `json:"records" form:"records"`
}

type RequestParams struct {
	Page         int64  `json:"page" form:"page"  extensions:"x-order=1"`
	Size         int64  `json:"size" form:"size" extensions:"x-order=2"`
	ConfigQuery  int64  `json:"config_query" form:"config_query" extensions:"x-order=3"`
	SortBy       string `json:"sort_by" form:"sort_by" extensions:"x-order=4"`
	SortOrder    string `json:"sort_order" form:"sort_order" extensions:"x-order=5"`
	SortMultiple string `json:"sort_multiple" form:"sort_multiple" extensions:"x-order=6"`
}
