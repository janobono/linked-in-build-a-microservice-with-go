package database

import (
	"context"
	"github.com/janobono/linked-in-build-a-microservice-with-go/internal/models"
)

func (c client) GetAllProducts(ctx context.Context, vendorID string) ([]models.Product, error) {
	var products []models.Product
	result := c.DB.WithContext(ctx).
		Where(models.Product{VendorID: vendorID}).
		Find(&products)
	return products, result.Error
}
