server-spark-submit
================

This server expose an API to create spark jobs with `spark-submit` command. Due that this server must run in a environment where that command is available.

Checkout the `example-submit.sh` to understand how to call this API.

### Running the server with default args

```
SPARK_SUBMIT_DEFAULT_ARGS="--deploy-mode client --packages com.databricks:spark-csv_2.11:1.4.0,com.datastax.spark:spark-cassandra-connector_2.10:1.6.0 --conf spark.cassandra.connection.host=cassandra.production-cassandra.svc.cluster.local --master spark://10.2.93.2:7077 --verbose" server-spark-submit
```

### Building

```
go build  -ldflags "-X main.version=1"
```

Or

```
goxc
```
