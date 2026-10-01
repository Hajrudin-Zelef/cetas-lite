---
id: collect-261001-general-networking/general-networking/platform-management-sm-endpoint-management-install-and-get-started-device-enroll-b49a6d74
title: "platform-management-sm-endpoint-management-install-and-get-started-device-enroll-b49a6d74"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-general-networking/platform-management-sm-endpoint-management-install-and-get-started-device-enroll-b49a6d74.md
source_anchor: ""
source_lines: [1, 70]
sha256: 0f2fa4abe0ac419111c0c0d33297a481318f2f1186e0e103051e147ac92091ba
---

# platform-management-sm-endpoint-management-install-and-get-started-device-enroll-b49a6d74

In large environments, it isn't time efficient to install software on individual PCs one at a time. In order to perform tasks like deploying the Systems Manager agent in bulk, administrators of Windows environments with Active Directory can make use of Active Directory Group Policy Objects to administratively push software out to a large number of devices.


For more information on GPOs, requirements, and any Microsoft or Active Directory questions, please consult Microsoft's documentation for software installation using Active Directory GPO.


This article will cover the steps required in a typical Active Directory environment to push out the System Manager client for Windows to devices across an AD domain.


1. Choose your Systems Manager network from the Network drop-down, then go to **Systems Manager > Add devices > Windows.**
2. Download the Windows installer image by clicking on the **Download** link and saving the .msi file
3. Move the .msi file to __a shared location__ that is accessible by all target devices; in this example, it's a "Software" folder on a file server.

## Agent version 1.0 - 3.0.3

1. On a domain controller, start the **Group Policy Management** tool.
 
2. Navigate to [your_company_name]-Computers, right-click, and choose "Create a GPO in this domain and link it here..."
3. Give the GPO a name, for example "Systems Manager" or "Meraki Systems Manager"
 
4. Under "Group Policy Objects", find the GPO you created in step 3, right-click on it, and choose "Edit..."
 
5. In the window that pops up, navigate to "Computer Configuration" to "Policies" to "Software Installation"; right-click it and choose "New" and "Package..."
 
6. Browse to- or type the full path of the .msi package on the share.
7. In the "Deploy Software" window that pops up, leave the setting option to "Assigned" and click OK

## Agent version 3.1.0+

To deploy Agent v3.1.0+ via GPO, administrators will need to create and distribute a transform file (.mst) with the Agent package.  The transform file defines the SM network environment the Agent will connect to after installation.


Transform files can be created in a variety of different msi editing tools. Steps to create a transform file in Microsoft Orca are outlined below.

#### Create transform file in Orca

1. Open Orca.  Click **File > Open**  in the main toolbar, and open the Agent installer .msi  (SMAgent-*x.x.x* )
2. In the main toolbar, click **Transform > New Transform**
3. Select the **Property** table from the list on the left, then click**Tables > Add Row** in the main toolbar
4. In the Property field enter **ENROLLMENT_CODE**  (case sensitive). In the Value field, enter your network enrollment ID or network enrollment string. If enrollment authentication is enabled, add**ENROLL_TOKEN** as a Property, with your network's bulk enrollment token as a value.
5.  In the main toolbar, click **Transform > Generate Transform** . Name the transform file and save it to a network share folder that is accessible by all target devices

#### Deploy transform file with GPO

1. On a domain controller, start the **Group Policy Management** tool.
 
2. Navigate to [your_company_name]-Computers, right-click, and choose "Create a GPO in this domain and link it here..."
3. Give the GPO a name, for example "Systems Manager" or "Meraki Systems Manager"
 
4. Under "Group Policy Objects", find the GPO you created in step 3, right-click on it, and choose "Edit..."
 
5. In the window that pops up, navigate to "Computer Configuration" to "Policies" to "Software Installation"; right-click it and choose "New" and "Package..."
 
6. Browse to or type the full path of the .msi package on the network share
 
7. In the Deploy Software window, change the setting option to **Advanced**
 
8. In the Properties window, click the **Modifications** tab, then click**Add**
 
9. Select the Transforms (*.mst) file from the network share location, and click Ok

## Initiate Deployment

1. Right-click the entry you created and enable "Enforced"
 
2. Start a shell by clicking on the Start button (then "Run..." on some versions of Windows") and typing "cmd"
 
3. In the cmd shell that pops up, type "gpupdate /force"
 
4. Wait until gpupdate completes. You are done.
