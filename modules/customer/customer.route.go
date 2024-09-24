package customer

import "github.com/gin-gonic/gin"

func Route(r *gin.Engine) {
	transaction := r.Group("api/")
	transaction.POST("customers", create)
	transaction.GET("customers", getCustomers)
}
