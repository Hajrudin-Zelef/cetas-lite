---
id: collect-261001-general-networking/general-networking/infrastructure-security-and-segmentation-1
title: "infrastructure-security-and-segmentation"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["exploit", "memory"]
source: docs/RAG/collect-261001-general-networking/infrastructure-security-and-segmentation.md
source_anchor: ""
source_lines: [1, 60]
sha256: 3eba12a5a733bceb34a6188e9591a6d1692e00cc70409e144ebfff66b0cf8403
---

# infrastructure-security-and-segmentation

## Securing the Data Plane

The purpose of a network is to move data from its source to its destination. Network devices are geared toward that one function alone, and the data plane is where it is executed. Hence, it should not be surprising that most attacks are targeted toward this plane, and most security is also implemented here.

The difference between the data plane of a Layer 2 device and that of a Layer 3 device is far more distinct than in any other plane. The security objectives, challenges, and solutions are very different between those two layers.

Layer 2 security is primarily focused on unauthorized or spoofed access, segmentation, and resource protection. On the other hand, Layer 3 security is primarily focused on the data itself and ensuring that the right kind from the right source is allowed. Even with the differences, some general security considerations apply to the data plane of both layers:

- **Zero trust and minimum required access:** This principle applies to the data plane also. No traffic or its source should be trusted, and each device should be given the minimum required access. Access control and segmentation of traffic play a crucial role in this.
- **Protocol validation:** Most of the common protocols used in the data plane have absolutely no security built in. They are usually easy to spoof and exploit. In fact, some of the most effective attacks are executed by exploiting transport protocols. Measures should be taken to prevent attacks using these protocols.
- **Audit trail collection:** Cisco devices provide various methods to generate and export audit trails in the form of logs. Common examples of these are access lists with logging and NetFlow. These should be enabled, exported to central collectors, and analyzed for incidence response and policy correction.

### Security at the Layer 2 Data Plane

The data plane on a switch is focused on switching frames in the local network segment. To accomplish this, it needs to know where a given MAC address is located and, sometimes, what IP address it is mapped to. To keep track of these two important details, a switch uses some protocols and tables, such as the CAM and ARP tables. The next few sections discuss the inherent vulnerabilities in the switching mechanism and the security features used to protect them.

#### The CAM Table and Port Security

The content addressable memory (CAM) table is a list of all MAC addresses known to the switch, along with their physical location and VLAN ID. When an endpoint sends a frame, its MAC address, source interface, and VLAN ID are recorded in the CAM table. The switch searches this table to find the location of a MAC address when it needs to deliver a frame. If the destination MAC address is not found in the CAM table, the frame is flooded out all interfaces except the one where the frame was received.

The process of learning a MAC address and its location has no security built into it. Anyone can send a frame with a spoofed source MAC address to cause the switch to update its CAM table. This can be used for a MITM attack such that frames are switched to the spoofed location instead of the intended destination. The traffic continues to flow to the attacking host until the original endpoint sends a frame, causing the switch to update the CAM table again. This is called a *MAC spoofing attack*.

Another problem with the CAM table is its limited size. An attacker can send hundreds of thousands of frames with difference source MAC addresses within a short period of time. This causes the switch to add these spoofed MAC addresses in the CAM table. Once the CAM table has reached its limits, it stops learning new MAC addresses and starts flooding frames with unknown MAC addresses out all its interfaces. This can cause severe degradation of the network to the point of a DoS situation. This attack is called *CAM flooding*. Each switch has a different CAM table size limit. The size of the CAM table on a Cisco IOS switch can be found by using the **show mac address-table count** command, as shown in Example 2-38.

#### **Example 2-38** *MAC Address Table Limit*

`SW1#**sh mac address-table count**
Mac Entries for Vlan 1:
---------------------------
Dynamic Address Count  : 17
Static  Address Count  : 37
Total Mac Addresses    : 54
Total Mac Address Space Available: 16303`

CAM table attacks can be prevented with the port security feature in Cisco switches. With port security, an interface can be configured to learn only a limited number of MAC addresses. The addresses an interface can learn can be dynamic, static, or a combination of these. Port security can be enabled on an interface with the **switchport port-security** command.

The following MAC address learning restrictions can be configured with port security:

- **Static Secure MAC Addresses:** With this option, an interface is configured to learn only the manually configured MAC addresses. This option limits the devices that can connect to a given interface and is best used for static devices such as printers, cameras, servers, and other devices that do not change locations. This also prevents MAC spoofing by preventing the configured MAC address from connecting to another switch port. Static secure MAC addresses can be configured on an interface with the**switchport port-security mac-address***mac-address* command.
- **Dynamic Secure Addresses:** This option configures the maximum addresses that an interface will learn. The addresses are learned dynamically as devices connect to the interface. This option can be used to prevent CAM flooding attacks and the addition of unauthorized switches or hubs. The number of maximum allowed addresses can be configured on an interface with the**switchport port-security maximum***number* command.
- **Sticky Secure MAC Addresses:** Configuring each interface to allow a static MAC address is administratively prohibitive and restrictive to users. The Sticky Secure MAC Addresses for port security solves this by dynamically learning the MAC addresses that connect to it and storing them in the configuration as static. Sticky secure learning can be enabled on an interface with the**switchport port-security mac-address sticky** command.

The three options can be combined to meet different requirements. For example, an interface can be configured to learn a maximum of two addresses, while one is statically defined, as shown in Example 2-39. Such a configuration allows the interface to learn one dynamic address along with the statically configured one.

#### **Example 2-39** *Port Security with Static and Dynamic Address Learning*

`SW1(config)#**interface Gi0/5**
SW1(config-if)#**switchport port-security mac-address 1001.1001.1001**
SW1(config-if)#**switchport port-security maximum 2**`

Example 2-40 shows another combination in which an interface is configured for a maximum of three addresses with sticky learning enabled. This means the switch will learn and save the MAC addresses of the first three devices that connect to it. After that, only those three devices will be able to connect to that interface.

#### **Example 2-40** *Three Sticky Secure Addresses*

`SW1(config)#**interface Gi0/9**
SW1(config-if)#**switchport port-security maximum 3**
SW1(config-if)#**switchport port-security mac-address sticky**`

Port security defines three actions that can be taken when a violation occurs. An interface can be configured to take one of these actions:

