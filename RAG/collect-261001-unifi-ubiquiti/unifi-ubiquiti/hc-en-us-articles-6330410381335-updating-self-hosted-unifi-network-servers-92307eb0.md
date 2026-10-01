---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-6330410381335-updating-self-hosted-unifi-network-servers-92307eb0
title: "hc-en-us-articles-6330410381335-updating-self-hosted-unifi-network-servers-92307eb0"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-6330410381335-updating-self-hosted-unifi-network-servers-92307eb0.md
source_anchor: ""
source_lines: [1, 16]
sha256: fbecd4b58c4860732a619abd942e7de500bec407eb8e7dbc009f12b4e81439d3
---

# hc-en-us-articles-6330410381335-updating-self-hosted-unifi-network-servers-92307eb0

(Legacy) Updating Self-Hosted UniFi Network Servers
| Looking for the next generation of UniFi self-hosting? The UniFi OS Server is the new standard for self-hosting UniFi, offering support for advanced features like Organizations, IdP Integration, and Site Magic SD-WAN. Learn more. | 
Updates for a self-hosted UniFi Network Server must be manually checked for and initiated each time.
If your UniFi Network application is running on a dedicated UniFi Console, such as our recommended UniFi Cloud Gateways, this updating guide does not apply (see our UniFi Updating article instead).
Users self-hosting a UniFi Network Server on a Linux machine should follow our Linux updating instructions.
Updating the Network Application
- Download the latest Network Application version from here. Note that this must be downloaded on the original machine hosting the UniFi Network Server.
- 
Close any instances of the Network Application that are running, prior to the installation.
  - Your network will continue to function as normal (devices will remain connected with internet access, and traffic will continue to be routed).
- 
Run the downloaded file and follow the setup wizard to finish the update.
  - We recommend downloading a backup file, found in the UniFi Network System Settings.
Note: macOS users may be required to move the downloaded file into the Applications folder, or right-click > open the file in order to begin the installation.
(Advanced) Updating via CLI on Linux-hosted Applications
It is also possible to use APT to manage updates on Debian- and Ubuntu-based installations. Refer to this article for more details. This should only be attempted by users with appropriate knowledge of Linux.
