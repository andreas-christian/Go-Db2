// connect.go

package main

import (
    _ "github.com/ibmdb/go_ibm_db"
    "database/sql"
    "fmt"
)

func main(){
        con := "HOSTNAME=localhost;PORT=50000;DATABASE=SAMPLE;UID=DB2INST1;PWD=db2inst1"
        db, err:=sql.Open("go_ibm_db", con)
        if err != nil {
                fmt.Println(err)
        }
        fmt.Println("Success!")
        db.Close()
}
