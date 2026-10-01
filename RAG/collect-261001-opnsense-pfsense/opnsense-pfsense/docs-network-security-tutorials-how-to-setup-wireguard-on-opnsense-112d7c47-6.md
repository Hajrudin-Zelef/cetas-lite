---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/docs-network-security-tutorials-how-to-setup-wireguard-on-opnsense-112d7c47-6
title: "docs-network-security-tutorials-how-to-setup-wireguard-on-opnsense-112d7c47"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["SpaceX"]
dates: []
keywords: ["benchmark", "latency", "lean", "memory", "throughput"]
source: docs/RAG/collect-261001-opnsense-pfsense/docs-network-security-tutorials-how-to-setup-wireguard-on-opnsense-112d7c47.md
source_anchor: ""
source_lines: [500, 575]
sha256: 3d54ae4694fb85272ce501d3e265f9850cd6427efa09b6743d5cd675c5c018ac
---

# docs-network-security-tutorials-how-to-setup-wireguard-on-opnsense-112d7c47

Regularly Audit Active and Inactive Peers: To maintain a clean and secure VPN configuration, it's important to routinely audit peer activity. You can do this by checking the “Last Handshake” time for each peer under VPN → WireGuard → Status in the OPNsense dashboard.If a peer has shown no activity for 30 days or more, and this aligns with your internal policy, consider removing or disabling that peer to reduce unnecessary exposure and keep the configuration lean.
- 
Use DNS Servers Explicitly to Avoid Leaks: Without proper DNS configuration, client devices may bypass the VPN tunnel when resolving DNS queries, which can lead to privacy breaches and data leaks. To prevent this, make sure to explicitly define DNS resolvers within the WireGuard configuration files. Use trusted and privacy-respecting DNS servers such as 10.0.0.1, 1.1.1.1 , or9.9.9.9 .In OPNsense, you can do this by navigating to VPN → WireGuard → Peers , and entering the preferred DNS servers in the DNS Servers field.
- 
Protect and Limit Access to WireGuard Admin Interface: The WireGuard management interface, especially within platforms like OPNsense Web UI, is a critical control point for your VPN infrastructure. It is essential to restrict access only to trusted administrators and implement multi-factor authentication (MFA) to prevent unauthorized changes or breaches. Recommended Security Measures are visible below.
- 
Limit access to the admin interface by specific IP addresses or internal networks to avoid exposure over the internet.
- 
Assign admin roles with only the minimum necessary privileges, following the principle of least privilege.
- 
Enable Two-Factor Authentication (2FA) using TOTP (Time-based One-Time Password) for added security during login. noteNever expose the admin panel to the public internet without strict firewall rules and access restrictions in place.
How can I Customize MTU and Optimize WireGuard Performance?
Optimizing WireGuard performance ensures a fast, stable, and secure VPN experience, especially in environments with varying network conditions. Below are key practices and adjustments you can implement to maximize WireGuard efficiency and reliability.
- 
Adjust the MTU (Maximum Transmission Unit) for Your Network: The MTU defines the largest size of a packet that can be transmitted without fragmentation. Using an inappropriate MTU may result in packet loss, latency issues, or connectivity problems. Manually test your ideal MTU value by sending pings with different payload sizes. On a Linux machine, you can use the following command. ping -M do -s 1472 your_vpn_server_ip
Gradually decrease the -s value (starting from 1472) until you find the highest size that doesn’t return a fragmentation needed error. Once identified, set the MTU in your WireGuard configuration file.tipLower MTUs (like 1280–1400) work better on LTE, mobile, or PPPoE networks where smaller packets reduce the risk of fragmentation.
- 
Optimize UDP Buffer Sizes to Improve Throughput: WireGuard uses the UDP protocol, which can be sensitive to system buffer limitations under heavy load. Increase your system’s send and receive buffer sizes to avoid packet drops with the following commands. sysctl -w net.core.rmem_max=2500000
 sysctl -w net.core.wmem_max=2500000
To make this change persistent across reboots, add the values to your /etc/sysctl.conf file. net.core.rmem_max = 2500000
 net.core.wmem_max = 2500000
This is especially important for high-throughput environments, such as data centers or file transfers over VPN, where small default buffers can become bottlenecks.
- 
Allocate System Resources to WireGuard Efficiently: When WireGuard runs on resource-constrained systems like routers or virtual machines, performance may degrade due to CPU or memory bottlenecks. Give WireGuard processes higher priority using nice with the following command. nice -n -5 wg-quick up wg0
In OPNsense, regularly monitor system resource usage (CPU, memory, and I/O) via the dashboard to ensure WireGuard isn’t competing with other processes. Consider disabling unused services or increasing hardware resources if you're running WireGuard on virtual machines or low-end devices.
- 
Adapt Configuration for Unstable or High-Latency Connections: Connections over LTE, satellite (e.g., Starlink), or public Wi-Fi networks often suffer from high latency and interruptions. Enable the PersistentKeepalive option to maintain a stable connection across NAT devices or aggressive firewalls.[Peer]
 PersistentKeepalive = 25
Consider reducing MTU further (e.g., 1280) to accommodate the unpredictable nature of such connections.
- 
Monitor and Benchmark Performance Continuously: Measuring performance regularly helps detect bottlenecks, packet loss, or routing misconfigurations early. Use the network tools provided below.  iperf3 to measure throughput ping andtraceroute for latency and path analysis wg show to inspect peer stats and handshake statusIn OPNsense, navigate to VPN → WireGuard → Status to view real-time data transfer rates, handshake times, and peer activity. Use these logs to assess if optimizations are effective.
How do I Set Up a WireGuard Failover Cluster or Redundant VPN Setup?
Setting up a** redundant or failover WireGuard VPN** ensures your devices stay connected even if the main VPN server goes down. While WireGuard does not support automatic multi-server switching out of the box, you can create a reliable failover system using smart routing, monitoring tools, and some configuration tricks.
1. Configure Multiple Endpoints for Failover (Manually or Dynamically): In WireGuard, you can only define one Endpoint per [Peer] block. This means you can't add two servers directly into a single config. However, you have options.
- 
Option 1: Use Dynamic DNS to point to whichever server is active: Instead of hardcoding a server IP, use a Dynamic DNS (DDNS) hostname as shown below. [Peer]
 PublicKey = SERVER_PUBLIC_KEY
 Endpoint = vpn.example.com:51820
When your primary server goes offline, update the DNS record to point to the backup server. Your client will connect to the new server without changing its config.
- 
Option 2: Use Separate Config Files and Switch with Scripts: You can create two separate .conf files.Primary Configuration: wg0.conf[Interface]
 PrivateKey = your_private_key
 Address = 10.0.0.2/32
 [Peer]
 PublicKey = PRIMARY_PUBLIC_KEY
 Endpoint = server1.example.com:51820
 AllowedIPs = 0.0.0.0/0
Backup Configuration: wg1.conf[Interface]
 PrivateKey = your_private_key
 Address = 10.0.0.2/32
 [Peer]
 PublicKey = BACKUP_PUBLIC_KEY
 Endpoint = server2.example.com:51820
 AllowedIPs = 0.0.0.0/0
tipBoth configurations use the same client private key and IP address, but connect to different servers.
You can switch from the primary to the backup server manually or automate the process using scripts or monitoring tools. The command below deactivates the current VPN and activates the backup.
wg-quick down wg0 && wg-quick up wg1
You can integrate this command into a health-check script or use a tool like Monit or systemd to automate the failover.
2. Set Up Redundant Routing in OPNsense (or Linux): To ensure uninterrupted VPN access, you can configure redundant routing so that if your primary WireGuard server becomes unreachable, traffic is automatically routed through a backup server.
OPNsense provides built-in support for routing failover using Gateway Groups. Follow these steps.
- 
Navigate to System → Gateways → Group .
- 
Click **+**Add to create a new gateway group.
- 
Add both of your WireGuard VPN gateways. 
  - 
Set Tier 1 for the primary gateway.
  - 
Set Tier 2 for the backup gateway (used only if Tier 1 fails).
- 
- 
For the Trigger Level , choose when failover should occur:
  - Options include Packet Loss, High Latency, or Down.
- 
Save the group and assign it to the appropriate firewall rule (e.g., for LAN or WireGuard interface).
