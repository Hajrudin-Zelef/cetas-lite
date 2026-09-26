---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-79-2
title: "Summary"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-79.md
source_anchor: ""
source_lines: [88, 186]
sha256: f9ba195fb7e679ad15b93dcd8da8effe92e80e8d4357cd6baa9a7d9acd00f402
---

# Summary

| **Property** | Description | 
|---|---|
| **port** (Default:**443** ) | TCP port (for both tcp-conn and http-get probes) | 
| **certificate** (Default:"") | Certificate from local store that should be used for host verification. | 
| **check-certificate** (*yes* \|*no* ; Default**no** ) | Enables trust chain validation from local certificate store. | 

## HTTP-GET/HTTPS-GET probe pass/fail criteria

| Property | Description | 
|---|---|
| **thr-http-time** (Default:**10s** ) | Fail threshold for http-resp-time | 
| **http-code-min** (Default:**100** ) | OK/fail criteria for HTTP response code. | 
| **http-code-max** (Default:**299** ) | Response in the range [ `http-code-min` ,`http-code-max` ] is a probe pass/OK; outside - a probe fail. See mozilla-http-status or rfc7231 | 

| Property | Description | 
|---|---|
| **host** (Default:"") | DNS name that should be resolved. | 
| **record-type** (*A* \|*AAAA* \|*MX* \|*NS* ; Default:**A** ) | Record type that will be used for DNS probe. | 
| **dns-server** | The DNS server that the probe should send its requests to, if not specified it will use the value from " `/ip dns` ". | 

# Probe statistics/variables

You can view statistics and use these variables in scripting, keep in mind that variables containing "-" must be written like this, for example, "done-tests" would be $"done-tests"

## Generic:

| Property | Description | 
|---|---|
| **name** | user added name for the Netwatch entry | 
| **comment** | user added comment | 
| **host** | host that was probed | 
| **type** | probe type | 
| **interval** | interval | 
| **timeout** | timeout | 
| **since** | The last time the status change happened | 
| **status** | The current status of probe | 
| **done-tests** | total count of probe tests already done so far | 
| **failed-tests** | count of failed probe tests | 

## ICMP:

| Property | Description | 
|---|---|
| **sent-count** | ICMP packets sent out | 
| **response-count** | Matching/valid ICMP packet responses received | 
| **thr-loss-count** | number of lost packets | 
| **thr-loss-percent** | number of lost packets in percent | 
| **thr-avg** | mean value of round trip time (rtt-avg) | 
| **thr-max** | max round trip time (rtt-max) | 
| **thr-jitter** | jitter ( = max - min) of round trip time (rtt-jitter) | 
| **thr-stdev** | standard deviation of round trip time (rtt-stdev) | 

## TCP:

| Property | Description | 
|---|---|
| **tcp-connect-time** | time taken to establish a TCP connection | 

## HTTP:

| Property | Description | 
|---|---|
| **http-status-code** | HTTP response status code (200 OK, 404 Not Found, etc.). See mozilla-http-status or RFC7231 | 

## HTTPS:

| Property | Description | 
|---|---|
| **http-status-code** | HTTP response status code (200 OK, 404 Not Found, etc.). See mozilla-http-status or RFC7231 | 

## DNS:

| Property | Description | 
|---|---|
| **ip** | IP address - the result of A record-type probe | 
| **ip6** | IPv6 address - the result of AAAA record-type probe | 
| **mail-servers** | Mail servers along with their priority - the result of MX record-type probe | 
| **name-servers** | Name servers - the result of NS record-type probe | 

# Logs

On each probe's OK/fail state change:

- probe identification info and OK->fail or fail->OK is printed to info level
- detailed probe stats and config is printed to debug level

# Status

Command */tool/netwatch/print* will show the current status of Netwatch and **read-only** properties:

- since - Indicates when a state of the host changed last time;
- status - Shows the current status of the host;
- host - address being monitored

# Quick Example

Here we will use a simple ICMP check to host with IP 8.8.8.8:

Afterward, in the logging section we can see Netwatch executed script:
