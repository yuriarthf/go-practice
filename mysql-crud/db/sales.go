package db

import (
	"errors"
	"fmt"
	"strings"
)

type Sale struct {
	ID         uint64 `json:"id"`
	CustomerID uint64 `json:"customerId"`
	ItemID     uint64 `json:"itemId"`
	Quantity   uint64 `json:"quantity"`
}

type SaleItem struct {
	ItemID   uint64 `json:"itemId"`
	Quantity uint64 `json:"quantity"`
}

func (c *Sale) Fields() []interface{} {
	return []interface{}{
		&c.ID,
		&c.CustomerID,
		&c.ItemID,
		&c.Quantity,
	}
}

func RegisterSale(custid uint64, sale []SaleItem) error {
	if len(sale) == 0 {
		return errors.New("No item sold")
	}
	_, err := GetCustomerById(custid)
	if err != nil {
		return errors.New("Customer not found")
	}

	var ids []uint64
	var args []interface{}
	for _, s := range sale {
		ids = append(ids, s.ItemID)
		args = append(args, s.ItemID, s.Quantity)
	}

	f, _ := itemsExist(ids)
	if !f {
		return errors.New("Unregistered item present")
	}

	query := "INSERT INTO (CUSTOMER_ID, ITEM_ID, QUANTITY) VALUES " +
		strings.TrimSuffix(strings.Repeat(fmt.Sprintf("(%d, ?, ?),", custid), len(ids)), ",")

	_, err = db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}
