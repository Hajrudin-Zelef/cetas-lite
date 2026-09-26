---
id: collect-260926-mikrotik/mikrotik/setting-up-a-mikrotik-hotspot-with-usermanager-step-by-step-1
title: "setting-up-a-mikrotik-hotspot-with-usermanager-step-by-step"
domain: mikrotik
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/dhcp-dns/setting-up-a-mikrotik-hotspot-with-usermanager-step-by-step.md
source_anchor: ""
source_lines: [1, 284]
sha256: 94a07a7eae96fafd0426b34d5aefd5330476893a1b8b0f993dd809493774c4cd
---

# setting-up-a-mikrotik-hotspot-with-usermanager-step-by-step

### Setting up a Mikrotik Hotspot with UserManager (Step-By-Step)

Mikrotik RouterOS provides a very powerful Hotspot Feature. This can be used with the Mikrotik built in Radius server (Userman) or with a remote Radius/Freeradius Server.


You will need:


Mikrotik RouterBOARD:







Notes:














This example uses Google's DNS service. You can easily make use of OpenDNS to implement simple filtering, or use your own DNS servers if needed.








Check if you can ping a public ip address like 8.8.8.8








I use any unused private ip range for this, it is used simply as an interface to run the RADIUS server on.





If you are planning to run the hotspot on a single interface you may skip this step.


I make use of the private 192.168.0.1/24 range for the hotspot network, but you can use whatever is suitable in your setup.







If you would like multiple interfaces to have access to the hotspot, you can repeat this process, only changing the interface each time.

If you are running on a RouterBOARD 750 or similar, you will need to add the ports that you AP's are connected to, to the bridge.

























I will soon do a UserManager Tutorial as well.

You will need:

Mikrotik RouterBOARD:

- Level 4 or better licence (Lower licences will allow only a single Hotspot client)
- RouterOS 6.x (5.x will also work, but this tutorial is based on v6.7)

The network will be configured as below. You may need to adjust the IP Addresses to suit your needs

Notes:

*The RouterBOARD CPU and RAM will directly affect the performance of your Hotspot, so consider beforehand how many clients you wish to connect.**A RouterBOARD 750 can comfortably run about 25-50 users.**In my example I will use a RouterBOARD 532 with one 2.4ghz WLAN card***Step 1: Configure internet access on the router**
**Add Router IP Address:***Change the IP to match your network configuration*
1. Click on the IP Menu
2. Click on the Addresses Menu
3. Click "+"
4. Enter the IP Address you wish to assign to the router, this will be the outward facing IP, so make sure to select the ethernet interface that will give the router internet access.
5. Click on "Apply"

**Configure Upstream DNS Server:**

1. Click on the IP Menu
2. Click on the DNS Menu
3. Enter your desired DNS server - here I am using Google's DNS
4. Click on "Apply"

**Configure Default Route:**
1. Click on IP
2. Click on Routes
3. Click on "+"
4. Enter 0.0.0.0/0 as the Dst. Address
5. Enter 10.0.0.1 as the Gateway

**Test:**
Check if you can ping a public ip address like 8.8.8.8

1. Click on Tools
2. Click on Ping
3. Enter a publicly available address
4. Click Start

## Step 2: Install User Manager and Hotspot

If you plan to use a stand alone Radius Server, you may skip this step.

Download the firmware package from Mikrotik

Extract the zip file on your local drive


Extract the zip file on your local drive

1. Make sure that the version of the file matches the version and architecture of your device
2. Open the Files window on winbox
3. Drag the "user-manager-X.X-xxxxxx.npk" to the files window.
4. Do the same for "hotspot-X.X-xxxxxx.npk".
5. Reboot the router (/system reboot)

## Step 3: Configure interfaces

First, we need to configure two Bridge interfaces. The first one will be a loopback interface. I have found in the past that if you use the normal loopback address (127.0.0.1), or one of the other static addresses, for the Radius (Usermanager) server, you may experience some difficulties.




**3.1.1 - Create Loopback Bridge**
1. Click on the "Bridge" menu
2. Click on "+"
3. Enter "Loopback" for the bridge name
4. Click "Apply"

**3.1.2 - Add Loopback Bridge IP Address**

I use any unused private ip range for this, it is used simply as an interface to run the RADIUS server on.

1. Click on the IP Menu
2. Click on the Addresses menu
3. Click the "+" button
4. Enter "10.10.0.1/32" as the IP Address
5. Select the "Loopback" Interface
6. Click "OK"

**3.2.1 - Create Hotspot Bridge**

If you are planning to run the hotspot on a single interface you may skip this step.

1. Click on the "Bridge" menu
2. Click on "+"
3. Enter "Hotspot" for the bridge name
4. Click "Apply"

**3.2.2 - Add Hotspot Bridge IP Address**

I make use of the private 192.168.0.1/24 range for the hotspot network, but you can use whatever is suitable in your setup.

1. Click on the IP Menu
2. Click on the Addresses menu
3. Click the "+" button
4. Enter "192.168.0.1/24" as the IP Address
5. Select the "Hotspot" Interface
6. Click "OK"

**3.2.3 - Add Hotspot Ports to Bridge**

If you would like multiple interfaces to have access to the hotspot, you can repeat this process, only changing the interface each time.

If you are running on a RouterBOARD 750 or similar, you will need to add the ports that you AP's are connected to, to the bridge.

**3.3 - Configure the Access Point**

If you are using a RouterBOARD 750 or similar, you will not be using this section.

You may choose to implement security on your access point, but since this is a captive portal, you should not need to use any security. This tutorial will not include any security settings.

1. Click on the "Wireless" Menu
2. Double click on the Wireless Interface that you will be using
3. Set the mode to "ap-bridge"
4. Set the band to 2Ghz-B/G (or otherwise if needs be)
5. Change the SSID to "Hotspot", or whatever suits you.
6. Click "OK"

## 4 - Configure the Hotspot

1. Click on the "IP" menu. If this option is not available refer to step 2
2. Click on the "Hotspot" item
3. Click on "Hotspot Setup". This will start the Hotspot Setup Wizard

**4.1 - The Hotspot Wizard**

1. Select the Hotspot bridge as the Hotspot Interface
2. Click Next

1. Click next - The address range should be filled in automatically as per our network configuration.

1. Click Next - the address pool should be pre-populated with the right settings

1. This tutorial will not cover the use of Certificates, so you may select "none" and click next

1. Enter the IP-Address of your SMTP server. Many providers do not allow use of their SMTP servers outside their own network, so this option allows you to circumvent the SMTP server configured on the client's device in favor of your own. (You may even specify the SMTP server of you own provider in some cases)
2. Click "Next"

These are the upstream DNS servers used by the hotspot.

1. Enter one or more upstream DNS servers, you can use OpenDNS to provide you with a basic filtering service. Here I use Google's public DNS.
2. Click "Next"

1. Enter a host name for the local Hotspot. I am using hotspot.example.com, but this could be anything you want.
2. Click "Next"

1. Enter a name for your administrative Hotspot user.
2. Enter a password for your administrative user.
3. Click "Next"

1. Click "OK" to complete your hotspot setup.

Congratulations, you have now set up basic functionality for a Mikrotik Wireless Hotspot, you can create users under "IP->Hotspot->Users. But alas, you still need to configure the Usermanager for a fully featured hotspot.

## Step 5 - Configuring UserManager

**5.1 Setting up the Hotspot to use RADIUS**

1. Click on the "IP" menu
2. Click on "Hotspot"
3. Select the "Server Profiles" tab
4. Double click on "hsprof1"
5. Select the "RADIUS" tab
6. Tick the "Use RADIUS" tickbox
7. Click "OK"


1. Click on "RADIUS"
2. Click on "+"
3. Tick the "hotspot" tickbox
4. Add the loopback bridge IP to the address field, in this tutorial 10.10.0.1
5. Choose a secure password
6. Click "OK"


1. Using your browser of choice, connect to http://router-ip/userman
2. Click "Log In" - The default username is ***admin***  with no password

1. Once you have logged in, click on the "Routers" menu
2. Click "Add" then "New"
3. Enter "Local Router" as the name
4. Enter the Loopback Bridge IP address
5. Enter the password you chose earlier.
6. Click "OK"

