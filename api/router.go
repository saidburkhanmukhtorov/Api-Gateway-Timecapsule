package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/time_capsule/Api-Gateway-Timecapsule/api/docs"
	"github.com/time_capsule/Api-Gateway-Timecapsule/api/handlers"
	"github.com/time_capsule/Api-Gateway-Timecapsule/config"
	"google.golang.org/grpc"
)

func NewRouter(timelineGrpcConn, memoryGrpcConn *grpc.ClientConn) *gin.Engine {
	cfg := config.Load()
	router := gin.Default()

	handler := handlers.NewHandler(timelineGrpcConn, memoryGrpcConn, &cfg)

	// Swagger documentation
	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API versioning
	v1 := router.Group("/v1")
	{
		// Milestone routes
		milestones := v1.Group("/milestones")
		{
			milestones.POST("", handler.MilestoneHandler.CreateMilestone)
			milestones.GET(":id", handler.MilestoneHandler.GetMilestone)
			milestones.PUT(":id", handler.MilestoneHandler.UpdateMilestone)
			milestones.PATCH(":id", handler.MilestoneHandler.PatchMilestone)
			milestones.DELETE(":id", handler.MilestoneHandler.DeleteMilestone)
			milestones.GET("", handler.MilestoneHandler.ListMilestones)
		}

		// Memory routes
		memories := v1.Group("/memories")
		{
			memories.POST("", handler.MemoryHandler.CreateMemory)
			memories.GET(":id", handler.MemoryHandler.GetMemory)
			memories.PUT(":id", handler.MemoryHandler.UpdateMemory)
			memories.PATCH(":id", handler.MemoryHandler.PatchMemory)
			memories.DELETE(":id", handler.MemoryHandler.DeleteMemory)
			memories.GET("", handler.MemoryHandler.ListMemories)
		}

		// Comment routes
		comments := v1.Group("/comments")
		{
			comments.POST("", handler.CommentHandler.CreateComment)
			comments.GET(":id", handler.CommentHandler.GetComment)
			comments.PUT(":id", handler.CommentHandler.UpdateComment)
			comments.PATCH(":id", handler.CommentHandler.PatchComment)
			comments.DELETE(":id", handler.CommentHandler.DeleteComment)
			comments.GET("", handler.CommentHandler.ListComments)
		}

		// Custom Event routes
		customEvents := v1.Group("/custom-events")
		{
			customEvents.POST("", handler.CustomEventHandler.CreateCustomEvent)
			customEvents.GET(":id", handler.CustomEventHandler.GetCustomEvent)
			customEvents.PUT(":id", handler.CustomEventHandler.UpdateCustomEvent)
			customEvents.PATCH(":id", handler.CustomEventHandler.PatchCustomEvent)
			customEvents.DELETE(":id", handler.CustomEventHandler.DeleteCustomEvent)
			customEvents.GET("", handler.CustomEventHandler.ListCustomEvents)
		}

		// Historical Event routes
		historicalEvents := v1.Group("/historical-events")
		{
			historicalEvents.POST("", handler.HistoricalEventHandler.CreateHistoricalEvent)
			historicalEvents.GET(":id", handler.HistoricalEventHandler.GetHistoricalEvent)
			historicalEvents.PUT(":id", handler.HistoricalEventHandler.UpdateHistoricalEvent)
			historicalEvents.PATCH(":id", handler.HistoricalEventHandler.PatchHistoricalEvent)
			historicalEvents.DELETE(":id", handler.HistoricalEventHandler.DeleteHistoricalEvent)
			historicalEvents.GET("", handler.HistoricalEventHandler.ListHistoricalEvents)
		}

		// Media routes
		media := v1.Group("/media")
		{
			media.POST("", handler.MediaHandler.CreateMedia)
			media.GET(":id", handler.MediaHandler.GetMedia)
			media.PUT(":id", handler.MediaHandler.UpdateMedia)
			media.PATCH(":id", handler.MediaHandler.PatchMedia)
			media.DELETE(":id", handler.MediaHandler.DeleteMedia)
			media.GET("", handler.MediaHandler.ListMedia)
		}
	}

	return router
}
