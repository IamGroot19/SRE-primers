Topics:

A: (20-30 minutes)
- Intro to loki; how is it different from other solutions & why is it better?
  - Multi-tenancy
  - Cover the idea of labels 

- Explain loki architecture in overall 
  - Components in read & write path
  
B: 30-40 minutes
- Talk about different clients && segue into promtail
  - Talk about service discovery & other good aspects of promtail

- Brief hands-on using Grafana cloud
  - Show the promtail dashboard (localhost:9080) & explain various things
  - Show around Grafana: Adding datasource, adding a panel & using explore

- Show basics of logQL  

- Briefly talk about logs based alerting & also metric queries
  - Show how to add alerting from grafana

c: 15-20 minutes
- Production setup
  - Show some stats
  - Some Loki Operational dashboards
  - Production incidents??


# Grafana cloud + local promtail setup


## Steps

Download this file https://github.com/IamGroot19/sre-primers/blob/master/elk/apache-logs-small  into 2 locations: a) `/tmp` b) `~/Documents` 

Create a grafana cloud account (single click signup - no credit card required). Create a new connection > search for "loki Hosted Logs" and you will get a URL. Add it to promtail config below.

Then, run the docker container

`docker run  -d --name promtail --volume "$PWD:/etc/promtail" --volume "/var/log:/var/log" -v "/var/lib/docker/containers:/var/lib/docker/containers" -v "/var/log/journal:/var/log/journal" -v "/tmp:/tmp" --net host  grafana/promtail:2.8.0 -config.file=/etc/promtail/promtail-grafana-cloud.yaml`

Promtail config at promtail-grafana-cloud.yaml
- for docker based detup



# LogQL

LogQL: PromQL inspired lang for fetching logs from loki.

2 types:
- Log queries: Fetch log lines
- Metric Queries: Extend log queries to calculate values based on query results.


## Basic Log queries (searching for stuff in logs)

Get all loglines from a given job
```
{job="sample-apache-accesslog-2"}
```


Get all loglines from a given job & a host
```
{job="sample-apache-accesslog-2", host="tpx13"} |~ `404`
```


Get count of all requests  for the path `/glob/*` over the past 7days
```
count_over_time({job="sample-apache-accesslog-2"} |~ `/blog/.*` [7d])
```


Search for particular range of IP Address
```
{job="sample-apache-accesslog-2"} |= ip("82.165.139.0-82.165.150.255")
```


Search for particular CIDR of IPs
``` 
{job="sample-apache-accesslog-2"} |= ip("82.165.139.0/8")
```


Attempting to query raw docker logs
```
{job="docker-new", filename="/var/lib/docker/containers/0ec9485bedd96a6b6d169e210aac63ac4364e0c9adff392513842ded491faf97/0ec9485bedd96a6b6d169e210aac63ac4364e0c9adff392513842ded491faf97-json.log"} | logfmt
```
- will have to do a lot of guesswork since container IDs are not informative enough

Search the OS login logs
- For a signed in user
```
{job="systemd-services-journal-sd"} |= "session opened"
```
- For wrong logins
```
{job="systemd-services-journal-sd"}
|~ "Invalid user.*"
  | regexp "(^(?P<user>\\S+ {1,2}){8})"
  | regexp "(^(?P<ip>\\S+ {1,2}){10})"
  | line_format "IP = {{.ip}}\tUSER = {{.user}}"
```
- NOTICE: We have multi-stage filtering here:
  - Stream selector: `{job="systemd-services-journal-sd"}`
  - Line filters: `|~ "Invalid user.*"  | regexp "(^(?P<user>\\S+ {1,2}){8})" | regexp "(^(?P<ip>\\S+ {1,2}){10})"`
  - Label filters: `| line_format "IP = {{.ip}}\tUSER = {{.user}}"`


## JSON parsing

Add the nginx-json-logs/nginx-access.log file to your promtail config

If logline is valid json, then aadding `| json` extracts all properties of the json document
- NOTE: Except for json arrays (which are skipped)

```
{job="nginx-json"} | json
```














