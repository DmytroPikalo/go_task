## how to run
You must first download the modules you do not have(Golang version 1.20)
<code>
go mod tidy
</code>
After that, you can start the server
<code>
go run main.go
</code>

## Documentation
### The Project has two structures to help ease the work
JsonReader handles requests from users 
<code>
type JsonReader struct {
	IpAddress string   `json:"ip_address"`
	Add       []string `json:"add"`
	Remove    []string `json:"remove"`
}
</code>

For the refresh function, we use the fields 
<code>
	Add       []string `json:"add"`
	Remove    []string `json:"remove"`
</code>

For the сontains function, we use the fields 
<code>
	IpAddress string   `json:"ip_address"`
</code>

MemoryData we use to store data in our memory.
<br>sync.RWMutex is needed because map is not "reliable" when using goroutines.<br/>
<br>the Url and Ips fields immediately acquire values using the GetConf and LinesToMap functions, respectively.<br/>
The Ips field may change during program execution.
<code>
type MemoryData struct {
	sync.RWMutex
	URL     string `yaml:"url"`
	Ips     map[string]bool
	Changes map[string][]string
}
</code>

### Project structure

all the main logic on the project is in the **helpers/** directory<br>
handling user requests in the **api/** directory<br>
for constant values/environment values **config/** directory

### Test run
To run all tests
<code>
go test test_task/helpers
</code>

To run all tests
<code>
go test test_task/helpers
</code>

To run a single test
<code>
go test -run ^TestValidateIp$ test_task/helpers
</code>

### Test endpoints Postman
If you using Postman you can import **end_points.postman_collection.json**

### Test endpoints curl
Example
<code>
curl http://localhost:3000/contains -d '{"ip_address": "1.1.1.1"}'
</code>
