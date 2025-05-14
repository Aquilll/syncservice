package provider

import (
	"fmt"
	"time"
	"syncservice/models"
)

type HubSpotProvider struct{}

func (h *HubSpotProvider) Create(customer models.ExternalCustomer) error {
	fmt.Printf("[HubSpot] Creating customer: %v\n", customer)
	time.Sleep(100 * time.Millisecond)
	return nil
}

func (h *HubSpotProvider) Update(customer models.ExternalCustomer) error {
	fmt.Printf("[HubSpot] Updating customer: %v\n", customer)
	time.Sleep(100 * time.Millisecond)
	return nil
}

func (h *HubSpotProvider) Delete(customerID string) error {
	fmt.Printf("[HubSpot] Deleting customer ID: %s\n", customerID)
	time.Sleep(100 * time.Millisecond)
	return nil
}
