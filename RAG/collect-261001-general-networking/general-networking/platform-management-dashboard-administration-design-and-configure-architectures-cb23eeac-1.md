---
id: collect-261001-general-networking/general-networking/platform-management-dashboard-administration-design-and-configure-architectures-cb23eeac-1
title: "platform-management-dashboard-administration-design-and-configure-architectures--cb23eeac"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-general-networking/platform-management-dashboard-administration-design-and-configure-architectures--cb23eeac.md
source_anchor: ""
source_lines: [1, 151]
sha256: 94c44e25a265dec63a6e36df02446daea97931ff9fec57bc9c51b682013d620a
---

# platform-management-dashboard-administration-design-and-configure-architectures--cb23eeac

Designing Meraki MV Smart Camera Solutions
There are a number of ways to design an IP surveillance system. The most important part of the design is identifying areas of security concern and positioning cameras to cover those areas. There are a variety of ways to design camera coverage for the same building. Too many designs are primarily built with the cost of installation in mind, instead of looking at the system as an investment in protecting assets and/or people. A good IP surveillance system should be focused on protecting people and assets. The first steps to building such a system are analyzing the building or facility and conducting a site survey.
Site Survey
Conducting a site survey helps provide an understanding of the security needs of a building/facility, and determines the requirements to address those needs. At the conclusion of a site survey, there should be a clear understanding of what needs to be monitored, the materials/parts needed, labor required, total number and locations of cameras to be installed, and an estimated cost for deploying the solution.
Pre-site Survey Requirements
The list below is a starting point to identifying deployment needs, and will help ensure a better site survey outcome:
- 
    Acquire to-scale floor plans of the building(s) 
  - 
        A purpose of the site survey is to determine the mounting locations of cameras to be installed. Having floor plans helps the site administrator and the installer both fully understand the intent of the design and requirements. If outdoor camera coverage is also required, try to obtain external building plans as well.
- 
        
- 
    Locate all network equipment closets 
  - 
        During the site survey it is necessary to understand existing network equipment, as the cameras will most likely be powered by and connected to the network. Identifying these locations beforehand is necessary.
- 
        
- 
    Ask the simple questions: 
  - 
        Why is an IP video surveillance system needed?
  - 
        What purpose does the system serve?
  - 
        What are the requirements the system MUST have and why MUST the deployment have them?
  - 
        When does the project need to be done?
  - 
        Where is camera coverage required?
  - 
        Who and/or what needs to be protected?
  - 
        Are there any restrictions around mounting cameras to walls and/or ceilings?
  - 
        Are there any restrictions around running CAT5e/6/7 cabling throughout the building/facility?
- 
        
- 
    Get keys and access to ALL parts of the building/facility 
  - 
        This may be difficult as some companies/entities have high levels of security. Be sure to coordinate with management/security/facilities, and explain the purpose of and requirements for survey.
- 
        
- 
    Schedule and treat the site survey like a project 
  - 
        Coordination and communication is key to making the site survey a success. Identifying things in the pre-site survey meeting can save a lot of time during the site survey.
- 
        
- 
    Understand and know the County/State/Local regulations around cameras/surveillance for each facility to be surveyed
Conducting a Site Survey
The site survey determines where to place the cameras. It may also uncover additional suggestions or recommendations that were not initially considered.
Site Survey Checklist
- 
    Have a system to take extensive notes and make recommendations
- 
    Have several copies of the floor plans handy for marking where cameras should be installed
- 
    Take lots of pictures! Pictures can help convey design a lot more easily and are extremely helpful for documentation.
- 
    Consider the following to determine placement for each camera: 
  - 
        Take into consideration camera position and areas of high contrast - bright natural light and shaded darker areas.
  - 
        It is HIGHLY recommended to have at least two (2) vantage points on each ingress and egress point. Having multiple cameras covering the same area is a GOOD thing, as it creates redundancy for backup.
  - 
        Consider the following areas as examples: 
    - 
            Entrances, exits, loading, and/or delivery entrances to the building(s)
    - 
            Any gateway(s) or parking areas/lots where employees/guests park vehicles
    - 
            Cashier, ATM, and other locations handling monetary transactions
    - 
            Highly used areas such as lunch rooms, reception areas, break rooms, waiting areas, hallways, etc.
    - 
            Perimeter coverage of building, including walkways, fences, patio, and outdoor accessible areas used by employees, guests, and others
    - 
            High value items and areas where security and visibility are a necessity
  - 
            
  - 
        The closer a camera is positioned with a narrow field of view, the easier things are to detect and recognize. General purpose coverage provides overall views.
  - 
        Will camera placement require a man lift for installation? 
    - 
            How difficult and high is the mount? This may affect future maintenance.
  - 
            
  - 
        What mounting equipment is required for each camera and how will they be mounted? 
    - 
            Height of installation
    - 
            Distance to targets
    - 
            Lighting
    - 
            Angle of placement
    - 
            Mounting style and mounts (ceiling vs. wall)
  - 
            
  - 
        Determine if new cable runs are required and the distance of the cable run to the nearest networking closet
  - 
        If vandalism of cameras is a concern, position the IP66 rated vandal resistant cameras offered in the Meraki portfolio
- 
        
- 
    When it comes to audio recording be aware of local laws
- 
    Determine if existing networking equipment is adequate for new security cameras 
  - 
        Are there enough ports on the switch(es)?
  - 
        Are additional network switches required?
  - 
        Are these ports PoE and does the equipment have adequate PoE budget to compensate for security cameras?
  - 
        Is the UPS adequate enough to handle the additional power draw and meet company policy standards for run time when down?
  - 
        Are power injectors required?
  - 
        Create network diagrams and document internet connections
  - 
        What is the WAN design and LAN uplink connections and are they adequate to support the necessary streaming?
- 
        
- 
    Are there dedicated computers for video surveillance? 
  - 
        What are the specifications of these computers and do they meet the minimum requirements to run a dedicated video wall?
  - 
        Which employees have access now and who will require access to the new security camera system?
- 
        
- 
    Other notable considerations: 
  - 
        Use tools such as IPVM.com as resources for determining proper placement for cameras and illustrating the outcome of an installation could look like
  - 
        Consider using a live camera-on-a-stick deployment, similar to how some active wireless site surveys might be conducted. Connect the MV camera to the PoE port on an MX67C for example, which can provide both PoE+ power and have cellular WAN for Internet connectivity (you also need a working SIM card for the MX.) This kind of setup can also be part of a mobile cart with a UPS system. 
    - 
            Note: This is very time consuming but provides an exact idea of what camera footage will look like after installation. Consider using an MV22 and/or MV72 for example, as those models have varifocal lenses and will allow you to simulate different deployment scenarios and fields of view including what you may see through a fixed-lens camera.
  - 
            
- 
        
