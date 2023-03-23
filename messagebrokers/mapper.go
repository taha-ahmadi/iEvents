package messagebrokers

type EventMapper interface {
	MapEvent(string, interface{}) (Event, error)
}

func NewEventMapper() EventMapper {
	return &StaticEventMapper{}
}
