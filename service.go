package main

import (
	"log"
	"secauto-gui/config"
	"secauto-gui/sdk"
)

// Service provides access to the remote SecAuto server via the SDK
type Service struct {
	client *sdk.Client
	config *config.Config
}

// NewService creates a new service instance
func NewService() *Service {
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
			cfg.Remote.URL = "http://localhost:8000"
			cfg.Remote.APIKey = ""
		}
	}

	// Create client with configuration
	client := sdk.NewClientFromConfig(cfg.Remote.URL, cfg.Remote.APIKey)

	return &Service{
		client: client,
		config: cfg,
	}
}

// GetConfig returns the current configuration
func (s *Service) GetConfig() *config.Config {
	return s.config
}

// GetHealth checks the health of the remote SecAuto server
func (s *Service) GetHealth() (*sdk.HealthResponse, error) {
	return s.client.GetHealth()
}

// GetAutomations retrieves all automations from the remote server
func (s *Service) GetAutomations() (*sdk.AutomationsResponse, error) {
	return s.client.GetAutomations()
}

// UploadAutomation uploads a new automation script
func (s *Service) UploadAutomation(filePath string) (*sdk.UploadAutomationResponse, error) {
	return s.client.UploadAutomation(filePath)
}

// DeleteAutomation deletes an automation script
func (s *Service) DeleteAutomation(name string) (*sdk.DeleteAutomationResponse, error) {
	return s.client.DeleteAutomation(name)
}

// GetIntegrations retrieves all integrations from the remote server
func (s *Service) GetIntegrations() (*sdk.IntegrationsResponse, error) {
	return s.client.GetIntegrations()
}

// GetIntegration retrieves a specific integration by name
func (s *Service) GetIntegration(name string) (*sdk.IntegrationResponse, error) {
	return s.client.GetIntegration(name)
}

// CreateIntegration creates a new integration
func (s *Service) CreateIntegration(integration sdk.Integration) (*sdk.IntegrationResponse, error) {
	return s.client.CreateIntegration(integration)
}

// UpdateIntegration updates an existing integration
func (s *Service) UpdateIntegration(name string, integration sdk.Integration) (*sdk.IntegrationResponse, error) {
	return s.client.UpdateIntegration(name, integration)
}

// DeleteIntegration deletes an integration
func (s *Service) DeleteIntegration(name string) error {
	return s.client.DeleteIntegration(name)
}

// UploadIntegration uploads a new integration file
func (s *Service) UploadIntegration(filePath string) (*sdk.UploadPluginResponse, error) {
	return s.client.UploadIntegration(filePath)
}

// DeleteIntegrationFile deletes an integration file
func (s *Service) DeleteIntegrationFile(name string) (*sdk.DeletePluginResponse, error) {
	return s.client.DeleteIntegrationFile(name)
}

// GetPlaybooks retrieves all playbooks from the remote server
func (s *Service) GetPlaybooks() (*sdk.PlaybooksResponse, error) {
	return s.client.GetPlaybooks()
}

// UploadPlaybook uploads a new playbook file
func (s *Service) UploadPlaybook(filePath string) (*sdk.UploadPlaybookResponse, error) {
	return s.client.UploadPlaybook(filePath)
}

// DeletePlaybook deletes a playbook
func (s *Service) DeletePlaybook(name string) (*sdk.DeletePlaybookResponse, error) {
	return s.client.DeletePlaybook(name)
}

// ExecutePlaybook executes a playbook synchronously
func (s *Service) ExecutePlaybook(req sdk.PlaybookExecutionRequest) (*sdk.PlaybookExecutionResponse, error) {
	return s.client.ExecutePlaybook(req)
}

// ExecutePlaybookAsync executes a playbook asynchronously
func (s *Service) ExecutePlaybookAsync(req sdk.PlaybookExecutionRequest) error {
	return s.client.ExecutePlaybookAsync(req)
}

// GetJobs retrieves all jobs with optional filtering
func (s *Service) GetJobs(status sdk.JobStatus, limit, offset int) (*sdk.JobsResponse, error) {
	return s.client.GetJobs(status, limit, offset)
}

// GetSchedules retrieves all schedules
func (s *Service) GetSchedules() (*sdk.SchedulesResponse, error) {
	return s.client.GetSchedules()
}

// CreateSchedule creates a new schedule
func (s *Service) CreateSchedule(schedule sdk.Schedule) error {
	return s.client.CreateSchedule(schedule)
}

// GetPlugins retrieves all plugins
func (s *Service) GetPlugins() (*sdk.PluginsResponse, error) {
	return s.client.GetPlugins()
}

// UploadPlugin uploads a new plugin file
func (s *Service) UploadPlugin(pluginType sdk.PluginType, filePath string) (*sdk.UploadPluginResponse, error) {
	return s.client.UploadPlugin(pluginType, filePath)
}

// DeletePlugin deletes a plugin
func (s *Service) DeletePlugin(pluginType sdk.PluginType, name string) (*sdk.DeletePluginResponse, error) {
	return s.client.DeletePlugin(pluginType, name)
}

// ValidatePlaybook validates a playbook without executing
func (s *Service) ValidatePlaybook(req sdk.ValidationRequest) (*sdk.ValidationResponse, error) {
	return s.client.ValidatePlaybook(req)
}

// ConfigureWebhook configures webhook notifications
func (s *Service) ConfigureWebhook(webhook sdk.Webhook) error {
	return s.client.ConfigureWebhook(webhook)
}

// GetCluster retrieves cluster information
func (s *Service) GetCluster() (map[string]interface{}, error) {
	return s.client.GetCluster()
}

// GetJobMetrics retrieves database performance metrics
func (s *Service) GetJobMetrics() (map[string]interface{}, error) {
	return s.client.GetJobMetrics()
}

// GetJobStats retrieves comprehensive job statistics
func (s *Service) GetJobStats() (map[string]interface{}, error) {
	return s.client.GetJobStats()
}

// Global service instance
var service *Service

// InitService initializes the global service
func InitService() {
	service = NewService()
	cfg := service.GetConfig()
	log.Printf("SecAuto service initialized - connecting to remote server at %s", cfg.Remote.URL)
	if cfg.Remote.APIKey != "" {
		log.Printf("Using API key: %s", cfg.Remote.APIKey[:8]+"...")
	} else {
		log.Printf("No API key configured")
	}
}
