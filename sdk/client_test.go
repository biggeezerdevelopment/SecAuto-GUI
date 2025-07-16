package sdk

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient("http://localhost:8000")
	if client == nil {
		t.Fatal("NewClient returned nil")
	}
	if client.BaseURL != "http://localhost:8000" {
		t.Errorf("Expected BaseURL to be http://localhost:8000, got %s", client.BaseURL)
	}
}

func TestNewDefaultClient(t *testing.T) {
	client := NewDefaultClient()
	if client == nil {
		t.Fatal("NewDefaultClient returned nil")
	}
	if client.BaseURL != "http://localhost:8000" {
		t.Errorf("Expected BaseURL to be http://localhost:8000, got %s", client.BaseURL)
	}
}

func TestNewClientWithAuth(t *testing.T) {
	client := NewClientWithAuth("http://localhost:8000", "secauto-api-key-2024-07-14")
	if client == nil {
		t.Fatal("NewClientWithAuth returned nil")
	}
	if client.BaseURL != "http://localhost:8000" {
		t.Errorf("Expected BaseURL to be http://localhost:8000, got %s", client.BaseURL)
	}
	if client.APIKey != "secauto-api-key-2024-07-14" {
		t.Errorf("Expected APIKey to be test-api-key, got %s", client.APIKey)
	}
}

func TestJobStatusConstants(t *testing.T) {
	statuses := []JobStatus{
		JobStatusPending,
		JobStatusRunning,
		JobStatusCompleted,
		JobStatusFailed,
		JobStatusCancelled,
	}

	for _, status := range statuses {
		if status == "" {
			t.Errorf("JobStatus constant is empty")
		}
	}
}

func TestScheduleTypeConstants(t *testing.T) {
	types := []ScheduleType{
		ScheduleTypeCron,
		ScheduleTypeInterval,
		ScheduleTypeOneTime,
	}

	for _, scheduleType := range types {
		if scheduleType == "" {
			t.Errorf("ScheduleType constant is empty")
		}
	}
}

func TestScheduleStatusConstants(t *testing.T) {
	statuses := []ScheduleStatus{
		ScheduleStatusActive,
		ScheduleStatusInactive,
	}

	for _, status := range statuses {
		if status == "" {
			t.Errorf("ScheduleStatus constant is empty")
		}
	}
}

func TestPluginTypeConstants(t *testing.T) {
	types := []PluginType{
		PluginTypeLinux,
		PluginTypeWindows,
		PluginTypePython,
		PluginTypeGo,
	}

	for _, pluginType := range types {
		if pluginType == "" {
			t.Errorf("PluginType constant is empty")
		}
	}
}

func TestWebhookEventConstants(t *testing.T) {
	events := []WebhookEvent{
		WebhookEventJobStarted,
		WebhookEventJobCompleted,
		WebhookEventJobFailed,
		WebhookEventJobCancelled,
	}

	for _, event := range events {
		if event == "" {
			t.Errorf("WebhookEvent constant is empty")
		}
	}
}
