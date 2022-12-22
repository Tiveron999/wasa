package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	//rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	rt.router.GET("/users/", rt.wrap(rt.listUserInfo))
	rt.router.POST("/users/:id", rt.wrap(rt.createUser))
	rt.router.PUT("/users/:id", rt.wrap(rt.updateUser))

	rt.router.GET("/users/photos/", rt.wrap(rt.listUserPhotos))
	rt.router.GET("/photos/", rt.wrap(rt.getPhotoInfo))
	rt.router.POST("/photos/", rt.wrap(rt.createPhoto))
	rt.router.DELETE("/photos/:id", rt.wrap(rt.deletePhoto))

	rt.router.GET("/likes/", rt.wrap(rt.listLikes))
	rt.router.PUT("/likes/:id", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/likes/:id", rt.wrap(rt.unlikePhoto))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
