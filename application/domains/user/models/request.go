package models

import (
	"errors"

	"github.com/go-playground/validator"
)

type RegisterReq struct {
	FullName        string `json:"full_name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

func (req *RegisterReq) ValidateRequest() error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			switch field {
			case "FullName":
				return errors.New("full name is required")
			case "Email":
				if err.Tag() == "required" {
					return errors.New("email is required")
				} else if err.Tag() == "email" {
					return errors.New("email format is invalid")
				}
			case "Password":
				if err.Tag() == "required" {
					return errors.New("password is required")
				} else if err.Tag() == "min" {
					return errors.New("password must be at least 8 characters long")
				}
			case "ConfirmPassword":
				if err.Tag() == "required" {
					return errors.New("confirm password is required")
				} else if err.Tag() == "eqfield" {
					return errors.New("confirm password must match the password")
				}
			}
		}
	}
	return nil
}

type LoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (req *LoginReq) ValidateRequest() error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			switch field {
			case "Email":
				if err.Tag() == "required" {
					return errors.New("email is required")
				} else if err.Tag() == "email" {
					return errors.New("email format is invalid")
				}
			case "Password":
				return errors.New("password is required")
			}
		}
	}
	return nil
}
