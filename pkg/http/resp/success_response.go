package resp

import "github.com/gin-gonic/gin"

type SuccessResp struct {
	StatusCode int         `json:"status_code" extensions:"x-order=1"`
	Message    string      `json:"message" extensions:"x-order=2"`
	Metadata   interface{} `json:"metadata,omitempty" extensions:"x-order=3"`
}

func NewSuccessResp(statusCode int, msg string, data interface{}) *SuccessResp {
	return &SuccessResp{
		StatusCode: statusCode,
		Message:    msg,
		Metadata:   data,
	}
}

func Response(ctx *gin.Context, response *SuccessResp) {
	ctx.JSON(response.StatusCode, response)
}
