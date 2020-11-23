# Accessing Db2 with the Go programming language

The sample programs in this repository show how to access and modify data in **Db2** using the **Go** programming language. *Go* was designed at Google to improve programming productivity. You can use the *Db2 Community Edition* to work with the sample programs. It is a full-featured version of Db2 but with automatically enforced sizing limitations. It can deploy up to 4 cores, 16 GB memory and 100 GB storage for the database.

Read through the following tutorial to understand how the sample programs select data from *Db2* and how they insert, update and delete data. To understand the program examples, basic knowledge of *Go* language or *C* programming language is recommended.

If you just want to get hands-on programming experience with *Go* language and *Db2* you can also take the tutorial on a virtual machine in the cloud. This has all required software pre-installed and can be used for up to one week: https://www.ibm.com/cloud/garage/dte/tutorial/go-db2-use-go-language-develop-db2-database-applications

## Overview of the sample programs

Here is an overview of the sample programs. More details are provided in the tutorial sections further down:
1. ```hello_world.go``` Prints the *Hello world* message.
2. ```connect.go``` Connects to the Db2 *sample* database.
3. ```count_rows.go``` Counts the number of records in some of the tables.
4. ```select_rows.go``` Executes a SELECT statement and retrieves the result set.
5. ```prepare_and_select.go``` Executes a SELECT statement multiple times using different parameter values in the WHERE clause in each execution.
6. ```insert_one_row.go``` Executes a simple INSERT statement.
7. ```insert_multiple_rows.go``` Prepares an INSERT statement and then executes that statement multiple times to insert multiple rows into a table.
8. ```delete_rows.go``` Deletes multiple rows in a loop.
9. ```create_table.go``` Executes a CREATE TABLE statement.
10. ```get_column_names.go``` Returns the names of the columns in a table.
11. ```update_row.go``` Updates exactly one row in a table. 
12. ```update_multiple_rows_with_autocommit.go``` Updates multiple rows in a loop. Each update is immediately commited.
13. ```update_multiple_rows_in_one_unit_of_work.go``` Updates multiple rows in one unit of work and uses the *Begin()* and *Commit()* functions.

## Download and Install the required software
Here are the prerequisites to run the sample programs on your local machine:
- Install a *Go language* binary release suitable for your operating system: https://golang.org/
- Install Db2. The Db2 Community Edition is available as standard download or as a docker container which can be installed with a single command as described here: https://medium.com/@ajstorm/installing-db2-on-your-coffee-break-5be1d811b052. You find more infos about the different deployment options of *Db2 Community Edition* in the following article: https://www.ibm.com/cloud/blog/announcements/ibm-db2-developer-community-edition
- Install the *go_ibm_db* cli driver (also check the API documentation): https://github.com/ibmdb/go_ibm_db
- Download and extract this Github repository. If you have the ```git``` command line tool installed you can clone the repository to your local machine like this: ```git clone https://github.com/andreas-christian/Go-Db2```

## Create the Db2 sample database
After you have installed the required software you also need to create the Db2 *sample* database. On Linux for example, you can run the following commands in a shell:
``` 
su - db2inst1
db2sample
``` 
The sample database contains a number of sample tables that include some data you can work with. The tables are located in schema *DB2INST1*. To see all the tables that are included in the sample database you can take the following steps (Linux):
``` 
su - db2inst1
db2start
db2 connect to sample
db2 "select substr(tabname,1,20) from syscat.tables where tabschema='DB2INST1'"
``` 

Details around the ```db2sample``` command can be found here:
https://www.ibm.com/support/knowledgecenter/SSEPGG_11.5.0/com.ibm.db2.luw.apdv.samptop.doc/doc/t0006757.html

## How to execute the sample programs

The following example shows how to execute the sample programs. We assume that you have downloaded and extracted the git repository in the home directory of your user. To run program *hello_world.go* on Linux for example, you take the following steps:
```
cd ~/Go-Db2/programs
go run hello_world.go
```

# Go Db2 Tutorial

In the following sections, we explain each of the sample programs in more detail. The *Go* code of each sample program is listed at the end of each section.

## connect.go

```connect.go``` is a simple *Go* program that connects to the *SAMPLE* database. It imports the following packages which are required to deploy the Db2 driver API:
```
import _ "github.com/ibmdb/go_ibm_db"  
import "database/sql"  
```
**Note:** The underscore before the package *github.com/ibmdb/go_ibm_db* is required. It ensures that the init function of the package is executed and package-level variables are created.

The function *sql.Open()* is executed to setup a database connection. It requires the driver name *go_ibm_db* and the connection string *con* as input parameters. The connection string specifies *hostname, port number, database name, user name,* and *password*. If *sql.Open()* was executed successfully, the database handle *db* is initialized. Otherwise it will be set to *nil*. Before the program terminates it calls function *db.Close()*. It closes the database connection and cleans up the database handle.

Since the database handle *db* is always required to execute a Db2 API function, we define the connection related variables *db, err,* and *con* outside of function *main()*. This makes sure we can access the database handle in all functions that are defined in package *main*.

We put the *defer* keyword in front of the call of function *db.Close()*. This makes sure that the function is automatically executed as soon as a return statement is executed anywhere in function main().

Execute *connect.go* from the shell as described in the previous section.

```
// connect.go

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

        fmt.Println("Success!")
}
```

# count_rows.go

In the next example, we use function *db.QueryRow()* to run a query that is expected to return at most one row. If there are multiple rows in the query result set, the function will only access the first row and discard the rest. We use the following type of *SELECT* statement in this example:
```
select count(*) statement from <tabname>
```
You notice that the SQL statement text is created by appending the table name in variable *tabname* to the rest of the SQL text. That means the SQL statement is created at runtime (not at compile time). The statement always returns exactly one row which contains the number of records that are stored in the table.

The *Scan()* function copies the columns from the current row into the values pointed. Since we expect a single integer value in the query result set, we define variable *count* of type *int32* and pass a pointer to that variable into function *scan()*.

```
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
```
 # select_rows.go

Next, you learn how to run a *SELECT* statement that returns more than one row. In this example, we use the following select statement:
```
select firstnme, lastname, job from employee where job='MANAGER'
```
Function *db.Query()* prepares and executes the *SELECT* statement. It is also possible to separate the preparation and execution of a SQL statement. This can help to reduce the overhead for statement preparation and is shown in a later example.

Function *rows.Next()* iterates over the result set and prepares the next result row for reading with the *Scan()* api. Since the *SELECT* statement returns three values in each row of the result set, we define variables *a,b,* and *c* and pass their address to function *Scan()*. The function copies the columns from the current row into variables *a,b,* and *c*.

```
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
        if err != nil {
                return
        }
        // make sure that the "rows" handle is released when main returns
        defer rows.Close()

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
```
# prepare_and_select.go

In some cases you may want to execute a query several times but with different filter conditions. In this case, you can prepare the statement once and then execute it multiple times and use a different filter condition in each execution. This allows you to reuse the prepared statement and reduces the overhead of statement preparation. To do this you need to use a parameter marker (?) as shown in the example below:
```
 select firstnme, lastname, job, workdept from employee where workdept = ?
 ```
The statement is prepared by executing function *db.Prepare()*. The function returns a handle st to the prepared statement:
```
st, err := db.Prepare("select firstnme, lastname, job, workdept from employee where workdept = ?")
```
Function *st.Query()* prepares and executes the SQL statement. We use statement handle st to reference the prepared statement. We also have to pass the appropriate number of parameters to the function. Since we prepared the statement with one parameter marker, we pass one parameter dept to the function:
```
rows,err := st.Query(dept)
```
Here is the corresponding sample program:
```
// prepare_and_select.go

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
```

# insert_one_row.go

Executes a simple INSERT statement.
```
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
```
# insert_multiple_rows.go

Prepares an INSERT statement and then executes that statement multiple times to insert multiple rows into a table.
```
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
```

# delete_rows.go

Deletes multiple rows in a loop.
```
// delete_rows.go

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

        st, err := db.Prepare("delete from lineitem where name=?")
        if err !=nil {
                fmt.Println("Error in Prepare: ")
                fmt.Println(err)
                return
        }

        lineitems:= []string{"Shirt","Coffee"}
        for _,item := range lineitems{
                _,err = st.Exec(item)
                if err != nil{
                        fmt.Println("Error:")
                        fmt.Println(err)
                        return
                }
                fmt.Println("Item deleted:")
                fmt.Println(item)
        }
}
```

# create_table.go

Executes a CREATE TABLE statement.
```
// create_table.go

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

        _,err:=db.Exec("create table LINEITEM(ID int,NAME varchar(20),QTY int)")
        if err != nil{
                fmt.Println("Error:")
                fmt.Println(err)
                return
        }
        fmt.Println("TABLE CREATED")
}
```

# get_column_names.go

Returns the names of the columns in a table.
```
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
        if err != nil {
                fmt.Printf("db.Query(): error!")
                return
        }
        // make sure that the "rows" handle is released when main returns
        defer rows.Close()

        cols, err := rows.Columns()
        fmt.Println("Number of columns: ",len(cols))
        // print the whole array at once
        fmt.Printf("%v\n",cols)
        // print each column name on a separate line
        for _,name := range cols {
                fmt.Printf("%s\n",name)
        }
}
```

# update_row.go

Updates exactly one row in a table. 

```
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
```

# update_multiple_rows_with_autocommit.go

Updates multiple rows in a loop. Each update is immediately commited.
```
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
        if err != nil {
                return
        }
        defer rows.Close()

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
```

# update_multiple_rows_in_one_unit_of_work.go

Updates multiple rows in one unit of work and uses the *Begin()* and *Commit()* functions.
```
// update_multiple_rows_in_one_unit_of_work.go

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
        if err != nil {
                return
        }
        defer rows.Close()

        var currentTime int64
        currentTime = time.Now().Unix()
        newqty := currentTime & 0x00000000000000FF
        fmt.Println("Quantity:", newqty)

        var id int32
        var name string
        var qty int32

        // Begin Unit of Work (UoW)
        uow, err := db.Begin()
        if err != nil {
                fmt.Println(err)
                return
        }
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
                _,err = uow.Exec("update lineitem set qty=? where id=?",newqty,id)
                if err != nil{
                        fmt.Println(err)
                        return
                }

        }
        // End Unit of Work
        err = uow.Commit()
        if err != nil {
                fmt.Println(err)
                return
        }
}
```

