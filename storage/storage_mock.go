package storage

import(
	"errors"
	"github.com/caesargusti/todo2/model"
)

type mock struct{}

func (o mock) List() ([]model.Customers, error) {
	result := []model.Customers{}
	result = append(result, model.Customers{
		ID:          1,
		Name:       "Mock Title\n",
		Address: "Mock Description\n",
		Phone:   "008880",
	}, model.Customers{
		ID:          2,
		Name:       "Mock Title\n",
		Address: "Second mock\n",
		Phone:   "009990",
	})
	return result, nil
}
func (o mock) Create(obj model.Customers) error {
	return nil
}
func (o mock) Detail(id int) (model.Customers, error) {
	if id == 1 {
		return model.Customers{ID: 1, Name: "Mock Title\n", Address: "Arab\n", Phone: "008880"}, nil
	}
	return model.Customers{}, errors.New("Todo tidak ditemukan")
}
func (o mock) Delete(id int) error {
	return nil
}

func (o mock) Update(obj model.Customers) error {
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
