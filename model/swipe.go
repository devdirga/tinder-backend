package model

import (
	"log"
	"time"

	"github.com/google/uuid"
)

type Swipe struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	TargetUserID uuid.UUID `json:"target_user_id"`
	SwipeType    string    `json:"swipe_type"`
	CreatedAt    time.Time `json:"created_at"`
}

func SwipeCreate(swipe Swipe) error {
	query := `INSERT INTO swipes (user_id, target_user_id, swipe_type) VALUES ($1, $2, $3)`
	err := DB.QueryRow(query, swipe.UserID, swipe.TargetUserID, swipe.SwipeType).Err()
	if err != nil {
		log.Println("Error creating swape:", err)
		return err
	}
	SwipeDailyCreate(Swipe{UserID: swipe.UserID})
	return nil
}

func SwipeData(user_id string) ([]User, error) {
	var users []User
	query := `SELECT id, username, bio, profile_image  
		FROM users 
		WHERE id NOT IN (
			SELECT target_user_id 
			FROM swipes 
			WHERE user_id = $1 
			AND DATE(created_at) = CURRENT_DATE
		)
		AND id != $2`
	rows, err := DB.Query(query, user_id, user_id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Bio, &user.ProfileImage); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
