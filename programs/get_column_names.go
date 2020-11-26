// get_column_names.go

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

        rows,err := db.Query("select * from employee fetch first 1 row only")
        if rows != nil {defer rows.Close()}
        if err != nil {
                fmt.Printf("db.Query(): error!")
                return
        }

        cols, err := rows.Columns()
        fmt.Println("Number of columns: ",len(cols))
        // print the whole array at once
        fmt.Printf("%v\n",cols)
        // print each column name on a separate line
        for _,name := range cols {
                fmt.Printf("%s\n",name)
        }
}
