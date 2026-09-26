---
id: collect-260926-mikrotik/mikrotik/questions-1030329-why-is-port-forwarding-in-mikrotik-routeros-stuck-at-syn-recv-b6b9cb32
title: "questions-1030329-why-is-port-forwarding-in-mikrotik-routeros-stuck-at-syn-recv-b6b9cb32"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/questions-1030329-why-is-port-forwarding-in-mikrotik-routeros-stuck-at-syn-recv-b6b9cb32.md
source_anchor: ""
source_lines: [1, 24]
sha256: 3e9f875b20568cf5f832a1c4389a8d8291b68ee3d33774590af05749c95348bb
---

# questions-1030329-why-is-port-forwarding-in-mikrotik-routeros-stuck-at-syn-recv-b6b9cb32

I'd like to set up port forwarding of tcp port 8000 -> 192.168.1.16:4200 on my Mikrotik RouterOS.
I've done the following:
/ip firewall nat add dstnat chain=dstnat action=dst-nat to-addresses=192.168.1.16 to-ports=4200 protocol=tcp dst-address=<PUBLIC_IP> dst-port=8000
When I try to use the service from the Internet then the following command just hangs:
curl <PUBLIC_IP>:8000
I can see the counters moving on the Mikrotik's NAT rule (via WebBox).
On the target machine, I can see the following in netstat -an | grep 4200:
tcp        0      0 0.0.0.0:4200            0.0.0.0:*               LISTEN
tcp        0      0 192.168.1.16:4200       <REMOTE_HOST>:37720     SYN_RECV
I verified that I am able to connect to the machine locally via curl 192.168.1.16:4200.
I can't figure out what can be wrong :(
UPDATE: Firewall filter rules:
/ip firewall filter
add action=accept chain=input comment="defconf: accept established,related,untracked" connection-state=established,related,untracked
add action=drop chain=input comment="defconf: drop invalid" connection-state=invalid
add action=accept chain=input comment="defconf: accept ICMP" protocol=icmp
add action=drop chain=input comment="defconf: drop all not coming from LAN" in-interface-list=!LAN
add action=accept chain=forward comment="defconf: accept in ipsec policy" ipsec-policy=in,ipsec
add action=accept chain=forward comment="defconf: accept out ipsec policy" ipsec-policy=out,ipsec
add action=fasttrack-connection chain=forward comment="fasttrack - except for ipsec" connection-mark=!ipsec connection-state=established,related
add action=accept chain=forward comment="defconf: accept established,related, untracked" connection-state=established,related,untracked
add action=drop chain=forward comment="defconf: drop invalid" connection-state=invalid
add action=drop chain=forward comment="defconf: drop all from WAN not DSTNATed" connection-nat-state=!dstnat connection-state=new in-interface-list=WAN
fasttrack-connectionis before it. To test, I'd try changingdst-address=<PUBLIC_IP>toin-interface-list=WANand start disabling other rules temporarily./ip firewall manglerule which broke some of my outgoing connections and the hosts I tried to verify my NAT connection were all involved. Sorry guys, the setup above is working indeed. Not sure, should I answer my question?
