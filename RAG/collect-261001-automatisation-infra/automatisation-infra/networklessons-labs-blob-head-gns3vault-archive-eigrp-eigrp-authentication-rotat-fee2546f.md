---
id: collect-261001-automatisation-infra/automatisation-infra/networklessons-labs-blob-head-gns3vault-archive-eigrp-eigrp-authentication-rotat-fee2546f
title: "networklessons-labs-blob-head-gns3vault-archive-eigrp-eigrp-authentication-rotat-fee2546f"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: ["2020-02", "2021-02"]
keywords: ["parameters"]
source: docs/RAG/collect-261001-automatisation-infra/networklessons-labs-blob-head-gns3vault-archive-eigrp-eigrp-authentication-rotat-fee2546f.md
source_anchor: ""
source_lines: [1, 11]
sha256: 0d811771369394617f17f86074c5d34d051d6c836457caa7922f9835a7cdbfe7
---

# networklessons-labs-blob-head-gns3vault-archive-eigrp-eigrp-authentication-rotat-fee2546f

As the senior security officer you decide all routing protocols should be configured as secure as possible. The company you work for has a single vendor policy and since you only have Cisco equipment you are running EIGRP (Enhanced Interior Gateway Routing Protocol). EIGRP has more advanced features for authentication since it uses a key-chain. The key-chain supports rotating keys which makes it more secure than having a single static key. Before implementing this for your whole organization you decide to test your enhanced security in a lab environment.
- All IP addresses have been preconfigured for you.
- EIGRP has been preconfigured for you (AS12).
- Enable EIGRP authentication between router Jack and Johnson. Use the following parameters:
  - Key-chain should be called: GNS3VAULT
  - Key1: password VAULT
  - Key2: password SAFE
- Key1 should be sent until 9:00AM on the 2nd of February 2020 and should be accepted 15 minutes past this time.
- Key2 should be valid from 8:50AM on the 2nd of February 2020 and should be valid till the 1st of February 2021.
- Make sure routing adjacencies do not drop when the keys are switched.
- c3640-jk9s-mz.124-16.bin
