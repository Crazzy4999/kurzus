package workers

import (
	"hangryAPI/internal/models"
	"hangryAPI/internal/repositories/db"
	"hangryAPI/internal/util"
	"hangryAPI/internal/workers/jobs"
	"log"
	"sync"
	"time"
)

type OrderDriverWorkerPool struct {
	orderRepo  *db.OrderRepository
	driverRepo *db.DriverRepository
}

func NewOrderDriverWorkerPool(orderRepo *db.OrderRepository, driverRepo *db.DriverRepository) *OrderDriverWorkerPool {
	return &OrderDriverWorkerPool{
		orderRepo:  orderRepo,
		driverRepo: driverRepo,
	}
}

var orderJobs = make(chan jobs.OrderJob, 100)
var orderResults = make(chan jobs.OrderJob, 100)

func (w *OrderDriverWorkerPool) assignDriver(job jobs.OrderJob) *jobs.OrderJob {
	driver := &models.Driver{
		ID:           job.Driver.ID,
		IsDelivering: true,
		FirstName:    job.Driver.FirstName,
		LastName:     job.Driver.LastName,
		Password:     job.Driver.Password,
	}

	err := w.driverRepo.Update(driver)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	order := &models.Order{
		ID:          job.Order.ID,
		DriverID:    util.NullInt64(int64(job.Driver.ID)),
		StatusID:    models.DELIVERING,
		Note:        job.Order.Note,
		DeliveredAt: job.Order.DeliveredAt,
	}

	err = w.orderRepo.Update(order)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	job.Order = order
	job.Driver = driver

	return &job
}

func (w *OrderDriverWorkerPool) allocate(jobsNum int) {
	orders, err := w.orderRepo.GetOrdersWithCreatedStatus()
	if err != nil {
		log.Fatal(err)
		return
	}

	drivers, err := w.driverRepo.GetAllNotDelivering()
	if err != nil {
		log.Fatal(err)
		return
	}

	for i := 0; i <= jobsNum; i++ {
		if 0 < len(orders) && 0 < len(drivers) && i < len(orders) && i < len(drivers) {
			job := jobs.OrderJob{Id: i, Order: orders[i], Driver: drivers[i]}
			orderJobs <- job
		}
	}

	close(orderJobs)
}

func (w *OrderDriverWorkerPool) worker(wg *sync.WaitGroup) {
	for job := range orderJobs {
		orderResults <- *w.assignDriver(job)
	}
	wg.Done()
}

func (w *OrderDriverWorkerPool) createWorkerPool(workersNum int) {
	wg := sync.WaitGroup{}
	wg.Add(workersNum)
	for i := 0; i < workersNum; i++ {
		go w.worker(&wg)
	}
	wg.Wait()
}

func (w *OrderDriverWorkerPool) Start(min int) {
	for {
		orderJobs = make(chan jobs.OrderJob, 100)
		orderResults = make(chan jobs.OrderJob, 100)
		<-time.After(time.Second * time.Duration(min)) //TODO: CHANGE BACK TO MINTUE
		w.allocate(100)
		w.createWorkerPool(10)
	}
}
