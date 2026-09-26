---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-36-3
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-36.md
source_anchor: ""
source_lines: [94, 154]
sha256: 19514bf77b9e6ef2ff90c99542e603e56cbda04e5521cee4464e8c616f57da73
---

# Introduction

| Property | Description | 
|---|---|
| **add-mac-cookie** (*yes\|no* ; Default:**yes** ) | Allows to add mac cookie for users. | 
| **address-list** (*string* ; Default: ) | Name of the address list in which users IP address will be added. Useful to mark traffic per user groups for queue tree configurations. | 
| **address-pool** (*string \|none* ; Default:**none** ) | IP pool name from which the user will get IP. When user has improper network settings configuration on the computer, HotSpot server makes translation and assigns correct IP address from the pool instead of incorrect one | 
| **advertise** (*yes \| no* ; Default:**no** ) | Enable forced advertisement popups. After certain interval specific web-page is being displayed for HotSpot users. Advertisement page might be blocked by browsers popup blockers. | 
| **advertise-interval** (*time[,time[,..]]* ; Default:**30m,10m** ) | Set of intervals between advertisement popups. After the list is done, the last value is used for all further advertisements, 10 minutes | 
| **advertise-timeout** (*time \| immediately \| never* ; Default:**1m** ) | How long advertisement is shown, before blocking network access for HotSpot client. Connection to Internet is not allowed, when advertisement is not shown. | 
| **advertise-url** (*string[,string[,..]]* ; Default: ) | List of URLs that is show for advertisement popups. After the last URL is used, list starts from the begining. | 
| **idle-timeout** (*time \| none* ; Default:**none** ) | Maximal period of inactivity for authorized HotSpot clients. Timer is counting, when there is no traffic coming from that client and going through the router, for example computer is switched off. User is logged out, dropped of the host list, the address used by the user is freed, when timeout is reached. | 
| **incoming-filter** (*string* ; Default: ) | Name of the firewall chain applied to incoming packets from the users of this profile, jump rule is required from built-in chain (input, forward, output) to chain=hotspot | 
| **incoming-packet-mark** (*string* ; Default: ) | Packet mark put on incoming packets from every user of this profile | 
| **keepalive-timeout** (*time \| none* ; Default: ) | Keepalive timeout for authorized HotSpot clients. Used to detect, that the computer of the client is alive and reachable. User is logged out, when timeout value is reached | 
| **mac-cookie-timeout** (*time* ; Default:**3d** ) | Selects mac-cookie timeout from last login or logout. | 
| **name** (*string* ; Default: ) | Descriptive name of the profile | 
| **on-login** (*string* ; Default:**""** ) | Script name to be executed, when user logs in to the HotSpot from the particular profile. It is possible to get username from internal **user** and**interface** variable. For example,*:log info "User $user logged in!"* . If hotspot is set on bridge interface, then**interface** variable will show bridge as actual interface unless**use-ip-firewall'** is set in bridge settings. List of available variables:  | 
| **on-logout** (*string* ; Default:**""** ) | Script name to be executed, when user logs out from the HotSpot.It is possible to get username from internal **user** and**interface** variable. For example,*:log info "User $user logged in!"* . If hotspot is set on bridge interface, then**interface** variable will show bridge as actual interface unless**use-ip-firewall** is set in bridge settings. List of available variables:  Starting with v6.34rc11 some additional variables are available:  | 
| **open-status-page** (*always \| http-login* ; Default:**always** ) | Option to show status page for user authenticated with mac login method. For example to show advertisement on status page (alogin.html)  | 
| **outgoing-filter** (*string* ; Default: ) | Name of the firewall chain applied to outgoing packets from the users of this profile, jump rule is required from built-in chain (input, forward, output) to chain=hotspot | 
| **outgoing-packet-mark** (*string* ; Default: ) | Packet mark put on outgoing packets from every user of this profile | 
| **rate-limit** (*string* ; Default:**""** ) | Simple dynamic queue is created for user, once it logs in to the HotSpot. Rate-limitation is configured in the following form **[rx-rate[/tx-rate] [rx-burst-rate[/tx-burst-rate] [rx-burst-threshold[/tx-burst-threshold] [rx-burst-time[/tx-burst-time] [priority] [rx-rate-min[/tx-rate-min]]]]** . For example, to set 1M download, 512k upload for the client, rate-limit=512k/1M | 
| **session-timeout** (*time* ; Default:**0s** ) | Allowed session time for client. After this time, the user is logged out unconditionally | 
| **shared-users** (*integer* ; Default:**1** ) | Allowed number of simultaneously logged in users with the same HotSpot username | 
| **status-autorefresh** (*time \| none* ; Default:**none** ) | HotSpot status page autorefresh interval | 
| **transparent-proxy** (*yes \|* ; Default:**yes** ) | Use transparent HTTP proxy for the authorized users of this profile | 

# HotSpot Users

This is the menu, where client's user/password information is actually added, additional configuration options for HotSpot users are configured here as well.

| Property | Description | 
|---|---|
| **address** (*IP* ; Default:**0.0.0.0** ) | IP address, when specified client will get the address from the HotSpot one-to-one NAT translations. Address does not restrict HotSpot login only from this address | 
| **comment** (*string* ; Default: ) | descriptive information for HotSpot user, it might be used for scripts to change parameters for specific clients | 
| **email** (*string* ; Default: ) | HotSpot client's e-mail, informational value for the HotSpot user | 
| **limit-bytes-in** (*integer* ; Default:**0** ) | Maximal amount of bytes that can be received from the user. User is disconnected from HotSpot after the limit is reached. | 
| **limit-bytes-out** (*integer* ; Default:**0** ) | Maximal amount of bytes that can be transmitted from the user. User is disconnected from HotSpot after the limit is reached. | 
| **limit-bytes-total** (*integer* ; Default:**0** ) | (limit-bytes-in+limit-bytes-out). User is disconnected from HotSpot after the limit is reached. | 
| **limit-uptime** (*time* ; Default:**0** ) | Uptime limit for the HotSpot client, user is disconnected from HotSpot as soon as uptime is reached. | 
| **mac-address** (*MAC* ; Default:**00:00:00:00:00:00** ) | Client is allowed to login only from the specified MAC-address. If value is *00:00:00:00:00:00* , any mac address is allowed. | 
| **name** (*string* ; Default: ) | HotSpot login page username, when MAC-address authentication is used name is configured as client's MAC-address | 
| **password** (*string* ; Default: )*sensitive* | User password | 
| **profile** (*string* ; Default:**default** ) | User profile configured in */ip hotspot user profile* | 
| **routes** (*string* ; Default: ) | Routes added to HotSpot gateway when client is connected. The route format **dst-address gateway metric** (for example,*192.168.1.0/24 192.168.0.1 1* ) | 
| **server** (*string \| all* ; Default:**all** ) | HotSpot server's name to which user is allowed login | 
| **otp-secret** (*string* ; Default: )*sensitive* | A one-time password token that is used for HotSpot user authorization, it could be used as separate "password" for HotSpot user authentication. | 

**Read-only proterties**

| Property | Description | 
|---|---|
| **bytes-in** (*integer* ) |  | 
| **bytes-out** (*integer* ) |  | 
| **packets-in** (*integer* ) |  | 
| **packets-out** (*integer* ) |  | 
| **uptime** (*time* ) |  | 

# HotSpot Active

HotSpot active menu shows all clients authenticated in HotSpot, the menu is informational (read-only) it is not possible to change anything here, except user can be logged out with the remove command.

