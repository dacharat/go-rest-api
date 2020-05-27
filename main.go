package main

import (
	"net/http"

	"github.com/dacharat/go-rest-api/controllers"
	"github.com/dacharat/go-rest-api/models"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func main() {
	r := engine()

	r.Run()
}

func engine() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	models.ConnectDataBase()

	r.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))
	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.POST("/login", controllers.Login)
	r.POST("/signup", controllers.Signup)
	r.GET("/logout", controllers.Logout)
	r.GET("/users", controllers.GetUsers)
	r.DELETE("/users/:id", controllers.DeleteUser)

	private := r.Group("/private")
	private.Use(AuthRequired)
	{
		private.GET("/me", controllers.Me)
	}
	return r
}

// AuthRequired is a simple middleware to check the session
func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}
