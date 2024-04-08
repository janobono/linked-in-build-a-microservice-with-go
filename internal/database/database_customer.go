package database

import (
	"context"

	"github.com/janobono/linked-in-build-a-microservice-with-go/internal/models"
)

func (c client) GetAllCustomers(ctx context.Context, emailAddress string) ([]models.Customer, error) {
	var customers []models.Customer
	result := c.DB.WithContext(ctx).
		Where(models.Customer{Email: emailAddress}).
		Find(&customers)
	return customers, result.Error
}
