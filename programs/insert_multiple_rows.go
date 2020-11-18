// insert_multiple_rows.go

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

        // prepare the statement once with parameter markers
        st, err := db.Prepare("insert into lineitem values (?,?,?)" )
        if err !=nil {
                fmt.Println("Error in Prepare: ")
                fmt.Println(err)
                return
        }

        lineitems:= []string{"Shirt","Bicycle","Laptop","Coffee","Burger","Watch"}
        for idx,item := range lineitems{
                _,err = st.Exec(idx,item,5)
                if err != nil{
                        fmt.Println("Error:")
                        fmt.Println(err)
                        return
                }
                fmt.Println("Row inserted.")
        }
}
