---
id: collect-261001-fortinet/fortinet/questions-58734-fortigate-enabling-directed-broadcast-to-broadcast-conversion-on-a153df0d-1
title: "questions-58734-fortigate-enabling-directed-broadcast-to-broadcast-conversion-on-a153df0d"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: ["2020-07-21"]
keywords: []
source: docs/RAG/collect-261001-fortinet/questions-58734-fortigate-enabling-directed-broadcast-to-broadcast-conversion-on-a153df0d.md
source_anchor: ""
source_lines: [1, 46]
sha256: cc7f24f1eb7c9a97c9d4f61129cfbf3c96ef5da00e90cef5cc05ef5ff88eff8f
---

# questions-58734-fortigate-enabling-directed-broadcast-to-broadcast-conversion-on-a153df0d

I'm not quite certain how to achieve the equivalent of ip directed broadcast with a FortiGate. In this case a FortiGate 60E with FortiOS 5.6.7. So I started to dig a little.
QUESTION:
Can anyone confirm that, on a FortiGate, set broadcast-forward enable on the egress interface does actually forward a directed broadcast packet to the given subnet as broadcast (as in: DstMAC ff:ff:ff:ff:ff:ff) out of that interface? Suitable firewall policies assumed to be in place, of course.
EDIT: That part of the question is answered: No, set broadcast-forward enable on the egress interface does not have this
desired effect.
Sideline Question: Is there another way to achieve this on a FortiGate?
EDIT 2020-07-21: Yes, it is possible. See "ADDON-2" below. It is based on Lukas' answer (see below).
Please note: I am perfectly familiar with ip directed-broacast <someACL> on Cisco routing gear, and I've successfully deployed WoL support many times with that.
Also note: I'm also not trying to make something like a broadcast-helper or WoL relay work on a FortiGate interface facing the WoL Magic Packet sending host. That host knows the remote subnet's directed broadcast address and sends to it.
I keep finding hints (such as next door on serverfault) that set broadcast-forward enable were to add support to have directed broadcasts forwarded as broadcasts in the attached subnet.
The documentation (or its equivalent for FortiOS 5.6) quoted with that has this to say:
ARP: by default, ARP broadcasts and ARP reply packets are flooded/forwarded on all ports or VLANs belonging to the same forwarding domain, without the need of firewall policies between the ports. This default behavior is necessary to allow the population of the FDB and allow further firewall policy lookup (see section Transparent mode Firewall processing for more details). This option is configurable at the interface settings level with the parameter arpforward (enabled by default).
Non-ARP: To forward non-ARP broadcasts, the following CLI command is used:
config system interface
edit "port2"
set broadcast-forward enable
next
end
BUT this quote is from the Networking in Transparent Mode section of the documentation (see --> Packet Forwarding --> Broadcast, Multicast, Unicast Forwarding), and we're not running transparent mode, here.
AND I do get the impression that set broadcast-forward enable is more an ingress thing than something for egress. The "best answer" in this thread on  the Fortinet community kind of confirms this gut feeling.
I have also read the FortiNet KB article, which is also being quoted and referenced elsewhere, but ... static ARP entries? Really? We have dozens of clients at that site! (Well, I could still add a static ARP entry for the directed broadcast address with ff:ff:ff:ff:ff:ff, but that seems somewhat wrong.)
Thanks for your answers, comments and pointers.
[ADDON-1]:
As suggested in zac67's answer, I tried with a multicast address, multicast policy, plus a narrow unicast policy (allowing source to directed-broadcast). None had the desired effect.
Also: set broadcast-forward enable on the egress interface has no effect. Packets get dropped upon ingress because of an ip forwarding check failure.
id=20085 trace_id=35 func=fw_local_in_handler line=402 msg="iprope_in_check() check failed on policy 0, drop"
Interestingly this happens despite the fact that the firewall does have a entry in the routing table mapping 192.168.10.255/32 to the correct egress interface.
RemoteFirewall # diag ip route list | grep 192.168.10
tab=255 vf=0 scope=253 type=3 proto=2 prio=0 0.0.0.0/0.0.0.0/0->192.168.10.0/32 pref=192.168.10.1 gwy=0.0.0.0 dev=8(internal1)
tab=255 vf=0 scope=254 type=2 proto=2 prio=0 0.0.0.0/0.0.0.0/0->192.168.10.1/32 pref=192.168.10.1 gwy=0.0.0.0 dev=8(internal1)
tab=255 vf=0 scope=253 type=3 proto=2 prio=0 0.0.0.0/0.0.0.0/0->192.168.10.255/32 pref=192.168.10.1 gwy=0.0.0.0 dev=8(internal1)
tab=254 vf=0 scope=253 type=1 proto=2 prio=0 0.0.0.0/0.0.0.0/0->192.168.10.0/24 pref=192.168.10.1 gwy=0.0.0.0 dev=8(internal1)
It is only with set broadcast-forward enable on the ingress interface (sic! Just don't get me started on the implications of this!) of the last hop Fortigate that I see a change in behaviour.
Please note: My tests were done with ICMP. Near the WoL sender, I only have access to systems that can send ICMP, not udp/9. Should be of no relevance, here.
filters=[host 192.168.115.1]
3.881184 WAN1 in 192.168.115.1 -> 192.168.10.255: icmp: echo request
3.881469 internal1 out 192.168.115.1 -> 192.168.10.255: icmp: echo request
3.882015 internal1 in 192.168.10.3 -> 192.168.115.1: icmp: echo reply
3.882016 internal1 in 192.168.10.2 -> 192.168.115.1: icmp: echo reply
3.882033 internal1 in 192.168.10.4 -> 192.168.115.1: icmp: echo reply
4.896689 WAN1 in 192.168.115.1 -> 192.168.10.255: icmp: echo request
4.896771 internal1 out 192.168.115.1 -> 192.168.10.255: icmp: echo request
4.897309 internal1 in 192.168.10.2 -> 192.168.115.1: icmp: echo reply
4.897325 internal1 in 192.168.10.3 -> 192.168.115.1: icmp: echo reply
4.899076 internal1 in 192.168.10.4 -> 192.168.115.1: icmp: echo reply
    
