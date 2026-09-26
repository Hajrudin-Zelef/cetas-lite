---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-85-2
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-85.md
source_anchor: ""
source_lines: [159, 264]
sha256: 6a7bb1aece91dd634eaf178d950c0dfe211d259452886be4876b9f3089573307
---

# Overview

```
$ curl -k -u admin: -X POST https://10.155.101.214/rest/ip/address/print\
  --data '{"_proplist": ["address","interface"]}' -H "content-type: application/json"
[{"address":"192.168.99.2/24","interface":"dummy"},
{"address":"172.16.5.1/24","interface":"sfpplus1"},
{"address":"172.16.6.1/24","interface":"sfp2"},
{"address":"172.16.7.1/24","interface":"sfp3"},
{"address":"10.155.101.214/24","interface":"sfp12"},
{"address":"192.168.111.111/32","interface":"dummy"}]
```
 ### Query

The '`.query`' key is used to create a query stack. The value is a list of query words. For example this POST request :

```
POST https://router/rest/interface/print
{".query":["type=ether","type=vlan","#|!"]}
```
 is equivalent to this API sentence

/interface/print
?type=ether
?type=vlan
?#|!

 For example, let's combine '*query*' and '*proplist*', to return '.id', 'address', and 'interface' properties for all dynamic records and records with the network 192.168.111.111

```
$ curl -k -u admin: -X POST https://10.155.101.214/rest/ip/address/print \
  --data '{".proplist": [".id","address","interface"], ".query": ["network=192.168.111.111","dynamic=true","#|"]}'\
  -H "content-type: application/json"
[{".id":"*8","address":"10.155.101.214/24","interface":"sfp12"},
{".id":"*A","address":"192.168.111.111/32","interface":"dummy"}]
```
 ### Timeout

If the command runs indefinitely, it will timeout and the connection will be closed with an error. The current timeout interval is 60 seconds. To avoid timeout errors, add a parameter that would sufficiently limit the command execution time. 

For example, let's see what we get when the ping command exceeds the timeout and how to prevent this by adding a count parameter:

```
$ curl -k -u admin: -X POST https://10.155.101.214/rest/ping \
  --data '{"address":"10.155.101.1"}' \
  -H "content-type: application/json"
{"detail":"Session closed","error":400,"message":"Bad Request"}
$ curl -k -u admin: -X POST https://10.155.101.214/rest/ping \
  --data '{"address":"10.155.101.1","count":"4"}' \
  -H "content-type: application/json"
[{"avg-rtt":"453us","host":"10.155.101.1","max-rtt":"453us","min-rtt":"453us","packet-loss":"0","received":"1","sent":"1","seq":"0","size":"56","time":"453us","ttl":"64"},
{"avg-rtt":"417us","host":"10.155.101.1","max-rtt":"453us","min-rtt":"382us","packet-loss":"0","received":"2","sent":"2","seq":"1","size":"56","time":"382us","ttl":"64"},
{"avg-rtt":"495us","host":"10.155.101.1","max-rtt":"650us","min-rtt":"382us","packet-loss":"0","received":"3","sent":"3","seq":"2","size":"56","time":"650us","ttl":"64"},
{"avg-rtt":"461us","host":"10.155.101.1","max-rtt":"650us","min-rtt":"359us","packet-loss":"0","received":"4","sent":"4","seq":"3","size":"56","time":"359us","ttl":"64"}]
```
 Another example is a bandwidth test tool, which can be limited by providing run duration:

```
$ curl -k -u admin: -X POST 'https://10.155.101.214/rest/tool/bandwidth-test' \
  --data '{"address":"10.155.101.1","duration":"2s"}' \
  -H "content-type: application/json"
[{".section":"0","connection-count":"20","direction":"receive","lost-packets":"0",
"random-data":"false","rx-10-second-average":"0","rx-current":"0","rx-size":"1500",
"rx-total-average":"0",
"status":"connecting"},
{".section":"1","connection-count":"20","direction":"receive","duration":"1s",
"lost-packets":"0","random-data":"false","rx-10-second-average":"0","rx-current":"0",
"rx-size":"1500","rx-total-average":"0",
"status":"running"},
{".section":"2","connection-count":"20","direction":"receive","duration":"2s",
"lost-packets":"581175","random-data":"false","rx-10-second-average":"854372352",
"rx-current":"854372352","rx-size":"1500","rx-total-average":"854372352",
"status":"running"},
{".section":"3","connection-count":"20","direction":"receive","duration":"3s",
"lost-packets":"9014","random-data":"false","rx-10-second-average":"891979008",
"rx-current":"929585664","rx-size":"1500","rx-total-average":"891979008",
"status":"done testing"}]
```
 # Errors

The success or failure of the API calls is indicated in the HTTP status code. In case of failure (status code 400 or larger), the body of the response contains a JSON object with the error code, a description of the error, and optional error details. For example, trying to delete an interface will return

`{"error":406,"message":"Not Acceptable","detail":"no such command or directory (remove)"}`
 # Examples

Below we have added a short collection of REST API calls you can make to your devices:

```
Make a log entry:
curl -k -u <username>:<password> -X POST http://<ip-address>/rest/execute --data '{"script":"/log/info test"}' -H "content-type: application/json"
Run a script:
curl -k -u <username>:<password> https://<ip-address>/rest/system/script/run --data '{".id":"*1"}' -H "content-type: application/json"
LTE monitor once:
curl -k -u <username>:<password> https://<ip-address>/rest/interface/lte/monitor -d '{"numbers":"0", "once":""}' -H "content-type: application/json"
WiFi monitor once:
curl -k -u <username>:<password> -X POST "http://<ip-address>/rest/interface/wifi/monitor" -H "Content-Type: application/json" -d '{ "numbers": "wifi1", "once":"" }'
Export device configuration:
curl -k -u <username>:<password> https://<ip-address>/rest/export --data '{"compact":"","file":"test.rsc"}' -H "content-type: application/json"
Move a firewall entry (switch positions):
curl -k -u <username>:<password> -X POST http://<ip-address>/rest/ip/firewall/nat/move --data '{".id":"*9",".id":"*C"}' -H "content-type: application/json"
LTE firmware update:
curl -k -u <username>:<password> -X POST 'http://<ip-address>/rest/interface/lte/firmware-upgrade'   --data '{"number":"lte2"}' -H "content-type: application/json"
Get OIDs from /system resource:
curl -k -u <username>:<password> -X POST http://<ip-address>/rest/system/resource/print --data '{"oid":""}' -H "content-type: application/json"
```
 An example of using `/tool fetch` to run a REST API POST to another RouterOS device:

`/tool fetch http-method=post url="http://<ip-address>/rest/execute" http-data="{\"script\":\"/log info fetchtest\"}" http-header-field="Content-Type:application/json" output=user user=<username> password=<password>`
