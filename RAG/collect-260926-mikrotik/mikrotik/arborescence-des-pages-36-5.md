---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-36-5
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-36.md
source_anchor: ""
source_lines: [265, 315]
sha256: e4d2974f310ed3f3a3651406919246cf2e3b490397d1243d9c70f02ea0969263
---

# Introduction

MAC cookie is a hotspot feature, designed to improve accessibility for smartphones, laptops and other mobile devices.

When MAC cookie feature is enabled (**login-by**=mac-cookie, **add-mac-cookie**=yes set in user profile), following actions are taken:

- **first successful login** . Mac cookie keeps record of username and password for the MAC address if there is only one host with such MAC. Cookie timeout is set to value equal to mac-cookie-timeout.
- **new host appears** . Hotspot checks if there is a mac cookie record for the MAC address and logs in host using recorded username and password. If there is more than one host with the same MAC address, user will not be logged in and MAC cookie record for this address will be deleted.

When **user logs out** mac cookie is removed in following cases:

- user-request - user clicked on logout button.
- admin-reset - disconnected from radius server or user is removed from hotspot active menu
- nas-request - traffic limit reached
- session-timeout

To debug problems with mac-cookies you will need to enable hotspot debug logs and look for reasons why mac-cookie login didn't work for certain host.

Reasons when mac cookie is removed by server:

- /ip hotspot cookie remove <x>
- /ip hotspot client remove <x>
- Radius server sends Disconnect-Request
- End-User has logged out him self via hotspot status page
- End user has reached his data cap ("traffic limit reached")
- Session-Timeout
- If mac-cookie login fails
- If server detects that in host table there is more than one entry with the same mac-address

# Using DHCP option to advertise HotSpot URL

Most devices, such as modern smartphones, do some kind of background checking to see if they are behind a captive portal. They do this by requesting a known webpage and comparing the contents of that page, to what they should be. If contents are different, the device assumes there is a login page and creates a popup with this login page.

This does not always happen, as this "known webpage" could be blocked, whitelisted, or not accessible in internal networks. To improve on this mechanism, RFC 7710 was created, allowing the HotSpot to inform all DHCP clients that they are behind a captive-portal device and that they will need to authenticate to get Internet access, regardless of what webpages they do or do not request.

This DHCP option field is enabled automatically, but only if the router has a DNS name configured and has a valid SSL certificate (so that the login page can be accessed over HTTPS). When these requirements are met, a special DHCP option will be sent, containing a link to `https://<dns-name-of-hotspot>/api`. This link contains information in JSON format, instructing the client device of the captive portal status, and the location of the login page.

Contents of `https://<dns-name-of-hotspot>/api` are as follows:

Some devices require venue-info URL as well, so you are free to modify the api.json file to your liking, just like any other hotspot files. It is located in the router files menu.

Important

If you have set up Hotspot before RouterOS v7.3 when RFC 7710 was implemented, you will have to use "Reset HTML" function, or manually add/edit the api.json file to have the above contents, for Hotspot detection to work.

# Load balancing with mangle on Hotspot server

In a multi-uplink Hotspot setup (for example, PCC), default routing often might fail to work as expected because the router prioritises mangle rules over the local destination check.

To fix this, ensure traffic destined for the Hotspot server is routed via the **local** table before the **mangle** chain is evaluated. You can achieve this by:

1. Adjusting the "`/routing settings"` (Routing table lookup).
2. Or, adding specific "`/ip route rule"` entries to force the Hotspot server's traffic to use the local table first (including HTTP 80 traffic which is used to detect the Hotspot server).
