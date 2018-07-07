package service

import (
    "time"
    "github.com/martinlebeda/taskmaster/termout"
    . "github.com/martinlebeda/taskmaster/model"
    "strconv"
    "strings"
)

func TmrSet(dateOpt, timeArg, title string) {
    goal, err := time.Parse("2006-01-02 15:04", dateOpt+" "+timeArg)
    CheckErr(err)

    insertNewTimer(title, goal)
}

func TmrAdd(duration, title string) {
    parseDuration, err := time.ParseDuration(duration)
    CheckErr(err)

    goal := time.Now().Add(parseDuration)

    insertNewTimer(title, goal)
}

func insertNewTimer(title string, goal time.Time) {
    termout.Verbose("New goal for ", title, " set to ", goal.String())
    db := OpenDB()
    stmt, err := db.Prepare("INSERT INTO timer(note, goal) values(?,?)")
    CheckErr(err)
    _, err = stmt.Exec(title, goal)
    CheckErr(err)
    termout.Verbose("New timer inserted")
}

func TmrDel(args []string)  {
    db := OpenDB()
    stmt, err := db.Prepare("delete from timer where rowid in (" + strings.Join(args,",") + ")")
    CheckErr(err)
    _, err = stmt.Exec()
    CheckErr(err)
    termout.Verbose("Timer deleted: ", strings.Join(args,","))
}

func TmrGetDistance(pastOpt, nextOpt bool) []TimerDistance {
    db := OpenDB()
    sql := "select rowid, distance, goal, note from timer_distance "

    if pastOpt {
        sql += " where distance < 0 "
    }
    if nextOpt {
        sql += " where distance > 0 "
    }

    sql +=    " order by distance "

    if nextOpt {
        sql += " limit 1 "
    }

    rows, err := db.Query(sql)
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
    timerDistances := TmrGetDistance(false, false)
    termout.TmrListDistance(timerDistances, false)
}