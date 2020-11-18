// insert_one_row.go

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

func main() {
        if connect() != nil { return } else { defer db.Close() }

        _,err:=db.Exec("insert into lineitem values (99,'Flowers',5)")
        if err != nil{
                fmt.Println("Error:")
                fmt.Println(err)
                return
        }
        fmt.Println("Row inserted.")
}
