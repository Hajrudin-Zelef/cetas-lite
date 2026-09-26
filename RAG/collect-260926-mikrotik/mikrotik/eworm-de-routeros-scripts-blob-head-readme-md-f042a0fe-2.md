---
id: collect-260926-mikrotik/mikrotik/eworm-de-routeros-scripts-blob-head-readme-md-f042a0fe-2
title: "eworm-de-routeros-scripts-blob-head-readme-md-f042a0fe"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["liability", "license"]
source: docs/RAG/lot-mikrotik/RouterOS/eworm-de-routeros-scripts-blob-head-readme-md-f042a0fe.md
source_anchor: ""
source_lines: [113, 177]
sha256: e1fca322017d8d951878dd56d3455ac4e28677384af371d4c50280f6f0678f6c
---

# eworm-de-routeros-scripts-blob-head-readme-md-f042a0fe

- Run other scripts on IPv4 DHCP server lease (dhcpv4-server-lease )
- Run other scripts on IPv6 DHCP client lease (dhcpv6-client-lease )
- Automatically upgrade firmware and reboot (firmware-upgrade-reboot )
- Download, import and update firewall address-lists (fw-addr-lists )
- Wait for global functions und modules (global-wait )
- Send GPS position to server (gps-track )
- Use WPA network with hotspot credentials (hotspot-to-wpa ,hotspot-to-wpa-lease &hotspot-to-wpa-cleanup )
- Create DNS records for IPSec peers (ipsec-to-dns )
- Update configuration on IPv6 prefix change (ipv6-update )
- Manage IP addresses with bridge status (ip-addr-bridge )
- Manage LEDs dark mode (leds-day-mode ,leds-night-mode &leds-toggle-mode )
- Forward log messages via notification (log-forward )
- Mode button with multiple presses (mode-button &mode-button-scheduler )
- Manage DNS and DoH servers from netwatch (netwatch-dns )
- Notify on host up and down (netwatch-notify )
- Visualize OSPF state via LEDs (ospf-to-leds )
- Manage system update (packages-update )
- Run scripts on ppp connection (ppp-on-up )
- Act on received SMS (sms-action )
- Forward received SMS (sms-forward )
- Play Super Mario theme (super-mario-theme )
- Chat with your router and send commands via Telegram bot (telegram-chat )
- Install LTE firmware upgrade (unattended-lte-firmware-upgrade )
- Update GRE configuration with dynamic addresses (update-gre-address )
- Update tunnelbroker configuration (update-tunnelbroker )
- Manage ports in bridge (mod/bridge-port-to )
- Manage VLANs on bridge ports (mod/bridge-port-vlan )
- Inspect variables (mod/inspectvar )
- IP address calculation (mod/ipcalc )
- Send notifications via e-mail (mod/notification-email )
- Send notifications via Gotify (mod/notification-gotify )
- Send notifications via Matrix (mod/notification-matrix )
- Send notifications via Ntfy (mod/notification-ntfy )
- Send notifications via Telegram (mod/notification-telegram )
- Download script and run it once (mod/scriptrunonce )
- Import ssh keys for public key authentication (mod/ssh-keys-import )
My scripts cover a lot of use cases, but you may have your own ones. You can
still use my scripts to manage and deploy yours, by specifying base-url
(and url-suffix) for each script.
This will fetch and install a script hello-world.rsc from the given url:
$ScriptInstallUpdate hello-world "base-url=https://git.eworm.de/cgit/routeros-scripts-custom/plain/";
For a script to be considered valid it has to begin with a magic token. Have a look at any script and copy the first line without modification.
Starting a script's name with mod/ makes it a module and it is run
automatically by global-functions.
⚠️ Warning: These links are being provided for your convenience only; they do not constitute an endorsement or an approval by me. I bear no responsibility for the accuracy, legality or content of the external site or for that of subsequent links. Contact the external site for answers to questions regarding its content.
- Hello World (This is a demo script to show how the linking to external documentation will be done.)
ℹ️ Info: You have your own set of scripts and/or modules and want these to be listed here? There should be a general info page that links here, and documentation for each script. You can start by cloning my Custom RouterOS-Scripts (or fork on GitHub or GitLab) and make your changes. Then please get in contact...
There is no specific function for script removal. Just remove it from configuration...
/system/script/remove to-be-removed;
Possibly a scheduler and other configuration has to be removed as well.
We have a Telegram Group RouterOS-Scripts 
Get help, give feedback or just chat - but do not expect free professional support!
Thanks a lot for past contributions! ❤️
Feel free to contact me via e-mail or open an issue or pull request at github.
This project is developed in private spare time and usage is free of charge
for you. If you like the scripts and think this is of value for you or your
business please consider to
donate with PayPal 
Thanks a lot for your support!
This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.
Our website contains links to the websites of third parties ("external links"). As the content of these websites is not under our control, we cannot assume any liability for such external content. In all cases, the provider of information of the linked websites is liable for the content and accuracy of the information provided. At the point in time when the links were placed, no infringements of the law were recognisable to us. As soon as an infringement of the law becomes known to us, we will immediately remove the link in question.
💡️ Hint: All external links are marked with an arrow pointing
diagonally in an up-right (or north-east) direction (
↗️ ).
