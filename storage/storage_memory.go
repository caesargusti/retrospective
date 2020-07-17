package storage

import(
	"errors"
	"github.com/caesargusti/todo2/model"
)

//Objek
type memory struct {
	store map[int]model.Customers
}

//ini dihapus tadinya pake pointer karena tidak pakai pointer yang lain jadi tidak konsisten
func newMemory() memory {
	return memory{
		store: make(map[int]model.Customers),
	}
}

func (o memory) List() ([]model.Customers, error) {
	//looping
	result := []model.Customers{}
	for _, v := range o.store {
		result = append(result, v)
	}
	return result, nil
}
func (o memory) Create(obj model.Customers) error {
	//menyimpan
	o.store[obj.ID] = obj
	return nil
}
func (o memory) Detail(id int) (model.Customers, error) {
	//untuk ngecek apakah id tertentu exist dalam sebuah array atau map untuk menyambungkan if pakai semicolon
	//ini gaperlu dicasting karena karena static
	if obj, ok := o.store[id]; ok {
		return obj, nil
	}
	return model.Customers{}, errors.New("Customers tidak ditemukan")
}
func (o memory) Delete(id int) error {
	delete(o.store, id)
	return nil
}

func (o memory) Update(obj model.Customers) error {
	// db := connection()
	// _, err := db.Exec("UPDATE customers SET name = $1, address = $2, phone = $3, birthdate = $4 WHERE id = $5",
	// 	obj.Name,
	// 	obj.Address,
	// 	obj.Phone,
	// 	obj.BirthDate,
	// 	obj.ID,
	// )
	// return err 
	panic("NOT IMPLEMENTED")
}

