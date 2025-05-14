package models

import "time"

// Note we are using customer data as a record here as the scope of this project is for bi-directional record sync
// and not  record transformation or conflict resolution.
type InternalCustomer struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ExternalCustomer struct {
	CustomerID   string `json:"customer_id"`
	FullName     string `json:"full_name"`
	EmailAddress string `json:"email_address"`
	Phone        string `json:"phone"`
	LastModified string `json:"last_modified"`
}
