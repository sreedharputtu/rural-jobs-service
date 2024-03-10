package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	// Import mysql driver
)

// Define connection details (replace with your actual credentials)
const (
	dbUser     = "your_username"
	dbPassword = "your_password"
	dbHost     = "localhost"
	dbPort     = "3306"
	dbName     = "your_database_name"
)

// Define structs for users and jobs
type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	UserType     string `json:"user_type"`
	PhoneNumber  string `json:"phone_number"`
	RuralArea    string `json:"rural_area"`
}

type Job struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CompanyName string `json:"company_name"`
	Location    string `json:"location"`
	Salary      int    `json:"salary"`
	PostedDate  string `json:"posted_date"`
	Category    string `json:"category"`
	IsActive    bool   `json:"is_active"`
	UserID      int    `json:"user_id"`
}

var db *sql.DB

func main() {
	// Open database connection
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize Gin router
	router := gin.Default()

	// User API routes
	router.POST("/users", createUser)

	// Job API routes
	router.POST("/jobs", createJob)
	router.GET("/jobs", getJobs) // Add route for GET /jobs

	// Run the server
	router.Run(":8080")
}

func createUser(c *gin.Context) {
	var user User
	// Bind JSON request body to user struct
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password before storing (implementation omitted for brevity)
	// hashedPassword, err := hashPassword(user.PasswordHash)
	// if err != nil {
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//     return
	// }

	hashedPassword := ""
	// Insert user data into database
	stmt, err := prepareUserInsert(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.Email, hashedPassword, user.UserType, user.PhoneNumber, user.RuralArea)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get the newly created user ID
	userID, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.ID = int(userID)
	c.JSON(http.StatusCreated, user)
}

func createJob(c *gin.Context) {
	var job Job
	// Bind JSON request body to job struct
	if err := c.BindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert job data into database
	stmt, err := prepareJobInsert(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = stmt.Exec()
	if err != nil {

	}
}

func getJobs(c *gin.Context) {
	// Prepare SQL statement to fetch jobs
	stmt, err := db.Prepare("SELECT * FROM Jobs WHERE is_active = true")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	// Execute the statement and retrieve results
	rows, err := stmt.Query()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var jobs []Job
	// Scan each row and append job data to slice
	for rows.Next() {
		var job Job
		err := rows.Scan(&job.ID, &job.Title, &job.Description, &job.CompanyName, &job.Location, &job.Salary, &job.PostedDate, &job.Category, &job.IsActive, &job.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		jobs = append(jobs, job)
	}

	// Respond with list of jobs in JSON format
	c.JSON(http.StatusOK, jobs)
}

// Helper functions to prepare SQL statements (implementation omitted for brevity)
func prepareUserInsert(db *sql.DB) (*sql.Stmt, error) {
	// ... (implement SQL statement to insert user data)
	return nil, nil
}

func prepareJobInsert(db *sql.DB) (*sql.Stmt, error) {
	// ... (implement SQL statement to insert job data)
	return nil, nil
}
