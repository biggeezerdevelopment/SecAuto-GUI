package main

import (
	"log"
	"net/http"
	"time"

	"secauto-gui/auth"

	"github.com/gin-gonic/gin"
)

// homePage serves the main dashboard page
func homePage(c *gin.Context) {
	// Get current user
	user := auth.GetCurrentUser(c)

	// Get SecAuto server URL from config
	cfg := service.GetConfig()
	secautoURL := cfg.Remote.URL

	// Check if this is an HTMX request
	if c.GetHeader("HX-Request") == "true" {
		log.Printf("HTMX request detected for dashboard")
		// Return only the dashboard content
		c.HTML(http.StatusOK, "dashboard-content", gin.H{
			"secauto_url": secautoURL,
		})
		return
	}

	log.Printf("Non-HTMX request for dashboard")

	// Return full page for direct navigation
	c.HTML(http.StatusOK, "base.html", gin.H{
		"title":       "SecAuto Dashboard",
		"timestamp":   time.Now(),
		"user":        user,
		"secauto_url": secautoURL,
	})
}

// automationsPage serves the automations list page
func automationsPage(c *gin.Context) {
	// Get current user
	user := auth.GetCurrentUser(c)

	// Get SecAuto server URL from config
	cfg := service.GetConfig()
	secautoURL := cfg.Remote.URL

	// Check if this is an HTMX request
	if c.GetHeader("HX-Request") == "true" {
		log.Printf("HTMX request detected for automations")
		// Return only the automations content
		c.HTML(http.StatusOK, "automations-content", gin.H{
			"secauto_url": secautoURL,
		})
		return
	}

	// Return full page for direct navigation
	c.HTML(http.StatusOK, "base.html", gin.H{
		"title":       "Automations",
		"timestamp":   time.Now(),
		"user":        user,
		"secauto_url": secautoURL,
	})
}

// integrationsPage serves the integrations list page
func integrationsPage(c *gin.Context) {
	// Get current user
	user := auth.GetCurrentUser(c)

	// Get SecAuto server URL from config
	cfg := service.GetConfig()
	secautoURL := cfg.Remote.URL

	// Check if this is an HTMX request
	if c.GetHeader("HX-Request") == "true" {
		// Return only the integrations content
		c.HTML(http.StatusOK, "integrations-content", gin.H{
			"secauto_url": secautoURL,
		})
		return
	}

	// Return full page for direct navigation
	c.HTML(http.StatusOK, "base.html", gin.H{
		"title":       "Integrations",
		"timestamp":   time.Now(),
		"user":        user,
		"secauto_url": secautoURL,
	})
}

// playbooksPage serves the playbooks list page
func playbooksPage(c *gin.Context) {
	// Get current user
	user := auth.GetCurrentUser(c)

	// Get SecAuto server URL from config
	cfg := service.GetConfig()
	secautoURL := cfg.Remote.URL

	// Check if this is an HTMX request
	if c.GetHeader("HX-Request") == "true" {
		// Return only the playbooks content
		c.HTML(http.StatusOK, "playbooks-content", gin.H{
			"secauto_url": secautoURL,
		})
		return
	}

	// Return full page for direct navigation
	c.HTML(http.StatusOK, "base.html", gin.H{
		"title":       "Playbooks",
		"timestamp":   time.Now(),
		"user":        user,
		"secauto_url": secautoURL,
	})
}

// pluginsPage serves the plugins list page
func pluginsPage(c *gin.Context) {
	// Get current user
	user := auth.GetCurrentUser(c)

	// Get SecAuto server URL from config
	cfg := service.GetConfig()
	secautoURL := cfg.Remote.URL

	// Check if this is an HTMX request
	if c.GetHeader("HX-Request") == "true" {
		// Return only the plugins content
		c.HTML(http.StatusOK, "plugins-content", gin.H{
			"secauto_url": secautoURL,
		})
		return
	}

	// Return full page for direct navigation
	c.HTML(http.StatusOK, "base.html", gin.H{
		"title":       "Plugins",
		"timestamp":   time.Now(),
		"user":        user,
		"secauto_url": secautoURL,
	})
}

// automationDetailPage serves the automation detail page
func automationDetailPage(c *gin.Context) {
	// Get current user
	user := auth.GetCurrentUser(c)

	id := c.Param("id")

	// Get SecAuto server URL from config
	cfg := service.GetConfig()
	secautoURL := cfg.Remote.URL

	// Check if this is an HTMX request
	if c.GetHeader("HX-Request") == "true" {
		// Return only the automation detail content
		c.HTML(http.StatusOK, "automation-detail-content", gin.H{
			"automation_id": id,
			"secauto_url":   secautoURL,
		})
		return
	}

	// Return full page for direct navigation
	c.HTML(http.StatusOK, "base.html", gin.H{
		"title":         "Automation Details",
		"timestamp":     time.Now(),
		"user":          user,
		"automation_id": id,
		"secauto_url":   secautoURL,
	})
}

// logsPage serves the logs page
func logsPage(c *gin.Context) {
	// Get current user
	user := auth.GetCurrentUser(c)

	// Get SecAuto server URL from config
	cfg := service.GetConfig()
	secautoURL := cfg.Remote.URL

	// Check if this is an HTMX request
	if c.GetHeader("HX-Request") == "true" {
		// Return only the logs content
		c.HTML(http.StatusOK, "logs-content", gin.H{
			"secauto_url": secautoURL,
		})
		return
	}

	// Return full page for direct navigation
	c.HTML(http.StatusOK, "base.html", gin.H{
		"title":       "System Logs",
		"timestamp":   time.Now(),
		"user":        user,
		"secauto_url": secautoURL,
	})
}

// metricsPage serves the metrics page
func metricsPage(c *gin.Context) {
	// Get current user
	user := auth.GetCurrentUser(c)

	// Get SecAuto server URL from config
	cfg := service.GetConfig()
	secautoURL := cfg.Remote.URL

	// Check if this is an HTMX request
	if c.GetHeader("HX-Request") == "true" {
		// Return only the metrics content
		c.HTML(http.StatusOK, "metrics-content", gin.H{
			"secauto_url": secautoURL,
		})
		return
	}

	// Return full page for direct navigation
	c.HTML(http.StatusOK, "base.html", gin.H{
		"title":       "System Metrics",
		"timestamp":   time.Now(),
		"user":        user,
		"secauto_url": secautoURL,
	})
}

// integrationNewPage serves the new integration page
func integrationNewPage(c *gin.Context) {
	// Get current user
	user := auth.GetCurrentUser(c)

	// Get SecAuto server URL from config
	cfg := service.GetConfig()
	secautoURL := cfg.Remote.URL

	// Check if this is an HTMX request
	if c.GetHeader("HX-Request") == "true" {
		// Return only the integration new content
		c.HTML(http.StatusOK, "integration-new-content", gin.H{
			"secauto_url": secautoURL,
		})
		return
	}

	// Return full page for direct navigation
	c.HTML(http.StatusOK, "base.html", gin.H{
		"title":       "Create New Integration",
		"timestamp":   time.Now(),
		"user":        user,
		"secauto_url": secautoURL,
	})
}
