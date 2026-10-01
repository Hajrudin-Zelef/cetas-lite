---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1563064-how-to-open-ports-as-opposed-to-forwarding-them-on-ubiquiti-ed-5632d1d2
title: "questions-1563064-how-to-open-ports-as-opposed-to-forwarding-them-on-ubiquiti-ed-5632d1d2"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1563064-how-to-open-ports-as-opposed-to-forwarding-them-on-ubiquiti-ed-5632d1d2.md
source_anchor: ""
source_lines: [1, 8]
sha256: bf9b2eaae17491ac09a7cb27fe9b23830914bb99b21aa083890f0c65a29beb2b
---

# questions-1563064-how-to-open-ports-as-opposed-to-forwarding-them-on-ubiquiti-ed-5632d1d2

First off, A port is open because a process is listening on it, and there is a network pathway between the client and the service.
As such, unless you are talking about a service being hosted on the router itself, it cannot "open a port". The router can block what would otherwise be an open port (that's kind of its job in most home use-cases), but it cannot do anything more than allow traffic on the path.
Also note, forwarding a port to a IP that doesn't have a processes on that port does not cause a port to become open either. The thing that makes it "open" is that there is a service listening there. otherwise it doesn't matter if the port is made inaccessible by NAT or by the internal host itself; its closed either way, just from different perspectives.
Secondly, what people call port-forwarding, is technically called DNAT. what you describe is a DNAT scenario (as opposed to an SNAT scenario, where you are connecting outbound from the LAN; SNAT is uncommon in home networking, beyond the automatic mechanisms of stateful NAT).
So to make your service accept incoming connections from the internet, you need to have the three elements; a Process, a Port, and a Pathway.
- create a forwarding rule for each port, accepting traffic from the WAN, and forwarding to your internal host on 9001 ad 9002. EdgeOS has a reasonably good interface for creating forwarding rules.
- Allow traffic incoming on that port from the Internet in your Firewall. in EdgeOS terms, place a ruleset on the WAN interface with a direction of "IN" and a rule to accept New and Established-state traffic destined for the port.
- verify that the service is listening on 0.0.0.0 or another suitable network interface (127.0.0.1 will not work). you can check which interface your service is bound to with netstat -ntlup (linux) ornetstat -abno | findstr LISTENING (powershell as admin), and looking at the local IP address. if its 0.0.0.0 or an IP on your LAN, it should work.
