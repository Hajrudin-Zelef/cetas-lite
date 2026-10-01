---
id: collect-261001-general-networking/general-networking/mv-mv-smart-camera-faq-d1c0d599-4
title: "mv-mv-smart-camera-faq-d1c0d599"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2020-06"]
keywords: ["license", "research", "training"]
source: docs/RAG/collect-261001-general-networking/mv-mv-smart-camera-faq-d1c0d599.md
source_anchor: ""
source_lines: [205, 234]
sha256: fc6c31dc0db5c2c2ad72fc50815053a9b9deb6db7c7309bf24caaa35da5352a5
---

# mv-mv-smart-camera-faq-d1c0d599

Meraki smart cameras use deep learning, a type of machine learning at the forefront of artificial intelligence research, to drive our computer vision object detection. The smart camera development teams continually show a computer thousands of examples of what objects look like and it "learns" how to identify them more and more accurately over time. The model improves as we provide it with additional training data.
The smart camera analytics models are trained on data that is legally owned or licensed by Cisco Meraki, and only uses data from customers who have explicitly opted-in to continually improve our models.
Does this do person identification or facial recognition?
No! Meraki smart cameras do not identify or track specific individuals. Persons are detected as objects (if the analytics model supports it) and tracked in the scene until the object leaves the camera's field of view.
Does this take any extra bandwidth? Do I need another server?
No! All of the image processing is done right on the camera's processor. Like all Meraki products, the hardware sends small amounts of metadata back to the dashboard for further processing and storage.
Why am I not getting the numbers I would expect?
Once a camera is physically installed in a proper location and adjusted for the correct FoV, the MV will automatically start to gather analytics data. For more information on installing an MV camera to optimize object detection performance, see the deployment guidelines below. In any deployment, you should use the data comparatively for observing trends and anomalies, as opposed to using it as an absolute measurement. Refer to the Object Detection Features > Entrances section above for some more explanation.
How can I troubleshoot MV object detection?
New in June 2020, you can now view the detailed output of the running object detection model. This new capability is instrumental to advanced troubleshooting and debugging of applications consuming this data via either MQTT or Dashboard API calls.
An MV Sense license is required to be applied to a MV camera to access the advanced analytics debug mode.
In order to access the advanced analytics debug mode, navigate to the "Show People” tool when viewing historical footage. You should now see an additional toggle for enabling/disabling the debug mode overlay. The debug mode will render each detected object's Object ID #, confidence %, and bounding coordinates (X0,Y0),(X1,Y1).
If there are no detections on the camera and no MQTT output for a camera, please ensure the camera is receiving the correct PoE power as stated in its datasheet. To see if you're hitting this issue, search the Event Log for Event Type "PoE power error. Incorrect PoE standard detected." which will be logged when the camera boots.
External RTSP
Will 3rd Parties Be Able to Store Video?
Once the RTSP stream leaves our ecosystem, the 3rd party system can handle the video however it would like.
What Type of Stream Delay Can Be Expected?
Since RTSP is not using HTTP Live Streaming (HLS), which is what is used on the dashboard, the delay will be significantly reduced. Enabling RTSP will not affect the dashboard delay for video streaming.
How Do I Secure the RTSP Stream?
While the External RTSP stream itself may not be secured or encrypted, there are various configurations that can secure access to the cameras. Some examples of these configurations are:
- Placing the cameras on a separate VLAN
- Configuring port isolation
- ACLs on the switches or edge firewall to allow/deny traffic to the cameras
Configuring Camera FOVs on a Floor Plan FAQs
My camera network is already set up with cameras placed on our floorplan. How will this update affect my cameras?
- A: If your cameras have already been placed on a floorplan, they will have their default FOV cones represented facing South:
 
 
 You can simply rotate them to the appropriate coverage angle and save your configuration changes.
 
