package model

import (
	"errors"
	"gotinder/config"
	"gotinder/producer"
	"gotinder/util"
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
	query := `INSERT INTO verification_token
		(email, token, expired) 
		VALUES
		($1, $2, $3) RETURNING id`
	token := uuid.New().String()
	err := DB.QueryRow(query,
		vtoken.Email,
		token,
		util.GetNow().Add(1*time.Hour)).Scan(&vtoken.ID)
	if err != nil {
		return err
	}

	if config.GetConf().IsQueue {
		conf := config.GetConf()
		req := util.RequestMessage{
			To:      vtoken.Email,
			Subject: "Confirm Email",
			Message: conf.URL + "verification/" + token,
		}
		err = producer.ProducerMessage(req)

	} else {
		conf := config.GetConf()
		err = util.SendMail(map[string]interface{}{
			"to":      vtoken.Email,
			"subject": "Confirmation Email",
			"message": conf.URL + "verification/" + token,
		})
	}

	return err
}

func VerifTokenConfirm(token string) error {
	var vt VerifToken
	var us User
	query := `SELECT
		email,
		token,
		expired
		FROM verification_token 
		WHERE token = $1`
	row := DB.QueryRow(query, token)
	if err := row.Scan(&vt.Email, &vt.Token, &vt.Exp); err != nil {
		return errors.New("token does not exist")
	}
	if vt.Exp.Before(util.GetNow()) {
		return errors.New("token has expired")
	}

	rowUser := DB.QueryRow(`SELECT 
		id,
		email 
		FROM users 
		WHERE email = $1`, vt.Email)
	if err := rowUser.Scan(&us.ID, &us.Email); err != nil {
		return errors.New("user does not exist")
	}

	if err := UserUpdate(User{ID: us.ID, Email: us.Email}); err != nil {
		return err
	} else {
		return err
	}
}
