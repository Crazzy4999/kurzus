package server

import (
	"encoding/json"
	config "hangryAPI/configs"
	handler "hangryAPI/internal/endpoint/http/handlers"
	"hangryAPI/internal/endpoint/http/responses"
	dbrepo "hangryAPI/internal/repositories/db"
	"hangryAPI/internal/router"
	"hangryAPI/internal/service/token"
	"net/http"
	"net/http/httptest"
)

func StartTest() *httptest.Server {
	cfg := config.NewConfig(".env.testing")
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
	_ = handler.NewUserHandler(ur, ar, cfg)
	_ = handler.NewAddressHandler(ar)
	_ = handler.NewDriverHandler(dr)
	_ = handler.NewSupplierHandler(sr, str)
	_ = handler.NewOrderHandler(or, osr)
	_ = handler.NewCategoryHandler(cr)
	_ = handler.NewItemHandler(ir, imr)
	_ = handler.NewMenuHandler(mr, imr)
	_ = handler.NewOrderMenuHandler(omr)

	router.GET("/profile", func(w http.ResponseWriter, r *http.Request) {
		tokenService := token.NewTokenService(cfg)
		claims, err := tokenService.GetClaims(r, cfg.AccessSecret)
		if err != nil {
			http.Error(w, handler.INVALID_CREDENTIALS, http.StatusBadRequest)
			return
		}

		resp := &responses.UserResponse{
			ID:        claims.ID,
			FirstName: "Test",
			LastName:  "Test",
			Email:     "test@test.test",
		}

		json.NewEncoder(w).Encode(resp)
	}, nil)

	return httptest.NewServer(router)
}
