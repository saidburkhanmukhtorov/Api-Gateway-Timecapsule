package handlers

import (
	"google.golang.org/grpc"

	"github.com/time_capsule/Api-Gateway-Timecapsule/config"
	"github.com/time_capsule/Api-Gateway-Timecapsule/kafka"
)

// Handler struct holds all the individual entity handlers.
type Handler struct {
	// Timeline service handlers.
	MilestoneHandler       *MilestoneHandler
	CustomEventHandler     *CustomEventHandler
	HistoricalEventHandler *HistoricalEventHandler

	// Memory service handlers.
	MemoryHandler  *MemoryHandler
	CommentHandler *CommentHandler
	MediaHandler   *MediaHandler
}

// It accepts gRPC connections and initializes Kafka producers within each handler.
func NewHandler(timelineGrpcConn, memoryGrpcConn *grpc.ClientConn, cfg *config.Config) *Handler {
	// Create Kafka producers
	milestoneProducer := kafka.NewProducer(cfg.KafkaMilestoneTopic)
	customEventProducer := kafka.NewProducer(cfg.KafkaCustomEventTopic)
	historicalEventProducer := kafka.NewProducer(cfg.KafkaHistoricalEventTopic)
	memoryProducer := kafka.NewProducer(cfg.KafkaMemoryTopic)
	commentProducer := kafka.NewProducer(cfg.KafkaCommentTopic)
	mediaProducer := kafka.NewProducer(cfg.KafkaMediaTopic)

	return &Handler{
		// Timeline service handlers.
		MilestoneHandler:       NewMilestoneHandler(milestoneProducer, timelineGrpcConn),
		CustomEventHandler:     NewCustomEventHandler(customEventProducer, timelineGrpcConn),
		HistoricalEventHandler: NewHistoricalEventHandler(historicalEventProducer, timelineGrpcConn),

		// Memory service handlers.
		MemoryHandler:  NewMemoryHandler(memoryProducer, memoryGrpcConn),
		CommentHandler: NewCommentHandler(commentProducer, memoryGrpcConn),
		MediaHandler:   NewMediaHandler(mediaProducer, memoryGrpcConn),
	}
}
