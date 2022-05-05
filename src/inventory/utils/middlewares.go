package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Validator(c *gin.Context) {
	fmt.Println("I'm a dummy test test")
	c.Next()
}
