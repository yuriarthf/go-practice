package db

import "database/sql"

type Customer struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	Age         uint8          `json:"age"`
	Address     sql.NullString `json:"address"`
	LastVisited string         `json:"lastVisited"`
	LastUpdated string         `json:"lastUpdated"`
	CreatedAt   string         `json:"createdAt"`
}

func (c *Customer) Fields() []interface{} {
	return []interface{}{
		&c.ID,
		&c.Name,
		&c.Age,
		&c.Address,
		&c.LastVisited,
		&c.LastUpdated,
		&c.CreatedAt,
	}
}

func GetAllCustomers() ([]Customer, error) {
	r, err := db.Query("SELECT * FROM Customers")
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var result []Customer
	for r.Next() {
		var cust Customer
		if err := r.Scan(cust.Fields()...); err != nil {
			return nil, err
		}
		result = append(result, cust)
	}

	return result, nil
}

func GetCustomerById(id uint64) (Customer, error) {
	r := db.QueryRow("SELECT * FROM Customers WHERE ID = ?", id)

	var result Customer
	if err := r.Scan(result.Fields()...); err != nil {
		return result, err
	}

	return result, nil
}

func NewCustomer(name string, age uint8, address string) (int64, error) {
	var r sql.Result
	var err error
	if len(address) == 0 {
		r, err = db.Exec(
			`INSERT INTO Customers (NAME, AGE) VALUES (?, ?)`,
			name, age,
		)
	} else {
		r, err = db.Exec(
			`INSERT INTO Customers (NAME, AGE, ADDRESS) VALUES (?, ?, ?)`,
			name, age, address,
		)
	}

	if err != nil {
		return -1, err
	}

	return r.LastInsertId()
}
