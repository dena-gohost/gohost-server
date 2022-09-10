package validator

import (
	"errors"

	"github.com/dena-gohost/gohost-server/gen/api"
)

type User struct {
	*api.User
}

func (r *User) Register() error {
	if r.User == nil {
		return errors.New("request cannot be empty")
	}
	if r.Id != nil && *r.Id != "" {
		return errors.New("id field cannot be filled")
	}
	if r.FirstName == nil || *r.FirstName == "" {
		return errors.New("first name field cannot be empty")
	}
	if r.LastName == nil || *r.LastName == "" {
		return errors.New("last name field cannot be empty")
	}
	if r.UserName == nil || *r.UserName == "" {
		return errors.New("user name field cannot be empty")
	}
	if r.Email == nil || *r.Email == "" {
		return errors.New("email field cannot be empty")
	}
	if r.Password == nil || *r.Password == "" {
		return errors.New("password field cannot be empty")
	}
	if r.UniversityId == nil || *r.UniversityId == "" {
		return errors.New("university field cannot be empty")
	}
	if r.BirthDate == nil {
		return errors.New("birth_date field cannot be empty")
	}
	if r.Year == nil || *r.Year == 0 {
		return errors.New("year field cannot be empty")
	}
	if r.GenderId == nil || *r.GenderId == "" {
		return errors.New("gender field cannot be empty")
	}
	if r.IconUrl == nil || *r.IconUrl == "" {
		return errors.New("icon_url field cannot be empty")
	}
	if r.InstagramId == nil || *r.InstagramId == "" {
		return errors.New("instagram id field cannot be empty")
	}
	return nil
}

func (r *User) Login() error {
	if r == nil {
		return errors.New("request cannot be empty")
	}
	if r.Email == nil || *r.Email == "" {
		return errors.New("email field cannot be empty")
	}
	if r.Password == nil || *r.Password == "" {
		return errors.New("password field cannot be empty")
	}
	return nil
}
