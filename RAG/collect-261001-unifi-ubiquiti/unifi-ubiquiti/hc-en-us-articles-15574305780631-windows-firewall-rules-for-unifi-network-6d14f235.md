---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-15574305780631-windows-firewall-rules-for-unifi-network-6d14f235
title: "hc-en-us-articles-15574305780631-windows-firewall-rules-for-unifi-network-6d14f235"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-15574305780631-windows-firewall-rules-for-unifi-network-6d14f235.md
source_anchor: ""
source_lines: [1, 22]
sha256: c2a2652c3962f190dd9e58baf583e8ff39c6418f37565feb6d13fb1775210ca0
---

# hc-en-us-articles-15574305780631-windows-firewall-rules-for-unifi-network-6d14f235

Windows Firewall Rules for UniFi Network
When using a self-hosted UniFi Network Server on Windows, the UniFi Network Application needs to be able to communicate with the UniFi devices on the network and allowed through the Windows Firewall.
Note: This guide applies only to self-hosted UniFi Network, not Cloud Gateways. We strongly recommend UniFi Cloud Gateways, for the most seamless experience possible.
Which Programs need to be Allowed?
Both the Java and Java Server programs need to be allowed through the Windows Firewall. The steps below assume that you installed the UniFi Network Server using the Help Center article here and used OpenJDK Java.
1. Open the Windows Control Panel and navigate to Windows Defender Firewall.
2. Select Allow an app or feature through Windows Defender Firewall.
3. If the options are greyed out, select Change settings.
4. Check both OpenJDK Platform binary entries and the Private and Public boxes.
5. Apply the changes.
6. To verify if the rules are present, navigate to Windows Defender Firewall > Advanced Settings and select Inbound Rules.
7. Scroll down and the OpenJDK Platform binary entries will be shown.
Adding the Rules Manually
The Windows Firewall rules can also be added manually. The commands below assume that you installed the UniFi Network Server using the Help Center article here and used OpenJDK Java.
1. Open an Administrative Command Prompt.
2. Copy and add the below entries, the version and path may differ on your system.
netsh advfirewall firewall add rule name="UniFi Network Server (Java)" protocol=TCP dir=in action=allow program="C:\Program Files\Eclipse Adoptium\jre-11.0.19.7-hotspot\bin\java.exe" profile=private,public enable=yes
netsh advfirewall firewall add rule name="UniFi Network Server (Java)" protocol=UDP dir=in action=allow program="C:\Program Files\Eclipse Adoptium\jre-11.0.19.7-hotspot\bin\java.exe" profile=private,public enable=yes
netsh advfirewall firewall add rule name="UniFi Network Server (Java Server)" protocol=TCP dir=in action=allow program="C:\Program Files\Eclipse Adoptium\jre-11.0.19.7-hotspot\bin\javaw.exe" profile=private,public enable=yes
netsh advfirewall firewall add rule name="UniFi Network Server (Java Server)" protocol=UDP dir=in action=allow program="C:\Program Files\Eclipse Adoptium\jre-11.0.19.7-hotspot\bin\javaw.exe" profile=private,public enable=yes
3. To verify if the rules are present, open Windows Defender Firewall with Advanced Security and select Inbound Rules.
4. Scroll down and the UniFi Network Server (Java Server) and (Java) entries will be shown.
