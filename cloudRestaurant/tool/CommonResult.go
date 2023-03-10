package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS int = 0 //操作成功
	FAILED  int = 0 //操作失败
)

func Success(ctx *gin.Context, v any) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": SUCCESS,
		"msg":  "成功",
		"data": v,
	})
}

func Failed(ctx *gin.Context, v any) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": FAILED,
		"msg":  v,
	})
}
