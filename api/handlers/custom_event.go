package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/time_capsule/Api-Gateway-Timecapsule/genproto/timeline"
	"github.com/time_capsule/Api-Gateway-Timecapsule/helper"
	"github.com/time_capsule/Api-Gateway-Timecapsule/kafka"
	"github.com/time_capsule/Api-Gateway-Timecapsule/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CustomEventHandler handles requests related to Custom Events.
type CustomEventHandler struct {
	kafkaProducer   *kafka.Producer
	timelineService timeline.CustomEventServiceClient // gRPC client for Timeline Service
}

// NewCustomEventHandler creates a new CustomEventHandler.
func NewCustomEventHandler(kafkaProducer *kafka.Producer, grpcConn *grpc.ClientConn) *CustomEventHandler {
	return &CustomEventHandler{
		kafkaProducer:   kafkaProducer,
		timelineService: timeline.NewCustomEventServiceClient(grpcConn), // Initialize gRPC client
	}
}

// CreateCustomEvent godoc
// @Summary     Create a new Custom Event
// @Description Create a new custom event with the provided details.
// @Tags        CustomEvents
// @Accept      json
// @Produce     json
// @Param       customEvent body     models.CreateCustomEventModel true "Custom Event details"
// @Success     202     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/custom-events [post]
func (h *CustomEventHandler) CreateCustomEvent(c *gin.Context) {
	var createCustomEvent models.CreateCustomEventModel
	if err := c.ShouldBindJSON(&createCustomEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body" + err.Error()})
		return
	}

	// No need to generate ID here, let the service handle it
	// createCustomEvent.ID = uuid.New().String()

	// Publish to Kafka
	if err := h.kafkaProducer.ProduceMessage(c.Request.Context(), "custom_event.create", createCustomEvent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create custom event: " + err.Error()}) // Log the specific error
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Custom event creation request accepted"})
}

// GetCustomEvent godoc
// @Summary     Get Custom Event by ID
// @Description Get details of a custom event by its ID.
// @Tags        CustomEvents
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Custom Event ID"
// @Success     200     {object} timeline.CustomEvent
// @Failure     400     {object} map[string]interface{}
// @Failure     404     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/custom-events/{id} [get]
func (h *CustomEventHandler) GetCustomEvent(c *gin.Context) {
	customEventID := c.Param("id")

	// Use gRPC to get the custom event from the Timeline Service
	grpcResponse, err := h.timelineService.GetCustomEventById(context.Background(), &timeline.GetCustomEventByIdRequest{Id: customEventID})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Custom event not found" + err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()}) // Return the gRPC error message
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get custom event: " + err.Error()}) // Log the specific error
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}

// UpdateCustomEvent godoc
// @Summary     Update Custom Event
// @Description Update an existing custom event.
// @Tags        CustomEvents
// @Accept      json
// @Produce     json
// @Param       id     path     string                  true "Custom Event ID"
// @Param       customEvent body     models.UpdateCustomEventModel true "Updated custom event data"
// @Success     202     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/custom-events/{id} [put]
func (h *CustomEventHandler) UpdateCustomEvent(c *gin.Context) {
	customEventID := c.Param("id")
	var updateCustomEvent models.UpdateCustomEventModel
	if err := c.ShouldBindJSON(&updateCustomEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Ensure the ID in the URL matches the ID in the payload
	if updateCustomEvent.ID != customEventID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID mismatch"})
		return
	}

	// Publish to Kafka
	if err := h.kafkaProducer.ProduceMessage(c.Request.Context(), "custom_event.update", updateCustomEvent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update custom event: " + err.Error()}) // Log the specific error
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Custom event update request accepted"})
}

// PatchCustomEvent godoc
// @Summary     Partially update Custom Event
// @Description Partially update an existing custom event.
// @Tags        CustomEvents
// @Accept      json
// @Produce     json
// @Param       id     path     string                  true "Custom Event ID"
// @Param       customEvent body     models.PatchCustomEventModel true "Patched custom event data"
// @Success     202     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/custom-events/{id} [patch]
func (h *CustomEventHandler) PatchCustomEvent(c *gin.Context) {
	customEventID := c.Param("id")
	var patchCustomEvent models.PatchCustomEventModel
	if err := c.ShouldBindJSON(&patchCustomEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// You might want to add the ID to the patchModel for clarity:
	patchCustomEvent.ID = customEventID

	// Publish to Kafka
	if err := h.kafkaProducer.ProduceMessage(c.Request.Context(), "custom_event.patch", patchCustomEvent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to patch custom event: " + err.Error()}) // Log the specific error
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Custom event patch request accepted"})
}

// DeleteCustomEvent godoc
// @Summary     Delete Custom Event
// @Description Delete a custom event by its ID.
// @Tags        CustomEvents
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Custom Event ID"
// @Success     202     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/custom-events/{id} [delete]
func (h *CustomEventHandler) DeleteCustomEvent(c *gin.Context) {
	customEventID := c.Param("id")

	// Publish to Kafka
	if err := h.kafkaProducer.ProduceMessage(c.Request.Context(), "custom_event.delete", map[string]string{"id": customEventID}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete custom event: " + err.Error()}) // Log the specific error
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Custom event delete request accepted"})
}

// ListCustomEvents godoc
// @Summary     List Custom Events
// @Description Get a list of custom events.
// @Tags        CustomEvents
// @Accept      json
// @Produce     json
// @Param        page     query    int    false  "Page number"                minimum(1)
// @Param        limit    query    int    false  "Number of items per page"  minimum(1)
// @Param        user_id  query    string false  "Filter by user ID"
// @Param        title    query    string false  "Filter by title"
// @Param        description query    string false  "Filter by description"
// @Param        category query    string false  "Filter by category"
// @Param        start_date query string false  "Filter by start date (YYYY-MM-DD)"
// @Param        end_date   query string false  "Filter by end date (YYYY-MM-DD)"
// @Success     200     {object} timeline.GetAllCustomEventsResponse
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/custom-events [get]
func (h *CustomEventHandler) ListCustomEvents(c *gin.Context) {
	// Get query parameters for pagination and filtering
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	userID := c.Query("user_id")
	title := c.Query("title")
	description := c.Query("description")
	category := c.Query("category")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	// Validate date formats if provided
	if startDate != "" {
		if _, err := time.Parse("2006-01-02", startDate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format. Use YYYY-MM-DD"})
			return
		}
	}
	if endDate != "" {
		if _, err := time.Parse("2006-01-02", endDate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format. Use YYYY-MM-DD"})
			return
		}
	}

	// Use gRPC to get the custom events from the Timeline Service
	grpcResponse, err := h.timelineService.GetAllCustomEvents(context.Background(), &timeline.GetAllCustomEventsRequest{
		Page:        helper.StringToInt(page),  // Use helper function to convert string to int32
		Limit:       helper.StringToInt(limit), // Use helper function to convert string to int32
		UserId:      userID,
		Title:       title,
		Description: description,
		Category:    category,
		StartDate:   startDate,
		EndDate:     endDate,
	})
	if err != nil {
		log.Println("Failed to get custom events:", err) // Log the specific error from the gRPC call
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get custom events"})
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}
