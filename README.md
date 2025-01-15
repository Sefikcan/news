CBES is a tool that enables real-time data synchronization between Couchbase and Elasticsearch. Changes in Couchbase are listened to using a CDC (Change Data Capture) approach and transferred to Elasticsearch.

# How CBES Works?

* **Change Streams**: All CRUD operations in Couchbase are captured using a DCP (Data Change Protocol).
* **DCP**: Couchbase sends changes to CBES via DCP, which then processes the data and sends it to Elasticsearch.
* **Data Transfer to Elasticsearch**: CBES sends the data to an index in Elasticsearch.
* **Using the Bulk API**: CBES uses the BULK API to send data to Elasticsearch. This improves efficiency and reduces latency.

# CBES Features

* **Real-time replication**: Data is transferred and synchronized to Elasticsearch almost in real-time.
* **Matching Data Models**: Converts the JSON from Couchbase into a format that matches Elasticsearch.
* **Supported Operations**: Insert, update, and delete operations are reflected in Elasticsearch.
* **Bulk Processing**: CBES is optimized for sending data in bulk to Elasticsearch.
* **Mapping**: CBES supports mapping configuration to define and customize how data from Couchbase is stored in Elasticsearch.
* **Flexibility**: One or more Couchbase buckets can be connected to one or more Elasticsearch indices.

# CBES Configuration
CBES configuration is done using a connector.yml file. An example configuration structure might look like this:

```toml
couchbase:
  connectionString: couchbase://localhost
  username: admin
  password: password
  bucket: news  # Couchbase bucket name

elasticsearch:
  hosts:
    - http://localhost:9200
  bulkRequest:
    maxActions: 2000  # Maximum number of documents in a bulk operation
    maxSize: 5mb      # Maximum size of data in a bulk operation
    flushInterval: 1s # Frequency of data transfer to Elasticsearch
```

# Advantages
* Easy integration: Data flow between Couchbase and Elasticsearch can be achieved without complex setups.
* High performance with bulk operations and asynchronous data transfer.
* Both Couchbase and Elasticsearch systems can be scaled independently.
* CBES organizes the data flow to accommodate these scalability requirements.

# Limitations
* Data flow is one-way (Couchbase -> Elasticsearch).
* CBES is compatible only with certain versions of Elasticsearch.
* CBES uses Couchbase's DCP feature. Since it relies on this database source, high DCP flow can impact system performance.

# What is DCP (Data Change Protocol)?
DCP is a protocol used by Couchbase that provides a mechanism to notify application layers of changes made in the database. This allows for dynamic and low-latency transmission of data from Couchbase.

# How DCP Works?
DCP provides a special channel for transmitting Couchbase changes to other systems. The mechanism works as follows:
* Couchbase logs every record. DCP uses these logs to track all changes in the database.
* DCP creates a stream for these changes, containing create, update, and delete operations, along with some metadata (e.g., timestamps).
* The relevant connectors receive this data.
* Finally, consumers receive and process the data.

# CBES Configuration File
* CBES configurations are stored in a YAML file, where the relevant settings can be made.
* An example YAML configuration file is as follows:
``` yml
couchbase:
  host: "localhost" # Couchbase host address
  port: 8091        # Couchbase port address
  user: "Administrator" # Couchbase user name
  password: "password" # Couchbase password
  bucket: "news_bucket" # Couchbase bucket name
  collections: ["xxx"] # List of collections to process
  index: "news_index" # Elasticsearch index name
  ssl:
    enabled: true # Enable SSL/TLS for secure connections

elasticsearch:
  host: "localhost"
  port: 9200
  index: "news_index"
  bulk_flush_max_actions: 5000 # Maximum data count for bulk operations
  bulk_flush_max_size_mb: 15 # Maximum memory size (MB) for bulk operations
  bulk_flush_interval_ms: 1000 # Maximum time interval (ms) for bulk operations

connector:
  log_level: "INFO" # Log level 
  metrics:
    enabled: true # Enables performance metrics for tools like Prometheus.
    port: 31415 # Metrics port
  event_filter:
    ignore_deleted: true # Determines whether deleted data from Couchbase should be sent to Elasticsearch.
    ignore_expired: true # Determines whether expired documents should be processed.
  document_routes:
    - matcher: "type == 'news'"  # Condition!
      index: "news-index" # Target Elasticsearch index name
  pipeline: # Defines the sequence of operations for data transfer from Couchbase to Elasticsearch. Multiple operations can be defined. Each operation describes data transfer between source and target.
    - name: "document-pipeline" # Pipeline name
      operations: # Defines operations to be performed.
        - type: "create" # Operation type
          source: "couchbase" # Source
          destination: "elasticsearch" # Target
```

# Additional Configuration and Settings

* CBES also includes the following configurations:
  * **data_type_mapping**: Settings for transforming data into a format suitable for Elasticsearch. For example, converting date formats or handling custom data types.
Controlling Data Flow from Couchbase (DCP)

``` yml
dcp:
  enabled: true # Enables or disables DCP
  flow_control_buffer: 128 # Memory size used during data transfer (MB)
  persistence_polling_interval: 100 # Time interval (ms) to check if data in Couchbase is persistent
  checkpoint_interval: 100 # Used to create checkpoints at specified intervals for DCP.
  start_from: "beginning" # Defines the starting point for DCP. "beginning" means starting from all current data in the database, "now" means starting from new changes.
```

# SSL/TLS Settings

``` yml
ssl:
  enabled: true # Enables SSL
  keystore: "/path/to/keystore" # Path to certificate file
  keystore_password: "password" # Password for keystore
  truststore: "/path/to/truststore" # Path to truststore file
  truststore_password: "password" # Truststore password
```

# Error and Retry Settings
``` yml
error_tolerance:
  max_retries: 5 # Maximum number of retries in case of errors
  retry_backoff_ms: 500 # Time interval (ms) between retries
```

# Advanced Settings
``` yml
advanced:
  batch_size: 1000 # Maximum number of documents to process per operation
  buffer_memory_mb: 512 # Memory allocated for processing
  kv_endpoints: 1 # Number of endpoints used for key-value protocol
```

# Audit Log

``` yml
audit:
  enabled: true # Enable audit log
  log_path: "/path/to/audit.log" # Path to log file
  log_rotation_size_mb: 100 # Maximum size of each log file (MB)
  log_retention_days: 7 # Retention period for log files (days)
``` 
# Custom Transformation
* In some cases, data may need to be transformed before writing to Elasticsearch.
* However, this transformation cannot be done directly in a TOML file.
* Why? The TOML configuration format for CBES is designed to support the connection setup between Elasticsearch and Couchbase. However, complex operations like data transformation are not natively supported.
* To perform transformations, we can define a pipeline in Elasticsearch and reference it in the TOML file.
* First, create a pipeline in Elasticsearch:
``` yml
PUT _ingest/pipeline/news-transform-pipeline
{
  "processors": [
    {
      "script": {
        "source": """
          ctx.id = ctx.id;
          ctx.title = ctx.title;
          ctx.content = ctx.content;
          ctx.author = ctx.author;
          ctx.timestamp = ctx.created_at;
        """
      }
    }
  ]
}
```

* Then, simply reference the pipeline in the TOML file under the [elasticsearch.type] property:

``` toml
[[elasticsearch.type]]
 prefix = ''
 index = 'news'
 pipeline = 'news-transform-pipeline'
```
