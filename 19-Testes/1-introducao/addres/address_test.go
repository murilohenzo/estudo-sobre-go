package address

import "testing"

type TestScenery struct {
	received string
	expect   string
}

func TestTypeAddress(it *testing.T) {
	testSceneriesMock := []TestScenery{
		{"Rua ABC", "Rua"},
		{"RUA ABC", "Rua"},
		{"AVENIDA ABC", "Avenida"},
		{"Avenida ABC", "Avenida"},
		{"ABC", "Invalid type"},
		{"", "Invalid type"},
	}

	it.Run("Validation type address", func(t *testing.T) {
		for _, scenery := range testSceneriesMock {
			got := TypeAddress(scenery.received)
			if got != scenery.expect {
				t.Errorf("\nO tipo recebido eh diferente do esperado!"+
					"\nexpect = [%s]\nreceived = [%s]", scenery.expect, got)
			}
		}
	})

}
