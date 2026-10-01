---
id: collect-261001-meraki/meraki/news-how-to-replace-an-existing-meraki-mx-device-0d7c85b9-2
title: "news-how-to-replace-an-existing-meraki-mx-device-0d7c85b9"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/news-how-to-replace-an-existing-meraki-mx-device-0d7c85b9.md
source_anchor: ""
source_lines: [65, 143]
sha256: 098e215fdaa89e346898069e5de97d895ee72d135e063d6d3b5d9beadd73f5be
---

# news-how-to-replace-an-existing-meraki-mx-device-0d7c85b9

Cold Swap: How To Replace an Existing MX Device With a Different Model
The cold swap methodology allows you to use a different MX model than your current one. It’s an excellent alternative to HA pairing if you intend to upgrade your device or switch to another network.
There are two ways to perform a cold swap replacement:
- Quick Swap
- Clone and Replace
1. Quick Swap
Since the dashboard network only has room for one MX device at a time, you need to remove the existing MX before installing the replacement.
Here are a few things to keep in mind:
- The replacement will be installed in the same network as the original MX.
- Quick swap preserves the original, non-local configurations.
Step 1: Removing the Original MX
- Go to Security & SD-WAN > Monitor > Appliance Status.
- Scroll down the page and click Remove Appliance From Network.
Step 2: Adding the New MX Via the Inventory Page
- Go to Organization > Inventory Page.
- Select the MAC address and Serial Number of the new MX device you want to add to the network.
- Click Add to… > Existing Network > find the network where the original MX belongs > click Add to existing.
Adding to a New Network
If you want to create a new network:
- Go to the Organization tab > Create Network.
- Create a name for the new network.
- Select your preferred Network Type.
- Select your preferred Configuration.
- Go to the Devices section > Check the MX devices you want to add to the network > click Claim > enter the serial numbers > click Create network.
Step 2 Alternative: Adding the New MX Via the Network Configuration Page
- From the Administrator view (with multiple organizations), go to the Search for network drop-down menu > choose the network where you want to add the device.
- Go to Network-wide > Configure > Add devices.
- Select the devices you want to add > click Add devices
Step 3: Add the Secondary MX as a Warm Spare
- From the Administrator view (with multiple organizations), go to the Search for network drop-down menu > choose the network where you added the new Meraki device.
- Go to Security Appliance > Appliance Status > Configure Warm Spare.
If Site-to-site VPN was previously disabled, it should be re-enabled through the device’s configuration interface after adding the new MX.
Step 4: Physically Swap the Old and New Devices
- Allow the device to check into the dashboard properly.
- Finish any necessary device upgrades.
- Transfer the WAN uplinks.
- Next, transfer all the LAN connections.
IMPORTANT: All the cables from the old MX should go to the same slots/ports in the new MX.
Advantages of Quick Swap:
- There’s no need to reconfigure the replacement MX in the dashboard.
- All previous client tracking data will be retained.
- Adding new networks or deleting existing ones is not required.
- It’s the easiest way to replace an existing Meraki MX device in a Combined Network.
- This is the best option if your existing MX device malfunctions and is already causing network downtime, and you have no contingency plans yet.
Disadvantages of Quick Swap:
- There’s downtime when you configure and physically install the new MX.
- The Clone and Replace method results in less downtime.
2. Clone and Replace
As the name suggests, this process turns the replacement device into a clone of the original Meraki MX. This method is best done early as a contingency rather than a solution after hardware failure. It involves pre-staging, configuring the replacement, and letting it check into a network identical to the existing MX network.
Step 1: Clone the Existing Network
- Go to the Organization tab > Create Network.
- Create a name for the new network.
- Select your preferred Network Type.
- Go to Network Configuration > Clone from existing network > choose the original network you want to clone.
Step 2: Add the Replacement MX Device to the Clone Network
- From the Administrator view (with multiple organizations), go to the Search for network drop-down menu > choose the network where you want to add the device.
- Go to Network-wide > Configure > Add devices.
- Select the devices you want to add > click Add devices
Step 3: Add the Secondary MX as a Warm Spare
- From the Administrator view (with multiple organizations), go to the Search for network drop-down menu > choose the network where you added the new Meraki device.
- Go to Security Appliance > Appliance Status > Configure Warm Spare.
- Bring the replacement device online.
- Allow the device to pull configurations and firmware updates.
Step 4: Physically Swap the Old and New Devices
- Allow the device to check into the dashboard properly.
- Finish any necessary device upgrades.
- Transfer the WAN uplinks.
- Next, transfer all the LAN connections.
IMPORTANT: All the cables from the old MX should go to the same slots/ports in the new MX.
Advantages of Clone and Replace:
- Minimal downtime; since the replacement is already configured and tested, the only downtime is when you physically swap the old device with the new one.
- Less disruptive to your operations.
Disadvantages of Clone and Replace:
- The clone network will not have any historical client tracking data (the data will only exist in the original network).
Get IT Management Support from the Experts
Our guide above is here to help anyone managing Cisco Meraki security appliances and networks. If you are more comfortable entrusting these tasks to experienced IT specialists, Stratus Information Systems can help.
We offer Meraki support and consulting services for on-site and remote management. Our knowledgeable team at Stratus can assist you with hardware installation, network configuration, monitoring, and more. We also have Cisco-certified engineers who can walk you through the process of replacing Meraki devices.
Our security service can help protect your network and websites by detecting and blocking malicious bots during security verification processes.
Our services ensure network integrity and minimize disruptive and costly downtimes. Schedule a consultation with one of our experts today.
