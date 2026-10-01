---
id: collect-261001-general-networking/general-networking/questions-73850-i-cant-ping-the-gateway-ip-fe801-from-the-internal-port-in-my-fo-f287e713
title: "questions-73850-i-cant-ping-the-gateway-ip-fe801-from-the-internal-port-in-my-fo-f287e713"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-general-networking/questions-73850-i-cant-ping-the-gateway-ip-fe801-from-the-internal-port-in-my-fo-f287e713.md
source_anchor: ""
source_lines: [1, 18]
sha256: b2dcce321f9d070450ee1187d76772a62d0c70609704288b84a15dae59a78d50
---

# questions-73850-i-cant-ping-the-gateway-ip-fe801-from-the-internal-port-in-my-fo-f287e713

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
4
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I am using a fortigate 60f firewall. I have configured it for both ipv4 and ipv6. I have used the SD wan (by adding wan1 and wan2) for load-balancing. Moreover, I have configure a default static route for both ipv4 and ipv6 (for the SD-wan). Further more I have configured ipvr4 and ipv6 policies too. This network works perfectly fine for the ipv4 configurations. Only my wan2 has ipv6 connectivity and the gateway is fe80::1. As I mentioned above I have configure a default static route(gateway fe80::1), policies(from internal => sdwan) and sd wan rules(To send all the v6 traffic through wan2) for ipv6. More over I have configured my internal interface ipv6 address as fd00:1/64 and host(PC) connected to the internal as fd00::2/64. My wan2 interface ipv6 address is fe80::2/64. ***
Still I can't ping the gateway (fe80::1). Why is that? Where have I made the mistake?**
First, your ISP will simply drop any packets with ULA addresses (fc00::/7, which includes packets your fd00::/64 network), so you cannot ping to your ISP from your ULA network.
Next, you are misusing the ULA addressing. You can locally assign addresses in the fd00::/8 range, but the next 40 bits after the fdmust be randomly chosen to give you a /48 prefix from which you can get 65,536 /64 networks for your use. Obviously, your fd00::/64 is not using the required random 40-bit Global ID.
Also, you can only ping a Link-Local address in the Link-Local range (fe80::/10) from the same link of the address you are trying to ping Traffic to and from Link-Local addresses cannot leave the link (be routed). All IPv6 interfaces use the same Link-Local network, so you must use a Zone ID (host OS-specific) when using Link-Local addresses. For example, on a Windows PC, you would ping fe80::1%2 or something like ping fe80::1%eth0 on a Linux host.
I can't ping the gateway IP (fe80::1) from the internal port in my
fortigate 60f firewall
That would be correct because you are trying to ping a Link-Local address from a different Link, and you cannot do that. You are also wanting your ISP to respond to a ULA addressed host, but the ISP should properly drop any packets with ULA addresses.
Still I can't ping the gateway (fe80::1). Why is that? Where have I
made the mistake?
I explained why, and I have pointed out your mistakes. You will never be able to ping a Link-Local address from a different link, and your ISP would drop any ULA traffic, anyway.
ping-options?
