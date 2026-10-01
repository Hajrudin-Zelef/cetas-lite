---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/set-up-a-unifi-hotspot-portal-and-guest-wifi
title: "set-up-a-unifi-hotspot-portal-and-guest-wifi"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Stripe"]
dates: []
keywords: ["fine-tuning", "parameters"]
source: docs/RAG/collect-261001-unifi-ubiquiti/set-up-a-unifi-hotspot-portal-and-guest-wifi.md
source_anchor: ""
source_lines: [1, 105]
sha256: a12fe37c53cfb9e0abb4d0853ce32e9dd1391de2f6746c829b3dcf35afa15502
---

# set-up-a-unifi-hotspot-portal-and-guest-wifi

Setting up a UniFi Hotspot Portal (aka “Captive Portal”) and Guest WiFi offers a seamless and secure internet experience for guests while keeping your core network isolated. Whether for a café, business, or event, UniFi’s hotspot capabilities allow you to manage guest access through portals, vouchers, or even payment integration. This guide combines advanced Wi-Fi settings and step-by-step configuration to create a scalable and customizable guest WiFi network.

Prerequisites

UniFi Controller (latest version preferred)

UniFi Access Points (e.g., U6-LR, U6-Pro)

VLAN for Guests (optional but recommended)

UXG Gateway or equivalent for routing and firewall setup

Step 1: Log in to the UniFi Controller

Access your UniFi Controller via browser: https://<controller-ip>:8443

Log in with your admin credentials.

Navigate to Settings from the left-hand sidebar.

Step 2: Create the Guest WiFi Network

Go to WiFi Networks:

In Settings, select WiFi under Networks.

Add New Network:

Click Create New WiFi Network.

Assign a name (SSID), e.g., “Guest WiFi”.

Security Options:

Choose WPA3 or WPA2 based on device compatibility.

If it’s an open network, leave the password field blank for public access.

VLAN Assignment (Optional):

Use VLAN tagging to segment guest traffic (e.g., VLAN ID 10 for guests).

Enable Captive portal:

Select the Captive portal in Hotspot 2.0.

A captive portal forces users to interact with a login page (Guest Portal) before accessing the network, which is useful for public WiFi, guest authentication, and managing network traffic.

Step 3: Configure the Hotspot Portal

Configure Guest Portal:

Navigate to Insights Menu → Select Hotspot and click on Landing page button.

Branding:

UniFi provides an intuitive way to add custom branding to your landing page. Configurable options include Title, Welcome Text, Login Button, Logo, Success Text, and complete Color Schemes.

Authentication:

Easily select one or multiple ways for your guests to connect

Facebook: Guests can authenticate using a Facebook account.

Password: Guests must enter a password to connect.

Payment: Guests must pay to use the WiFi. This is currently supported by Stripe.

Vouchers: Provide guests with vouchers that can be used to authenticate. Customize vouchers to support various expiration times, bandwidth limits, or data consumption quotas.

RADIUS (advanced): Preconfigure guest authentication via a RADIUS server.

External Portal Server (advanced): Integrate with a third-party portal server.

Step 4: Create Vouchers (Optional)

In the Hotspot Manager, click on Vouchers.

Set parameters such as:

Time limit, data usage, or bandwidth restrictions.

Print and distribute the generated codes for guest access.

Step 5: Bandwidth Limiting

Go to Settings → Profiles → User Groups.

Create a group called “Guests” and set Rate Limits (e.g., 5 Mbps download, 1 Mbps upload).

Apply this user group to your guest WiFi network.

Step 6: Monitoring and Fine-Tuning

Guest Insights:

Use the Insights tab in the controller to monitor guest connections, traffic, and bandwidth usage.

Adjust Advanced Settings:

Fine-tune features like Minimum RSSI, Airtime Fairness, and Band Steering to improve guest experience and optimize network performance based on device density and traffic volume.

Conclusion

By following these steps, you’ve successfully configured a UniFi Hotspot Portal and Guest WiFi with enhanced security, user management, and a customizable guest experience. Leveraging UniFi’s powerful tools, you can ensure a reliable, fast, and secure guest network with tailored access control and a branded portal.
