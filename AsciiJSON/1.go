package main
import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		data:=map[string]interface{}{
			"lang":"go语言",
			"tag":"<br>",
		}
		c.AsciiJSON(http.StatusOK,data)
	})
	r.Run(":3000") 
}