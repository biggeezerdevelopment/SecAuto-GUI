package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// Client represents the SecAuto API client
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	APIKey     string
}

// NewClient creates a new SecAuto API client
func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// NewDefaultClient creates a new SecAuto API client with the default remote server
func NewDefaultClient() *Client {
	return NewClient("http://localhost:8000")
}

// NewClientWithAuth creates a new SecAuto API client with authentication
func NewClientWithAuth(baseURL, apiKey string) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		APIKey: apiKey,
	}
}

// NewClientFromConfig creates a new SecAuto API client from configuration
func NewClientFromConfig(baseURL, apiKey string) *Client {
	client := NewClient(baseURL)
	client.APIKey = apiKey
	return client
}

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`
}

// Automation represents an automation script
type Automation struct {
	Name          string    `json:"name"`
	Filename      string    `json:"filename"`
	FileType      string    `json:"file_type"`
	Language      string    `json:"language"`
	LineCount     int       `json:"line_count"`
	FunctionCount int       `json:"function_count"`
	ImportCount   int       `json:"import_count"`
	Size          int       `json:"size"`
	IsValid       bool      `json:"is_valid"`
	ModifiedAt    time.Time `json:"modified_at"`
}

// AutomationsResponse represents the response for listing automations
type AutomationsResponse struct {
	Automations []Automation `json:"automations"`
	Count       int          `json:"count"`
	Message     string       `json:"message"`
	Success     bool         `json:"success"`
	Timestamp   time.Time    `json:"timestamp"`
}

// UploadAutomationResponse represents the response for uploading an automation
type UploadAutomationResponse struct {
	AutomationName string    `json:"automation_name"`
	Filename       string    `json:"filename"`
	Message        string    `json:"message"`
	Size           int       `json:"size"`
	Success        bool      `json:"success"`
	Timestamp      time.Time `json:"timestamp"`
}

// DeleteAutomationResponse represents the response for deleting an automation
type DeleteAutomationResponse struct {
	AutomationName string    `json:"automation_name"`
	Dependencies   []string  `json:"dependencies"`
	Message        string    `json:"message"`
	Success        bool      `json:"success"`
	Timestamp      time.Time `json:"timestamp"`
}

// Integration represents an integration configuration
type Integration struct {
	Name        string                 `json:"name"`
	Type        string                 `json:"type"`
	Description string                 `json:"description"`
	Enabled     bool                   `json:"enabled"`
	Version     string                 `json:"version"`
	URL         string                 `json:"url"`
	Username    string                 `json:"username"`
	Password    string                 `json:"password"`
	APIKey      string                 `json:"apikey"`
	Token       string                 `json:"token"`
	Secret      string                 `json:"secret"`
	Settings    map[string]interface{} `json:"settings"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
}

// IntegrationsResponse represents the response for listing integrations
type IntegrationsResponse struct {
	Integrations []Integration `json:"integrations"`
	Message      string        `json:"message"`
	Success      bool          `json:"success"`
	Timestamp    time.Time     `json:"timestamp"`
}

// IntegrationResponse represents the response for a single integration
type IntegrationResponse struct {
	Integration Integration `json:"integration"`
	Message     string      `json:"message"`
	Success     bool        `json:"success"`
	Timestamp   time.Time   `json:"timestamp"`
}

// Playbook represents a playbook
type Playbook struct {
	Name       string         `json:"name"`
	Filename   string         `json:"filename"`
	IsValid    bool           `json:"is_valid"`
	RuleCount  int            `json:"rule_count"`
	Operations map[string]int `json:"operations"`
	Size       int            `json:"size"`
	ModifiedAt time.Time      `json:"modified_at"`
}

// PlaybooksResponse represents the response for listing playbooks
type PlaybooksResponse struct {
	Playbooks []Playbook `json:"playbooks"`
	Count     int        `json:"count"`
	Message   string     `json:"message"`
	Success   bool       `json:"success"`
	Timestamp time.Time  `json:"timestamp"`
}

// UploadPlaybookResponse represents the response for uploading a playbook
type UploadPlaybookResponse struct {
	PlaybookName string    `json:"playbook_name"`
	Filename     string    `json:"filename"`
	Message      string    `json:"message"`
	Size         int       `json:"size"`
	Success      bool      `json:"success"`
	Timestamp    time.Time `json:"timestamp"`
}

// DeletePlaybookResponse represents the response for deleting a playbook
type DeletePlaybookResponse struct {
	PlaybookName string    `json:"playbook_name"`
	Message      string    `json:"message"`
	Success      bool      `json:"success"`
	Timestamp    time.Time `json:"timestamp"`
}

// JobStatus represents the status of a job
type JobStatus string

const (
	JobStatusPending   JobStatus = "pending"
	JobStatusRunning   JobStatus = "running"
	JobStatusCompleted JobStatus = "completed"
	JobStatusFailed    JobStatus = "failed"
	JobStatusCancelled JobStatus = "cancelled"
)

// Job represents a job
type Job struct {
	ID          string                 `json:"id"`
	Status      JobStatus              `json:"status"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
	StartedAt   *time.Time             `json:"started_at,omitempty"`
	CompletedAt *time.Time             `json:"completed_at,omitempty"`
	Context     map[string]interface{} `json:"context"`
	Results     []interface{}          `json:"results"`
	Error       string                 `json:"error,omitempty"`
}

// JobsResponse represents the response for listing jobs
type JobsResponse struct {
	Jobs      []Job     `json:"jobs"`
	Total     int       `json:"total"`
	Limit     int       `json:"limit"`
	Offset    int       `json:"offset"`
	Message   string    `json:"message"`
	Success   bool      `json:"success"`
	Timestamp time.Time `json:"timestamp"`
}

// ScheduleType represents the type of schedule
type ScheduleType string

const (
	ScheduleTypeCron     ScheduleType = "cron"
	ScheduleTypeInterval ScheduleType = "interval"
	ScheduleTypeOneTime  ScheduleType = "one_time"
)

// ScheduleStatus represents the status of a schedule
type ScheduleStatus string

const (
	ScheduleStatusActive   ScheduleStatus = "active"
	ScheduleStatusInactive ScheduleStatus = "inactive"
)

// Schedule represents a job schedule
type Schedule struct {
	ID              string                 `json:"id"`
	Name            string                 `json:"name"`
	Description     string                 `json:"description"`
	ScheduleType    ScheduleType           `json:"schedule_type"`
	Status          ScheduleStatus         `json:"status"`
	CronExpression  string                 `json:"cron_expression,omitempty"`
	IntervalSeconds int                    `json:"interval_seconds,omitempty"`
	Playbook        []interface{}          `json:"playbook"`
	Context         map[string]interface{} `json:"context"`
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
}

// SchedulesResponse represents the response for listing schedules
type SchedulesResponse struct {
	Schedules []Schedule `json:"schedules"`
	Message   string     `json:"message"`
	Success   bool       `json:"success"`
	Timestamp time.Time  `json:"timestamp"`
}

// PluginType represents the type of plugin
type PluginType string

const (
	PluginTypeLinux   PluginType = "linux"
	PluginTypeWindows PluginType = "windows"
	PluginTypePython  PluginType = "python"
	PluginTypeGo      PluginType = "go"
)

// Plugin represents a plugin
type Plugin struct {
	Name      string     `json:"name"`
	Type      PluginType `json:"type"`
	Filename  string     `json:"filename"`
	Size      int        `json:"size"`
	CreatedAt time.Time  `json:"created_at"`
}

// PluginsResponse represents the response for listing plugins
type PluginsResponse struct {
	Plugins   []Plugin  `json:"plugins"`
	Message   string    `json:"message"`
	Success   bool      `json:"success"`
	Timestamp time.Time `json:"timestamp"`
}

// UploadPluginResponse represents the response for uploading a plugin
type UploadPluginResponse struct {
	PluginName string    `json:"plugin_name"`
	PluginType string    `json:"plugin_type"`
	Filename   string    `json:"filename"`
	Message    string    `json:"message"`
	Size       int       `json:"size"`
	Success    bool      `json:"success"`
	Timestamp  time.Time `json:"timestamp"`
}

// DeletePluginResponse represents the response for deleting a plugin
type DeletePluginResponse struct {
	PluginName string    `json:"plugin_name"`
	PluginType string    `json:"plugin_type"`
	Message    string    `json:"message"`
	Success    bool      `json:"success"`
	Timestamp  time.Time `json:"timestamp"`
}

// WebhookEvent represents a webhook event
type WebhookEvent string

const (
	WebhookEventJobStarted   WebhookEvent = "job_started"
	WebhookEventJobCompleted WebhookEvent = "job_completed"
	WebhookEventJobFailed    WebhookEvent = "job_failed"
	WebhookEventJobCancelled WebhookEvent = "job_cancelled"
)

// Webhook represents a webhook configuration
type Webhook struct {
	URL        string            `json:"url"`
	Events     []WebhookEvent    `json:"events"`
	Headers    map[string]string `json:"headers,omitempty"`
	RetryCount int               `json:"retry_count,omitempty"`
	Timeout    int               `json:"timeout,omitempty"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
}

// WebhooksResponse represents the response for listing webhooks
type WebhooksResponse struct {
	Webhooks  []Webhook `json:"webhooks"`
	Message   string    `json:"message"`
	Success   bool      `json:"success"`
	Timestamp time.Time `json:"timestamp"`
}

// PlaybookExecutionRequest represents a request to execute a playbook
type PlaybookExecutionRequest struct {
	Playbook []interface{}          `json:"playbook"`
	Context  map[string]interface{} `json:"context,omitempty"`
	Options  *PlaybookOptions       `json:"options,omitempty"`
}

// PlaybookOptions represents options for playbook execution
type PlaybookOptions struct {
	Priority string `json:"priority,omitempty"`
	Timeout  int    `json:"timeout,omitempty"`
}

// PlaybookExecutionResponse represents the response for playbook execution
type PlaybookExecutionResponse struct {
	Context   map[string]interface{} `json:"context"`
	Results   []interface{}          `json:"results"`
	Success   bool                   `json:"success"`
	Timestamp time.Time              `json:"timestamp"`
}

// ValidationRequest represents a request to validate a playbook
type ValidationRequest struct {
	Playbook []interface{}          `json:"playbook"`
	Context  map[string]interface{} `json:"context,omitempty"`
}

// ValidationResponse represents the response for validation
type ValidationResponse struct {
	Valid     bool              `json:"valid"`
	Message   string            `json:"message"`
	Success   bool              `json:"success"`
	Timestamp time.Time         `json:"timestamp"`
	Errors    []ValidationError `json:"errors,omitempty"`
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Value   string `json:"value"`
}

// APIError represents an API error response
type APIError struct {
	Error string `json:"error"`
}

// GetHealth checks the health status of the SecAuto API
func (c *Client) GetHealth() (*HealthResponse, error) {
	resp, err := c.makeRequest("GET", "/health", nil)
	if err != nil {
		return nil, err
	}

	var health HealthResponse
	if err := json.Unmarshal(resp, &health); err != nil {
		return nil, fmt.Errorf("failed to unmarshal health response: %w", err)
	}

	return &health, nil
}

// GetAutomations retrieves all automation scripts
func (c *Client) GetAutomations() (*AutomationsResponse, error) {
	resp, err := c.makeRequest("GET", "/automations", nil)
	if err != nil {
		return nil, err
	}

	var automations AutomationsResponse
	if err := json.Unmarshal(resp, &automations); err != nil {
		return nil, fmt.Errorf("failed to unmarshal automations response: %w", err)
	}

	return &automations, nil
}

// UploadAutomation uploads a new automation script
func (c *Client) UploadAutomation(filePath string) (*UploadAutomationResponse, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	part, err := writer.CreateFormFile("automation", filepath.Base(filePath))
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, fmt.Errorf("failed to copy file: %w", err)
	}

	writer.Close()

	resp, err := c.makeMultipartRequest("POST", "/automation", &buf, writer.FormDataContentType())
	if err != nil {
		return nil, err
	}

	var uploadResp UploadAutomationResponse
	if err := json.Unmarshal(resp, &uploadResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal upload response: %w", err)
	}

	return &uploadResp, nil
}

// DeleteAutomation deletes an automation script
func (c *Client) DeleteAutomation(name string) (*DeleteAutomationResponse, error) {
	resp, err := c.makeRequest("DELETE", fmt.Sprintf("/automation/%s", name), nil)
	if err != nil {
		return nil, err
	}

	var deleteResp DeleteAutomationResponse
	if err := json.Unmarshal(resp, &deleteResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal delete response: %w", err)
	}

	return &deleteResp, nil
}

// GetIntegrations retrieves all integrations
func (c *Client) GetIntegrations() (*IntegrationsResponse, error) {
	resp, err := c.makeRequest("GET", "/integrations", nil)
	if err != nil {
		return nil, err
	}

	var integrations IntegrationsResponse
	if err := json.Unmarshal(resp, &integrations); err != nil {
		return nil, fmt.Errorf("failed to unmarshal integrations response: %w", err)
	}

	return &integrations, nil
}

// GetIntegration retrieves a specific integration
func (c *Client) GetIntegration(name string) (*IntegrationResponse, error) {
	resp, err := c.makeRequest("GET", fmt.Sprintf("/integrations/%s", name), nil)
	if err != nil {
		return nil, err
	}

	var integration IntegrationResponse
	if err := json.Unmarshal(resp, &integration); err != nil {
		return nil, fmt.Errorf("failed to unmarshal integration response: %w", err)
	}

	return &integration, nil
}

// CreateIntegration creates a new integration
func (c *Client) CreateIntegration(integration Integration) (*IntegrationResponse, error) {
	body, err := json.Marshal(integration)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal integration: %w", err)
	}

	resp, err := c.makeRequest("POST", "/integrations", body)
	if err != nil {
		return nil, err
	}

	var integrationResp IntegrationResponse
	if err := json.Unmarshal(resp, &integrationResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal integration response: %w", err)
	}

	return &integrationResp, nil
}

// UpdateIntegration updates an existing integration
func (c *Client) UpdateIntegration(name string, integration Integration) (*IntegrationResponse, error) {
	body, err := json.Marshal(integration)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal integration: %w", err)
	}

	resp, err := c.makeRequest("PUT", fmt.Sprintf("/integrations/%s", name), body)
	if err != nil {
		return nil, err
	}

	var integrationResp IntegrationResponse
	if err := json.Unmarshal(resp, &integrationResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal integration response: %w", err)
	}

	return &integrationResp, nil
}

// DeleteIntegration deletes an integration
func (c *Client) DeleteIntegration(name string) error {
	_, err := c.makeRequest("DELETE", fmt.Sprintf("/integrations/%s", name), nil)
	return err
}

// UploadIntegration uploads a new integration file
func (c *Client) UploadIntegration(filePath string) (*UploadPluginResponse, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	part, err := writer.CreateFormFile("integration", filepath.Base(filePath))
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, fmt.Errorf("failed to copy file: %w", err)
	}

	writer.Close()

	resp, err := c.makeMultipartRequest("POST", "/integrations/upload", &buf, writer.FormDataContentType())
	if err != nil {
		return nil, err
	}

	var uploadResp UploadPluginResponse
	if err := json.Unmarshal(resp, &uploadResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal upload response: %w", err)
	}

	return &uploadResp, nil
}

// DeleteIntegrationFile deletes an integration file
func (c *Client) DeleteIntegrationFile(name string) (*DeletePluginResponse, error) {
	resp, err := c.makeRequest("DELETE", fmt.Sprintf("/integrations/delete/%s", name), nil)
	if err != nil {
		return nil, err
	}

	var deleteResp DeletePluginResponse
	if err := json.Unmarshal(resp, &deleteResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal delete response: %w", err)
	}

	return &deleteResp, nil
}

// GetPlaybooks retrieves all playbooks
func (c *Client) GetPlaybooks() (*PlaybooksResponse, error) {
	resp, err := c.makeRequest("GET", "/playbooks", nil)
	if err != nil {
		return nil, err
	}

	var playbooks PlaybooksResponse
	if err := json.Unmarshal(resp, &playbooks); err != nil {
		return nil, fmt.Errorf("failed to unmarshal playbooks response: %w", err)
	}

	return &playbooks, nil
}

// UploadPlaybook uploads a new playbook file
func (c *Client) UploadPlaybook(filePath string) (*UploadPlaybookResponse, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	part, err := writer.CreateFormFile("playbook", filepath.Base(filePath))
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, fmt.Errorf("failed to copy file: %w", err)
	}

	writer.Close()

	resp, err := c.makeMultipartRequest("POST", "/playbook/upload", &buf, writer.FormDataContentType())
	if err != nil {
		return nil, err
	}

	var uploadResp UploadPlaybookResponse
	if err := json.Unmarshal(resp, &uploadResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal upload response: %w", err)
	}

	return &uploadResp, nil
}

// DeletePlaybook deletes a playbook
func (c *Client) DeletePlaybook(name string) (*DeletePlaybookResponse, error) {
	resp, err := c.makeRequest("DELETE", fmt.Sprintf("/playbook/%s", name), nil)
	if err != nil {
		return nil, err
	}

	var deleteResp DeletePlaybookResponse
	if err := json.Unmarshal(resp, &deleteResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal delete response: %w", err)
	}

	return &deleteResp, nil
}

// ExecutePlaybook executes a playbook synchronously
func (c *Client) ExecutePlaybook(req PlaybookExecutionRequest) (*PlaybookExecutionResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.makeRequest("POST", "/playbook", body)
	if err != nil {
		return nil, err
	}

	var executionResp PlaybookExecutionResponse
	if err := json.Unmarshal(resp, &executionResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal execution response: %w", err)
	}

	return &executionResp, nil
}

// ExecutePlaybookAsync executes a playbook asynchronously
func (c *Client) ExecutePlaybookAsync(req PlaybookExecutionRequest) error {
	body, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	_, err = c.makeRequest("POST", "/playbook/async", body)
	return err
}

// GetJobs retrieves all jobs with optional filtering
func (c *Client) GetJobs(status JobStatus, limit, offset int) (*JobsResponse, error) {
	url := "/jobs?"
	if status != "" {
		url += fmt.Sprintf("status=%s&", status)
	}
	if limit > 0 {
		url += fmt.Sprintf("limit=%d&", limit)
	}
	if offset > 0 {
		url += fmt.Sprintf("offset=%d&", offset)
	}

	resp, err := c.makeRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var jobs JobsResponse
	if err := json.Unmarshal(resp, &jobs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal jobs response: %w", err)
	}

	return &jobs, nil
}

// GetSchedules retrieves all schedules
func (c *Client) GetSchedules() (*SchedulesResponse, error) {
	resp, err := c.makeRequest("GET", "/schedules", nil)
	if err != nil {
		return nil, err
	}

	var schedules SchedulesResponse
	if err := json.Unmarshal(resp, &schedules); err != nil {
		return nil, fmt.Errorf("failed to unmarshal schedules response: %w", err)
	}

	return &schedules, nil
}

// CreateSchedule creates a new schedule
func (c *Client) CreateSchedule(schedule Schedule) error {
	body, err := json.Marshal(schedule)
	if err != nil {
		return fmt.Errorf("failed to marshal schedule: %w", err)
	}

	_, err = c.makeRequest("POST", "/schedules", body)
	return err
}

// GetPlugins retrieves all plugins
func (c *Client) GetPlugins() (*PluginsResponse, error) {
	resp, err := c.makeRequest("GET", "/plugins", nil)
	if err != nil {
		return nil, err
	}

	var plugins PluginsResponse
	if err := json.Unmarshal(resp, &plugins); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugins response: %w", err)
	}

	return &plugins, nil
}

// UploadPlugin uploads a new plugin file
func (c *Client) UploadPlugin(pluginType PluginType, filePath string) (*UploadPluginResponse, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	part, err := writer.CreateFormFile("plugin", filepath.Base(filePath))
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, fmt.Errorf("failed to copy file: %w", err)
	}

	writer.Close()

	resp, err := c.makeMultipartRequest("POST", fmt.Sprintf("/plugin/%s", pluginType), &buf, writer.FormDataContentType())
	if err != nil {
		return nil, err
	}

	var uploadResp UploadPluginResponse
	if err := json.Unmarshal(resp, &uploadResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal upload response: %w", err)
	}

	return &uploadResp, nil
}

// DeletePlugin deletes a plugin
func (c *Client) DeletePlugin(pluginType PluginType, name string) (*DeletePluginResponse, error) {
	resp, err := c.makeRequest("DELETE", fmt.Sprintf("/plugin/%s/%s", pluginType, name), nil)
	if err != nil {
		return nil, err
	}

	var deleteResp DeletePluginResponse
	if err := json.Unmarshal(resp, &deleteResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal delete response: %w", err)
	}

	return &deleteResp, nil
}

// ValidatePlaybook validates a playbook without executing
func (c *Client) ValidatePlaybook(req ValidationRequest) (*ValidationResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.makeRequest("POST", "/validate", body)
	if err != nil {
		return nil, err
	}

	var validationResp ValidationResponse
	if err := json.Unmarshal(resp, &validationResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal validation response: %w", err)
	}

	return &validationResp, nil
}

// ConfigureWebhook configures webhook notifications
func (c *Client) ConfigureWebhook(webhook Webhook) error {
	body, err := json.Marshal(webhook)
	if err != nil {
		return fmt.Errorf("failed to marshal webhook: %w", err)
	}

	_, err = c.makeRequest("POST", "/webhooks", body)
	return err
}

// GetCluster retrieves cluster information
func (c *Client) GetCluster() (map[string]interface{}, error) {
	resp, err := c.makeRequest("GET", "/cluster", nil)
	if err != nil {
		return nil, err
	}

	var cluster map[string]interface{}
	if err := json.Unmarshal(resp, &cluster); err != nil {
		return nil, fmt.Errorf("failed to unmarshal cluster response: %w", err)
	}

	return cluster, nil
}

// GetJobMetrics retrieves database performance metrics
func (c *Client) GetJobMetrics() (map[string]interface{}, error) {
	resp, err := c.makeRequest("GET", "/jobs/metrics", nil)
	if err != nil {
		return nil, err
	}

	var metrics map[string]interface{}
	if err := json.Unmarshal(resp, &metrics); err != nil {
		return nil, fmt.Errorf("failed to unmarshal metrics response: %w", err)
	}

	return metrics, nil
}

// GetJobStats retrieves comprehensive job statistics
func (c *Client) GetJobStats() (map[string]interface{}, error) {
	resp, err := c.makeRequest("GET", "/jobs/stats", nil)
	if err != nil {
		return nil, err
	}

	var stats map[string]interface{}
	if err := json.Unmarshal(resp, &stats); err != nil {
		return nil, fmt.Errorf("failed to unmarshal stats response: %w", err)
	}

	return stats, nil
}

// makeRequest performs an HTTP request to the API
func (c *Client) makeRequest(method, path string, body []byte) ([]byte, error) {
	url := c.BaseURL + path
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if c.APIKey != "" {
		req.Header.Set("X-API-Key", c.APIKey)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		var apiError APIError
		if err := json.Unmarshal(respBody, &apiError); err != nil {
			return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(respBody))
		}
		return nil, fmt.Errorf("API request failed: %s", apiError.Error)
	}

	return respBody, nil
}

// makeMultipartRequest performs a multipart HTTP request to the API
func (c *Client) makeMultipartRequest(method, path string, body *bytes.Buffer, contentType string) ([]byte, error) {
	url := c.BaseURL + path
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", contentType)
	if c.APIKey != "" {
		req.Header.Set("X-API-Key", c.APIKey)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		var apiError APIError
		if err := json.Unmarshal(respBody, &apiError); err != nil {
			return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(respBody))
		}
		return nil, fmt.Errorf("API request failed: %s", apiError.Error)
	}

	return respBody, nil
}
