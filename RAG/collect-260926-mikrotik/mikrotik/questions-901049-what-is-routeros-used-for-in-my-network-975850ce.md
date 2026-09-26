---
id: collect-260926-mikrotik/mikrotik/questions-901049-what-is-routeros-used-for-in-my-network-975850ce
title: "questions-901049-what-is-routeros-used-for-in-my-network-975850ce"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/questions-901049-what-is-routeros-used-for-in-my-network-975850ce.md
source_anchor: ""
source_lines: [1, 43]
sha256: 72f2558375e6e4e3722b0bc969b96465eb2dd7ae96760d89784242f0a018d62a
---

# questions-901049-what-is-routeros-used-for-in-my-network-975850ce

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
At http://192.168.0.2/ to http://192.168.0.7/, they show the configuration pages for different versions of RouterOS. The username and password are not the same as those I used to access my router's configuration page.
What is RouterOS used for in my network? Is it installed on my router? Is it the OS or firmware running on my router?
What are the default username and password to access its configuration webpage?
What does its configuration webpage allow me to do?
Thanks.
$ traceroute superuser.com
traceroute to superuser.com (190.93.247.58), 30 hops max, 60 byte packets
1 192.168.1.1 (192.168.1.1) 1.289 ms 1.788 ms 2.244 ms
2 10.235.28.1 (10.235.28.1) 6.717 ms 6.931 ms 7.084 ms
3 cci-209150226181.clarityconnect.net (209.150.226.181) 12.114 ms 12.481 ms 12.662 ms
4 cci-209150226245.clarityconnect.net (209.150.226.245) 13.747 ms 13.605 ms 13.940 ms
5 clarityconnect.fltg.net (98.159.210.49) 15.421 ms 15.630 ms 16.526 ms
6 xe-8-3-0.bar1.Syracuse2.Level3.net (4.26.24.33) 52.766 ms 32.621 ms 10.453 ms
7 * * *
8 telia-level3-10.NewYork1.Level3.net (4.68.110.82) 16.098 ms 20.152 ms 19.890 ms
9 nyk-bb1-link.telia.net (213.155.131.136) 20.449 ms 20.688 ms 21.085 ms
10 nyk-b2-link.telia.net (213.155.130.28) 21.393 ms 22.181 ms 22.412 ms
11 cloudflare-ic-301663-nyk-b2.c.telia.net (213.248.77.162) 22.662 ms 124.887 ms 16.149 ms
12 190.93.247.58 (190.93.247.58) 16.543 ms 17.175 ms 16.745 ms
$ traceroute 192.168.0.2
traceroute to 192.168.0.2 (192.168.0.2), 30 hops max, 60 byte packets
1 192.168.1.1 (192.168.1.1) 4.103 ms 7.418 ms 7.526 ms
2 10.235.28.1 (10.235.28.1) 10.167 ms 10.399 ms 10.614 ms
3 cci-209150226181.clarityconnect.net (209.150.226.181) 11.126 ms 11.338 ms 12.709 ms
4 cci-209150226142.clarityconnect.net (209.150.226.142) 15.422 ms 15.613 ms 15.855 ms
5 10.150.228.9 (10.150.228.9) 17.153 ms 16.759 ms 17.333 ms
6 10.150.228.1 (10.150.228.1) 25.243 ms 24.085 ms 25.122 ms
7 192.168.0.2 (192.168.0.2) 25.390 ms 25.574 ms 27.308 ms
If your router’s IP is 192.168.1.1 while the devices running RouterOS are indeed accessible via 192.168.0.2-7, they’re most definitely not in your network.
RouterOS is a MikroTik product and not related to your Linksys router in any way.
Update
From the traceroute output, we can see that your ISP, which provides wireless internet access, manages its internal network in the 10.0.0.0/8 range and then performs carrier-grade NAT. (This is common practice because the IPv4 address space has been fully exhausted and only small blocks remain for assignment.)
Your internal network is 192.168.1.0/24, that means all devices with IPs from 192.168.1.1 to 192.168.1.254 could belong to devices local to your network.
It’s very likely these MikroTik devices are used by your ISP to run its backbone network and/or wireless base stations.
Update 2
From the traceroute 192.168.0.2 output, we can discern these devices are definitely not local to your network or even anywhere near you. Because the path never leaves your ISP’s network, it’s likely networking equipment maintained by the ISP. It’s, as you call it, an “intermediate network”, although your regular internet traffic doesn’t go through it.
It’s rather questionably why they would route their traffic like that or use common home network subnets, but that’s a different matter. :)
RouterOS is a Linux distro made by MikroTik for their proprietary routing hardware.
My guess is that you have one or more MikroTik devices on your network. These were likely provided by your ISP. Since they are in a different subnet (192.168.0.0/24 vs 192.168.1.0/24), that usually indicates the router hardware that directly connects to the ISP network, like your cable/DSL modem.
The default username and password for RouterOS is admin and a blank password. However, the default IP for RouterOS is 192.168.88.1, so I believe it has probably been configured by your ISP, so the username/password are likely changed.
The reason you may be seeing multiple IP addresses is that the MikroTik box may have multiple interfaces. Unless you provide more details, there's not much more we can tell you at this point.
tracert superuser.comand include the output in your question.tracerton my Ubuntu. Which package to install?traceroute.traceroute 192.168.0.2and add the output, too.
