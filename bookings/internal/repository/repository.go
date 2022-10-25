package repository

import (
	"time"

	"github.com/yazuka000/bookings/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)

	InsertRoomRestriction(r models.RoomRestriction) error

	SearchAvailabilityByDatesByRoomId(start, end time.Time, roomId int) (bool, error)

	SearchAvailabilityForAllRooms(start, end time.Time)([]models.Room, error)
}
