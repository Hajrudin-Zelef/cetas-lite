---
id: collect-261001-fortinet/fortinet/index-php-document-fortigate-7-2-4-cli-reference-399620-config-webfilter-profile-c776463e-3
title: "index-php-document-fortigate-7-2-4-cli-reference-399620-config-webfilter-profile-c776463e"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/index-php-document-fortigate-7-2-4-cli-reference-399620-config-webfilter-profile-c776463e.md
source_anchor: ""
source_lines: [363, 492]
sha256: 7e478d1bb926b7bdab38a5d54cc8c63d20efdb1b890a380718494b6f910c7c49
---

# index-php-document-fortigate-7-2-4-cli-reference-399620-config-webfilter-profile-c776463e

| per-category | Per-category warnings. | 
| Option | Description | 
|---|---|
| session | After session ends. | 
| timeout | After timeout occurs. | 
config quota
| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| id | ID number. | integer | Minimum value: 0 Maximum value: 4294967295 | 0 | 
| category | FortiGuard categories to apply quota to (category action must be set to monitor). | user | Not Specified |  | 
| type | Quota type. | option | - | time | 
|  |  |  |  |  | 
| unit | Traffic quota unit of measurement. | option | - | MB | 
|  |  |  |  |  | 
| value | Traffic quota value. | integer | Minimum value: 1 Maximum value: 4294967295 | 1024 | 
| duration | Duration of quota. | user | Not Specified | 5m | 
| override-replacemsg | Override replacement message. | string | Maximum length: 28 |  | 
| Option | Description | 
|---|---|
| time | Use a time-based quota. | 
| traffic | Use a traffic-based quota. | 
| Option | Description | 
|---|---|
| B | Quota in bytes. | 
| KB | Quota in kilobytes. | 
| MB | Quota in megabytes. | 
| GB | Quota in gigabytes. | 
config override
| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| ovrd-cookie | Allow/deny browser-based (cookie) overrides. | option | - | deny | 
|  |  |  |  |  | 
| ovrd-scope | Override scope. | option | - | user | 
|  |  |  |  |  | 
| profile-type | Override profile type. | option | - | list | 
|  |  |  |  |  | 
| ovrd-dur-mode | Override duration mode. | option | - | constant | 
|  |  |  |  |  | 
| ovrd-dur | Override duration. | user | Not Specified | 15m | 
| profile-attribute | Profile attribute to retrieve from the RADIUS server. | option | - | Login-LAT-Service | 
|  |  |  |  |  | 
| ovrd-user-group <name> | User groups with permission to use the override. User group name. | string | Maximum length: 79 |  | 
| profile <name> | Web filter profile with permission to create overrides. Web profile. | string | Maximum length: 79 |  | 
| Option | Description | 
|---|---|
| allow | Allow browser-based (cookie) override. | 
| deny | Deny browser-based (cookie) override. | 
| Option | Description | 
|---|---|
| user | Override for the user. | 
| user-group | Override for the user's group. | 
| ip | Override for the initiating IP. | 
| browser | Create browser-based (cookie) override. | 
| ask | Prompt for scope when initiating an override. | 
| Option | Description | 
|---|---|
| list | Profile chosen from list. | 
| radius | Profile determined by RADIUS server. | 
| Option | Description | 
|---|---|
| constant | Constant mode. | 
| ask | Prompt for duration when initiating an override. | 
| Option | Description | 
|---|---|
| User-Name | Use this attribute. | 
| NAS-IP-Address | Use this attribute. | 
| Framed-IP-Address | Use this attribute. | 
| Framed-IP-Netmask | Use this attribute. | 
| Filter-Id | Use this attribute. | 
| Login-IP-Host | Use this attribute. | 
| Reply-Message | Use this attribute. | 
| Callback-Number | Use this attribute. | 
| Callback-Id | Use this attribute. | 
| Framed-Route | Use this attribute. | 
| Framed-IPX-Network | Use this attribute. | 
| Class | Use this attribute. | 
| Called-Station-Id | Use this attribute. | 
| Calling-Station-Id | Use this attribute. | 
| NAS-Identifier | Use this attribute. | 
| Proxy-State | Use this attribute. | 
| Login-LAT-Service | Use this attribute. | 
| Login-LAT-Node | Use this attribute. | 
| Login-LAT-Group | Use this attribute. | 
| Framed-AppleTalk-Zone | Use this attribute. | 
| Acct-Session-Id | Use this attribute. | 
| Acct-Multi-Session-Id | Use this attribute. | 
config web
| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| bword-threshold | Banned word score threshold. | integer | Minimum value: 0 Maximum value: 2147483647 | 10 | 
| bword-table | Banned word table ID. | integer | Minimum value: 0 Maximum value: 4294967295 | 0 | 
| urlfilter-table | URL filter table ID. | integer | Minimum value: 0 Maximum value: 4294967295 | 0 | 
| content-header-list | Content header list. | integer | Minimum value: 0 Maximum value: 4294967295 | 0 | 
| blocklist | Enable/disable automatic addition of URLs detected by FortiSandbox to blocklist. | option | - | disable | 
|  |  |  |  |  | 
| allowlist | FortiGuard allowlist settings. | option | - |  | 
|  |  |  |  |  | 
| safe-search | Safe search type. | option | - |  | 
|  |  |  |  |  | 
| youtube-restrict | YouTube EDU filter level. | option | - | none | 
|  |  |  |  |  | 
| vimeo-restrict | Set Vimeo-restrict ("7" = don't show mature content, "134" = don't show unrated and mature content). A value of cookie "content_rating". | string | Maximum length: 63 |  | 
| log-search | Enable/disable logging all search phrases. | option | - | disable | 
|  |  |  |  |  | 
| keyword-match <pattern> | Search keywords to log when match is found. Pattern/keyword to search for. | string | Maximum length: 79 | ** | 
| Option | Description | 
|---|---|
| enable | Enable setting. | 
| disable | Disable setting. | 
| Option | Description | 
|---|---|
| exempt-av | Exempt antivirus. | 
| exempt-webcontent | Exempt web content. | 
| exempt-activex-java-cookie | Exempt ActiveX-JAVA-Cookie. | 
| exempt-dlp | Exempt DLP. | 
| exempt-rangeblock | Exempt RangeBlock. | 
| extended-log-others | Support extended log. | 
| Option | Description | 
|---|---|
| url | Insert safe search string into URL. | 
| header | Insert safe search header. | 
| Option | Description | 
|---|---|
| none | Full access for YouTube. | 
| strict | Strict access for YouTube. | 
| moderate | Moderate access for YouTube. | 
| Option | Description | 
|---|---|
| enable | Enable setting. | 
| disable | Disable setting. |
