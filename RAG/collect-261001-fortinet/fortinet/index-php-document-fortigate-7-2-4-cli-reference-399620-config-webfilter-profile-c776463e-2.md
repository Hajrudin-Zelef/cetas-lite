---
id: collect-261001-fortinet/fortinet/index-php-document-fortigate-7-2-4-cli-reference-399620-config-webfilter-profile-c776463e-2
title: "index-php-document-fortigate-7-2-4-cli-reference-399620-config-webfilter-profile-c776463e"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-fortinet/index-php-document-fortigate-7-2-4-cli-reference-399620-config-webfilter-profile-c776463e.md
source_anchor: ""
source_lines: [179, 362]
sha256: 8c4224bd2d026ad8afa752cff4b82ec5204ba7a72e614e4ebf6eae55f46e888c
---

# index-php-document-fortigate-7-2-4-cli-reference-399620-config-webfilter-profile-c776463e

| wisp-algorithm | WISP server selection algorithm. | option | - | auto-learning | 
|  |  |  |  |  | 
| wisp-servers <name> | WISP servers. Server name. | string | Maximum length: 79 |  | 
| Option | Description | 
|---|---|
| enable | Enable setting. | 
| disable | Disable setting. | 
| Option | Description | 
|---|---|
| flow | Flow feature set. | 
| proxy | Proxy feature set. | 
| Option | Description | 
|---|---|
| enable | Enable setting. | 
| disable | Disable setting. | 
| Option | Description | 
|---|---|
| activexfilter | ActiveX filter. | 
| cookiefilter | Cookie filter. | 
| javafilter | Java applet filter. | 
| block-invalid-url | Block sessions contained an invalid domain name. | 
| jscript | Javascript block. | 
| js | JS block. | 
| vbs | VB script block. | 
| unknown | Unknown script block. | 
| intrinsic | Intrinsic script block. | 
| wf-referer | Referring block. | 
| wf-cookie | Cookie block. | 
| per-user-bal | Per-user block/allow list filter | 
| Option | Description | 
|---|---|
| bannedword-override | Banned word override. | 
| urlfilter-override | URL filter override. | 
| fortiguard-wf-override | FortiGuard Web Filter override. | 
| contenttype-check-override | Content-type header override. | 
| Option | Description | 
|---|---|
| normal | Normal, POST requests are allowed. | 
| block | POST requests are blocked. | 
| Option | Description | 
|---|---|
| enable | Enable setting. | 
| disable | Disable setting. | 
| Option | Description | 
|---|---|
| enable | Enable web proxy WISP. | 
| disable | Disable web proxy WISP. | 
| Option | Description | 
|---|---|
| primary-secondary | Select the first healthy server in order. | 
| round-robin | Select the next healthy server. | 
| auto-learning | Select the lightest loading healthy server. | 
config antiphish
| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| status | Toggle AntiPhishing functionality. | option | - | disable | 
|  |  |  |  |  | 
| default-action | Action to be taken when there is no matching rule. | option | - | exempt | 
|  |  |  |  |  | 
| check-uri | Enable/disable checking of GET URI parameters for known credentials. | option | - | disable | 
|  |  |  |  |  | 
| check-basic-auth | Enable/disable checking of HTTP Basic Auth field for known credentials. | option | - | disable | 
|  |  |  |  |  | 
| check-username-only | Enable/disable username only matching of credentials. Action will be taken for valid usernames regardless of password validity. | option | - | disable | 
|  |  |  |  |  | 
| max-body-len | Maximum size of a POST body to check for credentials. | integer | Minimum value: 0 Maximum value: 4294967295 | 65536 | 
| authentication | Authentication methods. | option | - | domain-controller | 
|  |  |  |  |  | 
| domain-controller | Domain for which to verify received credentials against. | string | Maximum length: 63 |  | 
| ldap | LDAP server for which to verify received credentials against. | string | Maximum length: 63 |  | 
| Option | Description | 
|---|---|
| enable | Enable AntiPhishing functionality. | 
| disable | Disable AntiPhishing functionality. | 
| Option | Description | 
|---|---|
| exempt | Exempt requests from matching. | 
| log | Log all matched requests. | 
| block | Block all matched requests. | 
| Option | Description | 
|---|---|
| enable | Enable checking of GET URI for username and password fields. | 
| disable | Disable checking of GET URI for username and password fields. | 
| Option | Description | 
|---|---|
| enable | Enable checking of HTTP Basic Auth field for known credentials. | 
| disable | Disable checking of HTTP Basic Auth field for known credentials. | 
| Option | Description | 
|---|---|
| enable | Enable username only credential matches. | 
| disable | Disable username only credential matches. | 
| Option | Description | 
|---|---|
| domain-controller | Domain Controller to verify user credential. | 
| ldap | LDAP to verify user credential. | 
config inspection-entries
| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| name | Inspection target name. | string | Maximum length: 63 |  | 
| fortiguard-category | FortiGuard category to match. | user | Not Specified | 0 | 
| action | Action to be taken upon an AntiPhishing match. | option | - | exempt | 
|  |  |  |  |  | 
| Option | Description | 
|---|---|
| exempt | Exempt requests from matching. | 
| log | Log all matched requests. | 
| block | Block all matched requests. | 
config custom-patterns
| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| pattern | Target pattern. | string | Maximum length: 255 |  | 
| category | Category that the pattern matches. | option | - | username | 
|  |  |  |  |  | 
| type | Pattern will be treated either as a regex pattern or literal string. | option | - | regex | 
|  |  |  |  |  | 
| Option | Description | 
|---|---|
| username | Pattern matches username fields. | 
| password | Pattern matches password fields. | 
| Option | Description | 
|---|---|
| regex | Pattern will be treated as a regex pattern. | 
| literal | Pattern will be treated as a literal string. | 
config ftgd-wf
| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| options | Options for FortiGuard Web Filter. | option | - | ftgd-disable | 
|  |  |  |  |  | 
| exempt-quota | Do not stop quota for these categories. | user | Not Specified | 17 | 
| ovrd | Allow web filter profile overrides. | user | Not Specified |  | 
| max-quota-timeout | Maximum FortiGuard quota used by single page view in seconds (excludes streams). | integer | Minimum value: 1 Maximum value: 86400 | 300 | 
| rate-javascript-urls | Enable/disable rating JavaScript by URL. | option | - | enable | 
|  |  |  |  |  | 
| rate-css-urls | Enable/disable rating CSS by URL. | option | - | enable | 
|  |  |  |  |  | 
| rate-crl-urls | Enable/disable rating CRL by URL. | option | - | enable | 
|  |  |  |  |  | 
| Option | Description | 
|---|---|
| error-allow | Allow web pages with a rating error to pass through. | 
| rate-server-ip | Rate the server IP in addition to the domain name. | 
| connect-request-bypass | Bypass connection which has CONNECT request. | 
| ftgd-disable | Disable FortiGuard scanning. | 
| Option | Description | 
|---|---|
| disable | Disable rating JavaScript by URL. | 
| enable | Enable rating JavaScript by URL. | 
| Option | Description | 
|---|---|
| disable | Disable rating CSS by URL. | 
| enable | Enable rating CSS by URL. | 
| Option | Description | 
|---|---|
| disable | Disable rating CRL by URL. | 
| enable | Enable rating CRL by URL. | 
config filters
| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| id | ID number. | integer | Minimum value: 0 Maximum value: 255 | 0 | 
| category | Categories and groups the filter examines. | integer | Minimum value: 0 Maximum value: 255 | 0 | 
| action | Action to take for matches. | option | - | monitor | 
|  |  |  |  |  | 
| warn-duration | Duration of warnings. | user | Not Specified | 5m | 
| auth-usr-grp <name> | Groups with permission to authenticate. User group name. | string | Maximum length: 79 |  | 
| log | Enable/disable logging. | option | - | enable | 
|  |  |  |  |  | 
| override-replacemsg | Override replacement message. | string | Maximum length: 28 |  | 
| warning-prompt | Warning prompts in each category or each domain. | option | - | per-category | 
|  |  |  |  |  | 
| warning-duration-type | Re-display warning after closing browser or after a timeout. | option | - | timeout | 
|  |  |  |  |  | 
| Option | Description | 
|---|---|
| block | Block access. | 
| authenticate | Authenticate user before allowing access. | 
| monitor | Allow access while logging the action. | 
| warning | Allow access after warning the user. | 
| Option | Description | 
|---|---|
| enable | Enable setting. | 
| disable | Disable setting. | 
| Option | Description | 
|---|---|
| per-domain | Per-domain warnings. | 
