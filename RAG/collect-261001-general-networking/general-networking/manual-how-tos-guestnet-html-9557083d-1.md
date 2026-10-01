---
id: collect-261001-general-networking/general-networking/manual-how-tos-guestnet-html-9557083d-1
title: "manual-how-tos-guestnet-html-9557083d"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/manual-how-tos-guestnet-html-9557083d.md
source_anchor: ""
source_lines: [1, 125]
sha256: b74ccab486136051bd03bea6a3b4e45f616f8c2a73b081b0c28003571feecc45
---

# manual-how-tos-guestnet-html-9557083d

Setup a Guest Network
This how to will explain how to setup a guest network using the captive portal. Guest Networks are widely used to allow guests controlled internet access at hotels, RV Parks or businesses.
Note
For the example we expect the GUESTNET interface to be connected with your actual guest network switch or access point. This tutorial does not explain how to setup a wireless network.
Security considerations
The Captive portal functionality operates entirely on the information available in the network. Protocols such as ARP or NDP do not provide any inherent proof of ownership. Since the portal has no control over the client, it is impossible to cryptographically prove a relationship between identity and device. Therefore, to improve security, the layer 2 network(s) attached to the portal must be properly isolated on the access point/switch level.
These features are often called “layer 2 isolation” for access points or “port isolation” for switches and prevent direct client-to-client communication. These features cannot prevent spoofing, but lower visibility in the network.
If stronger identity is a requirement, the Captive Portal likely isn’t for you. Consider using 802.1X network access control backed by a RADIUS/ policy server instead.
Note that this is less relevant if you use a layer 3 network such as WireGuard.
Businesses
Businesses usually want to share internet access with their guest and show them a landing page with a welcome message and some usage guidelines (policy). At the same time it is important to make sure guests won’t be able to access the company’s local network and limit the maximum internet usage.
Hotels and RV Parks
Hotels and RV parks usually utilize a captive portal to allow guests (paid) access to internet for a limited duration. Guests need to login using a voucher they can either buy or obtain for free at the reception. OPNsense has built-in support for vouchers and can easily create them on the fly. With this example we will show you how to setup the Guest Network for this purpose and setup a reception account for creating new vouchers.
Prerequisites
We will start configuration with a fresh OPNsense installation. You will need a system with a minimum of 3 ports (LAN/WAN/GUESTNET) for this tutorial.
Good to know
As the Hotel/RV Parks setup is almost identical to the business setup we will start with that and after finishing add/change the specifics to match the Hotel Guest setup.
Step 1 - Configure Interface
For the Guest Network we will add a new interface. Go to And use the + to add a new interface. Press Save. The new interface will be called OPT1, click on [OPT1] in the left menu to change its settings.
Select Enable Interface and fill in the following data for our example:
| Description | GUESTNET | A descriptive name for the interface | 
| Block Private networks | unselected |  | 
| Block bogon networks | unselected |  | 
| IPv4 Configuration Type | Static IPv4 | Set a static IPv4 address for the example | 
| IPv6 configuration Type | None |  | 
| MAC address | (Leave Blank) |  | 
| MTU | (Leave Blank) |  | 
| MSS | (Leave Blank) |  | 
| Speed and duplex | Default | You may also select the speed when known | 
| Static IPv4 address | 192.168.200.1/24 | We will use this segment for our guests | 
| IPv4 Upstream Gateway | Default |  | 
Press Save and then Apply changes.
Step 2 - Configure DHCP Server
Go to .
- Fill in the following to setup the DHCP server for our guest net (leave everything
- else on its default setting):
| Enable | Checked | Enable the DHCP server on GUESTNET | 
| Range | 192.168.200.100 to 192.168.200.200 | Serve IPs from this range | 
| DNS servers | 192.168.200.1 | Supply a DNS with the lease | 
| Gateway | 192.168.200.1 | Supply a gateway with the lease | 
Click Save.
Step 3 - Add Firewall Rules
Note
Rules to allow DNS and access to the captive portal zone webserver are installed automatically. If you are overriding this behavior, install the rules as listed in Captive Portal firewall rules before any other rules.
Go to to add a new rule.
Now add the following rules (in order of prevalence):
Block Local Networks
| Action | Block | Block this traffic | 
| Interface | GUESTNET | The GuestNet Interface | 
| Protocol | any |  | 
| Source | GUESTNET net |  | 
| Destination | LAN net |  | 
| Category | GuestNet Basic Rules | Category used for grouping rules | 
| Description | Block Local Networks |  | 
Click Save.
| Action | Block | Block this traffic | 
| Interface | GUESTNET | The GuestNet Interface | 
| Protocol | any |  | 
| Source | GUESTNET net |  | 
| Destination | This Firewall |  | 
| Category | GuestNet Basic Rules | Category used for grouping rules | 
| Description | Block Firewall Access |  | 
Click Save.
Note
These rules are used to block access to our local LAN network and firewall access from the Guests. If you have multiple local networks then you need to block each of them with multiple rules or use a bigger subnet to cover them all.
Allow Guest Networks
| Action | Pass | Allow this traffic | 
| Interface | GUESTNET | The GuestNet Interface | 
| Protocol | any |  | 
| Source | GUESTNET net |  | 
| Destination | any |  | 
| Destination port range | any |  | 
| Category | GuestNet Basic Rules | Category used for grouping rules | 
| Description | Allow Guest Network |  | 
Click Save and then Apply changes
Step 4 - Create Captive Portal
Go to
To add a new Zone press the + in the lower right corner of the form.
Note
When using multiple interfaces with the captive portal then each interface can have its own zone or multiple interfaces can share a zone.
For the Business setup we will start with the following settings:
| Enabled | Checked |  | 
| Interfaces | GUESTNET | Remove the default and add GUESTNET | 
| Authenticate using | (blank) | Remove any default setting | 
| Idle timeout | 0 | Disable Idle Timeout | 
| Hard timeout | 0 | No hard timeout | 
| Concurrent user logins | Unchecked | A user may only login once | 
| SSL certificate | none | Use plain http | 
| Hostname | (leave blank) | Used for redirecting login page | 
| Allowed addresses | (leave blank) |  | 
| Custom template | none | Use default template | 
| Description | Guest Network | Choose a description for the zone | 
Save and the Apply
Step 5 - Create Template
The template feature is one of the most powerful features of OPNsense’s Captive Portal solution and it’s very easy to work with.
Let’s create a custom landing page, to do so click on the tab Templates and click on the download icon in the lower right corner ( ).
Now download the default template, we will use this to create our own. Unpack the template zip file, you should have something similar to this:
Most files of the template can be modified, but some are default and may not be changes. Upon upload any changes to the files listed in exclude.list will be ignored. Currently these include the bootstrap JavaScript and some fonts.
With the captive portal enabled the default screen looks like:
Let’s change this default with a new logo and a welcome message, to this:
To do so use your favourite editor and open the index.html file to make the changes.
Let’s make the following changes to the template:
- Change the logo to company-logo.png
- Remove the navigation bar on the top
- Remove the height and width from the <img> tag
- Add a welcome text
- Make a link to the company website
Find the following part:
<header class="page-head">
<nav class="navbar navbar-default" >
    <div class="container-fluid">
        <div class="navbar-header">
            <a class="navbar-brand" href="#">
                <img class="brand-logo" src="images/default-logo.png" height="30" width="150">
            </a>
        </div>
    </div>
</nav>
</header>
And change to:
<header class="page-head">
    <div align="center">
      <a href="#">
          <img class="brand-logo" src="images/company-logo.png">
      </a>
