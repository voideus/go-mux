package router

import (
	// "fmt"

	"github.com/gin-gonic/gin"
	"github.com/voideus/golang-mux-rest/api/middlewares"
	"github.com/voideus/golang-mux-rest/controller"
)

// type ginRouter struct {
// 	*gin.Engine
// }

type GinHandler struct {
	*gin.Engine
	controller.PostController
	authMiddleWare middlewares.FirebaseAuthMiddleware
}

// var (
// 	ginDispatcher := gin.Default()
// )

func NewGinRouter(gin *gin.Engine, PostController controller.PostController, authMiddleware middlewares.FirebaseAuthMiddleware) GinHandler {
	return GinHandler{
		gin, PostController, authMiddleware,
	}
}

// func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
// 	ginDispatcher.HandleFunc(uri, f).Methods("GET")
// }
// func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
// 	ginDispatcher.HandleFunc(uri, f).Methods("POST")
// }
// func (*muxRouter) SERVE(port string) {
// 	fmt.Printf("Mux HTTP server running on port %v", port)
// 	http.ListenAndServe(port, ginDispatcher)
// }

func (h *GinHandler) SetupRoutes() {
	h.GET("/posts", h.PostController.GetPosts)
	// h.mux.HandleFunc()
}
