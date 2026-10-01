---
id: collect-261001-huawei/huawei/hc-en-us-articles-204911354-unifi-qos-and-traffic-shaping-ae99abf1
title: "hc-en-us-articles-204911354-unifi-qos-and-traffic-shaping-ae99abf1"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["cyber", "ethernet", "latency", "throughput"]
source: docs/RAG/collect-261001-huawei/hc-en-us-articles-204911354-unifi-qos-and-traffic-shaping-ae99abf1.md
source_anchor: ""
source_lines: [1, 86]
sha256: 6a5365389ba99f099cc344a23be60933924bad4c60dc34d80766e133d25e9ee8
---

# hc-en-us-articles-204911354-unifi-qos-and-traffic-shaping-ae99abf1

UniFi QoS and Traffic Shaping
UniFi offers advanced Quality of Service (QoS) and Traffic Shaping tools that let you prioritize critical applications and limit nonessential traffic, helping ensure optimal performance across all connected devices.
If you’re deploying A/V equipment and want to use pre-configured Port Profiles for Dante, Q-SYS, SDVoE, and other AV systems, refer to our guide on Pro AV Traffic Optimization on UniFi Switches.
For a full overview of UniFi’s Traffic and Policy Management capabilities, see here.
For a full overview of UniFi's Network and Cyber Security capabilities, see here.
Configuring QoS and Traffic Shaping Rules On The Gateway
In general, we recommend applying QoS and Traffic Shaping at the gateway, as it allows for granular control over all Internet-bound traffic and inter-VLAN routing. However, for deployments without a UniFi gateway—or when more localized control is needed for traffic that stays within the LAN—we also support applying QoS and traffic shaping on UniFi access points and switch ports. To do so:
- 
Navigate to QoS Rules: Follow the path depending on your UniFi Network version:
  - Network 9.4: Settings > Policy Table > Create New Policy > QoS
  - Network 9.3: Settings > Policy Engine > QoS > Configure
- Create Name: Assign a name to the rule.
- 
Select an Objective
  - Prioritize: Moves selected traffic into a higher-priority queue to improve performance.
  - Limit: Restricts selected traffic to a specific upload/download speed.
  - Prioritize and Limit: Moves selected traffic into a higher-priority queue and enforces a maximum upload/download speed.
- Select Interface(s): The rule can apply to a VPN Tunnel, WAN interface, or all WANs.
- 
Configure Source and Destination: The Objective configured in (2) will apply to traffic meeting all specified criteria.
  - 
Source
    - Any: Applies to any traffic originating from inside the local network.
    - Device: Specify certain client devices.
    - Network: Specify certain network(s) (VLANs).
  - 
Destination
    - Any: Applies to any, no matter what its destination is.
    - App: Specify specific applications or categories of applications.
    - IP: Specify certain IP addresses or IP address ranges, and ports associated with the destination traffic.
    - 
Domain: Specify one or more domains associated with the destination traffic.
Note: Requires the client devices to use the UniFi gateway as the DNS server.
    - Region: Specify one or more countries associated with the destination traffic.
  - (Optional) Set a Schedule: Set up a Schedule for when the rule should apply.
- 
Source
- Customize the rule for your desired outcome.
  - Source: Choose a Network, Device, etc.
  - Destination: Choose a specific App, Domain, IP, or Region.
- Click Add.
Once added, your rule will appear in the table and take effect immediately.
Example QoS Configuration
- 
Bandwidth Limit Devices: Restrict specific devices or applications from consuming excessive bandwidth.
  - Objective: Limit
  - Source: Select Device
  - Destination: Select App (e.g., Windows Update)
  - Download/Upload Limits: 10 Mbps each
  - Schedule: Always
- 
Prioritize Streaming Services: Improve video call or streaming stability.
  - Objective: Prioritize
  - Source: Select Device (e.g., work laptop)
  - Destination: Select App (e.g., Zoom, Netflix) or Category (Media Streaming Service)
Important Considerations for QoS Actions
Enabling QoS rules disables hardware offloading on the gateway, which may reduce maximum throughput across all clients and networks—not just those affected by the rule. The performance impact varies based on your gateway model and traffic conditions, and may reduce speeds by 24–45% for traffic exceeding 1 Gbps.
WiFi Speed Limits
- Navigate to Settings > Profiles > WiFi Speed Limits > Create New.
- Name the rule.
- Specify Upload and Download bandwidth limits.
- Click Add.
- Apply the bandwidth limitation to a WiFi SSID: Settings > WiFi > Select WiFi > Enable WiFi Speed Limit > Select Profile.
Switch Port Egress Limiting and DSCP Configuration
Egress Rate Limiting
- Navigate to the specific port, or Port Profile the configuration will apply to.
  - Individual Port: Ports > Select a Port > Manual > Enable Egress Rate Limit.
  - Port Profile: Settings > Profiles > Ethernet Ports > Manual > Enable Egress Rate Limit.
- Specify the limit.
DSCP Tags and IP Precedence
- Navigate to the specific port, or Port Profile the configuration will apply to.
  - Individual Port: Ports > Select a Port > Manual > Enable QoS.
  - Port Profile: Settings > Profiles > Ethernet Ports > Manual > Enable QoS.
- Specify the Protocol of traffic that should be added to this QoS queue.
- Select the Match Type:Match on a specific DSCP or IP Precedence value. None matches on all traffic regardless of value.
- Select a Code.
- Remark Traffic: Optionally remark matched traffic with a specific DSCP or CoS value. If not specified, forwarded traffic keeps its originally set value.
- Specify a Queue: Assign matched traffic to a queue in order to prioritize it over other traffic. The queue range is 0 to 7 with 0 being the lowest priority.
If you’re deploying A/V equipment and want to use pre-configured Port Profiles for Dante, Q-SYS, SDVoE, and other AV systems, refer to our guide on Pro AV Traffic Optimization on UniFi Switches.
802.11e / WMM Specifications for UniFi APs
All UniFi APs support 802.11e / WMM by default with the following mapping of WMM values and DSCP codes. There is no admin action required to take advantage of this traffic prioritization for latency sensitive traffic.
| DSCP Range | WMM Category | 
| 8-9, 11, 13, 15, 17, 19, 21, 23 | AC_BK | 
| 0-7, 10, 12, 14, 16, 18, 20, 22, 25, 27, 29, 31 | AC_BE | 
| 24, 26, 28, 30, 32-43, 45, 47 | AC_VI | 
| 44, 46, 48-63 | AC_VO |
