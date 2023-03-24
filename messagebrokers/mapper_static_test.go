package messagebrokers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taha-ahmadi/iEvents/contracts"
	"github.com/taha-ahmadi/iEvents/messagebrokers"
)

func TestStaticEventMapper_MapEvent(t *testing.T) {
	mapper := &messagebrokers.StaticEventMapper{}

	// Test unmarshalling from JSON bytes
	eventJSON := []byte(`{"id":"123","name":"Test Event"}`)
	event, err := mapper.MapEvent("eventCreated", eventJSON)
	assert.NoError(t, err)
	createdEvent, ok := event.(*contracts.EventCreatedEvent)
	assert.True(t, ok)
	assert.Equal(t, "123", createdEvent.ID)
	assert.Equal(t, "Test Event", createdEvent.Name)

	// Test decoding from a map
	locationMap := map[string]interface{}{
		"id":   "456",
		"name": "Test Location",
	}
	location, err := mapper.MapEvent("locationCreated", locationMap)
	assert.NoError(t, err)
	locationCreatedEvent, ok := location.(*contracts.LocationCreatedEvent)
	assert.True(t, ok)
	assert.Equal(t, "456", locationCreatedEvent.ID)
	assert.Equal(t, "Test Location", locationCreatedEvent.Name)

	// Test unknown event type
	unknownEvent, err := mapper.MapEvent("unknownEvent", nil)
	assert.Nil(t, unknownEvent)
	assert.Error(t, err)
	assert.Equal(t, "unknown event type unknownEvent", err.Error())
}
