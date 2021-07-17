package services

import "testing"

var items []int

func init() {
	if items == nil {
		items = []int{250, 500, 1000, 2000, 5000}
	}
}

func TestOrder_1(t *testing.T) {
	order := Calculate(items, 1)

	if len(order) != 1 {
		t.Errorf("There should only be 1 item")
	}

	if order[0] != 250 {
		t.Errorf("Item order quantity should be 250")
	}
}

func TestOrder_250(t *testing.T) {
	order := Calculate(items, 250)

	if len(order) != 1 {
		t.Errorf("There should only be 1 item")
	}

	if order[0] != 250 {
		t.Errorf("Item order quantity should be 250")
	}
}

func TestOrder_251(t *testing.T) {
	order := Calculate(items, 251)

	if len(order) != 1 {
		t.Errorf("There should only be 1 item")
	}

	if order[0] != 250 {
		t.Errorf("Item order quantity should be 500")
	}
}

func TestOrder_501(t *testing.T) {
	order := Calculate(items, 501)

	if len(order) != 2 {
		t.Errorf("There should be 2 items")
	}

	if order[0] != 500 {
		t.Errorf("First Item order quantity should be 500")
	}

	if order[1] != 250 {
		t.Errorf("Second Item order quantity should be 250")
	}
}

func TestOrder_12001(t *testing.T) {
	order := Calculate(items, 12001)

	if len(order) != 4 {
		t.Errorf("There should be 2 items")
	}

	if order[0] != 5000 {
		t.Errorf("First Item order quantity should be 500")
	}

	if order[1] != 5000 {
		t.Errorf("Second Item order quantity should be 250")
	}

	if order[2] != 2000 {
		t.Errorf("Second Item order quantity should be 250")
	}

	if order[3] != 250 {
		t.Errorf("Second Item order quantity should be 250")
	}
}
