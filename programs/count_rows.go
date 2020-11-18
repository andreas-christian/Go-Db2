// count_rows.go

package main
import (
    _ "github.com/ibmdb/go_ibm_db"
    "database/sql"
    "fmt"
)
var err error
var db *sql.DB
var con = "HOSTNAME=localhost;PORT=50000;DATABASE=SAMPLE;UID=DB2INST1;PWD=db2inst1"

func connect() error {
        db, err = sql.Open("go_ibm_db", con)
        if err != nil {
                fmt.Println(err)
                return err
        }
        return nil
}
func count(tabname string) error { 
        var count int32
        err = db.QueryRow("SELECT count(*) FROM "+tabname).Scan(&count)
        if err != nil {
                fmt.Println(err)
                return err
        }       
        fmt.Printf("Table \"%s\" contains %d rows.\n",tabname,count)
        return nil
}

func main() {
        if connect() != nil { return } else { defer db.Close() }

        count("ACT")
        count("DEPARTMENT")
        count("EMPLOYEE")
        count("ORG")
}
