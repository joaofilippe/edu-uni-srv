package enums

import "fmt"

type UserType int

const (
	Aluno UserType = iota
	Professor
	Responsavel
)

func (u UserType) String() string {
	return [...]string{"Aluno", "Professor", "Responsavel"}[u]
}

func ParseUserType(s string) (UserType, error) {
	switch s {
	case "Aluno":
		return Aluno, nil
	case "Professor":
		return Professor, nil
	case "Responsavel":
		return Responsavel, nil
	default:
		return -1, fmt.Errorf("invalid UserType: %s", s)
	}
}