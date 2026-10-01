---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-e4e1ebe2-3842-11f1-b31d-02356ffb40d9-fortimanag-9d0f6ae1-2
title: "docs-fortinet-com-v2-attachments-e4e1ebe2-3842-11f1-b31d-02356ffb40d9-fortimanag-9d0f6ae1"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-e4e1ebe2-3842-11f1-b31d-02356ffb40d9-fortimanag-9d0f6ae1.md
source_anchor: ""
source_lines: [237, 498]
sha256: cc45d9c971c45858362d1ec8d5e0929a9a0a362ef91f05b0708b97c1d5d02d0a
---

# docs-fortinet-com-v2-attachments-e4e1ebe2-3842-11f1-b31d-02356ffb40d9-fortimanag-9d0f6ae1

Best Practices
Best practice Description
Example: Using option names in a request
{
"method": "add",
"params": [
{
"url": "/pm/config/adom/root/pkg/test_
pkg/firewall/policy",
"data": [
{
"name": "Allow_HTTP",
"srcintf": "port1",
"dstintf": "port2",
"srcaddr": "all",
"dstaddr": "all",
"action": "accept",
"schedule": "always",
"service": "HTTP",
"logtraffic": "all"
}
]
}
],
"session": "...",
"id": 1
}
l Avoid the use of the numeric index values, such as0,1,2.
FortiManager 8.0.0 API Best Practices Guide 11
Fortinet Inc.

Best Practices
Best practice Description
Avoid: Using numeric index values in request
{
"method": "add",
"params": [
{
"url": "/pm/config/adom/root/pkg/test_
pkg/firewall/policy",
"data": [
{
"name": "Allow_HTTP",
"srcintf": "port1",
"dstintf": "port2",
"srcaddr": "all",
"dstaddr": "all",
"action": 1,
"schedule": "always",
"service": "HTTP",
"logtraffic": "all"
}
]
}
],
"session": "...",
"id": 1
}
l Add "verbose": 1 in your get request to receive option names in the
response instead of internal index values.
Example: Using verbose output
{
"method": "get",
"params": [ ... ],
"session": "...",
"id": 1,
"verbose": 1
}
FortiManager 8.0.0 API Best Practices Guide 12
Fortinet Inc.

Best Practices
Error handling and resilience
Best practice Description
Monitor HTTP status codes
l Monitor standard HTTP status codes:
l 100: Request header received and waiting for the request body.
l 200: Success.
l 401: Invalid session or token.
l 403: Permission denied.
l 404: Not found. Unable to find the specified resource.
l 405: Method not allowed for this resource.
l 500: Internal error.
l 503: Service unavailable.
Use retry logic with
exponential backoff
l Implement retry logic with exponential backoff for transient errors to
prevent flooding the FortiManager with excessive repeat requests.
Log request/response for
auditing and troubleshooting
l Record the details of each API interaction, both in requests and
responses so that they can be assessed if an API calls or fails to return
data, and provide a clear record of the actions taken using the API.
Compare activities/tasks
performed from the GUI when
there are errors with the API
l Compare activities performed using the API with the same activities
performed using the GUI. Troubleshooting issues from the GUI
interface is an easier way to isolate problems and determine if the issue
is platform use/configuration or API related. If you have problems
performing a tasks using the GUI (for example, installs) then you'll likely
also face issues performing the same task using the API.
Performance optimization
Best practice Description
Limit request frequency l Avoid unnecessary polling or tight loops.
Use filter and fields
l Usefields,filter,range, andloadsub to retrieve only the necessary
data when querying large tables.
FortiManager 8.0.0 API Best Practices Guide 13
Fortinet Inc.

Best Practices
Best practice Description
Example: Getting a firewall address group containing member
"host_001"
{
"method": "get",
"params": [
{
"fields": [
"name",
"member"
],
"filter": [
"member",
"contain",
"host_001"
],
"loadsub": 0,
"url": "/pm/config/adom/root/obj/firewall/addrgrp"
}
],
"session": "...",
"id": 1
}
l Using thefilter parameter with thedelete method allows you to bulk
delete objects that match specific criteria. This can improve
performance when compared with deleting objects using individual
JSON requests. "confirm": 1 must be included in the JSON request to
confirm and execute the deletion.
FortiManager 8.0.0 API Best Practices Guide 14
Fortinet Inc.

Best Practices
Best practice Description
Example: Deleting objects matching a filter
{
"method": "delete",
"params": [
{
"confirm": 1,
"filter": [
"name",
"like",
"west%" <----- virtual IP that starts with
string west
],
"url": "/pm/config/adom/root/obj/firewall/vip"
}
],
"session": "...",
"id": 1
}
l Avoid loading full tables when only a subset of data is needed that
could be obtained using filter/fields.
Avoid: Loading full table when filter and fields can be used
{
"method": "get",
"params": [
{
"url": "/pm/config/adom/root/obj/firewall/addrgrp"
}
],
"session": "...",
"id": 1,
"verbose": 1
}
Cache data l Cache static or rarely changing data locally when appropriate.
Use object URLs for specific
items
l Access specific objects using the full object URL.
FortiManager 8.0.0 API Best Practices Guide 15
Fortinet Inc.

Best Practices
Best practice Description
Example: Using an object's URL
{
"method": "get",
"params": [
{
"url":
"/pm/config/adom/root/obj/firewall/address/Vancouver"
}
],
"session": "...",
"id": 1,
"verbose": 1
}
l Avoid using general table URLs when you only need access to a
specific object.
Avoid: Using general table URLs to access specific objects
{
"method": "get",
"params": [
{
"url": "/pm/config/adom/root/obj/firewall/address"
}
],
"session": "...",
"id": 1,
"verbose": 1
}
Leverage the target attribute
when using/sys/proxy/json
for multi-device FortiOS REST
API requests
l In order to send FortiOS REST API requests to multiple FortiGates at the
same time, use the JSON URL/sys/proxy/json and put all device
names in the target to achieve the asynchronous effect.
FortiManager 8.0.0 API Best Practices Guide 16
Fortinet Inc.

Best Practices
Best practice Description
Example: Using /sys/proxy/json
{
"method": "exec",
"params": [
{
"data": {
"action": "get",
"resource": "/api/v2/monitor/router/ipv4"
"target": [
"adom/root/device/fgt1"
]
},
"url": "sys/proxy/json"
}
],
"session": "...",
"id": 1
}
Example: Setting the target when there are multiple
devices/device groups
"target": [
"adom/root/group/emea_devices",
"adom/root/group/apac_devices",
"adom/root/device/fgt1",
"adom/root/device/fgt2"
]
Monitor system resources
l The number of JSON API tasks that can run simultaneously depends on
the currently available resources. This includes the free memory size,
number of CPUs, clock speed of CPU, free disk size, etc. Avoid running
too many tasks simultaneously without considering system limitations.
Clone large objects
l Cloning is faster than manually recreating complex objects and all sub-
objects from scratch
FortiManager 8.0.0 API Best Practices Guide 17
Fortinet Inc.

