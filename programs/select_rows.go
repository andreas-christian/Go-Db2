// select_rows.go

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

        rows,err := db.Query("select firstnme, lastname, job from employee where job='MANAGER'")
        if rows != nil {defer rows.Close()}
        if err != nil {
                return
        }

        // iterate over all rows in the query result
        var a,b,c string
        for rows.Next() {
                err = rows.Scan(&a,&b,&c)
                if err != nil{
                        fmt.Println(err)
                        return
                }
                fmt.Printf("%-10s %-10s %-10s\n",a,b,c)
        }
}
