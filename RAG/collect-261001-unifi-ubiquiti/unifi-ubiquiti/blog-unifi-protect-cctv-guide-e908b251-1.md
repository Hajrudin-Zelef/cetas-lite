---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-unifi-protect-cctv-guide-e908b251-1
title: "blog-unifi-protect-cctv-guide-e908b251"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple", "United States"]
dates: ["2026-08", "2026-08-10"]
keywords: ["cost", "ethernet", "license"]
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-unifi-protect-cctv-guide-e908b251.md
source_anchor: ""
source_lines: [1, 51]
sha256: 7ccde9da95d8b884ec0249d28f8adb8e8714ec31c3376a89568c005ddc2254b2
---

# blog-unifi-protect-cctv-guide-e908b251

UniFi Protect Setup Guide 2026: Cameras, NVRs, Storage and Real Costs
Plan a UniFi Protect camera system with current 2026 NVR and camera choices, complete budgets, storage and PoE sizing, ONVIF support, installation steps and no-license-fee trade-offs.
What's Current in August 2026
- G6 cameras deliver 4K + AI face recognition from $199—and the G6 Pro line extends this further
- No mandatory subscription — all core Protect features run locally with no license fee
- ONVIF support now includes motion events, PTZ control, and audio for compatible third-party cameras
- Protect 7.x is the current release, bringing improved smart detections and archiving options
This guide walks you through building a UniFi Protect surveillance system—from initial planning to ongoing maintenance. Whether you're a home DIYer or a business owner, you'll find practical advice for creating a working security setup without mandatory recurring fees.
Last updated: August 10, 2026. Prices checked against the Ubiquiti US store and exclude tax and shipping.
| Feature | Benefit | Why It Matters | 
|---|---|---|
| No Mandatory Subscription | Core features (recording, live view, AI detections, app) have no recurring fee | Lower total cost of ownership and predictable budgeting | 
| Local Storage | Keeps your video data on your hardware | Enhanced privacy and control over your sensitive footage | 
| Scalable System | Start small and add cameras as needed | Cost-effective growth, the system adapts to your evolving needs | 
| AI-Powered Smart Detection | Reduces false alerts and focuses on relevant events (people, vehicles) | More efficient monitoring and faster response to actual threats | 
| User-Friendly Interface | Easy setup and management, even for non-experts | Less time spent on configuration and more on security | 
| Remote Access & Alerts | Monitor your property from anywhere via the mobile app | Proactive security and peace of mind on the go | 
| ONVIF Compatibility | Integrate existing third-party cameras without additional costs | Protect previous camera investments while upgrading gradually | 
Understanding the UniFi Protect Ecosystem
UniFi Protect is Ubiquiti's video surveillance platform, designed to be robust and user-friendly. It's a comprehensive system beyond basic recording, offering features catering to novice users and experienced IT professionals. Unlike many traditional CCTV systems, UniFi Protect operates without monthly subscription fees. This is possible because recorded footage is stored locally rather than in the cloud. Let's explore the key features and requirements:
- No Mandatory Subscription: There is no recurring license fee for core Protect features—recording, live view, smart detections, and mobile app access are included. Optional paid services (cloud storage tiers, UI Care) exist but are not required.
- Local Storage: Video footage is stored directly on your own hardware, giving you greater control over your data and faster access to recordings.
- AI-Driven Smart Detection: G6 and AI-series cameras distinguish between people, vehicles, and other objects. Advanced models include face recognition (excluding 360 models) and license plate detection, reducing false alerts.
- ONVIF Compatibility: UniFi Protect supports third-party ONVIF cameras with live view, recording, motion events, PTZ control, and audio. Compatibility varies by model. AI smart detections require native UniFi cameras or a compatible ONVIF camera paired with an AI Port.
- Remote Access: The UniFi Protect mobile app (iOS and Android) provides live camera feeds and recording playback from anywhere with an internet connection.
- Scalability: Start with a few cameras and expand as needed—consoles range from the compact UNVR Instant (up to 6 4K cameras) to the ENVR Core (up to 300 4K cameras).
System Requirements
- UniFi OS Console: At the heart of the UniFi Protect system is a UniFi OS Console. This device acts as the central hub, managing your cameras, storing recordings, and providing the interface for accessing the system. Options include the CloudKey+, Dream Machine Pro/SE/Pro Max, UNVR Instant, UNVR, and UNVR Pro—each offering different storage capacity and camera limits.
- Network Infrastructure: A stable network is crucial for optimal performance. Wi-Fi camera models such as the G6 Instant can connect wirelessly; consoles, NVRs, and PoE cameras normally use wired Ethernet. Power over Ethernet (PoE) switches provide both power and data to cameras through a single cable, simplifying installation.
UniFi Protect Overview
Planning Your UniFi CCTV Deployment
Before installing any equipment, careful planning must ensure your UniFi Protect system meets your security needs. This involves a thorough site assessment and selecting the appropriate hardware for your specific environment.
Site Assessment: Laying the Groundwork for Success
- Identify Vulnerable Areas: Begin by identifying areas susceptible to security breaches. Consider potential entry points, such as doors and windows, as well as areas where valuable assets are stored. Don't overlook areas like loading docks, parking lots, and less-trafficked parts of your property.
- Camera Placement Strategy:
  - Field of View: Each UniFi camera model has a specific field of view. Choose cameras with appropriate lenses to cover the desired areas effectively. Wider fields of view are suitable for large spaces, while narrower fields of view are better for focusing on specific points of interest.
  - Height and Angle: Mounting height and angle significantly impact coverage. Generally, cameras should be mounted high enough to prevent tampering but low enough to capture clear facial details. Experiment with different angles to optimize visibility and minimize blind spots.
  - Consider Obstructions: Be mindful of potential obstructions like trees, walls, or signage that could block the camera's view. Plan camera placement to avoid these obstacles or consider adjusting the environment if necessary.
- Lighting Conditions:
  - Day/Night Vision: Most UniFi cameras feature infrared (IR) LEDs for night vision. However, the effectiveness of IR can vary. Evaluate the ambient lighting in each area and consider cameras with strong low-light performance for areas with minimal illumination.
  - Supplemental Lighting: In areas with very poor lighting, consider installing external lighting to improve image quality and enhance the effectiveness of motion detection.
- Power Source Considerations:
  - PoE (Power over Ethernet): PoE is the recommended method for powering UniFi cameras, as it simplifies installation by providing both power and data through a single Ethernet cable.
  - PoE Switches/Injectors: Ensure you have a PoE switch with enough ports and a sufficient power budget for all your cameras. If you don't have a PoE switch, you can use PoE injectors to power individual cameras.
  - Alternative Power Options: While less common, some cameras may be powered by separate AC or DC adapters if PoE is not feasible.
- Environmental factors: What type of weather will these cameras have to endure? Is there a risk of vandalism? Make sure you pick a suitable camera for the environment.
UniFi Design Center Feature: Floor Plans
Using the UniFi Design Center for Optimal Camera Placement
Ubiquiti offers a free online tool called the UniFi Design Center for planning your camera deployment. It lets you upload your floor plan and virtually place UniFi devices, including cameras, to visualize coverage and optimize placement.
Here's how to use the UniFi Design Center for planning your surveillance system:
- Create a New Project or Sign In: If you don't have an account, you'll need to create a new one. After that, you can create a new project.
