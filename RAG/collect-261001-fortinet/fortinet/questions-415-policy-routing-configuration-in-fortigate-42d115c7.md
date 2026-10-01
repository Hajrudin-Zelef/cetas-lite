---
id: collect-261001-fortinet/fortinet/questions-415-policy-routing-configuration-in-fortigate-42d115c7
title: "questions-415-policy-routing-configuration-in-fortigate-42d115c7"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-fortinet/questions-415-policy-routing-configuration-in-fortigate-42d115c7.md
source_anchor: ""
source_lines: [1, 44]
sha256: dc0799943a8228c117a66dba64b425e98cc661087fec4cb43440312166ac82bf
---

# questions-415-policy-routing-configuration-in-fortigate-42d115c7

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
8
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have an scenario where a Fortigate firewall is used to separate internal networks from the Internet (FortiOS Version 4.0 MR3 patch 11). Right now there is a single Internet connection attached to the firewall and a default static route is used to get all Internet traffic through it. I would like to attach a second Internet connection to the firewall and then route only certain traffic through it, for example web browsing traffic.
For this setup, I keep the current static default route through the first link and then configure policy routing options in order to route traffic with destination port TCP/80 and TCP/443 through the second Internet link. As expected, policy routing is evaluated before routing table and all traffic destined to TCP/80 and TCP/443 is sent through to second link, including traffic between subnets directly connected to the Fortigate, what breaks communication between them.
In a Cisco environment I would adjust the ACL used to match traffic for policy routing, denying traffic between internal networks at the beginning of the ACL and adding a "permit any" statement at the end. However, I can not find the way to instruct the Fortigate to work in a similar manner.
Do you know how to make this scenario working with Fortigate?
Since policy routes are evaluated top-down, you can work around this limit by placing a more specific entry matching traffic from internal subnet A to internal subnet B.
However, this should be less than comfortable if you have many different networks attached to your internal interface.
In this case, I would recommend you a trick I once used:
since Fortigate devices ignore QoS marks, you should sign your "internet" packets on the firewall-facing port of your Cisco switch with a specific TOS and then use that mark in your policy-route.
"In case of a Fortinet firewall, its Policy Route: CLI version:
config router policy
edit 1
set input-device "port4"
set src 172.18.0.0 255.255.0.0
set dst 192.168.3.0 255.255.255.0
set protocol 6
set start-port 443
set end-port 443
set gateway 1.1.1.1
set output-device "port3"
next
end
For the GUI version, check the blog above. Can't post images until I get 10 rep points. :-/
Yes! you can make it possibly as per your requirement I understand that your want web -browsing traffic ie is tcp-80 & tcp-443 outbound traffic wants to flow from secondary internet link . And primary internet link will use remaining traffic other thàn web -traffic to configure this requirement .
policy - routing will have priority then default route as we know..
Create two outbound security policies for internet access for both isps
Assume
Port11 -first internet link
Port8 - second internet link
Port5 -Lan interface
Create policy for secondary internet link
Source interface : Lan Destination interface :port8 Source address : Any Destination :address :any Service : http &https Action : allowed Security profiles : on
Keep secondary internet link policy on top
Now create policy based route
Às source as : Lan subnet Destination :any Interface : Port 8 Action :allow
And then configure default route as
Ip route 0.0.0.0 0.0.0.0 pointing towards first isp gatewayip route 0.0.0.0 0.0.0.0 pointing towards second isp gateway
Let's see now traffic flow
Any user àccessing internet from LAN will first check policy based routing if ip matches packet will be send to policy of secondary link as per policy if traffic is 80 and 443 is allowed ànd other traffic is second on second policy that is first internet link policy .. in this you can configure your requirement all 80 and 443 traffic will flow from secondary link and remaining traffic will flow from first internet link unwanted traffic will be denied in implicitly deny policy at bottom... Like wise you can configure your requirement..
