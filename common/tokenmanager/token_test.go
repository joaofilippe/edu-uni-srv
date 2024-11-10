package tokenmanager

import (
	"fmt"
	"os"
	"testing"
)

func TestValidateToken(t *testing.T) {
	os.Setenv("SECRET_KEY", "adflaskdfahkjl21o43214y")
	token := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzE1MDU2MzQsInVzZXJJRCI6IjFjY2YxYzdiLTE4ZmYtNGVlMC04ODA5LTk3ODMxOWMyZTI1YiIsInVzZXJUeXBlIjoiQWRtaW5pc3RyYXRvciIsInVzZXJUeXBlSUQiOiIzMDZkZTllYS1iZTJmLTRjNjEtOTE1ZS1kMTZkYTYzOGVhOGIifQ.EM5ei67kifR29_cZksnORHJ6p8OZ6HmMfGeSi2178DU`

	c, err := ValidateToken(token)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	fmt.Println(c["userID"])
}
