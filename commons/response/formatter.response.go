package commonResponse

import (
	"github.com/gin-gonic/gin"
)

type MetaItemResponse struct {
	Code    int    `json:"code" example:"200"`
	Status  string `json:"status" example:"success"`
	Message string `json:"message" example:"create user success"`
}
type MetaResponse struct {
	Meta MetaItemResponse `json:"meta"`
}
type DataResponse struct {
	MetaResponse
	Data interface{} `json:"data"`
}
type DataResponseMember struct {
	MetaResponse
	Data interface{} `json:"items"`
}

func ResponseFormater(c *gin.Context, code int, status string, message string, data interface{}) {
	response := DataResponse{
		MetaResponse: MetaResponse{
			Meta: MetaItemResponse{
				Code:    code,
				Status:  status,
				Message: message,
			},
		},
		Data: data,
	}

	c.JSON(code, response)
}
