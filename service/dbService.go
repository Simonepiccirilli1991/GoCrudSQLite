package service

import (
	"NotifyGoBot/model"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() error {
	var err error
	db, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS notify (id INTEGER PRIMARY KEY AUTOINCREMENT, from_user TEXT, testo TEXT)")
	if err != nil {
		return err
	}

	return nil
}

func CloseDB() {
	db.Close()
}

func CreateNotify(notify *model.Notify) error {
	result, err := db.Exec("INSERT INTO notify (from, testo) VALUES (?, ?)", notify.From, notify.Testo)
	if err != nil {
		return err
	}
	notify.ID, _ = result.LastInsertId()
	return nil
}

func GetNotify() ([]model.Notify, error) {
	rows, err := db.Query("SELECT id, from, testo FROM notify")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifys []model.Notify
	for rows.Next() {
		var u model.Notify
		err := rows.Scan(&u.ID, &u.From, &u.Testo)
		if err != nil {
			return nil, err
		}
		notifys = append(notifys, u)
	}
	return notifys, nil
}

func UpdateNotifyTesto(notify *model.Notify, testo string) error {
	_, err := db.Exec("UPDATE users SET testo = ? WHERE id = ?", notify, notify.ID)
	if err != nil {
		return err
	}
	notify.Testo = testo
	return nil
}
