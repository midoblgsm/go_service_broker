package web_server

import (
	"fmt"
	"net/http"
)

type Controller struct {
}

func (c *Controller) Catalog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Dongdong hao bang")
}

func (c *Controller) CreateServiceInstance(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) RemoveServiceInstance(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) Bind(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) UnBind(w http.ResponseWriter, r *http.Request) {

}
