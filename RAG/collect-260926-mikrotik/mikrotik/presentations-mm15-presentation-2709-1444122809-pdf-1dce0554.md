---
id: collect-260926-mikrotik/mikrotik/presentations-mm15-presentation-2709-1444122809-pdf-1dce0554
title: "presentations-mm15-presentation-2709-1444122809-pdf-1dce0554"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2015-10-01"]
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/presentations-mm15-presentation-2709-1444122809-pdf-1dce0554.md
source_anchor: ""
source_lines: [1, 254]
sha256: 550192a3efb69671ab0c6b31f80bef3afcb762fc02f3cb03a348ad65f95d932c
---

# presentations-mm15-presentation-2709-1444122809-pdf-1dce0554

Hotel Wifi Solution with MikroTik 
CAPsMAN, HotSpot and User Manager 
MikroTik MUM Myanmar October 1, 2015 
Presented By Aung Myo

About Me 
• Mr. Aung Myo 
• Director of NETPRO Myanmar Co.,Ltd 
• Mikrotik Certified Consultant  
• MCSE, MCITP(EA), JNCIA(ER), JNCIS(ER), CCNA, CCNA-S,CCDA, CCDP , CCNP , 
MTCNA, MTCRE, MTCWE, MTCTCE, MTCINE 
• www.netpromyanmar.net 
• aungmyo@netpromyanmar.com 
• 09448009449, 09975558020 
• www.facebook.com/netpromyanmarcompany 
• www.facebook.com/mikrotikmyanmar

Hotel Wifi Solution 
1. Receptionist Create the Wifi Internet access account in 
Mikrotik User Manager depend on room type or 
whatever  
2. Generate voucher and give it to Guest for Internet 
Access

Hotel Wifi Solution with MikroTik 
• CAPsMAN (Controlled Access Point system 
Manager) also known as WLAN controller 
 
• HotSpot 
 
• User Manager

CAPsMAN Feature 
• Centralized management of RouterOS APs 
• Dual Band AP support 
• Provisioning of APs 
• MAC and IP Layer communication with APs 
• Certificate support for AP communication 
• Full and Local data forwarding mode 
• RADIUS MAC authentication 
• Custom configuration support

CAPsMAN Version 
• CAPsMAN v1 
• CAPsMAN v2  
NOTE: CAPsMAN v2 is NOT compatible with CAPsMAN v1 
(CAP) 
Controlled Access Points 
AP  
(CAPsMAN) 
Controlled Access Point system Manager 
Controller 
CAPsMAN+CAP CAP CAP

CAPsMAN LAB 
Internet
1
CAPsMAN+CAP1
CAP2 CAP3
SSID:HOTEL-GUEST(VLAN10-192.168.10.0/24)
SSID:HOTEL-MGMT(VLAN20-192.168.20.0/24
SSID:HOTEL-HR(VLAN30-192.168.30.0/24)
SSID:HOTEL-ADMIN(VLAN40-192.168.40.0/24)
HOTEL-GUEST
HOTEL-MGMT
HOTEL-HR
HOTEL-ADMIN
Internet
2

Require Package List 
CAPsMAN 
CAP

Network Infra Configuration 
• 1. Interface Master/Slave Configuration 
• 2. Bridge Configuration 
• 3. VLAN Configuration 
• 4. IP addressing 
• 5. WAN Configuration with Load Balancing 
• 6. NAT, NTP , Timezone and DNS 
 
 
CAPsMAN+CAP1, CAP2, CAP3

Internet
1
CAPsMAN+CAP1
CAP2 CAP3
SSID:HOTEL-GUEST(VLAN10-192.168.10.0/24)
SSID:HOTEL-MGMT(VLAN20-192.168.20.0/24
SSID:HOTEL-HR(VLAN30-192.168.30.0/24)
SSID:HOTEL-ADMIN(VLAN40-192.168.40.0/24)
HOTEL-GUEST
HOTEL-MGMT
HOTEL-HR
HOTEL-ADMIN
Internet
2
CAPsMAN LAB 
WAN1 WAN2 
Bridge1 
Bridge1 
Bridge1

CAPsMAN and CAPs configuration 
 Click CAPsMAN interface, check the “Enabled” for CAPsManager device. 
 Click  Wireless interface>click Interfaces tab>click the CAP button  
     and Check the “Enabled” for CAP device. 
 
* CAPsManager device itself can be enable CAP .

CAPsMAN see all CAPs 
 When we enabled CAPsMAN and CAP , you may see the remote AP under the tab of Remote CAP of 
CAPsMAN device.

Recipe of CAPsMAN 
• Channel 
 Can define Frequency and width, etc.. 
• Datapath 
Can define the data forwarding related settings, such as bridge to which particular 
interfaces should be added automatically as port. 
• Security 
 Can define authentication types or passphrase 
• Configuration 
Main wireless settings group, includes setting such as SSID and additionally binds 
together with other setting such as channel, datapath, security, etc…  
• Provisioning 
 Can define each AP for provisioning.

Channel Configuration

Datapath Configuration

Security Configuration

Virtual AP Configuration

Manual Provisioning

CAPsMAN Finished!!!! 
All created virtual AP can see with bounding condition.

HotSpot Configuration 
To use this feature, you must have installed and enabled  "hotspot" package. 
 
1.Go to IP>hotspot. Click the hotspot setup button under the server tab. 
So you will see the pop-up box will appear. Select a desire interface from the drop down list. Then, Click Next. 
 
2.It is automatically detect the IP address of the interface. Then, Click Next. 
 
3.Check the IP pool ranges. You can adjust your pool ranges to give out. Then, Click Next. 
 
4.We don't need certificate and SMTP server address. Leave as default. 
 
5.Set the DNS server IP address. And then, Click Next. 
 
6.Set the DNS Name.(Example :www.example.com ) 
 
7.You must create at least one hotspot user account. 
 If you set up those step correctly, you will be created hotspot successfully. 
  
                                 * You can manage the hotspot user by connecting with User Manager.

HotSpot Configuration

Customize the Wi-Fi Login Page 
1.Click the "Files" tab. In the "File List" ,find the hostpot folder . 
 
2.In the "hotpot" folder, you may see the subfolder such as "hotspot/ img","hotspot/lv, 
etc... and files such as "hotspot/login.html","hotspot/logout.html",etc.. 
 
3. You can simply drag and drop the "hostpsot/login.html" and "hotspot/img" to your 
desktop to edit. 
 
4.After finished your editing, place those file correctly to the destination folder. 
 
*If you have web programming experiences you can do more beautiful.

HotSpot Files

Customize the Login Page

RADIUS Setup for User Manager 
User Manager is one of built-in radius alike software includes in Router OS packages. 
 
To use this feature, you must have installed and enabled  "usermanager" package.  
 
1.Click the "Radius" tab. Simply press the + button to add "New Radius Server" 
 
2.Check you desire service, type the IP address or User Manager router address or 
127.0.0.1 and secret. 
And then click apply and ok. 
 
3.Open the web browser, type the User Manager router address/userman to access 
the User Manager Page. 
 
4.In the "User Manager "page, create new profile, limitation, user , etc... as per 
requirement.

RADIUS Setup for User Manager

User Manager Setup

User Manager Setup 
1.Add Routers  
 - IP address and secret key of income routers, logs , etc... 
2.Create “Profile Limitation”  
 - Bandwidth, Data upload and download limitation, etc… 
3.Create “profile” 
 - Shared user , Time Validity, etc… 
4.Create User 
 -Single user or Batch user with profile assignment  
5.Creat Voucher Code 
 -Generate one or many voucher code at a time. 
6.Customize the page 
 -this is optional

Continue.. Setup

Continue.. Setup

Continue.. Setup

Continue.. Setup

Customize the User Manager Page 
1.Log into your Mikrotik User Manager , go to settings tab. 
A new window will appear and see tabs “Appearances, Style, Template,” etc… 
 
2.Click on Style tab, modify the Main background, Logo, Logo text, etc… 
 
3.Click on Templates tab, edit the HTML code for voucher code template.

Customize the User Manager Page

Customize the Voucher Code

Real World Hotel Wifi Infra 
WAN1 WAN2 WAN3 WAN4
CAPsMAN
CAP1 CAP2 CAP3 CAP48
POE
10 Gbps

Any Question?

• References;  
 
• http://wiki.mikrotik.com/wiki/Manual:CAPsMAN 
• http://wiki.mikrotik.com/wiki/Manual:Hotspot_Introduction 
• http://wiki.mikrotik.com/wiki/Manual:User_Manager

Thank You
