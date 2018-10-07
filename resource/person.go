package resource

import "database/sql"

type Person struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}

func (p *Person) Save(db *sql.DB) {
	if _, err := db.Exec("INSERT INTO persons (id, name, age) VALUES ($1, $2, $3)", p.Id, p.Name, p.Age); err != nil {
		panic(err)
	}
}

func FindPersonByID(db *sql.DB, id string) *Person {
	rows, err := db.Query("SELECT * FROM persons WHERE id=$1", id)
	if err != nil {
		return nil
	}
	defer rows.Close()

	p := Person{}
	for rows.Next() {
		if err := rows.Scan(&p.Id, &p.Name, &p.Age); err != nil {
			return nil
		}

		return &p
	}

	return nil
}
