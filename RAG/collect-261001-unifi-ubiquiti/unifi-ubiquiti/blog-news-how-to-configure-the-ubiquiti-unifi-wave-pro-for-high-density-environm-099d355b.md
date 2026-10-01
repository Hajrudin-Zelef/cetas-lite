---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-news-how-to-configure-the-ubiquiti-unifi-wave-pro-for-high-density-environm-099d355b
title: "blog-news-how-to-configure-the-ubiquiti-unifi-wave-pro-for-high-density-environm-099d355b"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "ethernet"]
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-news-how-to-configure-the-ubiquiti-unifi-wave-pro-for-high-density-environm-099d355b.md
source_anchor: ""
source_lines: [1, 46]
sha256: 6b1d190da325a546e0ea40ededa4686ab05dd57fdefdabfc54713f030c16d2c5
---

# blog-news-how-to-configure-the-ubiquiti-unifi-wave-pro-for-high-density-environm-099d355b

The Ubiquiti UniFi Wave Pro is a high-performance access point (AP) designed to provide reliable Wi-Fi coverage in high-density environments, such as large office spaces, schools, stadiums, and shopping centers. By delivering fast connectivity and advanced management features, the UniFi Wave Pro helps ensure seamless connectivity for a large number of users and devices. 
In this post, we’ll walk through how to configure the Ubiquiti UniFi Wave Pro for optimal performance in high-density environments.
Step 1: Initial Setup and Network Integration
1.    Physical Placement and Installation:
•    In high-density environments, access point placement is crucial. Place the Wave Pro APs in strategic areas to ensure maximum coverage with minimal overlap. Avoid placing APs too close to each other, as this can create interference and lower performance. Ubiquiti recommends a grid-like layout with spacing adjusted based on room size and wall density.
•    Mount the Wave Pro APs at ceiling level if possible, as this provides the best distribution of the Wi-Fi signal across the area.
2.    Power and Ethernet Connection:
•    Connect the AP to a PoE (Power over Ethernet) switch to provide both power and data transmission. The UniFi Wave Pro supports PoE+, so make sure your switch can supply sufficient power to all devices on the network.
3.    UniFi Network Application Integration:
•    Log in to the UniFi Network application (available as a web interface or a mobile app) to begin configuration. Connect the Wave Pro APs to the network, and they should appear in the UniFi Network application as adoptable devices.
Step 2: Configuring Basic Wi-Fi Settings
1.    SSID and Frequency Bands:
•    Configure your SSID (network name) to be simple and recognizable. For high-density environments, it’s often beneficial to use one primary SSID for all users, simplifying network management.
•    Enable both 2.4 GHz and 5 GHz bands. The 2.4 GHz band covers a larger area and penetrates walls better but is more prone to interference, while the 5 GHz band supports higher speeds and works well in high-density settings.
2.    Channel Width Settings:
•    Set the channel width to 20 MHz for 2.4 GHz and 40 MHz for 5 GHz. Narrower channels in high-density environments reduce interference and allow more APs to operate without overlapping channels.
•    To avoid channel overlap, use non-overlapping channels (such as channels 1, 6, and 11 in the 2.4 GHz band). The UniFi Network application can automatically manage channel selection, or you can manually adjust it for optimal performance.
Step 3: Optimize for High-Density Traffic
1.    Adjust Transmit Power:
•    In crowded spaces with multiple APs, setting a lower transmit power reduces interference and allows clients to switch between APs more efficiently.
•    Start with medium power settings for both 2.4 GHz and 5 GHz bands, then adjust downwards if needed. UniFi’s Auto-Optimize Network feature can also assist by automatically tuning power levels and channel usage.
2.    Enable Band Steering:
•    Band steering encourages dual-band clients to connect to the less congested 5 GHz band, improving load distribution. In high-density areas, this feature helps balance the number of users across both frequency bands.
• Enable Band Steering in the Wi-Fi settings of the UniFi Network application.
3.    Set Up Load Balancing:
•    Load balancing helps ensure that no single AP becomes overloaded by limiting the number of connections per AP. The UniFi Wave Pro can handle a large number of devices, but configuring load balancing improves performance and user experience.
•    To enable this feature, go to the Settings in the UniFi Network application, navigate to the Wireless Networks section, and activate Load Balancing.
Step 4: Security and Access Management
1.    Guest Network Configuration:
•    If your environment has frequent guests, consider creating a dedicated guest network with its own SSID. Enable Guest Control settings to limit bandwidth and access to certain network resources.
• Implement captive portals or vouchers for secure guest access. The UniFi Network application offers these features to help you control and track guest activity.
2.    Enable WPA3 Security:
•    For the main network, enable WPA3 encryption to provide the most secure connection for users. WPA3 is especially beneficial in high-density environments, protecting sensitive data from potential threats.
3.    Advanced Threat Management:
•    The UniFi Threat Management suite offers firewall rules, intrusion detection, and intrusion prevention to safeguard your network from external threats.
•    Enable and configure these options under Settings > Security in the UniFi Network application. For high-density environments, configure firewall rules that prioritize business-critical applications or services.
Step 5: Monitor and Fine-Tune Performance
1.    Monitor Real-Time Network Usage:
•    The UniFi Network application provides real-time insights into connected devices, data usage, and AP performance. Use this data to identify any bottlenecks or overloaded APs and make adjustments.
2.    Optimize with Analytics and Insights:
•    The application’s analytics can reveal patterns in usage, helping you anticipate network load times and plan for optimal performance.
• For instance, if certain times of day are particularly busy, you may consider scheduling AP reboots or adjusting channel configurations accordingly.
3.    Regular Firmware Updates:
•    Ensure all APs are running the latest firmware. Ubiquiti regularly releases updates that provide performance improvements, security patches, and new features.
Configuring the Ubiquiti UniFi Wave Pro for a high-density environment requires thoughtful planning and optimization of key settings, including channel usage, power management, and security. By following these steps, you’ll ensure your network is prepared to handle heavy traffic while maintaining smooth connectivity for all users. 
Proper configuration allows the Wave Pro to perform optimally, making it an excellent choice for high-demand networking environments.
