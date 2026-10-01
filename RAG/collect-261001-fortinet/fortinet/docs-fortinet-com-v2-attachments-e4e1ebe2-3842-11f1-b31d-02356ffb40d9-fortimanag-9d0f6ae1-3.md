---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-e4e1ebe2-3842-11f1-b31d-02356ffb40d9-fortimanag-9d0f6ae1-3
title: "docs-fortinet-com-v2-attachments-e4e1ebe2-3842-11f1-b31d-02356ffb40d9-fortimanag-9d0f6ae1"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["embedding"]
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-e4e1ebe2-3842-11f1-b31d-02356ffb40d9-fortimanag-9d0f6ae1.md
source_anchor: ""
source_lines: [499, 728]
sha256: d1b11151f7a70d09422d2d0ebac30e3f36c9eed8e8f70bd345af1b54421d9a31
---

# docs-fortinet-com-v2-attachments-e4e1ebe2-3842-11f1-b31d-02356ffb40d9-fortimanag-9d0f6ae1

Best Practices
Best practice Description
Example: Cloning an ADOM
request = {
"method": "clone",
"params": [
{
"data": {
"name": "adom2"
},
"url": "/dvmdb/adom/adom1"
}
],
"session": "...",
"id": 1
}
Create a baseline
l Baseline your FortiManager with and without your API client tools
running. Depending on how you use FortiManager you might find
performance issues by creating a baseline with GUI only so you can
determine if your API calls are negatively impacting the system.
Use high-availability mode to
distribute installations
l If installation operations are slow, configure a FortiManager HA cluster
and enableperform-improve-by-ha {enable | disable} to distribute
the workload of installation tasks. See the CLI Reference for more
information.
Minimize table size to improve
performance in the GUI
l Consider how much data your tables present considering the API can
exponentially grow tables with ease. If possible, moving to a strategy
that creates fewer objects with common names but adds per-device
mapping per object will decrease the number of objects in the table
and therefore improve the experience in the GUI.
l As an example, if your address table(s) within an ADOM has tens or
hundreds of thousands of objects, you might notice a slowdown in
presenting this data in the GUI while the API with a proper filter does
not.
Optimize FortiManager
architecture for API efficiency
l Plan your FortiManager and FortiOS architecture before starting API
development to reduce complexity and technical debt. Leverage built-
in features like FortiManager variables and dynamic CLI/Jinja scripts to
minimize API calls and simplify configuration management.
l If you find yourself updating lots of objects in the form of lists and/or
groups that are time sensitive, consider creating dynamic lists in the
FortiGate using the SDN connectors. See Public and private
SDN connectors.
FortiManager 8.0.0 API Best Practices Guide 18
Fortinet Inc.

Best Practices
Best practice Description
Use multiplexing to regroup
multiple APIs into a single
request when appropriate
l Multiplexing can be used to group multiple APIs operations into a single
request.
For example, you can use data multiplexing when you need to create
multiple firewall addresses using one API call.
Example: Creating multiple firewall addresses using data
multiplexing
{
"method": "add",
"params": [
{
"data": [
{
"name": "test_001",
"type": "ipmask",
"subnet": [
"10.0.0.1",
"255.255.255.0"
]
},
{
"name": "test_002",
"type": "ipmask",
"subnet": [
"10.0.0.2",
"255.255.255.0"
]
},
{
"name": "test_003",
"type": "ipmask",
"subnet": [
"10.0.0.3",
"255.255.255.0"
]
}
],
"url": "/pm/config/adom/root/obj/firewall/address"
}
],
"session": "...",
"id": 1
}
l Parameter multiplexing can also be used to improve performance when
FortiManager 8.0.0 API Best Practices Guide 19
Fortinet Inc.

Best Practices
Best practice Description
all the JSON URLs in the request have the same prefix, for example
/pm/config.
Example: Retrieving information about specific objects in an
ADOM using parameter multiplexing
{
"method": "get",
"params": [
{
"url": "/pm/config/adom/root/obj/webfilter/profile",
"fields": ["name"],
"loadsub": 0
},
{
"url": "/pm/config/adom/root/obj/firewall/address",
"fields": ["name"],
"loadsub": 0
},
{
"url":
"/pm/config/adom/root/obj/firewall/service/group",
"fields": ["name"],
"loadsub": 0
}
],
"session": "...",
"id": 1
}
l Using parameter multiplexing in requests that contain URLs with
different prefixes is not supported.
FortiManager 8.0.0 API Best Practices Guide 20
Fortinet Inc.

Best Practices
Best practice Description
Avoid: Using parameter multiplexing with requests that
include different URL prefixes
{
"method": "get",
"params": [
{
"fields": [
"name"
],
"loadsub": 0,
"url": "/dvmdb/group"
},
{
"fields": [
"name"
],
"loadsub": 0,
"url": "/pm/config/adom/root/obj/firewall/address"
}
],
"session": "...",
"id": 1
}
When working in Workspace
mode, lock the ADOM once,
perform all necessary
changes, then commit and
unlock the ADOM
l Each commit performed in Workspace mode triggers a database copy
operation. To reduce the time taken to perform each copy operation,
plan your changes ahead, group them logically, and apply them in one
session before committing.
For example, lock the ADOM, add multiple objects, commit the
changes, and unlock the ADOM.
l Avoid locking the ADOM, making a single change, committing,
unlocking the ADOM, and repeating this process multiple times.
Security and compliance
Best practice Description
Change passwords and API
sessions
l Change admin passwords periodically to reduce the risk of
unauthorized access and exposure from compromised credentials.
l Rotate API sessions regularly by performinglogout andlogin
requests.
FortiManager 8.0.0 API Best Practices Guide 21
Fortinet Inc.

Best Practices
Best practice Description
Don't embed credentials l Avoid embedding credentials in scripts or config files.
Monitor your API activity
l Use the FortiManager Event Log to monitor API activity for anomalies.
For more information, see Event Log.
You can configure JSON API logging using theset jsonapi-log
command in the FortiManager CLI:
config system global
(global)# set jsonapi-log
all logging both jsonapi request & response.
disable disable jsonapi log.
request logging jsonapi request.
response logging jsonapi response.
Change management
Best practice Description
Maintain version compatibility
l Maintain version compatibility during FortiManager upgrades. Consult
the FortiManager Release Notes to verify which managed devices are
supported by your FortiManager version.
Support and resources
Best practice Description
Consult the Fortinet Developer
Network (FNDN)
l Refer to https://fndn.fortinet.net for API docs and tools.
l Join the Fortinet Developer Community for questions and examples.
Compare APIs between two
FortiManager versions using
the API Comparison feature
l Use the API Comparison feature in FortiAPI on the Fortinet Developer
Network to compare API information between two FortiManager
versions, and see APIs that have been updated, removed, or added.
FortiManager 8.0.0 API Best Practices Guide 22
Fortinet Inc.

Best Practices
Best practice Description
Contact Fortinet Support for
further assistance
l Contact Fortinet Support for API-related assistance, or use the
Fortinet Support Community to find answers to questions in the articles
and support forums.
FortiManager 8.0.0 API Best Practices Guide 23
Fortinet Inc.

