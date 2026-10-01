---
id: collect-261001-general-networking/general-networking/c-en-us-td-docs-security-firepower-quick-start-fp2100-firepower-2100-gsg-ftd-fmc-84933923-1
title: "c-en-us-td-docs-security-firepower-quick-start-fp2100-firepower-2100-gsg-ftd-fmc-84933923"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "licenses"]
source: docs/RAG/collect-261001-general-networking/c-en-us-td-docs-security-firepower-quick-start-fp2100-firepower-2100-gsg-ftd-fmc-84933923.md
source_anchor: ""
source_lines: [1, 52]
sha256: 66ab3394f766e23e5d6b1d0261165eb2ce181bc8dcabc065730bf2db6ee5725c
---

# c-en-us-td-docs-security-firepower-quick-start-fp2100-firepower-2100-gsg-ftd-fmc-84933923

Before You Start
Deploy and perform initial configuration of the Firewall Management Center. See the getting started guide for your model.
| Note | Version 7.4 is the final release for the Firepower 2100. | 
Is This Chapter for You?
To see all available applications and managers, see Which Application and Manager is Right for You?. This chapter applies to the Firewall Threat Defense with the Firewall Management Center.
This chapter explains how to manage the Firewall Threat Defense with a Firewall Management Center located on your management network. For remote branch deployment, where the Firewall Management Center resides at a central headquarters, see Firewall Threat Defense Deployment with a Remote Firewall Management Center.
About the Firewall
The hardware can run either Firewall Threat Defense software or ASA software. Switching between Firewall Threat Defense and ASA requires you to reimage the device. You should also reimage if you need a different software version than is currently installed. See Cisco Secure Firewall ASA and Secure Firewall Threat Defense Reimage Guide.
The firewall runs an underlying operating system called the Secure Firewall eXtensible Operating System (FXOS). The firewall does not support the FXOS Secure Firewall Chassis Manager; only a limited CLI is supported for troubleshooting purposes. See the Cisco FXOS Troubleshooting Guide for the Firewall Threat Defense for more information.
Privacy Collection Statement—The firewall does not require or actively collect personally identifiable information. However, you can use personally identifiable information in the configuration, for example for usernames. In this case, an administrator might be able to see this information when working with the configuration or when using SNMP.
See the following tasks to deploy the Firewall Threat Defense with the Firewall Management Center.
|  | Pre-Configuration | Install the firewall. See the hardware installation guide. | 
|  | Pre-Configuration | Review the Network Deployment. | 
|  | Pre-Configuration | Cable the Device. | 
|  | Pre-Configuration | Power on the Device. | 
|  | CLI | (Optional) Check the Software and Install a New Version. | 
|  | CLI or Firewall Device Manager | Complete the Firewall Threat Defense Initial Configuration | 
|  | Firewall Management Center | Log Into the Firewall Management Center. | 
|  | Cisco Commerce Workspace | Obtain Licenses for the Firewall Management Center: Buy feature licenses. | 
|  | Smart Software Manager | Obtain Licenses for the Firewall Management Center: Generate a license token for the Firewall Management Center. | 
|  | Firewall Management Center | Obtain Licenses for the Firewall Management Center: Register the Firewall Management Center with the Smart Licensing server. | 
|  | Firewall Management Center | Register the Firewall Threat Defense with the Firewall Management Center. | 
|  | Firewall Management Center | Configure a Basic Security Policy. | 
The Firewall Management Center communicates with the Firewall Threat Defense on the Management interface.
The dedicated Management interface is a special interface with its own network settings:
By default, the Management 1/1 interface is enabled and configured as a DHCP client. If your network does not include a DHCP server, you can set the Management interface to use a static IP address during initial setup at the console port.
Both the Firewall Threat Defenseand the Firewall Management Center require internet access from their management interfaces for licensing and updates.
| Note | The management connection is a secure, TLS-1.3-encrypted communication channel between itself and the device. You do not need to run this traffic over an additional encrypted tunnel such as Site-to-Site VPN for security purposes. If the VPN goes down, for example, you will lose your management connection, so we recommend a simple management path. | 
You can configure other interfaces after you connect the Firewall Threat Defense to the Firewall Management Center.
The following figure shows a typical network deployment for the firewall where the Firewall Threat Defense, Firewall Management Center, and management computer connect to the management network.
The management network has a path to the internet for licensing and updates.
The following figure shows a typical network deployment for the firewall where:
Inside acts as the internet gateway for Management and for the Firewall Management Center.
Management 1/1 connects to an inside interface through a Layer 2 switch.
The Firewall Management Center and management computer connect to the switch.
This direct connection is allowed because the Management interface has separate routing from the other interfaces on the Firewall Threat Defense.
To cable one of the above scenarios on the Firepower 2100, see the following steps.
| Note | Other topologies can be used, and your deployment will vary depending on your basic logical network connectivity, ports, addressing, and configuration requirements. | 
| Step 1 | Install the chassis. See the hardware installation guide. | 
| Step 2 | Cable for a separate management network: | 
| Step 3 | Cable for an edge deployment: | 
| Note | For version 6.5 and earlier, the Management 1/1 default IP address is 192.168.45.45. | 
The power switch is located to the left of power supply module 1 on the rear of the chassis. It is a toggle switch that controls power to the system. If the power switch is in standby position, only the 3.3-V standby power is enabled from the power supply module and the 12-V main power is OFF. When the switch is in the ON position, the 12-V main power is turned on and the system boots.
| Note | The first time you boot up the Firewall Threat Defense, initialization can take approximately 15 to 30 minutes. | 
It's important that you provide reliable power for your device (for example, using an uninterruptable power supply (UPS)). Loss of power without first shutting down can cause serious file system damage. There are many processes running in the background all the time, and losing power does not allow the graceful shutdown of your system.
| Step 1 | Attach the power cord to the device and connect it to an electrical outlet. | 
| Step 2 | Press the power switch on the back of the device. | 
| Step 3 | Check the PWR LED on the front of the device; if it is solid green, the device is powered on. | 
| Step 4 | Check the SYS LED on the front of the device; after it is solid green, the system has passed power-on diagnostics. | 
| Note | Before you move the power switch to the OFF position, use the shutdown commands so that the system can perform a graceful                                                       shutdown. This may take several minutes to complete. After the graceful shutdown is complete, the console displays It is safe to power off now . The front panel blue locator beacon LED lights up indicating the system is ready to be powered off. You can now move the                                                       switch to the OFF position. The front panel PWR LED flashes momentarily and turns off. Do not remove the power until the PWR                                                       LED is completely off. See the FXOS Configuration Guide for more information on using the shutdown commands. | 
To check the software version and, if necessary, install a different version, perform these steps. We recommend that you install your target version before you configure the firewall. Alternatively, you can perform an upgrade after you are up and running, but upgrading, which preserves your configuration, may take longer than using this procedure.
What Version Should I Run?
