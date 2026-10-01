---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-4-administration-guide-576158-configuring-firewall-authen-52aa3be9-1
title: "Configuring FSSO firewall authentication"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-4-administration-guide-576158-configuring-firewall-authen-52aa3be9.md
source_anchor: ""
source_lines: [1, 202]
sha256: e06a44f161869795e60a8a9ce189d5c946ccd72c3d9ac2e585c136aa2f651727
---

# Configuring FSSO firewall authentication

# Configuring FSSO firewall authentication

In this example, a Windows network is connected to the FortiGate on port 2, and another LAN, Network_1, is connected on port 3.


All Windows network users authenticate when they log on to their network. Engineering and Sales groups members can access the Internet without reentering their authentication credentials. The example assumes that you have already installed and configured FSSO on the domain controller.

LAN users who belong to the Internet_users group can access the Internet after entering their username and password. The example shows two users: User1, authenticated by a password stored in FortiOS; and User 2, authenticated on an external authentication server. Both users are local users since you create the user accounts in FortiOS.

User1 is authenticated by a password stored in FortiOS.

###### To create a locally authenticated user account in the GUI:

1. 
                                                    Go to *User & Authentication > User Definition* and click*Create New* .
2. 
                                                    Configure the following settings: User Type Local User User Name User1 Password hardtoguess1@@1 User Account Status Enabled
3. 
                                                    Click *Submit* .

###### To create a locally authenticated user account in the CLI:

config user local

edit user1

set type password

set passwd hardtoguess1@@1

next

end

You must first configure FortiOS to access the external authentication server, then create the user account.

###### To create a RADIUS-authenticated user account in the GUI:

1. 
                                                    Go to *User & Authentication > RADIUS Servers* and click*Create New* .
2. 
                                                    Configure the following settings: Name OurRADIUSsrv Authentication method Default Primary Server IP/Name 10.11.101.15 Secret OurSecret
3. 
                                                    Click *OK* .
4. 
                                                    Go to *User & Authentication > User Definition* and click*Create New* .
5. 
                                                    Configure the following settings: User Type Remote RADIUS User User Name User2 RADIUS Server OurRADIUSsrv User Account Status Enabled
6. 
                                                    Click *Submit* .

###### To create a RADIUS-authenticated user account in the CLI:

config user radius

edit OurRADIUSsrv

set server 10.11.102.15

set secret OurSecret

set auth-type auto

next

end

config user local

edit User2

set name User2

set type radius

set radius-server OurRADIUSsrv

next

end

This example assumes that you have already set up FSSO on the Windows network and that it used advanced mode, meaning that it uses LDAP to access user group information. You must do the following:

- 
                                                    Configure LDAP access to the Windows AD global catalog
- 
                                                    Specify the collector agent that sends user log in information to FortiOS
- 
                                                    Select Windows user groups to monitor
- 
                                                    Select and add the Engineering and Sales groups to an FSSO user group

###### To create an FSSO user group in the GUI:

1. 
                                                    Configure LDAP for FSSO: 
  1. 
                                                            Go to *User & Authentication > LDAP Servers* and click*Create New* .
  2. 
                                                            Configure the following settings: Name ADserver Server Name / IP 10.11.101.160 Distinguished Name dc=office,dc=example,dc=com Bind Type Regular Username cn=FSSO_Admin,cn=users,dc=office,dc=example,dc=com Password Enter a secure password.
  3. 
                                                            Leave other fields as-is. Click *OK* .
2. 
                                                            
3. 
                                                    Specify the collector agent for FSSO; 
  1. 
                                                            Go to *Security Fabric > External Connectors* and click*Create New* .
  2. 
                                                            Under *Endpoint/Identity* , select*FSSO Agent on Windows AD* .
  3. 
                                                            Configure the following settings: Name Enter the Windows AD server name. This name appears in the Windows AD server list when you create user groups. In this example, the name is WinGroups. Server IP/Name Enter the IP address or name of the server where the agent is installed. The maximum name length is 63 characters. In this example, the IP address is 10.11.101.160. Password Enter the password of the server where the agent is installed. You only need to enter a password for the collector agent if you configured the agent to require authenticated access. If the TCP port used for FSSO is not the default, 8000, you can run the `config user fsso` command to change the setting in the CLI.Collector Agent AD access mode Advanced LDAP Server Select the previously configured LDAP server. In this example, it is ADserver. User/Groups/Organization Units Select the users, groups, and OUs to monitor.
  4. 
                                                            Click *OK* .
4. 
                                                            
5. 
                                                    Create the FSSO_Internet_users user group: 
  1. 
                                                            Go to *User & Authentication > User Groups* and click*Create New* .
  2. 
                                                            Configure the following settings: Name FSSO_Internet_users Type Fortinet Single Sign-On (FSSO) Members Engineering, Sales
  3. 
                                                            Click *OK* .
6. 
                                                            

###### To create an FSSO user group in the CLI:

config user ldap

edit "ADserver"

set server "10.11.101.160"

set dn "cn=users,dc=office,dc=example,dc=com"

set type regular

set username "cn=administrator,cn=users,dc=office,dc=example,dc=com"

set password set_a_secure_password

next

end

config user fsso

edit "WinGroups"

set ldap-server "ADserver"

set password ********

set server "10.11.101.160"

next

end

config user group

edit FSSO_Internet_users

set group-type fsso-service

set member CN=Engineering,cn=users,dc=office,dc=example,dc=com CN=Sales,cn=users,dc=office,dc=example,dc=com

next

end

This example shows a firewall user group with only two users. You can add additional members.

###### To create a firewall user group in the GUI:

1. 
                                                    Go to *User & Authentication > User Groups* and click*Create New* .
2. 
                                                    Configure the following settings: Name Internet_users Type Firewall Members User1, User2
3. 
                                                    Click *OK* .

###### To create a firewall user group in the CLI:

config user group

edit Internet_users

set group-type firewall

set member User1 User2

next

end

###### To define policy addresses:

