package models

import (
	"fmt"
)

// A basic payment structure, with JSON spec.
type Payment struct {
	Id          int     `json:"id"`
	Beneficiary int     `json:"beneficiary"`
	Debtor      int     `json:"debtor"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
}

// Payment model type. A wrapper around a slice really.
type PaymentModel struct {
	store    map[int]*Payment
	sequence int
}

func NewPaymentModel() (model *PaymentModel) {
	model = &PaymentModel{
		store:    make(map[int]*Payment, 0),
		sequence: 0,
	}
	return
}

func NewPayment(beneficiary int, debtor int, amount float64) (payment *Payment) {
	payment = &Payment{
		Id:          -1,
		Beneficiary: beneficiary,
		Debtor:      debtor,
		Amount:      amount,
		Currency:    "GBP",
	}
	return
}

// Saves a payment to the payment model store
func (p *PaymentModel) CreatePayment(payment *Payment) (err error) {
	if payment.Id != -1 {
		err = fmt.Errorf("Cannot create existing payment")
	}

	var id int
	id, p.sequence = p.sequence, p.sequence+1

	p.store[id] = payment
	payment.Id = id

	return
}

// Updates a payment in the store
func (p *PaymentModel) UpdatePayment(payment *Payment) (err error) {
	if payment.Id == -1 {
		return fmt.Errorf("Inexistent payment")
	}

	if _, ok := p.store[payment.Id]; !ok {
		err = fmt.Errorf("Inexistent payment")
	}

	p.store[payment.Id] = payment

	return
}

// Deletes a payment from the store using its ID
func (p *PaymentModel) DeletePayment(payment *Payment) (err error) {
	if payment.Id == -1 {
		err = fmt.Errorf("uninitialized payment")
	}

	delete(p.store, payment.Id)

	return
}

// Retrieves a payment based on ID
func (p *PaymentModel) GetPaymentById(id int) (payment *Payment, err error) {
	var ok bool
	if payment, ok = p.store[id]; !ok {
		err = fmt.Errorf("payment not found")
	}

	return
}

// Converts the store of payments to a slice
func (p *PaymentModel) GetPayments() (payments []*Payment) {
	for _, el := range p.store {
		payments = append(payments, el)
	}

	if payments == nil {
		payments = make([]*Payment, 0)
	}

	return
}
