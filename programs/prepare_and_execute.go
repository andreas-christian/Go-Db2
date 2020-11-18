// prepare_and_execute.go

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

        // prepare the statement once with a parameter marker
        st, err := db.Prepare("select firstnme, lastname, job, workdept from employee where workdept = ?" )
        if err !=nil {
                fmt.Println("Error in Prepare: ")
                fmt.Println(err)
                return
        }
        // execute the statement multiple times and use a different
        // work department in the where clause for each query execution
        departments := []string{"A00","B01","C01","D11","D21","E11","E21"}
        for _,dept := range departments{
                fmt.Printf("\nSelect records for department '%s'\n", dept)
                rows,err := st.Query(dept)
                if err != nil {
                        fmt.Println("Error in Query: ")
                        fmt.Println(err)
                        return
                }

                // iterate over all rows in the query result
                for rows.Next() {
                        var a,b,c,d string
                        err = rows.Scan(&a,&b,&c,&d)
                        if err != nil{
                                fmt.Println(err)
                                return
                        }
                fmt.Printf("%-10s %-10s %-10s %-10s\n",a,b,c,d)
                }
                rows.Close()
        }
}
