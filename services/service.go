package services

import (
    "errors"
    "your_module_name/models" // Replace with your actual module path
    "your_module_name/repositories" // Replace with your actual module path
)

// Service defines the methods for managing services
type Service interface {
    GetAllServices() ([]models.Service, error)
    GetServiceByID(id string) (*models.Service, error)
    CreateService(service models.Service) (*models.Service, error)
    UpdateService(id string, service models.Service) (*models.Service, error)
    DeleteService(id string) error
}

// service is the implementation of the Service interface
type service struct {
    repo repositories.ServiceRepository // Assuming you have a repository layer
}

// NewService creates a new instance of service
func NewService(repo repositories.ServiceRepository) Service {
    return &service{repo: repo}
}

// GetAllServices retrieves all services from the repository
func (s *service) GetAllServices() ([]models.Service, error) {
    return s.repo.FetchAllServices()
}

// GetServiceByID retrieves a specific service by ID from the repository
func (s *service) GetServiceByID(id string) (*models.Service, error) {
    service, err := s.repo.FetchServiceByID(id)
    if err != nil {
        return nil, err
    }
    return service, nil
}

// CreateService creates a new service in the repository
func (s *service) CreateService(service models.Service) (*models.Service, error) {
    if service.Name == "" {
        return nil, errors.New("service name is required")
    }
    return s.repo.SaveService(service)
}

// UpdateService updates an existing service in the repository
func (s *service) UpdateService(id string, service models.Service) (*models.Service, error) {
    existingService, err := s.repo.FetchServiceByID(id)
    if err != nil {
        return nil, err
    }
    existingService.Name = service.Name
    existingService.Description = service.Description
    existingService.Type = service.Type
    existingService.InstructorID = service.InstructorID
    existingService.Duration = service.Duration
    existingService.Price = service.Price

    return s.repo.SaveService(*existingService)
}

// DeleteService deletes a service from the repository
func (s *service) DeleteService(id string) error {
    return s.repo.RemoveService(id)
}
