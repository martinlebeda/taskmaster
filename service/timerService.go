package service

import (
    "time"
    "github.com/martinlebeda/taskmaster/termout"
    . "github.com/martinlebeda/taskmaster/model"
    "strconv"
)

func TmrAdd(duration, title string) {

    parseDuration, err := time.ParseDuration(duration)
    CheckErr(err)

    goal := time.Now().Add(parseDuration)

    termout.Verbose("Nový cíl pro ", title, " nastaven na ", goal.String())

    db := OpenDB()
    stmt, err := db.Prepare("INSERT INTO timer(note, goal) values(?,?)")
    CheckErr(err)
    _, err = stmt.Exec(title, goal)
    CheckErr(err)
    termout.Verbose("New timer inserted")
}

func TmrGetDistance() []TimerDistance {
    db := OpenDB()
    rows, err := db.Query("select rowid, distance, goal, note from timer_distance order by distance")
    CheckErr(err)

    var result []TimerDistance
    for rows.Next() {
        var timerDistance TimerDistance
        rows.Scan(&timerDistance.Rowid, &timerDistance.Distance, &timerDistance.Goal, &timerDistance.Note)
        result = append(result, timerDistance)
    }

    return result
}

func TmrClean(deleteAll bool) {
    db := OpenDB()

    sql := "delete from timer"
    if !deleteAll {
        sql += " where rowid in (select rowid from timer_distance where distance < 0)"
    }

    result, err := db.Exec(sql)
    CheckErr(err)
    count, err := result.RowsAffected()
    CheckErr(err)
    termout.Verbose("Count of deleted timers: ", strconv.FormatInt(count, 10))
}

func TmrListAfterChange() {
    termout.EmptyLineOut()
    timerDistances := TmrGetDistance()
    termout.TmrListDistance(timerDistances)
}