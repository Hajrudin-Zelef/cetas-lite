---
id: collect-261001-meraki/meraki/news-how-to-replace-an-existing-meraki-mx-device-0d7c85b9-1
title: "news-how-to-replace-an-existing-meraki-mx-device-0d7c85b9"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-meraki/news-how-to-replace-an-existing-meraki-mx-device-0d7c85b9.md
source_anchor: ""
source_lines: [1, 64]
sha256: 10df9673780ec086d4d17574405838f10cf0e44e652e7bb264c436d4d977a24f
---

# news-how-to-replace-an-existing-meraki-mx-device-0d7c85b9

Like any electronic device, there comes a time when you may have to replace a Cisco Meraki MX hardware. If the warranty still applies, you may request a replacement unit from Cisco Meraki (you may be asked to return the broken device at the company’s expense within 45 days of the replacement’s shipment).
What’s unique about the Cisco Meraki system is you can configure the replacement device as a “warm spare” or a secondary device that will smoothly take the place of the primary device. However, if the primary device fails before you can configure the secondary device, you will have to perform an “mx cold swap”—a manual replacement process required when the secondary device is not pre-configured or during network upgrades. The same rule applies when switching your network or upgrading to a different Meraki MX model.
Below, we’ll show you how to replace your Cisco Meraki MX with a warm spare and a cold swap methodology.
Pre-Installation Configuration
Once you received or bought a replacement Meraki device, the first thing you should do is manually configure all local settings:
- Static WAN IP address
- Proxies
- Non-standard link speeds
These settings must be configured manually on the replacement device to ensure proper connectivity and security.
Similarly, you’ll need to configure the network to accept the new device and prevent existing security programs from blocking it. For example, remember to whitelist the replacement unit’s MAC (media access control) address so that other devices won’t flag it as a rogue server if you use it to run a DHCP (Dynamic Host Configuration Protocol).
Before deployment, you should also perform a firmware update to ensure the replacement device is running the latest software version.
Below, we’ll cover some of the steps for replacing an existing Meraki device.
Pre-Replacement Checklist
Before you begin replacing your existing MX device, it’s crucial to follow a comprehensive pre-replacement checklist. Taking these steps will help ensure a seamless transition, maintain your network’s security, and minimize the risk of unexpected network downtime.
1. Review Existing Network ConfigurationStart by thoroughly documenting your current network configuration. This includes noting all settings on your existing MX, such as VLANs, firewall rules, DHCP configurations, static WAN IP addresses, and any custom routing or security policies. Having a clear record of your existing network setup will make it easier to replicate these settings on the new device and avoid configuration errors.
2. Perform Security VerificationBefore making any changes, conduct a security verification of your existing MX and network environment. Check for any active security alerts, review recent logs for suspicious activity, and ensure that your security appliance is up to date with the latest firmware. This step helps protect your network from vulnerabilities during the transition and ensures that the new device will inherit a secure baseline.
3. Prepare the New DeviceUnbox and inspect the new MX device, confirming that it matches your intended model and is compatible with your existing network. Pre-configure the new device with essential settings—such as static IP addresses, management VLANs, and any required security credentials—before connecting it to your network. This preparation reduces the time your network is offline and helps prevent misconfigurations.
4. Backup and Document EverythingCreate backups of your current network configuration and document all changes you plan to make. This step is vital for troubleshooting and for restoring service quickly if any issues arise during the replacement process.
5. Notify Stakeholders and Schedule DowntimeCommunicate with your team and any affected users about the planned replacement. Schedule the swap during a maintenance window to minimize the impact of any unavoidable network downtime.
By carefully performing security verification and reviewing your existing network configuration before introducing the new device, you can safeguard your network’s integrity and ensure a smooth, secure transition to your new MX security appliance.
Warm Spare: How to Set Up a High-Availability (HA) Pair
A warm spare failover is designed to prevent downtime and ensure the integrity of the MX service and the functions it performs for your network. It establishes failover support at the appliance level by having a secondary MX on the ready. If the primary MX goes offline, the secondary unit will automatically take over.
A few reminders about HA Pairs:
- The secondary MX must be the exact same model as the primary device.
- The dashboard will show a swap button for primary and secondary devices. It is for assigning which device will act as the primary MX and not to test if the secondary device works.
- The only way to test a warm spare failover is to disconnect the uplink to the primary MX completely.
- The secondary MX should be pre configured with the necessary settings from the dashboard before physically connecting the device to the network. This ensures seamless failover if the primary device fails.
- When adding an MX to a network with an existing MX of the same model, it will automatically be set to warm spare mode.
Step 1: Dashboard Configuration
- From the dashboard, go to Security & SD-WAN > Monitor > Appliance status > Configure warm spare.
- Wait for the new window to open, then click Enabled.
- Enter the serial number of the secondary Meraki MX device.
- Choose the correct IP configuration.
  - MX Uplink IPs – This option allows the active MX device to use its distinct uplink IP. So if the primary MX fails and the secondary MX takes over, the latter will use its own IP, which is different from the primary.
  - Virtual Uplink IPs – This option uses a virtual IP (VIP) that both primary and secondary MXs use. It requires an additional public IP for each new uplink but ensures a seamless failover.
- Click Update.
Step 2: Configuration to Warm Spare
There are two ways to establish an HA pair:
- Via Passthrough or VPN Concentrator Mode
- Via Routed Mode
Let’s start with the VPN Concentrator method:
- Security & SD-WAN > Configure > Addressing & VLANs page
- Deploy the two MX devices through Passthrough or VPN concentrator mode.
- Check that the two devices:
  - Connect to the network only through Internet ports.
  - Are not connected to the network via LAN ports.
  - Are within the same IP subnet.
  - Communicate with one another and the dashboard.
- Set up the VIP that the primary and warm spare will share. Go to Security & SD-WAN > Monitor > Appliance status > SPARE
- Check that the VIP differs from the primary and spare devices’ IPs.
- Check that the VIP is in the same subnet as the primary and spare’s IPs.
Next are the steps via Routed Mode.
The process is similar to the VPN Concentrator method, except without the VPN configuration:
- Set up the VIP that the primary and warm spare will share. Go to Security & SD-WAN > Monitor > Appliance status > SPARE
- Check that the VIP differs from the primary and spare devices’ IPs.
- Check that the VIP is in the same subnet as the primary and spare’s IPs.
Advantages of HA Pairing:
- Only one license is needed for both MX units.
- There will be minimal to zero network downtime if the primary Meraki device fails.
- You can stay calm if existing MX hardware fails because the failover will automatically kick into action. The spare will recognize that the primary device has failed, so it will take over as the active MX, delegating the primary device as the passive MX.
- There’s also no need for manual, on-site intervention by network administration specialists.
Disadvantages of HA Pairing:
- You must obtain a similar model to your primary MX.
- Warm spares using different models are unsupported. If you have to use a different model, you must use the cold swap methodology to replace your Meraki MX device.
