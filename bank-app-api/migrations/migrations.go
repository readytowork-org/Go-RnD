package migrations

import (
	"fmt"

	"sumit.com/go-bank-backend/helpers"
	"sumit.com/go-bank-backend/interfaces"
)

func createAccounts() {
	db := helpers.ConnectDB()

	users := &[2]interfaces.User{
		{Username: "Sumit", Email: "sumitsapkota0@gmail.com"},
		{Username: "Saroj", Email: "Sarojsapkota0@gmail.com"},
	}

	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)

		account := &interfaces.Account{Type: "Daily", Name: string(users[i].Username + "'s" + " account"), Balance: uint(10), UserId: uint(users[i].ID)}
		db.Create(&account)

	}
	defer db.Close()
}

//Migration of database
func Migrate() {
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	db := helpers.ConnectDB()
	db.AutoMigrate(&User{}, &Account{})
	defer db.Close()
	fmt.Println("Database migrated")
	createAccounts()
}
