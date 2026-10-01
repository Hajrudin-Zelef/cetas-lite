---
id: collect-261001-general-networking/general-networking/mv-mv-smart-camera-faq-d1c0d599-5
title: "mv-mv-smart-camera-faq-d1c0d599"
domain: general-networking
role: reference
task: reference
actors: ["Qualcomm"]
dates: []
keywords: ["ethernet", "license", "memory"]
source: docs/RAG/collect-261001-general-networking/mv-mv-smart-camera-faq-d1c0d599.md
source_anchor: ""
source_lines: [235, 303]
sha256: eb28e946010945af660a99a7cf026558d09e6d4027a990c3a6632f0f1330499e
---

# mv-mv-smart-camera-faq-d1c0d599

First Year on Us FAQ
Sales Team FAQs
How do I add the offer in CCW?
This promotion will not appear in CCW and will be applied automatically. There is no change from your current deal registration process on any Meraki SKUs in CCW for this promotion.
Is this offer stackable with other Cisco partner promotions?
Yes, this promotion is stackable with all existing Cisco partner programs.
General MV FAQs
Where is video stored?
Video is stored on the camera itself. You also have the option of using Cloud Archive for continuous 24/7 recording for 30, 90, 180 or 365 days. Read the datasheet and this article to learn more.
How much footage can I store?
Retention day estimates based on quality settings can be found in this article. You can also read about different ways to extend your storage on that page.
What if I need longer storage?
You have the option to use Cloud Archive for continuous 24/7 recording, with storage durations of 30, 90, 180, or 365 days. To learn more, please refer to the datasheet and this article.
For customers in Europe or regions with strict storage limitations, the dashboard includes a setting to automatically discard footage after a specified number of days.
Lastly, if the recording options on the dashboard or Cloud Archive do not meet your needs, we also offer External RTSP, which allows footage to be sent to a third-party system. Read more here.
Is my video stored in the cloud?
Standard streaming video is only stored in the cloud when using Cloud Archive. In standard operations, streaming video is not stored in the cloud, but exported video clips are stored in the Meraki cloud ecosystem for a year. Metadata for video thumbnails and motion indexing are also stored in the cloud. Video will proxy through the cloud in the event of remote streaming and is cached in order to retrieve footage more quickly, but during standard recording operations, MV cameras will only save video locally.
Can I choose to backup my video in the cloud?
Yes. You have the option of using Cloud Archive for continuous 24/7 recording for 30, 90 , 180 or 365 days. Due to bandwidth considerations of storing video in the cloud, we recommend only using Cloud Archive on critical cameras. Read the datasheet and this article to learn more.
What resolution do these cameras record in? What is the frame rate?
Tables of resolutions, frame rates, and respective estimated retention days for each MV camera can be found in this article.
What happens if somebody steals my camera? What if it’s damaged or there’s a fire?
As with all security camera systems, care should be taken during installation to ensure cameras are out of reach of potential thieves or vandals. Cameras are rarely stolen and should be positioned out of easy reach. If a camera is stolen, its footage can be retrieved from anywhere when it comes back online thanks to the cloud proxy. This way you can easily retrieve the footage of it being stolen while also watching the live feed of the person using the camera.
Our IP67 rated MV72 should be used in areas where there is a high risk of theft/vandalism. Using motion alerts and/or the snapshot API when reasonable, in conjunction with other access control measures can help mitigate risk. Cloud Archive can be used also provide additional assurance for sensitive or at-risk cameras.
Like most other systems, in the event of a fire, if cameras become damaged to the point of inoperability, footage will be lost. All security camera systems carry some amount of risk, but unlike a deployment with a recording device, there is no single point of vulnerability in an MV deployment since each camera stores its own video.
Is stored video encrypted on the camera?
Yes, MV features full disk encryption. Additionally, cameras automatically purchase and provision their own publicly-signed SSL certificates upon initial boot up. Management data is encrypted as well. All encryption is turned on by default, and cannot be turned off.
Does my SSL certificates renew automatically?
Yes, MV cameras will automatically renew and provision their own publicly-signed SSL certificates, without any manual intervention.
What happens when a camera is removed from a network?
Removing a camera from a network deletes all the footage on the camera.
What if a camera loses access to the network?
MV cameras will continue to record, even if they lose access to the network, as long as they have power.
What will happen if the camera loses power?
Cameras will not operate without power. A switch with redundant power supplies is recommended if there is a concern about power loss.
Is there any way to back video up onto an NVR, DVR, NAS, or other centralized storage device?
The MV is designed to operate without an external recording solution. However, for certain use cases we also have External RTSP, which allows the footage to be sent to a 3rd party system. Read more here.
What if my bandwidth is limited?
The Meraki dashboard will intelligently determine if the viewing computer is in the same local network as the cameras. If this is the case, video traffic will stream directly over the LAN, saving WAN bandwidth. If video is viewed remotely, the dashboard will proxy it through the cloud. It’s recommended that customers with limited WAN bandwidth planning on doing frequent remote viewing should select the Standard Quality video setting.
What CODECs does the MV support?
H.264.
When I export video clips, what format are they exported in?
Video clips are exported as MP4 encoded with H.264.
Can the memory on my camera be upgraded or replaced?
No, the solid state memory on these cameras cannot be removed or replaced.
Are these cameras PTZ capable?
Meraki does not offer a PTZ camera. To see an area from multiple angles, we recommend the use of our MV32 fisheye camera which allows you to see a full 360 degree view of the surrounding area below the camera.
MV21 and MV71 are fixed varifocal dome cameras, so zoom, focus, and aperture settings can be adjusted remotely through the dashboard, but lens angle and positioning must be adjusted during installation.
MV12 and MV32 cameras use a completely fixed lens with autofocus.
Does MV offer analytics capabilities?
MV’s powerful motion search plus motion heat maps capability comes standard with all MV hardware plus a license.
All second generation cameras (MV12, MV22, MV72 and MV32) feature a Qualcomm Snapdragon processor, enabling machine-learning-based computer vision, including detecting people as objects.
With the release of the MV32, Motion Search 2.0 and Motion Recap were announced. These new features enhance the information gathered via motion search, enabling the scene of interest to be observed at a glance.
Are MV cameras NDAA compliant?
All second and third generation MV cameras are NDAA compliant.
Are these cameras wireless?
MV21 and MV71 require a standard Ethernet cable and PoE (for MV21) or PoE+ (for MV71) to operate.
All second generation cameras (MV12, MV22, MV72 and MV32) can be used with standard Ethernet and PoE, but they also feature 802.11ac wireless.
Meraki also has power injectors available for users who do not have a PoE enabled switch available.
Can you zoom in on historical footage?
Video clips are exported in the standard MP4 format, allowing you to use the zoom feature available in most standard video players during playback.
When video is exported from the MV32 in a hemispherical view, the zoom feature from standard video players will not dewarp the image.
What do the LED colors mean?
- Rainbow (solid, rotating through colors) - MV is booting up.
- Flashing Blue - MV is searching for WiFi network(s).
- Flashing Green - MV is upgrading or initializing for the first time.
- Solid Green - MV is connected via Ethernet.
- Solid Blue - MV is connected via WiFi.
- Solid Violet - MV has audio recording enabled.
