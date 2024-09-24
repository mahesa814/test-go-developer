package transaction

import (
	"github.com/gin-gonic/gin"
	"test-go-developer/commons"
	"test-go-developer/modules/transaction/request"
)

// -----------------LOAN CUSTOMER---------------------------
func createTransaction(c *gin.Context) {
	var payload request.TransactionRequest
	var validator = commons.Galidator.Validator(request.TransactionRequest{})
	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		c.JSON(400, gin.H{"error": validator.DecryptErrors(err)})
		return
	}
	customer, err := createTransactionService(payload)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": customer})
}

func getTransactions(c *gin.Context) {
	var payload request.TransactionQuery
	if err := c.ShouldBindQuery(&payload); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	transactions, err := getTransactionsService(payload)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": transactions})
}
