# Setup

Just use the docker compose to get ELK stack up. 
- As a small testcase / exploration, the logstash pipeline takes apache access logs from a file and dumps it into elasticsearch. 
- When using kibana for the first time, you need to setup index patterns to access the particular index. 

# Step1: Explore ES + Kibana

Before running the docker compose, comment out `command` for logstash container & instead use this: `command: [while true; do sleep 900; done]`

Also, set this up: `sysctl -w vm.max_map_count=262144`


## Instructions

Simply hit the ES endpoint && you will see basic details of the cluste: version, tag line etc.

Next, login to kibana. Just look around at various sections. 

Go to "Dev Tools". It's an interface where you can talk to Elasticsearch's REST API using curl-like commands. 

### Cmd to get cluster info
```

GET _cluster/health
GET _cluster/settings
GET _cluster/stats

```

### Cmd Around Indices
(mostly CRUD operations)

```
# List all indices of the cluster
GET _cat/indices?v
```

Create a doc & automatically its index as well
```
PUT /demo-index-1/_doc/1
{  
    "title": "Elasticsearch Indices Demo",
    "author": "Groot",
    "job": "SRE",
    "tagline": "I am Groot"
}
```
- PUT request overwrites existing doc (or creates new if it doesnt exist)
- Note: we are explicitly mentioning document's ID 

Create document using POST
```
POST /demo-index-1/_doc/
{  
    "title": "Elasticsearch Indices Demo",
    "author": "Rocket",
    "job": "SRE",
    "tagline": "The name's Rocket"
}
```

Fetch this particular index
```
GET _cat/indices/demo-index-1?v

GET demo-index-1

GET demo-index-1/_search
```


Look at data type of each field
- understand that this is dynamic mapping


Update / delete a document (we know it's ID)
```
PUT /demo-index-1/_doc/1
{  
    "title": "Elasticsearch Indices Demo",
    "author": "Teen Groot",
    "job": "SRE",
    "tagline": "I am GROOT"
} 

DELETE demo-index-1/_doc/1
```
- what to do if ID is unknown?
  - Use Update By Query API: https://www.elastic.co/guide/en/elasticsearch/reference/7.12/docs-update-by-query.html
  - use Delte by Query API: https://www.elastic.co/guide/en/elasticsearch/reference/7.12/docs-delete-by-query.html


Setup Cerebros
```
sudo docker run --rm --net host  lmenezes/cerebro:0.9.4

In the host, give: http://127.0.0.1:9200
```
