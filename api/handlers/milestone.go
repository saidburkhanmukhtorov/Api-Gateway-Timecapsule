package handlers

import (
	"context"
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

// MilestoneHandler handles requests related to Milestones.
type MilestoneHandler struct {
	kafkaProducer   *kafka.Producer
	timelineService timeline.MilestoneServiceClient // gRPC client for Timeline Service
}

// NewMilestoneHandler creates a new MilestoneHandler.
func NewMilestoneHandler(kafkaProducer *kafka.Producer, grpcConn *grpc.ClientConn) *MilestoneHandler {
	return &MilestoneHandler{
		kafkaProducer:   kafkaProducer,
		timelineService: timeline.NewMilestoneServiceClient(grpcConn), // Initialize gRPC client
	}
}

// CreateMilestone godoc
// @Summary     Create a new Milestone
// @Description Create a new milestone with the provided details.
// @Tags        Milestones
// @Accept      json
// @Produce     json
// @Param       milestone body     models.CreateMilestoneModel true "Milestone details"
// @Security    ApiKeyAuth
// @Success     202     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/milestones [post]
func (h *MilestoneHandler) CreateMilestone(c *gin.Context) {
	var createMilestone models.CreateMilestoneModel
	if err := c.ShouldBindJSON(&createMilestone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Publish to Kafka
	if err := h.kafkaProducer.ProduceMessage(c.Request.Context(), "milestone.create", createMilestone); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Milestone creation request accepted", "milestone_id": createMilestone.ID})
}

// GetMilestone godoc
// @Summary     Get Milestone by ID
// @Description Get details of a milestone by its ID.
// @Tags        Milestones
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Milestone ID"
// @Security    ApiKeyAuth
// @Success     200     {object} timeline.Milestone
// @Failure     400     {object} map[string]interface{}
// @Failure     404     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/milestones/{id} [get]
func (h *MilestoneHandler) GetMilestone(c *gin.Context) {
	milestoneID := c.Param("id")

	// Use gRPC to get the milestone from the Timeline Service
	grpcResponse, err := h.timelineService.GetMilestoneById(context.Background(), &timeline.GetMilestoneByIdRequest{Id: milestoneID})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Milestone not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}

// UpdateMilestone godoc
// @Summary     Update Milestone
// @Description Update an existing milestone.
// @Tags        Milestones
// @Accept      json
// @Produce     json
// @Param       id     path     string                  true "Milestone ID"
// @Param       milestone body     models.UpdateMilestoneModel true "Updated milestone data"
// @Security    ApiKeyAuth
// @Success     202     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/milestones/{id} [put]
func (h *MilestoneHandler) UpdateMilestone(c *gin.Context) {
	milestoneID := c.Param("id")
	var updateMilestone models.UpdateMilestoneModel
	if err := c.ShouldBindJSON(&updateMilestone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Ensure the ID in the URL matches the ID in the payload
	if updateMilestone.ID != milestoneID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID mismatch"})
		return
	}

	// Publish to Kafka
	if err := h.kafkaProducer.ProduceMessage(c.Request.Context(), "milestone.update", updateMilestone); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Milestone update request accepted"})
}

// PatchMilestone godoc
// @Summary     Partially update Milestone
// @Description Partially update an existing milestone.
// @Tags        Milestones
// @Accept      json
// @Produce     json
// @Param       id     path     string                  true "Milestone ID"
// @Param       milestone body     models.PatchMilestoneModel true "Patched milestone data"
// @Security    ApiKeyAuth
// @Success     202     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/milestones/{id} [patch]
func (h *MilestoneHandler) PatchMilestone(c *gin.Context) {
	milestoneID := c.Param("id")
	var patchMilestone models.PatchMilestoneModel
	if err := c.ShouldBindJSON(&patchMilestone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// You might want to add the ID to the patchModel for clarity:
	patchMilestone.ID = milestoneID

	// Publish to Kafka
	if err := h.kafkaProducer.ProduceMessage(c.Request.Context(), "milestone.patch", patchMilestone); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Milestone patch request accepted"})
}

// DeleteMilestone godoc
// @Summary     Delete Milestone
// @Description Delete a milestone by its ID.
// @Tags        Milestones
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Milestone ID"
// @Security    ApiKeyAuth
// @Success     202     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/milestones/{id} [delete]
func (h *MilestoneHandler) DeleteMilestone(c *gin.Context) {
	milestoneID := c.Param("id")

	// Publish to Kafka
	if err := h.kafkaProducer.ProduceMessage(c.Request.Context(), "milestone.delete", map[string]string{"id": milestoneID}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Milestone delete request accepted"})
}

// ListMilestones godoc
// @Summary     List Milestones
// @Description Get a list of milestones.
// @Tags        Milestones
// @Accept      json
// @Produce     json
// @Param        page     query    int    false  "Page number"                minimum(1)
// @Param        limit    query    int    false  "Number of items per page"  minimum(1)
// @Param        user_id  query    string false  "Filter by user ID"
// @Param        title    query    string false  "Filter by title"
// @Param        category query    string false  "Filter by category"
// @Param        start_date query string false  "Filter by start date (YYYY-MM-DD)"
// @Param        end_date   query string false  "Filter by end date (YYYY-MM-DD)"
// @Security    ApiKeyAuth
// @Success     200     {object} timeline.GetAllMilestonesResponse
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/milestones [get]
func (h *MilestoneHandler) ListMilestones(c *gin.Context) {
	// Get query parameters for pagination and filtering
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	userID := c.Query("user_id")
	title := c.Query("title")
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

	// Use gRPC to get the milestones from the Timeline Service
	grpcResponse, err := h.timelineService.GetAllMilestones(context.Background(), &timeline.GetAllMilestonesRequest{
		Page:      helper.StringToInt(page),
		Limit:     helper.StringToInt(limit),
		UserId:    userID,
		Title:     title,
		Category:  category,
		StartDate: startDate,
		EndDate:   endDate,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}
