package transaction

import "github.com/gin-gonic/gin"

func Route(r *gin.Engine) {
	transaction := r.Group("api/")
	transaction.POST("customers/transactions", createTransaction)
	transaction.GET("customers/transactions", getTransactions)
}
