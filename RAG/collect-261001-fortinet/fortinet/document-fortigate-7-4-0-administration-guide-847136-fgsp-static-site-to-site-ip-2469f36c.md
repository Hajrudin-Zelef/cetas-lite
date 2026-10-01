---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-0-administration-guide-847136-fgsp-static-site-to-site-ip-2469f36c
title: "document-fortigate-7-4-0-administration-guide-847136-fgsp-static-site-to-site-ip-2469f36c"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-0-administration-guide-847136-fgsp-static-site-to-site-ip-2469f36c.md
source_anchor: ""
source_lines: [1, 50]
sha256: 8efb9735744aacfd684d9cc9745254f7575e251da939e03a3c623cbcd9fd607e
---

# document-fortigate-7-4-0-administration-guide-847136-fgsp-static-site-to-site-ip-2469f36c

FGSP static site-to-site IPsec VPN setup
FGSP static site-to-site IPsec VPN setup
When configuring static site-to-site IPsec VPN between FGSP FortiGates and a remote gateway, the FGSP peers must have the  passive-mode setting enabled in the vpn ipsec phase1-interface configuration to function as an IPsec responder. This is a required configuration in this setup. If the FGSP peers act as initiators for tunnel setup when passive-mode is disabled and both FGSP peers initiate the tunnel with the same gateway IP, the remote IPsec gateway will be unable to process this, and the tunnel negotiation will fail. Likewise, when a failover occurs in FGSP and a new peer begins to initiate tunnel traffic, the remote IPsec gateway will be unable to handle the traffic initiated from the new peer.
Enabling passive-mode ensures the FGSP peers only respond to tunnel initiations from the remote IPsec gateway and do not initiate tunnel negotiations. This way, the preceding situations will not occur.
For dynamic tunnel configuration examples on FGSP peers, see the following topics:
- FGSP per-tunnel failover for IPsec
- FGCP over FGSP per-tunnel failover for IPsec
- Allow IPsec DPD in FGSP members to support failovers
Example
In this example, the FGSP peer has a loopback interface (Lo1) configured with the same IP address (192.168.202.31/32). All other interface IP addresses are different between the peers.
To configure IPsec on the FGSP peer FortiGates:
- Configure the phase 1 settings:config vpn ipsec phase1-interface
    edit "IPsec"
        set type static
        set interface "port1"
        set ike-version 2
        set local-gw 192.168.202.31
        set net-device disable
        set proposal aes256-sha256
        set dhgrp 14
        set passive-mode enable
        set remote-gw 10.10.100.100
    next
end
- Configure the phase 2 settings:config vpn ipsec phase2-interface
    edit "IPSec"
        set phase1name "IPsec"
        set proposal aes256-sha256
        set dhgrp 14
    next
end
To configure IPsec on the remote FortiGate:
- Configure the phase 1 settings:config vpn ipsec phase1-interface
    edit "IPsec"
        set type static
        set interface "port1"
        set ike-version 2
        set net-device disable
        set proposal aes256-sha256
        set dhgrp 14
        set remote-gw 192.168.202.31
    next
end
- Configure the phase 2 settings:config vpn ipsec phase2-interface
    edit "IPSec"
        set phase1name "IPsec"
        set proposal aes256-sha256
        set dhgrp 14
    next
end
