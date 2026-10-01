---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-360008976393-backups-and-migration-in-unifi-3d7a8e01-2
title: "hc-en-us-articles-360008976393-backups-and-migration-in-unifi-3d7a8e01"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-360008976393-backups-and-migration-in-unifi-3d7a8e01.md
source_anchor: ""
source_lines: [61, 72]
sha256: 0c552380d03c2b5e3777a3a263a7b8362d411d83d197cf7ac771b6c39bcafc68
---

# hc-en-us-articles-360008976393-backups-and-migration-in-unifi-3d7a8e01

Some older UniFi hosting options, such as CloudKey, Official UniFi Hosting, and self-hosted UniFI OS Servers, support cross-site management. Site export allows users to migrate specific sites to another UniFi system.
- Ensure that your new Network Application is up to date.
- Go to Network Settings > System > Site Management and click Export Site to download the site file and initiate the guided walkthrough.
- Ensure your new Cloud Gateway or UniFi OS Server has been set up (leave all other UniFi devices connected as they were previously).
- Open your new Network Application in a separate tab.
- Select Import Site from the site switcher in the new Network Application and upload the site export file from step (2).
  - If no site switcher is visible, enable Multi-Site Management in Settings > System > Site Management.
- Complete the guided walkthrough in the old Network Application by selecting the device(s) you wish to migrate.
- Enter the Inform URL (IP address of the new Cloud Gateway).
- Forget devices from the original Network Application.
- Wait for all devices to appear as Online.
  - If a device appears offline or 'Managed by Another Console,' factory reset and re-adopt it.
