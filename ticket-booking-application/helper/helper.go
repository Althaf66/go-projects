package helper

import(
  "strings"
)

func ValidateInput(userf string,userl string,email string,ticket int,remaining int)(bool,bool,bool){
  isValidName := len(userf) >= 2 && len(userl) >= 2
  isValidEmail := strings.Contains(email,"@")
  isValidTickets := ticket > 0 && ticket <= remaining
  
  return isValidName,isValidEmail,isValidTickets
}