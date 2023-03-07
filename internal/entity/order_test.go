package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_If_It_Gets_An_Error_If_ID_Is_Blank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "ID is required")
}
func Test_If_It_Gets_An_Error_If_Price_Is_Blank(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.Validate(), "Price is empty")
}
func Test_If_It_Gets_An_Error_If_Tax_Is_Blank(t *testing.T) {
	order := Order{ID: "123", Price: 10}
	assert.Error(t, order.Validate(), "Tax is empty")
}
func Test_WithAllValidParams(t *testing.T) {
	order := Order{ID: "123", Price: 10, Tax: 1.0}
	assert.NoError(t, order.Validate())
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 11.0, order.FinalPrice)
}
