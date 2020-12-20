package sensor

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gofrs/uuid"
)

type sensor struct {
	id          string
	name        string
	temperature float32
	humidity    float32
	battery     float32
	timestamp   string
}

func NewSensor(name string) *sensor {
	return &sensor{
		id:   uuid.Must(uuid.NewV4()).String(),
		name: name,
	}
}

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func (s *sensor) GetData() string {
	return fmt.Sprintf(`{ "id": %s, "name": %s, "temperature": %f, "humidity": %f, "battery": 95, "timestamp": %s }`,
		s.id, s.name, randFloat(21, 24), randFloat(40, 70), time.Now().Format(time.RFC3339))
}
