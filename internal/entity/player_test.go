package entity

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewPlayer(t *testing.T) {
	testCases := []Player{
		{"", "TheClip", 1, 1},
		{"", "TheClipBR", 5, 5},
		{"", "LongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNicknameLongNickname", 10, 10},
	}

	for _, tc := range testCases {
		t.Run(tc.Nickname, func(t *testing.T) {
			player := NewPlayer(tc.Nickname, tc.Life, tc.Attack)

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

			// Test Attack value
			if player.Attack != tc.Attack {
				t.Errorf("Expected player.Attack = %d; got %d", tc.Attack, player.Attack)
			}
			if player.Attack < 1 || player.Attack > 10 {
				t.Errorf("Expected player.Attack to be in range [1, 10]; got %d", player.Attack)
			}
		})
	}
}
