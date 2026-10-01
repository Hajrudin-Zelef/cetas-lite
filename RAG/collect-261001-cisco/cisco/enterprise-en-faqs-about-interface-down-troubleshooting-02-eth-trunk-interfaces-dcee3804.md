---
id: collect-261001-cisco/cisco/enterprise-en-faqs-about-interface-down-troubleshooting-02-eth-trunk-interfaces-dcee3804
title: "enterprise-en-faqs-about-interface-down-troubleshooting-02-eth-trunk-interfaces--dcee3804"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-faqs-about-interface-down-troubleshooting-02-eth-trunk-interfaces--dcee3804.md
source_anchor: ""
source_lines: [1, 35]
sha256: 584da3edf53fe10de6c1421e0260d885b1ae7718f9618dd19f7d5acbc3c5208d
---

# enterprise-en-faqs-about-interface-down-troubleshooting-02-eth-trunk-interfaces--dcee3804

This post was last edited by NE_Routers at 2018-5-9 16:04.
Common causes are as follows:
Troubleshooting Flowchart
On the network shown in Figure 1, Eth-Trunk interfaces cannot forward traffic.
Figure 1 Networking diagram of Eth-Trunk interfaces
NOTE: 
Interfaces 1 through 3 in this example are GE 1/0/1, GE 1/0/2, and GE 1/0/3, respectively.
The troubleshooting roadmap is as follows:
Figure 2 shows the troubleshooting flowchart. 
Figure 2 Flowchart for troubleshooting Eth-Trunk interfaces' traffic forwarding failures 
Run the display eth-trunk 1 command in any view to check the status of the Eth-Trunk interface.
If member interfaces are Down, troubleshoot the physical interfaces that are Down. For the detailed troubleshooting procedure, see "Physical Interconnection Troubleshooting".
If member interfaces are Up, go to Step 2.
Run the display eth-trunk 1 command to check configurations about the member interfaces of the Eth-Trunk interface on Device A.
Run the display eth-trunk 1 command to check configurations about the member interfaces of the Eth-Trunk interface on Device B.
If the number of Eth-Trunk member interfaces on Device A is different from that on Device B, add a required number of Eth-Trunk member interfaces so that the numbers on the two devices are the same.
If the number of Eth-Trunk member interfaces on Device A is the same as that on Device B, go to Step 3.
Run the display eth-trunk 1 command on Device A and Device B to check configurations of Eth-Trunk interfaces.
This command output shows that the minimum number of active member interfaces is set to 3. However, there are only two member interfaces in the Up state in the Eth-Trunk interface.
If the threshold for the minimum number of active member interfaces is configured, but the actual number of active member interfaces is less than the threshold, set the threshold to a proper value.
If the threshold for the minimum number of Up member interfaces is not configured, go to Step 4.
This command output shows that the maximum number of active member interfaces is set to 3. However, there are four member interfaces in the Up state in the Eth-Trunk interface, which causes the Eth-Trunk interface to go Down.
If the threshold for the maximum number of active member interfaces is configured, but the actual number of active member interfaces is greater than the threshold, set the threshold to a proper value.
If the threshold for the maximum number of Up member interfaces is not configured, go to Step 5.
Member interfaces are faulty, causing LACP negotiation to time out.
Troubleshoot the member interfaces.
The Eth-Trunk interface at one end of the Eth-Trunk link is configured to work in static LACP mode, whereas the Eth-Trunk interface at the other end of the Eth-Trunk link is not.
Correct the configurations of the two ends on the Eth-Trunk link to make them consistent.
If the preceding fault is rectified and LACP negotiation succeeds, the output of the display eth-trunk 1 command is as follows:
If LACP negotiation remains unsuccessfully after the preceding fault is rectified, go to Step 6.
If the Eth-Trunk interface is not configured to work in static LACP mode, go to Step 6.
To check whether BFD is configured, run the display current-configuration command to check the configuration files of the devices at the two ends of the Eth-Trunk link.
If bfd Eth-Trunk bind peer-ip default-ip interface gigabitethernet 1/0/1 is configured, the device has BFD configured.
Run the display bfd session all for-ip command to check information about a BFD session. The State field indicates the BFD session status.
If no device has BFD configured, go to Step 6.
