package database

import (
	"context"
	"github.com/janobono/linked-in-build-a-microservice-with-go/internal/models"
)

func (c client) GetAllServices(ctx context.Context) ([]models.Service, error) {
	var services []models.Service
	result := c.DB.WithContext(ctx).
		Find(&services)
	return services, result.Error
}
