---
id: collect-261001-general-networking/general-networking/mv-mv-smart-camera-faq-d1c0d599-1
title: "mv-mv-smart-camera-faq-d1c0d599"
domain: general-networking
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["training"]
source: docs/RAG/collect-261001-general-networking/mv-mv-smart-camera-faq-d1c0d599.md
source_anchor: ""
source_lines: [1, 64]
sha256: e4a03c2829994a63af152b7c3de4a991c63a9690428a2acf1128f0889ce254b1
---

# mv-mv-smart-camera-faq-d1c0d599

MV Smart Camera FAQ
Overview
This is a summary of the frequently asked questions on MV smart cameras. It is divided into the following sections:
MV52 (Bullet) Camera FAQs
How do I factory reset an MV52?
The factory reset button for the MV52 is located beneath the top hood of the camera. To access it, slide the hood away from the front lens in the direction of the mounting plate. Next, slide/remove the longer cover located under the hood. Once removed, you will find a pinhole that allows you to perform a factory reset on the camera.
It may require significant force to slide the plastic hood away from the camera body.
What firmware should the MV52 be on?
The MV52 must operate on the most recent Beta Firmware version.
My MV52 is out of focus
Drag the focus motor slider to its maximum position at the far end, then adjust it to the near end. After completing this, initiate the AutoFocus function for the entire frame. If the issue persists, contact the support team for further assistance.
My MV52 is extremely hot. Is this expected?
The MV52 has a high processing power as it tries to support all the latest intelligence features, 4K, HDR, and even long-range IR. Due to this, the thermal temperature on the camera is always high. Ensure safety while deploying it and keep it away from flammable materials even while deployed.
Does my MV52 connect wirelessly to the Access Point?
The MV52 supports 5GHz wireless connectivity exclusively and requires a 12V DC PoE adapter for operation.
My MV52 abruptly went offline? My MV52 seems to lose connectivity only during the night?
MV52 is powered by an 802.3at PoE+ connection with 30W of power wholly dedicated to its port. Certain non-enterprise switches display a maximum of 60W of power within their datasheet but share this load among other connected devices. Such situations can disallow the MV52 from coming online, especially during the night for IR consumption.
While running an export for 1 hour on my MV52, it tends to not record for a certain period?
Due to the processing limitations of the MV52, exports can become demanding, especially when a significant amount of motion is captured in each frame. To ensure smooth processing, consider breaking the export into smaller parts.
My MV52 video streaming is very unstable and choppy?
Adding a single MV52 video stream at 4K consumes approximately 8 Mbps for one tile on the video wall. Adding four tiles would require around 32 Mbps of WAN bandwidth for the video wall to operate.
- Ensure your uplink supports high-bitrate streaming, and verify that both your browser and drivers are up to date.
- The latest generation hardware must be used for 4K viewing.
IR reflection
To prevent the issues described below, ensure the camera is mounted in a way that avoids streetlights reflecting off the camera's IR LEDs or the presence of reflective surfaces, such as positioning it in front of a glass windowpane, etc.
MV32 (Fisheye) Camera FAQs
The Meraki MV32 is a fisheye camera equipped with the same edge analytics capabilities built into other second-generation cameras (MV12, MV22, and MV72) and stocked full of many new and exciting features. To get details on the camera’s hardware and resolution specifications, check out the datasheet.
Is the MV32 a Pan, Tilt, Zoom (PTZ) Camera?
The MV32 is a fisheye camera with a 180-degree view of the surrounding area. The term fisheye comes from the camera’s use of a fisheye lens which creates a wide-angle hemispherical image. Using digital zoom, the user can virtually/digitally pan, tilt, or zoom (DPTZ) around the scene without losing any information from areas they are currently not viewing. In the case of PTZ cameras, the only area being recorded is presently in view.
Do I require any special equipment or apps to view the video?
To view the video, log in to the dashboard and access the camera from your camera management page. Fisheye and DPTZ modes can be viewed without any special equipment. Virtual Reality (VR) mode is officially supported on standalone VR headsets, such as the Oculus Go, and is unofficially supported on mobile browsers when used with a mobile VR headset, such as Google Cardboard. To use VR mode, access the dashboard from the device and follow the steps outlined in this article.
What storage capacity does the MV32 have?
The camera has 256 GB of storage. 128 GB storage options are not available.
What video retention period estimations can I expect?
The MV32 runs a high-resolution video stream, boasting more than twice the recording resolution of the MV12. With 256 GB of storage, we can expect up to 20 days of video storage depending on the video quality when recording 24/7. To obtain more effective retention rates, enable motion-based retention or cloud archive.
Can I place the MV32 outdoors?
The MV32 has no weather-resistance rating and should be kept in a temperature and humidity-controlled indoor environments. The MV72, an IP67 weather-resistant varifocal camera, can be used for scenarios requiring a smart security camera outdoors.
Does the MV32 support wireless configuration?
Like other second-generation cameras, the MV32 can run through a wireless connection once it is configured and supplied with power. See this article for more information.
Does the MV32 have night vision?
The MV32 does not have night vision. It contains no built-in IR illuminators for very dark scenes. However, it has an IR-cut filter that allows it to be more sensitive to lower-light scenes.
Can I view multiple angles from a single MV32 in a video wall?
Yes! You can. read more here.
Can I mount the MV32 on the wall?
If placed on a wall, half of the camera’s field of view will face the perpendicular ceiling. This does not properly use of the 360-degree field of view of this camera. The MV32 is better installed on a ceiling, where the more of the field of view captures useful content. Additionally, the digital pan-tilt-zoom (DPTZ) feature of the MV32 is currently configured for ceiling mounting.
Furthermore, the camera’s computer vision analytics function on training algorithms based on a top-down view of its surroundings. Orientations differing from this model will limit the performance of the camera’s analytics. Therefore, it is not recommended to install the MV32 perpendicular to the ground.
What happens if I place the camera on a surface higher than the recommended 10-14 ft?
Mounting the MV32 higher than the recommended maximum height of 14 feet will result in decreased image quality of the surrounding area and impact the performance of the built-in people detection analytics. Placement at this height will enable an overview of larger areas but impact the amount of recognizable detail in the image. If your surface is very high, you can use the Telescoping Pendant Mount to lower the mounting height of the camera.
Can I replace all my cameras with a single MV32?
The MV32 is a great option to deploy if you desire a broad overview of the scene and to gain situational awareness. In general, we recommend utilizing the fisheye camera to detect an event of interest, and a fixed lens or varifocal camera to identify key details of the scene.
My MV32 gets very hot, is this normal?
It is normal for the camera to get warm in operation. Ensure the camera is properly deployed with the included mount plate, which is an important part of cooling the camera.
Vision Portal FAQs
What additional features does Vision Portal have over Dashboard?
- You can zoom into every tile by clicking on the enlarge icon on the Video wall on Vision
- You can create Multi-network video walls on Vision
How can the Vision Portal be accessed?
- 
    By clicking on the Camera Tab -> Monitor -> Vision Portal
- 
    Visiting vision.meraki.com
- 
    Progressive Web Application (once downloaded)
What is a progressive web application?
