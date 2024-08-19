package database

import (
	"context"
	"database/sql"
	types "farmaIA/internal/services"
	users "farmaIA/internal/user"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type Service interface {
	Health() map[string]string
    HelloWorld()  map[string]string
    GetUserTest()  map[string]string
    Gemini()
	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error
	GetTeste() (map[string]string, error)
	GetSurveys() ([]types.Survey, error)
    GetUsers() ([]users.User,error) 
}

type service struct {
	db *sql.DB
}

var (
	dbname     = os.Getenv("DB_DATABASE")
	password   = os.Getenv("DB_PASSWORD")
	username   = os.Getenv("DB_USERNAME")
	port       = os.Getenv("DB_PORT")
	host       = os.Getenv("DB_HOST")
	dbInstance *service
)

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname))
	if err != nil {
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf(fmt.Sprintf("db down: %v", err))
		return stats
	}

	stats["status"] = "up"
	stats["message"] = "It's healthy"

	dbStats := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}
	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

func (s *service) HelloWorld()  map[string]string{ 
	resp := make(map[string]string)
	resp["message"] = "Teste Farma"
    return resp
}

func (s *service) GetUserTest()  map[string]string{ 
	resp := make(map[string]string)
	resp["id"] = "1"
	resp["name"] = "jose"
	resp["email"] = "j@j.com"
	resp["password"] = "123"
	resp["role"] = "admin"
    return resp
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", dbname)
	return s.db.Close()
}

func (s *service) GetTeste() (map[string]string, error) {
	stats := make(map[string]string)
    stats["message"] = "banco"

    rows, err := s.db.Query("select * from test_table")
    if err != nil {
        log.Fatalf("error handling JSON marshal. Err: %v", err)
    }
    defer rows.Close()

    result := make(map[string]string)
    for rows.Next() {
        var id int
        var name string
        if err := rows.Scan(&id, &name); err != nil {
            log.Fatalf("error handling JSON marshal. Err: %v", err)
        }
        result[strconv.Itoa(id)] = name
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

	return result, nil
}

func (s *service) GetSurveys() ([]types.Survey, error) {
    rows, err := s.db.Query("select * from survey")
    if err != nil {
        log.Fatalf("Error on executing query. Err: %v", err)
    }
    defer rows.Close()

    result := make([]types.Survey, 0)
    for rows.Next() {
        survey := types.Survey{}
        if err := rows.Scan(&survey.Id, &survey.Nome, &survey.DataCriacao); err != nil {
            log.Fatalf("Error reading database. Err: %v", err)
        }
        result = append(result, survey)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

	return result, nil
}

func (s *service) GetUsers() ([]users.User, error) {
    rows, err := s.db.Query("select * from user")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    result := make([]users.User, 0)
    for rows.Next() {
        user := users.User{}
        if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role); err != nil {
            log.Fatalf("Error reading database. Err: %v", err)
        }
        result = append(result, user)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

	return result, nil
}

func (s *service) Gemini() {
    fmt.Println("do banco aqui")
}

