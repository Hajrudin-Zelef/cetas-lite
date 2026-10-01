---
id: collect-261001-fortinet/fortinet/fortigate-3-technical-tip-about-rest-api-97302-c3b8f1ec
title: "fortigate-3-technical-tip-about-rest-api-97302-c3b8f1ec"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/fortigate-3-technical-tip-about-rest-api-97302-c3b8f1ec.md
source_anchor: ""
source_lines: [1, 42]
sha256: 50d345fd15141adbeff8651ecb02ba2b17b547dbb4339b8718395f58c71cadc0
---

# fortigate-3-technical-tip-about-rest-api-97302-c3b8f1ec

Technical Tip: About REST API
Description
This article describes the FortiGate REST API.
Scope
FortiGate.
Solution
The REST API can be used to retrieve, create, update, and delete configuration settings, as well as to retrieve system logs and statistics, and to perform basic administrative actions such as reboot and shut down through programming script.
FortiOS versions below v7.0.13:
There are two ways the user can authenticate against the API:
- Session-based authentication.
- Token-based authentication.
FortiOS v7.0.13 and above:
The supported and recommended way of authenticating with FortiOS to gain REST API access is to use a REST API admin.
- Token-based authentication.
Authentication methods:
- Session-based authentication.
The authentication is valid per login session. The user needs to send a login request to obtain an authentication cookie and CSRF token to be used for subsequent requests. The user then needs to send a logout request to invalidate the authentication cookie and CSRF token.
The CSRF token is available in the session csrftoken cookie, which must be included in the request header under X-CSRFTOKEN.
Note: The HTTP (POST/PUT/DELETE) method require CSRF tokens. Read requests HTTP (GET) do not require CSRF tokens.
How to get CSRF token from the fortigate firewall:
Perform an HTTP POST Request:
Fortigate-IP/logincheck username=AdminUser&secretkey=AdminPassword&ajax=1
HTTP Response:
The FortiGate will respond with 3 cookies: variable APSCOOKIE_9538334086037707851, ccsrftoken and ccsrftoken_9538334086037707851.
$response = Invoke-WebRequest -Uri "http://x.x.x.x/logincheck?username=testuser&secretkey=fortinet&ajax=1" -Method Post -WebSession $session
Using the CSRF Token:
$headers = @{
 "X-CSRFTOKEN" = "EFE4FADF74599229187FA9EABACD8F"
 }
 $response = Invoke-WebRequest -Uri "http://x.x.x.x/api/v2/cmdb/webfilter/profile/" -WebSession $session -Headers $headers
-  Token-based authentication.
The authentication is done via a single API token. This token is only generated when creating an API admin. The user must store this token in a safe place because it cannot be retrieved again. The user can however regenerate the token at any time. Each API request must include the token to be authenticated as the associated API admin
Create an API admin:
The FortiOS REST APIs support the following HTTP methods:
| HTTP Method | Description | 
| GET | Retrieve a resource or collection of resources. | 
| POST | Create a resource or execute actions. | 
| PUT | Update a resource. | 
| DELETE | Delete a resource or collection of resources. | 
Note:
REST API tokens cannot be used when the FortiGate is in FIPS CC mode. The REST API admin account option is not available for FIPS-CC mode.
Related document:
