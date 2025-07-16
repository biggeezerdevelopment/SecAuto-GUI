package sdk

import (
	"fmt"
	"log"
)

// Example demonstrates how to use the SecAuto SDK
func Example() {
	// Create a new client for the remote SecAuto server
	client := NewDefaultClient()

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
	fmt.Printf("Found %d automations\n", automations.Count)

	// Get all integrations
	integrations, err := client.GetIntegrations()
	if err != nil {
		log.Fatalf("Failed to get integrations: %v", err)
	}
	fmt.Printf("Found %d integrations\n", len(integrations.Integrations))

	// Get all playbooks
	playbooks, err := client.GetPlaybooks()
	if err != nil {
		log.Fatalf("Failed to get playbooks: %v", err)
	}
	fmt.Printf("Found %d playbooks\n", playbooks.Count)

	// Get jobs with filtering
	jobs, err := client.GetJobs("", 10, 0)
	if err != nil {
		log.Fatalf("Failed to get jobs: %v", err)
	}
	fmt.Printf("Found %d jobs\n", len(jobs.Jobs))

	// Get schedules
	schedules, err := client.GetSchedules()
	if err != nil {
		log.Fatalf("Failed to get schedules: %v", err)
	}
	fmt.Printf("Found %d schedules\n", len(schedules.Schedules))

	// Get plugins
	plugins, err := client.GetPlugins()
	if err != nil {
		log.Fatalf("Failed to get plugins: %v", err)
	}
	fmt.Printf("Found %d plugins\n", len(plugins.Plugins))

	// Execute a playbook synchronously
	playbookReq := PlaybookExecutionRequest{
		Playbook: []interface{}{
			map[string]interface{}{
				"action":  "log",
				"message": "Hello from SDK example",
			},
		},
		Context: map[string]interface{}{
			"source": "sdk_example",
		},
		Options: &PlaybookOptions{
			Priority: "normal",
			Timeout:  30,
		},
	}

	execution, err := client.ExecutePlaybook(playbookReq)
	if err != nil {
		log.Fatalf("Failed to execute playbook: %v", err)
	}
	fmt.Printf("Playbook executed successfully: %v\n", execution.Success)

	// Validate a playbook
	validationReq := ValidationRequest{
		Playbook: []interface{}{
			map[string]interface{}{
				"action":  "log",
				"message": "Test message",
			},
		},
		Context: map[string]interface{}{
			"test": true,
		},
	}

	validation, err := client.ValidatePlaybook(validationReq)
	if err != nil {
		log.Fatalf("Failed to validate playbook: %v", err)
	}
	fmt.Printf("Playbook validation: %v\n", validation.Valid)

	// Get cluster information
	cluster, err := client.GetCluster()
	if err != nil {
		log.Fatalf("Failed to get cluster info: %v", err)
	}
	fmt.Printf("Cluster info: %+v\n", cluster)

	// Get job metrics
	metrics, err := client.GetJobMetrics()
	if err != nil {
		log.Fatalf("Failed to get job metrics: %v", err)
	}
	fmt.Printf("Job metrics: %+v\n", metrics)

	// Get job statistics
	stats, err := client.GetJobStats()
	if err != nil {
		log.Fatalf("Failed to get job stats: %v", err)
	}
	fmt.Printf("Job stats: %+v\n", stats)

	// Configure a webhook
	webhook := Webhook{
		URL: "https://example.com/webhook",
		Events: []WebhookEvent{
			WebhookEventJobStarted,
			WebhookEventJobCompleted,
			WebhookEventJobFailed,
		},
		Headers: map[string]string{
			"Authorization": "Bearer token123",
		},
		RetryCount: 3,
		Timeout:    30,
	}

	err = client.ConfigureWebhook(webhook)
	if err != nil {
		log.Fatalf("Failed to configure webhook: %v", err)
	}
	fmt.Printf("Webhook configured successfully\n")

	// Create a schedule
	schedule := Schedule{
		Name:           "Daily Security Scan",
		Description:    "Runs security scan every day at 2 AM",
		ScheduleType:   ScheduleTypeCron,
		Status:         ScheduleStatusActive,
		CronExpression: "0 2 * * *",
		Playbook: []interface{}{
			map[string]interface{}{
				"action":  "security_scan",
				"targets": []string{"192.168.1.0/24"},
			},
		},
		Context: map[string]interface{}{
			"scan_type": "vulnerability",
		},
	}

	err = client.CreateSchedule(schedule)
	if err != nil {
		log.Fatalf("Failed to create schedule: %v", err)
	}
	fmt.Printf("Schedule created successfully\n")

	fmt.Println("SDK example completed successfully!")
}
