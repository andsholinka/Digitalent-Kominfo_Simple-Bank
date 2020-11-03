package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//WrapAPIError dipanggil jika keluaran yang diinginkan berupa body yang berisi error
func WrapAPIError(c *gin.Context, message string, code int) {
	c.JSON(code, map[string]interface{}{
		"code":          code,
		"error_type":    http.StatusText(code),
		"error_details": message,
	})
}

//WrapAPISuccess dipanggil jika keluaran yang diinginkan berupa success message saja
func WrapAPISuccess(c *gin.Context, message string, code int) {
	c.JSON(code, map[string]interface{}{
		"code":   code,
		"status": message,
	})
}

//WrapAPIData dipanggil jika keluaran yang diinginkan berupa success message dan data output
func WrapAPIData(c *gin.Context, data interface{}, code int, message string) {
	c.JSON(code, map[string]interface{}{
		"code":   code,
		"status": message,
		"data":   data,
	})
}
