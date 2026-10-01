---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/home-network-unifi-protect-review-b6f9dfa0-2
title: "home-network-unifi-protect-review-b6f9dfa0"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple", "Google"]
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-unifi-ubiquiti/home-network-unifi-protect-review-b6f9dfa0.md
source_anchor: ""
source_lines: [57, 108]
sha256: f3af0b32e7beb4ac617f4a25dfa95147455a72cd602112f997eb95cdd3ca0c03
---

# home-network-unifi-protect-review-b6f9dfa0

Standard uses the H.264 codec, which is the most compatible option. It works in all major browser. The enhanced setting uses H.265/HVEC, this codec is indeed more efficient and uses less storage. But it’s only well supported on Apple devices.
So if you only have Apple devices, then you can safely enable the Enhanced encoding, but in other cases, it’s recommended to use the Standard encoding.
Zones and Lines
In the Zones and Lines tab (1), we can add Motion Zones, Smart Zones, Crossing Lines, and Privacy zones. Privacy zones are areas that you want to block (black) out from the recordings, for example, your neighbors garden.
You can create multiple zones for each camera, and fined per zone the sensitivity and what you wan to detect (Smart zones). Now the sensitivity is always a bit of trial and error. In general the further away the object is, the higher you need to set it, in relation to the object size you want to detect.
As you can see in the screenshot above, I have added multiple zones to this camera. The reasons for this is that I will need to use different sensitivity settings depending on how far away the object is, in order capture persons and not birds.
If you click Add Motion Zone you can draw an area in which you want to detect the motion. A zones can only have 8 points, so you are a bit limit to the shape you can draw. If you need more, simply split the zones in two motion zones.
Smart Detection Zones
Smart Detections zones work pretty much the same way as motion zones. The only difference is that you can select what it should detect, person, vehicle, animal, or multiple. Again you will need to play with the sensitivity settings to find what works for your situations or even use multiple zones.
What you can detect als depends on the camera model you have. You can also create Exclusion zones that ignore certain area wrom detection.
Crossing Lines
Lines crossing is another way to triggering alert. It allows you to draw a line and select what kind of detection should trigger an event. Line crossing are a great way to detect if somebody passes and entry gate for example.
Settings
The last tab, Settings, allow you to fine tune the image quality, night vision and audio recording of the camera. Most cameras have built in microphone. It’s always a good idea the enable the Noise reduction to clear up the audio recording. You can also change the sensitivity of the microphone.
For the night vision, by default it’s set to auto, but you can also change this to custom, allowing you to set the lux level when it should switch over. It depends a bit on the ambient light when you want to switch it over.
In night vision mode, the recordings are of course monochrome. But if you have some outdoor lights, then you might want to switch it over later, so you record longer in color mode, allowing you to see the color of a car for example.
Another setting I want to point out is the status light. A lot of people turn the status light off, because they want to make the camera less noticeable. But by leaving it on, people tend to look up to the camera, curious what the light is, giving you a good shot of there face. So I prefer to leave it on.
Under Advanced setting you can change the shutter exposure. If the recordings are a bit to dark or to bright, then you can fine tune it here.
Alarm Manager
With the camera configured and the zones created, the next step is to defined / fine-tune the alerts. Now as always with systems that send notifications / alerts, you should minimize to only the notifications that really matter, otherwise people get notification fatigue, where they start ignoring alerts.
Ubiquiti has expanded the capabilities of the alarm manager quite significantly the last years, so you can completely fine-tune it to your liking.
You can edit existing alarms or create new alarms. For each alarm you can defined when it needs to send the notification, based on the object detect or system event. You can even combine triggers, like a line crossing followed by a motion detection x seconds later.
In the scope section, you can defined the areas (zones/cameras) that are included in the detection. You can select multiple cameras and for each camera the zones you want to use.
And the last step is of course to define the Action. The actions that you can use depend on the systems you have. If you, for example, have a chime installed, you can let it play a sound when somebody crosses a line.
A good tip is to use schedules where possible. For example, during business hours, certain areas don’t need to trigger a motion alert. So using schedules allow you to minimize the noise.
Downloading and Looking back Events
The playback page allows you to view back recordings. At the bottom you have a couple of options to quickly find the correct recording, based on date and detected events.
We can also use this page to download recordings. To do this, first click on Download (1), this allows you to select the timeframe you want to export, using the sliders in the vertical timeline, or by entering the date/time in the lower part of the screen.
When done, click on Download to this Device, and select where you want to store the recording.
Two other great features are the heatmap and grid search. The heatmap button will generate an overlay, showing you where most activity was. The grid search on the other hand, allows you to quickly find all recordings with motion in the select grid(s).
Now when it comes to finding recordings (events), we also have the Find Anything page, which allows you to go through your recordings based on different filters.
The options I have a bit limited, because I don’t have any AI capable cameras, but with the newer generation cameras, you can also filter the recording on things like car color, license plate, persons with backpacks, etc.
This can really help with quickly finding the right events when you have a lot of footage to go through.
Storage Settings
The last setting I want to highlight are the storage settings. Here you can enable Enhanced Retention which will reduce the recording quality after the specified days. This allows you to extend how long you can retain recordings.
The other options is Continues Archiving. You can link another storage account to UniFi protect, like a OneDrive, Google Drive or NAS storage, and let Protect automatically archive footage to your second storage device.
Now I would recommend to use a cloud based (or at least remote) storage option, so in case something happens with your UniFi installation, you can at least still review the latest footage.
You can limit how much storage you want to use for this, as you can see in the screenshot above, I only allocated 1GB, and set it to only archive Person detection events.
Frequently Asked Questions
A couple of answer on some common questions when it comes to Unifi Protect.
Does Unifi Protect work with other cameras?
 Yes, you can now connect third-partry cameras to UniFi Protect. But you are limited to recording only.
How much storage does Unifi Protect has?
The amount of storage depends on the system and hard disk you choose.
Can I run UniFi Protect in Docker
No, UniFi Protect can only be used in combination with one of the supported UniFi Protect NVRs or Cloud Gateways
Can I install Unifi Protect on a Raspberry PI
No, unfortunatelly thas is also not possible.
Conclusion
I am really happy with the way UniFi Protect is going. I started many years ago with UniFi Video and have transitions along to the way to Protect and currently the G5 cameras. The feature set and capabilities keep expanding which is great.
I hope you liked this UniFi Protect Review, if you have any questions, just drop a comment below.
Hi Rudy, could you update this article with the latest features in 2025?
