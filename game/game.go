package game

import "fmt"

const (
	ECitizen  = "Citizen"
	EWarewolf = "Warewolf"
	ETeller   = "Teller"
	EKnight   = "Knight"
)

type Personer interface {
	GetID() int
	GetKind() string
	GetName() string
	Vote(people []Personer) int
	NightAction()
}

type Person struct {
	IsDead bool
	ID     int
	Kind   string
	Name   string
}

func (p *Person) GetID() int {
	return p.ID
}

func (p *Person) GetKind() string {
	return p.Kind
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) NightAction() {}

func (p *Person) Vote(people []Personer) int {
	for _, v := range people {
		fmt.Printf("%d: %s\n", v.GetID(), v.GetName())
	}
	return 1
}

func NewPerson(id int, kind, name string) Personer {
	switch kind {
	case ECitizen:
		return &Citizen{
			Person{
				ID:   id,
				Kind: kind,
				Name: name,
			},
		}
	case EWarewolf:
		return &Warewolf{
			Person{
				ID:   id,
				Kind: kind,
				Name: name,
			},
		}
	case ETeller:
		return &Teller{
			Person{
				ID:   id,
				Kind: kind,
				Name: name,
			},
		}
	case EKnight:
		return &Knight{
			Person{
				ID:   id,
				Kind: kind,
				Name: name,
			},
		}
	}
	return &Citizen{
		Person{
			ID:   id,
			Kind: ECitizen,
			Name: name,
		},
	}
}

type Citizen struct {
	Person
}

type Warewolf struct {
	Person
}

func (w *Warewolf) NightAction() {

}

type Teller struct {
	Person
}

type Knight struct {
	Person
}
