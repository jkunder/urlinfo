# urlinfo
## url lookup service
 
###Description
URL Lookup Service, provides a URL lookup service to a http proxy 
by looking up a database for URL reputation.

### URL database updates
Database updates are done through a HTTP POST request
POST {'url': url, 'status': 'ALLOW|BLOCK'}

THe status could be "ALLOW" or "BLOCK"


## URL Lookup
URL Lookups are achieved through a HTTP GET request
GET /urlinfo/1/{hostname_and_port}/{original_path_and_query_string}
 
Response of the Request is 
Status : 200 OK
Body : "URL {url} status {"ALLOW"|"DENY"|"UNKNOWN"}


## Source Code and Build
The Source code is under ws/src/github.com/urllookup
File : urllookup.go : Golang code to run HTTP server, Process GET and POST
requests and accordingly read/update a MAP datastructure.

Build : go build
Run : ./urllookup : Starts a http server listening on port 8080

## Test Scripts and steps to test
The tests scripts are under ws/test
Files : 

urllookup.py : Python script to lookup an URL
Ex: pyuthon urllookup.py cnn.com

urlpoststatus.py : Python script to Post status of an URL
Ex: python urlpoststatus.py cnn.com BLOCK


