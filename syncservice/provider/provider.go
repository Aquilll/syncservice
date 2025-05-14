package provider

import "syncservice/models"

type CRMProvider interface {
	Create(customer models.ExternalCustomer) error
	Update(customer models.ExternalCustomer) error
	Delete(customerID string) error
}
