---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-26-1
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "license"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-26.md
source_anchor: ""
source_lines: [1, 105]
sha256: 82994430ecb09e05fea681dd74e817075b1574e5dcd4136c937dd20a9053d323
---

# Overview

CAPsMAN allows applying wireless settings to multiple MikroTik AP devices from a central configuration interface.

More specifically, the Controlled Access Point system Manager (CAPsMAN) allows centralization of wireless network management and if necessary, data processing. When using the CAPsMAN feature, the network will consist of a number of 'Controlled Access Points' (CAP) that provide wireless connectivity and a 'system Manager' (CAPsMAN) that manages the configuration of the APs, it also takes care of client authentication and optionally, data forwarding.

When a CAP is controlled by CAPsMAN it only requires the minimum configuration required to allow it to establish a connection with CAPsMAN. Functions that were conventionally executed by an AP (like access control, client authentication) are now executed by CAPsMAN. The CAP device now only has to provide the wireless link layer encryption/decryption.

Depending on the configuration, data is either forwarded to CAPsMAN for centralized processing (*default*) or forwarded locally at the CAP itself (local forwarding mode).

Requirements

- Any RouterOS device can be a controlled wireless access point (CAP) as long as it has at least a Level 4 RouterOS license
- CAPsMAN server can be installed on any RouterOS device, even if the device itself does not have a wireless interface
- Unlimited CAPs (access points) supported by CAPsMAN
- 32 Radios per CAP maximum
- 32 Virtual interfaces per master radio interface maximum
- Not possible to use Nv2 and NStreme proprietary protocols

# Simple setup of a CAPsMAN system

Before deep-diving into the details of CAPsMAN operation, let us quickly illustrate how to set up the most basic system where you have a MikroTik router that manages two MikroTik AP devices. The benefit of CAPsMAN is that the CAP units don't need to be configured, all settings are done in the CAPsMAN server.

The CAPsMAN setup consists of defining configuration templates, which will then be pushed to the controllable AP devices (CAPs). Assuming your main router is already connected to the internet and works fine, you can proceed as follows.

In the central device, which will be your CAPsMAN server, create a new "Configuration" template with only the basic settings (network name, country, the local LAN bridge interface, the wireless password):

1. 2.

Then create a new "Provisioning" rule, which will assign the created configuration template to the CAP devices:

4.

All that remains to do on the CAPsMAN, is to enable it:

5.

Most MikroTik AP devices already support CAP mode out of the box, all you need to do, is make sure they are on the same network as your CAPsMAN, and then boot them up, while holding the reset button.

So, for example, connect the CAP device to one of the CAPsMAN device LAN ports while it is turned off, then hold the reset button, and power on the CAP device. Keep holding the button until the User LED turns solid, release now to turn on CAP mode. The device will now look for a CAPsMAN server (total time to hold the button, around 10 seconds).

The device will now show up in the CAPsMAN "Remote CAP" menu and will be "provisioned" with the configuration template, as per the provisioning settings. For more details on how to manually adjust all settings, keep reading this document.

## **CAP to CAPsMAN Connection**

For the CAPsMAN system to function and provide wireless connectivity, a CAP must establish management connection with CAPsMAN. A management connection can be established using MAC or IP layer protocols and is secured using 'DTLS'.

A CAP can also pass the client data connection to the Manager, but the data connection is not secured. If this is deemed necessary, then other means of data security needs to be used, e.g. IPSec or encrypted tunnels.

CAP to CAPsMAN connection can be established using 2 transport protocols (via Layer 2 and Layer3).

- MAC layer connection features:
  - no IP configuration necessary on CAP
  - CAP and CAPsMAN must be on the same Layer 2 segment - either physical or virtual (by means of L2 tunnels)
- IP layer (UDP) connection features:
  - can traverse NAT if necessary
  - CAP must be able to reach CAPsMAN using IP protocol
  - if the CAP is not on the same L2 segment as CAPsMAN, it must be provisioned with the CAPsMAN IP address, because IP multicast based discovery does not work over Layer3

In order to establish connection with CAPsMAN, CAP executes a discovery process. During discovery, CAP attempts to contact CAPsMAN and builds an available CAPsMANs list. CAP attempts to contact to an available CAPsMAN using:

- configured list of Manager IP addresses
- list of CAPsMAN IP addresses obtained from DHCP server
- broadcasting on configured interfaces using both - IP and MAC layer protocols.

When the list of available CAPsMANs is built, CAP selects a CAPsMAN based on the following rules:

- if **caps-man-names** parameter specifies allowed manager names (**/system identity** of CAPsMAN), CAP will prefer the CAPsMAN that is earlier in the list, if list is empty it will connect to any available Manager
- suitable Manager with MAC layer connectivity is preferred to Manager with IP connectivity

After Manager is selected, CAP attempts to establish DTLS connection. There are the following authentication modes possible:

- no certificates on CAP and CAPsMAN - no authentication
- only Manager is configured with certificate - CAP checks CAPsMAN certificate, but does not fail if it does not have appropriate trusted CA certificate, CAPsMAN must be configured with **require-peer-certificate=no** in order to establish connection with CAP that does not possess certificate
- CAP and CAPsMAN are configured with certificates - mutual authentication

After DTLS connection is established, CAP can optionally check CommonName field of certificate provided by CAPsMAN. **caps-man-certificate-common-names** parameter contains list of allowed CommonName values. If this list is not empty, CAPsMAN must be configured with certificate. If this list is empty, CAP does not check CommonName field.

If the CAPsMAN or CAP gets disconnected from the network, the loss of connection between CAP and CAPsMAN will be detected in approximately 10-20 seconds.

### **CAP Auto Locking to CAPsMAN**

CAP can be configured to automatically lock to a particular CAPsMAN server. Locking is implemented by recording certificate CommonName of CAPsMAN that CAP is locked to and checking this CommonName for all subsequent connections. As this feature is implemented using certificate CommonName, use of certificates is mandatory for locking to work.

Locking is enabled by the following command:

[admin@CAP] > /interface wireless cap set lock-to-caps-man=yes

Once CAP connects to suitable CAPsMAN and locks to it, it is reflected like this:

```
[admin@wtp] > /interface wireless cap print
...
        locked-caps-man-common-name: CAPsMAN-000C424C30F3
```
From now on CAP will only connect to CAPsMAN with this CommonName, until locking requirement is cleared, by setting **lock-to-caps-man=no**. This approach needs to be used if it is necessary to force CAP to lock to another CAPsMAN - by at first setting **lock-to-caps-man=no** followed by **lock-to-caps-man=yes**.

Note that CAP can be manually "locked" to CAPsMAN by setting **caps-man-certificate-common-names**.

### **Auto Certificates**

To simplify CAPsMAN and CAP configuration when certificates are required (e.g. for automatic locking feature), CAPsMAN can be configured to generate necessary certificates automatically and CAP can be configured to request certificate from CAPsMAN.

**Automatic certificates do not provide full public key infrastructure** and are provided for simple setups. If more complicated PKI is necessary - supporting proper certificate validity periods, multiple-level CA certificates, certificate renewal - other means must be used, such as manual certificate distribution or SCEP.

