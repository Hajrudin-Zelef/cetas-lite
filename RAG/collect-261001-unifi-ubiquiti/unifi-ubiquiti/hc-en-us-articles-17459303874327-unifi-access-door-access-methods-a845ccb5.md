---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-17459303874327-unifi-access-door-access-methods-a845ccb5
title: "hc-en-us-articles-17459303874327-unifi-access-door-access-methods-a845ccb5"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-17459303874327-unifi-access-door-access-methods-a845ccb5.md
source_anchor: ""
source_lines: [1, 56]
sha256: 9e8e818566716b16a3e63399d625626d997b78efdfd812e80eefccc7e65d87f1
---

# hc-en-us-articles-17459303874327-unifi-access-door-access-methods-a845ccb5

Configuring Door Unlock Methods in UniFi Access
UniFi Access offers a variety of door unlock methods to deliver a seamless experience for you and your users. This article provides an overview of each method with configuration links. To learn more about creating access policies and schedules, click here. To learn more about selecting the best UniFi Access Readers for your deployment, click here.
| Update your UniFi OS, as well as your Access and Protect applications and devices, to the latest versions for the newest features and optimal performance. | 
Supported Door Unlock Methods
By default, all supported unlock methods are enabled for a given reader. These settings can be configured in Access application > Devices > select a reader > Settings.
Below is an overview of the available unlock methods, along with a link to a detailed configuration guide for assigning these methods to users.
- NFC Card: Tap your NFC credential on a reader.
- 
Mobile Unlock:
  - Mobile Button: Tap a button in your UniFi Endpoint for mobile.
  - Mobile Tap: Tap your iOS or Android phone against a reader.
- Face Unlock: Scan your face at a reader.
- Touch Pass: Tap your phone or watch, with Touch Pass added, on a reader.
- PIN: Enter your PIN on a reader.
- QR Code: (For visitors only) Scan your QR code at a reader.
- Fingerprint: Scan your fingerprint at a reader.
- Hand Wave: (For exit only) Wave your hand in front of a reader. For security purposes, it cannot be used with other access methods simultaneously.
- License Plate Unlock: Position your vehicle within the detection zone of the nearby UniFi surveillance camera.
Configuring 2-Step Authentication
Version Requirements
- Access application 3.3.22 or later
- UniFi OS 4.1.22 or later
- UniFi Access Readers
  - G3 Intercom (UA-G3-Intercom) 1.8.19.0 or later
  - Intercom (UA-Intercom) 1.8.19.0 or later
  - G3 Reader Pro (UA-G3-Pro) 1.11.19.0 or later
  - G2 Reader Pro (UA-G2-Pro) 1.11.20.0 or later
- UniFi Access Control Hubs
  - Door Hub (UA-Hub-Door) 4.9.14.0 or later
  - Enterprise Access Hub (EAH-8) 5.7.17.0 or later
  - Elevator Hub (UA-Hub-Elevator) 4.9.14.0 or later
  - Gate Hub (UA-Hub-Gate) 2.3.15.0 or later
  - Door Hub Mini (UA-Hub-Door-Mini) 1.3.4.0 or later
Enable 2-Step Authentication in Access application > Devices > select a reader > Settings to enhance door unlock security using one of the following methods.
- Require Two Methods: The user must authenticate with any two enabled methods (e.g., NFC card and PIN) within 10, 20, or 30 seconds to unlock the door.
- Require Approver: The user authenticates first, then a designated approver (user or user group) must also authenticate within 10, 20, or 30 seconds to unlock the door.
UniFi Device Requirements
|  | G6 Pro Entry G6 Entry1 |              G3 Intercom Intercom  |              G3 Reader Pro G2 Reader Pro  |              G3 Reader G2 Reader  | G3 Reader Fingerprint | Reader Flex | Access Ultra | Reader Lite | Retrofit Reader Fingerprint2 | Retrofit Reader2 | Protect AI Camera3 | 
|  |  |  |  |  |  |  |  |  |  |  |  | 
| NFC Card | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | — | 
| Mobile Unlock | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | — | 
| Face Unlock | ✓ | ✓ | ✓ | — | — | — | — | — | — | — | — | 
| Touch Pass | ✓ | ✓ (G3 Only) | ✓ (G3 Only) | ✓ (G3 Only) | ✓ | ✓ | — | — | ✓ | ✓ | — | 
| PIN | ✓ (G6 Pro Entry Only) | ✓ | ✓ | — | ✓ | ✓ | — | — | ✓ | — | — | 
| QR Code | ✓ (G6 Pro Entry Only) | ✓ | ✓ | — | — | — | — | — | — | — | — | 
| Fingerprint | — | — | — | — | ✓ | — | — | — | ✓ | — | — | 
| Hand Wave | — | — | — | ✓ | — | — | — | ✓ | — | — | — | 
| License Plate Unlock | ✓ | — | — | — | — | — | — | — | — | — | ✓ | 
- To enable door unlock functions, G6 Entry must be adopted in the Protect application and directly connected to Access Control Hubs.
- Retrofit Reader must be used with the UniFi Retrofit Hub.
- See here for the supported AI Camera models.
- See UI Store to learn more about Access Readers.
FAQs
Does 2-step authentication apply to Admin Mode on a reader with a built-in camera?
No. 2-step authentication is enforced for standard door unlock workflows, but is not required to enter Admin Mode on readers with a built-in camera, such as the G3 Intercom and G3 Reader Pro.
Access admins can enter Admin Mode by tapping the Admin icon at the top of the reader screen and authenticating with their assigned NFC card or PIN, regardless of whether 2-step authentication is enabled on the reader.
