---
id: collect-261001-fortinet/fortinet/t5-forticloud-products-technical-tip-set-up-an-api-call-to-fortimanager-or-ta-p-b4c4039e
title: "t5-forticloud-products-technical-tip-set-up-an-api-call-to-fortimanager-or-ta-p--b4c4039e"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/t5-forticloud-products-technical-tip-set-up-an-api-call-to-fortimanager-or-ta-p--b4c4039e.md
source_anchor: ""
source_lines: [1, 38]
sha256: 66f570a92aa224ac373596b6a81f4805c8f4c307af3c4ee4fa6298c3735cab85
---

# t5-forticloud-products-technical-tip-set-up-an-api-call-to-fortimanager-or-ta-p--b4c4039e

Technical Tip: Set up an API call to FortiManager or FortiAnalyzer Cloud
Description
This article describes how to set up an API call to FortiManager Cloud or FortiAnalyzer Cloud.
Scope
FortiManager Cloud and FortiAnalyzer Cloud.
Solution
Open one of the following URLs depending on the platform:
After opening the corresponding page, locate the instance of choice. Select the API button at the bottom.
 
From the new panel that appears, select and copy the provided example of the JSON API.
To create an API user, go to Users.
Go to the IAM portal -> Permission Profiles -> Add New.
From Add Portal, choose FortiAnalyzer Cloud / FortiManager Cloud. Then enable access and the proper level of access.
Go to IAM portal -> User Groups -> Add Another Group.
Choose the name of the group and attach the proper permission profile under the IAM portal -> Users -> Add New -> API user.
Choose the proper user group for the user and then download the credentials.
To confirm that the API call is working correctly, use the following methods:
BASH script:
 
#!/bin/bashnn# IAM API usernurl_base="https://xxx.fortimanager.forticloud.com"nurl="$url_base/jsonrpc"nn# usernuser="6BD4E324xxx"npasswd="xxxxxx"nn# 1, get token from FortiCare, url is defined in admin portal Server Config -> API user setting -> FAC_OAUTH_API_URLntoken=$(curl --silent -k -X POST 
https://customerapiauth.fortinet.com/api/v1/oauth/token/ 
-H 'Content-Type: application/json' 
-d '{"username": "'$user'", "password": "'$passwd'", "client_id" : "FortiManager", "grant_type": "password" }' | jq -r '.access_token')necho 'access token is '$tokennn# 2, get session from instancensession=$(curl -k -X POST --http1.1 --silent $url_base/p/forticloud_jsonrpc_login/ 
-H 'Content-Type: application/json' 
-d '{"access_token": "'$token'"}' | jq -r '.session')nn# responsenecho 'session is '$sessionnnecho "******************System Status*************************"n# Get sys status from instancencurl -k -X POST --http1.1 $url 
-H 'Content-Type: application/json' 
-d '{"method": "get", "params": [{ "url": "/cli/global/system/status"}], "session": "'$session'", "id": 1}' | jq .
The output from the API call:
API Platform to send requests:
- Use the credentials and obtain an access token from the OAuth API URL 'https://customerapiauth.fortinet.com/api/v1/oauth/token/'.
{n	"content_type":	"application/json",n	"username":	"52C28D9C-XXXX-XXX-XXX-XXXX10F2186",n	"password":	"c60f4XXXXXXXXXXXXXXXXXXd2!1Aa",n	"client_id":	"FortiManager",n	"grant_type":	"password"n}
Note: The value for client_id can be 'FortiManager' or 'FortiAnalyzer', depending on the cloud service.
 
- Access tokens will be received in response to the above request (make sure the 'Headers' tab also has 'Content-Type' set to 'application/json').
- Using the access_token, send a login request to <FortiManager Cloud URL>/p/forticloud_jsonrpc_login/.
- The session key will then be received and can be used to send further requests.
Note: Steps are the same for both FortiManager and FortiAnalyzer Cloud.
Related documents:
