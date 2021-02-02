package mock

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type Clock struct {
	mock.Mock
}

func (m *Clock) Now() time.Time {
	args := m.Called()

	return args.Get(0).(time.Time)
}
