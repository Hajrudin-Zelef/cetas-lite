---
id: collect-261001-general-networking/general-networking/mv-mv-smart-camera-faq-d1c0d599-3
title: "mv-mv-smart-camera-faq-d1c0d599"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2021-11"]
keywords: ["license", "research", "training"]
source: docs/RAG/collect-261001-general-networking/mv-mv-smart-camera-faq-d1c0d599.md
source_anchor: ""
source_lines: [149, 204]
sha256: 12106bfd185edd961c0ec2722dd90863d69e32fc07b0f31cd1b4a81fd8cc126b
---

# mv-mv-smart-camera-faq-d1c0d599

As the Smart Codec applies various compression algorithms based on the amount of motion detected, the video retention noticed on the cameras will also depend on the motion detected. For most cameras, the results show double the improvement in video retention as of November 2021.
How is this going to affect the video quality?
The Smart Codec will perform various video compressions while maintaining perceivable video quality.
Which MV cameras will leverage the new Smart Codec?
All second-generation MV cameras (MV12, MV2, MV22(x), MV32, MV52, MV72(x)) on MV4.15+ will support Smart Codec.
Firmware MV4.18.1 does not support Smart Codec, however firmware MV4.18 does.
Third-generation MV cameras (MV13, MV33, MV63(x), MV93(x) will use Smart Codec as the default encoding.
Will Gen1 Cameras support Smart Codec?
Gen 1 cameras will not support Smart Codec.
MV Intelligence Training FAQs
How does it work?
Every Meraki camera transmits metadata related to object detection, motion, and light levels. Once a camera is registered into Intelligence Training, software on the Meraki Backend monitors for interesting events from cameras using only this metadata. This software will selectively collect a few frames for each interesting event and send this to our secure cloud storage for automatic annotation and training. Data is encrypted at all stages of the process, and strict access policies are in place across all resources. For information on our datacenters, storage architecture, regulatory compliance, and otherwise, please see our Trust page at https://meraki.cisco.com/trust.
Can I stop sharing data from a camera?
You can choose to stop sharing data from a camera or delete previously shared data at any time.
How do I opt in to Intelligence Training?
You can opt in to Intelligence Training by following the steps outlined here
How do I opt out to Intelligence Training?
If you prefer not to participate, you can opt out and disabling Intelligence Training at any time by following the steps process here
What do you mean by "machine learning"?
Meraki smart cameras use deep learning, a type of machine learning at the forefront of artificial intelligence research, to drive our computer vision object detection. The smart camera development teams continually show a computer thousands of examples of what objects look like, and it "learns" how to identify them more accurately over time. The model improves as we provide it with additional training data.
The smart camera analytics models are trained on data that is legally owned or licensed by Cisco Meraki, and only use data from customers who have explicitly opted-in to continually improve our models.
Does this take any extra bandwidth?
Data collection for MV Intelligence Training will use additional bandwidth beyond the normal camera operations. This process is designed to collect up to 3 GB of data per day. Once a diverse set of images has been gathered, data collection from the cameras will cease.
What is changing about Camera Intelligence Training?
To keep pace with advancements in AI and to continue delivering the highest quality models for our customers, we’re updating our data collection limits. Starting July 18th, the daily limit for MV Intelligence Training data collection will increase from 200MB to 3GB per day. This new limit represents a maximum threshold and is not one we anticipate reaching regularly. Our commitment remains to responsible data use while ensuring our AI solutions evolve to best serve your needs.
Vision Portal: Event Search FAQs
How do I change my timezone?
Contact your network admin to update the network timezone
What if I notice a lot of False positives?
For false positives on people or vehicle filters, note that our camera intelligence modeling improves with each new generation of our cameras. The best performance can be observed on our latest Gen 3 models, including the MV63 and MV93. We are actively working to enhance this experience! (Object Detection Documentation here).
Participate in our intelligent training to help make this even more accurate.
How can you sign up for the Intelligence Training?
You can sign up through the banner that appears on your network or by navigating to Configure -> Intelligence Training.
Third Generation MV Cameras: Overview and Specifications FAQs
How do I know if the MV is using a wireless connection?
The Status LED on every MV model will display a solid Blue LED if the device is connected using a wireless connection and a solid Green LED when using a hardwired connection.
How do I know if the MV is recording Audio?
Any MV that has audio recording enabled will always display a solid Purple LED to indicate it is recording audio, regardless of the connection type used by the device.
MV People Detection FAQs
How does it work?
Software on the camera analyzes images multiple times per second and identifies where people are located. The camera then tracks the location of these people over time to understand when they came, stayed, and left. The camera rolls up its findings and reports them to the dashboard, where you can view the data in a summary form. Our object detection is driven by computer vision and machine learning.
What do you mean by "machine learning"?
Meraki smart cameras use deep learning, a type of machine learning at the forefront of artificial intelligence research, to drive our computer vision object detection. The smart camera development teams continually show a computer thousands of examples of what people look like and it "learns" how to identify them more and more accurately over time. The model improves as we provide it with additional training data.
Right now the detector is trained on data from a small number of cameras. The smart camera development team expects to deploy a lot more of these cameras, and plans to use data from opted-in customers and all of their diverse deployment scenarios to continually make the model better.
Does this do person identification or facial recognition?
No. Meraki smart cameras don't identify or track specific individuals, they just detect them.
Does this take any extra bandwidth? Do I need another server?
No! All of the image processing is done right on the camera's processor. Like all Meraki products, the hardware sends small amounts of metadata back to the dashboard for further processing and storage.
Why am I not getting the numbers I would expect?
Once a camera is physically installed in a proper location and adjusted for the correct FoV, the MV will automatically start to gather analytics data. For more information on installing a MV camera to optimize people detection performance, see the deployment guidelines below. In any deployment, you should use the data comparatively for observing trends and anomalies, as opposed to using it as absolute measurement. Refer to the People Detection Features > Entrances section above for some more explanation.
How can I troubleshoot MV people detection?
There are no public facing logs to troubleshoot the analytics tools. You can use the "Show People” tool when viewing historical footage to get a rough idea of what the MV people detection model is detecting.
MV Object Detection FAQs
How does it work?
Software on the camera analyzes images multiple times per second and identifies where objects are located. The camera then tracks the location of these objects over time to understand when they entered, where they went within view, and when they left. The camera rolls up its findings and reports them to the dashboard, where you can view the data in a summary form. These sub-second detections can also be streamed for detailed analysis and storage via MQTT. This functionality requires a MV Sense license to be applied to the camera. Our object detection is driven by computer vision and machine learning.
What do you mean by "machine learning"?
