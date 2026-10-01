---
id: collect-261001-cisco/cisco/2022-10-01-creating-a-guest-wifi-network-using-a-vlan-with-a-cisco-aironet-acces-77b37a28-1
title: "2022-10-01-creating-a-guest-wifi-network-using-a-vlan-with-a-cisco-aironet-acces-77b37a28"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-cisco/2022-10-01-creating-a-guest-wifi-network-using-a-vlan-with-a-cisco-aironet-acces-77b37a28.md
source_anchor: ""
source_lines: [1, 80]
sha256: 5d4d53618036ac2dc68bb23f470c8953bb511d64b555314e46fda3621f9c2073
---

# 2022-10-01-creating-a-guest-wifi-network-using-a-vlan-with-a-cisco-aironet-acces-77b37a28

The aim of this howto is to explain the steps involved in setting up a second WiFi network in your home that does not have any access to all other devices within the home, such that your guests can have access to the internet but no access to your home network. I wrote it mainly because Cisco Aironet access points have a web interface that is very broken, causing me many dead ends and frustrated “Why isn’t it working? It should be working!” moments until I gave up and configured it via its console interface. The commands used are arcane if you are not well-versed in Cisco IOS so I wanted to write it all down whilst it was still fresh in my memory. I should also note that this guide assumes IPv4 only – I do not use IPv6.
The process breaks down into three sections :
- Configure the Cisco Aironet access point WiFi networks and VLANs.
- Configure your switch(es) to pass the VLAN traffic.
- Configure the OPNsense firewall to route the VLAN traffic appropriately.
1.  Configuring  the  Cisco  Aironet  access  point  WiFi  networks  and  VLANs 
I will assume that you already have a cable that converts the RJ-45 serial connector on the AP to something your computer can understand (old-skool serial or perhaps a serial-to-USB convertor). Such things are easily found online. I am configuring from a Linux machine; YMMV if you’re using Windows or a Mac, but both have equivalent programs to open a console on a serial connection. Under Linux the serial-to-USB adapter automagically creates a device entry in /dev, so I used screen to open a console connection, as follows:
screen /dev/ttyUSB0
Cisco IOS configuration operates using the concept of shells-within-shells. To first of all get into the ‘root shell’ from the initial shell we type:
en
Enter your root password if one is configured, when prompted. N.B. This password can be adjusted via the web interface – it is the one for the ‘cisco’ user. From here, the command to display the current configuration is:
show run
In terms of altering configuration, one either has to alter one of the existing lines or enter a prefix of ‘no’ to disable that item. IOS will simply throw an error and make no change if you get the syntax wrong. It is also important to bear in mind that one only alters the current running configuration, and a specific command is required to write that new config to NVRAM so that it survives a reboot of the device.
To start configuring, we need to enter the configuration shell with the command configure terminal, which can be abbreviated as follows:
conf t
We have the following items to configure, and due to the vagaries of Cisco APs they have to be done in this order:
- Define VLANs
- Define SSIDs, associating each with a VLAN
- Associate the SSIDs with a radio
Define  VLANs 
The way VLANs work, there is always a so-called ‘native’ VLAN which is that used by any traffic that is not explicitly tagged to use a VLAN. By convention, the native VLAN is number 1. We need therefore to explicitly name the native VLAN and create the one we’ll use for guest WiFi traffic:
dot11 vlan-name Native-VLAN vlan 1
dot11 vlan-name Guest-WiFi vlan 2
Define  SSIDs 
Here’s the syntax to define two SSIDs (your personal one and the one to be used for guest WiFi):
dot11 ssid personal ssid name
vlan 1
authentication open 
authentication key-management wpa version 2
guest-mode
wpa-psk ascii personal wpa2 password
dot11 ssid guest ssid name
vlan 2
authentication open 
authentication key-management wpa version 2
guest-mode
wpa-psk ascii guest wpa2 password
Obviously enough, the items in italics are for you to supply and cannot contain spaces. For those who’re interested, authentication open means that authentication is handled by the AP itself (i.e. not a captive portal or RADIUS or whatever) and guest-mode, bizarrely enough, means broadcast the SSID.
Associate  SSIDs 
We now associate the VLAN’ed SSIDs with a radio. For a dual band device the simplest approach is one network per radio. Cisco APs do support multiple SSIDs per radio, but I couldn’t work out how to configure it! Let me know if you have a good resource for this. Each time we configure an interface IOS drops us into a new shell, so we have to exit the sub-shell each time. Doing so also causes the radio configuration to be applied and the radio restarted. Dot11radio0 is the 2.4GHz radio, whilst 1 is the 5GHz. Channel can be set to your preference for radio0, whilst 5GHz auto-selects the least congested channel automatically.
int dot11radio0
ssid personal ssid name
channel 2422
encryption vlan 1 mode ciphers aes-ccm
exit
int dot11radio1
ssid guest ssid name
encryption vlan 2 mode ciphers aes-ccm
exit
Save  configuration  changes 
With all configuration done we need to write it to the EEPROM so that it will survive a reboot, as follows:
wr
Now you can disconnect from the access point.
Configure  your  switch(es)  to  pass  the  VLAN  traffic 
This item will depend upon the switch topology you have and the models of device you use. The rules to follow are:
- Ensure you’re using 802.1Q tagging and not port-based tagging.
- Ensure VLAN 1 is assigned to all ports and is untagged.
- Assign VLAN 2 as tagged to the port into which the Cisco AP is plugged.
- Assign VLAN 2 as tagged to the port into which your firewall (or router) is plugged.
In this way, VLANs 1 & 2 can pass traffic from the AP, 1 traffic can go anywhere, 2 traffic can only go to and from the firewall. In truth, VLAN 2 can be tagged to all ports if you’d like to avoid forgetting that certain ports are ‘special’; packets tagged as being for VLAN 2 by the AP are going to ‘travel down’ the second VLAN and end up treated uniquely by the firewall. The trade-off of this approach is that any device physically connected to the switch can snoop on VLAN 2. In this application that’s not really an issue – our desire is to prevent VLAN 2 traffic from entering VLAN 1.
Configure  the  OPNsense  firewall  to  route  the  guest  VLAN  traffic  appropriately 
Okay, your guest WiFi VLAN traffic is now battering itself against the LAN-side port of your firewall, so you need to configure the firewall to know what to do with it. Setting up OPNsense breaks down into:
- Define the existence of the VLAN.
- Create a virtual interface attached to the VLAN.
- Enable DHCP for the virtual interface.
- Create firewall rules to prevent the virtual interface seeing local networks but allowed out on the internet.
Define  the  existence  of  the  VLAN 
I followed this guide, which breaks down to going through the following menu path:
Interfaces > Other Types > VLAN > Add
The VLAN tag will be 2 (if you used the same as I did above) and the Parent Interface will be that of your LAN. Give it a meaningful name and leave the VLAN priority as-is.
Create  a  virtual  interface  attached  to  the  VLAN 
Follow this menu path:
Interfaces > Assignments
Then beside ‘New interface:’ at the bottom of the list select the VLAN name that you created in the previous step and click the + button.
It will now be on the list of interfaces in the menu tree on the left. Select the new interface, check the Enable box, choose a subnet that devices in the VLAN will use and enter it in the IPv4 Address (e.g. 192.168.100.1/24) and save.
Enable  DHCP  for  the  virtual  interface 
Follow this menu path:
Services > DHCPv4 > name_of_your_interface
Check the ‘Enable’ checkbox and enter a range of addresses that will be handed out to devices that join the guest WiFi network. Obviously these will need to be in the subnet you defined for the interface in the previous step.
Create  firewall  rules 
