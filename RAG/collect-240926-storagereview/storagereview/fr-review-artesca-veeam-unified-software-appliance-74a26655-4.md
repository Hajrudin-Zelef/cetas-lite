---
id: collect-240926-storagereview/storagereview/fr-review-artesca-veeam-unified-software-appliance-74a26655-4
title: "fr-review-artesca-veeam-unified-software-appliance-74a26655"
domain: storagereview
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/clean_en/storagereview/fr-review-artesca-veeam-unified-software-appliance-74a26655.md
source_anchor: ""
source_lines: [44, 64]
sha256: edd9283ea851cb496be5bebdd21bd432725b7e199496b8793ec341a764636e35
---

# fr-review-artesca-veeam-unified-software-appliance-74a26655

To test the ARTESCA+ Veeam unified software appliance, we connected the test unit to our Proxmox VE (PVE) server. Backing up virtual machines is done in three simple steps:
- Connecting the PVE cluster to Veeam
- Adding a worker virtual machine
- Creating a backup job
We will present the startup steps with Veeam, but a more detailed guide is available here: Veeam Plug-in for Proxmox VE 3 User Guide. To connect Veeam to PVE via the Backup & Replication console, go to the "Inventory" view, then select "Add Server" after right-clicking on the "Virtual Infrastructure" list item.
When the "Add Server" menu appears, follow the wizard's instructions to connect a Proxmox VE instance. You will be asked for the server's IP address or hostname and credentials, so make sure you have this information handy for your cluster.
After connecting the PVE, a worker virtual machine is needed to facilitate data transfers between the Proxmox VE server and Veeam. In the "Backup Infrastructure" view, right-click on "Backup Proxies," then select "Add Proxy" from the dropdown menu.
Add a worker process using the "Proxmox VE worker" wizard from the context menu, then select the desired host, CPU, memory, and network configuration. Name it (we used "artesca-worker") and complete the wizard.
Finally, we need to create a backup job to start backing up virtual machines on Proxmox VE. In the "Home" view of Veeam Backup & Replication, click on "Jobs," then "Backup." Right-click on "Backup," hover over "Backup" in the dropdown menu, and select "Virtual Machine."
A menu will then appear, allowing you to select the backup job settings. Be sure to include all virtual machines to be backed up using the "Add" button in the "Virtual Machines" step. Also select the ARTESCA+ S3 storage bucket as the backup repository. Once these steps are complete, the backup job will run according to your schedule.
Once you have configured Veeam to back up and store your hypervisors and data sources, the ARTESCA+ Veeam unified software appliance will run reliably in the background, protecting your organization's information and giving you peace of mind.
The future of Scality and Veeam
Scality previewed the next major version of ARTESCA+ and promised greater simplicity, better integration, and increased automation for configuring the embedded Veeam virtual machine. The company also plans to add support for deploying the Veeam software appliance (Linux-based) and many improvements to multi-node and high-availability features. This update will reduce the total deployment time of the solution and remove several configuration steps described in this article. Furthermore, although this document was written for Veeam B&R V12, organizations using V13 will see improved overall object storage performance and efficiency, and benefit from the latest platform protection and security features.
Conclusion
ARTESCA+ Veeam fully delivers on Scality's promises. In our lab, this unified software appliance eliminated backup and storage uncertainties by centralizing the S3 path, enabling versioning and object lock by default, and guiding installation through a simple and intuitive wizard usable by a generalist IT professional. The result: a functional and immutable repository in minutes, with minimal maintenance.
Security is focused on practicality rather than performance. Internal-only access points, no external DNS, credential isolation, Identity and Access Management (IAM), and optional multi-factor authentication (MFA) reduce exposure to risk while preserving day-to-day administrative simplicity. Grafana monitoring provides full visibility into services, nodes, and disks, making it easy to verify that backups are working properly and detect issues before they cause service disruptions.
From a business perspective, the added value lies in speed of implementation and reduced operational costs. Teams benefit from Veeam and an S3 target within a single infrastructure on standard x86, allowing them to maintain great flexibility in hardware choice and avoid vendor lock-in. Scality positions this unified appliance for usable capacity of approximately 20 to 440 TB, which corresponds to the needs of branch offices, edge sites, and SMBs that have outgrown the capabilities of a scale-up platform but do not want to deploy a separate object platform.
A few precautions should be taken into account. The unified appliance is a single node. For greater resilience or higher capacity, deploy ARTESCA in a multi-node cluster and keep Veeam as the control plane. External S3 exposure is possible, but it is an option to be enabled that must be documented and managed with separate credentials.
Overall, this is a mature approach, building on Scality's long experience and a team with decades of enterprise storage experience, combining sound defaults, strong immutability, and simple operations.
Platform demonstration
This report is sponsored by Scality. All views and opinions expressed in this report are based on our impartial evaluation of the product(s) studied.
