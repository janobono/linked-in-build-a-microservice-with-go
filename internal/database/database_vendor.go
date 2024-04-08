package database

import (
	"context"
	"github.com/janobono/linked-in-build-a-microservice-with-go/internal/models"
)

func (c client) GetAllVendors(ctx context.Context) ([]models.Vendor, error) {
	var vendors []models.Vendor
	result := c.DB.WithContext(ctx).
		Find(&vendors)
	return vendors, result.Error
}
