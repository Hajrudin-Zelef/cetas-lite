---
id: collect-261001-general-networking/general-networking/send-feedback-29-1
title: "Dialup IPsec VPN using custom TCP port"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-general-networking/send-feedback-29.md
source_anchor: ""
source_lines: [1, 86]
sha256: c38a403bf91b3cce1811d1a8b724ded4b949cc2a4288c4bc29e7bd9ff9a09ea4
---

# Dialup IPsec VPN using custom TCP port

# Dialup IPsec VPN using custom TCP port

Dialup IPsec VPN may operate over UDP or TCP, enabling VPN traffic from FortiClient to traverse restrictive firewalls that only permit TCP-based traffic. You can configure an IPsec VPN tunnel to exclusively use UDP, or you can configure the tunnel to automatically switch to TCP mode when the firewall blocks UDP.

In high-latency or congested networks, UDP-based VPN connections may suffer from packet loss or performance degradation. TCP, with its built-in error correction and retransmission mechanisms, enhances the reliability and stability of VPN connections in such environments.

Dialup IPsec over TCP is particularly advantageous in mobile or dynamic settings such as public WiFi, hotel networks, or cellular data where network conditions and restrictions often vary. This feature ensures more seamless and dependable VPN connectivity across a broader range of scenarios.

The custom TCP port functionality for IPsec is exclusively supported with IKE version 2 (IKEv2), and does not support NPU offloading.

## Example

In this example, FortiGate is configured as a dialup IPsec server using IKE version 2 (IKEv2) and operating on a custom TCP port (5500). IKEv2 is configured to use EAP for user authentication. The initial setup leverages the VPN wizard to create the dialup IPsec tunnel.

On the client side, FortiClient is managed by FortiClient EMS and configured to act as the dialup IPsec client. The client is configured to connect to the FortiGate server over the custom TCP port 5500. This feature requires FortiClient 7.4.1 or later.

For a detailed description of the steps to configure FortiClient EMS to use the custom TCP port 5500 for IPsec VPN connections, see IPsec VPN over TCP.


###### To enable TCP and configure a custom port in the GUI:

1. 
                                                    Go to *VPN > VPN Tunnels* and select the*Settings* tab.
2. 
                                                    Enable *Allow VPN negotiation over TCP* .
3. 
                                                    Set *TCP port for IKE/IPsec traffic* to port*5500* .
4. 
                                                    Click *Apply* .

###### To enable TCP and configure a custom port in the CLI:

```
config system settings
    set ike-tcp-service enable
    set ike-tcp-port 5500
end
```
                                            | Command | Description | 
|---|---|
| ike-tcp-servicer {enable \| disable} | Enable/disable VPN over TCP. This is a per-vdom setting. | 
| ike-tcp-port <port> | Set the TCP port for IKE/IPsec traffic (1 - 65535, default = 443). | 


When using TCP port 443 for IKE/IPsec traffic, GUI access can be affected for interfaces that are bound to an IPsec tunnel when the GUI admin port is also using port 443. To ensure continued functionality, change either the IKE/IPsec port or the administrative access port.

###### To change the administrative access port:

```
config system global
    set admin-sport <port>
end
```
                                                    | Command | Description | 
|---|---|
| admin-sport <port> | Set the administrative access port for HTTPS (1 - 65535, default = 443). | 

For port conflicts with ZTNA and Agentless VPN, ZTNA and Agentless VPN will take precedence. To avoid any port conflicts with other services, review the FortiOS Ports guide for other incoming ports used on the FortiGate.

###### To configure FortiGate as IPsec dialup server using VPN Wizard:

1. 
                                                    Go to *VPN > VPN Wizard* , and enter the following:Field Value Tunnel name v2_psk-120 Select a template FortiClient remote access
2. 
                                                    Click *Begin* .
3. 
                                                    In the *VPN Tunnel* section, enter the following:Field Value Secure Internet Access (SIA) Off FortiClient management type EMS Authentication method Pre-shared key Pre-shared key Enter suitable key IKE Version 2 NAT traversal Enable Keepalive frequency 10 EAP peer identification EAP identity request User authentication method Phase 1 interface Use dropdown to select user group *IPSEC* . To configure user groups for authentication, see User groups.(Optional) To use multiple user groups, select *Inherit from policy* . See Using single or multiple user groups for user authentication for details.DNS Server Specify Server IP 8.8.8.8
4. 
                                                    Click *Next* .
5. 
                                                    Under *Remote endpoint* section, enter the following:Field Value Addressing mode for connected endpoints Manual Address to assign to connected endpoints 9.5.6.7-9.5.6.70 Subnet for connected endpoints 255.255.255.255 FortiClient settings Security posture gateway matching Disable EMS SN verification Disable Save password Enable Auto Connect Enable Always up (keep alive) Enable
6. 
                                                    Click *Next* .
7. 
                                                    Under *Local FortiGate* section, enter the following:Field Value Incoming interface that binds to tunnel wan1(port1) Create and add interface to Zone Enable Local interface internal (port3) Local Address internal network
8. 
                                                    Click *Next* .
9. 
                                                    Under *Review* section, review the configuration pending configuration by the wizard.
10. 
                                                    Click *Submit* .The tunnel is configured and visible under *VPN > VPN Tunnels* .

###### To configure FortiGate as IPsec dialup server using the CLI:

