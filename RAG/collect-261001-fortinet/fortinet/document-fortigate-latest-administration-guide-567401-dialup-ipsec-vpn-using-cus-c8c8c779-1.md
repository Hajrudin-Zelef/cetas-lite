---
id: collect-261001-fortinet/fortinet/document-fortigate-latest-administration-guide-567401-dialup-ipsec-vpn-using-cus-c8c8c779-1
title: "Dialup IPsec VPN using custom TCP port"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-latest-administration-guide-567401-dialup-ipsec-vpn-using-cus-c8c8c779.md
source_anchor: ""
source_lines: [1, 62]
sha256: 2b3c1a38e56fda35ff2b9533f93e6bc09f469c2681cde14841cc1ee746263640
---

# Dialup IPsec VPN using custom TCP port

Dialup IPsec VPN traditionally relies on UDP but can now operate over TCP. This enhancement enables VPN traffic from FortiClient to traverse restrictive firewalls that only permit TCP-based traffic. You can configure an IPsec VPN tunnel to exclusively use UDP or TCP, or you can configure the tunnel automatically switch to TCP mode when the firewall blocks UDP.

In high-latency or congested networks, UDP-based VPN connections may suffer from packet loss or performance degradation. TCP, with its built-in error correction and retransmission mechanisms, enhances the reliability and stability of VPN connections in such environments.

Dialup IPsec over TCP is particularly advantageous in mobile or dynamic settings such as public WiFi, hotel networks, or cellular data where network conditions and restrictions often vary. This feature ensures more seamless and dependable VPN connectivity across a broader range of scenarios.

|  | The custom TCP port functionality for IPsec is exclusively supported with IKE version 2 (IKEv2), and does not support NPU offloading. | 

In this example, FortiGate is configured as a dialup IPsec server using IKE version 2 (IKEv2) and operating on a custom TCP port (5500). IKEv2 is configured to use EAP for user authentication. The initial setup leverages the VPN wizard to create the dialup IPsec tunnel. After the tunnel is created by the wizard, you use the CLI to customize the IKE settings and enable the use of TCP port 5500.

On the client side, FortiClient is managed by FortiClient EMS and configured to act as the dialup IPsec client. The client is configured to connect to the FortiGate server over the custom TCP port 5500. This feature requires FortiClient 7.4.1 or later.

For a detailed description of the steps to configure FortiClient EMS to use the custom TCP port 5500 for IPsec VPN connections, see IPsec VPN over TCP.


###### To configure FortiGate as IPsec dialup server using VPN Wizard:

1. 
                                                    Go to *VPN > VPN Wizard* , and enter the following:Field Value Tunnel name v2_psk-120 Select a template Remote Access
2. 
                                                    Click *Begin* .
3. 
                                                    Under *VPN Tunnel* section, enter the following:Field Value VPN client type FortiClient Authentication method Pre-shared key Pre-shared key Enter suitable key IKE Version 2 Transport Auto This can be changed to TCP encapsulation in the CLI. Use Fortinet encapsulation Disable The Fortinet propriety feature is designed to offload IPsec VPN traffic to Fortinet’s NP (Network Processor) ASICs to improve performance. This command enables or disables encapsulation of ESP (Encapsulating Security Payload) packets within non-standard TCP headers. 
  - 
                                                                            `disable` : Encapsulates ESP packets using standard TCP headers. This is the default option.
  - 
                                                                            `enable` : Encapsulates ESP packets using non-standard TCP headers.
 This feature is not supported in the following scenarios: 
  - 
                                                                                            When FortiGate is configured as Dialup IPsec VPN server for remote access when using FortiClient as dialup client.
  - 
                                                                                            In multi-vendor environments that use TCP encapsulation for ESP packets.
 Make sure that this setting is disabled for these two scenarios to ensure uninterrupted ESP packet flow encapsulated within real TCP headers. NAT traversal Enable Keepalive frequency 10 EAP peer identification EAP identity request User authentication method Phase 1 interface Use dropdown to select user group *IPSEC* . To configure user groups for authentication, see User groups.(Optional) To use multiple user groups, select *Inherit from policy* . See Using single or multiple user groups for user authentication for details.DNS Server Specify Server IP 8.8.8.8
4. 
                                                                            
5. 
                                                    Click *Next* .
6. 
                                                    Under *Remote Endpoint* section, enter the following:Field Value Address to assign to connected endpoints 9.5.6.7-9.5.6.70 Subnet for connected endpoints 255.255.255.255 FortiClient settings Security posture gateway matching Disable EMS SN verification Disable Save password Enable Auto Connect Enable Always up (keep alive) Enable
7. 
                                                    Click *Next* .
8. 
                                                    Under *Local FortiGate* section, enter the following:Field Value Incoming interface that binds to tunnel wan1(port1) Create and add interface to Zone Enable Local interface internal (port3) Local Address internal network
9. 
                                                    Click *Next* .
10. 
                                                    Under *Review* section, review the configuration pending configuration by the wizard.
11. 
                                                    Click *Submit* .The tunnel is configured and visible under *VPN > VPN Tunnels* .
12. 
                                                    Change the transport protocol to TCP encapsulation in the CLI: ```
config vpn ipsec phase1-interface
    edit "v2_psk-120"
        set transport tcp
    next
end
```

###### To configure FortiGate as IPsec dialup server using the CLI:

