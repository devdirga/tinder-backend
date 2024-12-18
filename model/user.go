package model

import (
	"database/sql"
	"errors"
	"fmt"
	"gotinder/util"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID `json:"id"`
	Username         string    `json:"username"`
	Email            string    `json:"email"`
	Password         string    `json:"password"`
	Bio              string    `json:"bio"`
	ProfileImage     string    `json:"profile_image"`
	SubscriptionType string    `json:"subscription_type"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	SwipeCount       int       `json:"swipe_count"`
}

func UserCreate(user User) error {
	query := `INSERT INTO users
		(username, email, password, bio, profile_image, subscription_type)
		VALUES
		($1, $2, $3, $4, $5, $6) RETURNING id`
	if err := DB.QueryRow(query,
		user.Username,
		user.Email,
		user.Password,
		user.Bio,
		user.ProfileImage,
		user.SubscriptionType).Scan(&user.ID); err != nil {
		return err
	}
	return nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	query := `SELECT 
		u.id,
		u.username,
		u.email,
		u.subscription_type,
		u.created_at,
		COALESCE(ds.swipe_count, 0) AS swipe_count,
		u.bio,
		u.profile_image,
		u.password 
		FROM users u
		left join daily_swipes ds on (ds.user_id = u.id AND ds.swipe_date = NOW()::DATE) 
		WHERE email = $1`
	row := DB.QueryRow(query, email)
	err := row.Scan(&user.ID,
		&user.Username,
		&user.Email,
		&user.SubscriptionType,
		&user.CreatedAt,
		&user.SwipeCount,
		&user.Bio,
		&user.ProfileImage,
		&user.Password,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func UserUpdate(user User) error {
	query := `UPDATE users 
		SET email = $1,
		email_verif = $2
		WHERE id = $3`
	err := DB.QueryRow(query, user.Email, util.GetNow(), user.ID).Err()
	if err != nil {
		return err
	}
	return nil
}

func UserUpdateByEmail(user User) error {
	query := `UPDATE users 
		SET 
		bio = $1,
		profile_image = $2
		WHERE email = $3`
	err := DB.QueryRow(query, user.Bio, user.ProfileImage, user.Email).Err()
	if err != nil {
		return err
	}
	return nil
}

func UserGetEmails(ids []string) ([]string, error) {
	// Build query with placeholders
	placeholders := make([]string, len(ids))
	for i := range ids {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}
	query := fmt.Sprintf("SELECT email FROM users WHERE id IN (%s)", strings.Join(placeholders, ", "))

	// Execute query
	rows, err := DB.Query(query, convertToInterfaceSlice(ids)...)
	if err != nil {
		log.Fatal("Failed to execute query:", err)
	}
	defer rows.Close()

	var res []string
	for rows.Next() {
		var mail string
		if err := rows.Scan(&mail); err != nil {
			return nil, err
		}
		res = append(res, mail)
	}

	return res, nil
}

func convertToInterfaceSlice(slice []string) []interface{} {
	result := make([]interface{}, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}
