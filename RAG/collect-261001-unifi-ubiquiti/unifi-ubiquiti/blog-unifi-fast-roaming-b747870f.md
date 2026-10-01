---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-unifi-fast-roaming-b747870f
title: "blog-unifi-fast-roaming-b747870f"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-unifi-fast-roaming-b747870f.md
source_anchor: ""
source_lines: [1, 39]
sha256: 3c7f6e3b8529c1bb3a1f13cc2622a9e050deab317e2ec44979b658ec952b08c5
---

# blog-unifi-fast-roaming-b747870f

UniFi fast roaming
If you're working with larger networks UniFi fast roaming might come in handy. Having multiple APs, it's important that your devices can easily switch as they move around. This is what fast roaming takes care off.
It basically improves the user experience by reducing the time it takes for devices to switch between APs. In this post, I'll (try to) show you the concept of fast roaming, its benefits, how it works within the UniFi, and ofcourse how to set it up!
Lets get it!
UniFi Fast Roaming
Fast roaming is a standard that enables devices to roam quickly from one AP to another within the same network. It does this by reducing the time it takes to authenticate to a new AP. Especially useful in environments where users or devices move frequently like big offices or campuses.
Benefits:
- Improve user experience: Provides smoother voice and video calls, and uninterrupted data streaming as users move around.
- Increased efficiency: Less authentication time as devices switch APs.
How Fast Roaming Works
Fast roaming uses a feature called Fast Basic Service Set Transition (FT). It allows the network to establish a security handshake before the device moves to the next AP. This pre-authentication process speeds up the time it takes for the device to reconnect and thus minimizing disruption in connectivity.
Setting up fast roaming in UniFi
Here’s how to enable ast roaming within the UniFi Controller for a network that supports and benefits from this feature:
Step 1: Access the UniFi Controller
Make sure you have admin access to your UniFi Controller. Open it via a web browser by connecting to the network address of your UniFi Controller.
Step 2: Open wireless networks settings
- Navigate to settings: In the UniFi Controller dashboard, find the settings section—this might be a gear icon or a sidebar menu depending on your version.
- Wireless networks: Go to the ‘Wireless Networks’ section where your Wi-Fi networks (SSIDs) are managed.
Step 3: Edit or create a wireless network
- Select a network: Choose the Wi-Fi network you wish to enable fast roaming on, or create a new one if needed.
- Advanced options: In the network settings, look for ‘Advanced Options’ or similar, which is often a checkbox that reveals more settings.
Step 4: Enable fast roaming
- Find fast roaming settings: Within the advanced options, look for ‘Fast Roaming’ or ‘802.11r’.
- Enable fast roaming: Check or toggle the option to enable it. This will activate fast roaming for the selected wireless network.
Step 5: Save and test
- Save changes: Make sure to apply and save the changes you have made.
- Test connectivity: Connect devices to the network and move between AP coverage areas to test if the transition is smooth. Just to make sure everything is configured correctly.
Other notes/tips
- Compatible devices: Ensure that both the client devices and all APs in your network support fast roaming to avoid any compatibility issues.
- Network design: Fast Roaming works best in environments where multiple APs overlap in coverage. Properly plan AP placement to optimize performance.
- Security protocols: Use secure and compatible security protocols like WPA2-Enterprise for the best results with Fast Roaming.
Final thoughts
It makes to implement fast roaming in large offices or campuses to improve the connectivity. If you follow the steps outlined above, you can improve your network performance by making smooth transitions across APs within their networks.
If you need more assistence with optimizing complex networks, consider using a managed service provider like UniHosted. We provide managed UniFi Controllers to improve your network efficiency. Especially, if you are an MSP or IT-service business looking for reliable, secure and scalable UniFi network deployments, this is something for you.
Related guides
Keep reading
- SSIDs in UniFiFull guide on how to set up SSIDs within UniFi Read guide
- UniFi Controller ExplainedWhat does the UniFi Controller actually do? Plain breakdown of self-hosted, cloud, and managed hosting setups with the pros and cons of each. Read guide
- Best practices for managing UniFi devicesDetailed guide on how to improve managing UniFi devices. List of best practices outlined. Read guide
