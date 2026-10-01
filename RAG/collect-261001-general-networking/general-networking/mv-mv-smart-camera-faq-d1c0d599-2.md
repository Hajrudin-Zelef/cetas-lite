---
id: collect-261001-general-networking/general-networking/mv-mv-smart-camera-faq-d1c0d599-2
title: "mv-mv-smart-camera-faq-d1c0d599"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["license", "licenses", "parameters", "research", "training"]
source: docs/RAG/collect-261001-general-networking/mv-mv-smart-camera-faq-d1c0d599.md
source_anchor: ""
source_lines: [65, 148]
sha256: eacc578490b12e6b7225ef457aad90cc9b6e2660317c30937a4ff5141b64d6d0
---

# mv-mv-smart-camera-faq-d1c0d599

A progressive web app (PWA) can run on a browser and be downloaded to run like a native app. It includes elements such as the service worker and a pre-caching dB.
What is the advantage of a PWA?
- 
    Reduced wait time for page load
- 
    Resilience in network conditions
- 
    Give a smooth and seamless experience in the browser
- 
    Give a native feeling experience with a small footprint when installed
What is the update button used for?
If the "Update" button isn't manually clicked, the app will eventually update to the latest version after a relaunch of the interface (assuming no other Meraki Vision applications are open). The button exists for those who wish for the update as soon as it is available.
How do I ensure my Vision Portal is on the latest release?
Ensure the application is upgraded to the latest version:
- 
    Restart the application to load the latest version
- 
    Command + shift + R on the keyboard should help reload it
- 
    One can check the currently running version by pressing the What's New panel on the top of the UI.
Can Cisco Meraki Support Engineers view my Vision Portal?
By Default, Cisco Meraki Support Engineers cannot view video or hear the audio in the Vision Portal. You may choose to allow temporary access to the Portal by following this document.
What devices will be supported?
No support within the Application Store for iOS or Android.
- 
    iPad’s (CANNOT be installed but can be pinned from Dashboard)
- 
    PC’s
- 
    Laptop
What browsers will be supported?
- The Vision Portal will always open in your default web browser when navigating from Dashboard
- The Web Application will be supported on the latest versions of Chrome / Safari and Firefox (recommended: Chrome). In most cases, compatibility will extend one version behind the most recent release.
I keep getting logged out of the Vision Portal. Help!
Today the Vision Portal uses the Organization Timeout that has been set on the Dashboard. If you have nothing set but you still keep getting logged out, please create a support ticket.
MV Sense FAQs
What is MV Sense?
It is a set of APIs (both RESTful and MQTT-based) that allow our customers to integrate the edge-computing capabilities of the MV into third-party business solutions. It is best not to call it an “open” API because it requires a license.
What is an API?
An API is a set of routines, protocols, and tools that allows applications to communicate with each other. Essentially, an API is a messenger that will deliver requests and responses from the provider you want to talk to. For example, you can request a list of zones from a specific camera from one of the MV Sense APIs, and the API will return this list to you. This greatly simplifies application development requiring information from external sources.
I don’t see the option for “Sense” on my camera’s settings. How do I use MV Sense?
Ensure the camera is a second-generation (MV*2) or later and on MV 3.22 firmware or later.
How does MV Sense work?
Several APIs are available through MV Sense which is outlined in this article. You can obtain three types of data insights through the MV Sense APIs about the people detection metadata generated through the MI edge analytics running on a second-generation Meraki camera. These APIs can be summarized as follows:
- 
    Historical Aggregate: How many people were here at X time?
- 
    Current Snapshot: How many people are here now?
- 
    Real-time Feed: A sub-second feed of people and their locations.
For more information, see the page on MV Sense on the Meraki Developer Portal.
Are there specific applications we recommend to integrate with MV Sense?
No. We provide a suite of APIs that customers can build solutions on top of. For more information, see the page for MV Sense on the Meraki Developer Portal. Some examples of current solutions using MV Sense can be viewed at the Meraki Marketplace. Furthermore, we are working with Cisco DevNet and other partners to develop more turnkey solutions.
How does licensing work?
MV Sense licensing works on a per-camera basis (it does not co-terminate). Each organization (new or old) with second-generation cameras (trials included) should already have ten free MV Sense licenses. These licenses do not expire.
Can I use MV Sense to funnel video to another application?
No. It is a way to leverage the MV’s people detection capabilities for additional analytics, triggers, or other desired customer applications.
MV Cloud Archive FAQs
How does it work?
Refer to this article on MV Cloud Archive. When enabled, the dashboard will automatically pull the footage from the onboard storage (if available) or the cloud archive without any customer intervention.
Can we assume the same secured access and encryption with Cloud Archive as with onboard storage?
Yes. Video uploaded to Cloud Archive is over an SSL/TLS connection. Stored video is also encrypted at rest. All rigorous security standards for the dashboard outlined here apply for Cloud Archive as we take the privacy of customer video very seriously.
Does the cloud archive store the analytics as well?
Yes.
Does this mean I can bulk offload video in private servers or third party software?
No. The interface for the customer will be almost exactly the same as without cloud archive. The user will still be able to do 1-hour video exports, but no bulk offloading.
How does licensing work?
It is available in 30-day, 90-day, 180-day, and 365-day storage license options, with 1-year, 3-year, and 5-year variants, applied on a per-camera basis (licenses do not co-terminate). The licensing is independent of the camera type and remains consistent across all fixed dome, varifocal, and fisheye cameras.
Do we provide licenses for trial?
No.
MV Audio Detection FAQs
How does it work?
Software on the camera processes the input audio stream, using windowed sub-sampling. These samples are then converted into visual representations of the input audio signal and processed by a model trained for specific audio classes. These audio detections can be streamed for detailed analysis and storage via MQTT. This functionality requires a MV Sense license to be applied to the camera. Our audio detection is driven by computer vision and machine learning.
What do you mean by "machine learning"?
Meraki smart cameras use deep learning, a type of machine learning at the forefront of artificial intelligence research, to drive our computer vision object detection. The smart camera development teams continually show a computer thousands of examples of what objects look like and it "learns" how to identify them more and more accurately over time. The model improves as we provide it with additional training data.
The smart camera analytics models are trained on data that is legally owned or licensed by Cisco Meraki, and only uses data from customers who have explicitly opted-in to continually improve our models.
Does this take any extra bandwidth? Do I need another server?
All audio processing is performed directly on the camera's processor. As with all Meraki products, the hardware transmits small amounts of metadata to the dashboard for further processing and storage. Additionally, MQTT messaging data can be sent to any MQTT broker resource of your choice.
MV Encoding Improvements: Smart Codec FAQs
What is a codec?
A codec is a software tool that encodes or decodes a data stream, such as video. Meraki MV smart cameras utilize the H.264 codec. Starting with firmware version MV4.15, MV cameras will begin using Smart Codec.
How do we deliver more video retention?
Smart Codec analyzes a scene in real-time and dynamically optimizes video retention and playback. It observes various levels of motions and other parameters to apply corresponding levels of compression logic to the video files by intelligently determining which scenes require more compression and which scenes do not.
How much can video retention improvement be expected?
