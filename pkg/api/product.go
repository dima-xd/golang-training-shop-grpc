package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"golang-training-shop-grpc/product_server/pkg/data"

	"github.com/gorilla/mux"
)

type productAPI struct {
	data *data.ProductData
}

func ServeProductResource(r *mux.Router, data data.ProductData) {
	api := &productAPI{data: &data}
	r.HandleFunc("/products", api.getAllProducts).Methods("GET")
	r.HandleFunc("/products", api.createProduct).Methods("POST")
	r.HandleFunc("/products/{id}", api.getProduct).Methods("GET")
	r.HandleFunc("/products/{id}/{unit_price}", api.updateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", api.deleteProduct).Methods("DELETE")
}

func (p productAPI) getAllProducts(writer http.ResponseWriter, request *http.Request) {
	products, err := p.data.ReadAll()
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get products"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(writer).Encode(products)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (p productAPI) createProduct(writer http.ResponseWriter, request *http.Request) {
	product := new(data.Product)
	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		log.Println("failed reading JSON: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if product == nil {
		log.Println("empty JSON")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = p.data.Create(*product)
	if err != nil {
		log.Println("product has not been created")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (p productAPI) getProduct(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := p.data.Read(id)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(writer).Encode(product)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (p productAPI) updateProduct(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	unitPrice := params["unit_price"]
	err = p.data.Update(id, unitPrice)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (p productAPI) deleteProduct(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = p.data.Delete(id)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}
