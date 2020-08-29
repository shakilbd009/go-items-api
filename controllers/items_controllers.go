package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/shakilbd009/go-items-api/domian/items"
	"github.com/shakilbd009/go-items-api/domian/queries"
	"github.com/shakilbd009/go-items-api/services"
	"github.com/shakilbd009/go-items-api/utils/http_utils"
	"github.com/shakilbd009/go-oauth-lib/oauth"

	"github.com/shakilbd009/go-utils-lib/logger"
	"github.com/shakilbd009/go-utils-lib/rest_errors"
)

var ItemsController itemsControllerInterface = &itemsController{}

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Search(http.ResponseWriter, *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {

	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.ResponseError(w, err)
		fmt.Println("error happend1")
		return
	}
	sellerID := oauth.GetCallerID(r)
	if sellerID == 0 {
		restErr := rest_errors.NewUnathorizedError("seller id is less then 0, seller id must be more than 0")
		logger.Error(restErr.Message(), fmt.Errorf("error happend2"))
		http_utils.ResponseError(w, restErr)
		return
	}
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restErr := rest_errors.NewBadRequestError("invalid request body")
		fmt.Println("error happend3")
		http_utils.ResponseError(w, restErr)
		return
	}
	defer r.Body.Close()
	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.ResponseError(w, restErr)
		fmt.Println("error happend4")
		return
	}
	itemRequest.Seller = sellerID

	result, restErr := services.ItemsService.Create(itemRequest)
	if restErr != nil {
		http_utils.ResponseError(w, restErr)
		fmt.Println("error happend5")
		return
	}
	http_utils.RespondJSON(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http_utils.ResponseError(w, rest_errors.NewBadRequestError("only id is allowed as Query param"))
		return
	}
	result, err := services.ItemsService.Get(strings.TrimSpace(id))
	if err != nil {
		http_utils.ResponseError(w, err)
		return
	}
	http_utils.RespondJSON(w, http.StatusOK, result)
}

func (c *itemsController) Search(w http.ResponseWriter, r *http.Request) {

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.ResponseError(w, restErr)
		return
	}
	defer r.Body.Close()
	var query queries.EsQuery
	if err := json.Unmarshal(bytes, &query); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.ResponseError(w, restErr)
		return
	}
	result, searchErr := services.ItemsService.Search(query)
	if searchErr != nil {
		http_utils.ResponseError(w, searchErr)
		return
	}
	http_utils.RespondJSON(w, http.StatusOK, result)
}

func (c *itemsController) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http_utils.ResponseError(w, rest_errors.NewBadRequestError("only id is allwed as query param"))
		return
	}
	result, err := services.ItemsService.Delete(id)
	if err != nil {
		http_utils.ResponseError(w, err)
		return
	}
	http_utils.RespondJSON(w, http.StatusOK, result)
}
