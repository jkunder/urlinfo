# urlinfo
## url lookup service
 
### Description
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
Body : "url : {url}, status : {"ALLOW"|"DENY"|"UNKNOWN"}  
       Ex: {url : cnn.com, status : ALLOW}


## Source Code and Containerized Build
The Source code is under ws/src/github.com/urllookup
File : urllookup.go :  
Golang code to run HTTP server, Process GET and POST requests and accordingly read/update a redis database

Build : ./build.sh : Output is a container urllookup:latest

## Test Scripts
The tests scripts are under ws/test
Files : 

urllookup.py : Python script to lookup an URL  
Ex: python urllookup.py cnn.com

urlpoststatus.py : Python script to Post status of an URL  
Ex: python urlpoststatus.py cnn.com BLOCK

## Bringing up the Test setup
under ws/test directory

```$xslt
./urllookup.sh start
```

This brings up the REDIS and urllookup containers

```$xslt
python urlpoststatus.py cnn.com BLOCK
  URL status updated = map[url:[cnn.com] status:[BLOCK]]

python urllookup.py cnn.com
  URL : /cnn.com Status : BLOCK
```

To clean up the containers
```$xslt
./urllookup.sh stop
```

## Performance Scaling Enhancements (PENDING)
1. Make the http server multi-threaded to handle requests parallely
2. Update the test scripts to open parallel connections and scale upto 1000 requests/sec
3. Run the containers in a Kubernetes environment for cluster and scaling




