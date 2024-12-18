package model

import (
	"gotinder/config"
	"gotinder/util"
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
	query := `INSERT INTO swipes
		(user_id,target_user_id,swipe_type)
		VALUES
		($1, $2, $3)`
	if err := DB.QueryRow(query, swipe.UserID, swipe.TargetUserID, swipe.SwipeType).Err(); err != nil {
		return err
	}

	if swipe.SwipeType == "like" {
		var matchx int
		qCheckLike := `SELECT 1 
			FROM swipes 
			WHERE target_user_id = $1 
			AND swipe_type = 'like'`
		if err := DB.QueryRow(qCheckLike, swipe.UserID).Scan(&matchx); err != nil {
			return err
		}
		if matchx == 1 {
			qInsert := `INSERT INTO matches
			(user_id, matched_user_id, created_at)
			VALUES
			($1, $2, NOW())`
			if err := DB.QueryRow(qInsert, swipe.UserID, swipe.TargetUserID).Err(); err != nil {
				return err
			}
			// TODO
			// send notification mail
			if config.GetConf().IsQueue {
				// using queue
			} else {
				ids := []string{swipe.UserID.String(), swipe.TargetUserID.String()}
				emails, err := UserGetEmails(ids)
				if err == nil {
					for _, email := range emails {
						go util.SendMail(map[string]interface{}{
							"to":      email,
							"subject": "Confirmation Email",
							"message": "Match",
						})
					}
				}
			}
		}
	}

	SwipeDailyCreate(Swipe{UserID: swipe.UserID})
	SwipeHistoryCreate(SwipeHistory{
		UserID:       swipe.UserID,
		TargetUserID: swipe.TargetUserID,
	})
	return nil
}

func SwipeData(user_id string) ([]User, error) {
	var users []User
	query := `SELECT id, email, username, bio, profile_image  
		FROM users 
		WHERE id NOT IN (
			SELECT target_user_id 
			FROM swipe_history 
			WHERE user_id = $1 
			AND DATE(swipe_date) = CURRENT_DATE
		)
		AND id != $2`
	rows, err := DB.Query(query, user_id, user_id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Email, &user.Username, &user.Bio, &user.ProfileImage); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
