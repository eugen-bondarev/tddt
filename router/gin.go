package router

import "github.com/gin-gonic/gin"

type GinCtx struct {
	ginCtx *gin.Context
}

func (c *GinCtx) GetBody(v any) error {
	return c.ginCtx.ShouldBindJSON(v)
}

func GinWrap(h Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &GinCtx{c}
		data, err := h(ctx)
		if err != nil {
			if httpErr, ok := err.(*HttpError); ok {
				c.JSON(httpErr.Code, gin.H{"error": httpErr.Message})
			} else {
				c.JSON(500, gin.H{"error": err.Error()})
			}
			return
		}
		c.JSON(200, data)
	}
}
