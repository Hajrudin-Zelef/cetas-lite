---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-80394-unifi6-lr-access-points-go-into-scan-loop-d612f0e2
title: "questions-80394-unifi6-lr-access-points-go-into-scan-loop-d612f0e2"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-80394-unifi6-lr-access-points-go-into-scan-loop-d612f0e2.md
source_anchor: ""
source_lines: [1, 27]
sha256: 15ce0372e83724b89bbc1a7fd7a384ceaf772d8d587fac28ff6673cac211e163
---

# questions-80394-unifi6-lr-access-points-go-into-scan-loop-d612f0e2

I have a Ubiquiti setup with four Unifi 6LR access points. I manage to adopt them in the controller just fine and I do briefly see the wifi come up, but then it immediately drops. SSH'ing into the access point and running logread, I see:
Sat Oct 15 00:03:47 2022 user.err : uplink-monitor[7903]: uplink-monitor.get_link_info(): uplink addr is null
Sat Oct 15 00:03:48 2022 daemon.notice wpa_supplicant[9283]: Scan requested (ret=0) - scan timeout 60 seconds
Sat Oct 15 00:03:48 2022 kern.warn kernel: [  461.049043] Scan Request for Apcli i/f proceed for scanning
Sat Oct 15 00:03:48 2022 kern.warn kernel: [  461.049147]  StaSiteSurvey go
Sat Oct 15 00:03:48 2022 kern.warn kernel: [  461.049195] SYNC - sync_fsm_scan_req_action:[1214] LAST_CH: 0, BAND: 0
Sat Oct 15 00:03:48 2022 kern.warn kernel: [  461.065001] ExtEventBeaconLostHandler::FW EVENT (60:22:32:27:16:33), Reason 0x10
Sat Oct 15 00:03:48 2022 kern.warn kernel: [  461.065071]   AP Beacon OFF!!!
Sat Oct 15 00:03:51 2022 kern.warn kernel: [  464.192612] ExtEventBeaconLostHandler::FW EVENT (60:22:32:27:16:33), Reason 0x10
Sat Oct 15 00:03:51 2022 kern.warn kernel: [  464.192710]   AP Beacon OFF!!!
Sat Oct 15 00:03:52 2022 kern.warn kernel: [  464.871661] rai0: 80211> cfg80211_scan_done
Sat Oct 15 00:03:53 2022 daemon.notice wpa_supplicant[9283]: Scan requested (ret=0) - scan timeout 60 seconds
Sat Oct 15 00:03:53 2022 kern.warn kernel: [  465.873345] Scan Request for Apcli i/f proceed for scanning
Sat Oct 15 00:03:53 2022 kern.warn kernel: [  465.873450]  StaSiteSurvey go
Sat Oct 15 00:03:53 2022 kern.warn kernel: [  465.874041] SYNC - sync_fsm_scan_req_action:[1214] LAST_CH: 0, BAND: 0
Sat Oct 15 00:03:53 2022 kern.warn kernel: [  465.874412] ExtEventBeaconLostHandler::FW EVENT (60:22:32:27:16:33), Reason 0x10
Sat Oct 15 00:03:53 2022 kern.warn kernel: [  465.874450]   AP Beacon OFF!!!
Sat Oct 15 00:03:56 2022 kern.warn kernel: [  469.192616] ExtEventBeaconLostHandler::FW EVENT (60:22:32:27:16:33), Reason 0x10
Sat Oct 15 00:03:56 2022 kern.warn kernel: [  469.192712]   AP Beacon OFF!!!
Sat Oct 15 00:03:57 2022 kern.warn kernel: [  469.695774] rai0: 80211> cfg80211_scan_done
Sat Oct 15 00:03:58 2022 daemon.notice wpa_supplicant[9283]: Scan requested (ret=0) - scan timeout 60 seconds
Sat Oct 15 00:03:58 2022 kern.warn kernel: [  470.697519] Scan Request for Apcli i/f proceed for scanning
Sat Oct 15 00:03:58 2022 kern.warn kernel: [  470.697621]  StaSiteSurvey go
Sat Oct 15 00:03:58 2022 kern.warn kernel: [  470.697670] SYNC - sync_fsm_scan_req_action:[1214] LAST_CH: 0, BAND: 0
Sat Oct 15 00:03:58 2022 kern.warn kernel: [  470.713438] ExtEventBeaconLostHandler::FW EVENT (60:22:32:27:16:33), Reason 0x10
Sat Oct 15 00:03:58 2022 kern.warn kernel: [  470.713508]   AP Beacon OFF!!!
so it appears the access point is going into some sort of scan loop for other wifi networks. However, I'm not asking it to do that? What's going on here?
