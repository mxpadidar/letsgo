package servicesimpl

import (
	"strings"
	"testing"
)

func TestBCryptHash(t *testing.T) {
	service := NewBcryptHashService()

	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{
			name:     "Valid password",
			password: "mxp",
			wantErr:  false,
		},
		{
			name:     "Empty password",
			password: "",
			wantErr:  false,
		},
		{
			name:     "Long password",
			password: strings.Repeat("a", 72), // bcrypt max length is 72 bytes
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := service.HashPassword(tt.password)

			// Check error expectation
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Verify hash properties
			if !tt.wantErr {
				if len(hash) < 60 {
					t.Errorf("Hash length too short: got %d, want >= 60", len(hash))
				}
				if !strings.HasPrefix(hash, "$2a$") && !strings.HasPrefix(hash, "$2b$") {
					t.Errorf("Invalid hash prefix: %s", hash[:4])
				}
			}
		})
	}
}

func TestBCryptCompare(t *testing.T) {
	service := NewBcryptHashService()

	tests := []struct {
		name      string
		password  string
		wantMatch bool
		wrongPass string
	}{
		{
			name:      "Correct password",
			password:  "mxp",
			wantMatch: true,
			wrongPass: "wrong",
		},
		{
			name:      "Empty password",
			password:  "",
			wantMatch: true,
			wrongPass: "wrong",
		},
		{
			name:      "Long password",
			password:  strings.Repeat("a", 72),
			wantMatch: true,
			wrongPass: "wrong",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// First hash the password
			hash, err := service.HashPassword(tt.password)
			if err != nil {
				t.Fatalf("Failed to hash password: %v", err)
			}

			// Test correct password
			if match := service.ComparePassword(hash, tt.password); match != tt.wantMatch {
				t.Errorf("ComparePassword() with correct password = %v, want %v", match, tt.wantMatch)
			}

			// Test wrong password
			if match := service.ComparePassword(hash, tt.wrongPass); match {
				t.Errorf("ComparePassword() with wrong password should return false")
			}
		})
	}
}
