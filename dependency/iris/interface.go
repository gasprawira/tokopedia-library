package iris

import (
	"context"

	"cloud.google.com/go/pubsub"
)

type LoggerWrapper interface {
	AddSchemaWithPubsub(eventName string, schema string, config ...pubsub.PublishSettings) error
	Log(eventName string, messageData Data) error
	LogWithContext(ctx context.Context, eventName string, messageData Data) error
	LogWithPubsub(eventName string, messageData Data) error
}
