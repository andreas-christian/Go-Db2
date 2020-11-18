# Accessing Db2 with the Go programming language

The sample programs in this repository explain how to access and modify data in **Db2** using the **Go** programming language. *Go* was designed at Google to improve programming productivity. 

Read through the following tutorial to understand how the sample programs select data from Db2 and how they insert, update and delete data. 

If you just want to get hands-on programming experience with Go language and Db2 you can also take the tutorial on a virtual machine in the cloud. This has all required software pre-installed and can be used for up to one week: https://www.ibm.com/cloud/garage/dte/tutorial/go-db2-use-go-language-develop-db2-database-applications

## Download and Install the required software
Here are the prerequisites to run the sample programs on your local machine:
- Install a *Go language* binary release suitable for your operating system: https://golang.org/
- Install Db2. The Db2 Community Edition is available as standard download or as a docker container which can be installed with a single command as described here: https://medium.com/@ajstorm/installing-db2-on-your-coffee-break-5be1d811b052. You find more infos about the different deployment options of *Db2 Community Edition* in the following article: https://www.ibm.com/cloud/blog/announcements/ibm-db2-developer-community-edition
- Install the *go_ibm_db* cli driver: https://github.com/ibmdb/go_ibm_db
- Download and extract this Github repository. If you have the ```git``` command line tool installed you can clone the repository to your local machine like this: ```git clone https://github.com/andreas-christian/Go-Db2```

## Create the Db2 sample database
After you have installed the required software you also need to create the Db2 *sample* database. On Linux, this can be done from a shell:
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

Perform the following steps to execute one of the sample programs:
- Open a shell window
- Change to the directory where the sample programs are located. On Linux for example, if you extracted the *Go-Db2* Git repository in the home directory of your user, execute the following command: ```cd ~/Go-Db2/programs```
- Then execute a program like this: ```go run hello_world.go```

## Overview of the sample programs

1. ```hello_world.go``` Prints the *Hello world* message.
2. ```connect.go``` Connects to the Db2 *sample* database.
3. ```count_rows.go``` Counts the number of records in some of the tables.
4. ```select_rows.go``` Executes a SELECT statement and retrieves the result set.
5. ```prepare_and_select.go``` Executes a SELECT statement multiple times using different parameter values in the WHERE clause in each execution.
6. ```insert_one_row.go``` Executes a simple INSERT statement.
7. ```insert_multiple_rows.go``` Prepares an INSERT statement and then executes that statement multiple times to insert multiple rows into a table.
8. delete_rows.go
9. create_table.go
10. get_column_names.go
11. ```update_row.go``` Updates exactly one row in a table. update_row.go
12. ```update_multiple_rows_with_autocommit.go``` Updates rows in a loop. Each update is immediately commited.
13. ```update_multiple_rows_in_one_unit_of_work.go``` 
