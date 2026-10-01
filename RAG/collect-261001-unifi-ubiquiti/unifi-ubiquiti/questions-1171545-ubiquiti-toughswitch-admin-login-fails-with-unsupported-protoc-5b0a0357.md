---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1171545-ubiquiti-toughswitch-admin-login-fails-with-unsupported-protoc-5b0a0357
title: "questions-1171545-ubiquiti-toughswitch-admin-login-fails-with-unsupported-protoc-5b0a0357"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1171545-ubiquiti-toughswitch-admin-login-fails-with-unsupported-protoc-5b0a0357.md
source_anchor: ""
source_lines: [1, 16]
sha256: c114083ec45ac62b99584b772608f08b36aa7e57a30d320b20f8d037b44425e4
---

# questions-1171545-ubiquiti-toughswitch-admin-login-fails-with-unsupported-protoc-5b0a0357

I have a Ubiquiti toughswitch used to provide connection in my home office. When I try to connect to the admin login from iMac (Retina 5K, 27inch, 2015) running Chrome I get the error "Unsupported protocol The client and server don't support a common SSL protocol version or cipher suite." Question is how do I determine what SSL version / level Chrome is supporting ?
2 Answers 2
To add to the above answer (I dont have enough rep to comment) you can update the toughswitch using the link below and then it will be accessible via chrome etc.
https://dl.ui.com/firmwares/edgemax/EdgeSwitchXP/v2.1.0/SW.v2.1.0.142.210208.1325.bin
- 
        1Please provide some more information. It is a valid and official link. But I for one, would want to know that it will fix the issue.Rohit Gupta– Rohit Gupta2025-10-22 23:17:39 +00:00Commented Oct 22, 2025 at 23:17
- 
        Indeed. It might be useful to mention that the toughswitch name was rebranded in 2018, which is why the firmware name is found under the EdgeSwitch name and point to for example the release notes for the 2.2.1 firmware version: "Improvements: ... Updated Web Server SSL cipher list. ..."HBruijn– HBruijn2025-10-23 13:18:43 +00:00Commented Oct 23, 2025 at 13:18
The root of the error is that that web interface probably only supports TLS 1.0 or 1.1. Most browsers dropped support for those TLS versions already a couple of years ago. AFAIK neither Chrome nor Safari have a setting re-enable TLS 1.1 support.
In contrast other browsers like for example Firefox still allow you to manually grant an exemption to access legacy sites that don't support modern/strong TLS versions but only TLS 1.1 or TLS 1.0 - via security.tls.version.enable-deprecated = true in about:config
That is not site specific and should be reset to false after you're done.
Update As answered here : Once you can access the switch again you can then upgrade the firmware, which should bring the used TLS versions up to current standards. Note that the ToughSwitch line has been re-branded and your device is now called the “EdgeSwitch XP” and firmware should therefore be downloadable from https://ui.com/download/edgemax
Question is how do I determine what SSL version / level Chrome is supporting ?
For example Wikipedia provides a history here: https://en.wikipedia.org/wiki/Version_history_for_TLS/SSL_support_in_web_browsers
Qualys SSL labs, that many know and use to test the SSL/TLS server configuration of their (web) servers also has a browser test: https://clienttest.ssllabs.com:8443/ssltest/viewMyClient.html
For the latest Firefox with security.tls.version.enable-deprecated = true that results in
