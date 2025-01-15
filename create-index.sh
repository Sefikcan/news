#!/bin/bash

# Connect to elasticsearch
HOST="http://elasticsearch:9200"
INDEX_NAME="news"

# If index doesn't exist, create it
if ! curl -s -o /dev/null -w "%{http_code}" -X GET "$HOST/$INDEX_NAME" | grep -q "200"; then
  echo "Index '$INDEX_NAME' not found, being created..."
  curl -X PUT "$HOST/$INDEX_NAME" -H 'Content-Type: application/json' -d'
  {
    "settings": {
      "number_of_shards": 1,
      "number_of_replicas": 1
    }
  }
  '
  echo "Index '$INDEX_NAME' created successfully."
else
  echo "Index '$INDEX_NAME' already exists."
fi
