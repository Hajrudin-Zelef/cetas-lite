---
id: collect-240926-storagereview/storagereview/fr-review-enabling-gen5-ssd-evaluations-with-serial-cables-57af70e4
title: "fr-review-enabling-gen5-ssd-evaluations-with-serial-cables-57af70e4"
domain: storagereview
role: reference
task: reference
actors: ["Broadcom", "Samsung"]
dates: []
keywords: ["consumer", "cost"]
source: docs/RAG/clean_en/storagereview/fr-review-enabling-gen5-ssd-evaluations-with-serial-cables-57af70e4.md
source_anchor: ""
source_lines: [1, 23]
sha256: a2238b5f1efd001ac55d5efcdd7a67986f24168142ba71b4fc720140e49cf1cc
---

# fr-review-enabling-gen5-ssd-evaluations-with-serial-cables-57af70e4

<!-- source: https://www.storagereview.com/fr/review/enabling-gen5-ssd-evaluations-with-serial-cables -->

Testing and reviewing SSDs brings a new set of complexities with each interface update or form factor change. These challenges have never been more widespread than today, as the industry moves not only to Gen5 SSDs, but also to the most diverse set of SSD form factors ever created.
In the Gen4 SSD space, we had just seen M.2 and U.2 SSDs. U.2 SSDs typically always appeared in a 15mm enclosure, but there were occasional exceptions, notably Samsung, which used the 7mm enclosure. As far as form factor diversity goes, that was about it. To be honest, we did review a set of E1.S EDSFF SSDs, but that review was an outlier. With EDSFF now gaining momentum, not only with hyperscalers but with all server vendors, the game is changing, and so is the test equipment.
This is where Serial Cables comes in, providing the critical infrastructure for our lab for 5th generation SSD evaluation. Our primary test platform is a Dell PowerEdge R760 server equipped with front-facing U.2 bays. However, Dell only supports 4th generation drives with its U.2 backplane, and reserves 5th generation for the E3.S backplane. The difficulties of standardizing our SSD test bench are already apparent. This problem is not unique to Dell: Lenovo, HPE, Supermicro and other manufacturers make design choices that make it nearly impossible to use a single test platform for all SSDs.
What Serial Cables has done is create a Gen5 JBOF with modular trays that can adapt to any form factor. The 8-bay unit is paired with a host card, cabling and a series of "paddle cards" to mount the drives. And with the exception of a few odd drive shapes like the enormous E.1L SSDs, the trays can accommodate most drives. what we need to test, very well. This means that with a single modern server, we can test any Gen5 SSD in exactly the same way.
Serial Cables Gen5 Equipment
As noted, this Gen5 test suite includes several pieces of Serial Cables equipment:
- PCIe Gen5 x16 MCIO Host Card
- Gen5 PCIO 8-Bay E3 Passive JBOF
- Gen5 PCIe U.2 Paddle Card
- Gen5 PCIe M.2 Paddle Card
PCIe Gen5 x16 MCIO Host Card
The host card comes with a Broadcom Atlas2 PCIe switch. It has 4 MCIO ports each providing an x4 connection to directly connect various PCIe devices. We use this card inside the Dell PowerEdge R760 to connect externally to the Serial Cables JBOF. The card costs $3,995.00.
The card itself has an FHHL footprint, allowing broad compatibility with most servers. For systems that only support half-height cards, they offer a similar offering in that form factor. On the system side, no drivers are required, making installation simple regardless of the operating system.
PCIe Gen5 8-Bay E3 Passive JBOF
The heart of this system is the Serial Cables Gen5 PCIe 8-Bay E3 Passive JBOD universal test platform. This JBOF supports E3, U.2, U.3 and M.2 form factors, with a common EDSFF backplane. Gen5 SSDs are more complicated than previous generations of SSDs due to the large number of shapes and sizes in which they come. Fortunately, this Gen5 JBOF is designed to solve this problem as it can support many different types of drives.
The JBOF has a total of eight bays, meaning we can test larger groups of SSDs if we wish, or multiples without having to swap drives. An out-of-band management function is built into the JBOD for basic operations such as enabling or disabling slots, as well as diagnostic functions. This device has a list price of $2,995.00.
PCIe Gen5 U.2, U.3 and M.2 Paddle Cards
These paddle cards are compatible with the JBOF and cost $75.00. We currently operate U.2/U.3, E3DSFF and M.2 paddle cards in the JBOD for enterprise and consumer PCIe Gen5 evaluations.
Final Thoughts
Overall, the Serial Cables PCIe Gen5 host card and JBOF combo have become a fantastic integral component of the StorageReview lab and allow us to test a wide range of emerging Flash products on a single test platform. This allows us to have reproducible performance reports, which is fundamental to presenting reliable data.
While this system is clearly not intended for the general public, any lab like ours that tests a wide variety of drives already knows the Serial Cables brand. They have been making adapter cards for many years that have become essential in many test environments. This Gen5 JBOD system is a great additional tool that we will rely on throughout our Gen5 SSD testing. You can check out the latest results obtained with this test system in our Memblaze PBlaze7 review.
