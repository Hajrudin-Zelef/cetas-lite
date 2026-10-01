---
id: collect-261001-general-networking/general-networking/anyconnect-authentication-methods-1
title: "anyconnect-authentication-methods"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/anyconnect-authentication-methods.md
source_anchor: ""
source_lines: [1, 147]
sha256: 7c78e11c717a12db65355584f80a12cecf13a6c50c5429fc35a69dca48e44144
---

# anyconnect-authentication-methods

**AnyConnect supports authentication with any of the following:**

For more details on AnyConnect configuration, refer to the **AnyConnect configuration guide.**

### **SAML Authentication**

SAML is an XML-based framework for exchanging authentication and authorization data between security domains. It creates a circle of trust between the user, a Service Provider (SP), and an Identity Provider (IdP) which allows the user to sign in a single time for multiple services. SAML can be used to authenticate with Identity Providers and MFA solutions such as, DUO, Okta, Onelogin, Entra ID.

SAML authentication requires MX firmware version **16.16+ or 17.5+.** For troubleshooting, see the **SAML Troubleshooting guide**

AnyConnect with SAML only supports **SP initiated SSO flow** and uses Redirect & POST during the exchange with the IdP

**Setting up AnyConnect Authentication with DUO**

**Setting up AnyConnect Authentication with Okta**

**Setting up AnyConnect Authentication with Onelogin**

**Setting up AnyConnect Authentication with Azure AD**

To configure AnyConnect on the MX Appliance to authenticate with DUO via SAML, see below. The directions below do not include configuration of an authentication source, which is a requirement if using DUO as an Identity Provider.

To set up SAML authentication, you need a **Service Provider (e.g. MX running AnyConnect), Identity Provider - DUO and a User.**

 

1. 
    **Configure your Identity Provider - IdP (DUO)**

- 
    Select configuration of a new Generic/Custom Application, ( **do not use AnyConnect presets in DUO for MX configuration** )
- 
    Configure only the Entity ID and ACS URL as follows:
- 
    ***e.g., If my AnyConnect Server hostname is "https://vtk-qpjgjhmpdh.dynamic-m.com", my DUO configuration for Entity ID and ACS URL will be configured as seen below:***
 
 **SP****Entity ID** : https://***vtk-qpjgjhmpdh.dynamic-m.com*** /saml/sp/metadata/SAML
 **ACS URL** : https://***vtk-qpjgjhmpdh.dynamic-m.com*** /saml/sp/acs
 
- 
    Ensure to add “:port” to the hostname when using any port other than the default 443, e.g https:// ***vtk-qpjgjhmpdh.dynamic-m.com*****:8443** , when referencing the AnyConnect Server in your configuration

After configuration, your Identity Provider will provide you with a SAML metadata file. e.g., seen from DUO IdP below.



1. 
    **Configure your AnyConnect Server on the MX**

- 
    Set authentication method to SAML 
- 
    Configure your AnyConnect URL - https:// ***vtk-qpjgjhmpdh.dynamic-m.com*** (add “:port” to the end of the hostname if using a port other than 443)
 Please ensure your AnyConnect URL starts with **https://**
 
- 
    Upload the SAML metadata XML file provided by your Identity Provider to the MX. 
 
- 
    Once your SAML metadata XML file is uploaded, the SAML metadata file section will go from "No File Found" to "SAML metadata File Uploaded". 
 

Save your configuration and attempt to connect to the VPN to verify configuration.

### **Certificate-based authentication with Username & password**

The AnyConnect server on the MX supports client certificate authentication as an authentication factor. If certificate authentication is enabled, the AnyConnect server will use the uploaded trusted CA certificate to validate authenticating clients before requesting for the users' credentials.


With certificate authentication, the administrator uploads a **.pem** or **.crt** file of the **issuing CA certificate** to the MX, and uploads a certificate signed by the same issuing CA to the **end user's device**.


On a Windows machine**, run MMC,** add **Certificates** Snap-in, navigate to **Personal > Certificates** folder, and import or request a new certificate.

This will enable only devices that have a certificate signed by the Root CA to successfully authenticate to VPN. A common use case is for filtering non-corporate devices from authenticating to the VPN.


Once certificate authentication succeeds, users must input credentials. If certificate authentication fails, the AnyConnect client will report certificate validation failure.

Certificate-only authentication is currently in beta, and Meraki Support is not currently able to assist with any issues related to it. Please request access here or please await the general availability release.

*Certificate Store Override:* Allows an administrator to direct AnyConnect to utilize certificates in the Windows machine's Local System certificate store for client certificate authentication. Certificate Store Override only applies to SSL, where the connection is initiated, by default, by the UI process. When using IPsec/IKEv2, this feature in the AnyConnect Profile is not applicable.

You must have a predeployed profile with this option enabled in order to connect with Windows using a machine certificate. If this profile does not exist on a Windows device prior to connection, the certificate is not accessible in the machine store, and the connection fails.

**Note:** Wildcard certificates are not supported.

### **Meraki Cloud Authentication**

Use this option if an Active Directory or RADIUS server is not available, or if VPN users should be managed via the Meraki Cloud. To add or remove users, use the **User Management** section at the bottom of the page. Add a user by clicking "Add new user" and entering the following information:

- 
    **Name** : Enter the user's name.
- 
    **Email** : Enter the user's email address.
- 
    **Password** : Enter a password for the user or click "Generate" to automatically generate a password.
- 
    **Authorized** : Select whether this user is authorized to use the client VPN.

To edit an existing user, click on the user under the **User Management** section. To delete a user, click the **X** next to the user on the right side of the user list.

When using Meraki-hosted authentication, the user's email address is the username that is used for authentication.

### 

**RADIUS Authentication**

Use this option to authenticate users on a RADIUS server. Click "**+RADIUS server**" to configure the server(s) to use. Enter the IP address of the RADIUS server, the port to be used for RADIUS communication, and the shared secret for the RADIUS server.


**Add MX security appliance as a RADIUS client on the NPS server.**

In order for the MX to act as an authenticator for RADIUS, it must be added as a client on NPS.

1. 
    Open the NPS server console by going to **Start** >**Programs** >**Administrative Tools** >**Network Policy Server** .
2. 
    In the left-side pane, expand the RADIUS Clients and Servers option.
3. 
    Right-click the RADIUS Clients option and select New.
4. 
    Enter a Friendly Name for the MX security appliance or Z teleworker gateway RADIUS client.
5. 
    Enter the IP address of your MX security appliance or Z teleworker gateway. This IP will differ depending on where the RADIUS server is located: 
  - 
        On a local subnet: use the IP address of the MX/Z on the subnet shared with the RADIUS server
  - 
        Over a static route: use the IP address of the MX/Z on the subnet shared with the next hop
  - 
        Over VPN: use the IP address of the MX/Z on the highest-numbered VLAN in VPN
6. 
        
7. 
    Create and enter a RADIUS Shared Secret (make note of this secret, you will need to add this to the dashboard).

**Note**: Currently only ASCII characters are supported for RADIUS shared secrets, Unicode characters are not supported.

1. 
    Press OK when finished. 

For additional information or troubleshooting assistance, please refer to Microsoft documentation on __RADIUS clients__.

 

**Configure a RADIUS Connection Request**

