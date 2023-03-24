package messagebrokers_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taha-ahmadi/iEvents/messagebrokers"
)

type TestEvent struct {
	EventId string `json:"eventId"`
	Name    string `json:"eventName"`
}

func (e *TestEvent) EventName() string {
	return "testEvent"
}

func TestDynamicEventMapper(t *testing.T) {
	mapper := messagebrokers.NewDynamicEventMapper()

	// Register event type
	err := mapper.RegisterMapping(reflect.TypeOf(TestEvent{}))
	assert.NoError(t, err)

	// Test unmarshalling from JSON bytes
	eventJSON := []byte(`{"eventId":"123","eventName":"Test Event"}`)
	event, err := mapper.MapEvent("testEvent", eventJSON)
	assert.NoError(t, err)
	testEvent, ok := event.(*TestEvent)
	assert.True(t, ok)
	assert.Equal(t, "123", testEvent.EventId)
	assert.Equal(t, "Test Event", testEvent.Name)

	// Test unknown event type
	unknownEvent, err := mapper.MapEvent("unknownEvent", nil)
	assert.Nil(t, unknownEvent)
	assert.Error(t, err)
	assert.Equal(t, "no mapping configured for event unknownEvent", err.Error())

	// Test registering invalid event type
	err = mapper.RegisterMapping(reflect.TypeOf("invalidType"))
	assert.Error(t, err)
	assert.Equal(t, "type *string does not implement the Event interface", err.Error())

	// Test decoding from a map
	eventMap := map[string]interface{}{
		"eventId":   "456",
		"eventName": "Test Event",
	}
	eventJSON, err = json.Marshal(eventMap)
	assert.NoError(t, err)
	event, err = mapper.MapEvent("testEvent", eventMap)
	assert.NoError(t, err)
	testEvent, ok = event.(*TestEvent)
	assert.True(t, ok)
	assert.Equal(t, "456", testEvent.EventId)
	assert.Equal(t, "Test Event", testEvent.Name)
}
