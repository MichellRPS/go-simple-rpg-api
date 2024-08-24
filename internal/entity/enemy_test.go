package entity

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewEnemy(t *testing.T) {
	testCases := []Enemy{
		{"", "A", 0, "8c3910f5-7248-40c7-b0c2-38446f328ce9", 8},
		{"", "ABC", 0, "9b48bd04-d944-463b-be4e-a4eef73d7e15", 7},
		{"", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", 0, "476544a3-029b-4f49-8631-e13ada552ec7", 10},
	}

	for _, tc := range testCases {
		t.Run(tc.Nickname, func(t *testing.T) {
			enemy := NewEnemy(tc.Nickname, tc.WeaponID, tc.WeaponDurability)

			// Test type
			if enemy == nil {
				t.Fatalf("enemy is nil")
			}

			// Test ID value
			if _, err := uuid.Parse(enemy.ID); err != nil {
				t.Errorf("Expected enemy.ID to be a valid UUID; got %s; %s", enemy.ID, err.Error())
			}

			// Test Nickname value
			if enemy.Nickname != tc.Nickname {
				t.Errorf("Expected enemy.Nickname = %s; got %s", tc.Nickname, enemy.Nickname)
			}
			if len(enemy.Nickname) > 255 {
				t.Errorf("Expected len(enemy.Nickname) <= 255; got: %d", len(enemy.Nickname))
			}

			// Test Life value
			if enemy.Life < 1 || enemy.Life > 10 {
				t.Errorf("Expected enemy.Life to be in range [1, 10]; got %d", enemy.Life)
			}

			// Test WeaponID value
			if _, err := uuid.Parse(enemy.WeaponID); err != nil {
				t.Errorf("Expected enemy.WeaponID to be a valid UUID; got %s; %s", enemy.WeaponID, err.Error())
			}

			// Test WeaponDurability value
			if enemy.WeaponDurability < 1 || enemy.WeaponDurability > 10 {
				t.Errorf("Expected enemy.WeaponDurability to be in range [1, 10]; got %d", enemy.WeaponDurability)
			}
		})
	}
}
