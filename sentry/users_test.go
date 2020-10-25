package sentry

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUserUnmarshal(t *testing.T) {
	data := []byte(`{
		"username": "test@example.com",
		"lastLogin": "2020-01-02T00:00:00.000000Z",
		"isSuperuser": false,
		"emails": [
			{
				"is_verified": true,
				"id": "1",
				"email": "test@example.com"
			}
		],
		"isManaged": false,
		"experiments": {},
		"lastActive": "2020-01-03T00:00:00.000000Z",
		"isStaff": false,
		"identities": [],
		"id": "1",
		"isActive": true,
		"has2fa": false,
		"name": "John Doe",
		"avatarUrl": "https://secure.gravatar.com/avatar/55502f40dc8b7c769880b10874abc9d0?s=32&d=mm",
		"dateJoined": "2020-01-01T00:00:00.000000Z",
		"options": {
			"timezone": "UTC",
			"stacktraceOrder": -1,
			"language": "en",
			"clock24Hours": false
		},
		"flags": {
			"newsletter_consent_prompt": false
		},
		"avatar": {
			"avatarUuid": null,
			"avatarType": "letter_avatar"
		},
		"hasPasswordAuth": true,
		"email": "test@example.com"
	}`)

	var user User
	err := json.Unmarshal(data, &user)
	assert.NoError(t, err)

	assert.Equal(t, User{
		ID:              "1",
		Name:            "John Doe",
		Username:        "test@example.com",
		Email:           "test@example.com",
		AvatarURL:       "https://secure.gravatar.com/avatar/55502f40dc8b7c769880b10874abc9d0?s=32&d=mm",
		IsActive:        true,
		HasPasswordAuth: true,
		IsManaged:       false,
		DateJoined:      time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		LastLogin:       time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
		Has2FA:          false,
		LastActive:      time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC),
		IsSuperuser:     false,
		IsStaff:         false,
		Avatar: UserAvatar{
			AvatarType: "letter_avatar",
			AvatarUUID: nil,
		},
		Emails: []UserEmail{
			{
				ID:         "1",
				Email:      "test@example.com",
				IsVerified: true,
			},
		},
	}, user)
}
