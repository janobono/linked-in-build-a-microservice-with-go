package database

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/janobono/linked-in-build-a-microservice-with-go/internal/dberrors"
	"gorm.io/gorm"

	"github.com/janobono/linked-in-build-a-microservice-with-go/internal/models"
)

func (c client) GetAllCustomers(ctx context.Context, emailAddress string) ([]models.Customer, error) {
	var customers []models.Customer
	result := c.DB.WithContext(ctx).
		Where(models.Customer{Email: emailAddress}).
		Find(&customers)
	return customers, result.Error
}

func (c client) AddCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	customer.CustomerID = uuid.NewString()
	result := c.DB.WithContext(ctx).
		Create(&customer)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	return customer, nil
}
