package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/voideus/golang-mux-rest/service"
	"go.uber.org/fx"
)

type controller struct{}

var (
	postService service.PostService
)

type PostController interface {
	GetPosts(c *gin.Context)
	AddPost(c *gin.Context)
}

func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

func (*controller) GetPosts(c *gin.Context) {

	c.Header("Content-type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		c.Abort()
		// json.NewEncoder(c).Encode(errors.ServiceError{Message: "Error getting the posts array"})
		return
	}
	c.JSON(200, gin.H{
		"data": posts,
	})

	// result, err := json.Marshal(posts)
	// if err != nil {
	// 	res.WriteHeader(http.StatusInternalServerError)
	// 	res.Write([]byte(`{"error":"Error marshalling the posts array"}`))
	// 	return
	// }

	// c.res.WriteHeader(http.StatusOK)
	// json.NewEncoder(c.res).Encode(posts)
}

func (*controller) AddPost(c *gin.Context) {
	// var post entity.Post
	// res.Header().Set("Content-type", "application/json")
	// err := json.NewDecoder(req.Body).Decode(&post)
	// if err != nil {
	// 	res.WriteHeader(http.StatusInternalServerError)
	// 	json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error marshalling the request"})
	// 	return
	// }

	// err1 := postService.Validate(&post)
	// if err1 != nil {
	// 	res.WriteHeader(http.StatusInternalServerError)
	// 	json.NewEncoder(res).Encode(errors.ServiceError{Message: err1.Error()})
	// 	return
	// }

	// result, err2 := postService.Create(&post)
	// if err2 != nil {
	// 	res.WriteHeader(http.StatusInternalServerError)
	// 	json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error saving the post"})
	// 	return
	// }

	// res.WriteHeader(http.StatusOK)
	// json.NewEncoder(res).Encode(result)
}

var Module = fx.Options(
	fx.Provide(NewPostController),
)
