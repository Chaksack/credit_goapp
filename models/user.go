package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id                    uint    `json:"id"`
	FirstName             string  `json:"first_name"`
	LastName              string  `json:"last_name"`
	Email                 string  `json:"email" gorm:"unique"`
	PhoneNumber           string  `json:"phone_number" gorm:"unique"`
	Age                   int     `json:"age"`
	MonthlyIncome         float32 `json:"monthly_income"`
	EmploymentStatus      string  `json:"employment_status"`
	EmploymentDuration    int     `json:"employment_duration"`
	NumberOfDependents    int     `json:"number_of_dep"`
	MaritalStatus         string  `json:"marital_status"`
	EducationalBackground string  `json:"educational_background"`
	HomeOwnershipStatus   string  `json:"home_status"`
	Password              []byte  `json:"-"`
	RoleId                uint    `json:"role_id"`
	Role                  Role    `json:"role" gorm:"foreignKey:RoleId"`
	GhcardId              uint    `json:"ghcard_id"`
	Ghcard                Ghcard  `json:"ghcard" gorm:"foreignKey:GhcardId"`
	UserTinId             uint    `json:"usertin_id"`
	UserTin               UserTin `json:"usertin" gorm:"foreignKey:UserTinId"`
	BankId                uint    `json:"bank_id"`
	Bank                  Bank    `json:"bank" gorm:"foreignKey:BankId"`
	MomoId                uint    `json:"momo_id"`
	Momo                  Momo    `json:"momo" gorm:"foreignKey:MomoId"`
	LoansId               uint    `json:"loans_id"`
	Loans                 Loans   `json:"loans" gorm:"foreignKey:LoansId"`
	CreditId              uint    `json:"credit_id"`
	Credit                Credit  `json:"credit" gorm:"foreignKey:CreditId"`
}

// func (user *User) SetPassword(password string) {
// 	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("1234"), 14)
// 	user.Password = hashedPassword
// }

// func (user *User) ComparePassword(password string) error {
// 	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
// }

func (user *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return nil
}

func (user *User) ComparePassword(password string) error {
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	return err
}
