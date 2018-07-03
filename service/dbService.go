// Copyright © 2018 Martin Lebeda <martin.lebeda@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package service

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "github.com/spf13/viper"
    "github.com/GuiaBolso/darwin"
    "log"
    "github.com/martinlebeda/taskmaster/termout"
    "github.com/spf13/cast"
)

var (
    // definition of DB migrations and schemas
    migrations = []darwin.Migration{
        {
            Version:     1,
            Description: "Creating table timer",
            Script:      `CREATE TABLE timer (note VARCHAR, goal DATETIME);`,
        },
        //{
        //	Version:     2,
        //	Description: "Adding column body",
        //	Script:      "ALTER TABLE posts ADD body TEXT AFTER title;",
        //},
    }
)

// Open database file
func OpenDB() *sql.DB {
    dbFileName := viper.GetString("dbfile")
    db, err := sql.Open("sqlite3", dbFileName)
    CheckErr(err)

    return db
}

// make DB schema actual
func DbUpgrade() {
    database := OpenDB()

    driver := darwin.NewGenericDriver(database, darwin.SqliteDialect{})
    d := darwin.New(driver, migrations, nil)
    err := d.Migrate()
    if err != nil {
        log.Println(err)
    }

    termout.Verbose("Database upgraded to version " + cast.ToString(getDbVersion(database)) + ".")

    database.Close()
}
func getDbVersion(db *sql.DB) float32 {
    var result float32
    rows, err := db.Query("select max(version) from darwin_migrations")
    CheckErr(err)
    for rows.Next() {
        err := rows.Scan(&result)
        CheckErr(err)
    }

    rows.Close()
    return result
}
