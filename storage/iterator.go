package storage

import "github.com/caesargusti/todo2/model"

type Iterator interface {
    HasNext() bool
    GetNext() model.Customers
}

type CustomerIterator struct {
    index int
    customerList []model.Customers
}

func (o *CustomerIterator) HasNext() bool {
    if o.index < len(o.customerList) {
        return true
    }
    return false
}

func (o *CustomerIterator) GetNext() model.Customers {
    if o.HasNext() {
        customer := o.customerList[o.index]
        o.index++
        return customer
    }
    return model.Customers{}
}


type Collection interface {
    CreateDataIterator() Iterator
}

type CustomersCollection struct {
    Cus []model.Customers
}

func (o CustomersCollection) CreateDataIterator() Iterator {
    return &CustomerIterator{
		index: 0,
        customerList: o.Cus,
    }
}


