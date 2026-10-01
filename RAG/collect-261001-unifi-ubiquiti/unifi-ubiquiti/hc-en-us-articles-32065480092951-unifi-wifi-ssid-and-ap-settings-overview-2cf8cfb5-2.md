---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-32065480092951-unifi-wifi-ssid-and-ap-settings-overview-2cf8cfb5-2
title: "hc-en-us-articles-32065480092951-unifi-wifi-ssid-and-ap-settings-overview-2cf8cfb5"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["exploit"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-32065480092951-unifi-wifi-ssid-and-ap-settings-overview-2cf8cfb5.md
source_anchor: ""
source_lines: [63, 103]
sha256: aa5fc17d26bf7c4e2ac2292573cadb88a6612ec77961c38be8cb31c14efbfcb5
---

# hc-en-us-articles-32065480092951-unifi-wifi-ssid-and-ap-settings-overview-2cf8cfb5

802.11 DTIM Period
Controls how frequently the access point sends buffered multicast and broadcast traffic, and can be configured separately for each WiFi band or left on Auto. Raising the DTIM interval can improve fast roaming performance in some cases, but it may negatively affect airtime efficiency, battery life, and connectivity for IoT devices.
This is an advanced setting and should only be adjusted if you're targeting specific roaming behavior or troubleshooting timing-related issues. Our defaults of 1 and 3 for 2.4Ghz and 5/6Ghz respectively have been tuned to balance connectivity and device performance.
Minimum Data Rate Control
Defines the lowest WiFi data rate that clients are allowed to use when connecting to the SSID. Increasing the minimum rate helps reduce airtime usage and encourages clients to roam sooner, improving overall efficiency—especially in large or high-density networks. It can also block legacy or low-speed IoT devices that rely on outdated data rates. Before adjusting this setting, consider the types of devices connecting to your network, as even high-performance clients may fall back to lower rates when signal strength is weak.
Security Settings
These options control how clients authenticate and how secure the WiFi connection is.
Security Protocol (WPA2/WPA3)
Defines the encryption and authentication standard for the network.
Use case:
- Use WPA2/WPA3 for most environments — this provides a good compromise between modern security and legacy device compatibility. It allows newer devices to use WPA3 while falling back to WPA2 for those that don’t support it.
- Use WPA3 only when enabling MLO.
- Use Enterprise for RADIUS-based authentication.
- WPA2 only may be necessary for legacy IoT devices—or when using features like PPSK, which currently require WPA2 compatibility.
- For guest networks, use OWE with an Open network - this gives modern clients the security of WPA3 without a password. Do note that all 6 GHz networks will be secured with OWE or WPA3, as required by standards - the OWE toggle in UniFi Network will enable it for 2.4 and 5 GHz on your SSID as well.
Protected Management Frames (PMF)
Adds encryption to WiFi management traffic—such as association and disassociation messages—to prevent spoofing and disassociation attacks. PMF is a required component of WPA3 and Fast Roaming, and can be configured with three options:
- Optional (default): Enables PMF when supported by the client, while maintaining compatibility with older devices that don’t support it.
- Required: Enforces PMF for all clients, providing the highest level of security. This is required for WPA3-only networks and OWE networks.
- Disabled: Turns off PMF entirely. This is only recommended when creating a separate WPA2 SSID for legacy or incompatible devices.
PMF is essential for securing modern networks, but disabling it should only be considered in cases where specific clients can’t connect otherwise.
Group Rekey Interval
Defines how often encryption keys are refreshed for broadcast traffic. We strongly recommend keeping the default setting.
MAC Address Filter
Restricts access to the network by allowing only client devices with specific MAC addresses to connect. This provides a basic layer of access control through allowlists managed within the UniFi interface. While useful for limiting which devices can join a network, it’s generally better suited for simple control scenarios with clients that do not rotate or randomize their MAC addresses. For more flexible or secure segmentation—such as assigning VLANs per device—PPSK is usually a more robust alternative.
RADIUS MAC Authentication
Authenticates client devices by verifying their MAC addresses against an external RADIUS server before allowing them to join the network. This method is commonly used in enterprise environments where centralized, policy-driven control over device access is required.
Unlike MAC Address Filter, which relies on a locally managed allowlist, RADIUS MAC Authentication offers centralized enforcement, integration with identity systems, and the ability to dynamically assign VLANs or policies based on MAC identity. Learn more here
SAE Anti-clogging
This is a required feature for WPA3 authentication by preventing denial-of-service (DoS) attacks that attempt to overwhelm the AP with handshake requests. This feature is rarely needed in typical deployments and is intended for high-security environments or cases where targeted attacks are a concern. It can safely be left disabled in most networks unless there is a specific threat model that requires it. We highly recommend not to adjust the default value of it.
SAE Sync Time
Configures the timeout window for WPA3's handshake retries, which determines how long the system waits for a response before dropping the attempt. This setting is designed to help mitigate denial-of-service (DoS) attacks that exploit handshake timing. Like SAE Anti-clogging, it is on by default and should only be adjusted in high-security environments or when specifically addressing targeted attack scenarios.
WiFi Blackout Scheduler
Allows you to automatically disable an SSID during specific times of the day or week. This is particularly useful for limiting access to guest or IoT networks outside of business hours, reducing unnecessary network activity, or enforcing screen time rules in home environments. Each SSID can have its own schedule, offering flexible control across different network types
Individual AP settings
Minimum RSSI
Tells the AP to disconnect clients based on signal strength (measured in dBm). This is helpful when attempting to keep client data rates up by enforcing strict limits to AP cell size (signal range) and control for sticky clients. However, this can have implications for general roaming if improperly tuned; some devices may refuse to connect to an AP from which they have been kicked multiple times. This is generally not recommended unless running a high density deployment.
Interference Blocker
An extension of Minimum RSSI which treats all connections below the minimum signal strength as noise - the AP will not waste airtime attempting to complete the association process with these clients. Recommended for high density deployments.
Roaming Assistant
New in Network 9.2, this uses BSS transition frames to inform clients that they will be dissociated from the AP once they drop below a certain signal strength. Unlike the “hard kick” of Minimum RSSI, this “soft kick” is tolerated much better by modern clients and is generally recommended to be set to a value below -70 dBm.
