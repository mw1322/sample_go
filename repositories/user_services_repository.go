package repositories

import (
    "database/sql"
    "your_project/models" // Adjust this import path according to your project structure
)

type UserServicesRepo struct {
    DB *sql.DB
}

// AddBooking adds a new booking to the user_services table
func (repo *UserServicesRepo) AddBooking(userID, serviceID int) error {
    _, err := repo.DB.Exec("INSERT INTO user_services (user_id, service_id) VALUES (?, ?)", userID, serviceID)
    return err
}

// GetBookingsByUserID retrieves bookings for a specific user
func (repo *UserServicesRepo) GetBookingsByUserID(userID int) ([]UserServiceBooking, error) {
    rows, err := repo.DB.Query("SELECT service_id, booking_date FROM user_services WHERE user_id = ?", userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var bookings []UserServiceBooking
    for rows.Next() {
        var booking UserServiceBooking
        if err := rows.Scan(&booking.ServiceID, &booking.BookingDate); err != nil {
            return nil, err
        }
        bookings = append(bookings, booking)
    }

    return bookings, nil
}

// RemoveBooking removes a booking from the user_services table
func (repo *UserServicesRepo) RemoveBooking(userID, serviceID int) error {
    _, err := repo.DB.Exec("DELETE FROM user_services WHERE user_id = ? AND service_id = ?", userID, serviceID)
    return err
}
