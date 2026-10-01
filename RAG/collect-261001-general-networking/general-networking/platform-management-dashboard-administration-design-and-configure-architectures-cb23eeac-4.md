---
id: collect-261001-general-networking/general-networking/platform-management-dashboard-administration-design-and-configure-architectures-cb23eeac-4
title: "platform-management-dashboard-administration-design-and-configure-architectures--cb23eeac"
domain: general-networking
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["attention", "ethernet"]
source: docs/RAG/collect-261001-general-networking/platform-management-dashboard-administration-design-and-configure-architectures--cb23eeac.md
source_anchor: ""
source_lines: [226, 283]
sha256: 7b8adc95cc1c65faa1fac1775b23c2d452c51ad2445dcb5d2b9f0aa5111e3cf6
---

# platform-management-dashboard-administration-design-and-configure-architectures--cb23eeac

Motion alerts send notifications via email when motion is detected either within the frame or a selected region of a camera. These can be extremely useful in monitoring areas where tracking motion is important. Motion alerts can be configured to be sent always, or only during a certain schedule. Scheduling motion alerts is useful when monitoring an area after hours. Motion alerts are disabled by default.
Create an Inventory List for All Cameras
This is multifaceted as it will help with pre-installation setup and deployment as well as help with post-installation documentation. Having something as simple as an excel spreadsheet listing a camera by model number, serial number, name, MAC address, location description, and other details in the PSSRP is a big plus when coordinating efforts with multiple hands. This can easily be done with the CSV export tool within the Meraki Dashboard from a network or organization inventory.
Define Necessary Video Walls for Monitoring Feeds
Video walls can be configured to allow for simultaneous viewing of multiple camera feeds. It may not be possible to determine all video wall needs prior to implementation, as customers may not know what they want or need. However, it is useful to have some basic video wall needs defined, as it is helpful in conveying how the system can operate/what it can do.
Determine which Users Need Alerts
A big advantage of the Meraki MV system is that the camera device (node) functions in relationship to other equipment in the dashboard. This enables alerting/notification of administrators in the event that the system goes offline. In many other systems, the security cameras, video management system, and storage solutions operate independently in silos. If one part of the system fails, there may not be a process to alert or notify administrators of the failure. This results in many administrators not knowing that a camera or storage system is offline until there is a need to pull footage. As part of the deployment, determine which users should be notified of issues, and configure alerts to be sent in the event of operational disruptions.
Dedicated Security or Monitoring Station(s)
Monitoring stations are common in instances that security guards need to watch multiple areas of a facility or campus. The Meraki dashboard can be configured with video walls to view up to 16 simultaneous camera streams at once per browser. System requirements for machines running video walls are available in our Hardware Guidelines for MV Streaming Workstations document.
Implementation and Installation of Meraki MV Cameras
Pre-Install Preparation for MV Cameras
Powering MVs
MV Cameras are powered by Power over Ethernet (PoE) via the ethernet cable. The consumption for MV12 and MV21 is within the 802.3af standard (PoE). The outdoor camera MV71 requires 802.3at (PoE+).
 
Time Synchronization
A vital part of a security camera system is time synchronisation for each camera. This is done automatically through the Meraki dashboard. No local NTP server necessary.
Assigning IP Addresses
Like any other network device, the MV camera requires an IP Address. MV units must be added to a subnet that uses DHCP and has available DHCP addresses to operate correctly. At this time, the MV cameras do not support static IP assignment. Consider using DHCP reservations if you need the IP address to remain constant.
Check and Conﬁgure Firewall Settings
If a ﬁrewall is in place, it MUST allow outgoing connections on particular ports to the Meraki dashboard. The most current list of outbound ports and IP addresses for your particular organization can be found on the Firewall information page.
DNS Configuration
Each MV will generate a unique domain name to allow for secured direct streaming functionality. These domain names resolve an A record for the private IP address of the camera. Any public recursive DNS server will resolve this domain. If using an on site DNS server, please allow *.devices.meraki.direct or configure a conditional forwarder so that local domains are not appended to *.devices.meraki.direct and these domain requests are forwarded to Google public DNS.
Conﬁguring a Network in Dashboard
All dashboard configurations should be done prior to installing any cameras. This allows the camera to associate with the correct organization/network in dashboard and download its configuration.
The following is a brief overview of the steps required to add an MV to your network. For detailed instructions on creating, conﬁguring and managing Meraki Camera networks, refer to the online documentation (https://documentation.meraki.com/MV).
- 
    Login to http://dashboard.meraki.com. If this is your ﬁrst time, create a new account.
- 
    Find the network you want to add your cameras to, or create a new network.
- 
    Add your cameras to your network. You will need your Meraki order number (found on your invoice) or the serial number of each camera, in an “xxxx-xxxx-xxxx” format, located on the bottom of the unit.
- 
    Verify that the camera is now listed under Cameras > Monitor > Cameras.
- 
    For an easy way to view and configure, install the Meraki App on your smartphone or tablet.
Note: During first-time setup, MV cameras will automatically update to the latest stable firmware. Some features may be unavailable until this automatic update is completed. This process may take up to 10 minutes, as it also includes whole-disk encryption. If you view cameras in dashboard during the setup, you may see the error message “Could not reach camera" while attempting autofocus. You will need to wait until the camera finishes upgrading to the latest stable firmware to use this feature. Please do not unplug cameras until they fully complete the upgrade process.
Installing and Configuring MV Cameras
Note: Leave the plastic film cover on all lenses until installation is complete. Take photos of installed cameras for reference. Add a name and physical address to each camera.
Installation Steps
- 
    Depending on your camera model, please follow the installation instructions in the respective Installation Guide.
- 
    Adjust and focus the camera. 
  - 
        Cisco Meraki App
  - 
        Computer (use any operating system and any browser - no plugins needed)
- 
        
- 
    Configure various video settings based on your design principles (see Design section):
- 
    Restrict the video viewing and export permissions for administrators as described here.
Post-Implementation and Troubleshooting
Run Dark Mode
Run dark disables the LED lights on all MVs. This feature is useful in situations where the lights may be annoying, distracting, or overly conspicuous. For example, the LEDs can be disabled to prevent outdoor MVs from drawing attention at night. To enable dark mode, add the tag “run_dark” to the camera. More details can be found in our Dark Mode documentation.
Troubleshooting IR Reflections
All Meraki MVs, with the exception of the MV32, are equipped with infrared (IR) illuminators. These can be turned on in Low light mode and used to provide an illuminated view. If the picture appears too blurry, please follow this Video Quality Troubleshooting guide.
