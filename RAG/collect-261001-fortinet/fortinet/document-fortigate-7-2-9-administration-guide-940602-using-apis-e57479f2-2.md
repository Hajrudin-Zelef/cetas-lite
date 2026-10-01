---
id: collect-261001-fortinet/fortinet/document-fortigate-7-2-9-administration-guide-940602-using-apis-e57479f2-2
title: "document-fortigate-7-2-9-administration-guide-940602-using-apis-e57479f2"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-2-9-administration-guide-940602-using-apis-e57479f2.md
source_anchor: ""
source_lines: [118, 196]
sha256: d796e1ea4419755976c3eba13d21e3e767312fe93591f03d6fd7723e004ae075
---

# document-fortigate-7-2-9-administration-guide-940602-using-apis-e57479f2

                                                    Go to Settings > General and turn off SSL certificate verification.
- 
                                                    Click on New and select HTTP
- 
                                                    In the new request, click on the Authorization tab, select Type as Bearer Token and enter <YOUR-API-TOKEN> in the Token field.
- 
                                                    Enter a URL like the one below: https://<FortiGate_address>/api/v2/cmdb/firewall/address/? format=name|comment.
- 
                                                    Click Send.
Results of the formatted API call:
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
            "comment": ""
        },
        {
            "name": "EMS_ALL_UNMANAGEABLE_CLIENTS",
            "q_origin_key": "EMS_ALL_UNMANAGEABLE_CLIENTS",
            "comment": ""
        },
        {
            "name": "FABRIC_DEVICE",
            "q_origin_key": "FABRIC_DEVICE",
            "comment": "IPv4 addresses of Fabric Devices."
        }, 
                                            Filtering an API call
The filter parameter can be used to specify a field and a keyword to limit what results match and are returned by a call. In this example, the preceding call is used with a filter to return only names and comments for address objects with the word Sales in the name.
To use the filter parameter in an API call using curl:
curl --insecure \
-H "Accept: application/json" \
-H "Authorization: Bearer <API-TOKEN>" https://<FortiGate_address>/api/v2/cmdb/firewall/address?format="name|comment&filter=name=@SSLVPN"
The backslash (\) allows for multiline commands in curl. You can choose to enter the backslashes or enter the commands into a single wrapping line.
To use the filter parameter in an API call using Postman:
- 
                                                    Open the Postman client.
- 
                                                    Go to Settings > General and turn off SSL certificate verification.
- 
                                                    Click on New and select HTTP
- 
                                                    In the new request, click on the Authorization tab, select Type as Bearer Token and enter <YOUR-API-TOKEN> in the Token field.
- 
                                                    Enter a URL like the one below: https://<FortiGate_address>/api/v2/cmdb/firewall/address/? format=name|comment&filter=name=@SSLVPN.
- 
                                                    Click Send.
Results of the formatted API call:
curl and Postman display the output similar to the following (output shortened for brevity):
  {
    "http_method": "GET",
    "size": 17,
    "limit_reached": false,
    "matched_count": 1,
    "next_idx": 5,
    "revision": "bd002ee1735120907182831e7528dc8b",
    "results": [
        {
            "name": "SSLVPN_TUNNEL_ADDR1",
            "q_origin_key": "SSLVPN_TUNNEL_ADDR1",
            "comment": ""
        }
    ],
    "vdom": "root",
    "path": "firewall",
    "name": "address",
    "status": "success",
    "http_status": 200,
    "serial": "****************",
    "version": "******",
    "build": ****}
                                            For a complete list of API calls, see the Fortinet Development Network (FNDN). A subscription is required to access the FNDN.
