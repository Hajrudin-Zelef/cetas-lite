---
id: collect-261001-meraki/meraki/news-how-to-set-up-cisco-meraki-dashboard-yourself-in-6-steps-11437331-2
title: "news-how-to-set-up-cisco-meraki-dashboard-yourself-in-6-steps-11437331"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "license", "licenses"]
source: docs/RAG/collect-261001-meraki/news-how-to-set-up-cisco-meraki-dashboard-yourself-in-6-steps-11437331.md
source_anchor: ""
source_lines: [103, 213]
sha256: 37d4eb02bf063f1489550794411e7a1487be949363eb19b6d18da4f5602c5142
---

# news-how-to-set-up-cisco-meraki-dashboard-yourself-in-6-steps-11437331

After creating a specific network, you can view and configure it from the network page at the network level, where you can manage devices, monitor metrics, and apply configurations.
4. Create a Systems Management Network
There’s a slightly different process for creating a systems manager network for endpoint management:
- 
Log into your Meraki dashboard account.
- 
Click “Set up Systems Manager.”
- 
Click “Next.”
- 
Name your network. Or, if there is a “Network type” drop-down menu, click “EMM (Systems Manager).”
- 
Review and click “Create network.”
Once this step is complete, you’re able to configure the following endpoint management features:
- 
Profiles
- 
Payloads
- 
Applications
- 
Etc.
5. Add Devices
To claim a new device, such as an MX security appliance or MS Series switch, log into your Meraki dashboard, click the ‘Organization’ button, and choose ‘Inventory’. Enter the serial number of the new device in the inventory list and click the blue “Claim” button. Once claimed, the new device will appear in your inventory list and can be added to a specific network. Devices must be added to a network to download their configuration in Cisco Meraki.
Step-by-step:
- 
Log into your Meraki dashboard.
- 
Click ‘Organization’ and select ‘Inventory’.
- 
Enter the serial number of the new device (such as an MX appliance, MS Series switch, or other devices) next to the blue “Claim” button.
- 
Click “Claim.”
- 
Click the checkbox next to any devices you want to be added to the network.
- 
Click “Add to.”
- 
Under “Existing network,” choose the network you created earlier.
- 
Click “Add to existing.”
Note: The MX security appliance is set to send a DHCP request on its internet port by default. Devices require DHCP to pull an IP address and connect to the Meraki cloud.
Switch ports and other devices, such as access points and switches, can also be managed and configured from the dashboard. MS Series switches support advanced features for port management, including storm control, adaptive policies, and secure authentication.
For more information on configuring specific devices, check out these resources:
- 
How to Add and Configure a Cisco Meraki Access Point
- 
This Is How to Configure a Meraki Firewall
- 
How to Connect a Meraki Switch Locally
6. Add Licenses
If you ordered your Meraki device license separately, you might need to manually add it:
- 
Log into your Meraki dashboard.
- 
Click “Organization.”
- 
Click “Configure.”
- 
Click “License info.”
- 
Click “Add another license.”
- 
Click “License more devices” for the “Operation.”
- 
Enter the “License key.”
- 
Click “Add license.”
- 
Review the details for accuracy.
- 
Click “Add license.”
Congrats! Your Meraki network and devices are now ready for configuration.
Managing Multiple Devices
Managing multiple devices across your organization doesn’t have to be complicated. The Meraki dashboard streamlines the process, allowing users to add and oversee multiple devices from a single dashboard account. With intuitive tools for device network and hardware network management, you can organize your Meraki devices by location, function, or department, ensuring every device is accounted for and configured correctly.
The systems manager network feature further enhances your ability to manage multiple devices, providing centralized control over device policies, security settings, and application deployments. The Meraki dashboard also supports advanced configuration options such as IP assignment and static IP setup, making it easy to tailor your network to your organization’s needs. Features like spanning tree help prevent network loops and ensure optimal performance across all connected devices. With these capabilities, users can efficiently configure, monitor, and manage multiple devices, keeping their network secure, scalable, and easy to maintain.
3 Common Configuration Issues and How to Solve Them
Whether you’re an IT veteran or not, there is always room for potential technical difficulties in the Meraki configuration process. Here are three potential configuration issues you may run into and some tips on how to solve them.
If you are unable to access the Meraki dashboard during initial setup, try connecting an ethernet cable directly from your computer to the MX appliance. This allows you to access the Meraki offline setup page, where the following page (login page) will appear for initial configuration.
Note: Once devices are online, they will automatically download the latest firmware, which is indicated by a flashing LED. To verify device connectivity, look for a solid green light on Access Points or a solid white light on MX/MS devices.
1. Misconfigured DNS
Because Meraki devices rely on DNS to repair dashboard URLs, if your device is reporting issues with its DNS configuration, it’s likely not receiving responses to DNS requests.
Try these things to address this issue:
- 
Verify that the device’s static IP has been working, and make sure that the address is still valid.
- 
Check for a typo or otherwise incorrect value when assigning the static IP.
- 
Ensure the correct VLAN is used for DHCP.
Lastly, you can try switching to DHCP. You would have received this error message only if the device found another working IP address. Once you switch the IP assignment to DHCP instead of static, the device will use the correct address, and the error will go away eventually.
2. Conflicting Uplink IP Address
If you’re getting an alert about a conflicting uplink IP address, another device in the network is using the same IP address as your Meraki hardware.
If this problem persists, try the following tips:
- 
Make sure your devices all have unique IP addresses.
- 
Check by opening the client’s page on your dashboard and finding the matching IP addresses.
3. Poorly Configured IP Assignment
Here are some quick troubleshooting tips to address this issue:
- 
Verify that the device’s static IP has been working, and make sure that the address is still valid.
- 
Check for a typo or otherwise incorrect value when assigning the static IP.
- 
Ensure the correct VLAN is used for DHCP.
Best Practices for Configuration
To maximize the performance and security of your Cisco Meraki deployment, it’s essential to follow best practices for configuration. Start by creating a dedicated Meraki dashboard account and promptly adding all your Meraki devices to the dashboard. This ensures centralized visibility and control from the outset. When configuring your devices, assign static IP addresses where appropriate to maintain consistent connectivity and simplify troubleshooting.
Implement spanning tree protocols to prevent network loops and maintain a resilient network topology. Always keep your Meraki devices updated with the latest firmware and software releases—this not only unlocks new features but also ensures your network benefits from the latest security enhancements. Regularly monitor your network and devices through the Meraki dashboard, using built-in analytics and alerts to proactively address issues before they impact your users. By adhering to these best practices, you’ll ensure your Meraki network is robust, secure, and ready to scale as your organization grows.
Get Expert Guidance
If you find setting up your Cisco Meraki Dashboard to be outside your comfort level, above or below your paygrade, or beyond the time you have to spare, no worries.
Stratus Information Systems can give you the expert guidance you need to work through it yourself or do it completely for you. Check out our options for Cisco Meraki configurations and make sure you get in touch if we can help.
