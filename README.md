# order
commands for run 
cd server                              [enters into server file]
$env:CHASSIS_HOME="$PWD"               [export]
go run main.go                         [run]


Payload for create or update an order
{
    "id":          int
    "status":      string
    "currency":    string
    "item":        string 
	"itemid":      int   
	"description": string 
	"price":       string 
	"quantity":    string 
	"total":       string
}

// service comb(Apache) need to be running mode to perform operations