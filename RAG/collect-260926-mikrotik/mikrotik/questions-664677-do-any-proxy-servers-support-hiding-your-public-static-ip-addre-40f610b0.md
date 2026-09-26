---
id: collect-260926-mikrotik/mikrotik/questions-664677-do-any-proxy-servers-support-hiding-your-public-static-ip-addre-40f610b0
title: "questions-664677-do-any-proxy-servers-support-hiding-your-public-static-ip-addre-40f610b0"
domain: mikrotik
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/dhcp-dns/questions-664677-do-any-proxy-servers-support-hiding-your-public-static-ip-addre-40f610b0.md
source_anchor: ""
source_lines: [1, 14]
sha256: 0f7c0d2b9f54a5de7772818c3261a43ba29b94997dd4630c7d91fc27c1cc9067
---

# questions-664677-do-any-proxy-servers-support-hiding-your-public-static-ip-addre-40f610b0

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I want to setup proxy server for my Mikrotik router. There is inbuilt web-proxy for Mikrotik router but I can extend this up to transparent proxy (kind of proxy server) only. We need High anonymity proxy so that we can hide our LAN static IPs (we don't have private IP) from outside Intruder/hackers.
And also I know I can setup NAT rule to hide our IP (only private IP not public/static IP) as per this link, but I can't hide static/public IP.
Essentially I want to hide our Public/Static IP (there is static/public IP for all systems in our company) from outside Internet. To achieve this I guess I need other software apart from Mikrotik router gateway setup.
Can anyone suggest me another software to achieve my requirement? I know about squid proxy but am not sure whether It can hide our static/public IP.
Note: we have assigned public/Static IP to all systems of our company since we have rights to access our company's system from anywhere by dedicated laptop(given by our company with more security) through VPN connection.
Firstly, there's nothing special about the private IP ranges; NAT will work just as well on any address, it's just that there's really no need for it with publicly-routable addresses.
For a proxy to use with MikroTik Router OS, I'd suggest Squid but probably any HTTP proxy server would work; if you're in a Microsoft environment you may prefer their proxy offering.
In terms of 'hiding' internal IPs, what I believe you're referring to is the HTTP X-Forwarded-For header. You should be able to configure most proxy servers to omit this header (it's not part of the standard anyway, hence the X- prefix). I'd strongly recommend testing this before deploying to production by connecting through the new proxy to a server you control and examining the HTTP headers the server receives.
