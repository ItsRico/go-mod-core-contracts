//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"

	"github.com/google/uuid"
)

// Notification and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/support-notifications/2.x#/Notification
type Notification struct {
	Id          string   `json:"id,omitempty" validate:"omitempty,uuid"`
	Created     int64    `json:"created,omitempty"`
	Modified    int64    `json:"modified,omitempty"`
	Category    string   `json:"category,omitempty" validate:"required_without=Labels,omitempty,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Labels      []string `json:"labels,omitempty" validate:"required_without=Category,omitempty,gt=0,dive,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Content     string   `json:"content" validate:"required,edgex-dto-none-empty-string"`
	ContentType string   `json:"contentType,omitempty"`
	Description string   `json:"description,omitempty"`
	Sender      string   `json:"sender" validate:"required,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Severity    string   `json:"severity" validate:"required,oneof='MINOR' 'NORMAL' 'CRITICAL'"`
	Status      string   `json:"status,omitempty" validate:"omitempty,oneof='NEW' 'PROCESSED' 'ESCALATED'"`
}

// NewNotification creates and returns a Notification DTO
func NewNotification(labels []string, category, content, sender, severity string) Notification {
	return Notification{
		Id:       uuid.NewString(),
		Labels:   labels,
		Category: category,
		Content:  content,
		Sender:   sender,
		Severity: severity,
	}
}

// ToNotificationModel transforms the Notification DTO to the Notification Model
func ToNotificationModel(n Notification) models.Notification {
	var m models.Notification
	m.Id = n.Id
	m.Created = n.Created
	m.Modified = n.Modified
	m.Category = n.Category
	m.Labels = n.Labels
	m.Content = n.Content
	m.ContentType = n.ContentType
	m.Description = n.Description
	m.Sender = n.Sender
	m.Severity = models.NotificationSeverity(n.Severity)
	m.Status = models.NotificationStatus(n.Status)
	return m
}

// ToNotificationModels transforms the Notification DTO array to the Notification model array
func ToNotificationModels(notifications []Notification) []models.Notification {
	models := make([]models.Notification, len(notifications))
	for i, n := range notifications {
		models[i] = ToNotificationModel(n)
	}
	return models
}

// FromNotificationModelToDTO transforms the Notification Model to the Notification DTO
func FromNotificationModelToDTO(n models.Notification) Notification {
	return Notification{
		Id:          n.Id,
		Created:     n.Created,
		Modified:    n.Modified,
		Category:    string(n.Category),
		Labels:      n.Labels,
		Content:     n.Content,
		ContentType: n.ContentType,
		Description: n.Description,
		Sender:      n.Sender,
		Severity:    string(n.Severity),
		Status:      string(n.Status),
	}
}

// FromNotificationModelsToDTOs transforms the Notification model array to the Notification DTO array
func FromNotificationModelsToDTOs(notifications []models.Notification) []Notification {
	dtos := make([]Notification, len(notifications))
	for i, n := range notifications {
		dtos[i] = FromNotificationModelToDTO(n)
	}
	return dtos
}
