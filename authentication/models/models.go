package models

type User struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	WalletAddress string `json:"walletAddress"`
	Email         string `json:"email"`
	UserPhoto     string `json:"userPhoto"`
	Age           int    `json:"age"`
	Profession    string `json:"profession"`
	Country       string `json:"country"`
	City          string `json:"city"`
	PropertyOwned int    `json:"propertyOwned"`
}
