<h1>Note Taking Application</h1>

<h2>Steps to run the application</h2>
<br />
<h4>1. Clone the repo</h4>
<br /><br />
<h4>2. run the command:</h4>

``` 
go mod tidy 
```

<br /><br />
<h5>3. DB operations</h5>
<br /><br />&emsp;
<h5>
  <<<< INSTALLING >>>> <br />&emsp;&emsp; a. MAC <br />&emsp;&emsp;&emsp;&emsp;
  1. install postgres:
</h5>
<br />

``` 
brew install postgresql 
```

<br />&emsp;&emsp;&emsp;&emsp;
<h5>2. start/stop postgres service:</h5>

``` 
brew services start/stop postgresql 
```

<br />&emsp;&emsp;&emsp;&emsp;
<h5>3.</h5>

``` 
psql postgres 
```

<br />&emsp;&emsp;&emsp;&emsp;
<h5>4. add postgres password:</h5>

``` 
\password {password}; 
```

<br /><br />&emsp;&emsp;
<h5>
  b. Windows <br />&emsp;&emsp;&emsp;&emsp; 1. install postgres:
  <!-- (https://www.postgresql.org/download/windows/) -->
  <a href="https://www.postgresql.org/download/windows/">[Link text Here]</a>
  <br />&emsp;&emsp;&emsp;&emsp; 2. port: 5432 (defaut), user: postgres (defaut)
  <br />&emsp;&emsp;&emsp;&emsp; 3. add postgres password
  <br />&emsp;&emsp;&emsp;&emsp; 4. open psql sql shell <br /><br />&emsp; <<<<
  CREATING >>>> <br />&emsp;&emsp;&emsp;&emsp; 5. create orm db:
</h5>

``` 
CREATE database job_scheduling; 
```

<br /><br />&emsp;
<h5><<<< CONNECTING >>>> <br />&emsp;&emsp;&emsp;&emsp; 6. connect to db:</h5>

``` 
\c "db" 
```

<br />&emsp;&emsp;&emsp;&emsp;
<h5>
  7. Edit .env file with postgres details <br /><br /><br />
  4. Migrate Tables
</h5>

``` 
go run .\migrations\migrate.go 
```

<br /><br />&emsp;
<h5>View DB Table Schemas:</h5>

``` 
\d "tablename" 
```

<br /><br />

<h5>5. Run the application</h5>

``` 
go run . 
```

<br /><br />
<h5>
  6. Call APIs
  <!-- https://docs.google.com/document/d/1yylvD4j20pHsJ7JqNsHjSVElq8S6DtAcmz9pSSjCIxg/edit?usp=sharing -->
  <a
    href="https://docs.google.com/document/d/1yylvD4j20pHsJ7JqNsHjSVElq8S6DtAcmz9pSSjCIxg/edit?usp=sharing"
    >[API DOC]</a
  >
</h5>
