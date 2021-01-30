package repository

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

type ProfileBrief struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Surname  string `db:"surname"`
	IsFriend bool   `db:"is_friend"`
}

type Profile struct {
	ProfileBrief
	Age       int    `db:"age"`
	Gender    string `db:"gender"`
	Interests string `db:"interests"`
	City      string `db:"city"`
}

type User struct {
	ID           string `db:"id"`
	PasswordHash string `db:"password_hash"`
	FilledForm   bool   `db:"filled_form"`
}

func (u userRepo) LoginUser(username string) (User, error) {
	var user User
	err := u.db.
		QueryRowx(`
		SELECT CONVERT(id, CHAR) AS id,
		       password_hash,
		       EXISTS(SELECT id from profile where profile.user_id = id) as filled_form
		FROM user WHERE username = ?
`, username).
		StructScan(&user)
	if err != nil && err != sql.ErrNoRows {
		return user, fmt.Errorf("failed to query user: %w", err)
	}

	return user, nil
}

func (u userRepo) CreateUser(username, hash string) (string, error) {
	res, err := u.db.Exec(
		`INSERT INTO user (username, password_hash) VALUES (?, ?);`, username, hash)
	if err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}

	idLast, err := res.LastInsertId()
	if err != nil {
		return "", fmt.Errorf("create user get last id:%w", err)
	}

	return fmt.Sprint(idLast), nil
}

func (u userRepo) GetProfile(userID int) (Profile, error) {
	var profile Profile
	err := u.db.Get(&profile, `
		SELECT 
		       profile.user_id AS id,
		       name,
		       surname,
		       age,
		       gender,
		       interests,
		       city
		FROM profile 
    	LEFT JOIN friends f ON profile.user_id = f.friend_id AND f.user_id = ?
        WHERE profile.user_id = ?
    	`, userID, userID)
	if err != nil {
		return profile, fmt.Errorf("query get profile: %w", err)
	}

	return profile, nil
}

func (u userRepo) List() ([]Profile, error) {
	var profiles []Profile
	err := u.db.Select(&profiles, "SELECT profile.name, surname, age, gender, interests, city FROM profile")
	if err != nil {
		return nil, fmt.Errorf("query list users: %w", err)
	}

	return profiles, nil
}

func (u userRepo) ListBrief(userID int) ([]ProfileBrief, error) {
	var profiles []ProfileBrief
	err := u.db.Select(&profiles, `
		SELECT profile.user_id AS id, name, surname, IFNULL(NOT NOT f.friend_id, FALSE) AS is_friend
		FROM profile 
    	LEFT JOIN friends f ON profile.user_id = f.friend_id AND f.user_id = ?
	    WHERE profile.user_id != ?
    	`, userID, userID)
	if err != nil {
		return nil, fmt.Errorf("query list profiles: %w", err)
	}

	return profiles, nil
}

func (u userRepo) CreateUpdateProfile(userID int, p Profile) error {
	_, err := u.db.Exec(`
		INSERT INTO profile (user_id, name, surname, age, gender, interests, city)
		VALUES (?, ?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE name      = ?,
		                        surname   = ?,
		                        age       = ?,
		                        gender    = ?,
		                        interests = ?,
		                        city      = ?;
    `, userID, p.Name, p.Surname, p.Age, p.Gender, p.Interests, p.City,
		p.Name, p.Surname, p.Age, p.Gender, p.Interests, p.City,
	)
	if err != nil {
		return fmt.Errorf("insert into profile: %w", err)
	}

	return nil
}

func (u userRepo) AddFriend(userID, friendID int) error {
	_, err := u.db.Exec(`INSERT INTO friends (user_id, friend_id) VALUES (?,?)`, userID, friendID)
	if err != nil {
		return fmt.Errorf("failed to exec add friend query: %w", err)
	}

	return nil
}

func (u userRepo) RemoveFriend(userID, friendID int) error {
	_, err := u.db.Exec(`DELETE FROM friends WHERE user_id = ? AND friend_id = ?`, userID, friendID)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("failed to exec add friend query: %w", err)
	}

	return nil
}

func (u userRepo) ListFriends(userID int) ([]ProfileBrief, error) {
	var friends []ProfileBrief
	err := u.db.Select(&friends, `
		SELECT  p.user_id as id, p.name, p.surname, TRUE as is_friend
		FROM friends JOIN profile p ON friends.friend_id = p.user_id
		WHERE friends.user_id = ?;
	`, userID)
	if err != nil {
		return nil, fmt.Errorf("query friends: %w", err)
	}

	return friends, nil
}
