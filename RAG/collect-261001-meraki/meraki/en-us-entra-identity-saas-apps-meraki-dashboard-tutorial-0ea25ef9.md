---
id: collect-261001-meraki/meraki/en-us-entra-identity-saas-apps-meraki-dashboard-tutorial-0ea25ef9
title: "en-us-entra-identity-saas-apps-meraki-dashboard-tutorial-0ea25ef9"
domain: meraki
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["consumer"]
source: docs/RAG/collect-261001-meraki/en-us-entra-identity-saas-apps-meraki-dashboard-tutorial-0ea25ef9.md
source_anchor: ""
source_lines: [1, 60]
sha256: b6fce24c6d77bfc4adeed7d7a8ef21fb3b021b27678ff0003a74830b69cca3be
---

# en-us-entra-identity-saas-apps-meraki-dashboard-tutorial-0ea25ef9

Configure Meraki Dashboard for Single sign-on with Microsoft Entra ID
In this article, you learn how to integrate Meraki Dashboard with Microsoft Entra ID. When you integrate Meraki Dashboard with Microsoft Entra ID, you can:
Control in Microsoft Entra ID who has access to Meraki Dashboard.
Enable your users to be automatically signed-in to Meraki Dashboard with their Microsoft Entra accounts.
Manage your accounts in one central location.
Prerequisites
The scenario outlined in this article assumes that you already have the following prerequisites:
A Microsoft Entra user account with an active subscription. If you don't already have one, you can Create an account for free.
Meraki Dashboard single sign-on (SSO) enabled subscription.
Scenario description
In this article, you configure and test Microsoft Entra SSO in a test environment.
Meraki Dashboard supports IDP initiated SSO.
Note
Identifier of this application is a fixed string value so only one instance can be configured in one tenant.
Adding Meraki Dashboard from the gallery
To configure the integration of Meraki Dashboard into Microsoft Entra ID, you need to add Meraki Dashboard from the gallery to your list of managed SaaS apps.
Configure and test Microsoft Entra SSO for Meraki Dashboard
Configure and test Microsoft Entra SSO with Meraki Dashboard using a test user called B.Simon. For SSO to work, you need to establish a link relationship between a Microsoft Entra user and the related user in Meraki Dashboard.
To configure and test Microsoft Entra SSO with Meraki Dashboard, perform the following steps:
Create Meraki Dashboard Admin Roles - to have a counterpart of B.Simon in Meraki Dashboard that's linked to the Microsoft Entra representation of user.
Test SSO - to verify whether the configuration works.
Browse to Entra ID > Enterprise apps > Meraki Dashboard > Single sign-on.
On the Select a single sign-on method page, select SAML.
On the Set up single sign-on with SAML page, select the edit/pen icon for Basic SAML Configuration to edit the settings.
On the Basic SAML Configuration section, perform the following steps:
In the Reply URL textbox, type a URL using the following pattern:
https://n27.meraki.com/saml/login/m9ZEgb/< UNIQUE ID >
Note
The Reply URL value isn't real. Update this value with the actual Reply URL value, which is explained later in the article.
Select the Save button.
Meraki Dashboard application expects the SAML assertions in a specific format, which requires you to add custom attribute mappings to your SAML token attributes configuration. The following screenshot shows the list of default attributes.
In addition to above, Meraki Dashboard application expects few more attributes to be passed back in SAML response which are shown below. These attributes are also pre populated but you can review them as per your requirements.
To understand how to configure roles in Microsoft Entra ID, see here.
In the SAML Signing Certificate section, select Edit button to open SAML Signing Certificate dialog.
In the SAML Signing Certificate section, copy the Thumbprint Value and save it on your computer. This value needs to be converted to include colons in order for the Meraki dashboard to understand it . For example, if the thumbprint from Azure is C2569F50A4AAEDBB8E it will need to be changed to C2:56:9F:50:A4:AA:ED:BB:8E to use it later in Meraki Dashboard.
On the Set up Meraki Dashboard section, copy the Logout URL value and save it on your computer.
In a different web browser window, sign in to your Meraki Dashboard company site as an administrator
Navigate to Organization > Settings.
Under Authentication, change SAML SSO to SAML SSO enabled.
Select Add a SAML IdP.
Paste the converted Thumbprint Value, which you have copied and converted in specified format as mentioned in step 9 of previous section into X.590 cert SHA1 fingerprint textbox. Then select Save. After saving, the Consumer URL will show up. Copy Consumer URL value and paste this into Reply URL textbox in the Basic SAML Configuration Section.
Create Meraki Dashboard Admin Roles
In a different web browser window, sign into meraki dashboard as an administrator.
Navigate to Organization > Administrators.
In the SAML administrator roles section, select the Add SAML role button.
Enter the Role meraki_full_admin, mark Organization access as Full and select Create role. Repeat the process for meraki_readonly_admin, this time mark Organization access as Read-only box.
Follow the below steps to map the Meraki Dashboard roles to Microsoft Entra SAML roles:
a. In the Azure portal, select App Registrations.
b. Select All Applications and select Meraki Dashboard.
c. Select App Roles and select Create App role.
d. Enter the Display name as Meraki Full Admin.
e. Select Allowed Members as Users/Groups.
f. Enter the Value as meraki_full_admin.
g. Enter the Description as Meraki Full Admin.
h. Select Save.
Test SSO
In this section, you test your Microsoft Entra single sign-on configuration with following options.
Select Test this application, and you should be automatically signed in to the Meraki Dashboard for which you set up the SSO
You can use Microsoft My Apps. When you select the Meraki Dashboard tile in the My Apps, you should be automatically signed in to the Meraki Dashboard for which you set up the SSO. For more information about the My Apps, see Introduction to the My Apps.
Learn to create an initial Azure Active Directory configuration to ensure all the identity solutions available in Azure are ready to use. This module explores how to build and configure an Azure AD system.
