---
id: collect-261001-cisco/cisco/sase-and-sd-wan-mx-integrations-security-service-edge-integrations-cisco-secure-cfea9725-2
title: "sase-and-sd-wan-mx-integrations-security-service-edge-integrations-cisco-secure--cfea9725"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/sase-and-sd-wan-mx-integrations-security-service-edge-integrations-cisco-secure--cfea9725.md
source_anchor: ""
source_lines: [108, 166]
sha256: 34e0ec3311399cf46e16a3b3c39312e507e9ad8cca39c10c2fcc819a8929659c
---

# sase-and-sd-wan-mx-integrations-security-service-edge-integrations-cisco-secure--cfea9725

- 
    Click Update & Save
Creating a Secondary or Backup Tunnel
Secondary tunnels can be enabled to provide redundancy if the primary tunnel fails.
To create a Secondary tunnel, navigate to the Primary tunnel you want to create a secondary tunnel for, and click on the dotted … menu to right of the peer, and click on Add option under Secondary. This will open a dedicated side drawer for the Secondary peer.
Configuring a secondary peer is the same as configuring a primary one. The inherit primary peer configuration makes it easy to inherit primary settings and minimizes error during configuration. IKE version, routing, private subnets and availability will be automatically inherited from the primary and set to read-only.
Save configuration.
Multi-Uplink Active-Active IPsec
This feature is mainly for Secure Internet use case with SASE/SSE providers. Considering traffic will be load balanced across all healthy tunnels, it’s important to ensure that the remote SASE/SSE peer sends the return traffic back through the tunnel it was received.
When Multi-Uplink IPsec is enabled, IPsec tunnels are established on all available uplinks, and traffic intelligently load-balanced across all healthy tunnels, providing both performance gains and built-in redundancy. If any tunnel becomes unavailable, traffic seamlessly redistributes across the remaining active connections.
To enable Multi Uplink IPsec:
- Create or edit an IPsec peer
- Enable Mutli-Uplink IPsec by toggling on Multi-Uplink IPsec VPN option on the peer side drawer as shown below.
Verification and Troubleshooting
In an AutoVPN and IPsec VPN environment, when configuration changes are applied on a Hub to a subnet that participates in the VPN, VPN connections will reset on every site connected to the Hub to update the VPN configuration.
The Network Tunnel Group will move from Disconnected Status to Connected. This change could take a few minutes.
- 
    Run ping tests from a tunnel enabled VLAN to the internet. For more information, see Using the Ping Live Tool.
- 
    Check the status of the VPN tunnel. For more information, see VPN Status Page.
- 
    Follow the VPN troubleshooting procedures. For more information, see Site-to-site VPN Troubleshooting.
IPsec VPN Firewall
You can add firewall rules to control what traffic is allowed to pass through the VPN tunnel. These rules will apply to outbound VPN traffic to/from all MX-Z appliances in the Organization that participate in site-to-site VPN. These rules are configured in the same manner as the Layer 3 firewall rules described on the Firewall Settings page of this documentation. Note that VPN Firewall rules will not apply to inbound traffic or to traffic that is not passing through the VPN.
Serviceability
Event Logs
If you have any issues or would like to know more about the Cisco Secure Access peering details, navigate to Network-wide > Monitor > Event log
- Select Event type Include - IPsec VPN Negotiation
- IPsec VPN Healthcheck event type can be used to identify the recently reported healthcheck status of the tunnels
Packet Captures
The following options are available for a packet capture on MX/Z platforms:
- 
    Appliance: The appliance the capture will run on.
- 
    Interface: Select the interface to run the capture on; the interface names will vary depending on the appliance configuration. A few examples of interfaces you may see are: 
  - 
        Internet 1 or Internet 2 - Capture traffic on one active WAN uplink. Internet 2 will only appear if there is a second WAN link.
  - 
        LAN - Captures traffic from all LAN ports
  - 
        Cellular - Captures cellular traffic from the integrated cellular interface. This does not apply to USB modems.
  - 
        Site-to-Site VPN - Captures AutoVPN traffic (MX/Z to MX/Z only). This does not apply to IPsec VPN peers.
- 
        
- 
    Output: Select how the capture should be displayed; view output or download .pcap.
- 
    Verbosity: Select the level of the packet capture (only available when viewing the output directly to Dashboard).
- 
    Ignore: Optionally ignore capturing broadcast/multicast traffic.
- 
    Filter expressions: Apply a capture filter.
To capture packets, select the WAN interface and use the filter expressions for UDP 500 for Phase 1 or UDP 4500 for Phase 2.
API
The Meraki dashboard API is an interface for software to interact directly with the Meraki cloud platform and Meraki-managed devices. The API contains a set of tools known as endpoints for building software and applications that communicate with the Meraki dashboard for use cases such as provisioning, bulk configuration changes, monitoring, and role-based access controls. The dashboard API is a modern, RESTful API using HTTPS requests to a URL and JSON as a human-readable format. The dashboard API is an open-ended tool that can be used for many purposes.
For more information, read here.
24/7 Support
Cisco Meraki Support is available 24/7 to Enterprise customers for assistance with resolving network issues and providing answers to questions not covered by the documentation. For more information, read here.
