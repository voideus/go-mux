package main

import (
	"fmt"
	"net/http"

	"github.com/voideus/golang-mux-rest/controller"
	router "github.com/voideus/golang-mux-rest/http"
	"github.com/voideus/golang-mux-rest/repository"
	"github.com/voideus/golang-mux-rest/service"
)

var (
	httpRouter     router.Router             = router.NewChiRouter()
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
)

func main() {
	const port string = ":8000"
	// os.Setenv("GOOGLE_APPIVATION_CREDENTIALS", "./firebase/token.json")
	httpRouter.GET("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Up and running...")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(port)
}
