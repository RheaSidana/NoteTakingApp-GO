<h1>
    Note Taking Application 
</h1>

<h3>
    Steps to run the application
</h3>

<h5>
    <br/>
    1. Clone the repo
    <br/><br/>
    2. run the command: <br/>
    ```
    go mod tidy
    ```
    <br/><br/>
    3. DB operations
    <br/><br/>&emsp;
    <<<< INSTALLING >>>>
    <br/>&emsp;&emsp;
    a. MAC 
    <br/>&emsp;&emsp;&emsp;&emsp;
    1. install postgres: <br/>
    ```
    brew install postgresql
    ```
    <br/>&emsp;&emsp;&emsp;&emsp;
    2. start/stop postgres service: <br/>
    ```
    brew services start/stop postgresql <br/>
    ```
    <br/>&emsp;&emsp;&emsp;&emsp;
    3. <br/>
    ```
    psql postgres
    ```
    <br/>&emsp;&emsp;&emsp;&emsp;
    4. add postgres password: <br/>
    ```
    \password {password};
    ```
    <br/><br/>&emsp;&emsp;
    b. Windows
    <br/>&emsp;&emsp;&emsp;&emsp;
    1. install postgres: 
    <!-- (https://www.postgresql.org/download/windows/) -->
    <a href="https://www.postgresql.org/download/windows/">[Link text Here]</a>
    <br/>&emsp;&emsp;&emsp;&emsp;
    2. port: 5432 (defaut), user: postgres (defaut)
    <br/>&emsp;&emsp;&emsp;&emsp;
    3. add postgres password
    <br/>&emsp;&emsp;&emsp;&emsp;
    4. open psql sql shell
    <br/><br/>&emsp;
    <<<< CREATING >>>>
    <br/>&emsp;&emsp;&emsp;&emsp;
    5. create orm db: <br/>
    ```
    CREATE database job_scheduling;
    ```
    <br/><br/>&emsp;
    <<<< CONNECTING >>>>
    <br/>&emsp;&emsp;&emsp;&emsp;
    6. connect to db: <br/>
    ```
    \c "db"
    ```
    <br/>&emsp;&emsp;&emsp;&emsp;
    7. Edit .env file with postgres details
    <br/><br/><br/>
    4. Migrate Tables <br/>
    ```
    go run .\migrations\migrate.go
    ```
    <br/><br/>&emsp;
    View DB Table Schemas: <br/>
    ```
    \d "tablename"
    ```
    <br/><br/>
    5. Seed Data to the Table <br/>
    ```
    go run .\dataSeeding\dataSeed.go
    ```
    <br/><br/>
    6. Run the application <br/>
    ```
    go run .
    ```
    <br/><br/>
    7. Call APIs
</h5>
