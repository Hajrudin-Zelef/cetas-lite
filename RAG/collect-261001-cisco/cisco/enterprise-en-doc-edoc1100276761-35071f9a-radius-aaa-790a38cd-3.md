---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd-3
title: "enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd.md
source_anchor: ""
source_lines: [70, 112]
sha256: 46df87ba6a29047ea80f2364eb2df73dcb41e5e7b038426ce9ce739125145ea1
---

# enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd

If multiple servers are configured in the RADIUS server template, the overall status detection time is related to the number of servers and the server selection algorithm. If a user terminal uses the client software for authentication and the timeout period of the terminal client software is less than the summary of all the status detection time, the terminal client software may dial up repeatedly and cannot access the network. If the user escape function is configured, the summary of all the status detection time must be less than the timeout period of the terminal client software to ensure that escape rights can be added to the users.
The following table lists the related commands.
| Command | Description | 
|---|---|
| radius-server { dead-interval dead-interval \| dead-count dead-count \| detect-cycle detect-cycle } | Configures conditions for marking the RADIUS server status as Down during the RADIUS server status detection.  | 
| radius-server max-unresponsive-interval interval | Configures the longest unresponsive interval of the RADIUS server. The default value is 300 seconds. If the interval for sending two consecutive RADIUS Access-Request packets is greater than the value of max-unresponsive-interval, the device marks the RADIUS server status as Down. | 
| radius-server dead-time dead-time | Configures the duration for which the RADIUS server status remains Down. dead-time: Specifies the duration for which the RADIUS server status remains Down after the server status is marked as Down. After the duration expires, the device marks the server status as Force-up. The default value is 5 minutes. | 
After the RADIUS server status is marked as Down, you can configure the automatic detection function to test the RADIUS server reachability.
The automatic detection function needs to be manually enabled. The automatic server status detection function can be enabled only if the user name and password for automatic detection are configured in the RADIUS server template view on the device rather than on the RADIUS server. Authentication success is not mandatory. If the device can receive the authentication failure response packet, the RADIUS server is properly working.
| Server Status | Whether Automatic Detection Is Supported | Time When an Automatic Detection Packet Is Sent | Condition for Switching the Server Status | 
|---|---|---|---|
| Down | Automatic detection is supported by default. | An automatic detection packet is sent after the automatic detection period expires. | If the device receives a response packet from the RADIUS server within the timeout period for detection packets, the device marks the RADIUS server status as Up; otherwise, the RADIUS server status remains Down. | 
| Up | Automatic detection can be enabled using the radius-server detect-server up-server interval command. | An automatic detection packet is sent after the automatic detection period expires. | If the conditions for marking the RADIUS server status as Down are met, the device marks the RADIUS server status as Down; otherwise, the RADIUS server status remains Up. | 
| Force-up | Automatic detection is supported by default. | An automatic detection packet is sent immediately. | If the device receives a packet from the RADIUS server within the timeout period, the device marks the RADIUS server status as Up; otherwise, the device marks the RADIUS server status as Down. | 
On a large-scale network, you are advised not to enable automatic detection for RADIUS servers in Up state. This is because if automatic detection is enabled on multiple NAS devices, the RADIUS server periodically receives a large number of detection packets when processing RADIUS Access-Request packets source from users, which may deteriorate processing performance of the RADIUS server.
After the radius-server testuser command is configured, the dead-time timer configured using the radius-server dead-time command does not take effect.
The following table lists commands related to automatic detection.
| Command | Description | 
|---|---|
| radius-server testuser username user-name password cipher password | Enables the automatic detection function.  | 
| radius-server detect-server interval interval | Specifies the automatic detection interval for RADIUS servers in Down status. The default value is 60 seconds. | 
| radius-server detect-server up-server interval interval | Enables the automatic detection function for the RADIUS server in Up status and configures the automatic detection interval. The default value is 0 seconds; that is, the device does not automatically detect RADIUS servers in Up status. | 
| radius-server detect-server timeout time-value | Specifies the timeout period for automatic detection packets. The default value is 3 seconds. | 
After the device marks the RADIUS server status as Down, you can configure the escape function to make users obtain escape authorization. After the device detects that the RADIUS server status reverts to Up, you can configure the reauthentication function to make users obtain authorization from the server through reauthentication, as shown in Figure 1-13.
For 802.1X authenticated users and MAC address authenticated users, after the RADIUS server status reverts to Up, users exist from escape authorization and are reauthenticated. For Portal authenticated users, after the RADIUS server status reverts to Up, users obtain pre-connection authorization and can be redirected to the Portal server for authentication only if the users attempt to access network resources.
The following table lists the commands for configuring the escape rights upon transition of the RADIUS server status to Down and configuring the reauthentication function, respectively.
| Command | Description | 
|---|---|
| authentication event authen-server-down action authorize { vlan vlan-id \| service-scheme service-scheme-name \| ucl-group ucl-group-name } [ response-fail ] | Configures the escape function upon transition of the RADIUS server status to Down. | 
| authentication event authen-server-up action re-authen | Configures the reauthentication function for users in escape status when the RADIUS server status reverts to Up. | 
Table 1-7 describes types of the CoA/DM packets.
| Packet Name | Description | 
|---|---|
| CoA-Request | When an administrator needs to modify the rights of an online user (for example, prohibit the user from accessing a website), the RADIUS server sends this packet to the RADIUS client, requesting the client to modify the user rights. | 
| CoA-ACK | If the RADIUS client successfully modifies the user rights, it returns this packet to the RADIUS server. | 
| CoA-NAK | If the RADIUS client fails to modify the user rights, it returns this packet to the RADIUS server. | 
| DM-Request | When an administrator needs to disconnect a user, the RADIUS server sends this packet to the RADIUS client, requesting the client to disconnect the user. | 
| DM-ACK | If the RADIUS client has disconnected the user, it returns this packet to the RADIUS server. | 
| DM-NAK | If the RADIUS client fails to disconnect the user, it returns this packet to the RADIUS server. | 
CoA allows the administrator to change the rights of an online user or perform reauthentication for the user through RADIUS after the user passes authentication. Figure 1-14 shows the CoA interaction process.
When a user needs to be disconnected, the RADIUS server sends a DM packet to the device. Figure 1-15 shows the DM interaction process.
The device returns a DM-ACK or DM-NAK packet as follows:
Different from the process in which authorization is performed for an online user or a user proactively goes offline, the server sends a request packet and the device sends a response packet in the CoA/DM process. If CoA/DM succeeds, the device returns an ACK packet. Otherwise, the device returns a NAK packet.
