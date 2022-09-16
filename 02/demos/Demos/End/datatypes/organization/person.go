package organization

import (
	"errors"
	"fmt"
	"strings"
)

type TwitterHandler string

func (th TwitterHandler) RedirectURL() string {
	return ""
}

type Identifiable interface {
	ID() string
}
type Name struct {
	first string
	last  string
}

type Employee struct {
	Name
}

type Person struct {
	Name
	twitterHandler TwitterHandler
}

func NewPerson(firstName, lastName string) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
	}
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.first, p.last)
}

func (p *Person) ID() string {
	return "12345"
}

func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("twitter handler must start with an @ symbol")
	}

	p.twitterHandler = handler
	return nil
}

func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}
