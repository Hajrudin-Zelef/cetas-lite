---
id: collect-260926-mikrotik/mikrotik/questions-1622238-mikrotik-hotspot-external-login-page-vs-radius-server-e92833a4
title: "questions-1622238-mikrotik-hotspot-external-login-page-vs-radius-server-e92833a4"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/dhcp-dns/questions-1622238-mikrotik-hotspot-external-login-page-vs-radius-server-e92833a4.md
source_anchor: ""
source_lines: [1, 2]
sha256: a141091a79bcf563da4c141e07ec4a3cf25c47b65593d1e1702bf18197440a9d
---

# questions-1622238-mikrotik-hotspot-external-login-page-vs-radius-server-e92833a4

I am fairly new to MikroTik RouterOS and I was wondering if the following logic is correct:
When configuring Hotspot feature of a MikroTik router with External login page as explained here, with the login page being a custom PHP/ASP/etc. hosted on the web, when user successfully authenticates on the said page, how does the web page then inform the MikroTik router about successful authentication? By setting a cookie? What I am trying to achieve is to have a website that hosts a separate database of usernames, and a login page on this website that the MikroTik router's Hotspot feature would redirect to when connection is established and authenticate the users there. I understand that using a RADIUS server would do the same but I am trying to build this using custom database of users.
