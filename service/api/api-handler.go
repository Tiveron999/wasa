package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes

	rt.router.GET("/hello/", rt.getHelloWorld)

	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// users
	rt.router.GET("/users/:usid", rt.wrap(rt.getUserInfo))
	rt.router.POST("/users/:usid", rt.wrap(rt.createUser))
	rt.router.PUT("/users/:usid", rt.wrap(rt.setMyUserName))

	// photos
	rt.router.GET("/users/:usid/photos/", rt.wrap(rt.getUserPhotos))
	rt.router.GET("/users/:usid/photos/:pid", rt.wrap(rt.getPhotoInfo))
	rt.router.POST("/users/:usid/photos/:pid", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/users/:usid/photos/:pid", rt.wrap(rt.deletePhoto))

	// likes
	rt.router.GET("/users/:usid/photos/:pid/likes/", rt.wrap(rt.getLikesList))
	rt.router.PUT("/users/:usid/photos/:pid/likes/:lid", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/users/:usid/photos/:pid/likes/:lid", rt.wrap(rt.unlikePhoto))

	// comments
	rt.router.POST("/users/:usid/photos/:pid/comments", rt.wrap(rt.commentPhoto))
	rt.router.GET("/users/:usid/photos/:pid/comments", rt.wrap(rt.getCommentsList))
	rt.router.GET("/users/:usid/photos/:pid/comments/:cid", rt.wrap(rt.getCommentInfo))
	rt.router.DELETE("/users/:usid/photos/:pid/comments/:cid", rt.wrap(rt.uncommentPhoto))

	// followers
	rt.router.GET("/users/:usid/followers/:userid/:followerid", rt.wrap(rt.getFollowers))
	rt.router.DELETE("/users/:usid/followers/:userid/:followerid", rt.wrap(rt.removeFollower))

	// banned
	rt.router.GET("/users/:usid/banned/:userid/:bannedid", rt.wrap(rt.getBannedList))
	rt.router.DELETE("/users/:usid/banned/:userid/:bannedid", rt.wrap(rt.unbanUser))
	rt.router.PUT("/users/:usid/banned/:userid/:bannedid", rt.wrap(rt.banUser))

	// following
	rt.router.GET("/users/:usid/following/:userid/:followerid", rt.wrap(rt.getFollowing))
	rt.router.DELETE("/users/:usid/following/:userid/:followerid", rt.wrap(rt.unfollowUser))
	rt.router.PUT("/users/:usid/following/:userid/:followerid", rt.wrap(rt.followUser))

	// feed
	rt.router.GET("/users/:usid/feed", rt.wrap(rt.getMyStream))

	// profile
	rt.router.GET("/users/:usid/profile", rt.wrap(rt.getUserProfile))

	// search
	rt.router.GET("/users/search", rt.wrap(rt.getSearchedUserInfo))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
