# Assignment


##  How to run

### Below command to download all the modules/packages used in the project 

``` go mod download ```   



### To run the rest api run this command
``` make dev ```          

 OR
 
```   go run main.go   ```   

* please add a .env file and fill this 
``` 
POSTGRESQL_ADDON_DB=""
POSTGRESQL_ADDON_HOST=""
POSTGRESQL_ADDON_PASSWORD=""
POSTGRESQL_ADDON_PORT=""
POSTGRESQL_ADDON_USER=""
  ```
## Folder structer
* Controller Folder Contains all the endpoint functions

* Database folder contains the database(PostgreSQL) connection configurations

* Model folder contains the Database Schemas

* Router folder, all the routes are defined in this folder

* Main.go- It is the main execution/entry point of the project

# RestAPI end points

*  ### (GET)localhost:4000/api/getdata -all the data available in the table will be displayed by hitting this route 

* ### (GET)lcoalhost:4000/api/:id/getadata - This route will show the available data against to the user id provided

* ### (POST)localhost:4000/api/createdata - This route will Insert the response given by the end user into the database table 

* ### (PUT)localhost:4000/api/:id/updatedata - This route will update data corresponds to the user id

* ### (DELETE)localhost:4000/api/:id/deletedata - This route will delete the data of a matched user id