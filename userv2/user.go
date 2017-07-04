package userv2

import "time"

type User struct {
	ID string `json:"id"`
	Blocked bool `json:"blocked"`
	Verified bool `json:"verified"`
	Username string `json:"username"`
	FullName string `json:"fullname"`
	Role Role `json:"role"`
	TimeZone string `json:"timeZone"`
	Locale string `json:"locale"`
	UserAddress UserAddress `json:"userAddress"`
	CreatedAt time.Time `json:"createdAt"`
}

type Role struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

type UserAddress struct {
	Country string `json:"country"`
	State string `json:"state"`
	City string `json:"city"`
	Line string `json:"line"`
	ZipCode string `json:"zipCode"`
}
