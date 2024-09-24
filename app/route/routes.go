package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"test-go-developer/commons"
	"test-go-developer/configs"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(func() string {
		if configs.AppMode == "production" {
			return gin.ReleaseMode
		}
		return gin.DebugMode // Set default mode if not in production
	}())

	r := gin.Default()

	// Collect route list for documentation or debugging
	routes := r.Routes()
	for _, item := range routes {
		commons.RouteList = append(commons.RouteList, fmt.Sprintf("%s %s", item.Method, item.Path))
	}
	return r
}
