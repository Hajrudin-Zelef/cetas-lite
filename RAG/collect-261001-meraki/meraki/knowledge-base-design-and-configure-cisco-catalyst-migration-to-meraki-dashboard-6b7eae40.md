---
id: collect-261001-meraki/meraki/knowledge-base-design-and-configure-cisco-catalyst-migration-to-meraki-dashboard-6b7eae40
title: "knowledge-base-design-and-configure-cisco-catalyst-migration-to-meraki-dashboard-6b7eae40"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/knowledge-base-design-and-configure-cisco-catalyst-migration-to-meraki-dashboard-6b7eae40.md
source_anchor: ""
source_lines: [1, 28]
sha256: bb2fabe5a6d2112d5f506823f9227b26ade4d5f5a9543bb1bdccb5aab7ff0a60
---

# knowledge-base-design-and-configure-cisco-catalyst-migration-to-meraki-dashboard-6b7eae40

Platforms: Catalyst Switches Serie 9200, 9200L, 9300, 9300L, 9300X, 9500.
Preparing the Meraki dashboard (pre-check)
Check if you have access to the Meraki dashboard. You can register for free, instructions here.
Obtain an API key from your existing dashboard. First, check that “Dashboard API Access” is checked in the Organization -> Settings section. If it is not, select it and save your changes. Then, in the My profile section, generate a new API key (only works for accounts with full admin permission).
You must be able to connect to api.meraki.com and meraki-cloud-monitoring-onboarding-app.s3.amazonaws.com via TCP port 443. It is recommended to disable HTTPS traffic inspection during migration. Any form of tampering with certificates may cause the process to fail.
If you still need to configure the dashboard, add a new “switch” or “combined” network (note: the “combined” network should include the “switch” network). Information on how to do this can be found here.
Download and install the migration application from the Meraki dashboard (Network wide -> Add Devices).
Switch preparation (pre-check)
Connect the switch to the Internet via the front port (not the management interface!).
Provide SSH access to the switch from the computer you will be migrating from.
Start ip routing. The default gateway set with the ip default-gateway command is not supported during migration!
Start AAA mode with the aaa new-model command.
Make sure you have level 15 clearance
Make sure you have a DNS server configured that can resolve the FQDN eu.tlsgw.meraki.com (for the European area). You can add it with the command ip name-server [DNS server IP address].
Start name resolution with the ip domain lookup command.
Make sure you have the correct time and date set on the switch and the NTP server configured. You can add it with the command ntp server [NTP server IP address]. You can check the correctness of the time and date with the show clock command.
Back up your switch configuration before starting the migration process.
Step-by step migration
Run the application. After accepting the regulations, we start by providing the previously generated API key.
Go through the onboarding process.
3. Check the configuration with the Start pre-check button. If the switch, dashboard and computer have been prepared correctly, the test should pass.
4. Assign the migrated switch to the network previously prepared in the Meraki dashboard.
5. Confirm migration.
The dashboard needs about 3-5 minutes to complete the operation. A clear message will inform you about the success of the process. Once you see it, you can add additional devices.
Result of Cisco Catalyst migration to Meraki dashboard
The switch is visible in the dashboard.
Now you can log into the switch’s CLI and make changes to it. If you want to remove a switch from the dashboard, you select it and select the Remove from Network option.
A more detailed description of the migration can be found on our blog.
