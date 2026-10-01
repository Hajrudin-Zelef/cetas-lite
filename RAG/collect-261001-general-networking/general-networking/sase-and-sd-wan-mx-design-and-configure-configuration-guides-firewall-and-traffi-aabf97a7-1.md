---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-aabf97a7-1
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-aabf97a7"
domain: general-networking
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["latency", "throughput"]
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-aabf97a7.md
source_anchor: ""
source_lines: [1, 48]
sha256: 27fb11cc020ad90e1c94dd167add534e3236e13b1535e0a5ec702e0269fcef9f
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-aabf97a7

SD-WAN and Traffic Shaping
Click 日本語 for Japanese
The MX appliance and Z-series gateway include an integrated Layer 7 packet inspection engine, enabling you to set QoS policies, load balancing, and prioritization based on traffic types and applications.
Uplink configuration
This section allows you to configure bandwidth settings, uplink statistics, and list update interval.
Uplink bandwidth settings
This option allows you to configure the upload and download bandwidth of the uplinks. This information is needed for traffic load balancing between the active WAN ports as well as for limiting upload and download traffic. You can configure WAN 1, WAN 2, and the cellular uplink individually. To configure specific upload and download bandwidths for a particular uplink, click the details button next to that uplink's bandwidth slider. The maximum value for each WAN interface is based on the rating of the specific model MX.
Uplink bandwidth settings apply to live tools and WAN connectivity monitoring (with the exception of the "dashboard throughput" tool) and available bandwidth will be split between client traffic flows and these appliance-sourced traffic flows. Extremely low upload bandwidth limits (50kbps or less) may interfere with live tool functionality.
Uplink Statistics
Clicking the Add Your Destination option allows you to add a custom destination for the MX to continually test ICMP connectivity to monitor latency and packet loss. These destinations cannot be private addresses across VPN tunnels and must be reachable through the WAN interface of the MX. Additionally, they must respond to ICMP traffic. It is good practice to include the MX's default gateway to monitor the directly connected link. The results for these tests will be visible at Security & SD-WAN > Monitor > Appliance status > Uplink > Historical data. Hostnames/FQDNs are not supported for uplink statistics monitoring.
MXs and Zs with integrated cellular will not send ICMP Uplink Statistics tests through their cellular interface. This is designed to help reduce the amount of traffic that is sent out of these uplinks. If there is a requirement to send ICMP tests for Uplink Statistics, please reach out to Meraki Support to have this functionality enabled. Please note that only tests to Google DNS (8.8.8.8) will be sent, not to any other Uplink Statistics destinations.
List update interval
This setting determines how often the MX should check for updates to security lists. You can specify an Hourly, Daily, or Weekly update interval. To specify different intervals depending on which uplink is being used to download lists, click details. This can be useful if you want to control bandwidth usage due to security list downloads on a low-bandwidth WAN link or cellular uplink.
Features affected by this setting include IDS/IPS and Malware Scanning.
Uplink selection
Primary uplink
This option determines which uplink should be the primary connection. VPN traffic and management traffic to the Meraki Dashboard use the primary uplink. If load balancing is disabled, all traffic will use the primary uplink unless an uplink preference specifies otherwise. Enhanced Failover logic will also apply in the event of a WAN flap.
Load balancing
When enabled, load balancing spreads Internet traffic across both uplinks proportional to the WAN 1 and WAN 2 bandwidths specified above.
Example: If WAN 1's bandwidth is 9 Mbps and WAN 2's bandwidth is 1 Mbps, the load-balancing algorithm sends 90% of the traffic through the WAN 1 uplink and 10% of the traffic through the WAN 2 uplink.
Multi-Uplink AutoVPN
This option is used to determine if AutoVPN tunnels should be formed over only the primary uplink or over all active uplinks simultaneously. There are two options that can be configured:
- Enabled: Create VPN tunnels over all of the available uplinks.
- Disabled: Do not create VPN tunnels over non-primary uplinks, unless the primary uplink fails.
Flow preferences
Use this option to direct traffic matching a Layer 3 definition via a particular uplink. Some common use cases involve sending traffic from different VLANs through different Internet uplinks, or sending a particular type of traffic such as FTP traffic out a particular uplink based on the destination port.
Note: ICMP traffic is not subject to traffic shaping rules. As a result, flow preferences will have no impact on ICMP traffic.
SD-WAN over Cellular Active Uplink
To use SD-WAN over cellular, the MX needs to be running MX16.2+ and have the feature enabled on an integrated cellular MX (MX67C and MX68CW only).
With this feature enabled, the cellular connection that was previously backup-only, can be configured as an active uplink in the SD-WAN & traffic shaping page as shown below:
When this toggle is set to Enabled, the cellular interface details, found on the Uplink tab of the Appliance status page, will show as Active even when a wired connection is also active, as shown below:
Given that this feature takes ownership of the WAN2 logic, enabling it means that the use of two wired networks is not supported. This is because only two WAN connections can be used concurrently.
This means that the wired connection can only be connected to Internet port 1 or WAN 1.
When using this feature on an MX67C, this results in the port LAN2 being unusable because LAN2 is a multi-use port that can also operate as WAN2.
When using this feature on an MX68CW, this results in the Internet 2 port being unusable.
As such, to configure an SD-WAN policy to utilize the cellular connection, associate it with WAN2 as shown below:
SD-WAN policies
Global bandwidth limits
This setting allows you to set limits on each client device's total network traffic (incoming / outgoing). The minimum throughput limit is 20 Kb/s. Click details or simple to switch between two possible modes.
- simple: Single setting that applies to both upload and download traffic throughput. Move the slider control right or left to set the limits.
- details: Allows you to set different limits on upload and download throughput. Enter the limits manually in Kb/s. You can also use this mode to create more precise per-client limits than in simple mode.
Enable SpeedBurst: To provide a better user experience in bandwidth-limited environments, an administrator can enable SpeedBurst by selecting the Enable SpeedBurst checkbox. SpeedBurst allows users to exceed their assigned limit in a burst for a short period of time. This provides a more satisfying Internet browsing experience. It also prevents any one user from using more than their fair share of bandwidth over the longer term. Users are allowed up to four times their allotted bandwidth limit for a period of up to five seconds.
Traffic shaping rules
To optimize your network, you can create shaping policies to apply per-user controls on a per-application basis. This allows you to reduce bandwidth for recreational applications such as peer-to-peer file sharing programs, and to prioritize bandwidth for your business-critical enterprise applications. The Meraki dashboard has a set of default rules that can be enabled/disabled.
You can also create custom rules and apply them to your desired traffic signatures. If a custom-defined rule is created that overlaps with a default rule, then the custom-defined rule will take effect.
Traffic shaping rules will also apply to traffic sent over an AutoVPN tunnel between Meraki devices, but do not apply to traffic that passes over a non-Meraki VPN tunnel.
The use of NBAR will add categories to the traffic shaping section that were not available before. More information about NBAR can be found at Next-gen Traffic Analytics - Network-Based Application Recognition (NBAR) Integration.
Creating Shaping Rules
