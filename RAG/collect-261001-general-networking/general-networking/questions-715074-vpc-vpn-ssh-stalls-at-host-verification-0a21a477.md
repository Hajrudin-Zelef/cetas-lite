---
id: collect-261001-general-networking/general-networking/questions-715074-vpc-vpn-ssh-stalls-at-host-verification-0a21a477
title: "questions-715074-vpc-vpn-ssh-stalls-at-host-verification-0a21a477"
domain: general-networking
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-general-networking/questions-715074-vpc-vpn-ssh-stalls-at-host-verification-0a21a477.md
source_anchor: ""
source_lines: [1, 68]
sha256: 3611ddd6c364f8b8fe5e9b267e51e4bea3b7fcf905f749d18443661494c30b7b
---

# questions-715074-vpc-vpn-ssh-stalls-at-host-verification-0a21a477

I have set up a VPN tunnel to our VPC on Amazon which consists of a public subnet with a couple machines, one NAT instance, and then several private subnets with an assortment of machines as well. I can see the machines in the VPC from our local network, Ive added ICMP and SSH from our local network IP range to the VPC security group and can ping/nmap the private IPs of the servers just fine (seeing port 22 is open.)
I can SSH into the public subnet instances that have public IPs with no problem. The issue I'm having is SSH'ing into any of the machines from the private/public subnet private IPs.
SSH spits out:
   ssh -v -i ~/.ssh/redacted.pem ec2-user@11.0.0.139
   OpenSSH_6.6, OpenSSL 1.0.2d 9 Jul 2015
   debug1: Reading configuration data /etc/ssh/ssh_config
   debug1: Connecting to 11.0.0.139 [11.0.0.139] port 22.
   debug1: Connection established.
   debug1: identity file /home/username/.ssh/redacted.pem type -1
   debug1: identity file /home/username/.ssh/redacted.pem-cert type -1
   debug1: Enabling compatibility mode for protocol 2.0
   debug1: Local version string SSH-2.0-OpenSSH_6.6p1-hpn14v5
   debug1: Remote protocol version 2.0, remote software version OpenSSH_6.2
   debug1: Remote is NON-HPN aware
   debug1: match: OpenSSH_6.2 pat OpenSSH* compat 0x14000000
   debug1: SSH2_MSG_KEXINIT sent
   debug1: SSH2_MSG_KEXINIT received
   debug1: AUTH STATE IS 0
   debug1: REQUESTED ENC.NAME is 'aes128-ctr'
   debug1: kex: server->client aes128-ctr hmac-md5-etm@openssh.com none
   debug1: REQUESTED ENC.NAME is 'aes128-ctr'
   debug1: kex: client->server aes128-ctr hmac-md5-etm@openssh.com none
   debug1: sending SSH2_MSG_KEX_ECDH_INIT
   debug1: expecting SSH2_MSG_KEX_ECDH_REPLY
At which point it sits for a while, finally returning:
   Connection closed by 11.0.0.139
The VPN is set up with static routing, and the tunnel is up. Route tables are propegating our local 192.18.0.0/16 network IP range through the Virtual Private Gateway. Networking ACLs are completely open IN and OUT.
Yet I'm not sure why I dont get a host verification step back from the server?
tcpdump shows this: 
   > sudo tcpdump -vvv 'tcp port 22 and host 11.0.0.139'
   tcpdump: listening on enp0s20u1, link-type EN10MB (Ethernet), capture size 262144 bytes
   12:36:26.096254 IP (tos 0x0, ttl 64, id 56385, offset 0, flags [DF], proto TCP (6), length 60)
       HOSTNAME.55391 > 11.0.0.139.ssh: Flags [S], cksum 0xb16e (correct), seq 989731254, win 29200, options [mss 1460,sackOK,TS val 3130871 ecr 0,nop,wscale 7], length 0
   12:36:26.141295 IP (tos 0x0, ttl 62, id 0, offset 0, flags [DF], proto TCP (6), length 60)
       11.0.0.139.ssh > HOSTNAME.55391: Flags [S.], cksum 0x3b34 (correct), seq 3868019661, ack 989731255, win 26847, options [mss 8961,sackOK,TS val 1060770 ecr 3130871,nop,wscale 7], length 0
   12:36:26.141372 IP (tos 0x0, ttl 64, id 56386, offset 0, flags [DF], proto TCP (6), length 52)
       HOSTNAME.55391 > 11.0.0.139.ssh: Flags [.], cksum 0xef39 (correct), seq 1, ack 1, win 229, options [nop,nop,TS val 3130885 ecr 1060770], length 0
   12:36:26.141828 IP (tos 0x0, ttl 64, id 56387, offset 0, flags [DF], proto TCP (6), length 83)
       HOSTNAME.55391 > 11.0.0.139.ssh: Flags [P.], cksum 0xb8b4 (correct), seq 1:32, ack 1, win 229, options [nop,nop,TS val 3130885 ecr 1060770], length 31
   12:36:26.185005 IP (tos 0x0, ttl 62, id 25699, offset 0, flags [DF], proto TCP (6), length 52)
       11.0.0.139.ssh > HOSTNAME.55391: Flags [.], cksum 0xef22 (correct), seq 1, ack 32, win 210, options [nop,nop,TS val 1060781 ecr 3130885], length 0
   12:36:26.188702 IP (tos 0x0, ttl 62, id 25700, offset 0, flags [DF], proto TCP (6), length 73)
       11.0.0.139.ssh > HOSTNAME.55391: Flags [P.], cksum 0x2e5c (correct), seq 1:22, ack 32, win 210, options [nop,nop,TS val 1060782 ecr 3130885], length 21
   12:36:26.188813 IP (tos 0x0, ttl 64, id 56388, offset 0, flags [DF], proto TCP (6), length 52)
       HOSTNAME.55391 > 11.0.0.139.ssh: Flags [.], cksum 0xeeeb (correct), seq 32, ack 22, win 229, options [nop,nop,TS val 3130899 ecr 1060782], length 0
   12:36:26.483445 IP (tos 0x0, ttl 64, id 56395, offset 0, flags [DF], proto TCP (6), length 1438)
       HOSTNAME.55391 > 11.0.0.139.ssh: Flags [.], cksum 0xdabd (correct), seq 32:1418, ack 1566, win 251, options [nop,nop,TS val 3130988 ecr 1060793], length 1386
   12:36:26.976751 IP (tos 0x0, ttl 64, id 56396, offset 0, flags [DF], proto TCP (6), length 1438)
       HOSTNAME.55391 > 11.0.0.139.ssh: Flags [.], cksum 0xda29 (correct), seq 32:1418, ack 1566, win 251, options [nop,nop,TS val 3131136 ecr 1060793], length 1386
   12:36:27.966780 IP (tos 0x0, ttl 64, id 56397, offset 0, flags [DF], proto TCP (6), length 1438)
       HOSTNAME.55391 > 11.0.0.139.ssh: Flags [.], cksum 0xd900 (correct), seq 32:1418, ack 1566, win 251, options [nop,nop,TS val 3131433 ecr 1060793], length 1386
   12:36:29.943400 IP (tos 0x0, ttl 64, id 56398, offset 0, flags [DF], proto TCP (6), length 1438)
       HOSTNAME.55391 > 11.0.0.139.ssh: Flags [.], cksum 0xd6af (correct), seq 32:1418, ack 1566, win 251, options [nop,nop,TS val 3132026 ecr 1060793], length 1386
   12:36:33.896741 IP (tos 0x0, ttl 64, id 56399, offset 0, flags [DF], proto TCP (6), length 1438)
       HOSTNAME.55391 > 11.0.0.139.ssh: Flags [.], cksum 0xd20d (correct), seq 32:1418, ack 1566, win 251, options [nop,nop,TS val 3133212 ecr 1060793], length 1386
   12:36:41.803450 IP (tos 0x0, ttl 64, id 56400, offset 0, flags [DF], proto TCP (6), length 1438)
       HOSTNAME.55391 > 11.0.0.139.ssh: Flags [.], cksum 0xc8c9 (correct), seq 32:1418, ack 1566, win 251, options [nop,nop,TS val 3135584 ecr 1060793], length 1386
   12:36:57.643373 IP (tos 0x0, ttl 64, id 56401, offset 0, flags [DF], proto TCP (6), length 1438)
       HOSTNAME.55391 > 11.0.0.139.ssh: Flags [.], cksum 0xb639 (correct), seq 32:1418, ack 1566, win 251, options [nop,nop,TS val 3140336 ecr 1060793], length 1386
   12:37:29.270027 IP (tos 0x0, ttl 64, id 56402, offset 0, flags [DF], proto TCP (6), length 1438)
       HOSTNAME.55391 > 11.0.0.139.ssh: Flags [.], cksum 0x9129 (correct), seq 32:1418, ack 1566, win 251, options [nop,nop,TS val 3149824 ecr 1060793], length 1386
   12:38:26.189231 IP (tos 0x0, ttl 62, id 25707, offset 0, flags [DF], proto TCP (6), length 64)
       11.0.0.139.ssh > HOSTNAME.55391: Flags [F.], cksum 0x86da (correct), seq 1566, ack 32, win 227, options [nop,nop,TS val 1090782 ecr 3130899,nop,nop,sack 1 {1418:2000}], length 0
   12:38:26.190197 IP (tos 0x0, ttl 64, id 56403, offset 0, flags [DF], proto TCP (6), length 132)
       HOSTNAME.55391 > 11.0.0.139.ssh: Flags [FP.], cksum 0x437d (correct), seq 2000:2080, ack 1567, win 251, options [nop,nop,TS val 3166900 ecr 1090782], length 80
   12:38:26.232692 IP (tos 0x0, ttl 62, id 41287, offset 0, flags [DF], proto TCP (6), length 40)
       11.0.0.139.ssh > HOSTNAME.55391: Flags [R], cksum 0x6dc0 (correct), seq 3868021228, win 0, length 0
What would be the cause to prevent SSH finishing the connection?
