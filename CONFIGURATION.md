# SecAuto GUI Configuration

The SecAuto GUI uses a YAML configuration file to manage settings for the server, remote API connection, and logging.

## Configuration File

The configuration file is located at `config.yaml` in the root directory of the application.

### Default Configuration

```yaml
server:
  port: "8080"
  host: "localhost"

remote:
  url: "http://localhost:8000"
  api_key: "secauto-api-key-2024-07-14"

logging:
  level: "info"
  file: "logs/secauto-gui.log"
```

## Configuration Options

### Server Configuration

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `server.port` | string | "8080" | Port for the GUI server |
| `server.host` | string | "localhost" | Host for the GUI server |

### Remote API Configuration

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `remote.url` | string | "http://localhost:8000" | URL of the remote SecAuto server |
| `remote.api_key` | string | "" | API key for authentication with the remote server |

### Logging Configuration

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `logging.level` | string | "info" | Log level (debug, info, warn, error) |
| `logging.file` | string | "logs/secauto-gui.log" | Log file path |

## API Key Configuration

### Setting the API Key

1. **Edit the config.yaml file:**
   ```yaml
   remote:
     url: "http://localhost:8000"
     api_key: "your-actual-api-key-here"
   ```

2. **Restart the application** for the changes to take effect.

### API Key Security

- The API key is stored in plain text in the configuration file
- The configuration endpoint (`/api/config`) returns a sanitized version without the actual API key
- Consider using environment variables for production deployments

### Environment Variables (Optional)

You can also set the API key using environment variables:

```bash
export SECAUTO_API_KEY="your-api-key-here"
export SECAUTO_REMOTE_URL="http://localhost:8000"
```

## Configuration Endpoint

The application provides a configuration endpoint at `/api/config` that returns the current configuration (without sensitive data):

```bash
curl http://localhost:8080/api/config
```

Response:
```json
{
  "server": {
    "port": "8080",
    "host": "localhost"
  },
  "remote": {
    "url": "http://localhost:8000",
    "api_key_configured": true
  },
  "logging": {
    "level": "info",
    "file": "logs/secauto-gui.log"
  }
}
```

## Automatic Configuration Creation

If the `config.yaml` file doesn't exist, the application will automatically create a default configuration file with the following structure:

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

## Configuration Validation

The application validates the configuration on startup and will:

1. **Load the configuration file** if it exists
2. **Create a default configuration** if the file doesn't exist
3. **Use hardcoded defaults** if configuration loading fails
4. **Log warnings** for any configuration issues

## Logging

Configuration-related messages are logged during startup:

```
SecAuto service initialized - connecting to remote server at http://localhost:8000
Using API key: secauto-...
```

or

```
SecAuto service initialized - connecting to remote server at http://localhost:8000
No API key configured
```

## Troubleshooting

### Common Issues

1. **Configuration file not found:**
   - The application will create a default `config.yaml` file
   - Check the logs for any warnings

2. **Invalid YAML syntax:**
   - The application will use hardcoded defaults
   - Check the `config.yaml` file for syntax errors

3. **API key not working:**
   - Verify the API key in the configuration file
   - Check that the remote server is accessible
   - Ensure the API key has the correct permissions

### Debugging

To enable debug logging, set the log level in the configuration:

```yaml
logging:
  level: "debug"
```

This will provide more detailed information about configuration loading and API requests. 