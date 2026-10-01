---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-7-administration-guide-567401-dialup-ipsec-vpn-using-cust-f240bffe-2
title: "document-fortigate-7-4-7-administration-guide-567401-dialup-ipsec-vpn-using-cust-f240bffe"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-7-administration-guide-567401-dialup-ipsec-vpn-using-cust-f240bffe.md
source_anchor: ""
source_lines: [76, 80]
sha256: f1c94b4d412a4af10b748229fa36e3ed7e5c53ac156951bc6dd085aebc804f4a
---

# document-fortigate-7-4-7-administration-guide-567401-dialup-ipsec-vpn-using-cust-f240bffe

                                                    On FortiGate, run diagnose vpn ike gateway list to verify the IPsec VPN tunnel status.Note that addr shows the custom TCP port value, andtransport showsTCP :vd: root/0 name: v2_psk-120_0 version: 2 interface: port1 3 addr: 10.152.35.150:5500 -> 10.152.35.193:54854 tun_id: 9.5.6.7/::10.0.0.23 remote_location: 0.0.0.0 network-id: 0 transport: TCP virtual-interface-addr: 169.254.1.1 -> 169.254.1.1 created: 592s ago eap-user: ipsec 2FA: no peer-id: 120 peer-id-auth: no FortiClient UID: B70BAD123010487E86DB102969115E99 assigned IPv4 address: 9.5.6.7/255.255.255.255 nat: me peer pending-queue: 0 PPK: no IKE SA: created 1/1 established 1/1 time 80/80/80 ms IPsec SA: created 1/1 established 1/1 time 0/0/0 ms id/spi: 23 93b6803bff7cff00/f89d6f9965fbf3a7 direction: responder status: established 592-592s ago = 80ms proposal: aes256-sha256 child: no SK_ei: f93108f3f8d9a94e-3e0a78289defb329-4d1ae67365f2cb56-e0d471a57ccb4f8d SK_er: 58b37cf4d2e96cb3-cb7e334a48905459-ac8e4ff743c86e5c-630454f2e35b97e6 SK_ai: fc83b139808121a2-1dd68396d804e28d-bd619c0c4778dbda-9a1eb9e6fdf13808 SK_ar: edad89ee56bf9ecc-81443426c00c78f5-0574d6b71163a43b-d9ebf04c3ae4b87f PPK: no message-id sent/recv: 0/124 QKD: no lifetime/rekey: 86400/85537 DPD sent/recv: 00000000/00000000 peer-id: 120
- 
                                                    Run a packet capture using the packet capture tool on FortiGate GUI under Network > Diagnostic tab for wan1(port1) interface with TCP port number 5500. For more information, see Using the packet capture tool.
- 
                                                    (Optional) Run the packet capture using the following command: diagnose sniffer packet wan1 "port 5500" 4 0 l. For more information, see Performing a sniffer trace or packet capture.
