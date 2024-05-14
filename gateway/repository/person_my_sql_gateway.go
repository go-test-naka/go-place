package repository

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-test-naka/go-place/log"
	"github.com/go-test-naka/go-place/model"
)

type PersonMySQLRepository struct {
	db *sql.DB
}

func CreatePersonMySQLRepository(db *sql.DB) *PersonMySQLRepository {
	return &PersonMySQLRepository{db}
}

func (pmr *PersonMySQLRepository) Save(person model.Person) error {
	INSERT := "INSERT INTO sys.Person(id, name, birthday, country) VALUES(0, ?, ?, ?);"
	stmt, err := pmr.db.Prepare(INSERT)
	stmt.Exec(person.Name, person.Birthday, person.Country)
	if err != nil {
		log.Error("Failed to insert data", err)
		panic(err.Error())
	}
	defer stmt.Close()
	return nil
}

func (pmr *PersonMySQLRepository) FindAll() []*model.Person {
	SELECT := "SELECT * FROM Person limit 100"
	res, err := pmr.db.Query(SELECT)
	if err != nil {
		log.Error("Failed to query data", err)
	}

	return rowsToModel(res)
}

func (pmr *PersonMySQLRepository) FindByName(name string) *model.Person {
	SELECT := "SELECT * FROM Person p WHERE p.name = ?"

	stmt, err := pmr.db.Prepare(SELECT)
	if err != nil {
		log.Error("Failed to query data", err)
	}

	res, err := stmt.Query(name)
	if err != nil {
		log.Error("Failed to query data", err)
	}

	person := rowsToModel(res)
	if len(person) > 0 {
		return person[0]
	}
	return nil
}

func rowsToModel(rows *sql.Rows) []*model.Person {
	persons := []*model.Person{}
	for rows.Next() {
		person, err := entityToModel(rows)
		if err != nil {
			log.Error("Failed to convert entity to model", err)
			continue
		}
		persons = append(persons, person)
	}
	return persons
}

func entityToModel(rows *sql.Rows) (*model.Person, error) {
	var id uint
	var name string
	var birthdayDB []uint8
	var country string

	if err := rows.Scan(&id, &name, &birthdayDB, &country); err != nil {
		return nil, err
	}

	birthdayStr := string(birthdayDB)
	birthday, err := time.Parse("2006-01-02 15:04:05", birthdayStr)
	if err != nil {
		return nil, err
	}

	p := model.Person{
		Id:       id,
		Name:     name,
		Birthday: birthday,
		Country:  country}

	return &p, nil
}
