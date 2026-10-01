---
id: collect-261001-general-networking/general-networking/anyconnect-client-download-and-deployment-2
title: "anyconnect-client-download-and-deployment"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/anyconnect-client-download-and-deployment.md
source_anchor: ""
source_lines: [106, 179]
sha256: 2fe81e1d1f2f948685793aee0c83e14882d41d276b0be510d684271d86a6e1ba
---

# anyconnect-client-download-and-deployment

When SBL is installed and enabled, AnyConnect starts before the Windows logon dialog box appears, ensuring users are connected to their corporate infrastructure before logging on. After VPN authentication, the Windows logon dialog appears, and the user logs in as usual.

For more details see **Start Before Logon**

**Note:** Currently Start Before Logon (SBL) is not supported with SAML Authentication.

**Configuration**

1. Install the AnyConnect Start Before Logon Module. There is a separate executable called "sbl-predeploy" file in the AnyConnect for Windows installation folder as shown below.



2. Once the SBL installation is complete, enable Start Before Logon (SBL) in the AnyConnect Profile and push profile to client.

- Open the VPN Profile Editor and choose **Preferences (Part 1)** from the navigation pane.
- Select **Use Start Before Logon** .
- (Optional) To give the remote user control over SBL, select **User Controllable.**
- 
    Click File, **Save** the profile, then**upload** it on the Dashboard > Security & SD-WAN > Client VPN > Cisco Secure Client Settings > "Profile update" toggle and save your configuration. Profiles can also be pushed to users via other methods e.g. via Systems Manager.The profile will get updated on the client after successfully connecting to the VPN or if manually updated on the client. Please note that profiles get overridden on the client if the new profile and the old one on the client share the same file name. **Please note,** the user must reboot the remote computer before SBL takes effect. After a reboot, users can use the network sign-in option to launch and connect to AnyConnect VPN.
 .
 

### **Android**

This is an example of installing the AnyConnect Client and configuring it on an Android device

**Install** 

AnyConnect VPN can be found in the Google Play Store

**Run Cisco Secure Client-AnyConnect**

Click on Connections to open the Connection Selector

**Create New Connection**

Use the + symbol to open the Connection Editor

**Create New Connection**

Enter the Description and Server Address

**Connecting**

Once the connection is created connect by toggling the connection in the AnyConnect app

**Login Prompt**

Enter the credentials to authorize the connection

**Connected**

You should now be connected

### iOS

Open the App Store and in the Search Box, enter Cisco Secure Client or AnyConnect

Tap on Cisco Secure Client

Tap Get, then install the application

The Cloud Icon in the image below indicates the app was installed previously, on first install the "Get" icon will be present.

Open the application and tap **Connections > Add VPN Connection ...**


Enter a friendly Description/Name for the connection, enter the Server address, and save changes

On the application home page toggle the **AnyConnect** **VPN** with the slider icon to enable the VPN and Login with your credentials

If successful we'll observe the connected status on the **Details** section of the home page

Further details regarding iOS setup and use for Secure Connect can be found here.
