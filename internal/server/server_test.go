package server

import (
	config "hangryAPI/configs"
	handler "hangryAPI/internal/endpoint/http/handlers"
	dbrepo "hangryAPI/internal/repositories/db"
	"hangryAPI/internal/router"
	"net/http/httptest"
)

func StartTest() *httptest.Server {
	cfg := config.NewConfig("../../../configs/.env.testing")
	router := router.NewRouter()
	db := dbrepo.GetDB(cfg)
	defer db.Close()

	ur := dbrepo.NewUserRepository(db)
	ar := dbrepo.NewAddressRepository(db)
	dr := dbrepo.NewDriverRepository(db)
	sr := dbrepo.NewSupplierRepository(db)
	str := dbrepo.NewSupplierTypesRepository(db)
	or := dbrepo.NewOrderRepository(db)
	osr := dbrepo.NewOrderStatusRepository(db)
	cr := dbrepo.NewCategoriesRepository(db)
	ir := dbrepo.NewItemRepository(db)
	imr := dbrepo.NewItemsMenusRepository(db)
	mr := dbrepo.NewMenuRepository(db)
	omr := dbrepo.NewOrderMenusRepository(db)

	_ = handler.NewAuthHandler(ur, cfg)
	userHandler := handler.NewUserHandler(ur, ar, cfg)
	_ = handler.NewAddressHandler(ar)
	_ = handler.NewDriverHandler(dr)
	_ = handler.NewSupplierHandler(sr, str)
	_ = handler.NewOrderHandler(or, osr)
	_ = handler.NewCategoryHandler(cr)
	_ = handler.NewItemHandler(ir, imr)
	_ = handler.NewMenuHandler(mr, imr)
	_ = handler.NewOrderMenuHandler(omr)

	router.GET("/profile", userHandler.GetProfile, nil)

	return httptest.NewServer(router)
}
