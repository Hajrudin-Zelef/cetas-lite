---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-aabf97a7-2
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-aabf97a7"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["throughput", "voice"]
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-aabf97a7.md
source_anchor: ""
source_lines: [49, 112]
sha256: 632b83cc2c68ac99ce8f6ace438d933bff36b795c6b3d2045131215c9013ec0c
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-aabf97a7

Click Create a new rule to add a traffic shaping rule. Traffic shaping policies consist of a series of rules that are performed in the order in which they appear in the policy, similar to custom firewall rules. There are two main components to each rule: the type of traffic to be limited or shaped (rule definition), and how that traffic should be limited or shaped (rule actions).
Note: Traffic shaping rules are applied per client flow. For example, setting a limit of 5 Mbps for three different traffic shaping rules will allow 5 Mbps to each client flow matching those respective rules.
Rule Definition
Rules can be defined in two ways:
- You can select from various predefined application categories such as Video & Music, Peer-to-Peer, or Email.
- You can create rules by specifying HTTP hostnames (for example, salesforce.com), port numbers (such as 80), IP ranges (such as 192.168.0.0/16), or IP address range and port combinations (such as 192.168.0.0/16:80).
The rule action is enforced on all traffic that matches the specifications you select. By clicking Add expression, you can create additional specifications for traffic that is shaped according to the same rule action.
Note: The rule for HTTP hostname does not require a catch-all "*" (asterisk) to be prefixed to include subdomains. For example, the hostname 'google.com' will include its subdomains as well.
Rule Actions
Traffic matching specified rule sets can be shaped or prioritized.
- 
    Bandwidth limits can be specified to ignore any limits specified for the whole network, to obey the specified limits, or to apply more-restrictive limits than the network limits. Use the bandwidth slider control to choose the appropriate limit for each type of traffic. To specify asymmetric limits on uploads and downloads, click details next to the bandwidth slider control.
- 
    Priority can be set to High, Normal, or Low, allowing the MX series to prioritize a given network flow relative to the rest of the network traffic. Note that Realtime is reserved for traffic tagged with a DSCP bit value mapping to EF46 only. 
  - 
        Firmware 18.2 and above: Class Based Weighted Fair Queueing with Deficit Round Robin is enforced on token bucket allocations. Allocation ratios are as follows: 
    - 
            Realtime: 8
    - 
            High: 4
    - 
            Normal: 2
    - 
            Low: 1
  - 
            
- 
        
With deficit round robin, queues that do not actively have traffic to send will see their tokens allocated to the next queue in line. This prevents starvation for lower priority queues while continuing to prioritize VoIP traffic and other manually defined high-priority traffic.
- 
    
  - 
        Firmware 18.1 and lower: Strict Priority Queueing is enforced with the following ratios on token bucket allocations 
    - 
            High: 4/7
    - 
            Normal: 2/7
    - 
            Low: 1/7
  - 
            
- 
        
- 
    Quality of Service (QoS) prioritization can be applied to Layer 3 traffic. To prioritize traffic at Layer 3, select a value for the DSCP tag in the IP header for all incoming and outgoing IP packets. This also affects the Wi-Fi Multimedia (WMM) priority of the traffic.
For the Priority feature to work as desired, ensure that uplink throughput settings are accurate.
For QoS prioritization to work as desired, ensure that upstream networking equipment also supports QoS prioritization.
Creating a Sample Traffic-Shaping Rule
Here is an example of how to set up a traffic shaping policy with multiple traffic-shaping rules. For additional examples, refer to our Simple Traffic Shaping Strategy article.
To prioritize VoIP and minimize peer-to-peer traffic and gaming, create a new traffic-shaping policy by following the steps below:
- In the Rule #1 Definition pull-down menu, choose VoIP & video conferencing.
- Under Bandwidth limit, choose Ignore network limit.
- In the Priority pull-down menu, choose High.
- Under DSCP tagging, choose 7 (WMM Voice).
- Click Add a new shaping rule.
- In the Rule #2 Definition pull-down menu, choose Peer-to-peer (P2P).
- Click Add an expression.
- In the new pull-down menu, choose Gaming.
- In the Bandwidth limit section, click Choose a limit and use the slider to choose a low throughput (the minimum is 20 Kb/s).
- Save your changes by clicking Save Changes at the bottom of the page.
Web cache
This option is not available on the MX64, MX64W, MX65, MX65W, MX67, MX67W, MX67C, MX68, MX68W, MX68CW, MX75, MX85, MX95, MX105, Z3, Z3C, Z4, Z4C devices.
When HTTP content caching is enabled, the MX will cache web content on its solid-state drive (SSD). This can improve end-user experience by reducing page load times and file download times for frequently accessed web content. Web caching only works for static HTTP content, so it will not be able to cache sites such as YouTube.
This feature is recommended only for sites with limited bandwidth. Locations with over 20 Mbps bandwidth will likely not benefit from content caching.
