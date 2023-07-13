package workers

import (
	"hangryAPI/internal/models"
	"hangryAPI/internal/repositories/db"
	"hangryAPI/internal/workers/jobs"
	"log"
	"sync"
	"time"
)

type OrderDeliveryWorkerPool struct {
	orderRepo  *db.OrderRepository
	driverRepo *db.DriverRepository
}

func NewOrderDeliveryWorkerPool(orderRepo *db.OrderRepository, driverRepo *db.DriverRepository) *OrderDeliveryWorkerPool {
	return &OrderDeliveryWorkerPool{
		orderRepo:  orderRepo,
		driverRepo: driverRepo,
	}
}

var deliveryJobs = make(chan jobs.DeliveryJob, 100)
var deliveryResults = make(chan jobs.DeliveryJob, 100)

func (w *OrderDeliveryWorkerPool) finishDelivery(job jobs.DeliveryJob) *jobs.DeliveryJob {
	order := &models.Order{
		ID:                job.Order.ID,
		DriverID:          job.Order.DriverID,
		StatusID:          models.DONE,
		Note:              job.Order.Note,
		EstimatedDelivery: job.Order.EstimatedDelivery,
		DeliveredAt:       job.Order.DeliveredAt,
	}

	date, err := time.Parse("2006-01-02T15:04:05.999999Z07:00", order.EstimatedDelivery)
	if err != nil {
		log.Fatal("Couldn't convert date string to time type!")
		return nil
	}

	if date.Unix() < time.Now().Unix() {
		driver, err := w.driverRepo.GetDriverByID(int(order.DriverID.Int64))
		if err != nil {
			log.Fatal(err)
			return nil
		}

		finishedDriver := &models.Driver{
			ID:           driver.ID,
			IsDelivering: false,
			FirstName:    driver.FirstName,
			LastName:     driver.LastName,
			Password:     driver.Password,
		}

		err = w.driverRepo.Update(finishedDriver)
		if err != nil {
			log.Fatal(err)
			return nil
		}

		order.DeliveredAt.String = time.Now().Format("2006-01-02T15:04:05.999999Z07:00")

		err = w.orderRepo.Update(order)
		if err != nil {
			log.Fatal(err)
			return nil
		}

		job.Order = order
	}

	return &job
}

func (w *OrderDeliveryWorkerPool) allocate(jobsNum int) {
	orders, err := w.orderRepo.GetOrdersWithDeliveringStatus()
	if err != nil {
		log.Fatal(err)
		return
	}

	for i := 0; i <= jobsNum; i++ {
		if 0 < len(orders) && i < len(orders) {
			job := jobs.DeliveryJob{Id: i, Order: orders[i]}
			deliveryJobs <- job
		}
	}

	close(deliveryJobs)
}

func (w *OrderDeliveryWorkerPool) worker(wg *sync.WaitGroup) {
	for job := range deliveryJobs {
		deliveryResults <- *w.finishDelivery(job)
	}
	wg.Done()
}

func (w *OrderDeliveryWorkerPool) createWorkerPool(workersNum int) {
	wg := sync.WaitGroup{}
	wg.Add(workersNum)
	for i := 0; i < workersNum; i++ {
		go w.worker(&wg)
	}
	wg.Wait()
}

func (w *OrderDeliveryWorkerPool) Start(min int) {
	for {
		deliveryJobs = make(chan jobs.DeliveryJob, 100)
		deliveryResults = make(chan jobs.DeliveryJob, 100)
		<-time.After(time.Second * time.Duration(min)) //TODO: CHANGE BACK TO MINTUE
		w.allocate(100)
		w.createWorkerPool(10)
	}
}
