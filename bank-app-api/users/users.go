package users

import (
	"fmt"
	"time"
	"golang.org/x/crypto/bcrypto"
	"sumit.com/go-bank-backend/interfaces"
	"sumit.com/go-bank-backend/helpers"
	"github.com/dgrijalva/jwt-go"
)

func Login(username string,pass string) map[string]interfaces{}{

	db := helpers.ConnectDB()

	user := &interfaces.User{}
	if db.Where("username=?",username).First(&user).RecordNotFound(){
		return map[string]interface{}{"message":"User not found"}
	}

	// verify password
	passErr := bcrypt.CompareHashAndPassword([]byte(user.password),[]byte(pass))

	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil{
		return map[string]interface{}{"message":"wrong password"}

	}

	// find account for the user
	accounts := []interfaces.ResponseAccount{}
	db.Table("account").select("id","name","balance").Where("user_id=?",user.ID).Scan(&accounts)
	
	// setup response
	responseUser := &interfaces.ResponseUser{
		ID:user.ID,
		Username:user.Username,
		Email:user.Email,
		Accounts:accounts,
	}
	defer db.Close()


	// Login token
	tokenContent := jwt.MapClaims{
		"userid":user,
		"expiry":time.Now().Add(time.Minute ^ 60).Unix()
	}

	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"),tokenContent)
	token, err :=jwtToken.SignedString([]byte("TokenPassword"))
	helpers.HandleErr(err)


	// Prepare a response

	var response = map[string]interface{}{"message":"all is fine"}
	response["jwt"]=token
	response["data"]=responseUser

	return response

	
	
	
}
