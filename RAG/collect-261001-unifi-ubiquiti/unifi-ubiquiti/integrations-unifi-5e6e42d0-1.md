---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/integrations-unifi-5e6e42d0-1
title: "integrations-unifi-5e6e42d0"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/integrations-unifi-5e6e42d0.md
source_anchor: ""
source_lines: [1, 34]
sha256: 7cb739f6f96ffe67c3d99904a19d412f04970962f46c7cc359065a320d452bfb
---

# integrations-unifi-5e6e42d0

UniFi Network
UniFi Network by Ubiquiti Networks, inc. is a software that binds gateways, switches and wireless access points together with one graphical front end.
With this integrationIntegrations connect and integrate Home Assistant with your devices, services, and more. [Learn more], you can bring your UniFi Network into Home Assistant to automate and monitor your network. Common use cases include:
- Use connected clients as presence detection to trigger automations when family members arrive home or leave.
- Control Wi-Fi availability on a schedule, for example to disable guest networks overnight or pause kids’ Wi-Fi during homework time.
- Monitor bandwidth usage and uptime of clients and network devices.
- Control PoE power on individual switch ports to remotely restart connected devices like cameras or access points.
- Toggle firewall rules, port forwarding, or traffic rules as part of broader home automations.
- Get notified about firmware updates and install them from Home Assistant.
Prerequisites
Hardware support
This integration supports all UniFi OS Consoles that run UniFi Network. It also supports self hosted versions of UniFi Network.
Software support
It is recommended to run latest stable versions of UniFi Network and UniFi OS.
Early Access and Release Candidate versions are not supported by Home Assistant.
Using Early Access Release Candidate versions of UniFi Network or UniFi OS can bring unexpected changes. If you choose to opt into either the Early Access or the Release Candidate release channel and anything breaks in Home Assistant, you will need to wait until that version goes to the official Stable Release channel before it is expected to work.
Local user
You need a local user created in your UniFi OS Console. Ubiquiti SSO cloud users will not work. Using an administrator or a user with full read/write access is recommended to get the most out of the integration, but it is not required. The entities that are created automatically adjust based on the permissions of the user you use.
- Sign in to your UniFi OS device.
- Go to People from the left-hand side menu.
- Select Create New.
- Fill in the user details:
  - Add Name.
  - Check Admin.
  - Set a username and password.
  - Check Restrict to local access only.
  - Uncheck Use a pre-defined role.
  - Set the first privilege level (Network) to Full Management.
  - Set the second privilege level (OS Settings) to None.
- In the bottom right, select Create.
There is currently support for the following device types within Home Assistant:
Configuration
To add the UniFi Network hub to your Home Assistant instance, use this My button:
      
