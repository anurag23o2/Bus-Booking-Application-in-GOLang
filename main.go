package main

import (
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type UserData struct {
	FirstName       string    `json:"firstName"`
	LastName        string    `json:"lastName"`
	Email           string    `json:"email"`
	NumberOfTickets uint      `json:"numberOfTickets"`
	Date            time.Time `json:"date"`        // Date of booking
	Source          string    `json:"source"`      // Source location
	Destination     string    `json:"destination"` // Destination location
}

var bookings []UserData
var mutex sync.Mutex

const maxTickets uint = 100

var totalTicketsSold uint = 0

func main() {
	router := gin.Default()

	// Enable CORS
	router.Use(cors.Default())

	// Add a route for the root URL
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Booking System API!",
		})
	})

	// Get all bookings
	router.GET("/bookings", func(c *gin.Context) {
		c.JSON(http.StatusOK, bookings)
	})

	// Add a new booking with validation
	router.POST("/bookings", func(c *gin.Context) {
		var booking UserData
		if err := c.ShouldBindJSON(&booking); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Validate the input
		if len(booking.FirstName) < 2 || len(booking.LastName) < 2 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "First name and last name must have at least 2 characters"})
			return
		}

		if !strings.Contains(booking.Email, "@") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email must contain '@'"})
			return
		}

		if booking.NumberOfTickets == 0 || booking.NumberOfTickets > 5 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "You can only book between 1 and 5 tickets per booking"})
			return
		}

		if booking.Source == "" || booking.Destination == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Source and destination cannot be empty"})
			return
		}

		if booking.Source == booking.Destination {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Source and destination cannot be the same"})
			return
		}

		if booking.Date.Before(time.Now()) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Date cannot be in the past"})
			return
		}

		mutex.Lock()
		defer mutex.Unlock()

		// Check if total tickets sold exceed the limit
		if totalTicketsSold+booking.NumberOfTickets > maxTickets {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough tickets available"})
			return
		}

		// Check if the email has already booked more than 5 tickets
		totalTicketsByEmail := uint(0)
		for _, b := range bookings {
			if b.Email == booking.Email {
				totalTicketsByEmail += b.NumberOfTickets
			}
		}
		if totalTicketsByEmail+booking.NumberOfTickets > 5 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "You cannot book more than 5 tickets with the same email"})
			return
		}

		// Add the booking
		bookings = append(bookings, booking)
		totalTicketsSold += booking.NumberOfTickets

		c.JSON(http.StatusCreated, gin.H{"message": "Booking added successfully"})
	})

	router.Run(":8080") // Backend runs on http://localhost:8080
}
