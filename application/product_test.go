package application_test

import (
	"testing"

	"github.com/leoCardosoDev/arquitetura_hexagonal_full_cycle/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enabled(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disabled(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal zero", err.Error())
}

func TestProduct_GetID(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	got := product.GetID()
	want := product.ID
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestProduct_GetName(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello World"
	got := product.GetName()
	want := "Hello World"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestProduct_GetStatus(t *testing.T) {
	product := application.Product{}
	product.Status = application.ENABLED
	got := product.GetStatus()
	want := product.Status
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestProduct_GetPrice(t *testing.T) {
	product := application.Product{}
	product.Price = 10.4
	got := product.GetPrice()
	want := 10.4
	if got != want {
		t.Errorf("got %.18f, wanted %.18f", got, want)
	}
}
