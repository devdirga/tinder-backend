package database

import (
	"gotinder/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migrate() {
	dsn := config.GetConf().DB
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	err = db.Exec(`
		-- CREATE ENUM subscription_type
		DO $$ 
			BEGIN
				IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'subscription_type') THEN
					CREATE TYPE subscription_type AS ENUM ('free', 'premium');
				END IF;
		END $$;

		-- CREATE ENUM swipe_type
		DO $$ 
			BEGIN
				IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'swipe_type') THEN
					CREATE TYPE swipe_type AS ENUM ('like', 'pass');
				END IF;
		END $$;

		-- CREATE EXT 
		CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
		
		--CREATE TABLE users
		CREATE TABLE IF NOT EXISTS public.users (
			id uuid NOT NULL DEFAULT uuid_generate_v4(),
			username varchar(255) NOT NULL,
			email varchar(255) NOT NULL,
			email_verif timestamptz NULL,
			password varchar(255) NOT NULL,
			is_verified boolean NOT NULL DEFAULT false,
			bio text NULL,
			profile_image text NULL,
			subscription_type public.subscription_type NOT NULL DEFAULT 'free'::subscription_type,
			created_at timestamptz NOT NULL DEFAULT NOW(),
			updated_at timestamptz NOT NULL DEFAULT NOW(),
			CONSTRAINT uni_users_email UNIQUE (email),
			CONSTRAINT users_pkey PRIMARY KEY (id)
		);

		-- CREATE TABLE verification_token
		CREATE TABLE IF NOT EXISTS public.verification_token (
			id uuid NOT NULL DEFAULT uuid_generate_v4(),
			email varchar(255) NOT NULL,
			token varchar(255) NOT NULL,
			expired timestamptz NOT NULL DEFAULT NOW(),
			CONSTRAINT verification_token_pkey PRIMARY KEY (id)
		);

		-- CREATE TABLE swipes
		CREATE TABLE IF NOT EXISTS swipes (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			user_id UUID NOT NULL,
			target_user_id UUID NOT NULL,
			swipe_type swipe_type NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			created_at_date DATE GENERATED ALWAYS AS (DATE(created_at)) STORED,
			CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
			CONSTRAINT fk_target_user FOREIGN KEY (target_user_id) REFERENCES users(id) ON DELETE cascade,
			CONSTRAINT unique_swipe UNIQUE (user_id, target_user_id, created_at_date )
		);

		-- CREATE TABLE daily_swipes
		CREATE TABLE IF NOT EXISTS daily_swipes (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			user_id UUID NOT NULL,
			swipe_date DATE NOT NULL,
			swipe_count INT NOT NULL,
			CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
			CONSTRAINT unique_daily_swipes UNIQUE (user_id, swipe_date)
		);

		-- CREATE TABLE swipe_history
		CREATE TABLE IF NOT EXISTS swipe_history (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			user_id UUID NOT NULL,
			target_user_id UUID NOT NULL,
			swipe_date DATE NOT NULL,
			CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
			CONSTRAINT fk_target_user FOREIGN KEY (target_user_id) REFERENCES users(id) ON DELETE cascade,
			CONSTRAINT unique_swipe_history UNIQUE (user_id, target_user_id, swipe_date)
		);

		CREATE table IF NOT EXISTS matches (
			id UUID PRIMARY key default uuid_generate_v4(),
			user_id UUID NOT NULL,
			matched_user_id UUID NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
			CONSTRAINT fk_matched_user FOREIGN KEY (matched_user_id) REFERENCES users(id) ON DELETE CASCADE,
			CONSTRAINT unique_match UNIQUE (user_id, matched_user_id)
		);
	`).Error
	if err != nil {
		log.Fatalf("Failed to create enum type: %v", err)
	}
	log.Println("Migration completed successfully.")
}
