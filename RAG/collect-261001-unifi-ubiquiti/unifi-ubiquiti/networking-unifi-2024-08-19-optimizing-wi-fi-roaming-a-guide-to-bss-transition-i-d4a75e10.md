---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/networking-unifi-2024-08-19-optimizing-wi-fi-roaming-a-guide-to-bss-transition-i-d4a75e10
title: "networking-unifi-2024-08-19-optimizing-wi-fi-roaming-a-guide-to-bss-transition-i-d4a75e10"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "latency"]
source: docs/RAG/collect-261001-unifi-ubiquiti/networking-unifi-2024-08-19-optimizing-wi-fi-roaming-a-guide-to-bss-transition-i-d4a75e10.md
source_anchor: ""
source_lines: [1, 44]
sha256: 9d880111a10257316e368eb304a22d7a269c0ee629b119ec56af64414e3187a3
---

# networking-unifi-2024-08-19-optimizing-wi-fi-roaming-a-guide-to-bss-transition-i-d4a75e10

In today’s world, seamless Wi-Fi connectivity is essential, especially in environments like offices, campuses, or large homes where users frequently move between different areas. Maintaining a strong and stable connection while transitioning between access points can be challenging, but BSS Transition (Basic Service Set Transition) is a feature designed to make this process smoother. In this article, we’ll explore what BSS Transition is, how it works in UniFi networks, and how to configure it for optimal Wi-Fi roaming performance.
Table of Contents
What is BSS Transition?
BSS Transition, sometimes referred to as Fast BSS Transition, is a feature that facilitates faster and more efficient Wi-Fi roaming. When a Wi-Fi device moves from the coverage area of one access point (AP) to another within the same network, it needs to reassociate with the new AP. This process can introduce delays, which may result in brief connectivity interruptions, especially for latency-sensitive applications like VoIP calls, video streaming, or online gaming.
BSS Transition aims to reduce the time it takes for a device to roam between APs by allowing the device to pre-authenticate with the next AP before leaving the coverage area of the current one. This pre-authentication process significantly reduces the time required for re-association, providing a more seamless user experience.
How Does BSS Transition Work?
In a typical Wi-Fi network, when a device moves out of the range of one AP, it needs to find and authenticate with a new AP before continuing data transmission. This can involve several steps, including scanning for available networks, authenticating with the new AP, and re-establishing secure connections.
BSS Transition simplifies and accelerates this process by allowing devices to negotiate the handoff between APs more efficiently. The device and the network work together to determine the best time and method to switch APs, ensuring that the transition happens as smoothly as possible.
Here’s a simplified breakdown of how BSS Transition works:
- Pre-authentication: While still connected to the current AP, the device identifies potential new APs and starts the authentication process in the background.
- Roaming Decision: The device and the network collaborate to decide when to switch to the new AP. This decision is typically based on signal strength, network load, and other factors.
- Seamless Handoff: Once the device moves into the new AP’s range, the connection is handed off quickly, with minimal disruption to ongoing activities like video calls or data transfers.
Benefits of BSS Transition in UniFi Networks
- Seamless Roaming: BSS Transition ensures that devices can move between APs with minimal disruption, providing a smooth experience for users who are on the move, such as during a video call or when streaming media.
- Reduced Latency: By pre-authenticating with the next AP, BSS Transition reduces the time required to reconnect, minimizing latency and improving the overall responsiveness of the network.
- Improved User Experience: In environments where users frequently move between coverage zones, BSS Transition helps maintain stable connections, reducing the chances of dropped connections or poor performance.
- Enhanced Network Efficiency: BSS Transition optimizes the use of available APs by ensuring that devices connect to the most suitable AP based on their location and network conditions, leading to better load balancing and overall network performance.
How to Enable BSS Transition in UniFi
Enabling BSS Transition in a UniFi network is a straightforward process that can be done through the UniFi Controller interface. Here’s how you can configure it:
- Access the UniFi Controller:
- Open your UniFi Controller interface through your browser.
- Log in with your credentials.
- Navigate to Wi-Fi Networks:
- From the main dashboard, go to the Settings section.
- Click on Wi-Fi Networks to view and manage your wireless networks.
- Edit the Wi-Fi Network:
- Select the Wi-Fi network where you want to enable BSS Transition.
- Click on Edit to modify the network settings.
- Enable BSS Transition:
- Scroll down to the Advanced Settings section.
- Look for the BSS Transition or Fast Roaming option (depending on your UniFi version).
- Enable this option to allow devices to take advantage of faster and more efficient roaming.
- Apply Changes:
- Once you’ve enabled BSS Transition, click Apply Changes to save and implement the new configuration.
- Monitor Performance:
- After enabling BSS Transition, monitor your network’s performance using the UniFi Controller’s Insights and Statistics features. Pay attention to how devices move between APs and ensure that roaming is smooth and efficient.
Best Practices for BSS Transition in UniFi Networks
- Optimize AP Placement: Ensure that your UniFi access points are placed strategically to provide overlapping coverage. This overlap allows devices to seamlessly transition between APs without experiencing a significant drop in signal strength.
- Test Roaming Performance: After enabling BSS Transition, test the roaming performance in your environment. Move around the space with a connected device, particularly while streaming video or making a VoIP call, to see how well the network handles transitions between APs.
- Consider Device Compatibility: While most modern devices support BSS Transition, some older devices may not fully benefit from the feature. In environments with mixed device capabilities, monitor how different devices respond to the change and adjust settings if necessary.
- Keep Firmware Updated: Ubiquiti frequently releases firmware updates that can improve BSS Transition performance and overall network stability. Ensure your UniFi devices are running the latest firmware to take advantage of these improvements.
- Combine with Other Features: BSS Transition works best when combined with other network optimization features, such as band steering and load balancing. These features help ensure that devices not only roam smoothly but also connect to the most appropriate band and AP.
Conclusion
BSS Transition is a valuable feature in UniFi networks that enhances Wi-Fi roaming by reducing the time and effort required for devices to switch between access points. Whether you’re managing a large office, a campus, or a multi-story home, enabling BSS Transition can significantly improve the user experience, especially for those who rely on stable and continuous connectivity. By following the steps outlined in this guide and adhering to best practices, you can optimize your UniFi network for seamless roaming and enhanced performance across all your devices.
