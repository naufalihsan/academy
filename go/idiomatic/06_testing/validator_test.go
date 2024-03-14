package main

import "testing"

func TestIsValidEmail(t *testing.T) {
	email := "email@example.com"
	if !IsValidEmail(email) {
		t.Errorf("IsValidEmail(%v) = false, expected true \n", email)
	}
}

func TestIsValidEmailTable(t *testing.T) {
	table := []struct {
		email    string
		expected bool
	}{
		{"email@invalid", false},
		{"email@valid.com", true},
	}

	for _, data := range table {
		res := IsValidEmail(data.email)
		if res != data.expected {
			t.Errorf("IsValidEmail(%v) = %t, expected %t \n", data.email, res, data.expected)
		}
	}
}
