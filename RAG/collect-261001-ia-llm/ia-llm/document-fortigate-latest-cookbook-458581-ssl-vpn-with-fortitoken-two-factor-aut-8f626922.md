---
id: collect-261001-ia-llm/ia-llm/document-fortigate-latest-cookbook-458581-ssl-vpn-with-fortitoken-two-factor-aut-8f626922
title: "Set up FortiToken multi-factor authentication"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-ia-llm/document-fortigate-latest-cookbook-458581-ssl-vpn-with-fortitoken-two-factor-aut-8f626922.md
source_anchor: ""
source_lines: [1, 35]
sha256: cb7fce7909078c9893dea58cb9b139929392292947c56ec1f8d38045810a24a5
---

# Set up FortiToken multi-factor authentication

# Set up FortiToken multi-factor authentication

This configuration adds multi-factor authentication (MFA) to the split tunnel configuration (SSL VPN split tunnel for remote user). It uses one of the two free mobile FortiTokens that is already installed on the FortiGate.

###### To configure MFA using the GUI:

1. Configure the user:
  1. Go to *User & Device > User Definition* and edit local user*sslvpnuser1* .
  2. Enter the user's *Email Address* .
  3. Enable *Two-factor Authentication* and select one mobile*Token* from the list,
  4. Enable *Send Activation Code* and select*Email* .
  5. Click *Next* and click*Submit* .
2. Go to 
3. Activate the mobile token:
  1. When a FortiToken is added to user *sslvpnuser1* , an email is sent to the user's email address. Follow the instructions to install your FortiToken mobile application on your device and activate your token.
4. When a FortiToken is added to user 

###### To configure MFA using the CLI:

1. Configure the user:```
config user local
    edit "sslvpnuser1"
        set type password
        set two-factor fortitoken
        set fortitoken <select mobile token for the option list>
        set email-to <user's email address>
        set passwd <user's password>
    next
end
```
2. Activate the mobile token:
  1. When a FortiToken is added to user *sslvpnuser1* , an email is sent to the user's email address. Follow the instructions to install your FortiToken mobile application on your device and activate your token.
3. When a FortiToken is added to user
