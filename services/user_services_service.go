package services

import (
    "errors"
    "your_project/repositories" // Adjust this import path according to your project structure
)

// UserService handles user services related operations
type UserService struct {
    UserServicesRepo repositories.UserServicesRepository // Assuming you have a repository layer
}

// NewUserService creates a new instance of UserService
func NewUserService(userServicesRepo repositories.UserServicesRepository) *UserService {
    return &UserService{UserServicesRepo: userServicesRepo}
}

// BookService books a service for a user
func (us *UserService) BookService(userID, serviceID int) error {
    // Here you might include business logic to check if the user can book the service
    // For example, checking if the user is already booked for that service

    // Add booking to the repository
    err := us.UserServicesRepo.AddBooking(userID, serviceID)
    if err != nil {
        return err
    }

    return nil
}

// GetUserBookings retrieves all services booked by a user
func (us *UserService) GetUserBookings(userID int) ([]UserServiceBooking, error) {
    bookings, err := us.UserServicesRepo.GetBookingsByUserID(userID)
    if err != nil {
        return nil, err
    }

    return bookings, nil
}

// CancelBooking allows a user to cancel a booking
func (us *UserService) CancelBooking(userID, serviceID int) error {
    // Here you might include business logic to check if the user can cancel the booking

    err := us.UserServicesRepo.RemoveBooking(userID, serviceID)
    if err != nil {
        return err
    }

    return nil
}
