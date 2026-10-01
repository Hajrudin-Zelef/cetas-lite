---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-4416276882327-how-to-set-up-unifi-8be28c32
title: "hc-en-us-articles-4416276882327-how-to-set-up-unifi-8be28c32"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["attention", "ethernet"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-4416276882327-how-to-set-up-unifi-8be28c32.md
source_anchor: ""
source_lines: [1, 51]
sha256: 74d7cf31529101d31482d973be74bc8a672eac3a3b38e6b36a733c4efcff0afc
---

# hc-en-us-articles-4416276882327-how-to-set-up-unifi-8be28c32

How to Set Up UniFi
UniFi OS is pre-installed on UniFi Consoles, streamlining the setup process so you can get your network up and running quickly. Although it is possible to self-host the UniFi Network Server or setup Access Points (APs) in Standalone Mode, these methods lack key advantages including automated backups, system updates, and more advanced software offerings.
If you are new to UniFi, we encourage you to check out Introduction to UniFi to learn more about device adoption, remote management, and the UniFi ecosystem in general.
Have You Purchased a UniFi Console?
Our UniFi Consoles are built to run UniFi OS and come pre-installed with all UniFi Applications. This includes things like the Dream Machine Pro, Dream Router and CloudKey G2 Plus. With these consoles, it's easy to automatically establish backups, update hardware and software, and ensure maximum compatibility by minimizing third-party inter-operation.
Set Up a Console
UniFi Consoles can be setup from the web or our mobile app. Regardless of which method you choose, you will be able to continue managing UniFi both via the mobile app and web following the initial setup.
- Download the UniFi Mobile App (iOS / Android).
- Connect your mobile device to your UniFi Console (Dream Machine Pro, Dream Router, CloudKey Gen2 Plus, etc.) via WiFi or Bluetooth. Remain within 3 meters (10 feet) if using Bluetooth.
- Follow the setup wizard to complete the process.
  - Though optional, we strongly recommend creating a UI Account, which gives you access to features like:
    - Secure management from anywhere in the world using our UniFi Site Manager, or one of our UniFi Mobile Apps
    - Automatic system backups linked to your UI Account
    - Pre-configured email notifications for convenient system monitoring
- Though optional, we strongly recommend creating a UI Account, which gives you access to features like:
- Once complete, you can proceed with adopting UniFi devices such as APs, switches, and cameras.
- Connect a laptop to your UniFi Console by joining the same WiFi, or using an ethernet cable.
- Open your internet browser and enter your UniFi Console's IP address. This can be found on the console's LCD screen, or by using the network scanner in our WiFiman Mobile App (iOS / Android). See Sign In to UniFi for more information.
 Note: UniFi Cloud Gateways can be configured by navigating tounifi/ in your browser.
- You may receive a privacy or security warning, which is normal. Select to proceed anyway.
- Follow the setup wizard to complete the process.
- Once complete, you can proceed with adopting UniFi devices such as APs, switches, and cameras.
Do You Have an Access Point Without a Console?
Standalone Mode should only be used when configuring a small quantity of APs because they do not support seamless roaming. Instead, you will need to manually connect to each AP as you move out of range. Also, APs must have a wired connection to your network; wireless meshing is not possible. Finally, remote management is not possible because it is done entirely through the UniFi Mobile app while locally connected to your WiFi.
For information on Standalone Access Points, check out our article here.
Do You Plan to Host a Server Yourself?
UniFi can also be self-hosted on a Windows/macOS/Linux machine, but this is only recommended for users proficient with network administration, because it requires constant and precise attention. You’ll need to manually configure your network for sufficient communication between UniFi Network and all of your networking devices. Success hinges on your ability to continuously navigate various third-party software, firewalls and other networking devices.
For information on Self-Hosting, check out our article here.
Are You Using Official UniFi Hosting?
Official UniFi Hosting is a cloud-only solution for UniFi Network device management, offering scalable and remote management of up to 1,000 UniFi devices.
For information on Official UniFi Hosting, check out our article here.
Common Troubleshooting
Our setup flow is designed to be as seamless as possible. However, some circumstances, often related to external network configurations, may cause setup failures.
Check Your Physical Connections
Ensure that your power and network connections are secure. If they are, but your issue persists, we recommend using our UniFi cables for optimal compatibility and performance.
Verify Your UniFi Console's Internet Connection
UniFi Cloud Gateways (UniFi Consoles such as our Dream products, which contain an integrated security gateway) are usually connected directly to the internet. However, non-gateway consoles, such as the CloudKey Gen2 Plus, can be incorrectly placed on a network or VLAN with restricted internet access. Please ensure this isn’t the case for your consoles.
Update to the Latest Version of UniFi OS
For the best setup experience, we recommend updating to the latest version and ensuring that automatic updates are enabled.
Open All Required Ports
Third-party gateways, firewalls, and Internet Service Provider (ISP) modems occasionally block traffic required for set up. Some of these include:
- TCP & UDP Port 443
- TCP Port 8883
- TCP & UDP Port 53
- UDP Port 123
- TCP Port 3478
See Required Port Reference for a complete listing of all requisite ports. You may wish to contact your ISP to ensure none of these are being blocked. We also recommend disabling any third-party firewalls.
Clone Your ISP Router or Modem’s MAC Address
Some ISPs restrict connectivity to certain MAC addresses. You may need to clone your modem or router’s MAC address in order to set up your UniFi Console. We recommend contacting your ISP first to confirm whether or not this is necessary.
Contact Your ISP
If you’ve tried all of the tips above but are still having setup issues, please contact your ISP to ensure that they have not restricted your UniFi Console's connectivity.
