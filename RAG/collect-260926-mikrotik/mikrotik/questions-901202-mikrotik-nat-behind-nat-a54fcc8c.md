---
id: collect-260926-mikrotik/mikrotik/questions-901202-mikrotik-nat-behind-nat-a54fcc8c
title: "questions-901202-mikrotik-nat-behind-nat-a54fcc8c"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["optics"]
source: docs/RAG/lot-mikrotik/forum/firewall-nat/questions-901202-mikrotik-nat-behind-nat-a54fcc8c.md
source_anchor: ""
source_lines: [1, 11]
sha256: 4fc0912272ebf1b56bd79a6979f3e94da2daf978ce86a254e90da3868ae22395
---

# questions-901202-mikrotik-nat-behind-nat-a54fcc8c

I have two routers, one given by the internet provider, which I cannot control ( only they can ) and the second is my Mikrotik HAP ac (RB962UiGS-5HacT2HnT).
My router was primary before, and I had several ports forwarded. Now it's secondary, so I asked them to forward to the router's IP. However, that's not really working.
So my configuration is:
Outside router: 192.168.1.1 Inside router's static IP: 192.168.1.10
Outside port 7722 forwarded to 7722 of 192.168.1.10
Internal router (Mikrotik)'s IP: 192.168.88.1, my server's IP in the internal network is 192.268.88.25
On the internal router (Mikrotik) I have the following rules:
  chain=dstnat action=dst-nat to-addresses=192.168.88.25 to-ports=22 protocol=tcp dst-address=192.168.1.1 dst-port=7722
Before adding another router in the loop, dst-address was the external IP address and this worked very well.
However, it's not working now. Any idea what I'm missing?
I cannot use their router, as it does not support 5Ghz network and wireless is extremely slow. However, I cannot plugin their optics directly to the Mikrotik, as they want to be able to "control" it ...
