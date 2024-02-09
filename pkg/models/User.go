package models

type User struct {
	Id                string  `json:"id"`
	Name              string  `json:"name"`
	Email             string  `json:"email"`
	ShopId            string  `json:"shop_id"`
	PhoneNumber       string  `json:"phone_number"`
	SexualOrientation string  `json:"sexual_orientation"`
	Address           Address `json:"address"`
}
