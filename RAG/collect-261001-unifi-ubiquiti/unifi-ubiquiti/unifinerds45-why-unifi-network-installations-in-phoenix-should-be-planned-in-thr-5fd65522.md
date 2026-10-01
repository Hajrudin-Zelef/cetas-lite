---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/unifinerds45-why-unifi-network-installations-in-phoenix-should-be-planned-in-thr-5fd65522
title: "unifinerds45-why-unifi-network-installations-in-phoenix-should-be-planned-in-thr-5fd65522"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-unifi-ubiquiti/unifinerds45-why-unifi-network-installations-in-phoenix-should-be-planned-in-thr-5fd65522.md
source_anchor: ""
source_lines: [1, 8]
sha256: 4ba04df677c276eecfefec4c3514602b17695e5be33e970c248e1d2219a085b8
---

# unifinerds45-why-unifi-network-installations-in-phoenix-should-be-planned-in-thr-5fd65522

Why UniFi Network Installations in Phoenix Should Be Planned in Threes
The Valley adds commercial space quickly, and companies here often go from one location to three inside a couple of years. Tempe, Chandler, north Phoenix. A business signs a second suite before the first one is properly furnished. That pattern should shape how the first network gets built.
Whatever is configured at location one tends to get copied to two and three, so discipline at the first site pays for itself. Consistent VLAN numbering, a naming convention for access points, documented address ranges. That is the argument for treating UniFi network installations Phoenix as a repeatable build rather than a series of one-offs. Skip it and each site ends up subtly different, which nobody notices until something breaks at the location you visit least.
Get unifinerds’s stories in your inbox
Join Medium for free to get updates from this writer.
Two local details shape the design. The first is heat. Ubiquiti rates a current access point to roughly 104°F and switches sit in a similar range, so a west-facing unconditioned closet in July clears that without trying, and a loaded PoE switch adds its own heat inside a sealed room. Hardware that runs hot rarely dies that week. It dies a year and a half early, which is much harder to trace.
The second is the buildings. Flex and small-bay tilt-up space around the Loop 101 typically runs open ceilings in the twenty-foot range, and newer distribution product goes considerably higher. Twenty feet is comfortably inside what an omnidirectional access point handles well; the practical limit sits nearer thirty-five, with directional antennas above roughly fifty. Tilt-up matters for a different reason than height. A reinforced concrete panel attenuates far more than drywall or brick, so exterior walls are close to opaque and coverage doesn’t reach a yard or the suite next door on optimism.
Once there’s more than one site, driving between them to check hardware stops making sense. A single controller covering every location puts firmware updates, alerts and configuration in one place. Set that up at site two, before the habit of managing each one separately takes hold. UniFiNerds.com has more on multi-site network planning.
