---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-unifi-protect-cctv-guide-e908b251-2
title: "blog-unifi-protect-cctv-guide-e908b251"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "latency", "license", "pricing"]
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-unifi-protect-cctv-guide-e908b251.md
source_anchor: ""
source_lines: [52, 95]
sha256: c91654a89c9e6542f3c2be3e741f07e5d66f0c4c9fae2bf0b8cf9cdde187f8d0
---

# blog-unifi-protect-cctv-guide-e908b251

- Upload Your Floor Plan: You can upload a floor plan image (PNG, JPG, or SVG) or create one from scratch using the tool's built-in design features. Ensure the scale of the floor plan is accurately set for realistic simulations.
- Select Your UniFi OS Console: Choose the console that will be the foundation of your system.
- Add Cameras to the Floor Plan: Select the UniFi Protect cameras you intend to use from the device list and drag and drop them onto your floor plan. You can adjust their position, rotation, and height to simulate real-world placement.
- Visualize Coverage: The Design Center will display the estimated coverage area for each camera, considering its field of view and mounting height. This helps you identify potential blind spots and ensure adequate coverage of critical areas.
- Experiment and Optimize: Try different camera positions and models to find the most effective layout. The Design Center allows you to easily move devices, change their settings, and see the real-time impact on coverage.
Choosing the Right Equipment
After completing a thorough site assessment, you can now select the appropriate UniFi Protect equipment for your specific needs. This includes choosing the right UniFi OS Console to manage the system and the best camera models for each location.
UniFi Protect Consoles
The UniFi OS Console is the central management device for your UniFi Protect system. It handles video recording, user access, and overall system configuration. Here are the primary options:
- UniFi CloudKey+ SSD (UCK-G2-SSD): A compact console suitable for smaller installations, supporting approximately 8 4K, 14 2K, or 24 HD cameras. The current $249 model includes a preinstalled 1TB SSD. (The older HDD variant at $199 is being phased out.) A good starting point for home users who don't need a full gateway.
- UniFi Dream Machine Pro (UDM-Pro): An all-in-one device combining NVR with a full-featured network gateway/controller. Supports approximately 8 4K, 14 2K, or 24 HD cameras (same general capacity range as the CloudKey+). Does not include built-in PoE — a separate PoE switch is required. Ideal for small to medium-sized businesses at $379.
- UniFi Dream Machine Special Edition (UDM-SE): Similar to the UDM-Pro, but includes built-in PoE ports and an integrated 128 GB SSD for faster UniFi OS responsiveness. Priced around $499.
- UniFi Dream Machine Pro Max (UDM-Pro-Max): The highest-capacity model in the UDM Pro family, with built-in Protect support. Designed for larger deployments with 200+ UniFi devices and 2,000+ clients, it supports up to 5 Gbps routing with IDS/IPS. Features Shadow Mode (VRRP) for high-availability failover. The built-in 128 GB SSD and dual 3.5" NVR HDD bays with RAID data protection make it a strong all-in-one choice. Priced at $599.
- UniFi Network Video Recorder (UNVR / UNVR Pro): Dedicated NVRs for larger deployments. The UNVR ($299) holds four 2.5/3.5" drives and supports up to 18 4K or 60 Full HD cameras. The UNVR Pro ($499) holds seven drives and supports up to 24 4K or 70 Full HD cameras. Both require a separate gateway for network routing. For a detailed comparison including the Gen 2 models, see our UNVR vs UNVR Pro vs UNVR G2 comparison.
- UniFi NVR Instant: The most affordable dedicated NVR at $199. Features an integrated 6-port PoE switch, HDMI output, and a single 3.5-inch drive bay. Supports up to 6 4K, 8 2K, or 15 Full HD cameras. Requires connection to a gateway or Layer 3 switch — the UNVR Instant is not itself a router.
- UniFi Enterprise NVR (ENVR): Enterprise-scale recorder with 16 drive bays supporting up to 70 4K or 210 Full HD cameras. Priced around $1,999. The newer ENVR Core ($4,999) extends capacity to 300 4K or 500 Full HD cameras with hot-swappable power supplies and optional storage expansion units.
UniFi Camera Options
UniFi offers a diverse range of cameras to suit various surveillance needs. Here's a breakdown of the current models:
G6 Series (Current Generation)
The G6 series brings 4K AI capabilities to an accessible price point. For detailed comparisons, see our G6 buying guide:
- G6 Bullet: Outdoor camera with a 1/1.8" 8MP sensor delivering 4K resolution at $199. Includes AI face recognition and license plate detection, IR illumination up to 30 m, and IP66 weather resistance. Available in black or white.
- G6 Turret: Same sensor and AI capabilities as the G6 Bullet in a compact turret form factor for $199. Suited to locations where a lower-profile camera is preferred. 3-axis manual adjustment for flexible installation.
- G6 Instant: Wi-Fi camera at $179 with 4K resolution and AI capabilities for locations without Ethernet access. Uses UniFi Auto-Link for setup and shares the same 1/1.8" 8MP sensor as the wired G6 models.
- G6 Dome: An ultra-robust dome camera at $279 with IK10 vandal-proof rating and 4K 8MP imaging, ideal for high-traffic areas requiring tamper resistance.
- G6 180: Combines two 1/1.8" 8MP sensors at $299 for a stitched 16MP, 180-degree ultra-wide image — useful for covering large areas with fewer cameras.
- G6 PTZ: An upgraded PTZ camera at $399 featuring dual cameras, 4K video capability, 10× hybrid zoom, integrated LED spotlight, speaker for two-way audio, and a MicroSD card slot for edge recording.
- G6 Pro Series: The G6 Pro Bullet, Pro Turret, and Pro Dome extend the G6 line with larger sensors, enhanced low-light performance, and optical zoom for more demanding environments. See the G6 buying guide for full specifications and pricing.
AI Series (Premium Line)
- AI Pro: 4K resolution with enhanced low-light performance from a large 1/1.8" sensor, 3× optical zoom, and AI features including facial recognition and license plate detection.
- AI Turret: An all-weather, vandal-proof 4K PoE+ turret camera with enhanced AI capabilities, long-range night vision, and a large image sensor with IR and visible LEDs.
- AI Dome: All-weather, vandal-proof 4K PoE dome camera with enhanced AI capabilities and long-range IR night vision.
- AI 360: This camera provides a wide, 360-degree view using a fisheye lens and de-warping software, ideal for covering large rooms or open spaces with AI-powered motion detection.
- AI Theta: A unique, discreet camera system with a separate lens and body connected by cable, supporting different lens options, including 360-degree coverage for specialized installations.
- AI PTZ Industrial: Built to withstand extreme environments, this $1,299 PTZ camera combines a hardened enclosure with 4K imaging and advanced AI features like person, vehicle, and animal detection.
- AI PTZ Precision: Industrial-grade 4K PTZ camera ($1,999) with 31× optical zoom, LiDAR-assisted autofocus, and 100 m adaptive IR night vision.
G5 Series (Value-Focused)
- G5 Bullet: This versatile camera is suitable for indoor and outdoor use. Its 5MP sensor provides sharp and detailed images, and its infrared LEDs ensure clear visibility even in low-light conditions.
- G5 Dome: Designed for discreet indoor or outdoor installations. It features a low-profile design and a 5MP sensor, similar to the G5 Bullet.
- G5 Flex: A compact and versatile 2K (4MP) PoE camera suitable for indoor or sheltered outdoor use. Flexible table or ceiling mounting. Currently $129.
- G5 Pro: Enhanced image quality with a large 1/1.8" sensor, 4K resolution at 30 FPS, 3× powered optical zoom, and long-range adaptive IR LEDs.
- G5 Turret Ultra: Ultra-compact, tamper-resistant, and weatherproof 2K HD PoE camera with long-range night vision. Currently $129.
- G5 Dome Ultra: This ultra-compact and tamper-resistant 2K HD PoE camera with night vision is designed for low-profile indoor security.
- G5 PTZ: This compact, all-weather camera features ultra-low latency pan-tilt-zoom control and versatile mounting options.
G4 Series (Professional Grade)
