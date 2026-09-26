---
id: collect-260926-mikrotik/mikrotik/rfc8910-captive-portal
title: "rfc8910-captive-portal"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/rfc8910-captive-portal.md
source_anchor: ""
source_lines: [1, 65]
sha256: 604896581b3f610aafe650d59d32006bbddbf2ec4b3ec4a1a3b5d7b4bc3cf451
---

# rfc8910-captive-portal

I know this has already been discussed, but has anyone had success?

This is what i have done for now.

We have valid certificate installed, and dhcp option 114 enabled.

I see requests for option 114 coming from the clients in wireshark

Option: (55) Parameter Request List

Length: 9

Parameter Request List Item: (1) Subnet Mask

Parameter Request List Item: (121) Classless Static Route

Parameter Request List Item: (3) Router

Parameter Request List Item: (6) Domain Name Server

Parameter Request List Item: (15) Domain Name

Parameter Request List Item: (108) Removed/Unassigned

Parameter Request List Item: (114) DHCP Captive-Portal

Parameter Request List Item: (119) Domain Search

Parameter Request List Item: (252) Private/Proxy autodiscovery

And the ACK looks like this:

Option: (114) DHCP Captive-Portal

Length: 38

Captive Portal: https://hotspot.xxxxxx.com/api

This is what i get when i enter this page, accesed via HTTPS.

{

“captive”: true,

“user-portal-url”: “https://hotspot.xxxxx.com/login”,

“can-extend-session”: true,

“info-venue-url”: “https://hotspot.xxxxxx.com”

}

The api.json is stored in hotspot folder.

But, I dont see any requests for the captive portal API, instead I see the old method (captive.apple.com - get-hotspot / Generate 204 +  302 redirections) happening.

Tested with Android 11, IOS 15, and Big Sur.

Any clue to go further? What more could I debug?

Is there anything more to do to make this happen? I read you have to make the client capable to reach oscp or ntp servers to validate your certificate, but how could I know if this is what is happening here?

Mikrotik says:

"*This DHCP option field is enabled automatically, but only if the router has a DNS name configured and has a valid SSL certificate (so that the login page can be accessed over HTTPS). When these requirements are met, a special DHCP option will be sent, containing a link to https:///api. This link contains information in JSON format, instructing the client device of the captive portal status, and the location of the login page."*
