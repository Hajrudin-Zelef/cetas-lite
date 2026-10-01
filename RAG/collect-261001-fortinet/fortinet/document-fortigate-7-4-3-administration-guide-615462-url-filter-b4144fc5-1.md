---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-3-administration-guide-615462-url-filter-b4144fc5-1
title: "URL filter"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-3-administration-guide-615462-url-filter-b4144fc5.md
source_anchor: ""
source_lines: [1, 169]
sha256: e3a2d190f771b236e168af5f7c3dbd696e979736ed2489396839ff295b6774c4
---

# URL filter

# URL filter

The URL filter uses specific URLs with patterns containing text and regular expressions so the FortiGate can process the traffic based on the filter action (exempt, block, allow, monitor) and web pages that match the criteria. Once a URL filter is configured, it can be applied to a firewall policy.

The following filter types are available:

| URL filter type | Description | 
|---|---|
| Simple | The FortiGate tries to strictly match the full context. For example, if you enter *www.facebook.com* in the*URL* field, it only matches traffic with*www.facebook.com* . It won't match*facebook.com* or*message.facebook.com* . When the FortiGate finds a match, it performs the selected URL action. | 
| Regular expression/ wildcard | The FortiGate tries to match the pattern based on the rules of regular expressions or wildcards. For example, if you enter **fa** in the*URL* field, it matches all the content that has*fa* such as*www.facebook.com* ,*message.facebook.com* ,*fast.com* , and so on. When the FortiGate finds a match, it performs the selected URL action. | 

For more information, see the URL Filter expressions technical tip in the Knowledge Base.

The following actions are available:

| URL filter action | Description | 
|---|---|
| Exempt | The traffic is allowed to bypass the remaining FortiGuard web filters, web content filters, web script filters, antivirus scanning, and DLP proxy operations. | 
| Block | The FortiGate denies or blocks attempts to access any URL that matches the URL pattern. A replacement message is displayed. | 
| Allow | The traffic is passed to the remaining FortiGuard web filters, web content filters, web script filters, antivirus proxy operations, and DLP proxy operations. If the URL does not appear in the URL list, the traffic is permitted. | 
| Monitor | The traffic is processed the same way as the *Allow* action. For the*Monitor* action, a log message is generated each time a matching traffic pattern is established. | 

The exempt URL filter action can be configured to bypass all or certain security profile operations. This setting can only be configured in the CLI.

If the `action` is set to `exempt`, use `set exempt` to select the security profile operations that exempt URLs skip.

```
config webfilter urlfilter
    edit <id>
        config entries
            edit <id>
                set action exempt
                set exempt {av web-content activex-java-cookie dlp fortiguard range-block pass antiphish all}
            next
        end
    next
end
```
                                            | Option | Description | 
|---|---|
| av | Antivirus scanning | 
| web-content | Web filter content matching | 
| activex-java-cookie | ActiveX, Java, and cookie filtering | 
| dlp | DLP scanning | 
| fortiguard | FortiGuard web filtering | 
| range-block | Range block feature | 
| pass | Pass single connection from all | 
| antiphish | Antiphish credential checking | 
| all | Exempt from all security profiles | 

|  | These exempt options are not visible in the GUI. Setting the URL filter *Action* to*Exempt* will exempt URLs from all security profiles. | 

In the following example, a URL filter will be created to block the facebook.com URL using a wildcard.

### Configuring a URL filter in the GUI

###### To create a URL filter for Facebook:

1. 
                                                    Go to *Security Profiles > Web Filter* and click*Create New* , or edit an existing profile.
2. 
                                                    In the *Static URL Filter* section, enable*URL Filter* .
3. 
                                                    Click *Create New* . The*New URL Filter* pane opens.
4. 
                                                    For URL, enter **facebook.com* , for*Type* , select*Wildcard* , and for*Action* , select*Block* .
5. 
                                                    Click *OK* . The entry appears in the table.
6. 
                                                    Configure the other settings as needed.
7. 
                                                    Click *OK* .

###### To apply the web filter profile to a firewall policy:

1. 
                                                    Go to *Policy & Objects > Firewall Policy* .
2. 
                                                    Edit a policy, or create a new one.
3. 
                                                    In the *Security Profiles* section, enable*Web Filter* and select the profile that you created.
4. 
                                                    Set *SSL Inspection* to*certificate-inspection* .The *no-inspection* profile does not perform SSL inspection, so it should not be selected with other UTM profiles.
5. 
                                                    Configure the other settings as needed.
6. 
                                                    Click *OK* .

#### Configuring a URL filter in the CLI

###### To create a URL filter for Facebook:

```
config webfilter urlfilter
    edit 1
        set name "webfilter"
        config entries
            edit 1
                set url "*facebook.com"
                set type wildcard
                set action block
            next
        end
    next
end
```
                                            ###### To apply the URL filter to a web filter profile:

```
config webfilter profile
    edit "webfilter" 
        config web
            set urlfilter-table 1
        end
        config ftgd-wf
            ...
        end
    next
end
```
                                            ###### To apply the web filter profile to a firewall policy:

```
config firewall policy
    edit 1
        set name "WF"
        set srcintf "wan2"
        set dstintf "wan1"
        set srcaddr "all"
        set dstaddr "all"
        set action accept
        set schedule "always"
        set service "ALL"
        set utm-status enable
        set inspection-mode proxy
        set logtraffic all
        set webfilter-profile "webfilter" 
        set ssl-ssh-profile "certificate-inspection"
        set nat enable
    next
end
```
                                            ### Verifying the URL filter results

Verify the URL filter results by going to a blocked website. For example, when you go to the Facebook website, the replacement message appears:


###### To customize the URL web page blocked message:

1. 
                                                    Go to *System > Replacement Messages* .
2. 
                                                    In the *HTTP* section, select*URL Block Page* and click*Edit* .
3. 
                                                    Edit the HTML to customize the message. See Replacement messages for more information.

###### To check web filter logs in the GUI:

1. 
                                                    Go to *Log & Report > Security Events* .
2. 
                                                    Click the *Web Filter* card name.
3. 
                                                    If there are a lot of log entries, click *Add Filter* and select*Event Type > urlfilter* to display logs generated by the URL filter.

###### To check web filter logs in the CLI:

