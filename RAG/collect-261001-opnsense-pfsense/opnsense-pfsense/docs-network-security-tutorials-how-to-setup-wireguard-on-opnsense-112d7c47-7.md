---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/docs-network-security-tutorials-how-to-setup-wireguard-on-opnsense-112d7c47-7
title: "docs-network-security-tutorials-how-to-setup-wireguard-on-opnsense-112d7c47"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["latency", "parameters", "throughput"]
source: docs/RAG/collect-261001-opnsense-pfsense/docs-network-security-tutorials-how-to-setup-wireguard-on-opnsense-112d7c47.md
source_anchor: ""
source_lines: [576, 674]
sha256: 4e0f5a7727a877f3ec87a22440cfc9ec59cbb9e5d003a43f01bae853507d9de5
---

# docs-network-security-tutorials-how-to-setup-wireguard-on-opnsense-112d7c47

Redundant routing ensures that VPN connectivity stays active, even during outages or server failures, without requiring manual intervention. This is especially valuable for remote workers, branch offices, or systems that require constant secure access.
3. Implement Health Monitoring to Trigger Failover: Even with redundant servers and routing, you still need a way to detect when the primary VPN server is down, and then automatically switch to the backup server. That’s where health monitoring comes in.
It constantly checks if your VPN connection is working. If it detects a failure (for example, no response from the main server), it can run a command or trigger a script to switch to the backup.
OPNsense includes a Gateway Monitoring system that checks the connection status of each VPN server.
- 
Go to System → Gateways .
- 
Edit your WireGuard gateway.
- 
Enable Monitor IP set it to something like 1.1.1.1 or 8.8.8.8 (these are stable DNS IPs to test reachability).
- 
Set conditions for failover. Choose when to trigger failover: Packet Loss (% of failed checks), Latency or Down.
- 
Enable email or system notifications if desired.
- 
Save and apply the settings.
When the primary gateway becomes unreachable, OPNsense will automatically switch to the Tier 2 (backup) gateway if you configured Gateway Groups earlier.
How can I Monitor WireGuard VPN Connections Using OPNsense's Diagnostic Tools?
Monitoring WireGuard VPN connections in OPNsense is crucial for ensuring stability, detecting issues, and maintaining continuous secure access. OPNsense provides several built-in diagnostic tools that allow you to view logs, observe real-time connection status, and analyze network traffic through packet capture. These tools help identify handshake failures, connectivity drops, or routing problems.
1. View WireGuard Logs Within OPNsense: To begin monitoring, you can view WireGuard logs directly within the OPNsense web interface. Follow these steps.
- 
Navigate to System → Log Files → General .
- 
In the log viewer, look for entries related to wireguard, wg-quick, or interface status updates.
- 
Depending on the WireGuard plugin version, there may also be a dedicated WireGuard log section under VPN → WireGuard → Log File .
These logs include valuable information such as successful handshakes, interface status, peer connectivity, and error messages. Reviewing these logs helps you understand when connections were established or if failures occurred.
These log types that are listed below are especially useful for troubleshooting.
- 
Handshake logs show if the connection between the client and server was successfully established.
- 
Interface up/down logs indicate whether the VPN interface was brought online or offline.
- 
Peer activity logs help confirm if data is being exchanged and if peers are staying connected.
By analyzing these logs, you can detect problems such as incorrect keys, failed tunnels, or issues caused by MTU mismatch.
2. Real-Time Monitoring of WireGuard Status*: For live monitoring, OPNsense provides a status page for WireGuard, follow these steps
- 
Go to VPN → WireGuard → Status .
- 
You can see each peer, their latest handshake timestamp, and data transfer statistics.
If the "Last Handshake" is recent and data transfer is active, your VPN connection is working as expected. However, if no handshake is detected for a prolonged time, it may indicate that the server is unreachable or that the peer is misconfigured.
3. Using Packet Capture for Deep Troubleshooting: When log files and status pages are not enough, OPNsense's built-in Packet Capture tool allows you to analyze the actual network traffic on the WireGuard interface.
To use the packet capture tool, follow these steps.
- 
Go to Interfaces → Diagnostics → Packet Capture. 
- 
Select the WireGuard interface (e.g., wg0) from the dropdown menu.
- 
Set a custom capture filter such as udp port 51820, which is the default port used by WireGuard.
- 
Start the capture and download the resulting .pcap file for further analysis using Wireshark or a similar tool.
Packet capture helps you identify the followings.
- 
Whether WireGuard packets are being sent from your device.
- 
Whether the remote server is responding.
- 
Any signs of packet loss, corruption, or duplication.
- 
Potential issues with NAT or firewall rules blocking return traffic.
This tool is essential when you need to troubleshoot advanced issues like asymmetric routing or intermittent connectivity.
To effectively monitor WireGuard VPN connections in OPNsense, you can combine system log analysis, real-time status monitoring, and packet-level inspection. These diagnostic tools provide deep insight into the health of your VPN tunnels and enable you to take action quickly when something goes wrong. This is especially important in critical environments where secure and continuous VPN connectivity is required
How do I Test the Performance and Speed of My WireGuard VPN Setup?
Testing the performance of your WireGuard VPN setup is crucial to ensure that the connection is stable, fast, and suitable for your needs. While WireGuard is known for its high-speed performance, various factors like server load, routing, bandwidth limits, and encryption overhead can impact real-world speeds. To get accurate measurements, it’s best to use dedicated tools and follow a structured approach.
One of the most effective tools to measure the speed between two endpoints is iperf3. It can test both upload and download throughput between your WireGuard server and client. You can use the tools shown below for WireGuard performance test.
- 
speedtest-cli : For testing internet speed from the client over the VPN tunnel.
- 
ping : To measure latency and packet loss.
- 
iperf3 : For controlled bandwidth testing between two hosts (ideal for VPN testing).
Among these, iperf3 is the most reliable for VPN testing since it eliminates external internet factors and focuses solely on the VPN tunnel.
Here’s how you can use iperf3 to test performance between your WireGuard server and client.
1. Install iperf3 on Both Systems: On Debian/Ubuntu-based systems, run the following command.
sudo apt install iperf3
On FreeBSD/OPNsense, use the following command.
pkg install iperf3
2. Start the iperf3 Server on One End: On your WireGuard server, run the following command.
iperf3 -s
This puts the server in listening mode, ready to receive traffic.
3. Run the iperf3 Client on the Other End: On the client, run the following command.
iperf3 -c <WireGuard_Server_IP>
This command starts the test and sends traffic through the VPN tunnel to the server. Also, optional parameters are listed below.
-t 30: Run the test for 30 seconds.
-R: Reverse mode to test upload from server to client.
Example command can be seen below. This will help you measure download speeds (from server to client).
iperf3 -c 10.0.0.1 -t 30 -R
4. Interpreting the Results: Once the test completes, you'll see the following details.
- 
Bandwidth (in Mbits/sec): This tells you the actual data throughput.
- 
Retransmissions (TCP only): In TCP mode, the output includes a “Retr” column, which shows how many packets had to be resent due to network issues.
- 
Jitter and Latency (if tested with UDP): Useful for VoIP or gaming.
- 
Packet loss: Any lost packets indicate instability or congestion.
Compare the results with your internet speed to determine if WireGuard introduces any significant bottlenecks.
Testing your WireGuard VPN’s performance is best done with iperf3, which allows controlled speed tests directly through the tunnel. By installing the tool on both ends and running client-server tests, you can accurately assess upload/download performance and identify any limitations in your setup. Regular performance testing ensures that your VPN delivers the speed and reliability needed for secure, high-throughput tasks.
Conclusion
