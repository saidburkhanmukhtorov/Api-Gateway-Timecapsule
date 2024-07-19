package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/time_capsule/Api-Gateway-Timecapsule/genproto/memory"
	"github.com/time_capsule/Api-Gateway-Timecapsule/helper"
	"github.com/time_capsule/Api-Gateway-Timecapsule/kafka"
	"github.com/time_capsule/Api-Gateway-Timecapsule/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// MemoryHandler handles requests related to Memories.
type MemoryHandler struct {
	kafkaProducer *kafka.Producer
	memoryService memory.MemoryServiceClient // gRPC client for Memory Service
}

// NewMemoryHandler creates a new MemoryHandler.
func NewMemoryHandler(kafkaProducer *kafka.Producer, grpcConn *grpc.ClientConn) *MemoryHandler {
	return &MemoryHandler{
		kafkaProducer: kafkaProducer,
		memoryService: memory.NewMemoryServiceClient(grpcConn), // Initialize gRPC client
	}
}

// CreateMemory godoc
// @Summary     Create a new Memory
// @Description Create a new memory with the provided details.
// @Tags        Memories
// @Accept      json
// @Produce     json
// @Param       memory body     models.CreateMemoryModel true "Memory details"
// @Security    ApiKeyAuth
// @Success     202     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/memories [post]
func (h *MemoryHandler) CreateMemory(c *gin.Context) {
	var createMemory models.CreateMemoryModel
	if err := c.ShouldBindJSON(&createMemory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body" + err.Error()})
		return
	}

	// No need to generate ID here, let the service handle it
	// createMemory.ID = uuid.New().String()

	// Publish to Kafka
	if err := h.kafkaProducer.ProduceMessage(c.Request.Context(), "memory.create", createMemory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create memory: " + err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Memory creation request accepted"})
}

// GetMemory godoc
// @Summary     Get Memory by ID
// @Description Get details of a memory by its ID.
// @Tags        Memories
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Memory ID"
// @Security    ApiKeyAuth
// @Success     200     {object} memory.Memory
// @Failure     400     {object} map[string]interface{}
// @Failure     404     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/memories/{id} [get]
func (h *MemoryHandler) GetMemory(c *gin.Context) {
	memoryID := c.Param("id")

	// Use gRPC to get the memory from the Memory Service
	grpcResponse, err := h.memoryService.GetMemoryById(context.Background(), &memory.GetMemoryByIdRequest{Id: memoryID})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Memory not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()}) // Return the gRPC error message
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get memory: " + err.Error()}) // Log the specific error
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}

// UpdateMemory godoc
// @Summary     Update Memory
// @Description Update an existing memory.
// @Tags        Memories
// @Accept      json
// @Produce     json
// @Param       id     path     string                  true "Memory ID"
// @Param       memory body     models.UpdateMemoryModel true "Updated memory data"
// @Security    ApiKeyAuth
// @Success     202     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/memories/{id} [put]
func (h *MemoryHandler) UpdateMemory(c *gin.Context) {
	memoryID := c.Param("id")
	var updateMemory models.UpdateMemoryModel
	if err := c.ShouldBindJSON(&updateMemory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Ensure the ID in the URL matches the ID in the payload
	if updateMemory.ID != memoryID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID mismatch"})
		return
	}

	// Publish to Kafka
	if err := h.kafkaProducer.ProduceMessage(c.Request.Context(), "memory.update", updateMemory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update memory: " + err.Error()}) // Log the specific error
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Memory update request accepted"})
}

// PatchMemory godoc
// @Summary     Partially update Memory
// @Description Partially update an existing memory.
// @Tags        Memories
// @Accept      json
// @Produce     json
// @Param       id     path     string                  true "Memory ID"
// @Param       memory body     models.PatchMemoryModel true "Patched memory data"
// @Security    ApiKeyAuth
// @Success     202     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/memories/{id} [patch]
func (h *MemoryHandler) PatchMemory(c *gin.Context) {
	memoryID := c.Param("id")
	var patchMemory models.PatchMemoryModel
	if err := c.ShouldBindJSON(&patchMemory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// You might want to add the ID to the patchModel for clarity:
	patchMemory.ID = memoryID

	// Publish to Kafka
	if err := h.kafkaProducer.ProduceMessage(c.Request.Context(), "memory.patch", patchMemory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to patch memory: " + err.Error()}) // Log the specific error
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Memory patch request accepted"})
}

// DeleteMemory godoc
// @Summary     Delete Memory
// @Description Delete a memory by its ID.
// @Tags        Memories
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Memory ID"
// @Security    ApiKeyAuth
// @Success     202     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/memories/{id} [delete]
func (h *MemoryHandler) DeleteMemory(c *gin.Context) {
	memoryID := c.Param("id")

	// Publish to Kafka
	if err := h.kafkaProducer.ProduceMessage(c.Request.Context(), "memory.delete", map[string]string{"id": memoryID}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete memory: " + err.Error()}) // Log the specific error
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Memory delete request accepted"})
}

// ListMemories godoc
// @Summary     List Memories
// @Description Get a list of memories.
// @Tags        Memories
// @Accept      json
// @Produce     json
// @Param        page         query    int     false  "Page number"                minimum(1)
// @Param        limit        query    int     false  "Number of items per page"  minimum(1)
// @Param        search_term  query    string  false  "Search by title, description, or tags"
// @Param        tags         query    string  false  "Filter by multiple tags (comma-separated)"
// @Param        start_date   query    string  false  "Filter by start date (YYYY-MM-DD)"
// @Param        end_date     query    string  false  "Filter by end date (YYYY-MM-DD)"
// @Param        user_id      query    string  false  "Filter by user ID"
// @Param        title        query    string  false  "Filter by title"
// @Param        description  query    string  false  "Filter by description"
// @Param        latitude     query    float64 false  "Filter by latitude"
// @Param        longitude    query    float64 false  "Filter by longitude"
// @Param        place_name   query    string  false  "Filter by place name"
// @Param        privacy      query    string  false  "Filter by privacy setting"
// @Security    ApiKeyAuth
// @Success     200     {object} memory.GetAllMemoriesResponse
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/memories [get]
func (h *MemoryHandler) ListMemories(c *gin.Context) {
	var tagsL []string
	// Get query parameters for pagination and filtering
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	searchTerm := c.Query("search_term")
	tags := c.Query("tags")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	userID := c.Query("user_id")
	title := c.Query("title")
	description := c.Query("description")
	latitudeStr := c.Query("latitude")
	longitudeStr := c.Query("longitude")
	placeName := c.Query("place_name")
	privacy := c.Query("privacy")

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

	// Parse latitude and longitude if provided
	var latitude, longitude float64
	if latitudeStr != "" {
		if lat, err := strconv.ParseFloat(latitudeStr, 64); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid latitude format"})
			return
		} else {
			latitude = lat
		}
	}
	if longitudeStr != "" {
		if lon, err := strconv.ParseFloat(longitudeStr, 64); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid longitude format"})
			return
		} else {
			longitude = lon
		}
	}
	if tags != "" {
		tagsL = strings.Split(tags, ",")
	}
	// Use gRPC to get the memories from the Memory Service
	grpcResponse, err := h.memoryService.GetAllMemories(context.Background(), &memory.GetAllMemoriesRequest{
		Page:        helper.StringToInt(page),
		Limit:       helper.StringToInt(limit),
		SearchTerm:  searchTerm,
		Tags:        tagsL, // Split comma-separated tags into a slice
		StartDate:   startDate,
		EndDate:     endDate,
		UserId:      userID,
		Title:       title,
		Description: description,
		Latitude:    latitude,
		Longitude:   longitude,
		PlaceName:   placeName,
		Privacy:     privacy,
	})
	if err != nil {
		log.Println("Failed to get memories:", err) // Log the specific error from the gRPC call
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get memories" + err.Error()})
		return
	}
	log.Println(grpcResponse.Memories)
	c.JSON(http.StatusOK, grpcResponse)
}
