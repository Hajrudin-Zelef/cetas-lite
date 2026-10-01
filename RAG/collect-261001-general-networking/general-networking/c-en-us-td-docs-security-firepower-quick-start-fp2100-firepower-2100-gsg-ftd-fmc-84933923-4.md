---
id: collect-261001-general-networking/general-networking/c-en-us-td-docs-security-firepower-quick-start-fp2100-firepower-2100-gsg-ftd-fmc-84933923-4
title: "c-en-us-td-docs-security-firepower-quick-start-fp2100-firepower-2100-gsg-ftd-fmc-84933923"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-general-networking/c-en-us-td-docs-security-firepower-quick-start-fp2100-firepower-2100-gsg-ftd-fmc-84933923.md
source_anchor: ""
source_lines: [111, 168]
sha256: 8a0d9560e8441ede47385108d6b22f61be8a628638dd3b275fecf671011362b1
---

# c-en-us-td-docs-security-firepower-quick-start-fp2100-firepower-2100-gsg-ftd-fmc-84933923

In any case, you need to perform additional interface configuration after you register the device. Enable the Firewall Threat Defense interfaces, assign them to security zones, and set the IP addresses. .
The following example configures a routed mode inside interface with a static address and a routed mode outside interface using DHCP.
| Step 1 | Choose , and click the Edit () for the firewall. | 
| Step 2 | Click Interfaces. | 
| Step 3 | Click Edit () for the interface that you want to use for inside. The General tab appears. | 
| Step 4 | Click the Edit () for the interface that you want to use for outside. The General tab appears. | 
| Step 5 | Click Save. | 
Enable the DHCP server if you want clients to use DHCP to obtain IP addresses from the Firewall Threat Defense.
| Step 1 | Choose , and click Edit () for the device. | 
| Step 2 | Choose . | 
| Step 3 | On the Server page, click Add, and configure the following options:  | 
| Step 4 | Click OK. | 
| Step 5 | Click Save. | 
The default route normally points to the upstream router reachable from the outside interface. If you use DHCP for the outside interface, your device might have already received a default route. If you need to manually add the route, complete this procedure. If you received a default route from the DHCP server, it will show in the IPv4 Routes or IPv6 Routes table on the page.
| Step 1 | Choose , and click Edit () for the device. | 
| Step 2 | Choose . | 
| Step 3 | Click Add Route, and set the following:  | 
| Step 4 | Click OK. The route is added to the static route table. | 
| Step 5 | Click Save. | 
A typical NAT rule converts internal addresses to a port on the outside interface IP address. This type of NAT rule is called interface Port Address Translation (PAT).
| Step 1 | Choose , and click . | 
| Step 2 | Name the policy, select the device(s) that you want to use the policy, and click Save. The policy is added the Firewall Management Center. You still have to add rules to the policy. | 
| Step 3 | Click Add Rule. The Add NAT Rule dialog box appears. | 
| Step 4 | Configure the basic rule options:  | 
| Step 5 | On the Interface Objects page, add the outside zone from the Available Interface Objects area to the Destination Interface Objects area. | 
| Step 6 | On the Translation page, configure the following options:  | 
| Step 7 | Click Save to add the rule. The rule is saved to the Rules table. | 
| Step 8 | Click Save on the NAT page to save your changes. | 
If you created a basic Block all traffic access control policy when you registered the Firewall Threat Defense, then you need to add rules to the policy to allow traffic through the device. The following procedure adds a rule to allow traffic from the inside zone to the outside zone. If you have other zones, be sure to add rules allowing traffic to the appropriate networks.
| Step 1 | Choose , and click Edit () for the access control policy assigned to the Firewall Threat Defense. | 
| Step 2 | Click Add Rule, and set the following parameters:  Leave the other settings as is. | 
| Step 3 | Click Apply. The rule is added to the Rules table. | 
| Step 4 | Click Save. | 
Deploy the configuration changes to the Firewall Threat Defense; none of your changes are active on the device until you deploy them.
| Step 1 | Click Deploy in the upper right. | 
| Step 2 | For a quick deployment, check specific devices and then click Deploy, or click Deploy All to deploy to all devices. Otherwise, for additional deployment options, click Advanced Deploy. | 
| Step 3 | Ensure that the deployment succeeds. Click the icon to the right of the Deploy button in the menu bar to see status for deployments. | 
Use the command-line interface (CLI) to set up the system and do basic system troubleshooting. You cannot configure policies through a CLI session. You can access the CLI by connecting to the console port.
You can also access the FXOS CLI for troubleshooting purposes.
| Note | You can alternatively SSH to the Management interface of the Firewall Threat Defense device. Unlike a console session, the SSH session defaults to the Firewall Threat Defense CLI, from which you can connect to the FXOS CLI using the connect fxos command. You can later connect to the address on a data interface if you open the interface for SSH connections. SSH access to data interfaces is disabled by default. This procedure describes console port access, which defaults to the FXOS CLI. | 
| Step 1 | To log into the CLI, connect your management computer to the console port. The Firepower 2100 ships with a DB-9 to RJ-45 serial cable, so you may need a third party DB-9-to-USB serial cable to make the connection. Be sure to install any necessary USB serial drivers for your operating system. The console port defaults to the FXOS CLI. Use the following serial settings:  You connect to the FXOS CLI. Log in to the CLI using the admin username and the password you set at initial setup (the default is Admin123). Example:  firepower login: admin Password: Last login: Thu May 16 14:01:03 UTC 2019 on ttyS0 Successful login attempts for user 'admin' : 1  firepower#    | 
| Step 2 | Access the Firewall Threat Defense CLI. connect ftd Example:  firepower# connect ftd >  After logging in, for information on the commands available in the CLI, enter help or ? . For usage information, see Cisco Secure Firewall Threat Defense Command Reference. | 
| Step 3 | To exit the Firewall Threat Defense CLI, enter the exit or logout command. This command returns you to the FXOS CLI prompt. For information on the commands available in the FXOS CLI, enter ? . Example:  > exit firepower#   | 
It's important that you shut down your system properly. Simply unplugging the power or pressing the power switch can cause serious file system damage. Remember that there are many processes running in the background all the time, and unplugging or shutting off the power does not allow the graceful shutdown of your firewall system.
You can power off the device using the Firewall Management Center device management page, or you can use the FXOS CLI.
It's important that you shut down your system properly. Simply unplugging the power or pressing the power switch can cause serious file system damage. Remember that there are many processes running in the background all the time, and unplugging or shutting off the power does not allow the graceful shutdown of your firewall.
You can shut down your system properly using the Firewall Management Center.
| Step 1 | Choose . | 
| Step 2 | Next to the device that you want to restart, click Edit (). | 
| Step 3 | Click the Device tab. | 
| Step 4 | Click Shut Down Device () in the System section. | 
| Step 5 | When prompted, confirm that you want to shut down the device. | 
| Step 6 | If you have a console connection to the firewall, monitor the system prompts as the firewall shuts down. You will see the following prompt:  System is stopped. It is safe to power off now.  Do you want to reboot instead? [y/N] If you do not have a console connection, wait approximately 3 minutes to ensure the system has shut down. | 
| Step 7 | You can now turn off the power switch and unplug the power to physically remove power from the chassis if necessary. | 
You can use the FXOS CLI to safely shut down the system and power off the device. You access the CLI by connecting to the console port; see Access the Firewall Threat Defense and FXOS CLI.
| Step 1 | In the FXOS CLI, connect to local-mgmt: firepower # connect local-mgmt | 
| Step 2 | Issue the shutdown command: firepower(local-mgmt) # shutdown Example: firepower(local-mgmt)# shutdown  This command will shutdown the system.  Continue? Please enter 'YES' or 'NO': yes INIT: Stopping Cisco Threat Defense......ok | 
| Step 3 | Monitor the system prompts as the firewall shuts down. You will see the following prompt:  System is stopped. It is safe to power off now. Do you want to reboot instead? [y/N]  | 
