---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-51-1
title: "Summary"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["latency", "license", "memory"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-51.md
source_anchor: ""
source_lines: [1, 88]
sha256: 84b5bcdeceef45e7d1cb4f3b336ce5336e3988b6e0f270831ec092e90c2bd944
---

# Summary

MikroTik offers multiple services for your RouterBOARD devices that are connected to the Internet. These services are meant to ease the inconveniences when configuring, setting up, controlling, maintaining, or monitoring your device. A more detailed list of available services that IP/Cloud can provide can be found below.

# Services

Be aware that if the router has multiple public IP addresses and/or multiple internet gateways, the exact IP used for communicating with MikroTik's Cloud server may not be as expected!

IP/Cloud requires a paid perpetual license for Cloud Hosted Router (CHR).

## DDNS

DDNS or Dynamic DNS is a service that updates the IPv4 address for A records and the IPv6 address for AAAA records periodically. Such a service is very useful when your ISP has provided a dynamic IP address that changes periodically, but you always need an address that you can use to connect to your device remotely. Below you can find operation details that are relevant to the IP/Cloud's DDNS service:

- Checks for outgoing IP address change: every 60 seconds
- Waits for the MikroTik's Cloud server's response: 15 seconds
- DDNS record TTL: 60 seconds
- Sends encrypted packets to **cloud2.mikrotik.com** using UDP/15252 port

Since RouterOS v6.43 if your device is able to reach **cloud2.mikrotik.com** using IPv6, then a DNS **AAAA** record is going to be created for your public IPv6 address. If your device is only able to reach cloud2.mikrotik.com using IPv4, then only a DNS **A** record is going to be created for your public IPv4 address. cloud.mikrotik.com is used for older RouterOS versions prior 6.44

To enable the DDNS service:

When the service is enabled, a DNS name will be stored on MikroTik's Cloud server permanently and this DNS name will resolve to the last IP that your RouterOS instance has sent to MikroTik's Cloud server.

To disable the DDNS service:

Before 7.17, the default value for ddns-enabled was "no". In versions after 7.17 and including, if you want to disable DDNS, make sure to disable the Back To Home feature first, if it was enabled, then set "ddns-enabled=auto"

As soon as you disable the service, your device sends a command to MikroTik's Cloud server to remove the stored DNS name.

To manually trigger a DNS update:

To actually connect to the device using the DNS name provided by the cloud server, a user must configure the router's firewall to permit such access from the WAN port. (Default MikroTik configuration does not permit access to services such as WebFig, WinBox, etc. from the WAN port).

## Update time

Correct time on a device is important, it causes issues with the system's logs, breaks HTTPS connectivity to the device, tunnel connectivity, and other issues. To have your system's clock updated, you can use NTP or SNTP, though it requires you to specify an IP address for the NTP Server. In most cases, NTP/SNTP is not required in order to simply have a correct time set on the device, for simplicity you can use the IP Cloud's update time service. Below you can find operation details that are relevant to the IP/Cloud's update time service:

- Approximate time (accuracy of several seconds, depends on UDP packet latency)
- Updates time after a reboot and during every DDNS update (when router's WAN IP address changes or after the force-update command is used)
- Sends encrypted packets to **cloud2.mikrotik.com** using UDP/15252 port
- Detects time-zone depending on the router's public IP address and our commercial database

To enable the time update service:

To enable automatic time zone detection:

## Backup

It is possible to store your device's backup on MikroTik's Cloud server. The backup service allows you to upload an encrypted backup file, download it and apply the backup file to your device as long as your device is able to reach MikroTik's Cloud server. Below you can find operation details that are relevant to the IP/Cloud's backup service:

- 1 free backup slot for each device
- Allowed backup size: 15MB
- Sends encrypted packets to **cloud2.mikrotik.com** using UDP/15252 and TCP/15252 port

To create a new backup and upload it the MikroTik's Cloud server:

The `create-and-upload` action command will create a new system's backup file, encrypt the backup file with AES using the provided password and upload it. For `upload` action command the password property has no effect since the `upload` action command uploads only already created system's backup files. 

To download the uploaded backup file and save it to the device's memory:

**Warning:** The secret-download-key is a unique identifier that can be used to download your encrypted backup to your other devices. Since you can download your encrypted backup from any location and any device by using the secret-download-key, then you should try to keep this identifier a secret. The downloaded backup is still encrypted using AES, nevertheless, make sure you are using a strong password! 

To remove the uploaded backup:

To replace an existing file with a new backup file, use the following command:

To upload an existing backup file (created previously):

Make sure that the backup was encrypted using AES, otherwise, the IP/Cloud will reject the backup upload. Since there is only 1 free backup slot per device, then you need to remove the existing backup before uploading a new one.

## Back to Home

For more info about Back to Home (BTH) service, see the separate documentation page.

## File share

For more info about File Share service, see the separate documentation page.

## Relay service

Back to home and File Share both partially rely on the MikroTik cloud relay service. All transmissions through the relay service are end-to-end encrypted, relay is purely to faclilitate connection and is designed to never require decryption of user data or metadata. See respective manuals for details on how each service uses the relay.

# Properties

**Sub-menu:** `/ip cloud`

