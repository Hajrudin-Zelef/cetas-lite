---
id: collect-261001-fortinet/fortinet/ha-with-fortigate-vm-and-third-party-products
title: "ha-with-fortigate-vm-and-third-party-products"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/ha-with-fortigate-vm-and-third-party-products.md
source_anchor: ""
source_lines: [1, 39]
sha256: 6398d5d7b5ecb0cd9496fa7d62f8ce20ab122b3793cdf7a63061e81f1e71554f
---

# ha-with-fortigate-vm-and-third-party-products

**H****A with FortiGate-VM and third-party products**

This chapter provides information about operating FortiOS VM cluster and operating FortiGate clusters with third party products such as layer-2 and layer-3 switches.


**Fo****r****t****iG****a****t****e****–****V****M for VMware HA configuration**

If you want to combine two or more FortiGate-VM instances into a FortiGate Clustering Protocol (FGSP) High Availability (HA) cluster the VMware server’s virtual switches used to connect the heartbeat interfaces must operate in promiscuous mode. This permits HA heartbeat communication between the heartbeat interfaces. HA heartbeat packets are non-TCP packets that use Ethertype values 0x8890, 0x8891, and 0x8890. The FGCP uses link-local IP4 addresses in the 169.254.0.x range for HA heartbeat interface IP addresses.


To enable promiscuous mode in VMware:

**1****.** In the vSphere client, select your VMware server in the left pane and then select the **C****on****f****i****gu****r****a****t****i****o****n** tab in the right pane.

**2****.** In **H****a****r****d****w****a****r****e**, select **N****e****t****w****o****r****k****i****ng**.

**3****.** Select **P****r****op****e****r****t****i****e****s** of a virtual switch used to connect heartbeat interfaces.

**4****.** In the **P****r****op****e****r****t****i****e****s** window left pane, select **v****S****w****i****t****c****h** and then select **E****d****i****t**.

**5****.** Select the **S****ec****u****r****i****t****y** tab, set **P****r****o****m****i****sc****uou****s Mode** to **A****cce****p****t**, then select **O****K**.

**6****.** Select **C****l****o****se**.


You must also set the virtual switches connected to other FortiGate interfaces to allow MAC address changes and to accept forged transmits. This is required because the FGCP sets virtual MAC addresses for all FortiGate interfaces and the same interfaces on the different VM instances in the cluster will have the same virtual MAC addresses.


To make the required changes in VMware:

**1****.** In the vSphere client, select your VMware server in the left pane and then select the **C****on****f****i****gu****r****a****t****i****o****n** tab in the right pane.

**2****.** In **H****a****r****d****w****a****r****e**, select **N****e****t****w****o****r****k****i****ng**.

**3****.** Select **P****r****op****e****r****t****i****e****s** of a virtual switch used to connect FortiGate VM interfaces.

**4****.** Set **MAC Address Changes**to**Accept**.

**5****.** Set **Fo****r****g****e****d Transmits** to **A****cce****p****t**.
