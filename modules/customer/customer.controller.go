package customer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator/v2"
	"net/http"
	"test-go-developer/modules/customer/request"
)

func create(c *gin.Context) {
	var payload request.CustomerRequests

	var g = galidator.G()

	var validator = g.Validator(request.CustomerRequests{})

	if err := c.ShouldBind(&payload); err != nil {
		fmt.Println(payload)
		c.JSON(http.StatusBadRequest, gin.H{"error": validator.DecryptErrors(err)})
		return
	}

	fmt.Println(payload)
	customer, err := createService(payload)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": customer})
}
func getCustomers(c *gin.Context) {
	var query request.CustomerQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	customer, err := getCustomerService(query)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": customer})
}
