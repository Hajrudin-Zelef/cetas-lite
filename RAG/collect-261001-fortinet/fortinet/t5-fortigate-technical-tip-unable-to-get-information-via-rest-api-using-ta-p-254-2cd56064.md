---
id: collect-261001-fortinet/fortinet/t5-fortigate-technical-tip-unable-to-get-information-via-rest-api-using-ta-p-254-2cd56064
title: "t5-fortigate-technical-tip-unable-to-get-information-via-rest-api-using-ta-p-254-2cd56064"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/t5-fortigate-technical-tip-unable-to-get-information-via-rest-api-using-ta-p-254-2cd56064.md
source_anchor: ""
source_lines: [1, 4]
sha256: 5e413dc9ee321b7f2e138cf6b75dce9b033dbf2b9e5e90886a05bd7da680a830
---

# t5-fortigate-technical-tip-unable-to-get-information-via-rest-api-using-ta-p-254-2cd56064

Technical Tip: Unable to get information via REST-API using Python/CURL when post-login-banner enabled
| Description | This article describes the failure of administrator to obtain information on FortiGate via REST-API using Python/CURL when post-login-banner is enabled. | 
| Scope | FortiGate, REST-API. | 
| Solution | 1) FortiGate is configured with a pre-login-banner and post-login-banner: # config system global set pre-login-banner enable set post-login-banner enable end  An administrator account with super_admin_readonly is configured to obtain information about FortiGate via CURL/Python command:    2) To login to Fortigate via CURL, the following can be used: curl -k -i -Z POST https://<FortiGate_IP>/logincheck -d "username=<username>&secretkey=<password>" --dump-header headers.txt -c cookies.txt 3) It will be possible to see the response of the login as successful:    4) However, error 401 Unauthorized would appear if there is an attempt to obtain information on the FortiGate. An attempt to retrieve FortiGate system information with the following command will be executed: curl -k -i -X GET https://<FortiGate_IP>/api/v2/monitor/system/status -b headers.txt Output:    5) The admin session is recorded in FortiGate, however, there will be no log indicating that the user logged in:    6) When logging into the web GUI via HTTPS, it would be possible to log in without issue:   7) The root cause is that the authentication with CURL is not complete when post-login-banner is enabled in the global setting. 8) The following log entry indicates successful login that was performed in step 6:    9) As a solution, it will be necessary to disable post-login-banner from global setting: # config system global set post-login-banner disable end  10) Once the above has been disabled, it would be possible to retrieve information on the FortiGate via CURL GET option: curl -k -i -X GET https://<FortiGate_IP>/api/v2/monitor/system/status -b headers.txt    11) In the system event log, the login event will be observed accordingly:    Note: This is expected behavior as the administrator login would only be considered complete after acknowledging the post-login-banner if the respective is being configured. |
