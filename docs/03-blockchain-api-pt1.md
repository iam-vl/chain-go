# Blockchain Server API 1 

Plan: 
* Launching web server
* How to create Blockchain API
* Launching UI Server for Wallets 
* Creating Wallet UI
* Creating wallets 
* Sending transactions from UI

## Launching web server 

blockchain server: 
```go
type BlockchainServer struct {
	port uint16
}
func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{port}
}
func (bcs *BlockchainServer) Port() uint16 {
	return bcs.port
}
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World!")
}
func (bcs *BlockchainServer) Run() {
	http.HandleFunc("/", HelloWorld)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bcs.Port())), nil))
}
```

Main: 
```go
func init() {
	log.SetPrefix("Blockchain: ")
}
func main() {
	port := flag.Uint("port", 5000, "TCP Port for Blockchain Server")
	flag.Parse()
	fmt.Println(port)
}
```
Running: 
```go
$ go run main.go -help
Usage of ~\AppData\Local\Temp\go-build3359835188\b001\exe\main.exe:
  -port uint
        TCP Port for Blockchain Server (default 5000)
$ go run main.go -port 5001
0xc00000a0f8
```
Update main: 
```go
fmt.Println(*port)
```
Result:  
```go
$ go run main.go -port 5001
5001
```
Update main:
```go
func main() {
	port := flag.Uint("port", 5000, "TCP Port for Blockchain Server")
	flag.Parse()
	fmt.Printf("Starting server on port ", str(*port))
	app := NewBlockchainServer(uint16(*port))
	app.Run()
}
```
Run now: 
```
go run main.go blockchain_server.go -port 5001
```

## How to create Blockchain API

## Launching UI Server for Wallets 

## Creating Wallet UI

## Creating wallets 

## Sending transactions from UI


