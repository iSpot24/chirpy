package auth

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	cases := []struct {
		input string
	}{
		{
			input: "test",
		},
	}

	for _, c := range cases {
		actual, err := HashPassword(c.input)

		if err != nil {
			t.Errorf("Hashing failed. Input: %v", c.input)
		}

		if actual == c.input {
			t.Errorf("Input equal to hash. Input: %v, Actual: %v", c.input, actual)
		}
	}
}

func TestCheckPasswordHash(t *testing.T) {
	expectedPass := "expected password"
	actual, err := bcrypt.GenerateFromPassword([]byte(expectedPass), bcrypt.DefaultCost)

	if err != nil {
		t.Errorf("Hashing failed. Input: %v", expectedPass)
	}

	cases := []struct {
		input    string
		expected error
	}{
		{
			input:    expectedPass,
			expected: nil,
		},
		{
			input:    "unexpected pass",
			expected: errors.New("crypto/bcrypt: hashedPassword is not the hash of the given password"),
		},
	}

	for _, c := range cases {
		err := CheckPasswordHash(actual, c.input)

		if err != nil && c.expected == nil || err == nil && c.expected != nil {
			t.Errorf("Comparing hash failed. Actual: %v, Expected: %v", err, c.expected)
		}
	}
}

func TestGetBearerToken(t *testing.T) {
	expectedPass := "expected password"
	token, err := HashPassword(expectedPass)

	if err != nil {
		t.Errorf("Hashing failed. Input: %v", expectedPass)
	}

	header1 := http.Header{}
	header1.Set("Authorization", "Bearer "+string(token))
	header2 := http.Header{}
	header2.Set("Authorization", string(token))

	type expectedStruct struct {
		token string
		err   error
	}

	cases := []struct {
		input    http.Header
		expected expectedStruct
	}{
		{
			input: header1,
			expected: expectedStruct{
				token: string(token),
				err:   errors.New(""),
			},
		},
		{
			input: header2,
			expected: expectedStruct{
				token: "",
				err:   errors.New("invalid authorization header"),
			},
		},
		{
			input: http.Header{},
			expected: expectedStruct{
				token: "",
				err:   errors.New("missing authorization"),
			},
		},
	}

	for _, c := range cases {
		token, err := GetHeaderToken(c.input, "Authorization", "Bearer")

		if err != nil && c.expected.err.Error() != err.Error() || err == nil && c.expected.err.Error() != "" {
			t.Errorf("get bearer token has errors. Actual: %v, Expected: %v", expectedStruct{
				token: token,
				err:   err,
			}, c.expected)
		}

		if token != c.expected.token {
			t.Errorf("get bearer token failed. Actual: %v, Expected: %v", expectedStruct{
				token: token,
				err:   err,
			}, c.expected)
		}
	}
}

func TestValidateJWT(t *testing.T) {
	userID := uuid.New()
	validToken, _ := MakeJWT(userID, "secret", time.Hour)

	tests := []struct {
		name        string
		tokenString string
		tokenSecret string
		wantUserID  uuid.UUID
		wantErr     bool
	}{
		{
			name:        "Valid token",
			tokenString: validToken,
			tokenSecret: "secret",
			wantUserID:  userID,
			wantErr:     false,
		},
		{
			name:        "Invalid token",
			tokenString: "invalid.token.string",
			tokenSecret: "secret",
			wantUserID:  uuid.Nil,
			wantErr:     true,
		},
		{
			name:        "Wrong secret",
			tokenString: validToken,
			tokenSecret: "wrong_secret",
			wantUserID:  uuid.Nil,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUserID, err := ValidateJWT(tt.tokenString, tt.tokenSecret)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUserID != tt.wantUserID {
				t.Errorf("ValidateJWT() gotUserID = %v, want %v", gotUserID, tt.wantUserID)
			}
		})
	}
}
