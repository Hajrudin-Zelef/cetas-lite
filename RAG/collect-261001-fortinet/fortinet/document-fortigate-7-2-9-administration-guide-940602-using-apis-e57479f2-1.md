---
id: collect-261001-fortinet/fortinet/document-fortigate-7-2-9-administration-guide-940602-using-apis-e57479f2-1
title: "document-fortigate-7-2-9-administration-guide-940602-using-apis-e57479f2"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-2-9-administration-guide-940602-using-apis-e57479f2.md
source_anchor: ""
source_lines: [1, 117]
sha256: b367d6e59932e6efebc91e4617f68029e90910e6948630565319c98d190b099d
---

# document-fortigate-7-2-9-administration-guide-940602-using-apis-e57479f2

Using APIs
Using APIs
Administrators can use API calls to a FortiGate to:
- 
                                                    Retrieve, create, update, and delete configuration settings
- 
                                                    Retrieve system logs and statistics
- 
                                                    Perform basic administrative actions, such as a reboot or shut down through programming scripts.
Token-based authentication
FortiGate supports only token-based authentication for API calls. Token-based authentication requires the administrator to generate a token, which is then included in each API request for authentication. A token is automatically generated when a new API administrator is created in FortiOS.
|  | Once the API administrator is created and the token displays, there is no way for the FortiGate to provide this token again. Ensure you record the token, and store it in a safe location; otherwise, you will have to generate a new token. | 
Creating the API administrator and generating the API token
When creating an API administrator, it is best practice to provide this account (and the associated token) with the minimum permissions required to complete the function. For example, if you only plan to use API calls to retrieve statistics or information from the FortiGate, the account should have read permissions.
|  | The API administrator account used in this topic's examples has full permissions strictly to illustrate various call types and does not adhere to the preceding recommendation. | 
See REST API administrator for detailed steps to create a REST API administrator.
Best Practices
Using API tokens with a request header
The API token can be included in any REST API request in either the request header or URL parameter. For added security, it is strongly recommended to use API tokens in the request header or transition your applications to include the API token in the request header instead of the URL parameter.
To pass the API token in the request header, the user needs to explicitly add the following field to the request header:
Authorization: Bearer <YOUR-API-TOKEN>
Configuring security options when creating a REST API administrator
In addition to using API tokens in the request header, for added security when creating a new REST API administrator, one or more of the following fields should be configured, listed in order of configuration difficulty from easy to difficult:
- 
                                                    Trusted Hosts To ensure that only trusted hosts/subnets can access the FortiGate REST API, you should configure the Trusted Hosts field when creating a new REST API administrator. You need your Source Address to create the trusted host.
- 
                                                    CORS Allow Origin Cross Origin Resource Sharing (CORS) allows third-party web apps to make API requests to the FortiGate using the token. For this field, specify the Access-Control-Allow-Origin on API responses. If possible, avoid using '* '.
- 
                                                    PKI Group Configure a PKI group to enable peer authentication using certificate matching, which provides an extra layer of security. Both the client certificate and token must match to be granted access to the API. The PKI group is typically specified as the name of a peer group defined under config user group which has PKI members. See Configuring a PKI user for details.
Making an API call to retrieve information from the FortiGate
The newly created API token is used to query the FortiGate for all firewall addresses. Many applications can be used for this query; this example uses the curl and Postman clients to demonstrate the functionality.
|  | To ensure maximum security when using API tokens, HTTPS is enforced. HTTP cannot be used. | 
Replace the placeholders below with values for your FortiGate:
- 
                                                    <FortiGate_address> is the IP address or hostname of your FortiGate as well as the HTTPS port number (default = 443 and does not need to be explicitly specified).
- 
                                                    <API-TOKEN> is the token you generated.
General API call
One of the simplest API calls is api/v2/cmdb/firewall/address, which returns all information about all firewall addresses.
To make a general API call using curl:
curl --insecure \
-H "Accept: application/json" \
-H "Authorization: Bearer <API-TOKEN>" https://<FortiGate_address>/api/v2/cmdb/firewall/address
The backslash (\) allows for multiline commands in curl. You can choose to enter the backslashes or enter the commands into a single wrapping line.
To make a general API call using Postman:
- 
                                                    Open the Postman client.
- 
                                                    Go to Settings > General and turn off SSL certificate verification.
- 
                                                    Click on New and select HTTP
- 
                                                    In the new request, click on the Authorization tab, select Type as Bearer Token and enter <YOUR-API-TOKEN> in the Token field.
- 
                                                    Enter a URL like the one below: https://<YOUR-FORTIGATE-ADDRESS>/api/v2/cmdb/firewall/address
- 
                                                    Click Send.
Results of the general API call:
curl and Postman display the output similar to the following (output shortened for brevity):
  {
    "http_method": "GET",
    "size": 17,
    "limit_reached": false,
    "matched_count": 17,
    "next_idx": 16,
    "revision": "bd002ee1735120907182831e7528dc8b",
    "results": [
        {
            "name": "EMS_ALL_UNKNOWN_CLIENTS",
            "q_origin_key": "EMS_ALL_UNKNOWN_CLIENTS",
            "uuid":"********-****-****-****-************"
            "type": "dynamic",
            "route-tag": 0,
            "sub-type": "ems-tag",
            "clearpass-spt": "unknown",
            "macaddr": [],
            "country": "",
            "cache-ttl": 0,
            "sdn": "",
            "fsso-group": [],
            "interface": "",
            "obj-tag": "",
            "obj-type": "ip",
            "tag-detection-level": "",
            "tag-type": "",
            "dirty": "clean",
            "hw-vendor": "",
            "hw-model": "",
            "os": "",
            "sw-version": "",
            "comment": "",
            "associated-interface": "",
            "color": 0,
            "filter": "",
            "sdn-addr-type": "private",
            "node-ip-only": "disable",
            "obj-id": "",
            "list": [],
            "tagging": [],
            "allow-routing": "disable",
            "fabric-object": "disable"
        },
        {
            "name": "EMS_ALL_UNMANAGEABLE_CLIENTS",
            "q_origin_key": "EMS_ALL_UNMANAGEABLE_CLIENTS",
            "uuid":"********-****-****-****-************",
                                            Formatting an API call
Since a general API call for address objects returns a large amount of information, it may be beneficial to format the API call to display certain information using the format parameter. In this example, the format parameter is used to display the name and comment for each firewall address.
To use the format parameter in an API call using curl:
curl --insecure \
-H "Accept: application/json" \
-H "Authorization: Bearer <API-TOKEN>" https://<FortiGate_address>/api/v2/cmdb/firewall/address?format="name|comment"
The backslash (\) allows for multiline commands in curl. You can choose to enter the backslashes or enter the commands into a single wrapping line.
To use the format parameter in an API call using Postman:
- 
                                                    Open the Postman client.
- 
