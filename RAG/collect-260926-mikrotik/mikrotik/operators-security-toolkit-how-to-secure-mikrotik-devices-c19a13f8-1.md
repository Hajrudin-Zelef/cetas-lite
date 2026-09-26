---
id: collect-260926-mikrotik/mikrotik/operators-security-toolkit-how-to-secure-mikrotik-devices-c19a13f8-1
title: "operators-security-toolkit-how-to-secure-mikrotik-devices-c19a13f8"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "throughput"]
source: docs/RAG/lot-mikrotik/forum/misc/operators-security-toolkit-how-to-secure-mikrotik-devices-c19a13f8.md
source_anchor: ""
source_lines: [1, 68]
sha256: 4068900365bcdba98c6e1785930262f11c7bb2472d167af948b9b542b9efe25c
---

# operators-security-toolkit-how-to-secure-mikrotik-devices-c19a13f8

(Version 1.0)
Mikrotik devices are wonderful networking tools. They offer flexibility and cost empowerment to solve networking problems. But, the way we deploy Mikrotiks in the industry is creating multiple security risks. People are not spending the time to secure Mikrotik devices.
- Organizations deploying Mikrotik devices are creating security risks inside their organization.
- ISPs who are not tracking the Mikrotik deployments on their customers are creating a risk for their business, network, and other customers.
- The Industry who are not pushing for better “out of the box” Mikrotic security is accepting massive DDoS attack, gateways for ransomware crews, and a range of other abuse.
It cannot be ignored that Mikrotik devices are 2022’s most dangerous malware platform. We hope this work will help engineers, administrators, and organizations seek out, secure, and clean up the Mikrotik devices in their network.
Step-by-Step Guide for Securing your Mikrotik Device
First, with any network device, there are core principles that lead to security and resiliency. These are general principles that we will point out in this guide. They are applied to any network vendor.
Second, we will present several deployment scenarios. In this scenario, we are working to secure a basic Mikrotik router connected to the Internet (see the figure).
Lastly, if you go through these steps and find problems, ASSUME YOUR ROUTER IS COMPROMISED! At the end of this process, you will need to do a NETINSTALL, rebuild the Mikrotik device, and start from scratch.
WHY?
The modern compromises to Mikrotik devices can get into the underlying Linux operating system. These RouterOS remediation steps cannot fix problems inserted into the device’s Linux. It must be rebuilt.
Step 0 – Assume your Mikrotik is infected, owned, and controlled by a Miscreant.
The number of exposed, vulnerable, and known Mikrotik devices is in the millions. Given this, it is best to assume that your device is controlled by a miscreant using your device for criminal activities. Criminal activities that put your network at risk.
Through this guide, we will continuously assume the Mikrotik device is compromised. Each step allows us to build confidence in the configuration and the deployment.
Step 1 – Backup Your Device
Yes! The first step is to back up your Mikrotik device and copy the backup to a safe location. This action is common sense before making any changes to a device or the network.
Mikrotik’s new BACKUP documentation provides details. (https://help.mikrotik.com/docs/display/ROS/Backup)
Don’t depend on the Mikrotik Cloud Backup!
Since RouterOS v6.44 it is possible to securely store your device’s backup file on MikroTik’s Cloud servers; read more about this feature on the Mikrotik IP/Cloud page. But what happens if you cannot get to the cloud backup? It is common sense to have a local backup location along with the convenience of a cloud backup. Do both.
Step 2 – Upgrade Mikrotik’s Winbox
Ease of use is one of the core reasons why there are so many Mikrotiks deployed. People do not need to be network experts to get value. Mikrotik’s Winbox application is one of the key reasons. Winbox is a small utility that allows the administration of MikroTik RouterOS using a fast and simple GUI. It is a native Win32 binary but can be run on Linux and macOS (OSX) using Wine. All Winbox interface functions are as close as possible, mirroring the console functions, which is why there are no Winbox sections in the manual.
Winbox loader can be downloaded from the MikroTik download page.
You now have confidence that you have the latest version of Winbox deployed on your network.
Step 3 – Upgrade RouterOS
Moving to an updated version of the RouterOS software is the first step. It is not a “security fix” but a critical security step.
Mikrotik’s Upgrading and Installation documentation provides details on how to upgrade. Upgrading to the latest stable version of Router OS software is recommended.
Step 4 – Username and Passwords
Mikrotik out-of-the-box has a default username of “admin” and NO PASSWORD! Yes, this contradicts the industry’s best common practices (BCPs). It exposes your brand new Mikrotik device to miscreants constantly scanning the Internet looking for new devices that have yet to change the default username and add a password. Mikrotik’s documentation has you connecting the device to the network with no security, then says ….
“Now anyone worldwide can access our router, so it is the best time to protect it from intruders and basic attacks.”
Just assume the miscreants have your Mikrotik username and password. Change your username and password!
Mikrotik’s Protecting the Router documentation has details to change the password. Here is a practical workflow:
- Change the Password of the user “admin” on the device. Remember the new password.
- Add new users to maintain the router. These new users would replace the “admin” user.
- Test the new username and passwords to ensure you have access. Have a separate username/password for your Winbox access.
- Test the Winbox configuration to ensure it works with the new usernames/passwords.
- Delete the user “admin.”
Test your Passwords to see if compromised
The Have I been PWND site is a public service that allows people to check their usernames, emails, and passwords. Testing any new password on Have I been PWND is prudent, validating the exposure.
https://haveibeenpwned.com/
Step 5 – Limit the Source IP Blocks to Connect to the Router
The world does not need to connect to your router. The world should not be telneting or trying to use Winbox from all over the world. Limit the IP Source address used for the username and the most critical services.
- Restrict username access for the specific IP address. In this example, we’re using the network from the illustration 192.168.2.0/24
    /user set 0 allowed-address=192.168.2.0/24
- Restrict Winbox service access to be only from a specific subnet:
    /ip service set winbox address=192.168.2.0/24
Step 6 – Change the port # for SSH
Changing the SSH port number is a trick to minimize brute-force password hacking on your router. Change and test the port before turning off the telnet service on the router!
The port number for SSH is 22 by default. Here, we change the default SSH port number from 22 to 2200. This will make it harder for miscreants’ “scanning tools” to find your SSH port and “brute force” password guessing.
/ip service set ssh port=2200
Test the SSH connection! Ensure you can SSH to the new port (2200 in this example).
Step 7 – Limit the Services Opened on the Router
Once you know SSH works on the new port, you can turn off the Telnet Service:
/ip service print
/ip service disable telnet
/ip service print
Don’t stop by turning off telnet. There is a range of services that Mikrotik turns on that are not needed.
The Best Common Practice (BCP) for security on network devices is to turn off everything! Then turn on specific functions/services as part of your network design. Minimize RISK by minimizing what is running in the background on your network device.
Start with disabling these services:
/ip service print
/ip service disable telnet,ftp,www,api,api-ssl
/ip service print
A bandwidth server is used to test throughput between two MikroTik routers. Disable it in the production environment. This can be used as a DDoS tool.
/tool bandwidth-server set enabled=no
Don’t let the router be a DNS DDoS Reflector. DNS cache might be configured on the router, which turns it into a DDoS Reflector. Make sure it it turned off.
/ip dns set allow-remote-requests=no
If you need the DNS Cache to have the router, explicitly set up the DNS function to restrict who can use the Mikrotik router for DNS traffic (i.e. only inside your network).
Step 8 – Turn off Direct Access via the MAC Address
