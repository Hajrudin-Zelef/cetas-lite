---
id: collect-260926-mikrotik/mikrotik/questions-1102662-mikrotik-how-send-the-mac-ap-to-external-captive-7d800983
title: "questions-1102662-mikrotik-how-send-the-mac-ap-to-external-captive-7d800983"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-1102662-mikrotik-how-send-the-mac-ap-to-external-captive-7d800983.md
source_anchor: ""
source_lines: [1, 19]
sha256: d4c87d8bee6252fd0e72a65bbd28b1305b27ee75d8472cc4ea320d530b72049d
---

# questions-1102662-mikrotik-how-send-the-mac-ap-to-external-captive-7d800983

I set up the MikroTik router hEX RB750Gr in hotspot mode, connected the WapR-2nD access point to it and distribute wifi. Hotspot is configured for an external authorization portal (Captive). Captive portal needs the mac address of this access point to identify the access point. But in the get request (in login.html) to the authorization portal, hotspot sends its data, not the access point data. I could solve this problem possible if in login.html add new variables in which to add the mac address of the access point by script. Please tell me if it is possible to add to login.html own variables?
<form name="redirect" action="https://auth.wifi........" method="get" style="display: none;">
        <input type="hidden" name="user-mac" value="$(mac)">
        <input type="hidden" name="user-ip" value="$(ip)">
        <input type="hidden" name="ap-mac" value="**$(ap_mac)**">      # hier i need custom variable ap_mac
        <input type="hidden" name="nas-id" value="......................">
        <input type="hidden" name="link" value="$(link-login-only)">
        <input type="hidden" name="err" value="$(error)">
        <input type="hidden" name="eor" value="$(error-orig)">
        <input type="hidden" name="type" value="mikrotik" />
        <input type="hidden" name="server-name" value="$(server-name)">
        <input type="hidden" name="server-address" value="$(server-address)">
        <input type="hidden" name="chap-id" value="$(chap-id)">
        <input type="hidden" name="chap-challenge" value="$(chap-challenge)">
        <input type="hidden" name="hostname" value="$(hostname)">
        <input type="hidden" name="ssl-login" value="$(ssl-login)">
        <input type="hidden" name="plain-passwd" value="$(plain-passwd)">
        <input type="submit" value="continue">
    </form>
