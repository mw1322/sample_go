package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "your_module_name/models" // Replace with your actual module path
    "your_module_name/services" // Replace with your actual module path
)

// ServiceController handles requests related to services
type ServiceController struct {
    service services.Service // Assuming you have a service layer
}

// NewServiceController creates a new instance of ServiceController
func NewServiceController(svc services.Service) *ServiceController {
    return &ServiceController{service: svc}
}

// GetAllServices retrieves all services
func (ctrl *ServiceController) GetAllServices(c *gin.Context) {
    services, err := ctrl.service.GetAllServices()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch services"})
        return
    }
    c.JSON(http.StatusOK, services)
}

// GetServiceByID retrieves a specific service by ID
func (ctrl *ServiceController) GetServiceByID(c *gin.Context) {
    id := c.Param("id")
    service, err := ctrl.service.GetServiceByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
        return
    }
    c.JSON(http.StatusOK, service)
}

// CreateService creates a new service
func (ctrl *ServiceController) CreateService(c *gin.Context) {
    var service models.Service
    if err := c.ShouldBindJSON(&service); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }
    createdService, err := ctrl.service.CreateService(service)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create service"})
        return
    }
    c.JSON(http.StatusCreated, createdService)
}

// UpdateService updates an existing service
func (ctrl *ServiceController) UpdateService(c *gin.Context) {
    id := c.Param("id")
    var service models.Service
    if err := c.ShouldBindJSON(&service); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }
    updatedService, err := ctrl.service.UpdateService(id, service)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update service"})
        return
    }
    c.JSON(http.StatusOK, updatedService)
}

// DeleteService deletes a service by ID
func (ctrl *ServiceController) DeleteService(c *gin.Context) {
    id := c.Param("id")
    if err := ctrl.service.DeleteService(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete service"})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}
