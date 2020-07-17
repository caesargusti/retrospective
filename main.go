package main

import (
	"log"
	"github.com/caesargusti/todo2/storage"
	"github.com/caesargusti/todo2/model"
	"bufio"
	"os"
	"fmt"
)

func main(){
	
	//Initialisasi
	//var memStore = storage.NewMemory()

	var memStore = storage.GetStorage(storage.StorageTypeDatabase)

	fmt.Println("\n1. Create Data")
	fmt.Println("2. View Data")
	fmt.Println("3. Find Data")
	fmt.Println("4. Update Data")
	fmt.Println("5. Delete Data")
	fmt.Println("6. Exit")

	fmt.Print("\nChoose : ")
	var proses int
	_, err := fmt.Scanf("%d", &proses)
	
	switch proses {
	case 1 :
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\nID Customer : ")
		
		var id int
		_, err = fmt.Scanf("%d", &id)

		fmt.Print("Name        : ")
		name, _ := reader.ReadString('\n')
		fmt.Print("Address     : ")
		address, _ := reader.ReadString('\n')
		fmt.Print("Phone       : ")
		phone, _ := reader.ReadString('\n')
		
		dataCreate := model.Customers{
			ID:          id,
			Name:       name,
			Address: address,
			Phone: phone,
		}
		
		//create
		
		if err := memStore.Create(dataCreate); err != nil {
			log.Fatal(err)
		}
		log.Println("Added Success")

	case 2 :
		list, err := memStore.List()
		if err != nil {
			log.Fatal(err)
		}
		for _, val := range list {
			log.Printf("\nID      : %d\nName    : %sAddress : %sPhone   : %s\n", val.ID, val.Name, val.Address, val.Phone)
		}

	case 3 :

		fmt.Print("\nID Customer : ")
		
		var id int
		_, err = fmt.Scanf("%d", &id)

		var obj model.Customers
		obj, err = memStore.Detail(id)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("FOUND !!")
		log.Printf("\nID      : %d\nName    : %sAddress : %sPhone   : %s\n", obj.ID, obj.Name, obj.Address, obj.Phone)

	case 4 :
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\nID Customer : ")
		var id int
		_, err = fmt.Scanf("%d", &id)

		fmt.Println("New Data\n")
		fmt.Print("Name : ")
		name, _ := reader.ReadString('\n')
		fmt.Print("Address : ")
		address, _ := reader.ReadString('\n')
		fmt.Print("Phone : ")
		phone, _ := reader.ReadString('\n')

		err = memStore.Update(model.Customers{
			ID:        id,
			Name:      name,
			Address:   address,
			Phone:     phone,
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Println("\nData Updated")

	case 5 :
		fmt.Print("\nID Customer : ")
		var id int
		_, err = fmt.Scanf("%d", &id)

		err = memStore.Delete(id)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("\nData Deleted")
	case 6 :
		log.Println("Thank U")
	default:
		log.Fatal("Pilihan Tidak Ada")
	}
	
}
