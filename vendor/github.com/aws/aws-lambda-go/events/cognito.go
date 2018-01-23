// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.

package events

// CognitoEvent contains data from an event sent from AWS Cognito
type CognitoEvent struct {
	DatasetName    string                          `json:"datasetName"`
	DatasetRecords map[string]CognitoDatasetRecord `json:"datasetRecords"`
	EventType      string                          `json:"eventType"`
	IdentityID     string                          `json:"identityId"`
	IdentityPoolID string                          `json:"identityPoolId"`
	Region         string                          `json:"region"`
	Version        int                             `json:"version"`
}

// CognitoDatasetRecord represents a record from an AWS Cognito event
type CognitoDatasetRecord struct {
	NewValue string `json:"newValue"`
	OldValue string `json:"oldValue"`
	Op       string `json:"op"`
}
