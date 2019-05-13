# Test script to POST a url and status to the lookup service
# usage "python urlpoststatus.py myurl status"

import httplib, urllib

# Setup Parser
import argparse
parser = argparse.ArgumentParser()
# First argument of string : "/url"
parser.add_argument("url")
# Second argument of string : "Status"
parser.add_argument("status")
args = parser.parse_args()
lookupurl = args.url
status = args.status


# Create a POST header
headers = {"Content-type": "application/x-www-form-urlencoded",
            "Accept": "text/plain"}

# Create a connection to URL lookup service
conn = httplib.HTTPConnection("localhost:8080")

# Sample URL status POST
params = urllib.urlencode({'url': lookupurl, 'status': status})

# Send the POST request
conn.request("POST", "", params, headers)

# Read Response and print status and data
response = conn.getresponse()
print response.status, response.reason
data = response.read()
print data
conn.close()

