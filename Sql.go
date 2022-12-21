package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type SQL struct {
    db *sql.DB
}

var sql_obj *SQL = nil

func sql_open() Result[*sql.DB] {
    db, err := sql.Open("sqlite3", "info.db")

    if err != nil {
        return Error[*sql.DB] { err.Error() }
    }

    return Success[*sql.DB] { db }
}

func sql_getversion() Result[string] {
    if sql_obj == nil {
        return Error[string] { "No database object" }
    }

    var version string
    err := sql_obj.db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)
    if err != nil {
        return Error[string] { err.Error() }
    }

    return Success[string] { version }
}

func sql_ensuretable() Result[*sql.Row] {
    if sql_obj == nil {
        return Error[*sql.Row] { "No database object" }
    }

    queryRes := sql_obj.db.QueryRow("CREATE TABLE IF NOT EXISTS `users` ( id int, firstname varchar(60), lastname varchar(60), occupancy varchar(120) );")
    if queryRes.Err() != nil {
        return Error[*sql.Row] { queryRes.Err().Error() }
    }

    return Success[*sql.Row] { queryRes }
}

func sql_do_work() {
    res := sql_open()
    if res.HasError() {
        fmt.Printf("%v", res.Error())
        return
    }

    sql_obj = &SQL { *res.Get() }
    defer sql_obj.db.Close()

    versionRes := sql_getversion()
    if versionRes.HasError() {
        fmt.Printf("%v", versionRes.Error())
    } else {
        fmt.Printf("Sql version: %s", *versionRes.Get())
    }

    ensureRes := sql_ensuretable()
    if ensureRes.HasError() {
        fmt.Printf("Could not create a table?(Not really sure what's going on)\n%v", ensureRes.Error())
        return
    }


}
