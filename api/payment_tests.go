package api

import (
	"testing"
)

func TestCreatePayment(t *testing.T) {
	p := NewPaymentModel()

	err := p.CreatePayment(NewPayment(123, 122, 100.00))

	if err != nil {
		t.Errorf("An error occured when creating a payment: %s", err)
	}
}

func TestCreatePaymentAlreadyExists(t *testing.T) {
	p := NewPaymentModel()

	payment1 := NewPayment(123, 122, 100.00)

	p.CreatePayment(payment1)
	err := p.CreatePayment(payment1)

	if err == nil {
		t.Error("Saving a duplicate payment should result in an error")
	}
}

func TestUpdatePayment(t *testing.T) {
	p := NewPaymentModel()

	payment := NewPayment(123, 122, 123.23)

	p.CreatePayment(payment)

	payment.Amount = 200.10
	err := p.UpdatePayment(payment)

	if err != nil {
		t.Errorf("Failed updating payment; error: %s", err)
	}
}

func TestUpdateInexistentPayment(t *testing.T) {
	p := NewPaymentModel()

	payment := NewPayment(123, 122, 123.23)

	payment.Amount = 200.10
	err := p.UpdatePayment(payment)

	if err == nil {
		t.Error("Updating a non-existent payment should fail")
	}
}

func TestRetrievePayment(t *testing.T) {
	p := NewPaymentModel()

	payment1 := NewPayment(123, 122, 123.23)

	p.CreatePayment(payment1)

	payment2, err := p.GetPaymentById(payment1.Id)

	if err != nil {
		t.Errorf("Failed to retrieve an existing payment, error: %s", err)
	}

	if payment2.Id != payment1.Id {
		t.Error("Retrieved payment was not the one requested")
	}
}

func TestRetrieveInexistentPayment(t *testing.T) {
	p := NewPaymentModel()

	_, err := p.GetPaymentById(1337)

	if err != nil {
		t.Errorf("Retrieving non-existent payment should fail, error: %s", err)
	}
}

func TestRetrievePaymentCollection(t *testing.T) {
	p := NewPaymentModel()

	p.CreatePayment(NewPayment(1, 2, 100.00))

	p.CreatePayment(NewPayment(2, 1, 50.00))

	payments := p.GetPayments()

	if len(payments) != 2 {
		t.Error("Payments collection should contain 2 items")
	}
}

func TestDeletePayment(t *testing.T) {
	p := NewPaymentModel()

	payment := NewPayment(123, 122, 123.23)

	p.CreatePayment(payment)

	err := p.DeletePayment(payment)

	if err != nil {
		t.Errorf("Failed to delete an existing payment, error: %s", err)
	}
}

func TestDeleteInexistentPayment(t *testing.T) {
	p := NewPaymentModel()

	payment := NewPayment(1, 2, 3)
	payment.Id = 1337

	err := p.DeletePayment(payment)

	if err != nil {
		t.Errorf("Deleting an inexistent payment should fail, error: %s", err)
	}
}
