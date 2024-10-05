package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "your_project/services" // Adjust this import path according to your project structure
)

// UserServicesController handles user services-related requests
type UserServicesController struct {
    UserService services.UserService // Assuming you have a UserService layer for handling business logic
}

// NewUserServicesController creates a new instance of UserServicesController
func NewUserServicesController(userService services.UserService) *UserServicesController {
    return &UserServicesController{UserService: userService}
}

// BookService allows a user to book a service
func (ctrl *UserServicesController) BookService(c *gin.Context) {
    var bookingRequest struct {
        UserID    int `json:"user_id"`
        ServiceID int `json:"service_id"`
    }

    if err := c.ShouldBindJSON(&bookingRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    err := ctrl.UserService.BookService(bookingRequest.UserID, bookingRequest.ServiceID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Service booked successfully"})
}

// GetUserBookings retrieves all services booked by a user
func (ctrl *UserServicesController) GetUserBookings(c *gin.Context) {
    userID := c.Param("user_id")

    bookings, err := ctrl.UserService.GetUserBookings(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, bookings)
}

// CancelBooking allows a user to cancel a booking
func (ctrl *UserServicesController) CancelBooking(c *gin.Context) {
    var cancelRequest struct {
        UserID    int `json:"user_id"`
        ServiceID int `json:"service_id"`
    }

    if err := c.ShouldBindJSON(&cancelRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    err := ctrl.UserService.CancelBooking(cancelRequest.UserID, cancelRequest.ServiceID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Booking cancelled successfully"})
}
