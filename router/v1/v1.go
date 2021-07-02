package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/truanter/yizhigo/http/controller"
	"github.com/truanter/yizhigo/http/middleware"
	"net/http"
)

func RegisterV1(router *gin.Engine) {
	goV1 := router.Group("/go/v1")
	{
		goV1.GET("/hello", func(ctx *gin.Context) {
			name := ctx.DefaultQuery("name", "yizhigo")
			ctx.String(http.StatusOK, fmt.Sprintf("hello, %s", name))
		})
		goV1.GET("/config", controller.GetConfig)
		goV1.GET("/index", controller.GetIndex)
		goV1.GET("/q", controller.Search)
		goV1.GET("/favorites", controller.GetFavoritesLocal)
		goV1.GET("/favorites_all", controller.GetFavorites)
		goV1.GET("/favorite_list", controller.GetFavoriteList)
		goV1.GET("/similar", controller.GetSimilarGoods)
		goV1.POST("/create_tpwd", controller.CreateTPWD)

		ops := goV1.Group("/ops", middleware.AccessTokenCheck)
		{
			ops.POST("/button_text", controller.OpsButtonText)
			ops.POST("/block_platform", controller.OpsBlockPlatform)
			ops.POST("/key_word", controller.OpsAddKeyWord)
			ops.POST("/big_brother_coming", controller.MyBigBrotherComing)
			ops.POST("/big_brother_leaving", controller.MyBigBrotherLeaving)
		}
	}
}
