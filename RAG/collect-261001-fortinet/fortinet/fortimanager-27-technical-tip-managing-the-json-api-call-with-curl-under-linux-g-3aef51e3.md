---
id: collect-261001-fortinet/fortinet/fortimanager-27-technical-tip-managing-the-json-api-call-with-curl-under-linux-g-3aef51e3
title: "fortimanager-27-technical-tip-managing-the-json-api-call-with-curl-under-linux-g-3aef51e3"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/fortimanager-27-technical-tip-managing-the-json-api-call-with-curl-under-linux-g-3aef51e3.md
source_anchor: ""
source_lines: [1, 123]
sha256: 2e8ea17c6c2f951160870c4002c6b41921b52c17a32463f3f652ee509cbf8a1d
---

# fortimanager-27-technical-tip-managing-the-json-api-call-with-curl-under-linux-g-3aef51e3

Technical Tip: Managing the JSON API call with cURL under Linux/Gnu and how to get, delete and create a user using it
- Last edited: 5 months agoApril 16, 2026 at 1:04 PM
- 0 replies
- 5780 views
Description
This article describes how to use the JSON API calls under Linux/Gnu using the cURL command line tool and how to get, delete, create a user, and the possible options to use to filter using fields or specific usernames and the output difference.
At the end of the article are shown some common errors.
Scope
Understand the JSON API calls used under the FortiManager.
Note.
There is an attached text file that contains the full JSON API calls and the simplified version used in the article.
Be aware that these examples could be implemented by other programs or programming languages.
The related articles and their links are at the bottom of the page.
Solution
It is possible to use already prepared requests from the Postman.
It is possible to copy a directly generated cURL command line as shown below.
Start with a login to the remote FortiManager:
To receive the output parsed, it is possible to use the 'jq' (apt install jq – for Debian users).
Then to view only the session, it is possible to use the 'grep'. Like '| grep session'.
To save the output directly as a file, use the cURL option -o the_name.
curl -k --location --request POST 'https://10.109.54.6/jsonrpc' \
--header 'Content-Type: application/json' \
--data-raw '{
  "method": "exec",
  "params": [
    {
      "data": {
        "passwd": "",
        "user": "admin"
      },
      "url": "/sys/login/user"
    }
  ],
  "session": "string",
  "id": 1
}' | jq "."
To get the current users and their profiles, use the cURL to get the data that was desired and to grep through the output.
curl -k --location --request POST 'https://10.109.54.6/jsonrpc' \
--header 'Content-Type: application/json' \
--data-raw '{
  "method": "get",
  "params": [
    {
        "url": "/cli/global/system/admin/user "
    }
  ],
  "session": "QTnToxshxQXrog+lxbMEFkViZkvRuYAE+6K/YMNrNKytZaj/KzpL2TW0mPCAPFjTp+puVBagII04PMw5atBMtg==",
  "id": 1
}' | jq "." |  grep -e userid -e profileid
To get only specific user data, it is possible to specify a username. Additionally, it is possible to use the 'fields' option. This option will reduce the output fields.
{
  "method": "get",
  "params": [
    {
        "fields": 
            [
                "userid"
            ],
        "url": "/cli/global/system/admin/user/admin"
    }
  ],
  "session": "{{session}}",
  "id": 1
}
If the output is compared with and without the fields for the user admin, it is possible to observe that with the first option, it will show 162 lines and without it, 254 lines of output are obtained.
Note: Due to reduced output when the filter is used, the field 'profiled' would be missing!
Now, delete a user.
curl -k --location --request POST 'https://10.109.54.6/jsonrpc' \
--header 'Content-Type: application/json' \
--data-raw '{
  "method": "delete",
  "params": [
    {
      "url": "/cli/global/system/admin/user/novosco"
    }
  ],
  "session": "string",
  "id": 1
}' | jq "."
Create a new user.
curl -k --location --request POST 'https://10.109.54.6/jsonrpc' \
--header 'Content-Type: application/json' \
--data-raw '{
"method": "add",
"params": [
{
"data": [
{
"adom": [
{
"adom-name": "all_adoms"
}
],
"change-password": "enable",
"profileid": "Super_User",
"trusthost1":["100.0.0.0","255.0.0.0"],
"password": "test",
"userid": "novosco"
}
],
"url": "/cli/global/system/admin/user"
}
],
"session": "QTnToxshxQXrog+lxbMEFkViZkvRuYAE+6K/YMNrNKytZaj/KzpL2TW0mPCAPFjTp+puVBagII04PMw5atBMtg==",
"id": 1
}'  |  jq "."
Note: In some cases the strings are case-sensitive!
Common errors:
- Self-signed certificate:
curl: (60) SSL certificate problem: self-signed certificate
More details are available in this document.
When the self-signed certificate is used, add to the cURL the following: --k /--insecure.
- Session ID error: Do not reuse a session ID generated on another device as a permission denied error will be shown.
{ "status": { "code": -11, "message": "No permission for the resource" }
- The object does not exist.
This error may be observed when the URL is not correct. In this case, it should be:
"url": "/cli/global/system/admin/user”
Related articles:
Technical Tip: Using FortiManager API.
Technical Tip: Managing the JSON API call with Postman and how to delete, create and update an ADOM using it.
Technical Tip: Using Firmware Manager CLI and API.
Technical Tip: How to run a proxy API call from FortiManager to a managed FortiGate.
Technical Tip: API requests by using the API token authentication method with Postman.
