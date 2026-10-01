---
id: collect-261001-cisco/cisco/sase-and-sd-wan-mx-integrations-security-service-edge-integrations-cisco-secure-cfea9725-1
title: "sase-and-sd-wan-mx-integrations-security-service-edge-integrations-cisco-secure--cfea9725"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/sase-and-sd-wan-mx-integrations-security-service-edge-integrations-cisco-secure--cfea9725.md
source_anchor: ""
source_lines: [1, 107]
sha256: c053eaa151554ddc459d9e5b88a988a524a0f05033e3eae0d29aea61162a7a13
---

# sase-and-sd-wan-mx-integrations-security-service-edge-integrations-cisco-secure--cfea9725

Cisco Secure Access Meraki SD-WAN Configuration Guide
Cisco Secure Internet Access Configuration Guide
Cisco Secure Access offers a security stack solution from the cloud for internet, SaaS, ZTNA, Remote access connections, etc. Cisco Secure Access acts as a security gateway where 0.0.0.0/0 traffic will be routed for inspection and enforcement prior to internet or site-to-site, private cloud termination.
This document covers configuration for Secure Internet access with Primary & Secondary static tunnels.
Meraki SDWAN + Secure Access Design Guide
Private Access BGP Configuration Guide
Prerequisites
- 
    Cisco Secure Access account
- 
    Meraki MX/Z device (running MX19.1.6+ firmware)
- 
    Meraki MX/Z Site-to-site VPN enabled
Caveats and Considerations
- 
    ECMP/Load balancing is not supported with this integration
Cisco Secure Access Configuration
Go to sse.cisco.com and login with your credentials and follow the steps outlined below.
1. Add a Network Tunnel Group - From the Secure Access console, navigate to Connect > Network Connections.
 
- 
    Select Network Tunnel Groups
- 
    Click Add. 
- 
    Enter the General Settings for your tunnel group:
- 
    Give your tunnel group a meaningful name.
- 
    Choose a Region.
- 
    Choose a Device Type.
- 
    Click Next.
- 
    Enter the Tunnel ID and Passphrase for your tunnel group:
- 
    Tunnel ID Format: Email.
- 
    Enter Passphrase
- 
    Re Enter Passphrase.
- 
    Click Next.
Note: The passphrase must be between 16 and 64 characters in length and contain at least one upper case letter, one lower case letter, and one number. The passphrase cannot include any special characters. 
 
4. Routing - choose Enable NAT / Outbound Only (for Secure Internet Access use case)
Static Routing configuration in Secure Access is not supported with Meraki SD-WAN. In Secure Access set Routing to "Enable NAT / Outbound Only" for Secure Internet Access use cases, and Dynamic Routing for Private Access use cases. Refer to the Design guide for supported Design topologies and configurations.
- 
    Click Save.
- 
    On the Data for Tunnel Setup page, review the network tunnel information for completeness. Click the Download CSV button to save configuration information needed for your Meraki Secure SD-WAN device.
Note: This is the only time that your passphrase is displayed.
- 
    Click Done.
A Network Tunnel Group will be configured but the Primary and Secondary Hub will be down until a tunnel is established to the Hubs.
Meraki Secure SD-WAN Configuration
In an AutoVPN and IPsec VPN environment, when configuration changes are applied on a Hub to a subnet that participates in the VPN, VPN connections will reset on every site connected to the Hub to update the VPN configuration.
Gather Details from Cisco Secure Access Portal
Information can be taken from the CSV downloaded from step 5 above or from the Data for Tunnel Setup page seen below:
On the Meraki Network, Navigate to Site-to-site VPN settings through the Security & SD-WAN > Configure > Site-to-site VPN page.
There are three options for configuring the MX-Z's role in the Auto VPN topology:
- 
    Off: The MX-Z device will not participate in site-to-site VPN.
- 
    Hub (Mesh): The MX-Z device will establish VPN tunnels to all remote Meraki VPN peers that are also configured in this mode, as well as any MX-Z appliances in hub-and-spoke mode that have the MX-Z device configured as a hub.
- 
    Spoke: This MX-Z device (spoke) will establish direct tunnels only to the specified remote MX-Z devices (hubs). Other spokes will be reachable via their respective hubs unless blocked by site-to-site firewall rules.
Select Hub (Mesh) or Spoke depending on your AutoVPN requirements, to enable VPN.
Configure Layer 7 Health Checks
Click on the Configure health check button as shown in the following image to start the process.
Next, configure a health check name and endpoint. Meraki Secure SD-WAN will probe the configured endpoint to track connectivity over the tunnel. The endpoint hostname will be resolved over the local WAN uplink with the DNS server IP configured on the WAN interface.
Next, create a new IPsec peer by clicking on the "Add peer" button on the IPsec VPN peers table
You can create Site-to-site VPN tunnels between a Security Appliance or a Teleworker Gateway and Cisco Secure Access under the IPsec VPN peers section on the Security & SD-WAN > Configure > Site-to-site VPN page.
Enter the following information:
- 
    Name—Provide a meaningful name for the tunnel.
- 
    IKE Version—Select IKEv2.
- 
    IPsec policies—Choose the predefined Umbrella configuration, see Supported IPsec Parameters.
- 
    Public IP—IP address to connect to Secure Access Network Tunnel Group Primary Data Center IP
- 
    Local ID—The Primary Tunnel ID for the Network Tunnel Group.
- 
    Remote ID—Leave this blank
- 
    Preshared secret—The Passphrase for the Network Tunnel Group created in Secure Access
- 
    Routing — Static, configure a default route 0.0.0.0/0
- 
    Availability— Add network tag for the MX network appliance that should build the tunnel to Secure Access. "All Networks" tag will configure all MXs in your Dashboard Organization to establish a tunnel to Secure Access
- 
    Tunnel Monitoring 
  - 
        Health check - select previously configured heath check. This endpoint will be used to monitor connectivity through the tunnel
- 
        
For IPSec configurations utilizing static routing (0.0.0.0/0) and health checks, health checks are currently limited to MX-to-single-headend architectures. This means that health checks will not influence routing failover as only one MX-to-headend health check will work as expected. We are expecting to address this in future firmware releases. Please note that this limitation does not apply to Meraki AutoVPN Secure Access deployments, which fully support these configurations.
Please see here for further detail: https://securitydocs.cisco.com/docs/...lh/121311.dita 
- 
    
  - 
        Failover directly to Internet - Enable if you prefer traffic to failover to the local Internet uplink when both primary and secondary tunnels are marked as failed.
- 
        
