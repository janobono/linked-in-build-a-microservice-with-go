package database

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/janobono/linked-in-build-a-microservice-with-go/internal/dberrors"
	"github.com/janobono/linked-in-build-a-microservice-with-go/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (c client) GetAllVendors(ctx context.Context) ([]models.Vendor, error) {
	var vendors []models.Vendor
	result := c.DB.WithContext(ctx).
		Find(&vendors)
	return vendors, result.Error
}

func (c client) AddVendor(ctx context.Context, vendor *models.Vendor) (*models.Vendor, error) {
	vendor.VendorID = uuid.NewString()
	result := c.DB.WithContext(ctx).Create(&vendor)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}

	return vendor, nil
}

func (c client) GetVendorById(ctx context.Context, ID string) (*models.Vendor, error) {
	vendor := &models.Vendor{}
	result := c.DB.WithContext(ctx).
		Where(&models.Vendor{VendorID: ID}).
		First(&vendor)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "vendor", ID: ID}
		}
		return nil, result.Error
	}
	return vendor, nil
}

func (c client) UpdateVendor(ctx context.Context, vendor *models.Vendor) (*models.Vendor, error) {
	var vendors []models.Vendor
	result := c.DB.WithContext(ctx).
		Model(&vendors).
		Clauses(clause.Returning{}).
		Where(&models.Vendor{VendorID: vendor.VendorID}).
		Updates(models.Vendor{
			Name:    vendor.Name,
			Contact: vendor.Contact,
			Phone:   vendor.Phone,
			Email:   vendor.Email,
			Address: vendor.Address,
		})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, &dberrors.NotFoundError{Entity: "vendor", ID: vendor.VendorID}
	}
	return &vendors[0], nil
}

func (c client) DeleteVendor(ctx context.Context, ID string) error {
	return c.DB.WithContext(ctx).Delete(&models.Vendor{VendorID: ID}).Error
}
