package provider

import (
	"fmt"
	"time"
	"syncservice/models"
)

type SalesforceProvider struct{}

func (s *SalesforceProvider) Create(customer models.ExternalCustomer) error {
	fmt.Printf("[Salesforce] Creating customer: %v\n", customer)
	time.Sleep(100 * time.Millisecond)
	return nil
}

func (s *SalesforceProvider) Update(customer models.ExternalCustomer) error {
	fmt.Printf("[Salesforce] Updating customer: %v\n", customer)
	time.Sleep(100 * time.Millisecond)
	return nil
}

func (s *SalesforceProvider) Delete(customerID string) error {
	fmt.Printf("[Salesforce] Deleting customer ID: %s\n", customerID)
	time.Sleep(100 * time.Millisecond)
	return nil
}
