---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-8-administration-guide-462154-a5c61c7b-1
title: "diagnose test application forticron 14"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-8-administration-guide-462154-a5c61c7b.md
source_anchor: ""
source_lines: [1, 105]
sha256: 574302173cd93f4b5f04a6844a967e2ef5710f7554e7f4fcb12592eca98b98fa
---

# diagnose test application forticron 14

Using the packet capture tool
Using the packet capture tool
Administrators can use the packet capture tool to select a packet and view its header and payload information in real-time. Once completed, packets can be filtered by various fields or through the search bar. The capture can be saved as a PCAP file that you can use with a third-party application, such as Wireshark, for further analysis.
Packet capture criteria can be stored for the re-initiation of packet captures multiple times using the same parameters, such as interface, filters, and so on.
Packet capture criteria can be created and stored in order to re-initiate packet captures in the GUI with the same parameters. Capture cards in the Network > Diagnostics page are sorted in alphabetical order of the configured name and colored depending on state:
- 
                                                    Green: The packet capture is running.
- 
                                                    Gray: The packet capture has not started yet, has completed, or the capture files have been deleted.
When creating the packet capture:
- 
                                                    The Name field must be a unique name for the packet capture criteria being configured.
- 
                                                    Enabling Include non-IP packets allows non-IP address packets to be captured when enabled. Supported non-IP address packet types include ARP, RARP, LLC, LLDP, VLAN, and LACPDU. When the packet capture is complete, non-IP address packets will include header information, however, unsupported types will display as Unknown.
- 
                                                    After configuring the packet capture criteria, you can choose to Start capture, Save settings for later, or Close. Starting a packet capture or saving the configured settings will both store criteria for future use.
To use the packet capture tool in the GUI:
- 
                                                    Go to Network > Diagnostics and select the Packet Capture tab.
- 
                                                    Optionally, select an Interface (any is the default).
- 
                                                    Optionally, enable Filters and select a Filtering syntax: 
  - 
                                                            Basic: enter criteria for the Host, Port, and Protocol number.
  - 
                                                            Advanced: enter a string, such as src host 172.16.200.254 and dst host 172.16.200.1 and dst port 443.
- 
                                                            
- 
                                                    Click Start capture. The capture is visible in real-time.
- 
                                                    While the capture is running, select a packet, then click the Headers or Packet data tabs to view more information. When the packet capture is running, disable Auto-scroll to stop automatic scrolling behavior when new packets arrive.
- 
                                                    When the capture is finished, click Save as pcap. The PCAP file is automatically downloaded.
- 
                                                    Optionally, use the Search bar or the column headers to filter the results further.
Multiple packet captures
Multiple packet captures can be run simultaneously for when many packet captures are needed for one situation. For example, ingress and egress interfaces can be captured at the same time to compare traffic or the physical interface and VPN interface can be captured using different filters to see if packets are leaving the VPN.
The packet capture dialog can be docked and minimized to run in the background. The minimized dialog aligns with other CLI terminals that are minimized.
|  | How many packet captures and the number of packets that can be captured depend on the device model. Whether the device model has disk storage affects when packet captures are deleted. Without disk storage, packet captures are deleted 24 hours after completion or immediately after reboot. With disk storage, packet captures are deleted after 7 days. To find the limit on the number of packet captures supported for a specific device model, use the Maximum Values Table, and search for the object firewall.on-demand-sniffer . | 
To run multiple packet captures at the same time:
- 
                                                    Go to Network > Diagnostics.
- 
                                                    Configure the first packet capture: 
  - 
                                                            Click New packet capture.
  - 
                                                            Select the Interface and configure other settings as needed.
  - 
                                                            Click Start capture. The first packet capture begins.
- 
                                                            
- 
                                                    Minimize the packet capture. The packet capture continues to run.
- 
                                                    Configure the second packet capture: 
  - 
                                                            Click New packet capture.
  - 
                                                            Select the Interface and configure other settings as needed.
  - 
                                                            Click Start capture. The second packet capture begins.
- 
                                                            
- 
                                                    When the captures are complete, expand the dialog and select Save as pcap for each packet capture.
Persistent packet captures
If the browser is closed or refreshed, users can return at a later time to view, stop, restart, and download the packet capture.
|  | The limit of simultaneous packet captures and the number of packets per capture that can be saved is dependent on the FortiGate model and capabilities. If this limit is reached, new packet captures cannot be created. Go to Network > Diagnostics to view the limit values. | 
To interact with a packet capture in the GUI:
- 
                                                    Go to Network > Diagnostics.
- 
                                                    Click New packet capture.
- 
                                                    Configure the packet capture fields.
- 
                                                    Click Start Capture.
- 
                                                    Stop the packet capture: 
  - 
                                                            If you have closed the packet capture, select the packet capture and click Stop.
  - 
                                                            If you are viewing the packet capture, click Stop capture.
- 
                                                            
- 
                                                    Download the packet capture: 
  - 
                                                            If you have closed the packet capture, select the packet capture and click Download.
  - 
                                                            If you are viewing the packet capture, click Save as .pcap.
- 
                                                            
- 
                                                    Restart the packet capture: 
  - 
                                                            If you have closed the packet capture, select the packet capture and click Restart.
  - 
                                                            If you are viewing the packet capture, click Restart capture.
 You will have the option to save the previous packet capture or discard it.
- 
                                                            
