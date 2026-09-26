---
id: collect-260926-mikrotik/mikrotik/questions-974829-how-to-set-a-default-page-within-hotspot-mikrotik-to-be-open-au-1241c925
title: "questions-974829-how-to-set-a-default-page-within-hotspot-mikrotik-to-be-open-au-1241c925"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/dhcp-dns/questions-974829-how-to-set-a-default-page-within-hotspot-mikrotik-to-be-open-au-1241c925.md
source_anchor: ""
source_lines: [1, 13]
sha256: 79e8d8e376faa4c988fb25a32a656c763fbae798b3f9ecee5b852560b9b63ec2
---

# questions-974829-how-to-set-a-default-page-within-hotspot-mikrotik-to-be-open-au-1241c925

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I'm trying to set a Default HTML Page to be opened after wifi connects automatically.
Kinda like the Login page. I've tried writing a few javascript lines to make the redirect work but after redict, there is no content within my HTML.
Is it possible to set an HTML page instead of the login page to be opened directly after wifi connect with MikroTik?
When using the Mikrotik hotspot functionality the login page is displayed which is stored on the Mikrotik router. This page is called login.html and can be modified to your needs.
If you want to use an external site to display to your users you can modify login.html with http meta refresh to redirect to an external url. You may want to add this URL to the walled garden so it can be accessed.
When the page is displayed, the user is not logged in. A HTTP POST must be done with a username and password to login the user and access other sites. You can do this by adding a form to the external site or the login.html page to login the user with one click. Without the login any request will be redirected to the login.html page and/or your external site.
MikroTik calls this the Hotspot Gateway. You can begin the setup with /ip hotspot setup or from IP > Hotspot from Winbox. Check the wiki page for several more in-depth examples.
