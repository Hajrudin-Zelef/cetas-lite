---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-e4e1ebe2-3842-11f1-b31d-02356ffb40d9-fortimanag-9d0f6ae1-1
title: "docs-fortinet-com-v2-attachments-e4e1ebe2-3842-11f1-b31d-02356ffb40d9-fortimanag-9d0f6ae1"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["2026-04-21"]
keywords: ["license", "parameters", "training"]
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-e4e1ebe2-3842-11f1-b31d-02356ffb40d9-fortimanag-9d0f6ae1.md
source_anchor: ""
source_lines: [1, 236]
sha256: 888fde4ccd4484232c629369d7ce17a69a413ec161a2202c5ab06f37510d7b6c
---

# docs-fortinet-com-v2-attachments-e4e1ebe2-3842-11f1-b31d-02356ffb40d9-fortimanag-9d0f6ae1

API Best Practices Guide
FortiManager 8.0.0

FORTINET DOCUMENT LIBRARY
https://docs.fortinet.com
FORTINET VIDEO LIBRARY
https://video.fortinet.com
FORTINET BLOG
https://blog.fortinet.com
CUSTOMER SERVICE & SUPPORT
https://support.fortinet.com
FORTINET TRAINING & CERTIFICATION PROGRAM
https://www.fortinet.com/training-certification
FORTINET TRAINING INSTITUTE
https://training.fortinet.com
FORTIGUARD LABS
https://www.fortiguard.com
END USER LICENSE AGREEMENT
https://www.fortinet.com/doc/legal/EULA.pdf
FEEDBACK
Email: techdoc@fortinet.com
April 21, 2026
FortiManager 8.0.0 API Best Practices Guide
02-800-1281541-20260421

TABLE OF CONTENTS
Change Log 4
Introduction 5
Getting started 6
Best Practices 7
Authentication and access control 7
API requests 8
Error handling and resilience 13
Performance optimization 13
Security and compliance 21
Change management 22
Support and resources 22
FortiManager 8.0.0 API Best Practices Guide 3
Fortinet Inc.

Change Log
Date Change Description
2026-04-21 Initial release.
FortiManager 8.0.0 API Best Practices Guide 4
Fortinet Inc.

Introduction
The FortiManager JSON API allows you to perform configuration and monitoring operations on a FortiManager
appliance or VM. The JSON API is based on JSON-RPC, a remote procedure call protocol encoded in JSON. It
provides access to a wide range of features, including device management, policy configuration, object
manipulation, and task execution.
By using structured JSON requests over HTTPS, you can efficiently interact with FortiManager to streamline
operations, and scale network management tasks across multiple Fortinet devices.
This document contains the following information:
l Getting started on page 6
l Best Practices on page 7
FortiManager 8.0.0 API Best Practices Guide 5
Fortinet Inc.

Getting started
Getting started
1. Access the Fortinet Developer Network: Find detailed information about the FortiManager API on the
Fortinet Developer Network site, https://fndn.fortinet.net. The Fortinet Developer Network site provides
documentation for getting started with FortiManager APIs, setting up API users, making API requests, and
more.
Access to the Fortinet Developer Network requires an account. Please contact your Fortinet
representative or Fortinet support for more information on getting account access to the
Fortinet Developer Network.
2. Create an API User: Configure a FortiManager API user. The type of API user required depends on the type
of authentication you will be using:
REST API Admin with a
predefined API key
1. Create a Rest API Admin with Read/Read-Write privileges configured for
JSON API Access.
2. Apply an Admin Profile with the required permissions to complete the
desired requests.
3. Copy the automatically generated API key to use in the request header
using the bearer authentication scheme.
The generated API key is permanent, which means the same
user account will always share the same session and you do
not need to use thelogin/logout endpoints. Because of this
benefit, predefined API keys are useful for simplifying
automation workflows.
Session-based
authentication
1. Create a regular administrator with Read/Read-Write privileges
configured for JSON API Access.
2. Apply an Admin Profile with the required permissions to complete the
desired requests.
3. Use thelogin endpoint with the configured administrator credentials to
generate the session ID to be used within API requests.
3. Use a lab environment to test APIs: When using FortiManager APIs, it is recommended to begin first in a
lab environment to test and validate workflows before deployment to production.
FortiManager 8.0.0 API Best Practices Guide 6
Fortinet Inc.

Best Practices
Best Practices
FortiManager offers powerful APIs to automate and integrate with your network and security management
workflows. Follow these best practices to ensure secure, efficient, and reliable use.
1. Authentication and access control on page 7
2. API requests on page 8
3. Error handling and resilience on page 13
4. Performance optimization on page 13
5. Security and compliance on page 21
6. Change management on page 22
7. Support and resources on page 22
Authentication and access control
Best practice Description
Uselogout requests when
using session-based
authentication
l Sendinglogout requests when using FortiManager’s JSON API is
critical to free up session slots. Neglecting this can inadvertently fill up
all available admin sessions, either at the individual user or system
level, resulting in denial of management access until sessions expire or
are forcefully ended.
l This is controlled by configurable parameters at both the user
(system.admin.user.login-max) and system level
(system.admin.setting.admin-max-login).
Example: Making a logout request
{
"method": "exec",
"params": [
{
"url": "/sys/logout"
}
],
"session": "...",
"id": 1
}
Leverage role-based access
control
l Use RBAC (Role-Based Access Control) to restrict API access to only
required resources. You can configure RBAC for your API administrator
using Admin Profiles configured in the FortiManager. For more
FortiManager 8.0.0 API Best Practices Guide 7
Fortinet Inc.

Best Practices
Best practice Description
information, see Administrator Profiles.
Restrict network access
l Limit access with an IP allowlist or using trusted source control. See
Trusted hosts.
l When additional flexibility is required (for example, if you require more
than 10 trusted hosts), you can leverage system.local-in policies.
API requests
Best practice Description
Use HTTPS l Use HTTPS only for all API communications.
Use proper formatting
l Structure requests using proper JSON format and required parameters.
Example: HTTP method and URL
POST https://<fmg_ip>/jsonrpc
Example: Standard JSON RPC body format
{
"method": "...",
"params": [ ... ],
"session": "...",
"id": 1
}
Leverage and monitor
asynchronous APIs
l Take advantage of the asynchronous mechanism available in certain
JSON API URLs.
If the API supports task creation, include flags likecreate_task and
nonblocking in your JSON request to initiate a new task. This approach
allows you to send multiple JSON requests sequentially, each creating
its own task. You can then wait for all tasks to finish asynchronously.
A common use case is adding multiple devices using separate JSON
requests. Alternatively, using theadd/dev-list method is also an
effective option.
FortiManager 8.0.0 API Best Practices Guide 8
Fortinet Inc.

Best Practices
Best practice Description
Example: Refresh multiple devices
{
"params": [
{
"url": "/dvm/cmd/update/dev-list",
"data": {
"adom": "root"
"flags": [
"create_task",
"nonblocking"
],
"update-dev-member-list": [
{
"name": "fgt_1"
},
{
"name": "hub_1"
},
{
"name": "hub_2"
}
]
}
}
],
"session": "...",
"id": 1
}
l Monitor asynchronous APIs to ensure task completion. If youlogout
while an API request run asynchronously, it might be reported as failed
in the Task Manager because FortiManager can't find a valid API
session to complete the task.
FortiManager 8.0.0 API Best Practices Guide 9
Fortinet Inc.

Best Practices
Best practice Description
Example: Monitor asynchronous APIs
{
"method": "get",
"params": [
{
"url": "/task/task/1"
}
],
"session": "...",
"id": 1,
"verbose": 1
}
Use option names in requests
l For any attribute that is defined as an option, always use name of
option instead of internal index value in JSON request. The name is
more intuitive and when used, it produces more readable and easier to
understand code.
For example, forfirewall/policy the action can includedeny,accept,
ipsec.
FortiManager 8.0.0 API Best Practices Guide 10
Fortinet Inc.

