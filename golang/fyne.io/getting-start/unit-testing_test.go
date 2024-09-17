package main

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

func TestGreeting(t *testing.T) {
	// Pega a função para realizar o teste
	out, in := makeUI()

	if out.Text != "Hello world!" {
		t.Error("Incorrect inital gretting")
	}
	// Teste definindo uma variavel para realizar o reste
	test.Type(in, "Felipe")
	if out.Text != "Hello Felipe!" {
		t.Error("Incorrect user gretting")
	}
}
