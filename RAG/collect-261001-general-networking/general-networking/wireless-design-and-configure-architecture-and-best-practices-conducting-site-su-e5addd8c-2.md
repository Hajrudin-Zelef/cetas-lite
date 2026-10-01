---
id: collect-261001-general-networking/general-networking/wireless-design-and-configure-architecture-and-best-practices-conducting-site-su-e5addd8c-2
title: "wireless-design-and-configure-architecture-and-best-practices-conducting-site-su-e5addd8c"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-general-networking/wireless-design-and-configure-architecture-and-best-practices-conducting-site-su-e5addd8c.md
source_anchor: ""
source_lines: [53, 66]
sha256: 07f29d6fc59b63c427b70e0eee3e5447119c73b0f86ab31dcd25713e28c17c0a
---

# wireless-design-and-configure-architecture-and-best-practices-conducting-site-su-e5addd8c

Site Survey Best Practices
- 
    Site surveys should be performed using the same Cisco Meraki access point model as that would be used in the customer infrastructure.
- 
    Ensure that the AP is not mounted close to any metal or concrete walls that can contribute to heavy attenuation of RF signals.
- 
    Have a blueprint of the location being surveyed handy, to document the signal readings, data rates and record any interference sources during the survey.
- With the proliferation of clients with varying wireless capabilities, it is important to survey for the ‘worst’ clients in order to ensure a consistent experience across all your clients once your wireless network is in production.
Additional Survey Options and Considerations
With so many variables that can affect the RF propagation of an AP and in turn affect the client performance and roaming behaviors, it is always recommended to have an onsite site survey performed to ensure optimal coverage and performance. Those who do not have access to a site survey tool can still leverage the statistics provided on the local status page on the access point. Admins can see signal-to-noise ratio for the client that is connected to the survey SSID. A classic approach is to design for a voice-grade network which uses cell edges around -67dBm and a consistent SNR (Signal to Noise Ratio) of 25dB or more. This is generally sufficient for normal data clients, while allowing for VoIP infrastructure to be introduced later.
The local status page also provides info about channel utilization on current channels. It also describes the amount of time that both WiFi and non-WiFi (interference) are present on each of those channels, and provides a list of all the other access points (Meraki and non-Meraki) that it can hear on the same channel (under the Neighbors tab). These statistics can help provide a good estimate of the coverage provided by the access point, along with any interference that may be present across those channels.
The following article outlines how to perform a full site survey, in-depth, and is strongly recommended for additional guidelines:
Useful Links and Resources
Please refer to the following articles for additional information about wireless behavior and troubleshooting:
