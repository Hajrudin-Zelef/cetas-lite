---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-21
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-21.md
source_anchor: ""
source_lines: [1, 11]
sha256: a50f1ad8aa0c813b40f89e3863bc3dd063efdbeae0cdd7efc44eba3963f85849
---

# Overview

**Sub-menu:** `/tool traceroute`

**Traceroute** displays the list of the routers that packet travels through to get to a remote host. The **traceroute** or **tracepath** tool is available on practically all Unix-like operating systems and **tracert** on Microsoft Windows operating systems.

Traceroute operation is based on TTL value and ICMP “Time Exceeded” message. Remember that TTL value in IP header is used to avoid routing loops. Each hop decrements TTL value by 1. If the TTL reaches zero, the packet is discarded and ICMP Time Exceeded message is sent back to the sender when this occurs.

Initially by traceroute, the TTL value is set to 1 when next router finds a packet with TTL = 1 it sets TTL value to zero, and responds with an ICMP "time exceeded" message to the source. This message lets the source know that the packet traverses that particular router as a hop. Next time TTL value is incremented by 1 and so on. Typically, each router in the path towards the destination decrements the TTL field by one unit TTL reaches zero.

Using this command you can see how packets travel through the network and where it may fail or slow down. Using this information you can determine the computer, router, switch or other network device that possibly causing network issues or failures.
