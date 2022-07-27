package handlers

import (
	"go_jwt_statefull/exception"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Recovery(ctx *gin.Context, recovered interface{}) {
	if internalServerError(ctx, recovered) {
		return
	}
	if badRequest(ctx, recovered) {
		return
	}
	if NotFound(ctx, recovered) {
		return
	}
}

func internalServerError(ctx *gin.Context, recovered interface{}) bool {
	result, ok := recovered.(*exception.InternalServer)
	if ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": result,
		})
		return true
	}
	return false
}

func badRequest(ctx *gin.Context, recovered interface{}) bool {
	result, ok := recovered.(*exception.BadRequest)
	if ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result,
		})
		return true
	}
	return false
}

func NotFound(ctx *gin.Context, recovered interface{}) bool {
	result, ok := recovered.(*exception.NotFound)
	if ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": result,
		})
		return true
	}
	return false
}
