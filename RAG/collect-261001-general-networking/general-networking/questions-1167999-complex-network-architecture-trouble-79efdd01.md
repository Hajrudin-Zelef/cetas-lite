---
id: collect-261001-general-networking/general-networking/questions-1167999-complex-network-architecture-trouble-79efdd01
title: "Complex Network Architecture Trouble"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-1167999-complex-network-architecture-trouble-79efdd01.md
source_anchor: ""
source_lines: [1, 47]
sha256: f6312e5ce93329f7b94dd0ec2c8de4dbfe4251d56d7985be2c10d5ceafc221ad
---

# Complex Network Architecture Trouble

*Score : 1 | Source : https://serverfault.com/questions/1167999/complex-network-architecture-trouble*

I hope that someone can help me because it's very complicated/messy. I'll explain the situation.
I have the 2 locations: SiteA and SiteB. They are connected in VXLAN over IPSEC between them with 2 OPNSense firewalls.
The network they share is 192.168.180.0/24
I set MTU to 1400 on 6 Proxmox servers in the 2 locations (3 in SiteA and 3 in SiteB) and MTU to 1400 on the OPNSense interfaces (both the VXLAN and the LAN interface on 192.168.180.0/24).
On OPNSense I also made some firewall normalization rules on the VXLAN and LAN interface to bring the mss to 1250. I did various tests with iperf and with transfer scp creating 2 LXC containers, one in SiteA 192.168.180.150 and one in SiteB 192.168.180.151.
With iperf the performance is 400MB/s. If I do a scp between the 2 containers the speed is good.
The problem is this: I am at home with my notebook (I have a Fiber connection 2.5Gbps) I use OpenVPN to connect on the OPNSense of SiteA and I made some rules to make sure I also reach the servers in SiteB. The ping is good on both locations.
SiteA
$ ping 192.168.180.150
PING 192.168.180.150 (192.168.180.150) 56(84) bytes of data. 
64 bytes from 192.168.180.150: icmp_seq=1 ttl=63 time=13.6 ms 
64 bytes from 192.168.180.150: icmp_seq=2 ttl=63 time=14.0 ms 
SiteB
$ ping 192.168.180.151 (SiteB)
PING 192.168.180.151 (192.168.180.151) 56(84) bytes of data. 
64 bytes from 192.168.180.151: icmp_seq=1 ttl=62 time=17.7 ms 
64 bytes from 192.168.180.151: icmp_seq=2 ttl=62 time=17.7 ms 
64 bytes from 192.168.180.151: icmp_seq=3 ttl=62 time=18.9 ms 
If I do an iperf from my PC on 192.168.180.150 in SiteA I have a good result, with 20Mbit/s. Even if I do a scp of a 100MB file to 192.168.100.150, the performance is acceptable: 3MB/s.
The problems begins when I connect to the 192.168.180.151 in SiteB: iperf is very slow:
iperf -c 192.168.180.151 
[ 1] 0.0000-20.8590 sec 336 KBytes 132 Kbits/sec 
But the worst is scp which remains in stalled for several seconds before starting:
$ scp 100mb.img root@192.168.180.151:/root/ 
100mb.img 0% 0 0.0KB/s - stalled - 
After several seconds (also 1 minute) of it remaining like this, at a certain point the transfer starts with a rather good speed. At the beginning it starts with 100KB/s, but then it reaches 2.0MB/s
$ scp 100mb.img root@192.168.180.151:/root/ 100mb.img 79% 75MB 2.1MB/s 00:09 ETA 
How can I solve this strange behavior? Why is it so slow at the beginning and then at a certain point it improves?
I launch a tcpdump on my PC to see the traffic:
tcpdump -nn -i tun0 host 192.168.180.151
These are the last packets I see during the "stalling":
14:37:53.360278 IP 10.70.155.6.36894 > 192.168.180.151.22: Flags [P.], seq 104:156, ack 57, win 331, options [nop,nop,TS val 3667267481 ecr 4040040446], length 52
14:37:53.380470 IP 192.168.180.151.22 > 10.70.155.6.36894: Flags [P.], seq 57:85, ack 156, win 160, options [nop,nop,TS val 4040060442 ecr 3667267481], length 28
14:37:53.380537 IP 10.70.155.6.36894 > 192.168.180.151.22: Flags [.], ack 85, win 331, options [nop,nop,TS val 3667267502 ecr 4040060442], length 0
And whe it starts the transfer these are firsts packets:
14:38:13.363765 IP 10.70.155.6.36894 > 192.168.180.151.22: Flags [P.], seq 156:208, ack 85, win 331, options [nop,nop,TS val 3667287485 ecr 4040060442], length 52
14:38:13.384506 IP 192.168.180.151.22 > 10.70.155.6.36894: Flags [P.], seq 85:113, ack 208, win 160, options [nop,nop,TS val 4040080446 ecr 3667287485], length 28
14:38:13.384587 IP 10.70.155.6.36894 > 192.168.180.151.22: Flags [.], ack 113, win 331, options [nop,nop,TS val 3667287506 ecr 4040080446], length 0
14:38:14.437494 IP 10.70.155.6.57580 > 192.168.180.151.22: Flags [.], seq 66559:67733, ack 4885, win 83, options [nop,nop,TS val 3667288559 ecr 4040023913], length 1174
I'm pretty sure that is a fragmentation problem but I've already tried a lot of things playing with mtu/mss but I am not able to solve.
My issue is not related to my PC and scp. The "real" problem instead is that I've a lot of Windows clients in remote working that needs to connects to the applications with VPN. I've tried shared folder copy and winscp and both have this strange behaviour. Copy from Windows Share to Windows11 PC is OK, but if I try to copy a file from client to the shared folder it hangs and after a while it gives me an error.
Any help is really appreciated
Thanks a lot
