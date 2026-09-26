---
id: collect-260926-mikrotik/mikrotik/best-scripts-for-firewall-and-router-protection
title: "best-scripts-for-firewall-and-router-protection"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/best-scripts-for-firewall-and-router-protection.md
source_anchor: ""
source_lines: [1, 46]
sha256: cbb8aa516e2364e0a1d9034fb15f90bcad6f7c7eb429e159c4fafa67e2352ca9
---

# best-scripts-for-firewall-and-router-protection

Be it a lesson that one should never use configs from internet, without understanding what they do. 

If you want something more secure than default firewall, good first step is to define why default firewall is not secure enough for you. There isn’t anything like one best firewall, only good or bad firewall for given purpose. Also there’s usually more than one way how to reach the goal. Sometimes there are some small functional differences, sometimes it’s just a matter of personal preferences.

For example, I upgraded one spare RB450 to current RouterOS 6.41.3 and default firewall config is the following.

Input (traffic to router itself):

```
/ip firewall filter
add action=accept chain=input comment="defconf: accept established,related,untracked" connection-state=established,related,untracked
add action=drop chain=input comment="defconf: drop invalid" connection-state=invalid
add action=accept chain=input comment="defconf: accept ICMP" protocol=icmp
add action=drop chain=input comment="defconf: drop all not coming from LAN" in-interface-list=!LAN
```

Forward (traffic going through router):

```
/ip firewall filter
add action=accept chain=forward comment="defconf: accept in ipsec policy" ipsec-policy=in,ipsec
add action=accept chain=forward comment="defconf: accept out ipsec policy" ipsec-policy=out,ipsec
add action=fasttrack-connection chain=forward comment="defconf: fasttrack" connection-state=established,related
add action=accept chain=forward comment="defconf: accept established,related, untracked" connection-state=established,related,untracked
add action=drop chain=forward comment="defconf: drop invalid" connection-state=invalid
add action=drop chain=forward comment="defconf:  drop all from WAN not DSTNATed" connection-nat-state=!dstnat connection-state=new in-interface-list=WAN
```

NAT config (just standard masquerade, nothing interesting):

```
/ip firewall nat
add action=masquerade chain=srcnat comment="defconf: masquerade" ipsec-policy=out,none out-interface-list=WAN
```

If you look at it, you can see the basic assumption that LAN is trusted, WAN isn’t. Other possible interfaces differ between input and forward.

Input accepts packets for established and related connections, it’s the good thing pretty much always. Also accepts untracked packets, but it doesn’t do anything by default, you first need to tell the router that some packets should not be tracked. Next is blocking of invalid packets (e.g. not part of any existing connection, or such that could be start of new connection), that’s also good for most non-advanced setups. Then it accepts ICMP, because usually accepting it will make your life easier more, than not accepting it makes life hard for evil hackers. Some people prefer to disable ICMP and then recoice when when some online scanner congratulates them for being “stealth”. Well, if they like it… Hopefully they won’t continue to do so with IPv6, because there it will break things for sure. If you know what you’re doing, you can do some sensible ICMP filtering, but I don’t think it’s worth it for most setups. Maybe if you’re connecting a nuclear facility or something.  Finally everything coming from somewhere else than LAN is dropped (LAN is defined as interface list elsewhere). So there really isn’t much to improve.

Forward accepts traffic from and to IPSec tunnels. It doesn’t do anything, if you don’t have any. And if you don’t plan to, you can remove these rules and save few CPU cycles. Fasttrack for established and related connections can speed things up and lower CPU usage. Next accept rule is same as for input, for stuff that can’t be fasttracked. Droping invalid connections is also the same. Finally everything from WAN is dropped, unless it’s forwarded port.

So as you can see, it’s also pretty secure. What I don’t like much is that everything else is allowed by default. E.g. if you connect to VPN, anything from there will be allowed to access LAN. On the other hand, connecting to VPN is extra step. If you don’t do it, there’s only LAN and WAN, nothing else. And it’s safe, because if you don’t forward any port inside, router won’t let anything pass from WAN to LAN. So again, nothing much to improve for simple setups.

Compare it with that loooong 4). I’m not saying it’s wrong. Some parts I’m alergic to, e.g. that “virus blocking” (try to find info about those names, some of them are no even from this millenium). Some ideas could be useful, e.g. to block traffic to unreachable private addresses leaking to WAN. But as whole it looks like it was made for slightly different purpose than simple home/office router. E.g. why bother with droppping bruteforcers, when you can simply not allow anything from WAN with one rule. Just because it’s longer doesn’t mean it’s better. 

As for the other links, 2) and 3) are written by MikroTik people, and they look ok. A lot of it is similar to default config. There are some extras, but it depends on your needs, what could be benefical for you. There are also some possibly useful ideas in 5), but again, it’s not good to just copy it without understanding.
