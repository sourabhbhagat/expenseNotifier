package model

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

//Expense - data model
type Expense struct {
	ExpenseTime time.Time `bson:"expenseTime" json:"expenseTime"`
	IsLunch     bool      `bson:"isLunch" json:"isLunch"`
	Total       float32   `bson:"total" json:"total"`
	Currency    string    `bson:"currency" json:"currency"`
	Place       string    `bson:"place" json:"place"`
}

//SaveExpense - funtion to save expenses
func SaveExpense(expenseObject Expense) (Expense,error) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		fmt.Println("Mongo error", err.Error())
		return expenseObject, errors.New("Mongo connection Error " + err.Error())
	}

	defer session.Close()

	// Collection Expense
	c := session.DB("test").C("Expense")

	//total is calculated by the flag IsLunch
	if expenseObject.IsLunch == true {
		expenseObject.Total += 70
	} else {
		expenseObject.Total += 30
	}

	expenseObject.ExpenseTime = time.Now()
	expenseObject.Currency = "INR"
	expenseObject.Place  = "Swami Residency"

	err = c.Insert(expenseObject)
	if err != nil {
		fmt.Println("DB insert error", err.Error())
		return expenseObject, errors.New("Cannot insert data into DB " + err.Error())
	}
	return expenseObject, err
}

//GetExpense - handler to get expenses
func GetExpense(expenseID string) (Expense,error) {
	expenseObject := Expense{}
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		fmt.Println("Mongo error", err.Error())
		return expenseObject, errors.New("Mongo connection Error " + err.Error())
	}

	defer session.Close()

	// Collection Expense
	err = session.DB("test").C("Expense").Find(bson.M{"_id": expenseID}).One(&expenseObject)
	if err != nil {
		fmt.Println("Unable to find expense by ID", err.Error())
		return expenseObject, errors.New("Unable to find expense by ID " + err.Error())
	}

	return expenseObject, err
}
