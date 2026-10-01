---
id: collect-261001-general-networking/general-networking/anyconnect-onelogin-saml-configuration
title: "anyconnect-onelogin-saml-configuration"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["consumer"]
source: docs/RAG/collect-261001-general-networking/anyconnect-onelogin-saml-configuration.md
source_anchor: ""
source_lines: [1, 54]
sha256: 6fd28d6d3f0c0ed53869de0d2469a259c89176b4631ec0779c5b6b6546dd0b6a
---

# anyconnect-onelogin-saml-configuration

### **AnyConnect VPN Onelogin SAML Configuration** 

**AnyConnect VPN Onelogin SAML Configuration**

This document highlights how to setup authentication with Onelogin using SAML for AnyConnect VPN on the MX Appliance. SAML is an XML-based framework for exchanging authentication and authorization data between security domains. It creates a circle of trust between the user, a Service Provider (SP), and an Identity Provider (IdP) which allows the user to sign in a single time for multiple services.

SAML authentication requires MX firmware version **16.16+ or 17.5+**

For additional information, refer to the __AnyConnect configuration guide.__

**Do not use** **AnyConnect predefined option in Onelogin** (that option only works for the ASA/FTD platforms) for AnyConnect SAML configuration when setting up AnyConnect authentication with the MX Appliance

To set up AnyConnect authentication on the MX with Onelogin, follow the steps below:

**Step 1.** Logon to Onelogin and click on 'Administration'

**Step 2.** Click on Applications → Applications

**Step 3.** Click on 'add app'

**Step 4.** In the search field, search for '**test connector**', and choose '**SAML Test Connector (Advanced)**' for **SAML 2.0** (not 1.1)

**Step 5.** Create a **SAML Test Connector (SP) or (Advanced)** and fill out an appropriate name e.g. Meraki AnyConnect VPN

**Step 6.** *If my AnyConnect Server URL is* *"**vtk-qpjgjhmpdh.dynamic-m.com",* *Onelogin should be configured as follows**:*

**Audience (EntityID)**: https://***vtk-qpjgjhmpdh.dynamic-m.com***/saml/sp/metadata/SAML

**Recipient:** https://***vtk-qpjgjhmpdh.dynamic-m.com***/saml/sp/metadata/SAML

**ACS (Consumer) URL Validator**: https:\/\/**vtk-*****qpjgjhmpdh***\***.dynamic-m***\.com.*.

**ACS (Consumer) URL:** https://**vtk-*****qpjgjhmpdh.dynamic-m.com***/saml/sp/acs

**Step 7.** Click on SSO (left pane), then **More Actions** (upper right) => **SAML Metadata**, and download the SAML metadata.

**Step 8.** Configure your **AnyConnect Server** on the Meraki Dashboard 

Set Authentication Type to **SAML**


Configure your **AnyConnect URL** - https://vtk-qpjgjhmpdh.dynamic-m.com 

(add “:port” to the end of the URL if using a port other than the default port 443)

Please ensure your AnyConnect URL starts with "https://"

 

- 
    Upload the  **SAML Metadata file** downloaded in**step 7**  above
 
- 
    Save your configuration.
