---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-80
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-80.md
source_anchor: ""
source_lines: [1, 45]
sha256: 496baadaf1fbfa6fc6e93fde48fa58fb9cca4bb9b34411bf41c5ce754f7a0652
---

# Introduction

Connection Rate is a firewall matcher that allows capturing traffic based on the present speed of the connection.

# Theory

Each entry in the connection tracking table represents bidirectional communication. Every time a packet gets associated with a particular entry, the packet size value (including the IP header) is added to the "connection-bytes" value for this entry. (in other words "connection-bytes" includes both - upload and download).

Connection Rate calculates the speed of connection based on the change of "connection-bytes". The connection rate is recalculated every second and does not have any averages.

Both options "connection-bytes" and "connection-rate" work only with TCP and UDP traffic. (you need to specify a protocol to activate these options). In the "connection-rate" you can specify a range of speed that you like to capture:

## Rule Example

These rules will capture TCP/UDP traffic that was going through the router when the connection speed was below 100kbps:

# Application Example - Traffic Prioritization

Connection-rate can be used in various ways, that still need to be realized, but the most common setup will be to detect and set lower priorities to the "heavy connections" (connections that maintain a fast rate for long periods (such as P2P, HTTP, FTP downloads). By doing this you can prioritize all other traffic that usually includes VoIP and HTTP browsing and online gaming.

The method described in this example can be used together with other ways to detect and prioritize traffic. As the connection-rate option does not have any averages we need to determine what will be the margin that identifies "heavy connections". If we assume that a normal HTTP browsing connection is less than 500kB (4Mb) long and VoIP requires no more than 200kbps speed, then every connection that after the first 500kB still has more than 200kbps speed can be assumed as "heavy".

(You might have different "connection-bytes" for HTTP browsing and different "connection-rate" for VoIP in your network - so, please, do your own research before applying this example)

For this example, let's assume that we have a 6Mbps upload and download connection to ISP.

## Quick Start for Impatient

### Explanation

In mangle, we need to separate all connections into two groups, and then mark packets from their 2 groups. As we are talking about client traffic most logical place for marking would be the mangle chain forward.

Keep in mind that as soon as a "heavy" connection will have lower priority and queue will hit max-limit - heavy connection will drop speed, and connection-rate will be lower. This will result in a change to a higher priority and the connection will be able to get more traffic for a short while, when again connection-rate will rise and that again will result in a change to lower priority). To avoid this we must make sure that once detected "heavy connections" will remain marked as "heavy connections" for all times.

### IP Firewall mangle

This rule will ensure that that "heavy" connections will remain heavy". and mark the rest of the connections with the default connection mark:

These two rules will mark all heavy connections based on our standards, that every connection that after the first 500kB still has more than 200kbps speed can be assumed as "heavy":

The last two rules in mangle will simply mark all traffic from corresponding connections:

### Queue

This is a simple queue tree that is placed on the Interface HTB - "public" is an interface where your ISP is connected, and "local" is where are your clients. If you have more than 1 "public" or more than 1 "local" you will need to mangle upload and download separately and place the queue tree in global-out:
