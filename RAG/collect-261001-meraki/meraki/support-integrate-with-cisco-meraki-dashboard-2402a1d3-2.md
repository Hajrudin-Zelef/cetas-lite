---
id: collect-261001-meraki/meraki/support-integrate-with-cisco-meraki-dashboard-2402a1d3-2
title: "support-integrate-with-cisco-meraki-dashboard-2402a1d3"
domain: meraki
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/support-integrate-with-cisco-meraki-dashboard-2402a1d3.md
source_anchor: ""
source_lines: [108, 178]
sha256: 32d88d4ab33963aea3bbbc025483fcfcd3882564bb884a54145841e4aa00dac3
---

# support-integrate-with-cisco-meraki-dashboard-2402a1d3

- In the Attributes section, confirm the JumpCloud Attribute Name matches the attribute used in Meraki Dashboard.
This will be either Email or Username.
- Under Service Provider Attribute Name, enter a value for the role attribute defined in the previous section.
A role attribute must be passed in the SAML token/assertion, specifically 'https://dashboard.meraki.com/saml/attributes/role'. This must match one of the Roles defined on the Organization > Administrators page. Dashboard will only accept one role attribute. If multiple roles or group memberships are provided, the first attribute matched will be used. If 'MemberOf' and 'role' attributes are both specified, 'MemberOf' will be prioritized.
- Under Constant Attributes, select OFF if you do not want to use JIT.
- Click Save.
Using JIT Provisioning
Additional attributes are required to use JIT provisioning. JIT required attributes are prepopulated and are on by default to enable JIT provisioning. You can’t edit the JIT required service provider attributes. You can customize the JumpCloud attribute name and the constant value for JIT required attributes. Toggle off the attributes to opt out of sending the attributes in the SAML assertion.
To complete the provisioning process
- Authorize a user’s access to the application in JumpCloud.
- Have the user log in to the application using SSO. The SAML assertion passes from JumpCloud to the service provider, and gives the service provider the information it needs to create the user account.
Authorizing SSO Application Access
Users are implicitly denied access to SSO Applications. After you connect an application to JumpCloud, you need to authorize user access to that application. You can authorize user access from the Applications, Users List or User Groups page.
To authorize user access from the SSO Application’s page
- Log in to the JumpCloud Admin Portal.
If your data is stored outside of the US, check which login URL you should be using depending on your region. If your organization uses LDAP, RADIUS, or requires firewall allow list configuration, the Fully Qualified Domain Names (FQDNs) will also be region specific. See JumpCloud Data Centers for the URLs, FQDNs, and IP addresses.
- Go to Access > SSO Applications, then select the application to which you want to authorize user access.
- Select the User Groups tab. If you need to create a new group of users, see Get Started: User Groups.
- Select the check box next to the desired group of users to which you want to give access.
- Click Save.
To learn how to authorize user access from the Users or User Groups pages, see Authorize Users to an SSO Application.
Validating SSO user authentication workflow(s)
Check your SP's documentation to ensure that both workflows are supported.
IdP-initiated user workflow
- Access the JumpCloud User Console
- Go to Applications and click an application tile to launch it
- JumpCloud asserts the user's identity to the SP and is authenticated without the user having to log in to the application
SP-initiated user workflow
- Go to the SP application login - generally, there is either a special link or an adaptive username field that detects the user is authenticated through SSO
This varies by SP.
- Login redirects the user to JumpCloud where the user enters their JumpCloud credentials
- After the user is logged in successfully, they are redirected back to the SP and automatically logged in
See Additional User Experience Considerations when setting up JumpCloud SSO.
Removing the SSO Integration
These are steps for removing the integration in JumpCloud. Consult your SP's documentation for any additional steps needed (like disabling "mandatory SSO login" settings) to remove the integration in the SP. Failure to remove the integration successfully for both the SP and JumpCloud may result in users, including admins, losing access to the application.
If your data is stored outside of the US, check which login URL you should be using depending on your region. If your organization uses LDAP, RADIUS, or requires firewall allow list configuration, the Fully Qualified Domain Names (FQDNs) will also be region specific. See JumpCloud Data Centers for the URLs, FQDNs, and IP addresses.
To deactivate the SSO Integration
- 
Log in to the JumpCloud Admin Portal.
- 
Go to Access > SSO Applications.
- 
Search for the application that you’d like to deactivate and click to open its details panel.
- 
Select the SSOtab.
- 
Scroll to the bottom of the configuration.
- 
Click Deactivate SSO.
- 
Click Save.
- 
If successful, you will receive a confirmation message.
To delete the application
- 
Log in to the JumpCloud Admin Portal.
- 
Go to Access > SSO Applications.
- 
Search for the application that you’d like to delete.
- 
Check the box next to the application to select it.
- 
Click Delete.
- 
Enter the number of the applications you are deleting
- 
Click Delete Application.
- 
If successful, you will see an application deletion confirmation notification.
Was this information helpful?
