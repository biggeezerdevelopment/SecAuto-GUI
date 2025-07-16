package main

import (
	"html/template"
	"log"
	"net/http"
	"secauto-gui/auth"
	"secauto-gui/config"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// @title SecAuto API
// @version 1.0
// @description SecAuto Automation Framework API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	r := gin.Default()

	// Initialize authentication
	if err := auth.InitDB(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer auth.CloseDB()

	auth.InitSessions()

	// Initialize the service to connect to remote SecAuto server
	InitService()

	// Load configuration
	cfg, err := config.LoadDefaultConfig()
	if err != nil {
		log.Printf("Warning: Failed to load config, using defaults: %v", err)
		// Create default config if it doesn't exist
		if err := config.CreateDefaultConfig("config.yaml"); err != nil {
			log.Printf("Warning: Failed to create default config: %v", err)
		}
		// Try to load again
		cfg, err = config.LoadDefaultConfig()
		if err != nil {
			log.Printf("Warning: Using hardcoded defaults: %v", err)
			cfg = &config.Config{}
			cfg.Server.Port = "8080"
			cfg.Server.Host = "localhost"
			cfg.Remote.URL = "http://localhost:8000"
			cfg.Remote.APIKey = ""
		}
	}

	// Set up custom template functions
	tmpl := template.New("").Funcs(template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"split": func(s, sep string) []string {
			return strings.Split(s, sep)
		},
		"hasPrefix": func(s, prefix string) bool {
			return strings.HasPrefix(s, prefix)
		},
		"trimPrefix": func(s, prefix string) string {
			return strings.TrimPrefix(s, prefix)
		},
	})

	// Load HTML templates with custom functions
	tmpl, err = tmpl.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal("Failed to parse templates:", err)
	}

	// Debug: Print all template names
	log.Printf("Loaded templates:")
	for _, name := range tmpl.DefinedTemplates() {
		log.Printf("  - %s", name)
	}

	r.SetHTMLTemplate(tmpl)

	// CORS configuration - allow requests to SecAuto server
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080", "http://localhost:8000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With", "HX-Request"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	// Serve static files
	r.Use(static.Serve("/static", static.LocalFile("./static", false)))

	// Authentication routes (no auth required)
	r.GET("/login", auth.LoginHandler)
	r.POST("/login", auth.LoginHandler)
	r.GET("/logout", auth.LogoutHandler)

	// Protected routes (require authentication)
	protected := r.Group("/")
	protected.Use(auth.AuthMiddleware())
	{
		// Configuration endpoint for frontend
		protected.GET("/api/config", func(c *gin.Context) {
			// Return SecAuto server URL for frontend to use directly
			config := map[string]interface{}{
				"secauto_server_url": cfg.Remote.URL,
				"api_key_configured": cfg.Remote.APIKey != "",
			}
			c.JSON(http.StatusOK, config)
		})

		// Proxy endpoints for HTMX requests
		protected.GET("/api/proxy/automations", proxyAutomations)
		protected.GET("/api/proxy/jobs", proxyJobs)
		protected.GET("/api/proxy/metrics", proxyMetrics)
		protected.GET("/api/proxy/integrations", proxyIntegrations)
		protected.GET("/api/proxy/playbooks", proxyPlaybooks)
		protected.GET("/api/proxy/plugins", proxyPlugins)
		protected.GET("/api/proxy/automations/:id", proxyAutomationDetail)
		protected.GET("/api/proxy/automations/:id/status", proxyAutomationStatus)
		protected.GET("/api/proxy/logs", proxyLogs)

		// Web frontend routes
		protected.GET("/", homePage)
		protected.GET("/web/automations", automationsPage)
		protected.GET("/web/automations/:id", automationDetailPage)
		protected.GET("/web/integrations", integrationsPage)
		protected.GET("/web/integrations/new", integrationNewPage)
		protected.GET("/web/playbooks", playbooksPage)
		protected.GET("/web/plugins", pluginsPage)
		protected.GET("/web/logs", logsPage)
		protected.GET("/web/metrics", metricsPage)

		// Admin routes (require admin role)
		admin := protected.Group("/admin")
		admin.Use(auth.AdminMiddleware())
		{
			admin.GET("/users", auth.AdminUsersHandler)
			admin.POST("/users", auth.AdminUsersHandler)
			admin.DELETE("/users/:id", auth.DeleteUserHandler)
		}
	}

	// Start server
	addr := cfg.Server.Host + ":" + cfg.Server.Port
	log.Printf("SecAuto GUI starting on http://%s", addr)
	log.Printf("Connecting to remote SecAuto server at %s", cfg.Remote.URL)
	log.Printf("Remote API Documentation available at %s/api-docs", cfg.Remote.URL)
	log.Printf("API endpoints available at http://%s/api/*", addr)
	log.Printf("Configuration endpoint available at http://%s/api/config", addr)
	log.Printf("Default admin login: admin / admin123")
	log.Fatal(r.Run(addr))
}
