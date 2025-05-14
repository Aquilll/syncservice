package models

import "testing"

func TestInternalCustomerJSONTags(t *testing.T) {
    c := InternalCustomer{}
    if c.ID != "" {
        t.Errorf("expected empty ID, got %s", c.ID)
    }
}

func TestExternalCustomerJSONTags(t *testing.T) {
    c := ExternalCustomer{}
    if c.CustomerID != "" {
        t.Errorf("expected empty CustomerID, got %s", c.CustomerID)
    }
}
