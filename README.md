# SecAuto GUI - Automation Framework

A modern, dark-themed web frontend for the SecAuto automation framework that connects to a remote SecAuto server. Features HTMX integration and a comprehensive Go SDK for interacting with the remote API.

## Features

- 🎨 **Dark Theme**: Modern, professional dark-themed UI using Bootstrap 5
- ⚡ **HTMX Integration**: Real-time updates without heavy JavaScript
- 📊 **Real-time Dashboard**: Live metrics and automation status
- 🔧 **Automation Management**: Create, run, monitor, and delete automations
- 📝 **Log Management**: Real-time log viewing with filtering
- 📈 **System Metrics**: CPU, memory, and automation statistics
- 🔌 **Go SDK**: Complete SDK library for programmatic access
- 📚 **Swagger Documentation**: Interactive API documentation

## Quick Start

### Prerequisites

- Go 1.21 or later
- Git
- Remote SecAuto server running at http://localhost:8000

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd SecAuto-Gui
```

2. Install dependencies:
```bash
go mod tidy
```

3. Generate Swagger documentation:
```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init
```

4. Run the application:
```bash
go run .
```

5. Open your browser and navigate to:
   - **Web Interface**: http://localhost:8080
   - **Remote API Documentation**: http://localhost:8000/api-docs

## API Endpoints

### Health Check
- `GET /api/v1/health` - Check API health status

### Automations
- `GET /api/v1/automations` - List all automations
- `POST /api/v1/automations` - Create new automation
- `GET /api/v1/automations/{id}` - Get automation details
- `PUT /api/v1/automations/{id}` - Update automation
- `DELETE /api/v1/automations/{id}` - Delete automation
- `POST /api/v1/automations/{id}/run` - Run automation
- `GET /api/v1/automations/{id}/status` - Get automation status

### Logs & Metrics
- `GET /api/v1/logs` - Get system logs (with optional filtering)
- `GET /api/v1/metrics` - Get system metrics

## Go SDK Usage

The SecAuto SDK provides a simple and intuitive way to interact with the API programmatically.

### Basic Usage

```go
package main

import (
    "fmt"
    "log"
    "secauto-gui/sdk"
)

func main() {
    // Create a new client for the remote SecAuto server
    client := sdk.NewDefaultClient()

    // Check API health
    health, err := client.GetHealth()
    if err != nil {
        log.Fatalf("Failed to check health: %v", err)
    }
    fmt.Printf("API Status: %s\n", health.Status)

    // Get all automations
    automations, err := client.GetAutomations()
    if err != nil {
        log.Fatalf("Failed to get automations: %v", err)
    }
    fmt.Printf("Found %d automations\n", len(automations))

    // Create a new automation
    req := sdk.CreateAutomationRequest{
        Name:        "Security Scan",
        Description: "Automated security scanning workflow",
        Config: map[string]interface{}{
            "scan_type": "vulnerability",
            "targets":   []string{"192.168.1.0/24"},
        },
    }

    automation, err := client.CreateAutomation(req)
    if err != nil {
        log.Fatalf("Failed to create automation: %v", err)
    }
    fmt.Printf("Created automation: %s\n", automation.ID)

    // Run the automation
    err = client.RunAutomation(automation.ID)
    if err != nil {
        log.Fatalf("Failed to run automation: %v", err)
    }

    // Get system metrics
    metrics, err := client.GetMetrics()
    if err != nil {
        log.Fatalf("Failed to get metrics: %v", err)
    }
    fmt.Printf("CPU Usage: %.1f%%\n", metrics.CPUUsage)
}
```

### Authentication

For authenticated requests, use the `NewClientWithAuth` function:

```go
client := sdk.NewClientWithAuth("http://localhost:8080", "your-api-key")
```

## Web Interface Features

### Dashboard
- Real-time system metrics
- Recent automation activity
- System status overview
- Auto-refreshing data

### Automation Management
- Create new automations with custom configurations
- View automation details and status
- Run automations with confirmation
- Delete automations with safety prompts
- Real-time status monitoring

### Log Management
- Real-time log viewing
- Filter by log level (INFO, WARN, ERROR)
- Filter by automation ID
- Auto-refresh toggle
- Color-coded log entries

### System Metrics
- CPU and memory usage
- Automation statistics
- Performance trends
- Real-time updates

## HTMX Integration

The web interface uses HTMX for seamless, dynamic updates without heavy JavaScript:

- **Real-time Updates**: Automatic refresh of data every 10-30 seconds
- **Progressive Enhancement**: Works without JavaScript, enhanced with HTMX
- **Server-side Rendering**: All UI updates handled server-side
- **Declarative Interactions**: HTMX attributes define behavior directly in HTML

### Key HTMX Features Used

- `hx-get` - Fetch data from server
- `hx-post` - Submit forms and trigger actions
- `hx-target` - Specify where to insert responses
- `hx-swap` - Control how content is swapped
- `hx-trigger` - Define when requests are sent
- `hx-confirm` - User confirmation dialogs
- `hx-push-url` - Update browser URL without page reload

## Development

### Project Structure

```
SecAuto-Gui/
├── main.go              # Application entry point
├── handlers.go          # API handlers with Swagger docs
├── web_handlers.go      # Web frontend handlers
├── go.mod              # Go module dependencies
├── templates/           # HTML templates
│   ├── base.html       # Base template with dark theme
│   ├── dashboard.html  # Dashboard page
│   ├── automations.html # Automation management
│   ├── logs.html       # Log viewing
│   ├── metrics.html    # System metrics
│   └── automation_detail.html # Automation details
├── sdk/                # Go SDK library
│   ├── client.go       # SDK client implementation
│   └── example.go      # Usage examples
└── docs/               # Generated Swagger docs
```

### Adding New Features

1. **API Endpoints**: Add handlers in `handlers.go` with Swagger documentation
2. **Web Pages**: Create new templates in `templates/` directory
3. **SDK Methods**: Add corresponding methods in `sdk/client.go`
4. **Frontend Routes**: Add routes in `main.go` and handlers in `web_handlers.go`

### Building

```bash
# Build the application
go build -o secauto-gui .

# Run with custom configuration
./secauto-gui
```

## Configuration

The application can be configured through environment variables:

- `PORT` - Server port (default: 8080)
- `HOST` - Server host (default: localhost)
- `LOG_LEVEL` - Logging level (default: info)

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For support and questions:
- Create an issue on GitHub
- Check the API documentation at http://localhost:8080/swagger/index.html
- Review the SDK examples in `sdk/example.go` 