package models

import "encoding/json"

type User struct {
	ID              int      `json:"id"`
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	Bio             string   `json:"bio"`
	PhoneNumber     string   `json:"phoneNumber"`
	SocialMediaURLs string   `json:"socialMediaUrls"`
	Templates       []string `json:"templates"`
}

func (user *User) ToJson() string {
	jbytes, err := json.Marshal(user)

	if err != nil {
		return "Something went wrong while encoding the user details."
	}

	return string(jbytes)
}
