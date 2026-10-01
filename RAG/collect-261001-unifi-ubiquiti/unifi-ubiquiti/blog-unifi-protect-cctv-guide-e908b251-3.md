---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-unifi-protect-cctv-guide-e908b251-3
title: "blog-unifi-protect-cctv-guide-e908b251"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["ethernet", "license"]
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-unifi-protect-cctv-guide-e908b251.md
source_anchor: ""
source_lines: [96, 157]
sha256: 5daaf3bd44edef0c67dcc04b63c14bbe897fe2e4f640742e72a7811b35620e5b
---

# blog-unifi-protect-cctv-guide-e908b251

- G4 Doorbell Pro: A video doorbell with dual cameras — a primary camera for visitors and a secondary downward-facing camera for package monitoring. Includes two-way audio and integrates with the UniFi Protect system like any other camera.
- G4 PTZ: A 4K Pan-Tilt-Zoom camera with 22× optical zoom, long-range IR night vision, motion tracking, and auto-patrol. Suitable for applications requiring variable framing from a single position.
AI Port - Transform Legacy Cameras
The AI Port at $199 brings AI capabilities to older UniFi cameras (G3/G4/G5) and third-party ONVIF cameras:
- Camera Capacity: Supports up to 5 Protect cameras (depending on resolution) or 1 4K / 2 2K / 3 HD ONVIF cameras
- Face Recognition & LPR: Adds facial recognition and license plate detection to connected cameras
- ONVIF Camera Support: Integrates third-party cameras with AI analytics
- Edge Recording: MicroSD slot for local backup recording (supported now)
- PTZ Control: Supports PTZ on compatible ONVIF cameras
- Power Output: When powered via PoE+, provides up to 12.95 W to the connected camera. PoE++ input is required for up to 25.5 W output
Storage Considerations
The amount of storage you need depends on the number of cameras, recording resolution, frame rate, and how long you want to retain recordings. UniFi Protect includes storage management features that allow you to optimize usage effectively.
- Hard Drive Selection: Surveillance-grade drives (Seagate SkyHawk, Western Digital Purple) are designed for continuous write workloads and are generally preferable for NVR use. Verify that any drive is compatible with the recorder and expected workload before purchasing.
- Storage Policies: UniFi Protect allows different retention periods for high-quality versus low-quality recordings, maximizing storage efficiency while preserving critical footage longer. For detailed planning, see our storage planning guide.
Step-by-Step Installation
With your equipment selected, it's time to begin the installation process. This involves physically mounting and connecting your cameras and configuring the UniFi Protect software.
Physical Installation
Mounting Cameras:
- Tools Required: Gather the necessary tools, typically including a drill, drill bits (appropriate for your mounting surface), a screwdriver, a level, a pencil or marker, and possibly wall anchors.
- Step-by-Step Mounting Instructions: Each camera model will have specific mounting instructions, usually found in the box or available online. Generally, this involves:
  - Positioning the Mount: Use the included mounting template (if provided) to mark the screw hole locations on the wall or ceiling.
  - Drilling Holes: Drill pilot holes at the marked locations. If mounting into drywall or masonry, use wall anchors.
  - Securing the Mount: Attach the camera mount to the surface using appropriate screws. Ensure it's firmly secured and level.
  - Attaching the Camera: Connect the camera to the mount, following the camera-specific instructions.
- Weatherproofing: Ensure all connections and openings are correctly sealed for outdoor installations to prevent water damage. Use weatherproof enclosures or sealant if necessary.
Pro Tip
When mounting G6 Turrets outdoors, apply dielectric grease on the Ethernet connection to prevent corrosion—even if you use the weather-sealing gland. It's a $5 insurance policy for your $199 camera.
Cabling:
- Running Ethernet Cables: Run Ethernet cables from your PoE switch or injector to each camera location. Consider using conduits or raceways to protect cables and maintain a neat appearance, especially in exposed areas.
- Connecting to PoE Switch/Injector: Connect one end of each Ethernet cable to a camera and the other to a corresponding port on your PoE switch or injector. Ensure the connections are secure.
Software Configuration
Setting up the UniFi OS Console:
- Initial Setup: Connect your UniFi OS Console to your network (and power it if it doesn't support PoE) and follow the initial setup instructions. This typically involves accessing the console's web interface through a browser and following an on-screen wizard.
- UniFi OS Interface: Familiarize yourself with the UniFi OS interface. This is where you'll manage all your UniFi devices, including your cameras.
Adopting Cameras:
- Discovery Process: Once powered on and connected to the network, your cameras should be automatically discovered by the UniFi Protect application within your UniFi OS Console.
- Adoption Process: The UniFi Protect interface shows a list of discovered cameras. Click the "Adopt" button next to each camera to add it to your system.
Basic Camera Settings:
- Naming Cameras: Assign descriptive names to each camera (e.g., "Front Door," "Backyard," "Warehouse Entrance") for easy identification.
- Setting Time Zone and Date: Ensure the correct time zone and date are configured on your UniFi OS Console for accurate timestamping of recordings.
Optimizing Camera Settings for Performance and Storage Efficiency
Once your cameras are installed and adopted, it's important to fine-tune their settings to achieve optimal performance and manage storage effectively.
Video Quality and Storage
- Resolution: Higher resolutions provide more detail but consume more storage space. For general surveillance, 2K is often sufficient, while critical areas (entrances, cash registers, license plates) benefit from 4K. The G6 series makes 4K accessible at $199 per camera.
- Frame Rate (FPS): A higher frame rate results in smoother video and increases storage usage. For most surveillance scenarios, 15-20 FPS is a good balance between fluidity and storage efficiency. You might choose a higher FPS for areas with fast-moving objects.
- Bitrate: This setting determines the amount of data used to encode each second of video. Higher bitrates generally result in better image quality but consume more storage. UniFi Protect typically offers automatic bitrate adjustment, which is suitable for most cases.
Motion Detection and Recording
- Motion Zones: Define specific areas where motion detection should be active within the camera's field of view. This helps to reduce false alerts triggered by irrelevant movements, such as swaying trees or passing cars outside your property.
- Sensitivity: Adjust the sensitivity of the motion detection to fine-tune its responsiveness. Higher sensitivity detects smaller movements but may increase false alerts. Finding the right balance depends on the specific environment and camera placement.
- Schedules: Schedules are configured separately and control when recording is active.
- Smart Detections: Utilize UniFi Protect's AI-powered smart detection features (available on G6, AI series cameras, and legacy cameras with AI Port) to identify specific objects like people, vehicles, or even faces and license plates.
Storage Management
- Recording Modes: UniFi Protect offers three recording modes per camera:
  - Always: Continuous 24/7 recording for complete coverage.
  - Detection Only: Records only when a smart detection or motion event is triggered, reducing storage use.
  - Never: Disables recording for specific cameras (useful for privacy or testing).
  - Schedules are configured separately and control when recording is active.
- Retention Policies: Define how long recordings are stored before automatic deletion. You can set different retention policies for high-quality versus low-quality recordings.
- Cloud Archiving: UniFi Protect supports archiving recordings to Google Drive, OneDrive, Dropbox, and NAS destinations for off-site redundancy.
ONVIF Integration and Third-Party Camera Support
UniFi Protect supports ONVIF (Open Network Video Interface Forum) cameras, allowing integration of third-party cameras from other manufacturers.
ONVIF Camera Integration
