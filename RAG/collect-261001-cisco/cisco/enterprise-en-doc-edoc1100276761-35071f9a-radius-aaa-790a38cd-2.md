---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd-2
title: "enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd.md
source_anchor: ""
source_lines: [43, 69]
sha256: 280d7e54753c15c47da6400603303f534c5ae2a7f7a021ce7b82913c1a62ee31
---

# enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd

The device does not definitely mark the status of the server that does not respond as Down. The device marks the server status as Down only if the corresponding conditions are met.
For the RADIUS server status introduction and conditions for a device to mark the server status as Down, see RADIUS Server Status Detection.
RADIUS packet retransmission discussed here applies only to a single server. If multiple servers are configured in a RADIUS server template, the overall retransmission period depends on the retransmission interval, retransmission times, RADIUS server status, number of servers, and algorithm for selecting the servers.
| Command | Description | 
|---|---|
| radius-server retransmit retry-times | Specifies the retransmission times. The default value is 5. | 
| radius-server timeout time-value | Specifies the retransmission interval. The default value is 2 seconds. | 
In addition, the algorithm for selecting a RADIUS server can be set to the single user-based or packet-based algorithm. If the algorithm for selecting a RADIUS server is set to the single user-based algorithm, authentication server information is saved in the authentication phase, and the device preferentially sends an accounting request to the accounting server in the accounting phase when the authentication server is also the accounting server. If the algorithm for selecting a RADIUS server is set to the packet-based algorithm, authentication server information is not saved in the authentication phase, and the accounting server is reselected in the accounting phase, which may result in that authentication and accounting for a user is not performed on the same server.
The primary and secondary roles are determined by the weights configured for the RADIUS authentication servers or RADIUS accounting servers. The server with the largest weight is the primary server. If the weight values are the same, the earliest configured server is the primary server. As shown in Figure 1-10, the device preferentially sends an authentication or accounting packet to the primary server among all servers in Up status. If the primary server does not respond, the device then sends the packet to the secondary server.
If this algorithm is used and a device sends an authentication or accounting packet to a server, the device selects a server based on the weights configured for the RADIUS authentication servers or RADIUS accounting servers. As shown in Figure 1-11, RADIUS server1 is in Up status and its weight is 80, and RADIUS server2 is also in Up status and its weight is 20. The possibility for the device to send the packet to RADIUS server1 is 80% [80/(80 + 20)], and that for RADIUS server2 is 20% [20/(80 + 20)].
Regardless of which algorithm is used, if all the servers in Up status do not respond to a packet sent by a device, the device retransmits the packet to a server among the servers whose status is originally marked as Down (to which the device has not sent any authentication or accounting packets) based on the server weight. If the device does not receive any response in the current authentication mode, the backup authentication mode is used, for example, local authentication mode. The backup authentication mode needs to be already configured in the authentication scheme. Otherwise, the authentication process ends.
Availability and maintainability of a RADIUS server are the prerequisites of user access authentication. If a device cannot communicate with the RADIUS server, the server cannot perform authentication or authorization for users. To resolve this issue, the device supports the user escape function upon transition of the RADIUS server status to Down. To be specific, if the RADIUS server goes Down, users cannot be authorized by the server but still have certain network access rights.
The user escape function upon transition of the RADIUS server status to Down can be enabled only after the device marks the RADIUS server status as Down. If the RADIUS server status is not marked as Down and the device cannot communicate with the RADIUS server, users cannot be authorized by the server and the escape function is also unavailable. As a result, users have no network access rights. Therefore, the device must be capable of detecting the RADIUS server status in a timely manner. If the device detects that the RADIUS server status transitions to Down, users can obtain escape rights; if the device detects that the RADIUS server status reverts to Up, escape rights are removed from the users and the users are reauthenticated.
The mechanism of checking the RADIUS server status by using accounting packets is the same as that of checking the RADIUS server status by using request packets.
After the automatic detection function is enabled, automatic detection is classified into the following conditions depending on differences of the RADIUS server status.
| Server Status | Whether the RADIUS Server Is Available | Reason for the Status | 
|---|---|---|
| Up | The RADIUS server is available. |  | 
| Down | The RADIUS server is unavailable. | The conditions for marking the RADIUS server status as Down are met. | 
| Force-up | When no RADIUS server is available, the device establishes a connection with a RADIUS server in Force-up state. | The timer specified by dead-time expires. | 
The RADIUS server status is initially marked as Up. After a RADIUS Access-Request packet is received and the conditions for marking the RADIUS server status as Down are met, the RADIUS server status transitions to Down. The RADIUS Access-Request packet that triggers the server status transition can be sent during user authentication or constructed by the administrator. For example, the RADIUS Access-Request packet can be a test packet sent when the test-aaa command is run or detection packet sent during automatic detection.
Whether the status of a RADIUS server can be marked as Down depends on the following factors:
The device marks the RADIUS server status as Down during the RADIUS server status detection.
After the system starts, the RADIUS server status detection timer runs. If the device does not receive any packet from the RADIUS server after sending the first RADIUS Access-Request packet to the server and the condition that the number of times the device does not receive any packet from the server (n) is greater than or equal to the maximum number of consecutive unacknowledged packets (dead-count) is met in a detection interval, a communication interruption is recorded. If the device still does not receive any packet from the RADIUS server, the device marks the RADIUS server status as Down when recording the communication interruption for the same times as the detection interval cycles.
If the device does not record any communication interruption in a detection interval, all the previous communication interruption records are cleared.
The device marks the status of a RADIUS server as Down if no response is received from the server for a long period of time.
If the user access frequency is low, the device receives only a few RADIUS Access-Request packets from users, conditions for marking the RADIUS server status as Down during the RADIUS server status detection cannot be met, and the interval for sending two consecutive unacknowledged RADIUS Access-Request packets is greater than the value of max-unresponsive-interval, the device marks the RADIUS server status as Down. This mechanism ensures that users can obtain escape authorization.
