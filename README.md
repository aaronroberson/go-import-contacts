# Go Importer

## Challenge
Create an simple app in [Go](https://golang.org) that reads an address CSV file (first, last, street address, city, state, zip) and uses Go channels to write records to a database.
Consider validation concerns, duplicate records etc. Report number of records imported successfully, and number of errors.
For extra credit to make a simple AngularJS interface to it.

## Environment

This project uses [Go](https://golang.org) to read in a CSV file and import contacts into a RDBMS [PostgreSQL](https://www.postgresql.org). 
To run this application, you will need to install [Go](https://golang.org) and [PostgreSQL](https://www.postgresql.org) ([Homebrew](http://braumeister.org/formula/postgresql) recommended for MacOS).

## Tasks

- [x] Ramp up on Go lang basics
- [x] Create source control repository
- [x] Create CSV reader and parse contacts
- [x] Choose a database and setup locally
- [ ] Learn Go channel basics
- [ ] Create channels for populating RDBMS
- [ ] Prevent duplicate imports
- [ ] Use transactions to roll back faulty imports
- [ ] Report number of unique records imported
- [ ] Report number of duplicated records ignored
- [ ] Report errors
