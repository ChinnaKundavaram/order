package resource

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"order/common"

	"order/models"
	OrderService "order/service"

	websocket "order/websocket"

	"github.com/go-chassis/go-archaius"
	"github.com/go-chassis/go-chassis/v2/server/restful"
	"github.com/go-chassis/openlog"
)

const jsonHeader string = "application/json"

type OrderResource struct {
	ts   OrderService.OrderServiceInterface
	Pool *websocket.Pool
}

// Services should get injected before using.
func (ur *OrderResource) Inject(ts OrderService.OrderServiceInterface) {
	ur.ts = ts
}

// creates the order
func (r *OrderResource) CreateOrder(context *restful.Context) {
	openlog.Info("Got a request to create order")
	order := models.Order{}
	// Read the payload into order from context
	err := context.ReadEntity(&order)
	if err != nil {
		openlog.Error(err.Error())
		// Send error response
		context.WriteHeaderAndJSON(http.StatusBadRequest, common.HTTPResponse{Status: 400, Msg: "Could not read paylaod"}, "application/json")
		return
	}
	fmt.Println(order)
	data := r.ts.CreateOrder(&order)
	context.WriteHeaderAndJSON(data.Status, data, jsonHeader)
}

func (r *OrderResource) UpdateOrder(context *restful.Context) {

	openlog.Info("Got a request to update order")
	order := make(map[string]interface{})

	//Read the payload into order from context
	err := context.ReadEntity(&order)
	if err != nil {
		openlog.Error(err.Error())
		// Send error response
		context.WriteHeaderAndJSON(http.StatusBadRequest, common.HTTPResponse{Status: 400, Msg: "Could not read paylaod"}, "application/json")
		return
	}
	Id := context.ReadPathParameter("id")
	res := r.ts.UpdateOrder(Id, order)
	context.WriteHeaderAndJSON(res.Status, res, "application/json")

}

func (r *OrderResource) DeleteOrder(context *restful.Context) {
	openlog.Info("got a request to delete order")

	Id := context.ReadPathParameter("id")
	res := r.ts.DeleteOrder(Id)
	context.WriteHeaderAndJSON(res.Status, res, "application/json")
}

func (r *OrderResource) FetchAll(context *restful.Context) {
	openlog.Info("Got a request to fetch all orders")
	pageno, size := common.GetPageDetails(context.ReadQueryParameter("page"), context.ReadQueryParameter("size"))
	OrderId, err := strconv.ParseInt(context.ReadPathParameter("orderId"), 10, 32)
	if err != nil {
		openlog.Error("Error occured while converitng the order ID")
		context.WriteHeaderAndJSON(400, common.HTTPResponse{Status: 400, Msg: "Invalid order Id"}, jsonHeader)
		return
	}
	//filters := make(map[string]interface{})
	// Need to add filters here
	data := r.ts.FetchAll(int(OrderId), pageno, size, nil)
	context.WriteHeaderAndJSON(data.Status, data, jsonHeader)
}

func (r *OrderResource) VersionInfo(context *restful.Context) {
	openlog.Info("Executing version info")
	dummy_res := struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}{
		Name:    "Template Service",
		Version: "0.0.1",
	}
	context.WriteJSON(dummy_res, "application/json")
}

func (r *OrderResource) WebsocketHandler(context *restful.Context) {
	openlog.Debug("connecting to web socket")
	conn, err := websocket.Upgrade(context.ReadResponseWriter(), context.Req.Request)
	if err != nil {
		openlog.Error("connection error")
		return
	}
	client := &websocket.Client{
		Conn:     conn,
		Pool:     r.Pool,
		GetData:  getData,
		Interval: 1,
	}
	r.Pool.Register <- client

}

// Define all APIs here.
func (r *OrderResource) URLPatterns() []restful.Route {
	// r.Pool = websocket.NewPool(3) // takes max connectionsconst jsonHeader string = "application/json"const jsonHeader string = "application/json"
	// go r.Pool.Start()
	baseurl := archaius.GetString("servicecomb.service.baseurl", "/api/V1.0")
	return []restful.Route{
		{Method: http.MethodGet, Path: "/info", ResourceFunc: r.VersionInfo},
		{Method: http.MethodGet, Path: "/websocket", ResourceFunc: r.WebsocketHandler},
		{Method: http.MethodPost, Path: baseurl + "/users", ResourceFunc: r.CreateOrder, Consumes: []string{"application/json"}, Produces: []string{"application/json"}},
		{Method: http.MethodPut, Path: baseurl + "/users/{id}", ResourceFunc: r.UpdateOrder, Consumes: []string{"application/json"}, Produces: []string{"application/json"}},
		{Method: http.MethodDelete, Path: baseurl + "/users/{id}", ResourceFunc: r.DeleteOrder, Consumes: []string{"application/json"}, Produces: []string{"application/json"}},
		{Method: http.MethodGet, Path: baseurl + "/users/{orderId}", ResourceFunc: r.FetchAll, Consumes: []string{"application/json"}, Produces: []string{"application/json"}},
	}
}
func getData(ctx *context.Context) []byte {
	current_time := time.Now()
	result := "time currenlty is " + current_time.String()
	return []byte(result)
}
