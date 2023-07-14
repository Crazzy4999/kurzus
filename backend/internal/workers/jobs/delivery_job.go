package jobs

import "hangryAPI/internal/models"

type DeliveryJob struct {
	Id    int
	Order *models.Order
}
