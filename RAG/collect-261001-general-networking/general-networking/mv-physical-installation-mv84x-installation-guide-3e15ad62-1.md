---
id: collect-261001-general-networking/general-networking/mv-physical-installation-mv84x-installation-guide-3e15ad62-1
title: "mv-physical-installation-mv84x-installation-guide-3e15ad62"
domain: general-networking
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-general-networking/mv-physical-installation-mv84x-installation-guide-3e15ad62.md
source_anchor: ""
source_lines: [1, 131]
sha256: 76e740c00fc55d9e3f48d40e5943496176ecc8fd008820abd0a2541286d45041
---

# mv-physical-installation-mv84x-installation-guide-3e15ad62

MV84X Installation Guide
This is a comprehensive guide for the MV84X camera, detailing its overview, power requirements, dashboard setup, and step-by-step installation instructions.
MV84X Overview
The Cisco Meraki MV84X is a multi-imager smart camera that is exceptionally simple to deploy and configure due to its integration into the Meraki dashboard and cloud-augmented edge storage. The MV Smart Camera family eliminates complex and expensive servers and video recorders required by traditional solutions, removing the limitations typically placed on video surveillance deployments.
What’s In the Box?
The following items are included with the MV84X.
Verify that all contents are present before beginning installation.
| Item | Specification | Quantity | 
|---|---|---|
| Camera | MV84X | 1 | 
| Camera Ceiling Base Plate | - | 1 | 
| IR Shield | - | 1 | 
| Installation Template | - | 1 | 
| Screw | M6x15 | 3 | 
| Screw | M5x30 | 3 | 
| Anchor | M5x50 | 3 | 
| Dessicant Pack | - | 2 | 
| Rubber Gasket (additional) | 6.3 - 8.3 mm | 1 | 
| Screwdriver | Secure T10 Torx | 1 | 
| Wrench | 26mm C-Socket | 1 | 
Powering the MV84X Camera
The MV84X features a 1000BASE-TX Ethernet port and requires 802.3bt Type 3 PoE++ for operation. Route the Ethernet cable from an active port on a Type 3 PoE++ switch or injector.
- Power over Ethernet supports a maximum cable length of 328 ft (100 m).
- The MV84X camera is capable of powering on with 802.3at (PoE+) when the IR illumination is disabled.
Pre-Install Preparation 
You must complete the following steps before going on-site to perform an installation.
Configure Your Network in the Dashboard
The following is a brief overview of the steps required to add an MV84X camera to your network. For detailed instructions about creating, configuring and managing the Meraki Camera networks, refer to the MV - Smart Cameras documentation.
- Login to http://dashboard.meraki.com. If this is your first time, create a new account.
- Find the network to which you plan to add your cameras or create a new network.
- Add your cameras to your network. You will need your Meraki order number (found on your invoice) or the serial number of each camera, which looks like Qxxx-xxxx-xxxx and is found on the unit label.
- Verify that the camera is now listed under Cameras > Monitor > Cameras.
Check and Configure Firewall Settings
If a firewall is in place, it must allow outgoing connections on particular ports to IP addresses. The most current list of outbound ports and IP addresses for your organization can be found under Help > Firewall Info on the dashboard.
DNS Configuration Best Practices for LAN Streaming
Each camera will generate a unique domain name for secured direct streaming functionality. These domain names resolve an A record for the private IP address of the camera. Any public recursive DNS server will resolve this domain.
If utilizing an on-site DNS server, please allow *.devices.meraki.direct or configure a conditional forwarder so that local domains are not appended to *.devices.meraki.direct and that these domain requests are forwarded to Google public DNS.
Assigning IP Addresses
Currently, the MV84X camera does not support static IP assignment. MV84X units must be added to a subnet that uses DHCP and has a pool of available DHCP addresses to operate correctly.
Mounting Instructions
If you are mounting the MV84X with additional accessories like a Wall Arm, Corner mount, Junction Box, Pole mount or the Parapet mount, please review the MV84X Mounting Options and Guidelines along with this installation document.
Step 1: Loosen the Dome Housing Screws
- 
    Using the provided T10 torx security-bit screwdriver, locate the eight (8) perimeter screws around the MV84X dome housing. 
  - 
        Loosen each screw to about 50%—AVOID removing them completely. This will allow the lifting or removal of the dome cover without losing the screws.
- 
        
- 
    Ensure the screws remain securely in their slots so they don’t fall out during the following steps.
Step 2: Rotate and Remove the Dome Cover
- 
    After partially loosening the eight perimeter screws in Step 1, gently rotate the dome cover in the direction indicated by the arrow.
- 
    The dome will disengage when the arrows on the dome and the base plate aligns.
- 
    Carefully lift the dome cover away from the rest of the housing.
Tip: If the dome doesn’t move freely, ensure each screw is loosened enough to allow rotation, but not too loose that it detaches from the base.
Step 3: Lift the Dome Housing
- 
    After rotating the dome in Step 2, carefully lift the dome housing straight up, as shown by the green arrows.
- 
    The dome is secured to the camera using a cable. Place the dome on the side while keeping the cable attached.
- 
    Verify that the internal assemblies (imagers, lenses, etc) remain secure and unobstructed.
Important: Avoid touching sensitive components inside the camera. Smudges or debris on the lenses can impact video quality.
Step 4: Unlock and Rotate the Camera Base
- 
    On the camera base plate, you’ll see a small button, gently press to unlock the camera from the base plate.
- 
    Turn the camera assembly in the direction of the green arrow to completely disengage the camera from the base plate.
Step 5: Detach the Mounting Plate
- 
    After completing the previous unlock/rotation step (Step 4), ensure the base is in the “unlock” position.
- 
    Separate the mounting plate (the large circular piece) away from the main camera assembly, as indicated by the upward arrows.
- 
    Place the mounting plate aside safely.
Step 6: Remove or Disassemble the Cable Gland
- 
    Locate C-Socket Nut on the camera base, this contains the cable gland assembly.
- 
    Unscrew the C-socket nut from the base of the camera. Use the provided wrench as needed.
- 
    Remove the seal basket (also known as a compression or seal ring) and the rubber gasket.
Keep these pieces safe and organized as they will be used later for reinstallation.
Step 7: Mount the Plate and Route the Cable
- 
    Prepare the Cable Opening 
  - 
        Ensure the hole in the ceiling or wall is large enough for your cable and connectors (e.g., RJ45 plug).
  - 
        If using anchors, insert them into pre-drilled holes in the ceiling or wall.
- 
        
- 
    Feed the Cable 
  - 
        Pull the cable through the hole, lining it up with the center opening of the mounting plate.
  - 
        Confirm that you have enough slack on the camera side to attach or adjust the cable later. Note that at approximately 20cm of cable is needed for proper installation after the camera is locked into its mounting plate.
- 
        
- 
    Align the Mounting Plate 
  - 
        Position the mounting plate against the ceiling or wall so the screw holes line up with the anchors (or pre-drilled holes).
- 
        
- 
    Secure with Screws 
  - 
        Insert the appropriate screws (e.g., M5x30) through the mounting plate into the anchors or backing surface.
- 
        
Step 8: Position the Camera Body
- 
    Hook the support cable on the camera assembly into the built-in clamp on the underside of the mounting plate (highlighted in the inset image).
- 
    Guide the Ethernet cable through the opening in the camera assembly. 
  - 
        Ensure there is enough slack to comfortably reach the camera’s port without stressing the cable.
- 
        
Step 9: Reattach the Camera Assembly
- 
    Gently lift the camera assembly toward the mounting plate. 
  - 
        Keep the network cable centered so it isn’t pinched or twisted.
- 
        
