package address

import "testing"

type TestScenery struct {
	received string
	expect   string
}

func TestTypeAddress(t *testing.T) {
	testSceneries := []TestScenery{
		{"Rua ABC", "Rua"},
		{"RUA ABC", "Rua"},
		{"AVENIDA ABC", "Avenida"},
		{"Avenida ABC", "Avenida"},
		{"ABC", "Invalid type"},
		{"", "Invalid type"},
	}

	for _, scenery := range testSceneries {
		receivedTypeAddress := TypeAddress(scenery.received)
		if receivedTypeAddress != scenery.expect {
			t.Errorf("\nO tipo recebido eh diferente do esperado!"+
				"\nexpect = [%s]\nreceived = [%s]", scenery.expect, receivedTypeAddress)
		}
	}
}
