---
id: collect-261001-general-networking/general-networking/ms-port-and-vlan-configuration-switch-ports-19ab0597-1
title: "ms-port-and-vlan-configuration-switch-ports-19ab0597"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "training", "voice"]
source: docs/RAG/collect-261001-general-networking/ms-port-and-vlan-configuration-switch-ports-19ab0597.md
source_anchor: ""
source_lines: [1, 74]
sha256: 696bb726e3ffb19ed0b2a92c3211d83d92327465c04cc0f136ec0647793ef993
---

# ms-port-and-vlan-configuration-switch-ports-19ab0597

Switch Ports
Learn more with these free online training courses on the Meraki Learning Hub:
Making Configuration Changes
On the Switching > Monitor > Switch Ports page, administrators can name ports, turn ports on/off, enable spanning tree (RSTP), define port types (access/trunk), and specify VLANs (data and voice). It is recommended to keep the total switch port count in a network to fewer than 8000 ports for reliable loading of the switch port page.
Switchport page may have issues loading if a dashboard network exceeds 400 switches per network.
Editing a port(s)
In order to make changes to a port or port group on an MS switch:
- Select the port or ports to be configured by checking their perspective check box(es).
- Choose Edit and make the desired changes. See the "Port configuration" section for all configurable items.
- Once the changes have been made, save them by selecting Update. This will instantly push the changes to the MS switches in the network.
Port configuration
The following fields are configurable on each switch port.
- Name: Description of the port.
- Tags: Labels that can be used to identify this port or a group of ports.
- Port status: Enable/Disable the port.
- Stacking: Enable flexible stacking on this port.
- RSTP: Rapid Spanning Tree Protocol (RSTP) and STP guards can be configured at the port level. For more information on port level spanning tree configuration, check out our article on Configuring Spanning Tree on Meraki Switches.
- PoE: Available on PoE switches only. Enable/Disable Power over Ethernet on this port.
- Link negotiation: Select the desired link speed.
Half Duplex is not supported on MS350 and MS355 series switches.
Cloud-managed Catalyst switches use automatic link negotiation on SFP/SFP+ ports, with Auto Negotiate as the only available option.
- Port Schedule: Apply a port schedule policy. Learn how to use port scheduling here.
- Port Isolation: Enabling this feature prevents any isolated port from communicating with other isolated ports.
- Trusted DAI: Enable/Disable the trusted status for Dynamic ARP Inspection.
- UDLD: Alert/Enforce Unidirectional Link Detection on the port.
- Type: Switch ports can be configured as one of two types:
    
  - Trunk: Configuring a trunk port will allow the selected port to accept/pass 802.1Q tagged traffic.  This type is usually used for connections to other switches or access points.
        
    - Access Policy: Apply a restriction policy to this port
    - Native VLAN: The switch will send traffic for this VLAN as untagged. (Note: MS390s/C9000s will use VLAN 1 as default if no native VLAN is specified).
    - Allowed VLANs: Only these VLANs will be able to traverse this link.
    - 
            PortFast Trunk: Allows a trunk port to more quickly transition to the STP Forwarding state. Note: This feature is supported starting in IOS XE 26.1.1. To utilize portfast trunk, please upgrade to IOS XE 26.1.1 or any later release. 
      - 
                Before enabling this feature, make sure that there are no loops in the network between the trunk port and the connected end-device as it may lead to network instability.
      - 
                Enabling this feature allows the IOS XE portfast trunk to be enabled on trunk ports, applying the spanning-tree portfast trunk option to the port(s) in question.
    - 
                
- Trunk: Configuring a trunk port will allow the selected port to accept/pass 802.1Q tagged traffic.  This type is usually used for connections to other switches or access points.
        
- Access: Configuring an access port will place all traffic on its defined VLAN and will only pass untagged traffic.  This type is usually used for connections to end-users.
    
  - Access Policy: Apply a restriction policy to this port.
        
    - Open: All devices will be able to access this port.
    - MAC allow list: MAC allow list allows users to enter up to 20 MAC addresses they want to be permitted to pass traffic on a particular interface, restricting traffic on that interface to the configured MAC addresses only. MAC addresses may be entered in aa:bb:cc:dd:ee:ff format or aaaa.bbbb.cccc format. (NOTE: Starting with MS18 MS classic switches will allow you to set a maximum number of MAC addresses allowed in the Allow MAC list. This is NON-sticky, for sticky please use the sticky MAC allow list option. )
            Configuration: 
      - 
                Navigate to Switching > Monitor > Switches and select your switch. In the mimic panel, select the port to configure and then click the pencil icon in the Configuration section.
      - Navigate to the Access Policy drop-down field and Select MAC allow List
      - 
                Enter up to 20 MAC addresses to allow on the interface and click Update
    - 
                
- Access Policy: Apply a restriction policy to this port.
        
- Sticky MAC allow list:
    Like MAC allow list, Sticky MAC allow list also allows users to configure between 1-20 MAC addresses allowed to pass traffic on a particular switch port, but Sticky MAC also allows MAC addresses to be dynamically learned on an interface. Users can either program the allowed MAC addresses statically into the Allowed listed MACs list, or allow for the switchport to dynamically learn the MACs. For example if you set the number of Sticky MACs to 5 and program 1 in the allow list, the next 4 MACs dynamically learned will be programmed into the stick MAC list. Any MACs learned after this will be denied access to that specific port. Configuration: 
  - 
        Navigate to Switching > Monitor > Switches and select your switch. In the mimic panel, Select the port to configure and then click the pencil icon in the Configuration section.
  - 
        Navigate to the Access Policy drop-down field and Select Sticky MAC allow List
  - 
        Enter a maximum number (between 1-20) of Sticky MACs to allow on the interface
  - 
        (Optional) Enter any static sticky MACs to allow on the interface in the Allow Listed MACs
- 
        
In this example the number of sticky MACs is set to 5 with one sticky MAC being hard coded. The switch will now learn the next 4 MACs that are seen on this switchport dynamically to make the total of 5 Sticky MACs. It can take up to 5 minutes for the learned MAC to appear in the dashboard.
Sticky MAC addresses persist through a device reboot.
- User-defined access policy: Administrators may define a policy for authentication via 802.1x or MAB.  Learn more about access policies here.
    
