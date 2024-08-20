package entity

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewPlayer(t *testing.T) {
	testCases := []Player{
		{"", "TheClip", 1, "1ccaac4a-3be5-442f-bba0-e7d78dfdc4a3"},
		{"", "TheClipBR", 5, "177b7b6f-ba70-4dc9-aae6-a92c07a02a2e"},
		{"", "LongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNickname", 10, "0a96a984-2f14-43d6-9563-1ec193f6e932"},
	}

	for _, tc := range testCases {
		t.Run(tc.Nickname, func(t *testing.T) {
			player := NewPlayer(tc.Nickname, tc.Life, tc.WeaponID)

			// Test type
			if player == nil {
				t.Fatalf("player is nil")
			}			

			// Test ID value
			if _, err := uuid.Parse(player.ID); err != nil {
				t.Errorf("Expected player.ID to be a valid UUID; got %s; %s", player.ID, err.Error())
			}

			// Test Nickname value
			if player.Nickname != tc.Nickname {
				t.Errorf("Expected player.Nickname = %s; got %s", tc.Nickname, player.Nickname)
			}
			if len(player.Nickname) > 255 {
				t.Errorf("Expected len(player.Nickname) <= 255; got: %d", len(player.Nickname))
			}

			// Test Life value
			if player.Life != tc.Life {
				t.Errorf("Expected player.Life = %d; got %d", tc.Life, player.Life)
			}
			if player.Life < 1 || player.Life > 10 {
				t.Errorf("Expected player.Life to be in range [1, 10]; got %d", player.Life)
			}

			// Test WeaponID value
			if _, err := uuid.Parse(player.WeaponID); err != nil {
				t.Errorf("Expected player.WeaponID to be a valid UUID; got %s; %s", player.WeaponID, err.Error())
			}
		})
	}
}
