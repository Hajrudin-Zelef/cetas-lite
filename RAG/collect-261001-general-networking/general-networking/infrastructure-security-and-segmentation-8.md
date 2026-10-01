---
id: collect-261001-general-networking/general-networking/infrastructure-security-and-segmentation-8
title: "infrastructure-security-and-segmentation"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/infrastructure-security-and-segmentation.md
source_anchor: ""
source_lines: [564, 621]
sha256: f7a972f4500e27a1463e02ef43331fc3982b1df63cd3444ba8fb2888f3382f67
---

# infrastructure-security-and-segmentation

NBAR uses its Protocol Discovery function to classify traffic. With Protocol Discovery, you can discover any protocol or application traffic that is supported by NBAR and obtain statistics that are associated with that protocol. Protocol Discovery maintains the following per-protocol statistics for enabled interfaces:

- Total number of input packets and bytes
- Total number of output packets and bytes
- Input bit rates
- Output bit rates

To use NBAR in QoS policies, Protocol Discovery needs to be enabled on router interfaces with the **ip nbar protocol-discovery** command. The collected data from Protocol Discovery can be verified with the **show ip nbar protocol-discovery interface** *interface* command.

While an in-depth discussion on QoS is beyond the scope of this book, as a quick refresher, QoS is configured using the Modular QoS CLI (MQC). Earlier in this chapter, you saw MQC used to configure CoPP. The MQC is a CLI that allows you to define traffic classes (class maps), create and configure traffic policies (policy maps), and attach the traffic policies to interfaces. Using MQC to configure NBAR consists of the following steps:

- **Step 1. Defining a traffic class with a class map:** Class maps are created using the**class-map***name* command and contain a**match protocol** statement to classify traffic using NBAR. Hundreds of applications and protocols are available, with numerous sub-options to create granular classification.
- **Step 2. Defining a traffic policy:** A traffic policy or policy map defines what action to take on defined traffic classes. Policy maps are created using the**policy-map***name* command. Within a policy map, multiple class maps can be associated with actions such as**drop, police** , and**shape** .
- **Step 3. Attaching the traffic policy to an interface:** A traffic policy or policy map needs to be attached to an interface in order to be effective. This can be done using the**service-policy** {**input** |**output** }*policy-map-name* command.

Example 2-60 shows how NBAR is used with MQC to classify and drop P2P file-sharing traffic. This example is based on the network shown in Figure 2-4. Interface Gi1 on R1 is configured for NBAR Protocol Discovery, and a policy map is applied to that interface to drop P2P file-sharing traffic.

#### **Example 2-60** *Configuring NBAR to Drop P2P File-Sharing Traffic*

`R1(config)#**interface Gi1**
R1(config-if)#**ip nbar protocol-discovery**
R1(config-if)#**exit**
R1(config)#**class-map p2p-traffic**
R1(config-cmap)#**match protocol bittorrent**
R1(config-cmap)#**match protocol fasttrack**
R1(config-cmap)#**match protocol gnutella**
R1(config-cmap)#**match protocol kazaa2**
R1(config-cmap)#**exit**
R1(config)#**policy-map drop-p2p**
R1(config-pmap)#**class p2p-traffic**
R1(config-pmap-c)#**drop**
R1(config-pmap-c)#**exit**
R1(config-pmap)#**exit**
R1(config)#**interface Gi1**
R1(config-if)#**service-policy output drop-p2p**
R1(config-if)#**service-policy input drop-p2p**`

#### TCP Intercept

TCP Intercept is an important security feature on Cisco routers that is used to protect TCP servers from SYN-flooding attacks. A SYN-flooding attack occurs when a hacker floods a server with requests for TCP connections sourced from spoofed addresses. Because these messages come from spoofed or unavailable addresses, the connections do not get established, and the server is forced to keep them open for a while. The resulting volume of unresolved open connections eventually overwhelms the server and can cause it to deny service to valid requests.

The TCP Intercept feature protects the servers, such as web servers, in a network from such SYN-flooding attacks. It does so by intercepting and validating TCP connection requests. In intercept mode, the router intercepts TCP synchronization (SYN) packets from clients to servers. It establishes a connection with the client on behalf of the destination server, and if that connection is successful, it establishes a connection with the server on behalf of the client and knits the two half-connections together transparently.

With the router intercepting all TCP connection requests, a SYN-flooding attack never reaches the servers itself. To protect itself from being overwhelmed in case of an attack, the router uses aggressive timeouts of half-open or embryonic connections.

Another way to protect the server and the router both is to use TCP Intercept in *watch* mode. In this mode, the router does not intercept TCP connections but passively watches each connection request. If the request is not completed within a configured time interval, the router intervenes and terminates the connection.

Before configuring TCP Intercept, an extended access list containing a list of servers to protect needs to be defined. You can choose to allow the whole inside network to be protected, but that may cause the router to be overwhelmed. It is recommended that you define the critical servers that require the protection.

After defining the access list, TCP Intercept can be enabled with the **ip tcp intercept list** *access-list* global configuration command. The TCP Intercept mode can be configured with the **ip tcp intercept mode** {**intercept**|**watch**} command. Example 2-61 shows how TCP Intercept is enabled to protect three servers in watch mode.

#### **Example 2-61** *Configuring TCP Intercept*

`R1(config)#**access-list 105 permit tcp any host 10.1.1.10**
R1(config)#**access-list 105 permit tcp any host 10.1.1.11**
R1(config)#**access-list 105 permit tcp any host 10.1.1.12**
R1(config)#**ip tcp intercept list 105**
R1(config)#**ip tcp intercept watch**`
