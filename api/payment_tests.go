package api

import (
	"techtest-payments-api/api"
	"testing"
)

func TestCreatePayment(t *testing.T) {
	p := NewPaymentModel()

	err := p.CreatePayment(&Payment{
		Beneficiary: 123,
		Debtor:      122,
		Amount:      100.00,
		Currency:    "GBP",
	})

	if err != nil {
		t.Error("An error occured when creating a payment: %s", err)
	}
}

func TestCreatePaymentAlreadyExists(t *testing.T) {
	p := NewPaymentModel()

	payment1 := &Payment{
		Beneficiary: 123,
		Debtor:      122,
		Amount:      100.00,
		Currency:    "GBP",
	}

	p.CreatePayment(payment1)
	err := p.CreatePayment(payment1)

	if err == nil {
		t.Error("Saving a duplicate payment should result in an error")
	} else if err != "Payment ID already exists" {
		t.Error("Error when saving duplicate payment was not as expected")
	}
}

func TestUpdatePayment(t *testing.T) {
	p := NewPaymentModel()

	payment := &Payment{
		Beneficiary: 123,
		Debtor:      122,
		Amount:      123.23,
		Currency:    "GBP",
	}

	p.CreatePayment(payment)

	payment.Amount = 200.10
	err = p.UpdatePayment(payment)

	if err != nil {
		t.Error("Failed updating payment")
	}
}

func TestUpdateInexistentPayment(t *testing.T) {
	p := NewPaymentModel()

	payment := &Payment{
		Beneficiary: 123,
		Debtor:      122,
		Amount:      123.23,
		Currency:    "GBP",
	}

	p.CreatePayment(payment)

	payment.Amount = 200.10
	err = p.UpdatePayment(payment)

	if err == nil {
		t.Error("Updating a non-existent payment should fail")
	} else if err != "Inexistent payment" {
		t.Error("Updating a non-existent payment should fail with the correct error")
	}
}

func TestRetrievePayment(t *testing.T) {
	p := NewPaymentModel()

	payment := &Payment{
		Beneficiary: 123,
		Debtor:      122,
		Amount:      123.23,
		Currency:    "GBP",
	}

	p.CreatePayment(payment)

	ok, payment2 := p.GetPaymentById(payment.Id)

	if !ok {
		t.Error("Failed to retrieve an existing payment")
	}

	if payment2.Id != payment1.Id {
		t.Error("Retrieved payment was not the one requested")
	}
}

func TestRetrieveInexistentPayment(t *testing.T) {
	p := NewPaymentModel()

	ok, _ := p.GetPaymentById(1337)

	if ok {
		t.Error("Retrieving non-existent payment should fail")
	}
}

func TestDeletePayment(t *testing.T) {
	p := NewPaymentModel()

	payment := &Payment{
		Beneficiary: 123,
		Debtor:      122,
		Amount:      123.23,
		Currency:    "GBP",
	}

	p.CreatePayment(payment)

	err := p.DeletePayment(payment)

	if err {
		t.Error("Failed to delete an existing payment")
	}
}

func TestDeleteInexistentPayment(t *testing.T) {
	p := NewPaymentModel()

	err := p.DeletePayment(1337)

	if !err {
		t.Error("Deleting an inexistent payment should fail")
	}
}
