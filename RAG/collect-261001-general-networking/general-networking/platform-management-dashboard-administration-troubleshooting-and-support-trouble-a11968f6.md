---
id: collect-261001-general-networking/general-networking/platform-management-dashboard-administration-troubleshooting-and-support-trouble-a11968f6
title: "platform-management-dashboard-administration-troubleshooting-and-support-trouble-a11968f6"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/platform-management-dashboard-administration-troubleshooting-and-support-trouble-a11968f6.md
source_anchor: ""
source_lines: [1, 76]
sha256: 54606a1216b222389d3912560232e269a0b2b92662ba59a2a5e972b7db4a407b
---

# platform-management-dashboard-administration-troubleshooting-and-support-trouble-a11968f6

This article describes how to use the Packet capture tool in Dashboard to troubleshoot client-side DHCP issues on your network. This requires Wireshark installed in order to open PCAP file that will be downloaded from Dashboard. In this example, a client device is connected to a Cisco Meraki switch port but is unable to get IP information from the DHCP server.


The DHCP handshake is illustrated in Figure 1 below.

**Figure 1: DHCP Handshake.**

1. Navigate to **Monitor > Packet capture.** In a combined network you will want to navigate to **N****etwork-wide > Packet capture** and select which Cisco Meraki Appliance you would like to capture off of:

**Figure 2: Packet Capture tool**


2. Set up your packet capture tool to gather data from the switch uplink port and the client on the same switch. You will need to set your packet capture tool to download file to PCAP file and capture for 60 seconds. If your DHCP server is connected to the same switch you will want to capture off the specific port the client is connected to instead of the uplink.

**Figure 3: Packet capture view on Dashboard**


3. Start capture.

4. Open the Command prompt from the client machine and perform an **ipconfig  /release** then **ipconfig  /renew.** This will force the client machine to perform a DHCP broadcast. Perform this step a few times in order to generate traffic that is being captured by the packet capture tool 

**Figure 4: IP release and renew**




5. Open the saved PCAP file which has been downloaded from Dashboard with Wireshark and enter the **bootp** display filter, click **Apply**. This filter will show any part of the DHCP process in the capture:  **DHCP discover, DHCP offer, DHCP request, DHCP acknowledge**.This will give you great insight of where the DHCP process is potentially failing. The figure bellow shows the four-way DHCP process as well as the the **Transaction ID** which is very important as this groups each DHCP process/handshake together:




**Figure 5: Packet capture illustrating the DHCP handshake**


 

6. Validate your host device is sending out a discovery broadcast, and 2-) we want to validate the DHCP server is responding back. Figure 5 shows my client device "Source: f0:de:f1:a3:5d:d6 " sending out a broadcast to "Destination ff:ff:ff:ff:ff:ff " and Figure 6 shows the available DHCP server "Source: 00:18:0a:42:3e:b5 " responding back to my client with a DHCP Offer.

**Figure 6:  DHCP Discovery message** 

**Figure 7:  DHCP offer message** 



The two most common problems are: 1-) The client device never receives a response from the DHCP server or 2-) The client device gets an IP from the wrong DHCP server.


For the latter, please take a look at the following KB - Tracking down a rogue DHCP server 




If you are not seeing a response back, here are some other things worth looking into:

- Verify 802.1q is correct setup on the Switch port.
- Ensure the switches uplink has the correct allowed VLANs if using VLANs and VTP within your infrastructure.
- Run packet captures off your other devices along the path between the client and the DHCP server
- Run wireshark on your DHCP server to verify you are seeing the clients DHCP discover making it to your server and that the response has the correct destination MAC address.
- Check routing setup on your Layer 3 devices to ensure the client has the correct path setup to the DHCP server.
- Review the DHCP server for leases problems, exhausted DHCP pool or out right DHCP services issues.

Note: This Article can also be applied to other Cisco Meraki Devices

MX: Set Up your packet capture tool to capture off the LAN side of the MX

**Figure 8:**  **Packet capture view on Dashboard**

MR: There are two potential places to capture data:

- If your SSID is running in bridge mode you will want to capture off the wired connection

**Figure 9:**  **Packet capture view on Dashboard**

If your SSID is configured to run in NAT mode you will want to setup your capture on the Wireless portion

**Figure 10:**  **Packet capture view on Dashboard**
