package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes

	rt.router.GET("/hello/", rt.getHelloWorld)

	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	rt.router.GET("/users/:usid", rt.wrap(rt.listUserInfo))
	rt.router.POST("/users/:usid", rt.wrap(rt.createUser))
	rt.router.PUT("/users/:usid", rt.wrap(rt.updateUser))

	rt.router.GET("/users/:usid/photos/", rt.wrap(rt.listUserPhotos))
	rt.router.GET("/users/:usid/photos/:pid", rt.wrap(rt.getPhotoInfo))
	rt.router.POST("/users/:usid/photos/:pid", rt.wrap(rt.createPhoto))
	rt.router.DELETE("/users/:usid/photos/:pid", rt.wrap(rt.deletePhoto))

	rt.router.GET("/users/:usid/photos/:pid/likes/", rt.wrap(rt.listLikes))
	rt.router.PUT("/users/:usid/photos/:pid/likes/:lid", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/users/:usid/photos/:pid/likes/:lid", rt.wrap(rt.unlikePhoto))

	rt.router.POST("/users/:usid/photos/:pid/comments", rt.wrap(rt.commentPhoto))
	rt.router.GET("/users/:usid/photos/:pid/comments", rt.wrap(rt.listComments))
	rt.router.GET("/users/:usid/photos/:pid/comments/:cid", rt.wrap(rt.getCommentInfo))
	rt.router.DELETE("/users/:usid/photos/:pid/comments/:cid", rt.wrap(rt.deleteComment))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
