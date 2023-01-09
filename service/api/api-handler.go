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
	rt.router.GET("/users/:user_id", rt.wrap(rt.getUserInfo))
	rt.router.POST("/users/:user_id", rt.wrap(rt.createUser))
	rt.router.PUT("/users/:user_id", rt.wrap(rt.setMyUserName))

	// photos
	rt.router.GET("/users/:user_id/photos", rt.wrap(rt.getUserPhotos))
	rt.router.GET("/users/:user_id/photos/:photo_id", rt.wrap(rt.getPhotoInfo))
	rt.router.POST("/users/:user_id/photos/:photo_id", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/users/:user_id/photos/:photo_id", rt.wrap(rt.deletePhoto))

	// likes
	rt.router.GET("/users/:user_id/photos/:photo_id/likes", rt.wrap(rt.getLikesList))
	rt.router.PUT("/users/:user_id/photos/:photo_id/likes/:liker_id", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/users/:user_id/photos/:photo_id/likes/:liker_id", rt.wrap(rt.unlikePhoto))

	// comments
	rt.router.POST("/users/:user_id/photos/:photo_id/comments", rt.wrap(rt.commentPhoto))
	rt.router.GET("/users/:user_id/photos/:photo_id/comments", rt.wrap(rt.getCommentsList))
	rt.router.GET("/users/:user_id/photos/:photo_id/comments/:comment_id", rt.wrap(rt.getCommentInfo))
	rt.router.DELETE("/users/:user_id/photos/:photo_id/comments/:comment_id", rt.wrap(rt.uncommentPhoto))

	// followers
	rt.router.GET("/users/:user_id/followers/:follower_id", rt.wrap(rt.getFollowers))
	rt.router.DELETE("/users/:user_id/followers/:follower_id", rt.wrap(rt.removeFollower))

	// banned
	rt.router.GET("/users/:user_id/banned/:banned_id", rt.wrap(rt.getBannedList))
	rt.router.DELETE("/users/:user_id/banned/:banned_id", rt.wrap(rt.unbanUser))
	rt.router.PUT("/users/:user_id/banned/:banned_id", rt.wrap(rt.banUser))

	// following
	rt.router.GET("/users/:user_id/following/:following_id", rt.wrap(rt.getFollowing))
	rt.router.DELETE("/users/:user_id/following/:following_id", rt.wrap(rt.unfollowUser))
	rt.router.PUT("/users/:user_id/following/:following_id", rt.wrap(rt.followUser))

	// feed
	rt.router.GET("/users/:user_id/feed", rt.wrap(rt.getMyStream))

	// profile
	rt.router.GET("/users/:user_id/profile", rt.wrap(rt.getUserProfile))

	// search
	rt.router.GET("/search/:username", rt.wrap(rt.getSearchedUserInfo))

	// login
	rt.router.GET("/session", rt.wrap(rt.login))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
