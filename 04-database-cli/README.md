# 04 Database CLI

The goal of this project is to understand databases and how to migrate up and down databases.

## Overview

This application is used to build a HR System to manager employees. It should be able to do simple CRUD operations and the data should be persisted after the app is closed.

## Requirements

Should be able to perform crud operations via a cli on a data file of tasks. The operations should be as follows:

```
$ go run main.go users add "Michael"
$ go run main.go users list
$ go run main.go users update 2 "Tommy"
$ go run main.go users delete 2 
```

## Packages
- `strconv` for turning types into strings and visa versa
- `text/tabwriter` for writing out tab aligned output
- `os` for opening and reading files
- `github.com/spf13/cobra` for the command line interface


## Additional Tasks

- Create a new table called department with similar CRUD operations. Users can be part of only 1 departments. Database Migration should be used.
- Create a new table called clients with similar CRUD operations. Users can have 1 or many clients at any time. 
- There should be a Migration seed runner.
- Try to do the same on a different database (postgresql) and use Docker. 
