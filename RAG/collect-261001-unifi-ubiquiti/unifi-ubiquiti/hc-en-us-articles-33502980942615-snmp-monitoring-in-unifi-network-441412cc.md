---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-33502980942615-snmp-monitoring-in-unifi-network-441412cc
title: "hc-en-us-articles-33502980942615-snmp-monitoring-in-unifi-network-441412cc"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-33502980942615-snmp-monitoring-in-unifi-network-441412cc.md
source_anchor: ""
source_lines: [1, 34]
sha256: d7efffe04f0c7d317a0f9bdaa4d78cf8e287765c33ef0f708bebdd41e1dd0ad0
---

# hc-en-us-articles-33502980942615-snmp-monitoring-in-unifi-network-441412cc

SNMP Monitoring in UniFi Network
Simple Network Management Protocol (SNMP) allows you to monitor UniFi devices using third-party tools like Zabbix, PRTG, or Nagios. This provides visibility into device performance and health metrics, helping you troubleshoot and optimize your network.
Once enabled, your SNMP Manager can poll UniFi devices for OID values (Object Identifiers). Uploading UniFi’s MIB (Management Information Base) files to your SNMP Manager improves readability by translating numeric OIDs into named metrics.
Requirements
To use SNMP, you’ll need:
- UniFi Network is running with compatible UniFi Devices adopted
  - Note: USW Flex and Ultra switches do not support SNMP.
- At least one SNMP-capable UniFi device
- An SNMP monitoring platform (e.g., Zabbix, PRTG, Nagios)
- Network access between the SNMP Manager and your UniFi devices
Configuring SNMP
- Within UniFi Network, go to to Settings > CyberSecure > Traffic Logging.
- Scroll to the SNMP section and enable it.
- Select the SNMP version(s) to enable:
  - v1/v2c: Use a community string to allow read-only access (default: public, but this should be changed).
  - 
v3: Enter a username and password for authenticated, optionally encrypted access.
    - Authentication Protocol: SHA
    - Privacy Protocol: AES-128
    - For some services (Datadog for example), a dedicated field for the password may not be present. If there are authentication and privacy key entries, please use the password for both fields.
- Click Apply Changes.
- (Optional) Configure Location and Contact metadata by navigating to Devices > select a device > Settings > toggle SNMP > enter Location and Contact > click Apply Changes.
- Set up your SNMP Manager to poll the UniFi devices using the configured community string or SNMP v3 credentials.
Note: SNMP traps are not currently available, but will be added in a future release.
MIB Files
UniFi provides a MIB file that can be uploaded to your SNMP Manager:
These allow your SNMP Manager to parse the OID responses in human-readable form (e.g., eth0 instead of .1.3.6.1.2.1.2.2.1.2.2). Polling works without these files, but output will be numeric.
Note: These MIBs are not comprehensive and may not include all OIDs, particularly for newer or model-specific devices.
Troubleshooting
If you are not receiving SNMP data:
- Verify SNMP is enabled in Settings > CyberSecure > Traffic Logging
- Check that your community string (v1/v2c) or credentials (v3) match on both ends
- If there is a firewall or other security device between the SNMP Manager and UniFi devices, make sure UDP port 161 is allowed
Restart the device or UniFi Network application if changes do not take effect
