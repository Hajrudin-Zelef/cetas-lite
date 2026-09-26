---
id: collect-260926-mikrotik/mikrotik/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87-20
title: "Copyright (c) 2016, Andrea Dainese"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87.md
source_anchor: ""
source_lines: [3566, 3785]
sha256: 444373d207ffb350d4fc902e1527c0f4fd62134f388c449dbd453d8dd0aea355
---

# Copyright (c) 2016, Andrea Dainese

Z-index: Used to change the object’s overlay pos ition on the “Topology Canvas.” An object 
with a higher numerically valued z -index will cover an object with a lower 
numerically valued z-indexed.   
Example: The blue object has a z-index of -1 and the orange object’s z-index is 0. 
Orange object is top over blue object. 
Border width: Used to change the object’s border width. 
Border type: Used to change the border style of the object between solid and dashed. 
Border colour: Used to change the colour of the object’s border 
Background colour: Used to change the background colour of the object 
Transparent: Turns off background colour (filling) and makes the 
object transparent. 
Rotate: Used to rotate the object on the topology.  
 
Name: Used to change the object’s name. 
To save the object, press Save (green button).  
 
10.1.7 Lock objects movement 
The “Lock Lab” feature prevents objects from being moved around on the canvas (among other 
things). For more information about this feature, refer to section 7.1.12. 
10.2 Custom design logical topology 
EVE Community includes a feature to upload your own custom topology picture and map nodes 
to it for easy access.

EVE-NG Community Cookbook 
Version 1.11 
Page 144 of 165 © EVE-NG LTD 
10.2.1 Custom design upload  
Before you upload a custom picture in the lab, make sure it is 
in .png or jpg format with resolution 130-150x130-150 pixels. 
TIP: It is best is to create a topology in the MS Visio and after 
convert it to the .png picture format with resolution 140x140. 
 
Step 1: Open “Add an Object” and then “Pictures” from the left sidebar or 
right click on a free area on topology canvas and hit “Add Picture.” 
 
 
 
 
 
 
 
 
Step 2:  Browse your PC for a .png or .jpg file and hit “Add”.  
 
 
 
Once the picture is added to the topology canvas, the sidebar will display a new option: “Logical 
maps” 
 
 
Step 3: Open the “Logical maps” menu item.  
 
 
 
Pictures window management  
 Delete uploaded picture from the lab 
 
Image Map: Map nodes to places in the 
picture 
 
Display uploaded picture. Work with lab and 
custom topology 
 
Zoom/unzoom uploaded custom topology 
 
Makes the window transparent to see  the 
“Topology Canvas” behind it. Clicking again 
returns to the normal view.

EVE-NG Community Cookbook 
Version 1.11 
Page 145 of 165 © EVE-NG LTD 
 
Close “Pictures” window. 
10.2.2 Custom topology mapping 
This feature allows you to map the lab nodes to your custom topology picture.  
Step 1: Open the Image Map window: 
 
Step 2: Select a node, from the dropdown menu, that you want to map to the topology. 
 
Step 3: Move your mouse over a node icon on the “Image Map” and click to map it. The grey 
circle means that the node is mapped. 
 
Step 4: Continue mapping the rest of the nodes. 
 
Step 5: OPTIONAL. You can also add a mapping for a device external to your EVE server in 
order to telnet, VNC, or RDP to it. This way you can open sessions to all your devices (whether 
external or internal) in one place. 
Select from menu:  
And map with node on topology.  
 
Change image map adding protocol, IP and port.

EVE-NG Community Cookbook 
Version 1.11 
Page 146 of 165 © EVE-NG LTD 
 
 
Step 6: Save your mapping and refresh the browser with F5.  
10.2.3 Delete topology or mapping  
To delete a single node mapping, right click on node mapping circle and click  “Delete.”  
 
To delete the entire custom topology, click delete.  
 
10.3 Configuration export feature 
EVE Community includes an export configuration feature that allows you to save  and manage 
configurations in a lab. The "Configuration Export" and “Startup-configs” features will allow you 
to set these saved configurations as startup configs for your nodes when they boot.  
 IMPORTANT NOTE : Before you start us ing the  “Configuration export ” feature, you 
must complete at least one configuration export. 
Nodes will be greyed out withou t the option to enable “Startup-
configs” until you complete at least one configuration export for 
each node. 
 
 
Node boot order:

EVE-NG Community Cookbook 
Version 1.11 
Page 147 of 165 © EVE-NG LTD 
NVRAM: NVRAM is used as writable permanent storage for the startup configuration . During 
the boot process, the node will alway s check NVRAM for a saved configuration. Saving the 
configuration to NVRAM requires a vendor specific command. Cisco: copy run startup (wr), 
Juniper: commit, etc. It is MANDATORY to save a node’s configuration before you can export 
it. 
Exported configuration: A node configuration that has been exported from the node. It can 
be used to backup configurations or to set them as startup-configs.  
Wipe node: Wiping a node will erase the NVRAM (running config) or the t emporary image 
snapshot, depending on the type of node. Upon a successful wipe, the node will boot with the 
factory default configuration or the configuration included in the base image you are using.  If 
you have the “Startup -config” feature enabled for th e node, then it will boot with the chosen 
config set. You must wipe a node after changing certain node template settings like the image 
or startup-config. You also must wipe the node the first time you want to enable the “Startup -
config” feature. 
Factory default configuration: The base configuration that is applied from the manufacturer.  
 
10.3.1 Supported nodes for configuration exports 
Cisco Dynamips all nodes 
Cisco IOL (IOS on Linux) 
Cisco ASA 
Cisco ASAv 
Cisco CSR1000v 
Cisco Nexus 9K 
Cisco Nexus Titanium 
Cisco vIOS L3 
Cisco vIOS L2 
Cisco XRv 
Cisco XRv9K 
Juniper VRR 
Juniper VMX 
Juniper vMX-NG 
JunipervQFX 
JunipervSRX 
Juniper vSRX-NG 
Mikrotik 
PFsense FW 
Timos Alcatel 
vEOS Arista 
10.3.2 Startup config management  
10.3.2.1 Global commands 
Configurations can be managed via the “Startup-configs window which can 
be accessed from the sidebar menu while on the Topology page.

EVE-NG Community Cookbook 
Version 1.11 
Page 148 of 165 © EVE-NG LTD 
 
Topology page, More Options: 
Export all CFGs – Exports all supported node configurations. 
Set nodes startup-cfg to default configset- Sets all 
supported nodes to boot from the default configuration set. 
Set nodes startup-cfg to none - Sets all supported nodes to 
boot from NVRAM configuration. 
Delete default configuration set. Warning, this will delete 
your exported default configuration set for all nodes. 
 
10.3.2.2 Individual node commands 
Select node, right click 
 
Wipe: Wipes the NVRAM for a single node 
Export CFG: Exports the configuration for a single node 
 
10.3.2.3 Multiple selected nodes commands 
 
Wipe Selected: Wipes the NVRAM for selected nodes 
Export all CFGs: Exports the configuration for selected nodes 
Set nodes startup-cfg to default configs set: Set selected nodes to the default config set 
Set nodes startup -cfg to none : Set nodes to boot from NVRAM or from factory default if 
wiped.

EVE-NG Community Cookbook 
Version 1.11 
Page 149 of 165 © EVE-NG LTD 
Delete nodes startup cfg: Delete selected node’s startup cfg. (clean default set) 
10.3.2.4 Startup-configuration window 
No configuration exports or manual configs loaded for nodes 
 
Startup-configs are exported and the “Configuration Export” feature can be used. 
 
10.3.2.5 Startup-config window information 
 Config set menu 
 No configuration is available for node. Grey 
node 
 
Configuration is available and can be used. 
Blue node. Exported configuration persist 
 Configuration persist but it is disabled. Node 
will boot from NVRAM or factory default if it 
is wiped  
 Configuration persists and node will boot 
from the configuration after being wiped 
 Ace Editor. Different vendor configuration 
edit option. Just Text visual format.

EVE-NG Community Cookbook 
Version 1.11 
Page 150 of 165 © EVE-NG LTD 
10.3.3 Export configuration 
Example: 
 
