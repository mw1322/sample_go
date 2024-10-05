package repositories

import (
    "database/sql"
    "your_module_name/models" // Replace with your actual module path
)

type ServiceRepository interface {
    FetchAllServices() ([]models.Service, error)
    FetchServiceByID(id string) (*models.Service, error)
    SaveService(service models.Service) (*models.Service, error)
    RemoveService(id string) error
}

type serviceRepository struct {
    db *sql.DB // Database connection
}

func NewServiceRepository(db *sql.DB) ServiceRepository {
    return &serviceRepository{db: db}
}

func (repo *serviceRepository) FetchAllServices() ([]models.Service, error) {
    var services []models.Service
    rows, err := repo.db.Query("SELECT * FROM services")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var service models.Service
        if err := rows.Scan(&service.ID, &service.Name, &service.Description, &service.Type, &service.InstructorID, &service.Duration, &service.Price, &service.CreatedAt, &service.UpdatedAt); err != nil {
            return nil, err
        }
        services = append(services, service)
    }
    return services, nil
}

func (repo *serviceRepository) FetchServiceByID(id string) (*models.Service, error) {
    var service models.Service
    row := repo.db.QueryRow("SELECT * FROM services WHERE id = ?", id)
    if err := row.Scan(&service.ID, &service.Name, &service.Description, &service.Type, &service.InstructorID, &service.Duration, &service.Price, &service.CreatedAt, &service.UpdatedAt); err != nil {
        return nil, err
    }
    return &service, nil
}

func (repo *serviceRepository) SaveService(service models.Service) (*models.Service, error) {
    if service.ID == 0 {
        // Create new service
        result, err := repo.db.Exec("INSERT INTO services (name, description, type, instructor_id, duration, price) VALUES (?, ?, ?, ?, ?, ?)",
            service.Name, service.Description, service.Type, service.InstructorID, service.Duration, service.Price)
        if err != nil {
            return nil, err
        }
        id, err := result.LastInsertId()
        if err != nil {
            return nil, err
        }
        service.ID = int(id)
    } else {
        // Update existing service
        _, err := repo.db.Exec("UPDATE services SET name = ?, description = ?, type = ?, instructor_id = ?, duration = ?, price = ? WHERE id = ?",
            service.Name, service.Description, service.Type, service.InstructorID, service.Duration, service.Price, service.ID)
        if err != nil {
            return nil, err
        }
    }
    return &service, nil
}

func (repo *serviceRepository) RemoveService(id string) error {
    _, err := repo.db.Exec("DELETE FROM services WHERE id = ?", id)
    return err
}
