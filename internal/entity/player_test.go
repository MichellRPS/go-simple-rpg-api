package entity

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewPlayer(t *testing.T) {
	testCases := []Player{
		{"", "TheClip", 1, "8c3910f5-7248-40c7-b0c2-38446f328ce9", 8},
		{"", "TheClipBR", 5, "9b48bd04-d944-463b-be4e-a4eef73d7e15", 7},
		{"", "LongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNickname", 10, "476544a3-029b-4f49-8631-e13ada552ec7", 10},
	}

	for _, tc := range testCases {
		t.Run(tc.Nickname, func(t *testing.T) {
			player := NewPlayer(tc.Nickname, tc.Life, tc.WeaponID, tc.WeaponDurability)

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

			// Test WeaponDurability value
			if player.WeaponDurability < 1 || player.WeaponDurability > 10 {
				t.Errorf("Expected player.WeaponDurability to be in range [1, 10]; got %d", player.WeaponDurability)
			}
		})
	}
}
