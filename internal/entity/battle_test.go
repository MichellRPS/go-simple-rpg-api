package entity

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewBattle(t *testing.T) {
	testCases := []Battle{
		{"", "8c7891c9-0c57-432c-9bd3-8116668bcf5c", "9def6009-e41f-4ad3-b08d-20f5990642ae", 0},
		{"", "5781eb61-35c1-43b2-9bd1-23c2836ac69c", "2116faf6-6da5-461b-a3db-c83263422e11", 0},
		{"", "97ae2f38-7d7e-4da3-bfe8-01866717b69c", "c84b1cee-dcea-4f66-b6ff-018a8949ba32", 0},
	}

	for _, tc := range testCases {
		t.Run(tc.EnemyID+tc.PlayerID, func(t *testing.T) {
			battle := NewBattle(tc.EnemyID, tc.PlayerID)

			// Test type
			if battle == nil {
				t.Fatalf("battle is nil")
			}

			// Test ID value
			if _, err := uuid.Parse(battle.ID); err != nil {
				t.Errorf("Expected battle.ID to be a valid UUID; got %s; %s", battle.ID, err.Error())
			}

			// Test Enemy ID value
			if _, err := uuid.Parse(battle.EnemyID); err != nil {
				t.Errorf("Expected battle.EnemyID to be a valid UUID; got %s; %s", battle.EnemyID, err.Error())
			}

			// Test Player ID value
			if _, err := uuid.Parse(battle.PlayerID); err != nil {
				t.Errorf("Expected battle.PlayerID to be a valid UUID; got %s; %s", battle.PlayerID, err.Error())
			}

			// Test DiceThrown value
			if battle.DiceThrown < 1 || battle.DiceThrown > 6 {
				t.Errorf("Expected battle.DiceThrown to be in range [1, 6]; got %d", battle.DiceThrown)
			}
		})
	}
}
