package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Login - SecAuto",
		})
		return
	}

	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"title":    "Login - SecAuto",
			"error":    "Username and password are required",
			"username": username,
		})
		return
	}

	user, err := AuthenticateUser(username, password)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"title":    "Login - SecAuto",
			"error":    "Invalid username or password",
			"username": username,
		})
		return
	}

	// Login user
	err = LoginUser(c, user)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{
			"title":    "Login - SecAuto",
			"error":    "Failed to create session",
			"username": username,
		})
		return
	}

	c.Redirect(http.StatusFound, "/")
}

func LogoutHandler(c *gin.Context) {
	LogoutUser(c)
	c.Redirect(http.StatusFound, "/login")
}

func AdminUsersHandler(c *gin.Context) {
	if c.Request.Method == "GET" {
		fmt.Printf("GET request to /admin/users\n")
		users, err := GetUsers()
		if err != nil {
			fmt.Printf("Failed to get users: %v\n", err)
			c.HTML(http.StatusInternalServerError, "error.html", gin.H{
				"title":   "Error",
				"message": "Failed to load users",
			})
			return
		}

		fmt.Printf("Retrieved %d users for display\n", len(users))
		// Get current user
		user := GetCurrentUser(c)

		// Check if this is an HTMX request
		if c.GetHeader("HX-Request") == "true" {
			// Return only the admin users content
			c.HTML(http.StatusOK, "admin-users-content", gin.H{
				"users": users,
			})
			return
		}

		// Return full page for direct navigation
		c.HTML(http.StatusOK, "base.html", gin.H{
			"title": "User Management - SecAuto",
			"user":  user,
			"users": users,
		})
		return
	}

	// Handle user creation
	fmt.Printf("Form data received: %+v\n", c.Request.Form)
	fmt.Printf("Content-Type: %s\n", c.GetHeader("Content-Type"))

	// Try to get form data manually
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	role := c.PostForm("role")

	fmt.Printf("Manual form data - Username: '%s', Email: '%s', Password: '%s', Role: '%s'\n",
		username, email, password, role)

	var userData UserCreate
	if err := c.ShouldBind(&userData); err != nil {
		// Get current user
		user := GetCurrentUser(c)
		c.HTML(http.StatusBadRequest, "base.html", gin.H{
			"title": "User Management - SecAuto",
			"user":  user,
			"error": "Invalid form data: " + err.Error(),
		})
		return
	}

	// Use manually extracted form data
	userData = UserCreate{
		Username: username,
		Email:    email,
		Password: password,
		Role:     role,
	}

	// Validate required fields
	if userData.Username == "" || userData.Email == "" || userData.Password == "" {
		// Get current user
		user := GetCurrentUser(c)
		c.HTML(http.StatusBadRequest, "base.html", gin.H{
			"title": "User Management - SecAuto",
			"user":  user,
			"error": "Username, email, and password are required",
		})
		return
	}

	// Set default role if not specified
	if userData.Role == "" {
		userData.Role = "user"
	}

	fmt.Printf("Creating user: %+v\n", userData)
	err := CreateUser(userData)
	if err != nil {
		// Get current user
		user := GetCurrentUser(c)
		c.HTML(http.StatusInternalServerError, "base.html", gin.H{
			"title": "User Management - SecAuto",
			"user":  user,
			"error": "Failed to create user: " + err.Error(),
		})
		return
	}

	// Get updated user list
	fmt.Printf("Retrieving users after creation...\n")
	users, err := GetUsers()
	if err != nil {
		// Get current user
		user := GetCurrentUser(c)
		c.HTML(http.StatusInternalServerError, "base.html", gin.H{
			"title": "User Management - SecAuto",
			"user":  user,
			"error": "Failed to load users after creation",
		})
		return
	}

	// Get current user
	user := GetCurrentUser(c)
	fmt.Printf("Rendering page with %d users\n", len(users))
	c.HTML(http.StatusOK, "base.html", gin.H{
		"title":   "User Management - SecAuto",
		"user":    user,
		"users":   users,
		"success": "User created successfully",
	})
}

func DeleteUserHandler(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	// Convert string to int (you might want to add proper validation)
	var id int
	_, err := fmt.Sscanf(userID, "%d", &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
