# Accessing a relational database- mysql

## Prerequisites
- An installation of the MySQL relational database management system (DBMS).
## Set up a database
1. At the command line, log into your DBMS, as in the following example for MySQL.
    ```
    $ mysql -u root -p
    Enter password:
    
    mysql>
    ```
2. At the mysql command prompt, create a database.
   ```
   mysql> source > db/file/path/create-tables.sql

   ```
3. At your DBMS command prompt, use a SELECT statement to verify you’ve successfully created the table with data.
   ```
   mysql> select * from album;
   +----+---------------+----------------+-------+
   | id | title         | artist         | price |
   +----+---------------+----------------+-------+
   |  1 | Blue Train    | John Coltrane  | 56.99 |
   |  2 | Giant Steps   | John Coltrane  | 63.99 |
   |  3 | Jeru          | Gerry Mulligan | 17.99 |
   |  4 | Sarah Vaughan | Sarah Vaughan  | 34.98 |
   +----+---------------+----------------+-------+
   4 rows in set (0.00 sec)

   ```
4. From the command prompt, set the DBUSER and DBPASS environment variables for use by the Go program.
   ```
   $ export DBUSER=username
   $ export DBPASS=password
   ```
5. From the command line in the directory containing main.go, run the code by typing go run with a dot argument to mean “run the package in the current directory.
   ```
   $ go run .
   Connected!
   ```
6. All Done.


### TODO: 
- Create a full relational table for api server. 