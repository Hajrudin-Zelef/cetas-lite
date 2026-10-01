---
id: collect-261001-meraki/meraki/architectures-and-best-practices-cisco-meraki-best-practice-design-best-practice-7806cb7c-2
title: "architectures-and-best-practices-cisco-meraki-best-practice-design-best-practice-7806cb7c"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/architectures-and-best-practices-cisco-meraki-best-practice-design-best-practice-7806cb7c.md
source_anchor: ""
source_lines: [161, 245]
sha256: 685dfef83fb27eaf8d3a20bb1977fd9be639eafb769fdd9440297c545f2df6f2
---

# architectures-and-best-practices-cisco-meraki-best-practice-design-best-practice-7806cb7c

Template Firewall Rules
When configuring layer 3 firewall rules, CIDR notation, VLAN Objects and Network Objects can be used. The VLAN Object is used when the entire subnet needs to be specified whereas CIDR notation is used when more flexibility is needed to specify the subnets.
- 
    Go to Security & SD-WAN > Configure > Firewall > Layer 3, click Add a rule
- 
    Choose the policy, specify if the rule matched should be allowed or denied
- 
    Select the protocol to match in outbound traffic
- 
    Specify the IP address or range using CIDR notation to match the outbound traffic. Note that also the name of the VLAN can be chosen as well
- 
    Choose the Src/dst port to match in outbound traffic
IP Offset
If you wish to specify a single IP within a VLAN using a VLAN Object an IP Offset can be used to calculate the IP address that the rule will apply to based on the network address for the specified VLAN. The IP address will be calculated using the Network Address + Offset and not the VLAN Interface IP + Offset.
For example, if you have a firewall rule containing the source Users.10 representing the Users VLAN with an IP Offset of 10 and the subnet of the Users VLAN is 192.168.100.0/24 then the source will be interpreted as 192.168.100.10.
IP Offset is not available when configuring a dual-stack firewall rule, to use an IP Offset it must be either an IPv4-only or IPv6-only rule.
Template SD-WAN Policies
- 
    SD-WAN policies can be configured to control and modify the flows for specific VPN traffic. You can have a specific type of traffic go over one Uplink over the other. 
- 
    Go to Security & SD-WAN > Configure > SD-WAN & traffic shaping > SD-WAN policies > VPN traffic, and choose Add a preference
- 
    You'll be prompted with the Uplink selection policy dialog box. From this box, you can define the type of traffic that should adhere to the policy on the Traffic filters section. You can either add Custom expressions to select traffic based on Protocol/Source/Destination criteria, or you can select traffic based on pre-defined applications. 
  - 
        To add a custom expression to select traffic. 
    - 
            Choose Add +
    - 
            The Custom expressions option should already be selected.
    - 
            Choose the Protocol. You can choose either TCP, UDP, ICMP or Any
    - 
            Choose Source to define the source address criteria. You can select one of the following: 
      - 
                You can choose Any
      - 
                You can type in the source in CIDR format( eg: 10.0.0.0/8), and then choose Add
      - 
                You can choose a VLAN from the drop-down menu with the list of VLANs and then choose Add VLAN
      - 
                You can choose a VLAN from the drop-down menu with the list of VLANs and then click Host, type in the last octet of the host address, then choose on Add host
    - 
                
    - 
            Choose the Src port. The Source port could be 'Any', a port number (eg: 2000), or a port range (eg: 2000-3000) within 1-65535.
    - 
            Click Destination to define the source address criteria. You can select one of the following: 
      - 
                You can choose Any
      - 
                You can type in the source in CIDR format( eg: 10.0.0.0/8), and then choose Add
      - 
                You can choose a VLAN from the drop-down menu with the list of VLANs and then choose Add VLAN
      - 
                You can choose a VLAN from the drop-down menu with the list of VLANs and then choose Host, type in the last octet of the host address, then choose Add host
    - 
                
    - 
            Choose the Dst port. The Destination port could be 'Any', a port number (eg: 2000), or a port range (eg: 2000-3000) within 1-65535.
  - 
            
  - 
        To add a pre-defined application to select traffic. 
    - 
            Select the application type from the menu and then the interesting application in question from the sub-menu (e.g., VoIP & video conferencing > Webex)
    - 
            Add all the applications which you want them to adhere to the policy and then choose Add+ to exit the applications menu
  - 
            
- 
        
- 
    Under the Policy section, you can select one of the following as the Preferred uplink 
  - 
        WAN1 or WAN2. If you choose WAN1 or WAN2, you'll have the opportunity to configure failover criteria under Fail over if drop-down menu. You can select either Poor performance and then choose one of the performance classes from the Performance class drop-down menu or you can choose Uplink down. By default, VoIP is the only pre-defined performance class. Any additional performance classes have to be defined under Security & SD-WAN > Configure > SD-WAN & traffic shaping > SD-WAN policies > Custom performance classes.
  - 
        Best for VoIP. The uplink that is best for VoIP traffic will be chosen.
  - 
        Load balance. The WAN Appliance will balance traffic across the uplinks that meet the performance class selected from On uplinks that meet performance class drop-down menu.
  - 
        Global preference. The uplink will be chosen based on the configuration under Security & SD-WAN > Configure > SD-WAN & traffic shaping > Uplink selection > Global preferences.
  - 
        Global bandwidth limits. When configuring this setting, keep in mind that the MX will apply and return the speed settings configured here at the template level as opposed to the network level. This may not reflect the affected device's hardware capabilities. Refer to https://developer.cisco.com/meraki/a...ink-bandwidth/ for more information
- 
        
