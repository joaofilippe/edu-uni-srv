package userentities

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/edu-uni-srv/domain/enums"
	"github.com/joaofilippe/edu-uni-srv/domain/interfaces"
)

func TestValidateUsername(t *testing.T) {
	tests := []struct {
		email string
		valid bool
	}{
		{"valid@example.com", true},
		{"invalid", false},
		{"@invalid.com", false},
		{"invalid@.com", false},
		{"invalid@com.", false},
		{"invalid@com", false},
		{"valid@sub.example.com", true},
	}

	for _, test := range tests {
		user := &User{
			id:          uuid.New(),
			username:    "username",
			password:    "password",
			email:       test.email,
			userType:    enums.Student,
			userDetails: interfaces.IUserDetails(nil),
			createdAt:   time.Now(),
			updatedAt:   time.Time{},
			active:      true,
		}

		if user.validateEmail() != test.valid {
			t.Errorf("validateEmail(%s) = %v; want %v", test.email, !test.valid, test.valid)
		}
	}
}
