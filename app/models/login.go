package models

import "github.com/revel/revel"

/**
Login Model
 */
type Login struct {
	Email      string
	Password   string
	RememberMe bool
}

/**
Validate Object
 */
func (l Login) Validate(v *revel.Validation) {
	v.Check(l.Email,
		revel.Required{},
		revel.ValidEmail(),
	)

	v.Check(l.Password,
		revel.Required{},
		revel.MinSize{5},
		revel.MaxSize{20},
	).Message("Password should 8 character long")
}
