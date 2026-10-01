---
id: collect-261001-cisco/cisco/cisco-basic-configuration-ciscozine
title: "cisco-basic-configuration-ciscozine"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "voice"]
source: docs/RAG/collect-261001-cisco/cisco-basic-configuration-ciscozine.md
source_anchor: ""
source_lines: [1, 198]
sha256: 5128ff730aa4179ae32d0bd7b3fced4a12f6ea5d77a99ce3d7edcfe55d272168
---

# cisco-basic-configuration-ciscozine

There are some “standards” steps used for basic configuration on your Cisco router/switch:

- Define the hostname
- Assign the privileged level
- Secure console port
- Secure VTY lines
- Encrypt the passwords


**Define hostname** It is very useful define the name of your Cisco switch/router. To do it:

Enable administrative privilege

Router>en

Enter the configuration mode:

Router#configure terminal

Define the hostname (the name router is changed to Ciscozine)

Router(config)#hostname Ciscozine
Ciscozine(config)#

**Assign the privileged level secret** The enable secret password allows you to enter enable mode (sometimes referred to as privileged mode or privileged-EXEC mode). To set a secret password:

Enable administrative privilege

Ciscozine>en

Enter the configuration mode:

Ciscozine#configure terminal

Define the privileged-EXEC mode password; remember to use a strong password (combine letters, numbers, nd symbols). The enable secret command uses a one-way encryption hash based on Message Digest 5 (MD5)

Ciscozine(config)#enable secret C1sc0z$n3

**Secure console port** For security reasons, you 

**must**define a console password:

Enable administrative privilege

Ciscozine>en

Enter the configuration mode:

Ciscozine#configure terminal

Enter in the console configuration

Ciscozine(config)#line console 0

Define your console password

Ciscozine(config-line)#password g4S$%A2

Enable the password defined

Ciscozine(config-line)#login

Use the “logging synchronous” command, so the messages appear only after you press Return

Ciscozine(config-line)#logging synchronous

**Secure VTY lines** For security reasons, you 

**must**define a vty line password. The vty lines defines the password needed for a remote access (telnet, ssh, …)

Enable administrative privilege

Ciscozine>en

Enter the configuration mode:

Ciscozine#configure terminal

Enter the vty line configuration. The “0 4” define the “first last” line number; in this case there are 5 allowed sessions.

Ciscozine(config)#line vty 0 4

Define your vty lines password

Ciscozine(config-line)#password g%S$%Aa3

Enable the password

Ciscozine(config-line)#login

Use the “logging synchronous” command, so the messages appear only after you press Return

Ciscozine(config-line)#logging synchronous

**Encrypt the passwords** Before you type the encrypt password command, the passwords stored in the running-config are not encrypted. To “encrypt” it:

Enable administrative privilege

Ciscozine>en

Enter the configuration mode:

Ciscozine#configure terminal

The actual encryption process occurs when the current configuration is written or when a password is configured. Password encryption (it uses Vigenere algorithm) is applied to all passwords, including username passwords, authentication key passwords, the privileged command password, console and virtual terminal line access passwords, and Border Gateway Protocol neighbor passwords. This command is primarily useful for keeping unauthorized individuals from viewing your password in your configuration file.

Ciscozine(config)#service password-encryption

**REMEMBER**:The Vigenere algorithm is used to obfuscate the passwords (not really encrypt them as there is no encryption key) in order to prevent “shoulder surfing” from exposing passwords to someone who briefly looks at a running configuration. If, however, someone gets hold of the configuration they can easily retrieve the passwords using the reverse translation of the Vigenere algorithm.

You could check your current configuration with the command “show running-config”.

Ciscozine#show running-config
Building configuration...
Current configuration : 782 bytes
!
version 12.2
service timestamps debug datetime msec
service timestamps log datetime msec
service password-encryption
!
hostname Ciscozine
!
enable secret 5 $1$8PBD$5nv39GC91koQbOcxfVFIe/
!
memory-size iomem 15
ip subnet-zero
!
!
!
ip audit notify log
ip audit po max-events 100
!
!
!
voice call carrier capacity active
!
!
!
!
!
!
!
!
!
mta receive maximum-recipients 0
!
!
!
!
interface FastEthernet0/0
 no ip address
 shutdown
 duplex auto
 speed auto
!
interface FastEthernet0/1
 no ip address
 shutdown
 duplex auto
 speed auto
!
ip classless
ip http server
!
!
!
call rsvp-sync
!
!
mgcp profile default
!
dial-peer cor custom
!
!
!
!
!
line con 0
 password 7 1410463848410B79
 logging synchronous
 login
line aux 0
line vty 0 4
 password 7 130252214F49252B78
 login
!
!
end
Ciscozine#

This is the basic configuration and… **remember to save the configuration** with:

Ciscozine#copy running-config startup-config

or

Ciscozine#write
