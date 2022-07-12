package router

import (
	// "fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/voideus/golang-mux-rest/controller"
)

type muxRouter struct {
}

type MuxHandler struct {
	*http.ServeMux
	controller.PostController
}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter(mux *http.ServeMux, PostController controller.PostController) MuxHandler {
	return MuxHandler{
		mux, PostController,
	}
}

// func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
// 	muxDispatcher.HandleFunc(uri, f).Methods("GET")
// }
// func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
// 	muxDispatcher.HandleFunc(uri, f).Methods("POST")
// }
// func (*muxRouter) SERVE(port string) {
// 	fmt.Printf("Mux HTTP server running on port %v", port)
// 	http.ListenAndServe(port, muxDispatcher)
// }

func (h *MuxHandler) SetupRoutes() {
	h.HandleFunc("/posts", h.PostController.GetPosts)
	// h.mux.HandleFunc()
}
