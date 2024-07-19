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

// CommentHandler handles requests related to Comments.
type CommentHandler struct {
	kafkaProducer *kafka.Producer
	memoryService memory.CommentServiceClient // gRPC client for Memory Service
}

// NewCommentHandler creates a new CommentHandler.
func NewCommentHandler(kafkaProducer *kafka.Producer, grpcConn *grpc.ClientConn) *CommentHandler {
	return &CommentHandler{
		kafkaProducer: kafkaProducer,
		memoryService: memory.NewCommentServiceClient(grpcConn), // Initialize gRPC client
	}
}

// CreateComment godoc
// @Summary     Create a new Comment
// @Description Create a new comment with the provided details.
// @Tags        Comments
// @Accept      json
// @Produce     json
// @Param       comment body     models.CreateCommentModel true "Comment details"
// @Security    ApiKeyAuth
// @Success     202     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/comments [post]
func (h *CommentHandler) CreateComment(c *gin.Context) {
	var createComment models.CreateCommentModel
	if err := c.ShouldBindJSON(&createComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Publish to Kafka
	if err := h.kafkaProducer.ProduceMessage(c.Request.Context(), "comment.create", createComment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment: " + err.Error()}) // Log the specific error
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Comment creation request accepted"})
}

// GetComment godoc
// @Summary     Get Comment by ID
// @Description Get details of a comment by its ID.
// @Tags        Comments
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Comment ID"
// @Security    ApiKeyAuth
// @Success     200     {object} memory.Comment
// @Failure     400     {object} map[string]interface{}
// @Failure     404     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/comments/{id} [get]
func (h *CommentHandler) GetComment(c *gin.Context) {
	commentID := c.Param("id")

	// Use gRPC to get the comment from the Memory Service
	grpcResponse, err := h.memoryService.GetCommentById(context.Background(), &memory.GetCommentByIdRequest{Id: commentID})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()}) // Return the gRPC error message
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comment: " + err.Error()}) // Log the specific error
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}

// UpdateComment godoc
// @Summary     Update Comment
// @Description Update an existing comment.
// @Tags        Comments
// @Accept      json
// @Produce     json
// @Param       id     path     string                  true "Comment ID"
// @Param       comment body     models.UpdateCommentModel true "Updated comment data"
// @Security    ApiKeyAuth
// @Success     202     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/comments/{id} [put]
func (h *CommentHandler) UpdateComment(c *gin.Context) {
	commentID := c.Param("id")
	var updateComment models.UpdateCommentModel
	if err := c.ShouldBindJSON(&updateComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Ensure the ID in the URL matches the ID in the payload
	if updateComment.ID != commentID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID mismatch"})
		return
	}

	// Publish to Kafka
	if err := h.kafkaProducer.ProduceMessage(c.Request.Context(), "comment.update", updateComment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update comment: " + err.Error()}) // Log the specific error
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Comment update request accepted"})
}

// PatchComment godoc
// @Summary     Partially update Comment
// @Description Partially update an existing comment.
// @Tags        Comments
// @Accept      json
// @Produce     json
// @Param       id     path     string                  true "Comment ID"
// @Param       comment body     models.PatchCommentModel true "Patched comment data"
// @Security    ApiKeyAuth
// @Success     202     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/comments/{id} [patch]
func (h *CommentHandler) PatchComment(c *gin.Context) {
	commentID := c.Param("id")
	var patchComment models.PatchCommentModel
	if err := c.ShouldBindJSON(&patchComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// You might want to add the ID to the patchModel for clarity:
	patchComment.ID = commentID

	// Publish to Kafka
	if err := h.kafkaProducer.ProduceMessage(c.Request.Context(), "comment.patch", patchComment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to patch comment: " + err.Error()}) // Log the specific error
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Comment patch request accepted"})
}

// DeleteComment godoc
// @Summary     Delete Comment
// @Description Delete a comment by its ID.
// @Tags        Comments
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Comment ID"
// @Security    ApiKeyAuth
// @Success     202     {object} map[string]interface{}
// @Failure     400     {object} map[string]interface{}
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/comments/{id} [delete]
func (h *CommentHandler) DeleteComment(c *gin.Context) {
	commentID := c.Param("id")

	// Publish to Kafka
	if err := h.kafkaProducer.ProduceMessage(c.Request.Context(), "comment.delete", map[string]string{"id": commentID}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment: " + err.Error()}) // Log the specific error
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Comment delete request accepted"})
}

// ListComments godoc
// @Summary     List Comments
// @Description Get a list of comments.
// @Tags        Comments
// @Accept      json
// @Produce     json
// @Param        page     query    int    false  "Page number"                minimum(1)
// @Param        limit    query    int    false  "Number of items per page"  minimum(1)
// @Param        memory_id query    string false  "Filter by memory ID"
// @Param        user_id  query    string false  "Filter by user ID"
// @Param        content  query    string false  "Filter by content"
// @Security    ApiKeyAuth
// @Success     200     {object} memory.GetAllCommentsResponse
// @Failure     500     {object} map[string]interface{}
// @Router      /v1/comments [get]
func (h *CommentHandler) ListComments(c *gin.Context) {
	// Get query parameters for pagination and filtering
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	memoryID := c.Query("memory_id")
	userID := c.Query("user_id")
	content := c.Query("content")

	// Use gRPC to get the comments from the Memory Service
	grpcResponse, err := h.memoryService.GetAllComments(context.Background(), &memory.GetAllCommentsRequest{
		Page:     helper.StringToInt(page),
		Limit:    helper.StringToInt(limit),
		MemoryId: memoryID,
		UserId:   userID,
		Content:  content,
	})
	if err != nil {
		log.Println("Failed to get comments:", err) // Log the specific error from the gRPC call
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comments"})
		return
	}

	c.JSON(http.StatusOK, grpcResponse)
}
