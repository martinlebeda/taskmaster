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
	"github.com/GuiaBolso/darwin"
	"github.com/martinlebeda/taskmaster/termout"
	"github.com/martinlebeda/taskmaster/tools"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"log"
)

var (
	// definition of DB migrations and schemas
	migrations = []darwin.Migration{
		{
			Version:     1,
			Description: "Creating table timer",
			Script:      `CREATE TABLE timer (note VARCHAR, goal DATETIME);`,
		},
		{
			Version:     2,
			Description: "view for timer",
			Script:      `create view timer_distance as SELECT rowid, strftime('%s', goal, 'localtime') - strftime('%s', 'now', 'localtime') as distance, goal, note from timer order by distance;`,
		},
		{
			Version:     3,
			Description: "add tag to timer",
			Script: `ALTER TABLE timer ADD COLUMN tag VARCHAR;
DROP VIEW IF EXISTS timer_distance;
CREATE VIEW timer_distance AS SELECT rowid, strftime('%s', goal, 'localtime') - strftime('%s', 'now', 'localtime') as distance, goal, tag, note FROM timer ORDER BY distance;
`,
		},
		{
			Version:     4,
			Description: "Creating table work",
			Script:      `CREATE TABLE work (category VARCHAR, code VARCHAR, desc VARCHAR, start DATETIME, stop DATETIME);`,
		},
		{
			Version:     5,
			Description: "Fix view timer_distance",
			Script: `DROP VIEW IF EXISTS timer_distance;
            CREATE VIEW timer_distance AS SELECT rowid, strftime('%s', goal, 'localtime') - strftime('%s', 'now', 'localtime') as distance, goal, tag, note FROM timer ORDER BY distance;`,
		},
		{
			Version:     6,
			Description: "Create table task",
			Script: `CREATE TABLE task (
id INTEGER PRIMARY KEY AUTOINCREMENT, 
prio VARCHAR,
status VARCHAR, 
code VARCHAR,
category VARCHAR,
desc VARCHAR, 
date_in DATETIME, 
date_done DATETIME, 
url VARCHAR, 
note VARCHAR, 
script VARCHAR
);`,
		},
		{
			Version:     7,
			Description: "Fix view timer_distance",
			Script: `DROP VIEW IF EXISTS timer_distance;
            CREATE VIEW timer_distance AS SELECT rowid, strftime('%s', goal, 'localtime') - strftime('%s', 'now', 'localtime') as distance, goal, tag, note FROM timer ORDER BY distance;`,
		},
		{
			Version:     8,
			Description: "Add estimate to task",
			Script:      `ALTER TABLE task ADD COLUMN estimate VARCHAR;`,
		},
		{
			Version:     9,
			Description: "big simplify task and work",
			Script: `DROP TABLE IF EXISTS workNew;
				CREATE TABLE workNew (id INTEGER PRIMARY KEY AUTOINCREMENT, desc VARCHAR, start DATETIME, stop DATETIME);
				insert into workNew (id, desc, start, stop) select rowid, CASE WHEN code is not NULL and code <> '' THEN code || " - " ELSE "" END || desc || 
						CASE WHEN category is not NULL and category <> '' THEN " +" || category ELSE "" END, start, stop from work;
				drop table work;
				ALTER TABLE workNew RENAME TO work;

                DROP TABLE IF EXISTS taskNew; 
				CREATE TABLE taskNew (id INTEGER PRIMARY KEY AUTOINCREMENT, status VARCHAR, desc VARCHAR, date_in DATETIME, date_done DATETIME, estimate VARCHAR);
                insert into taskNew (id, status, desc, date_in, date_done, estimate) 
                select id, status,
                        CASE WHEN prio is not NULL and prio <> '' THEN "(" || prio || ") " ELSE "" END ||
                        CASE WHEN code is not NULL and code <> '' THEN code || " - " ELSE "" END || desc || 
                        CASE WHEN category is not NULL and category <> '' THEN " +" || category ELSE "" END ||
                        CASE WHEN url is not NULL and url <> '' THEN " " || url ELSE "" END,
                        date_in, date_done, estimate from task;
				drop table task;
				ALTER TABLE taskNew RENAME TO task;`,
		},
		{
			Version:     10,
			Description: "drop estimate from task",
			Script: `DROP TABLE IF EXISTS taskNew; 
				CREATE TABLE taskNew (id INTEGER PRIMARY KEY AUTOINCREMENT, status VARCHAR, desc VARCHAR, date_in DATETIME, date_done DATETIME);
                insert into taskNew (id, status, desc, date_in, date_done) 
                select id, status, desc, date_in, date_done from task;
				drop table task;
				ALTER TABLE taskNew RENAME TO task;`,
		},
	}
)

// Open database file
func OpenDB() *sql.DB {
	dbFileName := viper.GetString("dbfile")
	db, err := sql.Open("sqlite3", dbFileName)
	tools.CheckErr(err)

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
	tools.CheckErr(err)
	for rows.Next() {
		err := rows.Scan(&result)
		tools.CheckErr(err)
	}

	rows.Close()
	return result
}
