package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {
	g := gin.Default()

	v1 := g.Group("/api/v1")
	{

		//Event route
		v1.POST("/events", app.createEvent)
		v1.GET("/events", app.getAllEvents)
		v1.GET("/events/:id", app.getEvent)
		v1.PUT("/events/:id", app.updateEvent)
		v1.DELETE("/events/:id", app.deleteEvent)

		//Attendence route
		v1.POST("/events/:id/attendence/:userId", app.addAttendenceToEvent)
		v1.GET("/events/:id/attendence", app.getAttendenceForEvent)

		//Auth route
		v1.POST("/auth/register", app.registerUser)
	}

	return g
}
