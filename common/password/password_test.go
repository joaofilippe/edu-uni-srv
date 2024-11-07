package password

import (
    "testing"
)

func TestHashPassword(t *testing.T) {
    password := "mysecretpassword"
    hash, err := HashPassword(password)
    if err != nil {
        t.Fatalf("Erro ao gerar hash da senha: %v", err)
    }

    if len(hash) == 0 {
        t.Fatalf("Hash gerado está vazio")
    }
}

func TestCheckPasswordHash(t *testing.T) {
    password := "mysecretpassword"
    hash, err := HashPassword(password)
    if err != nil {
        t.Fatalf("Erro ao gerar hash da senha: %v", err)
    }

    if !CheckPasswordHash(password, hash) {
        t.Fatalf("A senha não corresponde ao hash")
    }

    wrongPassword := "wrongpassword"
    if CheckPasswordHash(wrongPassword, hash) {
        t.Fatalf("A senha incorreta foi aceita")
    }
}