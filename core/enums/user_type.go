package enums

import "fmt"

type UserType int

const (
	Student UserType = iota
	Teacher
	Guardian
	Administrator
)

func (u UserType) String() string {
	return [...]string{"Student", "Teacher", "Guardian", "Administrator"}[u]
}

func ParseUserType(s string) (UserType, error) {
	switch s {
	case "Student":
		return Student, nil
	case "Teacher":
		return Teacher, nil
	case "Guardian":
		return Guardian, nil
	case "Administrator", "admin":
		return Administrator, nil
	default:
		return -1, fmt.Errorf("invalid UserType: %s", s)
	}
}
