package handler

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	// マルチパートフォームが利用できるメモリの制限を設定する(デフォルトは 32 MiB)
	router.MaxMultipartMemory = 1 << 30 // 1 GiB
	router.POST("/exec", ExecHandler)
	return router
}
