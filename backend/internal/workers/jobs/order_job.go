package jobs

import "hangryAPI/internal/models"

type OrderJob struct {
	Id     int
	Order  *models.Order
	Driver *models.Driver
}
