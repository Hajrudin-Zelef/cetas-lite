---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-34338807702679-configuring-the-unifi-retrofit-hub-in-unifi-acc-8f2364ee-2
title: "hc-en-us-articles-34338807702679-configuring-the-unifi-retrofit-hub-in-unifi-acc-8f2364ee"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-34338807702679-configuring-the-unifi-retrofit-hub-in-unifi-acc-8f2364ee.md
source_anchor: ""
source_lines: [135, 158]
sha256: c4708c826cb3b922af2904e036a57daddf1446d1e1327a9cf9d4bb56ee578a38
---

# hc-en-us-articles-34338807702679-configuring-the-unifi-retrofit-hub-in-unifi-acc-8f2364ee

Does the Retrofit Hub support gate access?
Yes. The gate can be unlocked using an NFC card, Touch Pass, and Mobile Unlock (Mobile Button or Mobile Tap). License Plate Unlock is not supported at this time.
Are Retrofit Readers automatically paired when connected to the Retrofit Hub?
Yes. Once physically connected and powered, Retrofit Readers are automatically recognized by the Retrofit Hub upon discovery.
How do I configure the Retrofit Reader connected to the Retrofit Hub?
Navigate to Devices > Readers > UA Retrofit Reader > Settings and configure the reader settings:
- Name
- Direction (Entry/Exit)
- Access Methods: NFC Card, Touch Pass, and Mobile Unlock (Mobile Button or Mobile Tap)
- Greeting: Set a custom message and salutation.
- Advanced: Configure Status Light and Status Sound.
For third-party Wiegand readers, all settings must be configured on the reader’s platform, not through UniFi Access.
How do I pair a Protect camera with the Retrofit Hub?
To monitor door activity with a UniFi Protect camera, navigate to Access application > Devices > Hubs > UA Retrofit Hub 2 > Overview > Locations > select a door > Paired Devices.
Does the Retrofit Hub support daisy-chaining readers?
Yes. You can daisy-chain up to two Retrofit Readers per door. If a second reader is added after the Retrofit Hub and the first reader have already been adopted, you will need to run Discover Reader in Access application > Devices > UA Retrofit Hub 2 > Terminal > Reader > Reader Type > Retrofit Reader to detect and adopt the second reader.
The Retrofit Hub does not currently support daisy-chaining Wiegand readers.
Can I add a backup battery to the Retrofit Hub for power redundancy?
Yes. You can connect a 12V SLA backup battery to the Retrofit PSU 12V (UACC-Retrofit-PSU-12V) to keep the Retrofit Hub running during power outages.
What happens to the registered NFC cards if I change the card format?
If the card format is changed in the Access application or on a third-party Wiegand platform, all NFC cards must be re-registered. A mismatch between the reader and card formats will prevent the reader from recognizing the cards, resulting in door unlock failures.
Can I connect a Retrofit Reader and a third-party Wiegand reader to the same Retrofit Hub?
We don't recommend mixing UniFi Retrofit Readers and third-party Wiegand readers on the same Retrofit Hub because NFC card registrations made on readers that use different card formats will not synchronize properly across readers. This can lead to data mismatches and door unlocking failures.
For the best compatibility and performance, you should typically stick to using all UniFi Readers, or all third-party Wiegand readers with the same card format.
