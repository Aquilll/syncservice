package transformer

import (
    "testing"
    "time"
    "syncservice/models"
)

func TestTransformationRoundTrip(t *testing.T) {
    now := time.Now()
    internal := models.InternalCustomer{
        ID:          "123",
        FirstName:   "John",
        LastName:    "Doe",
        Email:       "john@example.com",
        PhoneNumber: "1234567890",
        UpdatedAt:   now,
    }

    external := ToExternal(internal)
    got := ToInternal(external)

    if got.ID != internal.ID {
        t.Errorf("expected ID %s, got %s", internal.ID, got.ID)
    }
}
