---
id: collect-260926-mikrotik/mikrotik/questions-1149022-how-to-redirect-webfig-to-https-on-mikrotik-22950720
title: "questions-1149022-how-to-redirect-webfig-to-https-on-mikrotik-22950720"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/misc/questions-1149022-how-to-redirect-webfig-to-https-on-mikrotik-22950720.md
source_anchor: ""
source_lines: [1, 12]
sha256: a39e541708b1d8f7f6a7d32243b27581d12ca5990294cfa15463c74c12202e57
---

# questions-1149022-how-to-redirect-webfig-to-https-on-mikrotik-22950720

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I've just enabled SSL for logging in the router administration on my Mikrotik router (IP -> Services -> www-ssl + added new self-signed SSL certificate). I would like the connection on port 80 to my router administration to be redirected to the SSL port 443.
I have tried creating new Firewall NAT rule:
ip firewall nat add chain=dstnat dst-address=192.168.1.1 protocol=tcp dst-port=80 action=redirect to-ports=443
But it doesn't work. What should I do to make it right?
HTTPS redirects work by having the HTTP web server sending a HTTP redirect. You cannot accomplish this using a firewall. Even if you made port 80 traffic go to port 443, the result would be equivalent to going to http://example.org:443/.
Be sure, that www-ssl is enabled and fully configured. After that your NAT rule should work since router is now using port 80 and the traffic to this port can reach NAT chain
