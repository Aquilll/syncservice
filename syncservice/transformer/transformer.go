package transformer

import (
	"time"
	"syncservice/models"
)

func ToExternal(internal models.InternalCustomer) models.ExternalCustomer {
	return models.ExternalCustomer{
		CustomerID:   internal.ID,
		FullName:     internal.FirstName + " " + internal.LastName,
		EmailAddress: internal.Email,
		Phone:        internal.PhoneNumber,
		LastModified: internal.UpdatedAt.Format(time.RFC3339),
	}
}

func ToInternal(external models.ExternalCustomer) models.InternalCustomer {
	return models.InternalCustomer{
		ID:          external.CustomerID,
		FirstName:   external.FullName,
		LastName:    "",
		Email:       external.EmailAddress,
		PhoneNumber: external.Phone,
		UpdatedAt:   time.Now(),
	}
}
