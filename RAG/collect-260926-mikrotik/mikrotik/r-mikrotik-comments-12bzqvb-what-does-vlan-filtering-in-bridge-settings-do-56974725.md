---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-12bzqvb-what-does-vlan-filtering-in-bridge-settings-do-56974725
title: "r-mikrotik-comments-12bzqvb-what-does-vlan-filtering-in-bridge-settings-do-56974725"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/r-mikrotik-comments-12bzqvb-what-does-vlan-filtering-in-bridge-settings-do-56974725.md
source_anchor: ""
source_lines: [1, 76]
sha256: 290e6b348a7f6771b41f327c65b55b7e0e20db167c162cdcbb6288f9986b15f9
---

# r-mikrotik-comments-12bzqvb-what-does-vlan-filtering-in-bridge-settings-do-56974725

what does "VLAN Filtering" in bridge settings do? 
        
    I've been struggling to configure my new Internet gateway for multiple VLANs. I'm using Mikrotik hEX S and I want to use the built-in switch as a switch and have several VLANs for different purposes.
Most if not all tutorials that I've found end their configuration by enabling VLAN Filtering under bridge settings. However, I can't get my configuration working with this enabled and my current setup that I got working uses a single bridge with multiple VLANs, VLAN interfaces matching the VLANs, but has VLAN filtering disabled.
What am I missing here? What does this feature do and why is it recommended so much? Note that I don't need any retagging, Q-in-Q, filters based on VLAN IDs or anything like this. I just want to use the integrated switch as a switch with few VLANs, access/hybrid/trunk ports and inter-VLAN routing.
Section des commentaires
I will say, if you just create tagged VLAN interfaces on your bridge, it all works without doing any other configuration. If all you want is adding tagged VLANS on every port.
I want a mix of access ports with no VLAN tagging, hybrid ports (trunk with native VLAN in Cisco lingo) and trunks. It seems that I can get this working without VLAN filtering enabled so I am confused why it's recommended so much and if it's still a relevant config option today.
It's very relevant on switches where the VLAN stuff can happen in hardware. Not so much on a router where the VLANs aren't being switched anyway and most of them don't have switch chips that support VLAN-aware switching even if you wanted to do hardware offloading.
I made a little guide based on my Routerboard3011. You mentioned that you use a Cisco Product. I also have a Cisco switch connected to my SFP port.
Step 1 - Create bridge & add ports
If no bridge is created go to
/interface bridgeand create a new bridge
Add ports to newly created bridge
/interface bridge ports
Step 2 - Add VLAN Interfaces
Go to
/interface vlan
Create a new VLAN for each VLAN ID and select the created bridge as interface
Set MTU to 1508
Step 3 - Add Address List
Go to
/ip address list
Add desired IP and select the corresponding VLAN Interface which was created in step 2
Step 4 - Create DHCP Servers for VLANs
Go to
/ip dhcp serverclick onDHCP Setup
Create DHCP Server for each VLAN interface
Step 5 - Add VLAN IDs to bridge
Go to
/interface bridge vlan
Select correct bridge (if multiple bridges available)
Add VLAN IDs created in
/interface vlan
Add
tagged/trunkports. At minimum thebridge itselfmust be selected
Add
untaggedports if a device is directly connected to Routerboard (ex. Printer)
Step 6 - Change PVID for untagged ports (Devices directly connected to Routerboard)
Go to
/interface bridge ports
Select ether port
Go to tab VLAN
Enter PVID of the VLAN you want to assign
Change
Frame Typestoadmit only untagged and priority tagged
Step 7 - Enable VLAN filtering on bridge
Go to
/interface bridge
Select bridge
In the opened bridge menu go to VLAN and enable
VLAN Filtering
AFAIK depending on the MikroTik model there is an optimal method to set up VLAN to achieve near-wire-speed
Can you explain why this step is required? I know it is, just never understood why.
Adding the bridge itself to the VLAN makes the CPU receive the frames. If not, only the switch deal with them (pure L2): An example of that, having an CRS326 not doing routing on a VLAN, an external router doing it.
So if you want routing on a VLAN, you must add the bridge to itself. In a router supporting L3 switching (RB5009 for example, it has a Marvell 88E6393 switch chip), the CPU still have to receive at least the SYN packets in frames: it acts like an SDN controller and will tell the switch chip what to do with the new flow. Once the switch chip got its answer from CPU, it won't bother it after. The switch chip populates its flow forwarding table using information from the CPU/SDN Controller. It was a 10 000 feet view of the principle!
The bridge needs to be tagged on a vlan if you want to be able to reach/manage the router via the IP address on that vlans interface.
Definitely important on any CRS for maximum performance, it basically enforces your vlan mapping and sends everything hardware offload through the switch chip.
Exactly what it says it does. It enabled ‘VLAN filtering ‘ on all ports on that bridge and enables the vlan tab on the bridge to do its thing.
Watch this.
Video 0003 - Mikrotik VLANS and how to connect 3 of them on a home or SMB network https://youtu.be/YI0gPoCQDmM
Si solo quieres hacer inter-vlan-routing la configuración va así.
Crea un bridge. Crea todas tus interfaces VLAN (SVI) si quieres que tengan comunicación en capa 3.
Luego de eso tienes que decidir cuales de los puertos serán troncales o de acceso.
¿Ya lo pensaste bien? Entonces mete todos los puertos al bridge. (Al menos deja un puerto fuera del bridge para que no te quedes sin acceso al Mikrotik).
Ahora tienes que entrar al apartado de bridge. En cada uno de los puertos tienes que configurarlos como tagged o como untagged. En tus puertos troncales todas las VLANs deberían estar como tagged. En los puertos de acceso debes configurarlos como untagged y ademas debes cambiar el PVID.
Una vez tienes todo esto hecho ya puedes habilitar el bridge-vlan-filtering.
Tengo un tutorial espero te sirva --> http://foroisp.com/threads/1936-Bridge-en-Mikrotik-Uso-correcto-en-los-Cloud-Core-Switch-serie-300
Commentaire supprimé par le membre
Forgetting to include the bridge made for a frustrating couple of evenings
To this day is still don't understand why it's required and not automatic.
My current configuration seems to be working even without this enabled so I am wondering if I am missing something of if this is a config option that was relevant in the past but is not required anymore. I am testing my configuration using another router that sends frames in multiple VLANs with DHCP in those VLANs so I can confirm things actually work and VLAN tags are processed on the new router that I am setting up.
This is probably why nothing worked with VLAN filtering enabled as this configuration is not at all intuitive to me. I may give it another try using this style of config but I am still not convinced that I gain anything from it if VLANs are processed without it anyway.
If you want to use untagged vlan ports, then you can’t do this in hardware without using vlan filtering.
Yeah. The bridge needs to be tagged for your tagged VLANs and untagged for the native VLAN (1). That's how the CPU is able to connect to the bridge and assign addresses, and make forwarding decisions.
The frame type acceptance policy is a really important bit. A lot of examples out there will use “accept only vlan tagged”, but if you’re using hybrid mode (mixed native and tagged) then accept all is the only option.
