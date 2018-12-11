package pkg

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ParseRequest(c *gin.Context, request interface{}) {
	err := c.ShouldBindWith(request, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, "parse request error:"+err.Error())
		fmt.Println("ParseRequest Result", request)
		fmt.Println("ParseRequest Error", err)
		return
	}
}
