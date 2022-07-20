package routes

import (
	"github.com/voideus/golang-mux-rest/controller"
	"github.com/voideus/golang-mux-rest/infrastructure"
)

type PostRoutes struct {
	handler        infrastructure.Router
	postController controller.PostController
}

func NewPostRoutes(
	handler infrastructure.Router,
	postController controller.PostController,
) *PostRoutes {
	return &PostRoutes{
		handler:        handler,
		postController: postController,
	}
}

func (pr PostRoutes) Setup() {
	pr.handler.GET("/posts", pr.postController.GetPosts)
}
