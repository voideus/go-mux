package routes

import (
	"github.com/voideus/golang-mux-rest/api/middlewares"
	"github.com/voideus/golang-mux-rest/controller"
	"github.com/voideus/golang-mux-rest/infrastructure"
)

type PostRoutes struct {
	handler        infrastructure.Router
	postController controller.PostController
	authMiddleware middlewares.FirebaseAuthMiddleware
}

func NewPostRoutes(
	handler infrastructure.Router,
	postController controller.PostController,
	authMiddleware middlewares.FirebaseAuthMiddleware,
) *PostRoutes {
	return &PostRoutes{
		handler:        handler,
		postController: postController,
		authMiddleware: authMiddleware,
	}
}

func (pr PostRoutes) Setup() {
	pr.handler.GET("/posts", pr.authMiddleware.Handle(), pr.postController.GetPosts)
}
