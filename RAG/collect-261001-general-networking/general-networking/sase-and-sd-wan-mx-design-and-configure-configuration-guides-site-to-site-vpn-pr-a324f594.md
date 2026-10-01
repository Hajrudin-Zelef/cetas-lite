---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-pr-a324f594
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-pr-a324f594"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-pr-a324f594.md
source_anchor: ""
source_lines: [1, 80]
sha256: 4f6571d73a8c66341993c06f0a991346b0c2c2a01e628aa11dcf031498fd8c05
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-pr-a324f594

Primary and Secondary IPsec VPN Tunnels
Click 日本語 for Japanese
Overview
Primary and secondary IPsec tunnels can be created to provide redundant connectivity to a remote peer. Use cases include SASE, SSE, SWG, FWaaS, etc.
Prerequisites
- 
    IKEv2 must be configured
- 
    MX 19.1.4 or newer release
- 
    MX and Z4/C platforms that support MX 19.1.4 firmware and above
Note: The Primary and Secondary IPsec VPN Tunnel feature is designed for redundant connectivity to external peers (e.g., SASE, SSE). It is not supported for IPsec VPNs established between two Meraki MX devices.
Tunnel Monitoring
To enable failover between primary and secondary tunnels, tunnel monitoring is required. Tunnel monitoring with Layer 7 health check (HTTP probes) enables tracking of primary and secondary IPsec tunnels to determine Layer 3 and 7 connectivity over both tunnels. Meraki Secure SD-WAN uses the health check monitoring data to determine when to automatically failover to the Secondary tunnel.
Note: Layer 7 health check is not supported on dynamic routed BGP enabled tunnels. Tunnel monitoring with Layer 7 health check is only supported with static routed tunnels.
For IPSec configurations utilizing static routing (0.0.0.0/0) and health checks, health checks are currently limited to MX-to-single-headend architectures. This means that health checks will not influence routing failover as only one MX-to-headend health check will work as expected. We are expecting to address this in future firmware releases. Please note that this limitation does not apply to Meraki AutoVPN Secure Access deployments, which fully support these configurations.
Please see here for further detail: https://securitydocs.cisco.com/docs/...lh/121311.dita 
Configure Layer 7 Health Checks
- Click on the Configure health check button as shown in the following image to start the process.
2. Configure a health check name and endpoint. Meraki Secure SD-WAN will probe the configured endpoint to track connectivity over the tunnel. The endpoint hostname will be resolved over the local WAN uplink with the DNS server IP configured on the WAN interface.
3. Apply the health check to a tunnel. Navigate to the tunnel you want to monitor or create a new IPsec peer.
4. If you already have an IPsec peer configured, click on the … dotted menu next to the configured peer, and click on edit. On the side drawer under Tunnel monitoring, select the configured health check from the Health check drop-down and save. This will apply the health check to the tunnel.
Tunnel monitoring with Layer 7 health check must be configured to enable tunnel monitoring and failover to the Secondary tunnel.
- Once configured, the Layer 7 health check and failover to direct internet will be inherited by the secondary tunnel.
- Only one Layer 7 health check can be applied to a tunnel pair (primary and secondary)
The Enable failover to DIA (Direct Internet Access) allows you to control failover behavior if both primary and secondary tunnels are down, or marked failed by tunnel monitoring.
Failover to DIA is turned off by default, which means when the tunnel fails, traffic will not failover to the Internet uplink. However if the failover to DIA option is enabled, traffic will be rerouted to the local uplink.
Once tunnel monitoring has been set up, you can now add a secondary tunnel.
Creating a Secondary Tunnel
Secondary tunnels can be enabled to provide redundancy if the primary tunnel fails.
To create a secondary tunnel, navigate to the primary tunnel you want to create a secondary tunnel for, and click on the dotted … menu to right of the peer, and click on Add option under secondary. This will open a dedicated side drawer for the secondary peer.
- Secondary tunnels are not supported on BGP over IPsec peers. If you have BGP enabled on your primary peer, you will not see the (Primary or Secondary) option
Configuring a secondary peer is the same as configuring a primary one. The inherit primary peer configuration makes it easy to inherit primary settings and minimizes error during configuration. IKE version, routing, private subnets and availability will be automatically inherited from the primary and set to read-only.
Modifying an Existing Secondary Tunnel Peer Information
Note: The maximum recommended number of Dashboard configured IPsec peers is 1500.
Secondary tunnels can be modified by:
- Clicking the ">" arrow on the left-hand side next to the Primary peer name. This will expand and show the Secondary peer information.
- Then click on the three dots "..." sign on the right-hand side of the Secondary peer line to Edit secondary peer.
Tunnel failover criteria
Failover between primary and secondary tunnel occurs when:
- 
    Tunnel is up (IPsec is established) but the health check probes fail.
- 
    Tunnel is down. 
MX periodically applies the following policy based on the latest probe data:
- 
    If the primary tunnel is up, MX picks it.
- 
    If the primary tunnel is down, and the secondary tunnel is up, MX picks the secondary.
- 
    If both the primary and secondary tunnels are down, the MX picks Direct Internet Access if enabled.
Note: This failover mechanism is pre-emptive: traffic will automatically return to the primary tunnel once it is detected as "up" again.
“Up” means that the most recent probe has succeeded (returned a 1xx, 2xx or 3xx status code).
“Down” means that the most recent probe has failed (ICMP unreachable, TCP reset, HTTP 4xx or 5xx code or similar) or timed out (packets lost) after 10 seconds.
Note: The delay could be longer from an end user experience perspective as the end user’s device (laptop or smartphone) may take time to reestablish connections after we have completed the failover or failback.
Failover Timers
Probe interval - 10 secs
Failover time - between 1 - 30 secs
Note: Probe interval and failover time is not configurable at this time.
Probe IP
Tunnel monitoring probes are sourced from 192.0.2.3/32 by default for both primary and secondary tunnels. The Probe IP cannot be configured via Dashboard at this time. The ability to customize your probe IP via Dashboard will be available at a later date.
VPN Tunnel and Health Check Status
Navigate to Security & SD-WAN > Monitor > VPN Status – IPsec peers tab.
Here you can see the status of the IPsec tunnel.
| VPN status indicator | Meaning | 
|---|---|
| Green | Phase 1 and phase 2 are up | 
| Amber | Phase 1 is up but phase 2 is down | 
| Red | Phase 1 and phase 2 are both down | 
In an optimal scenario both the primary and secondary peer will be passing the health check and appear as follows:
Health check status is currently logged in the event log. The "Non-Meraki VPN Healthcheck" event type category can be used for health check logs.
Troubleshooting
- 
    IPsec tunnel status - Check if the IPsec tunnel is up.
- 
    Health check status - Check if health check probes are successful.
- 
    DNS resolution failure - Check if endpoint is resolvable with WAN uplink DNS, try a generic endpoint.
- 
    Remote peer issue - Troubleshoot the IPsec tunnel.
