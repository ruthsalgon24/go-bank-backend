package useraccounts

import (
	"fmt"

	"duomly.com/go-bank-backend/database"
	"duomly.com/go-bank-backend/helpers"
	"duomly.com/go-bank-backend/interfaces"
	"duomly.com/go-bank-backend/transactions"
)

// Refactor function updateAccount to use database package
func updateAccount(id uint, amount float64) interfaces.ResponseAccount {
	account := interfaces.Account{}
	responseAcc := interfaces.ResponseAccount{}

	database.DB.Where("id = ? ", id).First(&account)
	account.Balance = float64(amount)
	database.DB.Save(&account)

	responseAcc.ID = account.ID
	responseAcc.Name = account.Name
	responseAcc.Balance = float64(account.Balance)
	return responseAcc
}

// Refactor function getAccount to use database package
func getAccount(id uint) *interfaces.Account {
	account := &interfaces.Account{}
	if database.DB.Where("id = ? ", id).First(&account).RecordNotFound() {
		return nil
	}
	return account
}

// Create function Transaction
func Transaction(userId uint, from uint, to uint, amount float64, jwt string) map[string]interface{} {
	// Conver uint to string
	userIdString := fmt.Sprint(userId)
	// Validate ownership
	isValid := helpers.ValidateToken(userIdString, jwt)
	if isValid {
		// Take sender and receiver
		fromAccount := getAccount(from)
		toAccount := getAccount(to)
		// Handle errors
		if fromAccount == nil || toAccount == nil {
			return map[string]interface{}{"message": "Account not found"}
		} else if fromAccount.UserID != userId {
			return map[string]interface{}{"message": "You are not owner of the account"}
		} else if float64(fromAccount.Balance) < amount {
			return map[string]interface{}{"message": "Account balance is too small"}
		}
		// Update account
		updatedAccount := updateAccount(from, float64(fromAccount.Balance)-amount)
		updateAccount(to, float64(fromAccount.Balance)+amount)

		// Create transaction
		transactions.CreateTransaction(from, to, amount)

		// Return response
		var response = map[string]interface{}{"message": "all is fine"}
		response["data"] = updatedAccount
		return response
	} else {
		return map[string]interface{}{"message": "Not valid token"}
	}
}
