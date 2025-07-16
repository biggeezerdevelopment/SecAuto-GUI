package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

func InitSessions() {
	store = sessions.NewCookieStore([]byte("secauto-secret-key-change-in-production"))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days
		HttpOnly: true,
		Secure:   false, // Set to true in production with HTTPS
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip auth for login page and static assets
		if c.Request.URL.Path == "/login" || c.Request.URL.Path == "/auth/login" {
			c.Next()
			return
		}

		session, err := store.Get(c.Request, "secauto-session")
		if err != nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		// Check if user is authenticated
		userID, ok := session.Values["user_id"].(int)
		if !ok {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		// Get user from database
		user, err := GetUserByID(userID)
		if err != nil {
			session.Values["user_id"] = nil
			session.Save(c.Request, c.Writer)
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		// Add user to context
		c.Set("user", user)
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		u := user.(*User)
		if u.Role != "admin" {
			c.HTML(http.StatusForbidden, "error.html", gin.H{
				"title":   "Access Denied",
				"message": "You don't have permission to access this page.",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func LoginUser(c *gin.Context, user *User) error {
	session, err := store.Get(c.Request, "secauto-session")
	if err != nil {
		return err
	}

	session.Values["user_id"] = user.ID
	session.Values["username"] = user.Username
	session.Values["role"] = user.Role
	session.Options.MaxAge = 86400 * 7 // 7 days

	return session.Save(c.Request, c.Writer)
}

func LogoutUser(c *gin.Context) error {
	session, err := store.Get(c.Request, "secauto-session")
	if err != nil {
		return err
	}

	session.Values["user_id"] = nil
	session.Values["username"] = nil
	session.Values["role"] = nil
	session.Options.MaxAge = -1

	return session.Save(c.Request, c.Writer)
}

func GetCurrentUser(c *gin.Context) *User {
	user, exists := c.Get("user")
	if !exists {
		return nil
	}
	return user.(*User)
}
