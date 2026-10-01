---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f-4
title: "how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f.md
source_anchor: ""
source_lines: [183, 241]
sha256: 3624d20180e6f96dd7c4f1c7b9dc1f4093c5928355f2e34f21fc7d7db6748eff
---

# how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f

Next up is configuring the interfaces. I highly recommend you consider how you want to layout your network before creating the interfaces because you may find that it is very disruptive to reconfigure the interfaces at a later time since it likely involves temporarily taking down some or all of your network.
In this example, I will be configuring a LAG interface to demonstrate the process, but you may skip this section and go to the VLAN section if you do not want to use a LAG interface. The process of creating VLANs is the same whether or not you are using a LAG interface. The only difference is the physical interface you choose to associate to the VLAN.
Note
If you plan to use Zenarmor, I highly recommend you test if LAGs will function properly before committing to using a LAG. Since I tested my OPNsense configuration on a separate system wile using my old firewall appliance, I could verify that the LAG would function well with Zenarmor.
Your mileage may vary depending on the type of hardware you are using since there may be issues with netmap that prevent certain types of network interfaces from working properly in LAGs. You may be safe if you have igb based 1G interfaces such as those found on the Protectli VP2410 because it has been stable for me for an extended period of time.
Also in discussions on Twitter, it was mentioned that perhaps issues are more likely to occur with LAGs on networks that have heavy traffic on a regular basis such as saturated business/enterprise networks.
Interfaces: Settings
| Option | Value | 
|---|---|
| Hardware CRC | Check “Disable hardware checksum offload” (if not already checked) | 
| Hardware TSO | Check “Disable hardware TCP segmentation offload” (if not already checked) | 
| Hardware LRO | Check “Disable hardware large receive offload” (if not already checked) | 
| VLAN Hardware Filtering | Choose the “Disable VLAN Hardware Filtering” option | 
I have often seen the recommendation to disable hardware offloading on the network interfaces due to various issues that may be encountered. For most users, it is always best to leave it off unless you have thoroughly tested out these configuration options.
In a home network, hardware offloading (if it works) is likely less impactful than using it on a heavily saturated business or enterprise network unless you regularly saturate your network’s bandwidth.
If you are using IDS/IPS services such as Suricata or Zenarmor, hardware offloading should be disabled since it is incompatible with netmap.
Other Types: LAGG
As mentioned earlier, I am going to assume a default OPNsense installation where only the WAN and LAN interfaces are enabled and the remaining interfaces were not assigned during installation. You should be connected to the LAN interface while configuring the interfaces in OPNsense. This allows you to easily change the interface configuration without dropping your connection or accidentally locking yourself out. Hence, the advantage of dedicating one of the network interfaces for management and administration purposes as described in this guide!
The Protectli box in my example has 4 interfaces, which means I am able to use the remaining 2 interfaces to create a LAG interface. Go to the “Interfaces > Other Types > LAGG” page and click on the “+” button to configure a LAG interface.
In order to assign the parent interfaces to a LAG, they must not be currently assigned. The dropdown box will only show interfaces which are allowed to be added to a LAG interface.
| Option | Value | 
|---|---|
| Parent interface | Select the interfaces to be in a LAG (in this example: igb2 and igb3) | 
| Lag proto | LACP (you will need a switch which supports LACP – configuration is described later) | 
| Description | VLAN LAG (you may enter your own value here) | 
| MTU | Leave blank to use default value | 
Other Types: VLAN
With the LAG created, you can now create the VLANs that will be associated to the LAG interface. Once the VLANs are associated to the LAG interface, you will be able to assign VLAN interfaces in the same way as physical interfaces as shown in the next section.
In this example, I am going to create several VLANs to demonstrate different use cases you may want to use in your network. You definitely do not need to have this many VLANs or you can create even more than what I have shown. Think about the types of devices you wish to isolate and create the appropriate VLANs. You may find it helpful to consider the function of the devices when creating VLANs. Devices which perform similar functions and need to have similar restrictions/access are good candidates for being on the same network.
For the purposes of this guide, the following VLANs will be created along with some of the reasons why you may want such a VLAN on your network:
| VLAN Tag | VLAN Description | Purpose | 
|---|---|---|
| 10 | DMZ | For anything you want to expose to the Internet | 
| 20 | USER | For PCs, laptops, phones (to isolate from IoT devices) | 
| 30 | IOT | For Internet of Things (IoT) devices to protect other parts of your network | 
| 40 | GUEST | Guest network for visitors/untrusted devices | 
| 50 | IPCAM | Isolated network for IP cameras (for local access only) | 
Create the VLANs by navigating to the “Interfaces > Other Types > VLAN” page. To minimize the length of this guide, repeat the following configuration below for each VLAN in the table above:
| Option | Value | 
|---|---|
| Device | Leave empty to automatically generate a name | 
| Parent | lagg0 (use the LAG interface as the parent for all VLANs) | 
| VLAN tag | Use the values in the table above for each VLAN | 
| VLAN priority | You may use the default “Best Effort” or select priorities (not sure how much it impacts actual performance) | 
| Description | Use the values in the table above for each VLAN (or use your own) | 
Interfaces: Assignments
After the VLANs are created, you will be able to assign them to interfaces. You can think of an “interface” as not only the address of the physical port itself but also an entirely separate network. That concept may seem confusing to new users, but creating a new interface assignment is how you create separate physical or logical networks in OPNsense (and other router platforms). When creating an interface you can specify the size of the network, which limits the total number of devices that can be connected to each network. The interface acts as the gateway for each network where traffic may enter or exit.
On the “Interfaces > Assignments” page, you can create a new interface by clicking on the “+” button in the “New interface” section of the page. The dropdown box only shows unassigned physical/logical interfaces. Once you assign the interface, it will no longer be included in the dropdown.
The WAN and LAN interfaces should already be assigned from the OPNsense installation so I will only mention setting up the VLAN interface assignments.
Select each VLAN listed in the table below in the “Network port” dropdown box and add the appropriate “Description”. The “Description” is displayed on the “Interfaces” section in the left side menu so it is important to use a short name to indicate the purpose of each network. Otherwise, they will show up as “OPT1”, “OPT2”, etc., which will be very confusing when you have multiple networks to manage.
| Network Port | Description | 
|---|---|
| vlan01 DMZ (Parent: lagg0, Tag: 10) | DMZ | 
| vlan02 USER (Parent: lagg0, Tag: 20) | USER | 
| vlan03 IOT (Parent: lagg0, Tag: 30) | IOT | 
| vlan04 GUEST (Parent: lagg0, Tag: 40) | GUEST | 
| vlan05 IPCAM (Parent: lagg0, Tag: 50) | IPCAM | 
Click the “Save” button when you are finished.
Tip
