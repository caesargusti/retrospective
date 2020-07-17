package main

import (
	"log"
	"github.com/caesargusti/todo2/storage"
	"github.com/caesargusti/todo2/model"
)

func main(){
	
	//Initialisasi
	//var memStore = storage.NewMemory()

	var memStore = storage.GetStorage(storage.StorageTypeDatabase)
	
		dataCreate := model.Customers{
			ID:          1,
			Name:       "Caesar Pamungkas",
			Address: "Lampung",
			Phone: "081394442007",
		}
		
		//create
		
		if err := memStore.Create(dataCreate); err != nil {
			log.Fatal(err)
		}
		log.Println("Added Success")

		list, err := memStore.List()
		if err != nil {
			log.Fatal(err)
		}

		iteratorData := storage.CustomersCollection{Cus:list}.CreateDataIterator()
		for iteratorData.HasNext(){
			val := iteratorData.GetNext()
			log.Printf("\nID      : %d\nName    : %s\nAddress : %s\nPhone   : %s\n", val.ID, val.Name, val.Address, val.Phone)
		}
		// for _, val := range list {
		// 	log.Printf("\nID      : %d\nName    : %sAddress : %sPhone   : %s\n", val.ID, val.Name, val.Address, val.Phone)
		// }

		var obj model.Customers
		obj, err = memStore.Detail(1)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("FOUND !!")
		log.Printf("\nID      : %d\nName    : %s\nAddress : %s\nPhone   : %s\n", obj.ID, obj.Name, obj.Address, obj.Phone)

		// err = memStore.Update(model.Customers{
		// 	ID:        id,
		// 	Name:      name,
		// 	Address:   address,
		// 	Phone:     phone,
		// })
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// log.Println("\nData Updated")

		// err = memStore.Delete(id)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// log.Println("\nData Deleted")
	
}
