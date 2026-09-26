---
id: collect-260926-mikrotik/mikrotik/operators-security-toolkit-how-to-secure-mikrotik-devices-c19a13f8-3
title: "operators-security-toolkit-how-to-secure-mikrotik-devices-c19a13f8"
domain: mikrotik
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["cost", "open source", "research"]
source: docs/RAG/lot-mikrotik/forum/misc/operators-security-toolkit-how-to-secure-mikrotik-devices-c19a13f8.md
source_anchor: ""
source_lines: [130, 161]
sha256: 2f0ff7462b2c66233da05bba6d86d35820d8c7ec893b0d473efc7ebec20ac9c0
---

# operators-security-toolkit-how-to-secure-mikrotik-devices-c19a13f8

If the following data exists; it might indicate a Trickbot infection:
chain=dstnat action=dst-nat to-addresses=
to-ports=80 protocol=tcp dst-address=dst-port=449 
chain=srcnat action=masquerade src-address=
Run the following command to remove the potentially malicious NAT rule:
/ip firewall nat remove numbers=
Step FINAL – NETINSTALL and Start from Scratch
The final step is the reality check. There is no way to RouterOS commands can fix malware inserted into the device’s Linux operating system. This process helps to shape your troubleshooting, confirm if there is a compromise, and the process to configure Mikrotik to be secure.
Netinstall is a tool for installing and reinstalling MikroTik devices running RouterOS. Always try using Netinstall if you suspect your device is not working properly.
Tools to Check your Mikrotik Device
Severial companies have tool to help you track, check, and secure your Mikrotik devices.
Shadowserver’s Daily Network Report
Shadowserver provides any organization with IP addresses, Autonomous System Numbers (ASNs), or Domain Names with daily report based on their vast array of security telemetry. The Daily Network Reports are FREE and a Public Service to protect the Internet. Mikrotik devices that are exposed, vulnerable, or compromised are listed in the reports. Find out more about the Network Reports and How to Subscribe to the reports on wwww.shadowserver.org.
Meris RouterOS Checker
The Meris RouterOS Checker is a open source tool crafted by Eclypsium to help network admins check their Mikrotik devices.
https://github.com/eclypsium/mikrotik_meris_checker
Routeros-scanner
Microsoft crafted routeros-scanner as an open source tool to help organizations clean up Trickbot exploitations on their Mikrotik devices.
https://github.com/microsoft/routeros-scanner
RouterOS Security Research
Tenable is curating all their Mikrotik RouterOS security research on this RouterOS Security Research Github.
https://github.com/tenable/routeros.
What a Youtube Video – Securing Mikrotik Devices
There is a Youtube Playlist for a “how to secure your Mikrotik” device here: Mikrotik Router Security – How to keep control over your Mikrotik devices. It is good to watch and listen to many of your peers working to secure their devices. Taking action to protect your network is the most critical element!
- Mikrotik Tutorial no. 36 – 7 Things to do to Secure Mikrotik Router
- MikroTik RouterOS Securing Your Router and Good Security Practises
- Security analysis of recent RouterOS exploits
Are you looking for more practical, low-cost security Advice?
- You can sign up to the mailing list for updates here: Stay Connected with Senki’s Updates.
- Subscribe to Senki’s YOUTUBE Channel for videos on this and other security topics.
- Ask questions to Barry Greene – bgreene@senki.org
The materials and guides posted on www.senki.org here are designed to help organizations leverage the talent around them to get started with their security activities. Start with the Operator’s Security Toolkit and Meaningful Security Conversations with your Vendors. Each is no-nonsense security for all Operators. It provides details to help them build more security resilient networks. In the meantime, stay connected to the Senki Community to get updates on new empowerment and security insights.
