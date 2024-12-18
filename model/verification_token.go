package model

import (
	"errors"
	"gotinder/util"
	"log"
	"time"

	"github.com/google/uuid"
)

type VerifToken struct {
	ID    uuid.UUID `json:"id" bson:"_id,omitempty"`
	Email string    `json:"email" bson:"email"`
	Token string    `json:"token" bson:"token"`
	Exp   time.Time `json:"expired" bson:"expired"`
}

func VerifTokenCreate(vtoken VerifToken) error {
	query := `INSERT INTO verification_token (email, token, expired) 
		VALUES ($1, $2, $3) RETURNING id`
	token := uuid.New().String()

	err := DB.QueryRow(query, vtoken.Email, token, util.GetNow().Add(1*time.Hour)).Scan(&vtoken.ID)
	if err != nil {
		log.Println("Error creating user:", err)
		return err
	}
	// conf := config.GetConf()
	// err = util.SendMail(map[string]interface{}{
	// 	"to":      vtoken.Email,
	// 	"subject": "Confirmation Email",
	// 	"message": conf.URLFront + "verification/" + token,
	// })
	return err
}

func VerifTokenConfirm(token string) error {
	var vt VerifToken
	var us User
	query := `SELECT email, token, expired FROM verification_token WHERE token = $1`
	row := DB.QueryRow(query, token)
	if err := row.Scan(&vt.Email, &vt.Token, &vt.Exp); err != nil {
		return errors.New("token does not exist")
	}

	log.Println("db email", vt.Email)
	log.Println("db", vt.Exp)
	log.Println("system", util.GetNow())
	if vt.Exp.Before(util.GetNow()) {
		return errors.New("token has expired")
	}
	queryUser := `SELECT id, email FROM users WHERE email = $1`
	rowUser := DB.QueryRow(queryUser, vt.Email)
	if err := rowUser.Scan(&us.ID, &us.Email); err != nil {
		return errors.New("user does not exist")
	}

	if err := UserUpdate(User{
		ID:    us.ID,
		Email: us.Email,
	}); err != nil {
		return err
	} else {
		return err
	}
}
