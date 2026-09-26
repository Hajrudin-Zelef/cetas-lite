---
id: collect-260926-mikrotik/mikrotik/screenshots-how-to-reset-mikrotik-router-pdf-8dcd8d17-1
title: "screenshots-how-to-reset-mikrotik-router-pdf-8dcd8d17"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2014-21-01"]
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/RouterOS/screenshots-how-to-reset-mikrotik-router-pdf-8dcd8d17.md
source_anchor: ""
source_lines: [1, 128]
sha256: b8835cce23df70bd81e6c2dd6791e2766de01e6414d35bbd532fe14adacab885
---

# screenshots-how-to-reset-mikrotik-router-pdf-8dcd8d17

1/21/2014 Manual:Netinstall - MikroTik Wiki
http://wiki.mikrotik.com/wiki/Manual:Netinstall 1/7
Applies to
RouterOS:
2.9, v3, v4
Manual:Netinstall
Contents
1 NetInstall Description
2 Interface
3 Screenshot
4 NetInstall Example
4.1 Requirements
4.2 Connection process
4.3 Configuring RouterBOARD
4.3.1 Configuring RouterBOARD without COM port
4.3.2 Configuring RouterBOARD with COM port
4.4 Installation
4.5 Cleanup
4.6 Reset RouterOS Password
NetInstall Description
NetInstall is a program that runs on Windows computer that allows you to install MikroTiK RouterOS onto a PC or onto a
RouterBoard via an Ethernet network.
You can download Netinstall on our download page (http://www.mikrotik.com/download.html) .
NetInstall is also used to re-install RouterOS in cases where the the previous install failed, became damaged or access
passwords were lost.
Your device must support booting from ethernet, and there must be a direct ethernet link from the Netinstall computer
to the target device. All RouterBOARDs support PXE network booting, it must be either enabled inside RouterOS
"routerboard" menu if RouterOS is operable, or in the bootloader settings. For this you will need a serial cable.
Note: For RouterBOARD devices with no serial port, and no RouterOS access, the reset button can also start PXE booting
mode. See your RouterBOARD manual PDF for details. For example RB750 PDF
(http://www.routerboard.com/pricelist/download_file.php?file_id=118)
Netinstall can also directly install RouterOS on a disk (USB/CF/IDE/SATA) that is connected to the Netinstall Windows
machine. After installation just move the disk to the Router machine and boot from it.
Interface
The following options are available in the Netinstall window:
Routers/Drives - list of PC drives, and in the routers that were detected near the Netinstall PC
Make floppy - used to create a bootable 1.44" floppy disk for PCs which don't have Etherboot support
Net booting - used to enable PXE booting over network (your default choice)
Install/Cancel - after selecting the router and selecting the RouterOS packages below, use this to start install
SoftID - the SoftID that was generated on the router. Use this to purchase your key
Key / Browse - apply the purchased key here, or leave blank to install a 24h trial
Get key - get the key from your mikrotik.com account directly
Flashfig - launch Flashfig - the mass config utility which works on brand new devices
Keep old configuration - keeps the configuration that was on the router, just reinstalls software (no reset)
IP address / "Netmask - enter IP address and netmask in CIDR notation to preconfigure in the router
Gateway - default gateway to preconfigure in the router
Baud rate - default serial port baud-rate to preconfigure in the router
Configure script File that contains RouterOS CLI commands that directly configure router (e.g. commands produced by
export command). Used to apply default configuration

1/21/2014 Manual:Netinstall - MikroTik Wiki
http://wiki.mikrotik.com/wiki/Manual:Netinstall 2/7
Screenshot
for installation over network, don't forget to enable the PXE server, and make sure Netinstall is not blocked by your
firewall or antivirus. The connection should be directly from your Windows PC to the Router PC (or RouterBOARD), or at
least through a switch/hub.
NetInstall Example
This is a step by step example of how to install RouterOS on a RouterBoard 532 from a typical notebook computer.
Requirements
The Notebook computer must be equiped with the following ports and contain the following files:
Ethernet port.
Serial port.
Serial communications program (such as Hyper Terminal)
The .npk RouterOS file(s) (not .zip file) of the RouterOS version that you wish to install onto the Routerboard.
The NetInstall program available from the Downloads page at www.mikrotik.com
It is recommended to disable any other Network interfaces in your PC, leave only the one which is connected to your
router
Connection process
1. Connect the routerboard to a switch, a hub or directly to the Notebook computer via Ethernet. The notebook
computer Ethernet port will need to be configured with a usable IP address and subnet. For example: 10.1.1.10/24
2. Connect the routerboard to the notebook computer via serial, and establish a serial communication session with

1/21/2014 Manual:Netinstall - MikroTik Wiki
http://wiki.mikrotik.com/wiki/Manual:Netinstall 3/7
2. Connect the routerboard to the notebook computer via serial, and establish a serial communication session with
the RouterBoard. Serial configuration example in in the Serial console manual
3. Run the NetInstall program on your notebook computer.
4. Press the NetInstall "Net Booting" button, enable the Boot Server, and enter a valid, usable IP address (within the
same subnet of the IP address of the Notebook) that the NetInstall program will assign to the RouterBoard to
enable communication with the Notebook computer. For example: 10.1.1.5/24
5. Set the RouterBoard BIOS to boot from the Ethernet interface.
Configuring RouterBOARD
Configuring RouterBOARD without COM port
To boot RouterBOARD withtout COM port from Network, you can use reset button. Consult RouterBOARD.com and
specific RouterBOARD User Guide to find reset button location and usage instructions. For example RB751U-2HnD
etherboot instructions,
RouterBOARD 751U-2HnD RouterBOOT reset button (RES, front panel) has two functions to reset RouterOS configuration
and boot it from Etherboot: - Connect Netinstall PC to "ether1" port and hold this button during boot time longer, until LED
turns off, then release it to make the RouterBOARD look for Netinstall servers.
As well Etherboot can be configured by RouterOS (when you have access to it),
s y s t e m  r o u t e r b o a r d  s e t t i n g s  s e t  b o o t - d e v i c e = t r y - e t h e r n e t - o n c e - t h e n - n a n d
Configuring RouterBOARD with COM port
To access Routerboard BIOS configuration: reboot the Routerboard while observing the activity on the Serial Console. You
will see the following prompt on the Serial Console “Press any key within 2 seconds to enter setup” indicating that you
have a 1 or 2 second window of time when pressing any key will give you access to Routerboard BIOS configuration
options.
(press any key when prompted):
You will see the following list of available BIOS Configuration commands. To set up the boot device, press the 'o' key:
W h a t  d o  y o u  w a n t  t o  c o n f i g u r e ?
  d  -  b o o t  d e l a y
  k  -  b o o t  k e y
  s  -  s e r i a l  c o n s o l e
  l  -  d e b u g  l e v e l
  o  -  b o o t  d e v i c e
  b  -  b e e p  o n  b o o t
  v  -  v g a  t o  s e r i a l
  t  -  a t a  t r a n s l a t i o n
  p  -  m e m o r y  s e t t i n g s
  m  -  m e m o r y  t e s t
  u  -  c p u  m o d e
  f  -  p c i  b a c k - o f f
  r  -  r e s e t  c o n f i g u r a t i o n
  g  -  b i o s  u p g r a d e  t h r o u g h  s e r i a l  p o r t
  c  -  b i o s  l i c e n s e  i n f o r m a t i o n
  x  -  e x i t  s e t u p
Next Selection: Press the 'e' key to make the RouterBoard to boot from Ethernet interface:
S e l e c t  b o o t  d e v i c e :
*  i  -  I D E
  e  -  E t h e r b o o t
  1  -  E t h e r b o o t  ( t i m e o u t  1 5 s ) ,  I D E
  2  -  E t h e r b o o t  ( t i m e o u t   1 m ) ,  I D E
  3  -  E t h e r b o o t  ( t i m e o u t   5 m ) ,  I D E
  4  -  E t h e r b o o t  ( t i m e o u t  3 0 m ) ,  I D E
  5  -  I D E ,  t r y  E t h e r b o o t  f i r s t  o n  n e x t  b o o t  ( 1 5 s )
  6  -  I D E ,  t r y  E t h e r b o o t  f i r s t  o n  n e x t  b o o t  ( 1 m )
  7  -  I D E ,  t r y  E t h e r b o o t  f i r s t  o n  n e x t  b o o t  ( 5 m )
  8  -  I D E ,  t r y  E t h e r b o o t  f i r s t  o n  n e x t  b o o t  ( 3 0 m )

