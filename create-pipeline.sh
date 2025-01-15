#!/bin/bash

# Elasticsearch host and port information
ES_HOST="http://elasticsearch:9200"

# Define pipeline name
PIPELINE_NAME="news-transform-pipeline"

# Define pipeline configuration as JSON
PIPELINE_DEFINITION='{
  "processors": [
    {
      "script": {
        "source": "ctx.id = ctx.id; ctx.title = ctx.title; ctx.content = ctx.content; ctx.author = ctx.author; ctx.cat = ctx.created_at;"
      }
    }
  ]
}'

# Send pipeline to elasticsearch
curl -X PUT "$ES_HOST/_ingest/pipeline/$PIPELINE_NAME" \
  -H "Content-Type: application/json" \
  -d "$PIPELINE_DEFINITION"

# Check result
if [ $? -eq 0 ]; then
  echo "Pipeline $PIPELINE_NAME created successfully."
else
  echo "Pipeline $PIPELINE_NAME an error occurred while creating."
fi
