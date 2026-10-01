---
id: collect-261001-general-networking/general-networking/mv-physical-installation-mv84x-installation-guide-3e15ad62-2
title: "mv-physical-installation-mv84x-installation-guide-3e15ad62"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["alignment", "ethernet"]
source: docs/RAG/collect-261001-general-networking/mv-physical-installation-mv84x-installation-guide-3e15ad62.md
source_anchor: ""
source_lines: [132, 246]
sha256: 25bad2f565715efd4efb3d7b255558d90c098502bb2cf15b199d275dbebbb673
---

# mv-physical-installation-mv84x-installation-guide-3e15ad62

- 
    Twist the camera assembly in the direction of “TO LOCK” marker to the mounting plate and align the arrow marks to lock the parts together.
- 
    Ensure the plate and camera assembly seat flush against each other and against the ceiling.
- 
    Ensure you have approximately 20cm of the network cable dropping down from the camera assembly so that it reaches the ethernet port.
- 
    Double-check the cable’s path to ensure it is free of sharp edges and not compressed between the plate and the ceiling.
Step 10: Reassemble the Cable Gland
- 
    Gather the top cap, compression ring, and bottom insert (the parts you removed earlier in Step 6).
- 
    One by one, guide the components up the cable in the correct order: 
  - 
        The bottom rubber gasket.
  - 
        The seal basket.
  - 
        The threaded top cap.
- 
        
- 
    Reassemble the pieces together as shown in the image.
- 
    Push the reassembled gland pieces into the camera’s port or housing opening (where the cable emerges).
- 
    Screw the top cap into place using the wrench provided by rotating it clockwise.
Important: When using the wrench, make sure it does not hit the lens or imager.
Step 11: Power On and Verify Network Connection
- 
    Align the RJ45 connector with the camera’s Ethernet port. 
  - 
        Push it in firmly until you hear or feel the latch click into place.
- 
        
- 
    Ensure the power source is Type 3 PoE++ (802.3bt).
- 
    Look for an LED indicator to confirm the camera is powered.
- 
    Wait for the camera to show online on the Meraki dashboard and ensure the video feed is viewable.
Step 12: Adjust the Camera Angle and Orientation
- 
    Once the live view is available, the lenses can now be correctly aimed as needed by following these steps: 
  - 
        Move the camera imager vertically to capture more ceiling or floor as needed.
  - 
        Rotate the lens horizontally on the 360 track to aim the lenses at the areas of interest.
  - 
        Adjust the lens assembly to straighten the image if the scene appears tilted.
  - 
        Use the Meraki dashboard to Optical Zoom each imager if needed, if you want a narrower field of view.
- 
        
- 
    Use the Meraki dashboard to verify the camera’s angle is correct.
- 
    Make fine-tune adjustments until you have the optimal field of view.
After the camera is online for a few minutes, it will begin a firmware upgrade sequence. If this occurs during aiming, wait a few minutes for the video feed to appear after the firmware upgrade.
Step 13: Replace the Desiccant Packets
- 
    The MV84X ships with pre-installed desiccant packs to ensure the camera is moisture-free during transit and installation, these packs must be removed and replaced with a fresh set of desiccant packs provided in the box.
- 
    Remove the two desiccant packs and identify the designated spots inside the camera assembly where the desiccant packets should be placed.
- 
    Find the two new desiccant packs and remove it from its silver colored protective pouch. 
  - 
        Handle the packets carefully; avoid tearing or puncturing them.
- 
        
- 
    Place two desiccant packets into the recommended slots or areas, as indicated by the arrows in the diagram.
Tip: Desiccant helps mitigate condensation, especially in environments with varying temperatures or high humidity. Ensure that you replace them.
Step 14: Reattach the Dome Cover
- 
    Hold the dome cover (the large outer shell) and position it so the alignment markers match up with the camera assembly.
- 
    You will see small arrows on the dome and camera base indicating the correct orientation.
- 
    Ensure the dome edges are flush with the camera’s base, no gaps should be visible.
- 
    Identify the eight (8) perimeter screws around the dome cover (the same ones loosened in earlier steps).
- 
    Using the T10 torx security-bit screwdriver, turn each screw clockwise until it’s snug.
- 
    Tighten them in a star or alternating pattern to ensure even pressure (e.g., top-left, then bottom-right, etc.). Avoid Over-Torquing.
- 
    Stop once you feel moderate resistance as over-tightening can strip threads.
- 
    Once fully tightened, the dome should be firmly in place, with no visible gaps.
Step 15: Attach the IR Shield
- 
    Hold the Infrared (IR) shield below the camera dome, orienting it so the cutout matches the dome’s contours (Cisco logo).
- 
    Slide or push the shield onto the dome until it seats snugly around the camera housing. 
  - 
        You should hear four clicks from the corner of the IR Shield.
- 
        
- 
    The shield should fit flush against the dome edges, with minimal gaps for optimal IR reflection.
- 
    If it doesn’t sit evenly, remove and realign.
Tip: The shield directs infrared light correctly at night, and proper installation reduces glare or reflections on the dome.
MV84X LED Status Indicator
Your MV84X is equipped with an LED light on the front of the unit to convey system functionality and performance information.
The following colours and patterns indicate the various status conditions of an MV:
- 
    Rainbow (solid, rotating through colours) - MV is booting up.
- 
    Flashing Green - MV is upgrading or initialising for the first time.
- 
    Solid Green - MV is connected via Ethernet.
- 
    Solid Amber - MV is not able to connect to the Dashboard.
