# Accessing Db2 with the Go programming language

The sample programs in this repository explain how to access and modify data in Db2 using the Go programming language. Go was designed at Google to improve programming productivity. 

Read through the following tutorial to understand how the sample programs select data from Db2 and how they insert, update and delete data. 

If you just want to get hands-on programming experience with Go language and Db2 you can also work through the below tutorial on a virtual machine in the cloud. This has all required software pre-installed and can be used for up to one week: https://www.ibm.com/cloud/garage/dte/tutorial/go-db2-use-go-language-develop-db2-database-applications

Here are the prerequisites to run the sample programs on your local machine:
- Install a go language binary release suitable for your system: https://golang.org/
- Install Db2. The Db2 Community Edition is available as standard download or as a docker container which can be installed with a single command: https://www.ibm.com/cloud/blog/announcements/ibm-db2-developer-community-edition-
- Install the go_ibm_db cli driver: https://github.com/ibmdb/go_ibm_db

After you have setup your environment you also need to create the Db2 sample database. This can be done from the db2 command line:
  su - db2inst1
  db2sample

Details around the sample database can be found here:
https://www.ibm.com/support/knowledgecenter/SSEPGG_11.5.0/com.ibm.db2.luw.apdv.samptop.doc/doc/t0006757.html
