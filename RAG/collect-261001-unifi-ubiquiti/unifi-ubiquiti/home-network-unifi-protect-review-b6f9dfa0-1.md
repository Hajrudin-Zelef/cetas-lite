---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/home-network-unifi-protect-review-b6f9dfa0-1
title: "home-network-unifi-protect-review-b6f9dfa0"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "fine-tuning", "license"]
source: docs/RAG/collect-261001-unifi-ubiquiti/home-network-unifi-protect-review-b6f9dfa0.md
source_anchor: ""
source_lines: [1, 56]
sha256: 094adaa8ab7726d2457f9431659d00dbcd8d611e0bbb46a828750a52f1432824
---

# home-network-unifi-protect-review-b6f9dfa0

UniFi Protect is the video security system from Ubiquiti. Just like the Unifi Access Point and Network components, are the UniFi Camera and Protect system really easy to use. They integrate seamlessly together, but to be clear, you can use the Unifi Protect also on its own.
If we are talking about Unifi Protect, then we are talking about the security system behind the cameras. Protect is the NVR for recording the video streams, it sends out the motion alerts and allows you to control and manage your camera’s.
In this article
In this review, we are going to take a look at UniFi Protect, how to install it, and how to configure your camera’s.
Unifi Protect Overview
UniFi Protect does a great job of finding the perfect balance between flexibility and ease-of-use. You can choose between a wide variety of cameras, which you can all install with PoE (Power-over-Ethernet) and get it up and running in a couple of minutes.
The biggest advantage of UniFi Protect compared to other systems is that all the recordings are stored locally but can be viewed from anywhere in the world through the Unifi Protect App.
And another important different is that Protect is completely subscription-free, even for the motion-alerts or the cloud access is no subscription needed.
The product range of UniFi Protect has extend significantly the last couple years. When I wrote the first review of Protect, only a few cameras existed. Right now, you can choose from ~30 different cameras, connect different doorbells, and extend the system with sensors, floodlights and siren.
Advanced Features
Also the range of features have expended over time. Where it started with smart detection, the system can now read license plates and recognize faces.
Smart Detections
Motion detections are always hard to get right. Basic motion detection will result in a lot of false positives, due to light reflections or shadows for example. And fine-tuning the motion detection takes a lot of time.
Smart Detection, however, allows cameras to recognize familiar objects, like persons, vehicles, animals or even packages. All current UniFi Protect cameras support smart detections these days. But some camera are of course a bit better in it then others or are capable of detecting more.
You can find a complete overview of the difference here in the UniFi documentation.
License Plate Recognition
As the name says, camera with license plate recognition are capable of reading license plates from cars. This allows you to fully track a car through different camera views or quickly search recordings based on the license plate.
Face Recognition
The newest cameras also support Face Recognition. Now good to know is that the face recognition algoritme is running locally, so no data is send to the cloud for processing.
The feature allows you to track persons through difference cameras feeds, and can also be used to unluck doors when combining it with UniFi Access.
AI Recognition
The AI Recognition helps you with finding the right footage, also called the “Find anything” feature. You can for example search through your footage to find all cars that are red, or find all persons with a backpack.
Again, this logic behind this runs locally, so nothing is send to the cloud, which is great. But the number of cameras that support this is limited, and in most cases you will need to use the AI Key or AI Port to enhance the footage.
Unifi Protect Cameras
As mentioned in the beginning, you can currently choose between roughly 30 difference cameras, from ceiling mounted dome and turrets to full PTZ cameras. Each camera only needs a PoE connection to get it up and running.
Protect works best with UniFi Cameras, but you can also connect ONVIF-compatible third-party cameras. When using third-party cameras, features like motion detection, or PTZ are not supported. But this allows you to transition easily from an existing environment to a full UniFi Protect environment.
Now which camera you should choose really depends on the location you want to place it and the intend usages.
Getting started with Unifi Protect
To use UniFi Protect you first need an NVR or an Protect compatible Cloud Gateway. These days there are quite a few options, but the most commonly used ones are:
|  | Network Video Recorder Instant | Cloud Gateway Max | Network Video Recorder | 
|---|---|---|---|
| Number of cameras | 15 HD or 6 4K | 15 HD or 5 4K | 60 HD or 18 4k | 
| Storage | 1TB – to 24TB | 512 GB – to 2TB | 8TB to 4x 24TB | 
| PoE ports | 6 GbE Ports | No | No | 
| Price excl storage | $199 | $199 | $299 | 
The Cloud Gateway Max is interesting when you want to use UniFi Network as well. You can then use it as gateway (router) as well. The downside is that you will need an PoE capable switch to power your cameras.
If you only want to use UniFi Protect, then the the Network Video Recorder Instant is the best entry option. It gives you plenty of storage options, and comes with a built in PoE switch.
For enterprise grade installation, the Network Video Recorder is the way to go. It comes with 4 drive bays, allowing you to expand the storage capacity to 4x 24TB.
Other good options are the Dream Machine product lines for the larger office installations.
Installing UniFi Protect
Depending on the console you are using, Protect will be installed straight from the beginning or you will need to install the Protection application after you have set up UniFi Network.
You can install the application under Control Plane > Updates. Make sure it’s updated to the latest version as well.
After you have installed Protect, the first thing you need to do is adopt your cameras. In Protect, click on Devices. All cameras connected to your network will be listed. Simply click Adopt to add the camera to Protect.
With the cameras connected we can start with settings up the motion zones, fine-tuning the system, and setting up the alerts.
Configuration is done per camera, so for each camera we defined the recording zones, when to record, in which quality, etc.
Recording Settings
In the device page, click on an camera to open the recording setting. The first step is to decide when the camera should record. There are two settings that influence when the camera starts recording:
- When to Record
- Recording mode
Important here is to keep your storage in mind, how much storage do you have and how long do you want to retain it. Having two cameras record 24/7 on 2K or 4K will require signifying more storage compared to motion only recording.
I prefer to set the cameras to record Always, and the recording mode to Adaptive. This means that the camera will record in high quality when motion is detected, and otherwise in lower quality.
Next we can configure which AI and Audio Events we want to detect, I recommend enabling them all. It will make searching the footage later a lot easier.
But a setting more importantly is the Record Motion Events. Make sure to enable this, because this will highlight motion events in your recordings. And we need to defined the number of seconds before and after the event we want to record.
Now I prefer to set this to at least 2 seconds before and after. More then often I have seen a motion detection kicking in to late, or in sometimes you just want to now what happened before the recording.
The default settings for the Recording Quality are often good enough. If storage is an issue, and quality less important, you can lower the recording quality and/or FPS (Frame Per Second).
Under the Recording Quality you also have an option to set the Encoding. Now standard is perfectly fine for most cases, but if you look at the info description, then the Enhanced option will give you better video quality and it reduces storage size. Sound like a win, right? There is actual an important difference that you need to know.
