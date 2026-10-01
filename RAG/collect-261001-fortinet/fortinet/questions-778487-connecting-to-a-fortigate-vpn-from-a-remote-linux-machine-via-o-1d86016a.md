---
id: collect-261001-fortinet/fortinet/questions-778487-connecting-to-a-fortigate-vpn-from-a-remote-linux-machine-via-o-1d86016a
title: "questions-778487-connecting-to-a-fortigate-vpn-from-a-remote-linux-machine-via-o-1d86016a"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/questions-778487-connecting-to-a-fortigate-vpn-from-a-remote-linux-machine-via-o-1d86016a.md
source_anchor: ""
source_lines: [1, 95]
sha256: 11e29851fc809a38e038bf7cd7b260fab812059c12af8dc68db80bd6ae983041
---

# questions-778487-connecting-to-a-fortigate-vpn-from-a-remote-linux-machine-via-o-1d86016a

Here's the setup:
I have a FortiGate unit on a business network, which has a FortiGate VPN set up. Machines on a remote network that can run FortiClient (Windows and Mac machines) have no problem connecting to this VPN. I have been tasked with getting Linux machines to connect to the VPN, which is unsupported by Fortigate.
To try to figure out how, I have an Ubuntu 16.04 machine set up on a remote network, with OpenSwan running trying to connect to a specific tunnel I set up for it on the FortiGate.
The closest I can get it to connecting so far, though, is this:
002 "icms" #1: initiating Aggressive Mode #1, connection "icms"
113 "icms" #1: STATE_AGGR_I1: initiate
003 "icms" #1: received Vendor ID payload [RFC 3947] method set to=115 
003 "icms" #1: received Vendor ID payload [Dead Peer Detection]
003 "icms" #1: received Vendor ID payload [XAUTH]
003 "icms" #1: ignoring unknown Vendor ID payload [8299031757a36082c6a621de0005024d]
002 "icms" #1: Aggressive mode peer ID is ID_IPV4_ADDR: 'a.b.c.d'
003 "icms" #1: no suitable connection for peer 'a.b.c.d'
003 "icms" #1: initial Aggressive Mode packet claiming to be from a.b.c.d on a.b.c.d but no connection has been authorized
218 "icms" #1: STATE_AGGR_I1: INVALID_ID_INFORMATION
002 "icms" #1: sending notification INVALID_ID_INFORMATION to a.b.c.d:500
Where "icms" is the name of the connection, and 'a.b.c.d' is standing in for the public IP of the FortiGate.
My /etc/ipsec.d/icms.conf configuration:
conn icms  
    type=tunnel  
    authby=secret  
    pfs=no  
    ike=aes128-sha1;modp1536  
    phase2alg=aes128-sha1  
    aggrmode=yes  
    keylife=28800s  
    ikelifetime=1800s  
    right=a.b.c.d  
    rightnexthop=%defaultroute  
    rightsubnet=172.16.1.0/16
    left=e.f.g.h  
    leftnexthop=%defaultroute 
    auto=add  
'e.f.g.h' is the IP of the Ubuntu machine.
My /etc/ipsec.d/icms.secrets:
a.b.c.d : PSK "presharedsecret"
Any help or advice at all would be appreciated, and if I can provide any more information please tell me. I have tried multiple configurations of OpenSwan and FortiGate tunnels, to no avail so far.
EDIT 1: the FortiGate config info!
config vpn ipsec phase1-interface
    edit "icms"
        set type static
        set interface "wan1"
        set ip-version 4
        set ike-version 1
        set local-gw 0.0.0.0
        set nattraversal enable
        set keylife 86400
        set authmethod psk
        set mode aggressive
        set peertype any
        set mode-cfg disable
        set proposal aes128-sha1 aes192-sha256
        set localid "icms"
        set localid-type auto
        set negotiate-timeout 30
        set fragmentation enable
        set dpd enable
        set forticlient-enforcement disable
        set comments "Phase1 to Remote Linux"
        set npu-offload enable
        set dhgrp 14 5
        set wizard-type custom
--More--                  set xauthtype disable
        set mesh-selector-type disable
        set remote-gw '<IP of Ubuntu Machine>'
        set monitor ''
        set add-gw-route disable
        set psksecret ENC <encrypted string>
        set keepalive 10
        set auto-negotiate enable
        set dpd-retrycount 3
        set dpd-retryinterval 5
    next
 end
And the phase2 fortigate config:
config vpn ipsec phase2-interface
edit "@icms"
    set phase1name "icms"
    set proposal aes128-sha1 aes256-sha1 3des-sha1 aes128-sha256 aes256-sha256 3des-sha256
    set pfs disable
    set replay enable
    set keepalive disable
    set auto-negotiate enable
    set keylife-type seconds
    set encapsulation tunnel-mode
    set comments ''
    set protocol 0
    set src-addr-type subnet
    set src-port 0
    set dst-addr-type ip
    set dst-port 0
    set keylifeseconds 43200
    set src-subnet 172.16.1.0 255.255.255.248
    set dst-start-ip '<IP of Ubuntu Machine>'
next
end
