// update_multiple_rows_with_autocommit.go

package main
import (
    _ "github.com/ibmdb/go_ibm_db"
    "database/sql"
    "fmt"
    "time"
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

        rows,err := db.Query("select * from lineitem")
        if rows != nil {defer rows.Close()}
        if err != nil {
                fmt.Println(err)
                return
        }

        var currentTime int64
        currentTime = time.Now().Unix()
        newqty := currentTime & 0x00000000000000FF
        fmt.Println("Quantity:", newqty)

        var id int32
        var name string
        var qty int32

        for rows.Next() {
                err = rows.Scan(&id,&name,&qty)
                if err != nil{
                        fmt.Println(err)
                        return
                }
                fmt.Printf("Fetched one row:\n")
                fmt.Printf("%-5d %-10s %-5d\n",id,name,qty)
                time.Sleep(1*time.Second)

                fmt.Printf("Updating row with new quantity value: %d\n",newqty)
                _,err = db.Exec("update lineitem set qty=? where id=?",newqty,id)
                if err != nil{
                        fmt.Println(err)
                        return
                }

        }
}
