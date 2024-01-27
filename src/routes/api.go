package routes

import (
	"NetGo/src/controllers"
	"NetGo/src/types"
)

// A mapping of routes to controllers.
// The key is the route, the value is the controller function
// Define your API endpoints here...
var ApiRoutes = types.RouteList{
	"/posts": controllers.PostsController,
	"/posts/:id": controllers.PostsController,
	"/posts/:id/comments": controllers.PostCommentsController,
	"/comments": controllers.CommentsController,
	"/comments/:id": controllers.CommentsController,
}