---
id: collect-261001-general-networking/general-networking/switching-ms-switches-operate-and-maintain-monitoring-and-reporting-packet-captu-182aac87
title: "switching-ms-switches-operate-and-maintain-monitoring-and-reporting-packet-captu-182aac87"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["training"]
source: docs/RAG/collect-261001-general-networking/switching-ms-switches-operate-and-maintain-monitoring-and-reporting-packet-captu-182aac87.md
source_anchor: ""
source_lines: [1, 103]
sha256: 2dd785d1814925ded7957e26558b96384890dfb0dea0a9561874937364201c79
---

# switching-ms-switches-operate-and-maintain-monitoring-and-reporting-packet-captu-182aac87

How to Capture Packets Using Port Mirroring on MS
Overview
Workstations in promiscuous mode sniff LAN packets within their broadcast domain. A workstation connected to Cisco Meraki switches captures these packets through port mirroring.
This article explains how to capture traffic passed by an MS switch using the following steps:
- Enable port mirroring on your switch.
- Connect a workstation to your destination port.
- Capture packets in promiscuous mode.
The most effective way to capture traffic passed on a given switchport is to mirror that port to another available port, so the switch sends all traffic from the source port out on the mirrored destination port.
Learn more with this free online training course on the Meraki Learning Hub:
Port mirror egress modes
Meraki switches offer two egress modes for port mirrors:
- 
    True egress — Native VLAN traffic is untagged.
- 
    Tagged egress — Native VLAN is tagged.
The following switches support each mode:
- 
    True egress: MS22, MS42, MS120, MS220, MS320, MS350, MS390, MS410, Meraki managed and monitored Catalyst switches.
- 
    Tagged egress: MS225, MS250, MS420, MS425.
Capture filter types
You can filter a capture for a specific client's IP address or for a specific type of traffic. Filter either before or after the capture. The filters used for each differ.
- 
    Capture filter — Limits the type of data the switch captures and saves to the file. This filter is used less frequently. Its syntax differs from display filters. Refer to the Wireshark database for capture filter syntax.
- 
    Display filter — The more common filter type. It does not reduce the traffic captured, which eliminates the possibility of applying an incorrect filter and missing the traffic required to troubleshoot. Apply this filter inside Wireshark when viewing the completed capture. Refer to the Wireshark database for display filter syntax.
Understanding a rolling capture
A rolling capture automatically saves the output to files at set intervals and breaks up a large capture into multiple smaller files. This helps when running a long-term capture for troubleshooting intermittent issues such as choppy audio on VOIP.
For some issues, you may need to run port mirrors or span port captures for long periods until the issue occurs. The goal is to run a capture and stop it once the issue surfaces. A capture run for a long duration—6 hours, for example—produces a .pcap file too large for your computer to open, since captures larger than 100 MB become difficult to open on some computers. Configure the capture with multiple options to make this easier.
Understanding the ring buffer
Set ring buffers to ensure you do not fill up all the disk space on your device. The ring buffer starts overwriting the oldest file based on how many files you specify. You do not have to use it, but it helps ensure you do not fill up your HDD.
Prerequisites
- Two ports on the same MS switch or within the same switch stack—one source port and one destination port.
- A workstation to connect to the destination port, with DHCP enabled on the host.
- Wireshark installed on the workstation. Refer to Wireshark's download page to download Wireshark, then follow the prompts.
The MS switch supports utilizing aggregate ports as a source port only (not as a destination port). The uplink port cannot be used as the destination and is not listed as an option.
Step-by-step instructions
Enable port mirroring on your switch
Mirror one or more ports on an MS switch:
- 
    In the Meraki dashboard, navigate to Switch > Monitor > Switch port s.
- 
    Select one or more ports to be mirrored. You can mirror multiple source ports to a single destination port.
- 
    Select Mirror.
- 
    Specify the destination mirror port, which captures traffic on the source ports. Both ports must be on the same switch or within the same switch stack. You can have multiple source ports but only a single destination port.
- 
    Select Create port mirror.
Connect a workstation to your destination port
- 
    Physically connect a workstation to your destination port.
- 
    Make sure DHCP is enabled on the host.
- 
    Confirm the host receives a 169.254.X.X IP address.
Clients connected to a destination port of a port mirror do not have network connectivity, as the destination port does not serve clients.
Take a packet capture with Wireshark
Wireshark displays the packets a device sees. Packets contain the data transmitted between computers, and viewing this information often aids in diagnosing network issues. A hardwired device may not see all packets transmitted on a network; it may only see broadcast packets and packets addressed to itself due to the functionality of modern networking equipment.
- 
    Open Wireshark.
- 
    Select Capture Options.
- 
    Uncheck Enable promiscuous mode on all interfaces, check the Promiscuous option for your capture interface, and select the interface.
- 
    Select Start. A new window shows the packets the device picks up.
- 
    Select Stop once you obtain the desired packets.
- 
    Save the capture from the File menu with a distinct name.
Learn more with this free online training course on the Meraki Learning Hub:
Take a rolling capture
- 
    Open Wireshark.
- 
    Select Capture Options.
- 
    Uncheck Enable promiscuous mode on all interfaces, check the Promiscuous option for your capture interface, and select the interface.
- 
    In the Output tab, select Browse.
- 
    Enter a filename in the Save As field, select a folder to save captures to, and select Save.
6. Select Create a new file automatically after… and Use a ring buffer with x files. This creates a maximum of x files, with each file set to the size or timeframe configured. For example, creating a new file automatically after 32 megabytes with a ring buffer of 128 files provides 4 gigabytes of rolling captures.
- 
    Select Start. A new window shows the packets the device picks up.
Apply a display filter
- 
    Open your packet capture.
- 
    Select the filter box.
- 
    Input the filter string as provided by your support engineer, then select Apply.
- 
    To save the filtered data, go to File > Export Specified Packets.
- 
    Make sure the Displayed radio button is checked and the file has a unique filename, then select Save.
Verification
- After creating the port mirror, the workstation connected to the destination port receives a 169.254.X.X IP address, confirming the connection.
- After starting the capture in Wireshark, the capture window displays the packets the device picks up, confirming that mirrored traffic reaches the workstation.
Troubleshooting
- A workstation connected to a destination port does not have network connectivity, because the destination port does not serve clients. This is expected behavior.
- If a capture runs for a long duration, the resulting .pcap file may exceed 100 MB and become difficult to open. Use a rolling capture with a ring buffer to break the output into smaller files and avoid filling your disk.
