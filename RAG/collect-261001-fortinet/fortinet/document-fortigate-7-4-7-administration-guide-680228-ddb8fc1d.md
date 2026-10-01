---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-7-administration-guide-680228-ddb8fc1d
title: "diagnose sniffer packet <interface_name> <'filter'> <verbose> <count> <tsformat>"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-7-administration-guide-680228-ddb8fc1d.md
source_anchor: ""
source_lines: [1, 56]
sha256: 247703ecb810c63016707d7cbd71de61690693fb8a451e73e3cddc03e67a7fba
---

# diagnose sniffer packet <interface_name> <'filter'> <verbose> <count> <tsformat>

Performing a sniffer trace or packet capture
Performing a sniffer trace or packet capture
When you troubleshoot networks and routing in particular, it helps to look inside the headers of packets to determine if they are traveling the route that you expect them to take. Packet sniffing is also known as network tap, packet capture, or logic analyzing.
For more information on controlling GUI packet captures in the CLI, see Using the packet capture tool.
|  | For FortiGates with NP2, NP4, or NP6 interfaces that are offloading traffic, disable offloading on these interfaces before you perform a trace or it will change the sniffer trace. | 
Sniffing packets
To perform a sniffer trace in the CLI:
Before you start sniffing packets, you should prepare to capture the output to a file. A large amount of data may scroll by and you will not be able to see it without saving it first. One method is to use a terminal program like PuTTY to connect to the FortiGate CLI. Once the packet sniffing count is reached, you can end the session and analyze the output in the file.
The general form of the internal FortiOS packet sniffer command is:
# diagnose sniffer packet <interface_name> <'filter'> <verbose> <count> <tsformat>
To stop the sniffer, type CTRL+C.
| <interface_name> | The name of the interface to sniff, such as port1 orinternal . This can also beany to sniff all interfaces. | 
| <'filter'> | What to look for in the information the sniffer reads. none indicates no filtering, and all packets are displayed as the other arguments indicate. The filter must be inside single quotes ('). | 
| <verbose> | The level of verbosity as one of:  | 
| <count> | The number of packets the sniffer reads before stopping. If you don't put a number here, the sniffer will run until you stop it with <CTRL+C >. | 
| <tsformat> | The timestamp format.  | 
Simple sniffing example:
 # diagnose sniffer packet port1 none 1 3. 
This displays the next three packets on the port1 interface using no filtering, and verbose level 1. At this verbosity level, you can see the source IP and port, the destination IP and port, action (such as ack), and sequence numbers.
In the output below, port 443 indicates these are HTTPS packets and that 172.20.120.17 is both sending and receiving traffic.
Head_Office_620b # diagnose sniffer packet port1 none 1 3
interfaces=[port1]
filters=[none]
0.545306 172.20.120.17.52989 -> 172.20.120.141.443: psh 3177924955 ack 1854307757 
0.545963 172.20.120.141.443 -> 172.20.120.17.52989: psh 1854307757 ack 3177925808 
0.562409 172.20.120.17.52988 -> 172.20.120.141.443: psh 4225311614 ack 3314279933 
Using packet capture in a firewall policy
FortiGate can capture packets matching a firewall policy. You can enable capture-packet  in the firewall policy.
To use packet capture, the FortiGate must have a disk and logging must be enabled in the firewall policy.
For information about using the packet capture tool in the GUI, see Using the packet capture tool.
To enable packet capture in a policy in the GUI:
- 
                                                    Go to Policy & Objects > Firewall Policy and click Create New.
- 
                                                    Enter a name for the policy and configure the required settings.
- 
                                                    Enable Log Allowed Traffic and select Security Events or All Sessions.
- 
                                                    Enable Capture Packets.
- 
                                                    Click OK.
To enable packet capture in a policy in the CLI:
config firewall policy
    edit <id>
        set action accept
        set logtraffic {all | utm}
        set capture-packet enable
    next
end
                                            To view the packet capture:
- 
                                                    Go to Log & Report > Forward Traffic and select the log that matches the firewall policy.
- 
                                                    Select Details > Archived Data and click on the download button.
- 
                                                    Open the downloaded PCAP file in a packet analyzer tool, such as Wireshark.
