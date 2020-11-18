// update_row.go

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

        // prepare the statement with parameter markers
        st, err := db.Prepare("update lineitem set qty=? where id=?")
        if err !=nil {
                fmt.Println("Error in Prepare: ")
                fmt.Println(err)
                return
        }

        id := 0
        qty := 3
        _, err = st.Exec(id, qty)
        if err != nil{
                fmt.Println("Error:")
                fmt.Println(err)
                return
        }
        fmt.Println("Row updated.")
}
