package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// proxyJSONToHTML proxies a JSON endpoint and converts it to HTML using a template
func proxyJSONToHTML(c *gin.Context, endpoint string, templateName string) {
	// Get SecAuto server URL and API key from config
	if service == nil {
		c.HTML(http.StatusInternalServerError, "error-content", gin.H{
			"message": "Service not initialized",
		})
		return
	}
	cfg := service.GetConfig()
	secautoURL := cfg.Remote.URL
	apiKey := cfg.Remote.APIKey

	fmt.Printf("Making request to: %s%s\n", secautoURL, endpoint)

	// Create request with API key header
	req, err := http.NewRequest("GET", secautoURL+endpoint, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		c.HTML(http.StatusInternalServerError, "error-content", gin.H{
			"message": "Failed to create request to SecAuto server: " + err.Error(),
		})
		return
	}

	// Add API key header
	req.Header.Set("X-API-Key", apiKey)

	// Make request to SecAuto server
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		c.HTML(http.StatusInternalServerError, "error-content", gin.H{
			"message": "Failed to connect to SecAuto server: " + err.Error(),
		})
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Response status: %d\n", resp.StatusCode)

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		c.HTML(http.StatusInternalServerError, "error-content", gin.H{
			"message": "Failed to read response from SecAuto server",
		})
		return
	}

	fmt.Printf("Response body: %s\n", string(body))

	// Parse JSON response
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		c.HTML(http.StatusInternalServerError, "error-content", gin.H{
			"message": "Failed to parse response from SecAuto server: " + err.Error(),
		})
		return
	}

	fmt.Printf("Parsed data: %+v\n", data)

	// Render HTML template with the data
	c.HTML(http.StatusOK, templateName, data)
}

// proxyAutomations proxies the automations endpoint and converts to HTML
func proxyAutomations(c *gin.Context) {
	proxyJSONToHTML(c, "/automations", "automations-response")
}

// proxyJobs proxies the jobs endpoint and converts to HTML
func proxyJobs(c *gin.Context) {
	proxyJSONToHTML(c, "/jobs", "jobs-response")
}

// proxyMetrics proxies the metrics endpoint and converts to HTML
func proxyMetrics(c *gin.Context) {
	proxyJSONToHTML(c, "/jobs/metrics", "metrics-response")
}

// proxyIntegrations proxies the integrations endpoint and converts to HTML
func proxyIntegrations(c *gin.Context) {
	proxyJSONToHTML(c, "/integrations", "integrations-response")
}

// proxyPlaybooks proxies the playbooks endpoint and converts to HTML
func proxyPlaybooks(c *gin.Context) {
	proxyJSONToHTML(c, "/playbooks", "playbooks-response")
}

// proxyPlugins proxies the plugins endpoint and converts to HTML
func proxyPlugins(c *gin.Context) {
	proxyJSONToHTML(c, "/plugins", "plugins-response")
}

// proxyAutomationDetail proxies a specific automation endpoint and converts to HTML
func proxyAutomationDetail(c *gin.Context) {
	automationID := c.Param("id")
	proxyJSONToHTML(c, "/automations/"+automationID, "automation-detail-response")
}

// proxyAutomationStatus proxies an automation status endpoint and converts to HTML
func proxyAutomationStatus(c *gin.Context) {
	automationID := c.Param("id")
	proxyJSONToHTML(c, "/automations/"+automationID+"/status", "automation-status-response")
}

// proxyLogs proxies the logs endpoint and converts to HTML
func proxyLogs(c *gin.Context) {
	automationID := c.Query("automation_id")
	endpoint := "/logs"
	if automationID != "" {
		endpoint += "?automation_id=" + automationID
	}
	proxyJSONToHTML(c, endpoint, "logs-response")
}
