---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-6392732c
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-6392732c"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-6392732c.md
source_anchor: ""
source_lines: [1, 47]
sha256: 76f68debae6bea9dae2b50065ee64664d05dbe65c0410e77d2c43ecdd6e040ec
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-6392732c

Creating a DMZ with the MX Security Appliance
The MX Security Appliance can be used to create a DMZ zone using VLANs, Firewall rules, and 1:1 NAT mappings. To do this, three things need to be accomplished:
- Segment the network using VLANs.
- Restrict inter-VLAN traffic using ACLs.
- Forward desired traffic using NAT rules.
In this example, the network will be divided into two zones.
- Internal - Contains clients and other devices not directly reachable from the Internet, but able to initiate outbound communication.
- DMZ - Contains public facing servers and services.
Within the DMZ there is a web server at 172.16.32.2, which should be reachable by all internal clients and any Internet hosts. However, no communication should be allowed to Internal hosts that is initiated by the web server, and only web traffic should be allowed between Internal hosts and the web server in the DMZ. Clients and the DMZ server are both connected to a downstream managed switch. Refer to the topology below.
Segment the network into VLANs
- Navigate to Configure > Addressing & VLANs.
- Ensure that Mode is set to Routed.
- Set VLANs to "Enabled" if not already done.
- Create local VLANs for the Internal and DMZ networks, as shown below.
 
- Ensure that the LAN port connecting to the downstream switch is configured to correctly handle the two VLANs. In this case, VLAN 1 (Internal) is native and untagged, while VLAN 2 (DMZ) is tagged.
 Note: Ensure that the downstream switch is correctly configured to match these settings on the port connecting to the MX.
 
- Click Save Changes
Restrict inter-VLAN traffic using ACLs
- Navigate to Configure > Firewall.
- Under Outbound rules, add the following layer 3 firewall rules.
    
  - Allow TCP:80 traffic from the Internal VLAN to the web server.
  - Allow TCP:443 traffic from the Internal VLAN to the web server.
  - Block all other traffic from the Internal VLAN to the web server and DMZ VLAN.
  - Block all traffic from the DMZ VLAN to the Internal VLAN.
- Click Save Changes.
This will allow:
- Internal clients and DMZ servers to communicate freely with the Internet.
- Internal clients to access web resources on the web server.
- Internet hosts to access web resources on the web server.
...while preventing:
- Internal clients from access other resources on the web server or other DMZ servers (such as SSH or FTP).
- DMZ servers from accessing internal clients, unless in reply (to prevent allowing access to the internal network if the web server is compromised).
- Internet hosts from accessing internal clients.
Forward desired traffic using NAT rules
- Navigate to Configure > Firewall.
- Under 1:1 NAT, add a 1:1 NAT mapping as shown below.
    
  - The Public IP should be the IP address being directed to the selected Uplink, which will be forwarded to the web server.
Note: If using the public IP address on the MX itself, refer to the guide on port forwarding for this section.
  - The LAN IP should be the IP address of the web server.
  - Under Allowed inbound connections, select TCP ports 80 and 443 to forward web traffic to the web server.
  - For Remote IPs enter "any", unless restricting to specific IP addresses or ranges.
- The Public IP should be the IP address being directed to the selected Uplink, which will be forwarded to the web server.
- Click Save Changes.
