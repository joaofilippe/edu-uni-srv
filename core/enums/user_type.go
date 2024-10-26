package enums

import "fmt"

type UserType int

const (
	Student UserType = iota
	Professor
	Guardian
	Administrator
)

func (u UserType) String() string {
	return [...]string{"Student", "Professor", "Guardian", "Administrator"}[u]
}

func ParseUserType(s string) (UserType, error) {
	switch s {
	case "Student":
		return Student, nil
	case "Professor":
		return Professor, nil
	case "Guardian":
		return Guardian, nil
	case "Administrator":
		return Administrator, nil
	default:
		return -1, fmt.Errorf("invalid UserType: %s", s)
	}
}