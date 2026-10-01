---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-29457139215511-ai-lpr-setup-in-unifi-protect-adc5cc9b
title: "hc-en-us-articles-29457139215511-ai-lpr-setup-in-unifi-protect-adc5cc9b"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-29457139215511-ai-lpr-setup-in-unifi-protect-adc5cc9b.md
source_anchor: ""
source_lines: [1, 19]
sha256: 0d4442ec201b112202e1bf0a58443fb456e88a55556cbd65866dbfff5dbabb90
---

# hc-en-us-articles-29457139215511-ai-lpr-setup-in-unifi-protect-adc5cc9b

AI-LPR Setup in UniFi Protect
The UniFi AI-LPR camera is a dedicated license plate recognition (LPR) solution designed to deliver precise and reliable results, day and night. In this guide, we'll show how simple it is to set up and optimize your AI-LPR camera for maximum performance.
If you are interested in learning about LPR in the context of all of UniFi Protect's detections, click here.
Camera Installation
The AI LPR camera should be installed as close to the license plates while also minimizing the angle of capture. This is especially important for reliable captures at night. The IR does not reflect as well off of license plates that are far away or at a severe angle. Select a mounting point as close to the road or above the road if possible to meet these requirements:
- Max. mounting height: 5m
- Capture distance: 12m
- Max. horizontal and vertical camera angle: <25 degrees
- Angle for high-speed capture: <15 degrees
Note: For lower speed or stopped traffic, up to 15m distance and 30 degree angle may work.
Camera Configuration
- Power on your camera: Connect the camera using PoE+ power.
- Set up camera: Adopt it and allow it to automatically update to the latest version.
- Configure LPR Night Vision: Most regions use reflective license plates, in which case you should navigate to UniFi Device > AI LPR > Settings and enable LPR Night Vision. If you do not live in one of these regions, you should leave this setting disabled.
- Optionally configure LPR-Only Mode: Navigate to UniFi Device > AI LPR > Recording Settings and enable LPR Dedicated. This will turn off all other detection types and use all camera resources for license plate recognition.
- Adjust the optical zoom: Navigate to Devices, select a camera and click on its live feed to adjust the zoom so that the camera is as focused as possible on the region where plates will appear.
- Audit and adjust: Review the footage from daytime and nighttime. If detections aren't reliable enough, review your mounting, zoom and night vision settings.
Context Camera
Because the AI-LPR's glare-cut filter detects reflective license plates at night, its visibility is limited to license plates and headlights only. For broader surveillance coverage, consider installing a secondary "context camera" to capture the general area.
