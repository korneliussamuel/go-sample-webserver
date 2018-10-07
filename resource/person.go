package resource

import (
	"database/sql"
)

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
