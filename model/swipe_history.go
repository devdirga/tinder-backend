package model

import (
	"time"

	"github.com/google/uuid"
)

type SwipeHistory struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	TargetUserID uuid.UUID `json:"target_user_id"`
	SwipeDate    time.Time `json:"swipe_date"`
}

func SwipeHistoryCreate(swipeHistory SwipeHistory) error {
	qINSERT := `INSERT INTO swipe_history
		(user_id, target_user_id, swipe_date)
		VALUES
		($1, $2, CURRENT_DATE)`
	if err := DB.QueryRow(qINSERT, swipeHistory.UserID, swipeHistory.TargetUserID).Err(); err != nil {
		return err
	}
	return nil
}
