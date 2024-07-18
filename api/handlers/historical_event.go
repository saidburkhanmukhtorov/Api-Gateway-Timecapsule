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

// HistoricalEventHandler handles requests related to Historical Events.
type HistoricalEventHandler struct {
	kafkaProducer   *kafka.Producer
	timelineService timeline.HistoricalEventServiceClient // gRPC client for Timeline Service
}

// NewHistoricalEventHandler creates a new HistoricalEventHandler.
func NewHistoricalEventHandler(kafkaProducer *kafka.Producer, grpcConn *grpc.ClientConn) *HistoricalEventHandler {
	return &HistoricalEventHandler{
		kafkaProducer:   kafkaProducer,
		timelineService: timeline.NewHistoricalEventServiceClient(grpcConn), // Initialize gRPC client
	}
}

// CreateHistoricalEvent godoc
// @Summary     Create a new Historical Event
// @Description Create a new historical event with the provided details.
// @Tags        HistoricalEvents
// @Accept      json
// @Produce     json
// @Param       historicalEvent body     models.CreateHistoricalEventModel true "Historical Event details"
// @Success     202     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/historical-events [post]
func (h *HistoricalEventHandler) CreateHistoricalEvent(c *gin.Context) {
	var createHistoricalEvent models.CreateHistoricalEventModel
	if err := c.ShouldBindJSON(&createHistoricalEvent); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// No need to generate ID here, let the service handle it
	// createHistoricalEvent.ID = uuid.New().String()

	// Publish to Kafka
	if err := h.kafkaProducer.ProduceMessage(c.Request.Context(), "historical_event.create", createHistoricalEvent); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create historical event"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Historical event creation request accepted"})
}

// GetHistoricalEvent godoc
// @Summary     Get Historical Event by ID
// @Description Get details of a historical event by its ID.
// @Tags        HistoricalEvents
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Historical Event ID"
// @Success     200     {object} timeline.HistoricalEvent
// @Failure     400     {object} map[string]interface{}
// @Failure     404     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/historical-events/{id} [get]
func (h *HistoricalEventHandler) GetHistoricalEvent(c *gin.Context) {
	historicalEventID := c.Param("id")

	// Use gRPC to get the historical event from the Timeline Service
	grpcResponse, err := h.timelineService.GetHistoricalEventById(context.Background(), &timeline.GetHistoricalEventByIdRequest{Id: historicalEventID})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				log.Println(err)
				c.JSON(http.StatusNotFound, gin.H{"error": "Historical event not found"})
				return
			}
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
			return
		}
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get historical event"})
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}

// UpdateHistoricalEvent godoc
// @Summary     Update Historical Event
// @Description Update an existing historical event.
// @Tags        HistoricalEvents
// @Accept      json
// @Produce     json
// @Param       id     path     string   true "Historical Event ID"
// @Param       historicalEvent body     models.UpdateHistoricalEventModel true "Updated historical event data"
// @Success     202     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/historical-events/{id} [put]
func (h *HistoricalEventHandler) UpdateHistoricalEvent(c *gin.Context) {
	historicalEventID := c.Param("id")
	var updateHistoricalEvent models.UpdateHistoricalEventModel
	if err := c.ShouldBindJSON(&updateHistoricalEvent); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Ensure the ID in the URL matches the ID in the payload
	if updateHistoricalEvent.ID != historicalEventID {

		c.JSON(http.StatusBadRequest, gin.H{"error": "ID mismatch"})
		return
	}

	// Publish to Kafka
	if err := h.kafkaProducer.ProduceMessage(c.Request.Context(), "historical_event.update", updateHistoricalEvent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update historical event"})
		log.Println(err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Historical event update request accepted"})
}

// PatchHistoricalEvent godoc
// @Summary     Partially update Historical Event
// @Description Partially update an existing historical event.
// @Tags        HistoricalEvents
// @Accept      json
// @Produce     json
// @Param       id     path     string                  true "Historical Event ID"
// @Param       historicalEvent body     models.PatchHistoricalEventModel true "Patched historical event data"
// @Success     202     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/historical-events/{id} [patch]
func (h *HistoricalEventHandler) PatchHistoricalEvent(c *gin.Context) {
	historicalEventID := c.Param("id")
	var patchHistoricalEvent models.PatchHistoricalEventModel
	if err := c.ShouldBindJSON(&patchHistoricalEvent); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// You might want to add the ID to the patchModel for clarity:
	patchHistoricalEvent.ID = historicalEventID

	// Publish to Kafka
	if err := h.kafkaProducer.ProduceMessage(c.Request.Context(), "historical_event.patch", patchHistoricalEvent); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to patch historical event"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Historical event patch request accepted"})
}

// DeleteHistoricalEvent godoc
// @Summary     Delete Historical Event
// @Description Delete a historical event by its ID.
// @Tags        HistoricalEvents
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Historical Event ID"
// @Success     202     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/historical-events/{id} [delete]
func (h *HistoricalEventHandler) DeleteHistoricalEvent(c *gin.Context) {
	historicalEventID := c.Param("id")

	// Publish to Kafka
	if err := h.kafkaProducer.ProduceMessage(c.Request.Context(), "historical_event.delete", map[string]string{"id": historicalEventID}); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete historical event"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Historical event delete request accepted"})
}

// ListHistoricalEvents godoc
// @Summary     List Historical Events
// @Description Get a list of historical events.
// @Tags        HistoricalEvents
// @Accept      json
// @Produce     json
// @Param        page     query    int    false  "Page number"                minimum(1)
// @Param        limit    query    int    false  "Number of items per page"  minimum(1)
// @Param        title    query    string false  "Filter by title"
// @Param        category query    string false  "Filter by category"
// @Param        start_date query string false  "Filter by start date (YYYY-MM-DD)"
// @Param        end_date   query string false  "Filter by end date (YYYY-MM-DD)"
// @Success     200     {object} timeline.GetAllHistoricalEventsResponse
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/historical-events [get]
func (h *HistoricalEventHandler) ListHistoricalEvents(c *gin.Context) {
	// Get query parameters for pagination and filtering
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	title := c.Query("title")
	category := c.Query("category")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	// Validate date formats if provided
	if startDate != "" {
		if _, err := time.Parse("2006-01-02", startDate); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format. Use YYYY-MM-DD"})
			return
		}
	}
	if endDate != "" {
		if _, err := time.Parse("2006-01-02", endDate); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format. Use YYYY-MM-DD"})
			return
		}
	}

	// Use gRPC to get the historical events from the Timeline Service
	grpcResponse, err := h.timelineService.GetAllHistoricalEvents(context.Background(), &timeline.GetAllHistoricalEventsRequest{
		Page:  helper.StringToInt(page),
		Limit: helper.StringToInt(limit),

		Title:     title,
		Category:  category,
		StartDate: startDate,
		EndDate:   endDate,
	})
	if err != nil {
		log.Println("Failed to get historical events", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get historical events"})
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}
