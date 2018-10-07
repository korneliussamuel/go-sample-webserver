PREREQUISITE:
1. Install go
2. Install postgresql
3. Inside postgres, make appropriate role and database
4. Inside selected database, create table `persons` which contain: ID (int), Name (text), Age (int)

BUILD and RUN:
1. change database config inside db/conf.sh
2. source db/conf.sh
3. go build .
4. ./go-sample-webserver

