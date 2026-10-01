---
id: collect-261001-general-networking/general-networking/anyconnect-vpn-okta-saml-configuration
title: "anyconnect-vpn-okta-saml-configuration"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/anyconnect-vpn-okta-saml-configuration.md
source_anchor: ""
source_lines: [1, 53]
sha256: 5f359f4c0608b4fd8da1d52fdc939b9fc88ebfe59aec544ce744cd6db9d2e4b0
---

# anyconnect-vpn-okta-saml-configuration

### **AnyConnect VPN Okta SAML Configuration**

This document highlights how to setup authentication with Okta using SAML for AnyConnect VPN on the MX Appliance. SAML is an XML-based framework for exchanging authentication and authorization data between security domains. It creates a circle of trust between the user, a Service Provider (SP), and an Identity Provider (IdP) which allows the user to sign in a single time for multiple services.

SAML authentication requires MX firmware version **16.16+ or 17.5+**

For additional information, refer to the __AnyConnect configuration guide.__

 

**Step 1.** Create an Account with **Okta**

**Step 2.** Go to “Applications” -> "Applications" → “Create App Integration" → "**SAML 2.0**"

**Step 3.** Configure an App name e.g Meraki AnyConnect VPN  => Next.

**Step 5.** General Settings: For "**Sign On Method**" choose "**SAML 2.0**"

**Step 6.**  *If my AnyConnect Server URL is* *"**vtk-qpjgjhmpdh.dynamic-m.com",* *Okta should be configured as follows**:*

**Single sign on URL**: https://***vtk-qpjgjhmpdh.dynamic-m.com***/saml/sp/acs

**Audience URI (SP Entity ID)**: https://***vtk-qpjgjhmpdh.dynamic-m.com***/saml/sp/metadata/SAML

*If using a port other than the default port 443, Okta should be configured as follows:* 

**Single sign on URL**: https://***vtk-qpjgjhmpdh.dynamic-m.com:<PORT>***/saml/sp/acs

**Audience URI (SP Entity ID)**: https://***vtk-qpjgjhmpdh.dynamic-m.com:<PORT>***/saml/sp/metadata/SAML

Leave other advance settings to default

**Step 7.** "Sign On" tab --> In the **SAML Signing Certificates** section, select **Actions** --> Select **View IdP metadata** to open metadata file in a new tab --> Right-click in the newly opened tab and save the metadata page as an XML file

**Step 8.** "Assignments"  tab to assign the Users you have created to the app, If you have not created any user: Click on Directory --> Add Person

**Step 9.** Configure your **AnyConnect Server** on the Meraki Dashboard 

- Set Authentication Type to **SAML**

Configure your **AnyConnect URL** - https://vtk-qpjgjhmpdh.dynamic-m.com 

(add “:port” to the end of the URL if using a port other than the default port 443)

Please ensure your AnyConnect URL starts with "https://"



- 
    Upload the  **SAML Metadata file** downloaded in**step 7**  above
 
- 
    Save your configuration.
