---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-29221435686039-unifi-ai-key-setup-and-faqs-6c3b5845
title: "hc-en-us-articles-29221435686039-unifi-ai-key-setup-and-faqs-6c3b5845"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "transcription"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-29221435686039-unifi-ai-key-setup-and-faqs-6c3b5845.md
source_anchor: ""
source_lines: [1, 49]
sha256: 0e8acbfec164fdeea3746530dad3d5cdc85635ad851120493074b52ee5a52e2e
---

# hc-en-us-articles-29221435686039-unifi-ai-key-setup-and-faqs-6c3b5845

UniFi AI Key Setup and FAQs
AI Key is an advanced edge AI appliance designed to elevate your UniFi Protect security system. It enhances detection capabilities with powerful features like NeXT AI natural language search, AI alerts, speech transcription, image enhancement. Setting up the AI Key is quick and straightforward, allowing you to unlock advanced AI-driven insights that transform how you monitor and secure your environment.
For more information on UniFi Protect AI detections and capabilities, and how they are impacted by the AI Key and/or AI Port, see here.
Overview
AI Key works as an edge appliance by receiving smart detections from paired UniFi cameras. Traditionally, Protect cameras are only able to recognize very basic objects such as people and vehicles. AI Key takes these basic smart detections and runs them through a series of advanced AI models to provide powerful new features. While the models used by AI Key provide many useful new features, they might occasionally return inaccurate results. Over time these models will be optimized to improve their accuracy and add additional features.
Setup
- Connect the AI Key using PoE++ power.
- Adopt the AI Key to UniFi Protect. To learn more about adoption, go here.
- Allow several minutes for the initial firmware update to complete.
- Navigate to Settings > Intelligence.
- Select which cameras for each of the AI Key features. Note that AI Key can only process 1,000 detections per hour, so it’s important to be strategic with which cameras are paired.
  - Computer Vision Enhancement: Enables Next AI Search, AI Alerts, and NeXT AI Summary
  - NeXT AI Summary: Automatically generate summaries for cameras with Computer Vision Enhancement enabled. Enabling may reduce the number of analyzed detections per hour.
  - Speech to Text: Create transcripts from speech smart detections.
  - License Plate Recognition: Enable LPR on G4/G5 series cameras
  - Face Recognition: Enable Face Recognition on G4/G5 series cameras.
  - Face Enhancement: Enable manual enhancement of Face Recognition events.
Natural Language Search
Each smart detection processed by AI Key is analyzed and categorized in the background. This enables NeXT AI Natural Language Search in the Find Anything tab. Enter a query into the search bar to quickly locate footage based on a general description. Results are best when searches describe people or vehicles.
AI Enhanced Alarms
AI Key integrates with Alarm Manager to bring AI alerts to UniFi Protect. Configure up to 5 predefined search queries to receive alerts for any detection that matches the query. AI Alerts can be configured using the following steps:
- Navigate to Alarm Manager and create a new alarm.
- Under Trigger, select Objects -> AI Key Advanced
- Enter the desired search query and adjust the match confidence.
- Under scope, select which cameras to apply the query to.
- Under Action configure notifications and web hooks.
For more information about Alarm Manager, click here.
Speech to text
AI Key provides transcription of speech smart detections provided by AI series cameras. Easily read through a conversation instead of listening to the entire clip. Transcriptions can be found on the find anything page by selecting a detection and clicking on the show transcript button.
Face Enhancement
The face enhancement feature on AI Key uses advanced AI algorithms to clarify and refine facial details captured by UniFi Protect cameras. It improves image quality by reducing noise, enhancing resolution, and sharpening key facial features, making it easier to identify individuals in various lighting conditions. Face enhancement is manual action that can be triggered on the Faces page.
NeXT AI Summary
By default, AI Key will provide the option to manually generate event descriptions for each detection processed. This is available by clicking on the NeXT AI Summary button at the top of the screen in the detections popup.
NeXT AI Summaries can also be generated automatically by enabling them in AI Key settings. This will reduce the maximum number of events AI Key can process per hour.
Face Recognition and LPR
AI Key also has the ability to provide Face Recognition and LPR on Protect G4/G5 series cameras. This functionality is similar to AI Port but there will be delays in processing depending on the length of the queue.
Using AI Key with AI Port
AI Port is required to use Protect G3 or ONVIF third party cameras with AI Key. AI Port processes the streams in real time and adds smart detections. These detections are further processed by AI Key to enable more advanced functionality.
Additionally, AI Port is required to receive real time Face and LPR detections G4 and G5 series cameras. See the table below for a detailed comparison:
|  | AI Function | AI/G6 Series | G4/G5 Series | G3/ONVIF | 
| AI Key | Natural Language Search | ✓ | ✓ | ✓ * | 
|  | AI Enhanced Alarms | ✓ | ✓ | ✓ * | 
|  | Advanced AI Summaries | ✓ | ✓ | ✓ * | 
|  | Face Enhancement | ✓ | ✓ | ✓ * | 
|  | Speech to Text | ✓ | ✓ | ✓ * | 
| AI Port | Face Recognition and LPR | Built in to AI Series | ✓ ** | ✓ | 
|  | Smart Detections | Built in to AI Series | Built into G4/G5 Series | ✓ | 
*Requires AI Port
**Async, realtime with AI Port
