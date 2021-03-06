package Navigation

import (
	"Navigation-Web/controller"

	"github.com/gin-gonic/gin"
)

type navigationRouter struct {
}

var NavigationRouter = new(navigationRouter)

func (n *navigationRouter) InitNavigationRouter(r *gin.RouterGroup) {
	navigationRouter := r.Group("navigation")
	{
		navigationRouter.GET("/navigation", controller.NavigationInfoController.GetAll)
		navigationRouter.POST("/navigation",  controller.NavigationInfoController.Create)
		navigationRouter.PUT("/navigation/:id",  controller.NavigationInfoController.Update)
		navigationRouter.DELETE("/navigation",  controller.NavigationInfoController.Delete)
	}
}
