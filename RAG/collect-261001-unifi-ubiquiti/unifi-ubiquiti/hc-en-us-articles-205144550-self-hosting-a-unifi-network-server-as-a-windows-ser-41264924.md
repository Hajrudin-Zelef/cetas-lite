---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-205144550-self-hosting-a-unifi-network-server-as-a-windows-ser-41264924
title: "hc-en-us-articles-205144550-self-hosting-a-unifi-network-server-as-a-windows-ser-41264924"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-205144550-self-hosting-a-unifi-network-server-as-a-windows-ser-41264924.md
source_anchor: ""
source_lines: [1, 27]
sha256: 7be2cfc3f9130f64cfd2fd5a89d2112f3bffa7a74ca91bdd54770a0b882afc13
---

# hc-en-us-articles-205144550-self-hosting-a-unifi-network-server-as-a-windows-ser-41264924

(Legacy) Self-Hosting a UniFi Network Server as a Windows Service
| Looking for the next generation of UniFi self-hosting? The UniFi OS Server is the new standard for self-hosting UniFi, offering support for advanced features like Organizations, IdP Integration, and Site Magic SD-WAN. Learn more. | 
Running a UniFi Network Server on Windows operating systems can be done using one of two methods:
- From the launcher: UniFi Network runs in the foreground (default method).
- As a Windows service: UniFi Network runs in the background (advanced, continue below).
New Set Up
1. Download the UniFi Network Server (Windows) from the Downloads page and run the installer. We recommend installing the latest version of the UniFi Network Server for the best experience.
2. Download the appropriate JRE package depending on your UniFi Network version.
3. Run the Java installer and set the Set JAVA_HOME variable to "Will be installed on local hard drive" as seen below.
5. Enable TCP Port 8080, TCP Port 8843, UDP Port 10001, and UDP Port 3478 on any local firewall (including Windows Defender) or antivirus software. See our Required Ports Reference to learn more.
6. Open a Windows Command Prompt (CMD) window as an admin.
7. Change the directory to the location of UniFi installation.
cd "%UserProfile%\Ubiquiti UniFi\"
8. Once in the root of the UniFi folder, run the following command to install the service:
java -jar lib\ace.jar installsvc
9. Wait for the installation to complete, indicated by the "Complete Installation" log message.
10. Start the UniFi Network Server service with the command below:
java -jar lib\ace.jar startsvc
11. Open a browser and navigate to the server's IP address or https://localhost:8443 to configure the UniFi Network Server.
Upgrade an Existing Set Up
1. Download a backup of the UniFi Network Server and close it.
3. Change the directory to the location of UniFi installation.
cd "%UserProfile%\Ubiquiti UniFi\"
4. Once in the root of the UniFi folder, issue the following to uninstall the service:
java -jar lib\ace.jar uninstallsvc
5. Wait for the service uninstall process to complete.
6. Follow steps 1-11 from the New Set Up section above to install an upgraded version of the UniFi Network Server as a service.
