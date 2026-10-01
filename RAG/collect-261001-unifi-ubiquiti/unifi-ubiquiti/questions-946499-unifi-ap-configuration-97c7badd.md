---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-946499-unifi-ap-configuration-97c7badd
title: "questions-946499-unifi-ap-configuration-97c7badd"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-946499-unifi-ap-configuration-97c7badd.md
source_anchor: ""
source_lines: [1, 13]
sha256: 77764d69c7730f5338838ef1a65b94eb206ae4ffd7e24673ffe0b27ff6812649
---

# questions-946499-unifi-ap-configuration-97c7badd

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
So I have deployed a number of Unifi AP's at my client sites and they are working very well. I configure the AP's at my office using the Unifi software and then take out to the client site and install them.
I have to make a change to one of these units out in the field and I didn't want to have to bring the unit back to my office to do this. Is there anyway to make changes to a Unifi AP w/o the software configuration tool?
If you set them up correctly, you simply make the change from your office system, with no need to visit the site at all, since the controller can be remote from the APs, by design.
If you didn't set them up correctly, you'll be bringing it back to the office. And this might motivate you to set up a proper, reachable-from-client-locations controller instance either at your office or in the cloud; or not, depending on you. The details of how to do that are findable at the Ubiquiti UniFi forum.
All changes to UniFi APs are made on the controller and communicated from the controller to the APs. There is no backdoor, that's the way the system is designed to work.
Even nowadays you can assign a controller to your ubnt account and in the controller settings under 'Cloud Access'.
Afterwards you can access all online controllers assigned to your account directly from https://unifi.ubnt.com just signing in with your ubnt account and making changes to any device like you have it in the office over the GUI.
