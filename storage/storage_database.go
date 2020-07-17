package storage
import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"github.com/caesargusti/todo2/model" 
	"errors"
)

type database struct{}

func connection() *sql.DB{
	db, err := sql.Open("postgres", "host=postgres port=5432 user=postgres password=password dbname=db sslmode=disable")
	// db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=password dbname=db sslmode=disable")
	
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func (o database) Create(obj model.Customers) error{
	db := connection()
	// _, err := db.Exec("CREATE TABLE customers (id int, name varchar(20), address varchar(100), phone varchar(20))")
	_, err := db.Exec("INSERT into customers(id,name,address,phone) values ($1, $2, $3, $4);", 
		obj.ID,
		obj.Name,
		obj.Address,
		obj.Phone,
	)
	return err
}

func (o database) List() ([]model.Customers, error){
	db := connection()
	rows, err := db.Query("select id, name, address, phone from customers Limit 5")
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	res := make([]model.Customers, 0) // Inisialisasi Slice SIZE 0
	for rows.Next(){
		var cus model.Customers
		err := rows.Scan(&cus.ID,
			&cus.Name,
			&cus.Address,
			&cus.Phone)
		if err != nil{
			return nil, err
		}
		res = append(res, cus)
	}
	return res, nil
}

func (o database) Detail(id int) (model.Customers, error){
	db := connection()
	rows, err := db.Query("SELECT id,name,address,phone FROM customers WHERE id = $1;", id)
	if err != nil{
		return model.Customers{}, err
	}
	defer rows.Close()
	for rows.Next(){
		var cus model.Customers
		err := rows.Scan(&cus.ID,
			&cus.Name,
			&cus.Address,
			&cus.Phone)
		if err != nil{
			return model.Customers{}, err
		}
		return cus, nil
	}
	return model.Customers{}, errors.New("Customers tidak ditemukan")
}

func (o database) Update(obj model.Customers) error {
	db := connection()
	_, err := db.Exec("UPDATE customers SET name = $1, address = $2, phone = $3 WHERE id = $4",
		obj.Name,
		obj.Address,
		obj.Phone,
		obj.ID,
	)
	return err 
	// panic("aw")
}

func (o database) Delete(id int) error{
	db := connection()
	_, err := db.Exec("DELETE from customers WHERE id = $1",id)
	return err
}