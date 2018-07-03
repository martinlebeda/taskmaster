package service

import (
    "time"
)

func TmrAdd(duration, title string) {

    parseDuration, err := time.ParseDuration(duration)
    CheckErr(err)
    goal := time.Now().Add(parseDuration)

    db := OpenDB()
    stmt, err := db.Prepare("INSERT INTO timer(note, goal) values(?,?)")
    CheckErr(err)
    _, err = stmt.Exec(title, goal)
    CheckErr(err)
}
