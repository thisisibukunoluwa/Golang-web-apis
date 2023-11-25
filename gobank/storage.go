package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error 
	DeleteAccount(int) error 
	UpdateAccount(*Account) error 
	GetAccounts() ([]*Account, error)
	GetAccountByID(int)(*Account, error)
	GetAccountByNumber(int)(*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db,err := sql.Open("postgres",connStr)
	if err != nil {
		return nil,err
	}
	if err := db.Ping(); err != nil {
		return nil,err
	}
	return &PostgresStore{
		db:db,
	}, nil
}

func (s* PostgresStore) Init() error {
	return s.createAccountTable()
}


func (s* PostgresStore) createAccountTable() error {
	query := `create table if not exists account (
		id serial primary key,
		first_name varchar(100),
		last_name varchar(100),
		number serial,
		encrypted_password varchar(100),
		balance serial, 
		created_at timestamp
	)`
	_, err := s.db.Exec(query)
	// even when i remove encrypted_password varchar(100), its still the same thing 
	// if err != nil {
	// 	return fmt.Errorf("failed to create account table heree: %v", err)
	// }
	return err 
}

// it led to here 
func (s *PostgresStore) CreateAccount(acc *Account) error {
	query := `insert into account 
	(first_name, last_name, number, encrypted_password, balance, created_at)
	values ($1, $2, $3, $4, $5, $6)`

	resp,err := s.db.Query(
			query,
			acc.FirstName,
		 	acc.LastName, 
		 	acc.Number,
		  	acc.EncryptedPassword, 
		  	acc.Balance, 
		  	acc.CreatedAt)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)
	
	return nil
}
func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}
func (s *PostgresStore) DeleteAccount(id int) error {
	_, err := s.db.Query("delete from account where id = $1", id)
	return err
}

func (s *PostgresStore) GetAccountByNumber(number int) (*Account, error) {
	rows, err := s.db.Query("select * from account where number = $1", number)
	if err != nil {
		return nil,err
	}
	for rows.Next() {
		return scanIntoAccount(rows)
	}
	return nil , fmt.Errorf("account with number [%d] not found", number)
}

func (s *PostgresStore) GetAccountByID(id int) (*Account,error) {
	rows, err := s.db.Query("select * from account where id = $1", id)
	if err != nil {
		return nil,err
	}
	for rows.Next() {
		return scanIntoAccount(rows)
	}
	return nil , fmt.Errorf("account %d not found", id)
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	 rows, err := s.db.Query("select * from account")
	  if err != nil {
			return nil, err 
	 }
	 accounts := []*Account{}
	 for rows.Next() {
		 account, err := scanIntoAccount(rows)
		 if err != nil {
			return nil, err 
		 }
		 accounts = append(accounts, account)
		}
	 return accounts,nil
}

func scanIntoAccount( rows *sql.Rows) (*Account, error) {
	 	 account := new(Account)
		 err := rows.Scan(
			&account.ID, 
			&account.FirstName,
			&account.LastName, 
			&account.Number,
			&account.EncryptedPassword,
			&account.Balance,
			&account.CreatedAt,
		)
		return account, err 
}
// Start a postgres instance 
//docker run --name some-postgres -e POSTGRES_PASSWORD=gobank -p 5432:5432 -d postgres

//docker ps lists the Docker containers that are currently running on our system. 

// check if docker is running
// sudo launchctl list | grep docker

// Got an error
//Cannot connect to the Docker daemon at unix:///Users/ibukunoluwaakintobi/.docker/run/docker.sock. Is the docker daemon running?.

// before you can run an instance , you first have to start docker , oh to do that you log in to the desktop app and start it 


//https://pkg.go.dev/github.com/lib/pq


// was having this error 
//dial tcp [::1]:5432: connect: connection refused
// make: *** [run] Error 1

// Explanation
//The error message you're seeing, "dial tcp [::1]:5432: connect: connection refused," is related to a network connection issue. It suggests that your Go application is trying to connect to a PostgreSQL database on localhost ([::1]:5432), but the connection is being refused.

// so i ran the container as i noticed it was probably not on 
// Container name already in use 
//https://www.baeldung.com/ops/docker-name-already-in-use

//Error response from daemon: Conflict. The container name "/some-postgres" is already in use by container "38aba21d76522ab8ea8a47428a16716554bbb9c0bcb73432d9da016fe817d997". You have to remove (or rename) that container to be able to reuse that name.
// See 'docker run --help'.

//i did docker rm some-postgres to remove the instance 

// i reran "make run" then i saw this 
// 2023/10/01 21:39:33 pq: password authentication failed for user "pqgotest"
// make: *** [run] Error 1

// If i end a route with a slash that does not have a query param , i am going to get a 404 error 

// eg if i have a route localhost:3000/account , and instead i put in postman localhost:3000/account/

// i was having a recurrent error -> column "encrypted_password" of relation "account" does not exist tunrs out the problem was that i had updated the number of columns in my postgres databse , but i did not create a new instance , so it was stoill trying to access a non existent column on an old instance 