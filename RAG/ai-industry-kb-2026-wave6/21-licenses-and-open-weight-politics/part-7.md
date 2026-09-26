---
id: ai-industry-kb-2026-wave6/21-licenses-and-open-weight-politics/part-7
title: "§21. Licenses and Open-Weight Politics (part 7)"
domain: licenses-and-open-weight-politics
role: deep-dive
task: model-release
actors: ["Alibaba", "China", "DeepSeek", "EU", "Falcon", "MiniMax", "Moonshot", "Nvidia", "United States", "Xiaomi", "Z.ai", "vLLM"]
dates: ["2025-01-13", "2025-01-15", "2025-07", "2026-04-20", "2026-07-31", "2026-08-22", "2026-09-21"]
keywords: ["license", "licenses", "apache", "deepseek", "diffusion", "distribution", "foundry", "glm", "kimi", "nvidia", "open weights", "parameters"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [10244, 10259]
section: "§21. Licenses and Open-Weight Politics"
delta_of: ai-industry-kb-2026
sha256: 30b2da5ef009a16d54695f324f362c167c8e6b847d5f2ac73d807dfe7f8c3ebb
---

# §21. Licenses and Open-Weight Politics (part 7)

- The BIS legal article is "parameters", not "weights": ECCN 4E091 controls the parameters of AI models trained above 10^26 operations — the legal term covers full, pruned, and distilled parameter sets alike [SECONDARY]. Source: https://theaicounsel.net/wp-content/uploads/2025/06/01_25_bis.pdf
- BIS extended deemed-export rules so that sharing restricted model weights with foreign nationals (e.g. Chinese nationals) inside the United States is itself a controlled transfer — the control reaches domestic labs with foreign staff [SECONDARY]. Source: https://www.caixinglobal.com/2025-01-15/detailed-explanation-of-new-us-export-controls-on-ai-chips-overseas-ai-model-training-restricted-102279207.html
- July 2025 update: BIS added reporting requirements for exporters using the AIA license exception — the exception is not a paperwork-free path [SECONDARY]. Source: https://theaicounsel.net/wp-content/uploads/2025/06/01_25_bis.pdf
- The 4E091 rule was announced January 13, 2025 as part of the AI diffusion package; the press release pairs it with foundry due-diligence requirements [SECONDARY]. Source: https://admin.govexec.com/media/general/2025/1/ai_embargoed_press_release.pdf
- Qwen3.7MaxThinking is MIT-licensed per the KTransformers serving docs — the Qwen flagship reasoning model ships under MIT, not Apache [COMMUNITY]. Source: https://github.com/vLLM-X/ktransformers/blob/HEAD/doc/en/api/server.rst
- GLM-5.3-Flash is described as "the first 200B-class open-source reasoning model running natively on domestic chips" — MIT plus non-NVIDIA hardware is the news, not just the license [SECONDARY]. Source: https://pasqualepillitteri.it/en/news/13263/ox-alpha-revealed-zhipu-glm-5-3-flash-mit
- DeepSeek V4.1 Flash released July 31, 2026 ("0731"), MIT, 284B total/13B active, text — the versioned card to keep separate from generic V4 Flash [SECONDARY]. Source: https://kimi-k2.org/blog/56-kimi-k3-vs-deepseek-v4-flash-0731
- MiMo V2.6-Pro (2026-09-21/22): MIT, per Xiaomi's MIT-licensed MiMo line — shares V2.5-Pro's 1.02T/42B profile but is a distinct release [SECONDARY]. Source: wave6/03-model-weights-wave3.md (infoworld MIT line)
- Kimi K2.6 (2026-04-20) license not verified in this pass — do not assume it inherits K3's Modified MIT [UNVERIFIED].
- MiniMax M2.7, GLM-4.7, and Falcon-H1 2026 releases: licenses not pulled in this pass — marked UNVERIFIED rather than defaulted [UNVERIFIED].
- The H3 README's license line is exactly "Licensed under the MiniMax H3 Community License Agreement" with a link to the agreement — first-party confirmation of the license name [VENDOR]. Source: https://huggingface.co/OpenVDN/vdn-minimax-h3/blob/main/README.md
- The autodirector license-notes file transcribes the H3 exclusion as "the United States, the European Union, the United Kingdom, and South Korea" and the $20M revenue threshold — matching the other two transcriptions [SECONDARY]. Source: https://github.com/makhmudovmurod/autodirector/blob/HEAD/Docs/license-notes.md
- MiniMax Music 3's analysis notes the model was "released under an open weights license" with the territorial change as the story — the analysis treats the no-exclusion as a deliberate legal correction [SECONDARY]. Source: https://digitalmatters.me/artificial-intelligence-ai/minimax-music-3/
- The MiniMax X post (2026-08-22) confirming H3's open weights is cited by the hardware guide as the release announcement — social posts are now primary license-distribution channels [SECONDARY]. Source: https://runaihome.com/blog/minimax-h3-open-weights-local-ai-hardware-guide-2026/
- NVIDIA's nemotron-speech-streaming-en-0.6b carries the NOTICE-file requirement — even small NVIDIA models inherit the Open Model License formalities [COMMUNITY]. Source: https://github.com/mikeschirtzinger/silent-notetaker/blob/HEAD/docs/research/model-licenses.md

