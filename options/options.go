package options

import (
	"errors"
	"fmt"
	"regexp"
)

type Customer struct {
	id            string
	email         string
	firstName     *string
	lastName      *string
	loyaltyPoints int
}

type Option = func(c *Customer)

var ErrEmptyId = errors.New("id cannot be empty")
var ErrInvalidEmail = "email is invalid: %s"
var ErrEmptyFirstName = errors.New("first name cannot be empty")
var ErrEmptyLastName = errors.New("last lane cannot be empty")

func NewCustomer(id, email string, options ...Option) (*Customer, error) {
	c := &Customer{
		id:    id,
		email: email,

		// default values
		loyaltyPoints: 100,
	}

	for _, opt := range options {
		opt(c)
	}

	if err := c.validate(); err != nil {
		return nil, err
	}
	return c, nil
}

func WithName(firstName, lastName string) Option {
	return func(c *Customer) {
		c.firstName = &firstName
		c.lastName = &lastName
	}
}

func WithLoyaltyPoints(loyaltyPoints int) Option {
	return func(c *Customer) {
		c.loyaltyPoints = loyaltyPoints
	}
}

func (c *Customer) validate() error {
	if c.id == "" {
		return ErrEmptyId
	}
	if !c.IsValidEmail(c.email) {

		return fmt.Errorf(ErrInvalidEmail, c.email)
	}
	if c.firstName != nil && *c.firstName == "" {
		return ErrEmptyFirstName
	}
	if c.lastName != nil && *c.lastName == "" {
		return ErrEmptyLastName
	}
	return nil
}

func (c *Customer) IsValidEmail(email string) bool {

	emailPattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailPattern.MatchString(email)
}
