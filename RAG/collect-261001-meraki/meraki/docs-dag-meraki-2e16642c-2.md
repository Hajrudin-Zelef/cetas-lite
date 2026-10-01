---
id: collect-261001-meraki/meraki/docs-dag-meraki-2e16642c-2
title: "docs-dag-meraki-2e16642c"
domain: meraki
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/docs-dag-meraki-2e16642c.md
source_anchor: ""
source_lines: [70, 120]
sha256: 5a6cc0b062e50e8505975d9205df0ceae94930d29d3a1fd79035a2034da7c1cd
---

# docs-dag-meraki-2e16642c

| Universal Prompt | Traditional Prompt | 
Read the Universal Prompt Update Guide for more information about the update process and the new login experience for users.
Duo Access Gateway Universal Prompt support is available to Duo Federal customers only starting with version 2.0.0. Activating Universal Prompt for DAG applications requires the following steps:
- Enable Use frameless in the local DAG server admin console on the **General** page.
- Complete the federation steps in this document for your SAML app.
- Log in as an end user and complete Duo authentication to the SAML application now federated with your Duo Access Gateway server. This first authentication shows the traditional Duo prompt in a redirect instead of an iframe.
- You then activate the Universal Prompt for all users of that specific DAG SAML application.
The "Universal Prompt" section of your DAG SAML app shows the status as "Update Required" when you first create it. The status will change after you log in to the application and complete Duo authentication to the SAML application now federated with your Duo Access Gateway server.
Add the Meraki Application to Duo Access Gateway
Before you do this, verify that you updated the "Attributes" list for your Duo Access Gateway authentication source as specified here.
- 
Return to the Applications page of the DAG admin console session.
- 
Click the Choose File button in the "Add Application" section of the page and locate the Meraki SAML application JSON file you downloaded from the Duo Admin Panel earlier. Click the Upload button after selecting the JSON configuration file.
- 
The Meraki SAML application is added. Copy the Login URL for the Meraki application (it looks like https://yourserver.example.com/dag/saml2/idp/SSOService.php?spentityid=https://dashboard.meraki.com). That is the URL you will use to log in to the Meraki dashboard via DAG.
Verify SSO
You can log in to Meraki with IdP-initiated SSO using the Login URL provided by the DAG server when you created the Meraki application in the DAG admin console (https://yourserver.example.com/dag/saml2/idp/SSOService.php?spentityid=https://dashboard.meraki.com).
Enter your primary directory username and password on the Duo Access Gateway login page, approve Duo two-factor authentication, and get redirected to the Meraki site after authenticating.
Congratulations! Your Meraki users now authenticate using Duo Access Gateway.
Activate Universal Prompt
Once you authenticate to your newly-federated SAML application, the "Universal Prompt" section of the application's details page in the Admin Panel reflects this status as "Ready to activate", with these activation control options:
- Show traditional prompt: (Default) Your users experience Duo's traditional prompt via redirect when logging in to this application.
- Show new Universal Prompt: Your users experience the Universal Prompt via redirect when logging in to this application.
Enable the Universal Prompt experience by selecting Show new Universal Prompt, and then scrolling to the bottom of the page to click Save.
Once you activate the Universal Prompt, the application's Universal Prompt status shows "Activation Complete" here and on the Universal Prompt Update Progress report.
The next time your users log in to this application, they will see the new Universal Prompt experience instead of the traditional Duo prompt.
If you plan to permit use of WebAuthn authentication methods (security keys, U2F tokens, or Touch ID) in the traditional Duo Prompt, Duo recommends configuring allowed hostnames for this application and any others that show the inline Duo Prompt before onboarding your end-users.
The Duo Universal Prompt has built-in protection from unauthorized domains so this setting does not apply.
Add Ownership and Risk Information
Go to the "Ownership and Risk" section of the application's page in the Duo Admin Panel to assign application owners and classify the application's risk level. Cisco Identity Intelligence automatically imports this information to populate relevant fields for Duo Advantage and Premier customers. Duo Essentials and Duo Federal plans exclude Cisco Identity Intelligence features.
You may set any of the following:
- Technical Owner: Search for and assign Duo users responsible for technical configuration and maintenance of this application.
- Business Owner: Search for and assign Duo users responsible for business decisions and access approvals related to this application.
- Application Sensitivity: Select this application's risk level from the drop-down list. Default: Not Set.
- Compliance Requirements: Select any applicable regulatory frameworks for this application (SOX (Sarbanes-Oxley), HIPAA, PCI-DSS, etc.) from the list.
Scroll to the bottom of the page and click Save to apply your changes.
Grant Access to Users
If you did not already grant user access to the Duo users you want to use this application be sure to do that before inviting or requiring them to log in with Duo.
Meraki does not support SP-initiated SSO login at this time.
Microsoft AD FS
Microsoft's Active Directory Federation Services (AD FS) is a popular choice for SSO because it easily integrates with the AD identity store many organizations already have deployed. Duo's support for cloud applications and SSO drops in to an existing AD FS installation to provide secondary authentication after a user passes primary authentication (successful Active Directory logon).
If you don't already have AD federation running the first step is to install and configure Microsoft AD FS in your organization. Deployment Guides for AD FS versions 2.1, and 3.0/4.0 are available from Microsoft.
Once your AD FS services are up and running, the second step is to configure the SSO partnership between your AD FS service and the external cloud resource, in this case Meraki. Learn more about configuring Meraki SSO with AD FS at the Meraki Support site.
After you have successfully configured and tested AD FS SSO login to Meraki using your AD domain credentials, you can then install the Duo AD FS integration. AD FS protection is included with Duo Essentials, Duo Advantage, and Duo Premier plans.
With the Duo integration for AD FS installed, users pass primary authentication to the AD FS service as usual. Once primary authentication succeeds, users are forwarded to the Duo service for secondary authentication. After approving logon using one of Duo's authentication methods, the user is fully logged in to Meraki.
Other Identity Partners
Using a third-party SSO provider for cloud application access? Duo partners with leading cloud SSO providers like Okta and OneLogin to secure access with our strong and flexible authentication platform.
You can also use Duo two-factor authentication with CAS and Shibboleth on-premises IdPs.
Troubleshooting
Need some help? Try searching our Knowledge Base articles or Community discussions. For further assistance, contact Support.
