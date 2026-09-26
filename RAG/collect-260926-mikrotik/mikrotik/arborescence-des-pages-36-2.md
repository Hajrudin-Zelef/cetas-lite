---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-36-2
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-36.md
source_anchor: ""
source_lines: [64, 93]
sha256: 470a1cd78c5e071ca0573c5419f2e5080f229b6ce19b19abe09896f88890c6f5
---

# Introduction

| Property | Description | 
|---|---|
| **dns-name** (*string* ; Default:**""** ) | DNS name of the HotSpot server. This is the DNS name used as the name of the HotSpot server (i.e., it appears as the location of the login page). This name will automatically be added as a static DNS entry in the DNS cache. | 
| **hotspot-address** (*IP* ; Default:**0.0.0.0** ) | IP address of HotSpot service. | 
| **html-directory** (*string* ; Default:**hotspot** ) | Directory name in which HotSpot HTML pages are stored (by default *hotspot* directory). It is possible to specify different directory with modified HTML pages. To change HotSpot login page, get HotSpot files from your router, change and upload them back to same location. Full path must be typed in html-directory field, including "/flash/(hotspot_dir)" | 
| **html-directory-override** (*string* ; Default:**none** ) | Alternative path for hotspot html files. It should be used only when customized hotspot html files are stored on external storage. | 
| **http-cookie-lifetime** (*time* ; Default:**3d** ) | HTTP cookie validity time, the option is related to *cookie* HotSpot login method | 
| **http-proxy** (*IP:Port* ; Default:**0.0.0.0:0** ) | Address and port of the proxy server for HotSpot service, when default value is used all request are resolved by the local /ip proxy | 
| **https-redirect** (*yes \| no* ; Default:**yes** ) | Whether to redirect unauthenticated user to hotspot login page, if user is visiting a https:// url. Since certificate domain name will mismatch, often this leads to errors, so you can set this parameter to "no" and all https requests will simply be rejected and user will have to visit a http page. | 
| **login-by** (*cookie\|http-chap\|http-pap\|https\|mac\|trial\|mac-cookie* ; Default:**http-chap, cookie** ) | Used HotSpot authentication method  | 
| **mac-auth-password** (*string* ; Default: ) | Used together with MAC authentication, field used to specify password for the users to be authenticated by their MAC addresses. The following option is required, when specific RADIUS server rejects authentication for the clients with blank password | 
| **name** (*string* ; Default: ) | Descriptive name of the profile | 
| **nas-port-type** (*string* ; Default:**wireless-802.11** ) | NAS-Port-Type value to be sent to RADIUS server, NAS-Port-Type values are described in the RADIUS RFC 2865. This optional value attribute indicates the type of the physical port of the HotSpot server. | 
| **radius-accounting** (*yes \| no* ; Default:**yes** ) | Send RADIUS server accounting information for each user, when yes is used | 
| **radius-default-domain** (*string* ; Default: ) | Default domain to use for RADIUS requests. Allows to use separate RADIUS server per */ip hotspot profile* . If used, same domain name should be specified under /radius domain value. | 
| **radius-interim-update** (*time \| received* ; Default:**received** ) | How often to send accounting updates . When *received* is set, interim-time is used from RADIUS server.**0s** is the same as*received* . | 
| **radius-location-name** (*string* ; Default: ) | RADIUS-Location-Id to be sent to RADIUS server. Used to identify location of the HotSpot server during the communication with RADIUS server. Value is optional and used together with RADIUS server. | 
| **radius-mac-format** (*"XX XX XX XX XX XX"\|XX:XX:XX:XX:XX:XX\|XXXXXX-XXXXXX\|XXXXXXXXXXXX\|XX-XX-XX-XX-XX-XX\|XXXX:XXXX:XXXX\|XXXXXX:XXXXXX* ; Default:**XX:XX:XX:XX:XX:XX** ) | Option to set format of user mac-address, that is sent to RADIUS server during AAA session. | 
| **rate-limit** (*string* ; Default:**""** ) | Rate limitation in form of **rx-rate[/tx-rate] [rx-burst-rate[/tx-burst-rate] [rx-burst-threshold[/tx-burst-threshold] [rx-burst-time[/tx-burst-time]]]] [priority] [rx-rate-min[/tx-rate-min]]** from the point of view of the router (so "rx" is client upload, and "tx" is client download). All rates should be numbers with optional 'k' (1,000s) or 'M' (1,000,000s). If tx-rate is not specified, rx-rate is as tx-rate too. Same goes for tx-burst-rate and tx-burst-threshold and tx-burst-time. If both rx-burst-threshold and tx-burst-threshold are not specified (but burst-rate is specified), rx-rate and tx-rate is used as burst thresholds. If both rx-burst-time and tx-burst-time are not specified, 1s is used as default. rx-rate-min and tx-rate min are the values of limit-at properties | 
| **smtp-server** (*IP* ; Default:**0.0.0.0** ) | SMTP server address to be used to redirect HotSpot users SMTP requests. | 
| **split-user-domain** (*yes \| no* ; Default:**no** ) | Split username from domain name when the username is given in "user@domain" or in "domain\user" format from RADIUS server | 
| **ssl-certificate** (*string \| none* ; Default:**none** ) | Name of the SSL certificate on the router to to use only for HTTPS authentication. | 
| **trial-uptime** (*time/time* ; Default:**30m/1d** ) | Used only with *trial* authentication method. First time value specifies, how long trial user identified by MAC address can use access to public networks without HotSpot authentication. Second time value specifies amount of time, that has to pass until user is allowed to use trial again. | 
| **trial-user-profile** (*string* ; Default:**default** ) | Specifies **hotspot user profile** for trial users. | 
| **use-radius** (*yes \| no* ; Default:**no** ) | Use RADIUS to authenticate HotSpot users. | 

# HotSpot User Profiles

User profile menu is used for common HotSpot client settings. Profiles are like User groups with the same set of settings, rate-limit, filter chain name, etc.

