# Test script to lookup the status of a sample URL
# Usage : python urllookup.py myurl

import httplib

# Setup Parser
import argparse
parser = argparse.ArgumentParser()
# One argument of string : "/url"
parser.add_argument("url")
args = parser.parse_args()
lookupurl = '/'+args.url

# Setup Connection to lookup service
conn = httplib.HTTPConnection("localhost:8080")

# GET request for a sample URL
conn.request("GET", lookupurl)

# Print the Status and Response
r = conn.getresponse()
print r.status
data=r.read()
print data
