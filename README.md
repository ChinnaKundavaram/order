# order
commands for run............. 
cd server                              [enters into server file]
$env:CHASSIS_HOME="$PWD"               [export]
go run main.go                         [run]


Payload for create or update an order.......

{
    "id":int
    "status":string
    "currencyunit":string
}
