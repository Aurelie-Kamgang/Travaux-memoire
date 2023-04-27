#!/bin/bash

# List of URLs to send requests to
URLS="https://www.google.com/ https://cloud.google.com/apis"
# Number of requests to send
NUM_REQUESTS=10
# Time to wait between sending requests (in seconds)
WAIT_TIME=5

for URL in $URLS
 do
  for ((i=1; i<=$NUM_REQUESTS; i++))
   do
    echo "Sending request #$i to $URL"
    curl $URL
    sleep $WAIT_TIME
   done
 done
