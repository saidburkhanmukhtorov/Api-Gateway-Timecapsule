package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/time_capsule/Api-Gateway-Timecapsule/genproto/memory"
	"github.com/time_capsule/Api-Gateway-Timecapsule/helper"
	"github.com/time_capsule/Api-Gateway-Timecapsule/kafka"
	"github.com/time_capsule/Api-Gateway-Timecapsule/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// MediaHandler handles requests related to Media.
type MediaHandler struct {
	kafkaProducer *kafka.Producer
	memoryService memory.MediaServiceClient // gRPC client for Memory Service
}

// NewMediaHandler creates a new MediaHandler.
func NewMediaHandler(kafkaProducer *kafka.Producer, grpcConn *grpc.ClientConn) *MediaHandler {
	return &MediaHandler{
		kafkaProducer: kafkaProducer,
		memoryService: memory.NewMediaServiceClient(grpcConn), // Initialize gRPC client
	}
}

// CreateMedia godoc
// @Summary     Create new Media
// @Description Create new media with the provided details.
// @Tags        Media
// @Accept      json
// @Produce     json
// @Param       media body     models.CreateMediaModel true "Media details"
// @Security    ApiKeyAuth
// @Success     202     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/media [post]
func (h *MediaHandler) CreateMedia(c *gin.Context) {
	var createMedia models.CreateMediaModel
	if err := c.ShouldBindJSON(&createMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// No need to generate ID here, let the service handle it
	// createMedia.ID = uuid.New().String()

	// Publish to Kafka
	if err := h.kafkaProducer.ProduceMessage(c.Request.Context(), "media.create", createMedia); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create media: " + err.Error()}) // Log the specific error
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Media creation request accepted"})
}

// GetMedia godoc
// @Summary     Get Media by ID
// @Description Get details of media by its ID.
// @Tags        Media
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Media ID"
// @Security    ApiKeyAuth
// @Success     200     {object} memory.Media
// @Failure     400     {object} map[string]interface{}
// @Failure     404     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/media/{id} [get]
func (h *MediaHandler) GetMedia(c *gin.Context) {
	mediaID := c.Param("id")

	// Use gRPC to get the media from the Memory Service
	grpcResponse, err := h.memoryService.GetMediaById(context.Background(), &memory.GetMediaByIdRequest{Id: mediaID})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Media not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()}) // Return the gRPC error message
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get media: " + err.Error()}) // Log the specific error
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}

// UpdateMedia godoc
// @Summary     Update Media
// @Description Update existing media.
// @Tags        Media
// @Accept      json
// @Produce     json
// @Param       id     path     string                  true "Media ID"
// @Param       media body     models.UpdateMediaModel true "Updated media data"
// @Security    ApiKeyAuth
// @Success     202     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/media/{id} [put]
func (h *MediaHandler) UpdateMedia(c *gin.Context) {
	mediaID := c.Param("id")
	var updateMedia models.UpdateMediaModel
	if err := c.ShouldBindJSON(&updateMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Ensure the ID in the URL matches the ID in the payload
	if updateMedia.ID != mediaID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID mismatch"})
		return
	}

	// Publish to Kafka
	if err := h.kafkaProducer.ProduceMessage(c.Request.Context(), "media.update", updateMedia); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update media: " + err.Error()}) // Log the specific error
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Media update request accepted"})
}

// PatchMedia godoc
// @Summary     Partially update Media
// @Description Partially update existing media.
// @Tags        Media
// @Accept      json
// @Produce     json
// @Param       id     path     string                  true "Media ID"
// @Param       media body     models.PatchMediaModel true "Patched media data"
// @Security    ApiKeyAuth
// @Success     202     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/media/{id} [patch]
func (h *MediaHandler) PatchMedia(c *gin.Context) {
	mediaID := c.Param("id")
	var patchMedia models.PatchMediaModel
	if err := c.ShouldBindJSON(&patchMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// You might want to add the ID to the patchModel for clarity:
	patchMedia.ID = mediaID

	// Publish to Kafka
	if err := h.kafkaProducer.ProduceMessage(c.Request.Context(), "media.patch", patchMedia); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to patch media: " + err.Error()}) // Log the specific error
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Media patch request accepted"})
}

// DeleteMedia godoc
// @Summary     Delete Media
// @Description Delete media by its ID.
// @Tags        Media
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Media ID"
// @Security    ApiKeyAuth
// @Success     202     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/media/{id} [delete]
func (h *MediaHandler) DeleteMedia(c *gin.Context) {
	mediaID := c.Param("id")

	// Publish to Kafka
	if err := h.kafkaProducer.ProduceMessage(c.Request.Context(), "media.delete", map[string]string{"id": mediaID}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete media: " + err.Error()}) // Log the specific error
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Media delete request accepted"})
}

// ListMedia godoc
// @Summary     List Media
// @Description Get a list of media.
// @Tags        Media
// @Accept      json
// @Produce     json
// @Param        page     query    int    false  "Page number"                minimum(1)
// @Param        limit    query    int    false  "Number of items per page"  minimum(1)
// @Param        memory_id query    string false  "Filter by memory ID"
// @Param        types    query    string false  "Filter by media types (comma-separated)"
// @Security    ApiKeyAuth
// @Success     200     {object} memory.GetAllMediaResponse
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/media [get]
func (h *MediaHandler) ListMedia(c *gin.Context) {
	// Get query parameters for pagination and filtering

	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	memoryID := c.Query("memory_id")
	types := c.Query("types")
	// Use gRPC to get the media from the Memory Service
	grpcResponse, err := h.memoryService.GetAllMedia(context.Background(), &memory.GetAllMediaRequest{
		Page:     helper.StringToInt(page),
		Limit:    helper.StringToInt(limit),
		MemoryId: memoryID,
		Type:     types,
	})
	if err != nil {
		log.Println("Failed to get media:", err) // Log the specific error from the gRPC call
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get media" + err.Error()})
		return
	}
	log.Println(grpcResponse.Media)
	c.JSON(http.StatusOK, grpcResponse)
}
