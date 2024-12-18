package model

import (
	"gotinder/util"
	"log"
	"time"

	"github.com/google/uuid"
)

type SwipeDaily struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	SwipeDate  time.Time `json:"swipe_date"`
	SwipeCount int       `json:"swipe_count"`
}

func SwipeDailyCreate(sw Swipe) error {
	var swipeDaily SwipeDaily
	query := `SELECT user_id, swipe_count FROM daily_swipes WHERE user_id = $1 AND swipe_date = NOW()::DATE`
	row := DB.QueryRow(query, sw.UserID)
	err := row.Scan(&swipeDaily.UserID, &swipeDaily.SwipeCount)
	if err != nil {
		tm := util.GetNow()
		qINSERT := `INSERT INTO daily_swipes (user_id, swipe_date, swipe_count) VALUES ($1, $2, 1)`
		err := DB.QueryRow(qINSERT, sw.UserID, time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, time.UTC)).Err()
		if err != nil {
			log.Printf("Failed insert daily_swipes: %s", err.Error())
		}
	} else {
		newCount := swipeDaily.SwipeCount + 1
		qUPDATE := `UPDATE daily_swipes set swipe_count = $1 WHERE user_id = $2`
		err := DB.QueryRow(qUPDATE, newCount, sw.UserID).Err()
		if err != nil {
			log.Printf("Failed update daily_swipes: %s", err.Error())
		}
	}
	return nil
}
