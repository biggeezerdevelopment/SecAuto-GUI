# SecAuto SDK

A Go SDK for interacting with the SecAuto SOAR Rules Engine API.

## Features

- Complete API coverage matching the OpenAPI specification
- Type-safe client with proper error handling
- Support for file uploads (automations, integrations, playbooks, plugins)
- Authentication via X-API-Key header
- Comprehensive response types matching the API specification
- Configuration support via YAML file

## Installation

```bash
go get secauto-gui/sdk
```

## Usage

### Creating a Client

```go
import "secauto-gui/sdk"

// Create a client for the default remote server
client := sdk.NewDefaultClient()

// Create a client with custom URL
client := sdk.NewClient("http://localhost:8000")

// Create a client with authentication
client := sdk.NewClientWithAuth("http://localhost:8000", "your-api-key")

// Create a client from configuration
client := sdk.NewClientFromConfig("http://localhost:8000", "your-api-key")
```

### Configuration

The SDK supports configuration through a `config.yaml` file:

```yaml
server:
  port: "8080"
  host: "localhost"

remote:
  url: "http://localhost:8000"
  api_key: "your-api-key-here"

logging:
  level: "info"
  file: "logs/secauto-gui.log"
```

The application will automatically:
- Load configuration from `config.yaml`
- Create a default configuration file if it doesn't exist
- Use the API key for authentication with the remote server

### Health Check

```go
health, err := client.GetHealth()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Status: %s, Version: %s\n", health.Status, health.Version)
```

### Automations

```go
// Get all automations
automations, err := client.GetAutomations()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Found %d automations\n", automations.Count)

// Upload an automation script
resp, err := client.UploadAutomation("/path/to/script.py")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Uploaded: %s\n", resp.AutomationName)

// Delete an automation
resp, err := client.DeleteAutomation("script_name")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Deleted: %s\n", resp.AutomationName)
```

### Integrations

```go
// Get all integrations
integrations, err := client.GetIntegrations()
if err != nil {
    log.Fatal(err)
}

// Create an integration
integration := sdk.Integration{
    Name:        "virustotal",
    Type:        "virustotal",
    Description: "VirusTotal integration",
    Enabled:     true,
    URL:         "https://api.virustotal.com/v3",
    APIKey:      "your-api-key",
}

resp, err := client.CreateIntegration(integration)
if err != nil {
    log.Fatal(err)
}

// Upload integration file
resp, err := client.UploadIntegration("/path/to/integration.py")
if err != nil {
    log.Fatal(err)
}
```

### Playbooks

```go
// Get all playbooks
playbooks, err := client.GetPlaybooks()
if err != nil {
    log.Fatal(err)
}

// Upload a playbook
resp, err := client.UploadPlaybook("/path/to/playbook.json")
if err != nil {
    log.Fatal(err)
}

// Execute a playbook synchronously
req := sdk.PlaybookExecutionRequest{
    Playbook: []interface{}{
        map[string]interface{}{
            "action": "log",
            "message": "Hello from SDK",
        },
    },
    Context: map[string]interface{}{
        "source": "sdk_example",
    },
    Options: &sdk.PlaybookOptions{
        Priority: "normal",
        Timeout:  30,
    },
}

execution, err := client.ExecutePlaybook(req)
if err != nil {
    log.Fatal(err)
}

// Execute a playbook asynchronously
err = client.ExecutePlaybookAsync(req)
if err != nil {
    log.Fatal(err)
}
```

### Jobs

```go
// Get jobs with filtering
jobs, err := client.GetJobs(sdk.JobStatusRunning, 10, 0)
if err != nil {
    log.Fatal(err)
}

// Get job metrics
metrics, err := client.GetJobMetrics()
if err != nil {
    log.Fatal(err)
}

// Get job statistics
stats, err := client.GetJobStats()
if err != nil {
    log.Fatal(err)
}
```

### Schedules

```go
// Get all schedules
schedules, err := client.GetSchedules()
if err != nil {
    log.Fatal(err)
}

// Create a schedule
schedule := sdk.Schedule{
    Name:        "Daily Security Scan",
    Description: "Runs security scan every day at 2 AM",
    ScheduleType: sdk.ScheduleTypeCron,
    Status:      sdk.ScheduleStatusActive,
    CronExpression: "0 2 * * *",
    Playbook: []interface{}{
        map[string]interface{}{
            "action": "security_scan",
            "targets": []string{"192.168.1.0/24"},
        },
    },
    Context: map[string]interface{}{
        "scan_type": "vulnerability",
    },
}

err = client.CreateSchedule(schedule)
if err != nil {
    log.Fatal(err)
}
```

### Plugins

```go
// Get all plugins
plugins, err := client.GetPlugins()
if err != nil {
    log.Fatal(err)
}

// Upload a plugin
resp, err := client.UploadPlugin(sdk.PluginTypePython, "/path/to/plugin.py")
if err != nil {
    log.Fatal(err)
}

// Delete a plugin
resp, err := client.DeletePlugin(sdk.PluginTypePython, "plugin_name")
if err != nil {
    log.Fatal(err)
}
```

### Validation

```go
// Validate a playbook
req := sdk.ValidationRequest{
    Playbook: []interface{}{
        map[string]interface{}{
            "action": "log",
            "message": "Test message",
        },
    },
    Context: map[string]interface{}{
        "test": true,
    },
}

validation, err := client.ValidatePlaybook(req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Valid: %v\n", validation.Valid)
```

### Webhooks

```go
// Configure a webhook
webhook := sdk.Webhook{
    URL: "https://example.com/webhook",
    Events: []sdk.WebhookEvent{
        sdk.WebhookEventJobStarted,
        sdk.WebhookEventJobCompleted,
        sdk.WebhookEventJobFailed,
    },
    Headers: map[string]string{
        "Authorization": "Bearer token123",
    },
    RetryCount: 3,
    Timeout:    30,
}

err = client.ConfigureWebhook(webhook)
if err != nil {
    log.Fatal(err)
}
```

### Cluster Information

```go
// Get cluster information
cluster, err := client.GetCluster()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Cluster info: %+v\n", cluster)
```

## Error Handling

The SDK returns proper errors for all API calls. Check the error response for details:

```go
resp, err := client.GetAutomations()
if err != nil {
    // Handle error
    log.Printf("API Error: %v", err)
    return
}
```

## Constants

The SDK provides constants for various enums:

### Job Status
- `sdk.JobStatusPending`
- `sdk.JobStatusRunning`
- `sdk.JobStatusCompleted`
- `sdk.JobStatusFailed`
- `sdk.JobStatusCancelled`

### Schedule Types
- `sdk.ScheduleTypeCron`
- `sdk.ScheduleTypeInterval`
- `sdk.ScheduleTypeOneTime`

### Schedule Status
- `sdk.ScheduleStatusActive`
- `sdk.ScheduleStatusInactive`

### Plugin Types
- `sdk.PluginTypeLinux`
- `sdk.PluginTypeWindows`
- `sdk.PluginTypePython`
- `sdk.PluginTypeGo`

### Webhook Events
- `sdk.WebhookEventJobStarted`
- `sdk.WebhookEventJobCompleted`
- `sdk.WebhookEventJobFailed`
- `sdk.WebhookEventJobCancelled`

## API Endpoints

The SDK supports all endpoints from the SecAuto OpenAPI specification:

- **Health**: `/health`
- **Automations**: `/automation`, `/automations`
- **Integrations**: `/integrations`, `/integrations/upload`, `/integrations/delete/:name`
- **Playbooks**: `/playbooks`, `/playbook/upload`, `/playbook`, `/playbook/async`
- **Jobs**: `/jobs`, `/jobs/metrics`, `/jobs/stats`
- **Schedules**: `/schedules`
- **Plugins**: `/plugins`, `/plugin/:type`
- **Validation**: `/validate`
- **Webhooks**: `/webhooks`
- **Cluster**: `/cluster`

## Configuration

The SDK integrates with the application's configuration system. See `CONFIGURATION.md` for detailed information about:

- Configuration file format
- API key management
- Server settings
- Logging configuration

## Testing

Run the tests to verify the SDK functionality:

```bash
go test ./sdk
```

## Example

See `example.go` for a complete example of using the SDK. 