---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100413634-fen1022j-02-resources-en-us-topic-000000130617247-204fed34-4
title: "hedex-api-pages-edoc1100413634-fen1022j-02-resources-en-us-topic-000000130617247-204fed34"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100413634-fen1022j-02-resources-en-us-topic-000000130617247-204fed34.md
source_anchor: ""
source_lines: [165, 193]
sha256: fe1977047cbb81c2de4f5aa809991e9aa014eb5becc2af83d7dbe62344ea0eb0
---

# hedex-api-pages-edoc1100413634-fen1022j-02-resources-en-us-topic-000000130617247-204fed34

  - Run the interface atm interface-number command to enter the view of the VDSL interface working in ATM mode.
  - Run the shutdown command to deactivate the VDSL interface.
- Run the vdsl band v43 { off | on } command to disable or enable the v43 carrier band on the VDSL interface.By default, the v43 carrier band is enabled on a VDSL interface.
- Set uplink parameters for the VDSL interface.
  - Run the adsl standard { adsl2 [ annexa | annexm | annexb | annexj | annexl ] | adsl2+ [ annexa | annexm | annexb | annexj ] | auto | gdmt [ annexa | annexb ] | t1413 } command to configure a transmission standard for the VDSL interface in ATM mode.By default, the transmission standard for a VDSL interface in ATM mode is auto.
  - Run the adsl bitswap { off | on } command to enable or disable bit exchange on the VDSL interface.By default, bit exchange is enabled on a VDSL interface.
  - Run the adsl sra { off | on } command to enable or disable seamless rate adaptation on the VDSL interface.By default, seamless rate adaptation is disabled on a VDSL interface.
  - Run the adsl trellis { off | on } command to enable or disable trellis coding on the VDSL interface.By default, trellis coding is enabled on a VDSL interface.
- Run the undo shutdown command to activate the VDSL interface.
To access the Internet through VDSL (in PTM mode):
- Log in to the AR using STelnet or through the console port.
- Configure the VDSL interface to work in PTM mode.
  - Run the system-view command to enter the system view.
  - Run the set workmode slot slot-id vdsl ptm command to configure the VDSL interface to work in PTM mode.By default, a VDSL interface works in PTM mode.
- Deactivate the VDSL interface.
  - Run the interface ethernet interface-number command to enter the view of the VDSL interface working in PTM mode.
  - Run the shutdown command to deactivate the VDSL interface.
- Set uplink parameters for the VDSL interface.
  - Run the vdsl standard vdsl2 { annexa | annexb } command to configure a transmission standard for the VDSL interface in PTM mode.By default, a VDSL interface in PTM mode automatically uses the transmission standard of the peer interface.
  - Run the adsl bitswap { off | on } command to enable or disable bit exchange on the VDSL interface.By default, bit exchange is enabled on a VDSL interface.
  - Run the adsl sra { off | on } command to enable or disable seamless rate adaptation on the VDSL interface.By default, seamless rate adaptation is disabled on a VDSL interface.
  - Run the adsl trellis { off | on } command to enable or disable trellis coding on the VDSL interface.By default, trellis coding is enabled on a VDSL interface.
- Run the undo shutdown command to activate the VDSL interface.
Registering the AR with iMaster NCE-Campus
- Run the system-view command to enter the system view.
- Run the agile controller host host port port command to configure the IP address/URL and port number information for the AR to register with iMaster NCE-Campus. The value of host in this command can be an IP address or a URL.
- Run the agile controller bootstrap host host port port verifytype [ esn | code [ verifycode verifycode ] ] command to configure the IP address/domain name and port number of a bootstrap server and specify the voucher verification mode.
- Run the quit command to exit the system view.
- Run the save command to save the configuration.
