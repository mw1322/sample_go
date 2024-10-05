package repositories

type UserServicesRepository interface {
    AddBooking(userID, serviceID int) error
    GetBookingsByUserID(userID int) ([]UserServiceBooking, error)
    RemoveBooking(userID, serviceID int) error
}
