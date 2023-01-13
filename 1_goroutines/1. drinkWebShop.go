//see "7. Drink Web Shop Architecture.PNG"

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var orders = make(map[int]string)

func mainDrinkWebShop() {
	router := mux.NewRouter()
	router.HandleFunc("/viewOrder", viewOrderHandle).Methods("GET")
	router.HandleFunc("/addOrder/{orderID}/{orderItem}", addOrderHandle).Methods("POST")
	router.HandleFunc("/viewAllOrders", viewAllOrdersHandle).Methods("GET")
	http.ListenAndServe(":8080", router)
}

func viewAllOrdersHandle(w http.ResponseWriter, r *http.Request) {
	for id, item := range orders {
		fmt.Fprintf(w, "%d is %s\n", id, item)
	}
}

// Sent from Postman as: localhost:8080/viewOrder?orderID=11
// Using  r.URL.Query
func viewOrderHandle(w http.ResponseWriter, r *http.Request) {
	orderIDStr := r.URL.Query().Get("orderID")
	orderID, e := strconv.Atoi(orderIDStr)
	if e != nil {
		log.Printf("Can not convert to int value: %s", orderIDStr)
		return
	}
	log.Printf("View Order: %d", orderID)
	order, ok := orders[orderID]
	if ok == false {
		fmt.Fprintf(w, "%d is %s", orderID, "not exists")
	} else {
		fmt.Fprintf(w, "%d is %s", orderID, order)
	}
}

// Sent from Postman as: localhost:8080/addOrder/11/milk
// Using gorilla mux
func addOrderHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderIDStr := vars["orderID"]
	orderItemStr := vars["orderItem"]
	orderID, e := strconv.Atoi(orderIDStr)
	if e != nil {
		log.Printf("Can not convert to int value: %s", orderIDStr)
		return
	}
	log.Printf("Add Order: %d of item: %s", orderID, orderItemStr)
	orders[orderID] = orderItemStr
	fmt.Fprintf(w, "Add Order: %d of item: %s", orderID, orderItemStr)
}
