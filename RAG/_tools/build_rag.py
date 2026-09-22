#!/usr/bin/env python3
"""Build the RAG corpus for a briefing-ia document.

Partitions the source Markdown into contiguous line ranges cut on heading
boundaries, writes one file per chunk with a YAML metadata header, and emits
the INDEX/manifest. The script is deterministic and idempotent; it refuses to
finish if the partition is not an exact cover of the source (no gap, no overlap).

Source: docs/RAG/briefing-ia-2026-en.md
Output: RAG/briefing-ia-2026/
"""

from __future__ import annotations

import hashlib
import json
import re
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]
SOURCE = ROOT / "docs" / "RAG" / "briefing-ia-2026-en.md"
RAG_ROOT = ROOT / "RAG"
CORPUS_SLUG = "briefing-ia-2026"
CORPUS_TITLE = "AI News 2026 — Reference Dossier"
OUT = RAG_ROOT / CORPUS_SLUG

# (folder, slug, first_heading_line, title)
CHUNKS = [
    # 00-front-matter
    ("00-front-matter", "00-title", 1, "Reference Dossier — AI News 2026"),
    ("00-front-matter", "01-table-of-contents", 7, "Detailed table of contents (source)"),
    ("00-front-matter", "02-alphabetical-index", 285, "Alphabetical index (source)"),
    ("00-front-matter", "03-thematic-keywords", 381, "Thematic keywords (source)"),
    # 01-timeline-2026
    ("01-timeline-2026", "00-introduction-methodology-guide", 404, "Introduction, methodology and user guide"),
    ("01-timeline-2026", "01-january", 565, "January 2026: the month of latency"),
    ("01-timeline-2026", "02-february", 667, "February 2026: the capital deluge"),
    ("01-timeline-2026", "03-march", 831, "March 2026: the Californian regulatory lever"),
    ("01-timeline-2026", "04-april", 937, "April 2026: Anthropic's return and the pact of the century"),
    ("01-timeline-2026", "05-may", 1070, "May 2026: the model war and easy money"),
    ("01-timeline-2026", "06-june", 1206, "June 2026: record IPO, critical models and sovereignty"),
    ("01-timeline-2026", "07-july", 1417, "July 2026: multimodality, containment and the compute war"),
    ("01-timeline-2026", "08-august", 1658, "August 2026: consolidation, acquisitions and scientific agents"),
    ("01-timeline-2026", "09-september-part1", 1864, "September 2026 (part 1): models, science, acquisitions, advisory"),
    ("01-timeline-2026", "10-september-part2", 2053, "September 2026 (part 2): regulation, diplomacy, consumer agents, markets"),
    # 02-openai
    ("02-openai", "01-gpt56-stratification", 2189, "GPT-5.6: the Sol / Terra / Luna stratification"),
    ("02-openai", "02-june27-preview", 2334, "The June 27 preview: trusted partners, Trump decree and speculation"),
    ("02-openai", "03-containment-account", 2412, "The July 9–13 containment incident: account and technical mechanics"),
    ("02-openai", "04-containment-analysis", 2506, "The July 9–13 containment incident: chronology, impact and lessons"),
    ("02-openai", "05-astra-launch", 2671, "Astra: from the August slowdown to the September 3 launch"),
    ("02-openai", "06-astra-for-law", 2806, "Astra for Law (September 17): a configuration, not a model"),
    ("02-openai", "07-gpt-live", 2874, "GPT-Live (July 8): full-duplex voice"),
    ("02-openai", "08-stripe-openrouter", 2954, "Stripe acquires OpenRouter: payment rails of agentic AI"),
    ("02-openai", "09-ipo-filing", 3054, "The confidential IPO filing of June 8"),
    ("02-openai", "10-summary-threads", 3160, "OpenAI summary: master chronology, security and product threads"),
    ("02-openai", "11-summary-questions-method", 3255, "OpenAI summary: ecosystem thread, open questions, method note"),
    # 03-anthropic
    ("03-anthropic", "01-mythos-launch", 3364, "Anthropic: Mythos-class launch, safeguards, Glasswing, export control"),
    ("03-anthropic", "02-shutdown-restoration", 3552, "Anthropic: total shutdown, July 1 restoration, saga analysis"),
    ("03-anthropic", "03-september-iterations", 3681, "Anthropic: Fable 5.1 and Mythos 5.1 (September 1)"),
    ("03-anthropic", "04-model-lineup", 3753, "Anthropic 2026 model lineup: Opus, Sonnet, Fable, comparison"),
    ("03-anthropic", "05-funding", 3923, "Anthropic funding: Series G, Series H, run-rate"),
    ("03-anthropic", "06-compute-deals", 4021, "Anthropic compute deals: AWS, AMD, architecture comparison"),
    ("03-anthropic", "07-ipo-settlement", 4144, "Anthropic: IPO filings reported and the $1.5 billion settlement"),
    ("03-anthropic", "08-research-science", 4213, "Anthropic research: R&D share, protein design, Fermat, Zhipu accusation"),
    ("03-anthropic", "09-reference-tables", 4372, "Anthropic reference: master chronology and 2026 lineup table"),
    ("03-anthropic", "10-compute-lexicon", 4486, "Anthropic: seven gigawatts and the compliance lexicon"),
    ("03-anthropic", "11-outlook-synthesis-method", 4578, "Anthropic: 2027 outlook, synthesis, method note"),
    # 04-google-meta
    ("04-google-meta", "01-flash-cadence", 4725, "Google: the Gemini Flash release-train cadence"),
    ("04-google-meta", "02-gemini-live-voice", 4850, "Gemini 3.8 Live and Live Extended Thinking: voice as a native channel"),
    ("04-google-meta", "03-gemini-billion-users", 4957, "Gemini: one billion monthly users, distribution as a weapon"),
    ("04-google-meta", "04-irregular-incident", 5056, "The Irregular incident: a cybersecurity exercise out of the sandbox"),
    ("04-google-meta", "05-gemini-4-ghost", 5152, "Gemini 4: the ghost model, teased but not released"),
    ("04-google-meta", "06-muse-spark-pivot", 5223, "The Muse Spark pivot: Meta cuts into the LLaMA lineage"),
    ("04-google-meta", "07-meta-muse", 5298, "Meta Muse: Meta's personal agent"),
    ("04-google-meta", "08-muse-portfolio", 5418, "Muse Image, Video, Code and Glimmer: the creative portfolio"),
    ("04-google-meta", "09-muse-spark-13", 5520, "Muse Spark 1.3: agentic efficiency as a selling point"),
    ("04-google-meta", "10-aira3-kaggle", 5597, "AIRA₃ on Kaggle: the gold medal demonstration"),
    ("04-google-meta", "11-meta-infrastructure-summary", 5674, "Meta infrastructure: Graviton5 and the Vera Rubin pact, plus summary"),
    # 05-xai-microsoft
    ("05-xai-microsoft", "01-grok-4x-lineage", 5790, "xAI: the Grok 4.5→4.8 lineage"),
    ("05-xai-microsoft", "02-grok-copilot-precursors", 5936, "Grok in Copilot and the Microsoft-ecosystem precursors"),
    ("05-xai-microsoft", "03-xai-compute-finance", 6004, "xAI: Colossus v2, war chest and the SpaceX–xAI merger"),
    ("05-xai-microsoft", "04-build-mai-agents", 6086, "Microsoft Build 2026: the MAI family, Scout, Solara, Copilot Cowork"),
    ("05-xai-microsoft", "05-code-of-conduct", 6324, "Microsoft: the Code of Conduct for Humanist AI and the Suleyman–Amodei clash"),
    ("05-xai-microsoft", "06-master-timeline", 6429, "xAI and Microsoft: master timeline and conclusion"),
    ("05-xai-microsoft", "07-cross-analysis", 6492, "xAI/Microsoft cross-analysis: pricing, developer axis, Frontier, MAI doctrine"),
    ("05-xai-microsoft", "08-portraits-blindspots", 6619, "Portraits (Suleyman, Musk) and acknowledged blind spots"),
    ("05-xai-microsoft", "09-reference-tables", 6681, "Reference: glossary, Grok and MAI comparison tables"),
    ("05-xai-microsoft", "10-month-narratives", 6754, "Month narratives (June, September) and the economics of models"),
    ("05-xai-microsoft", "11-precursor-numbers-foundations", 6837, "Precursors, productivity battleground, quantitative memo, Scout foundations"),
    ("05-xai-microsoft", "12-synthesis-bets", 6934, "Synthesis: 2026 on one page, the rivalry, the five bets"),
    ("05-xai-microsoft", "13-reading-epilogue-colophon", 7016, "Reading note, epilogue and section colophon"),
    # 06-labs-china-europe
    ("06-labs-china-europe", "01-deepseek", 7049, "DeepSeek: V4-Flash, V4-Pro then V4.1 Flash"),
    ("06-labs-china-europe", "02-alibaba-qwen", 7206, "Alibaba / Qwen: accelerated image generation and distillation"),
    ("06-labs-china-europe", "03-moonshot-kimi", 7281, "Moonshot / Kimi: K3, open weights and the Hong Kong road"),
    ("06-labs-china-europe", "04-zai-glm", 7383, "Z.ai / GLM: raises, distillation accusation, security move"),
    ("06-labs-china-europe", "05-stepfun", 7469, "StepFun: Step 5 Preview"),
    ("06-labs-china-europe", "06-mistral", 7531, "Mistral: record raise, Microsoft and the Vibe platform"),
    ("06-labs-china-europe", "07-cohere-aleph", 7650, "Cohere: Command A+ and the Aleph Alpha merger"),
    ("06-labs-china-europe", "08-sakana-apple", 7755, "Sakana Fugu and Apple iOS 27 / Siri AI"),
    ("06-labs-china-europe", "09-stability-inflection-ant", 7844, "Stability AI, Inflection AI and Ant / inclusionAI"),
    ("06-labs-china-europe", "10-cross-prices-sizes", 7948, "Cross-cutting readings: API prices and open-weight sizes"),
    ("06-labs-china-europe", "11-cross-distill-license-capital", 8011, "Cross-cutting readings: distillation, licenses, capital"),
    ("06-labs-china-europe", "12-cross-calendar-conclusion", 8092, "Cross-cutting readings: September calendar and section conclusion"),
    # 07-infrastructure-compute
    ("07-infrastructure-compute", "01-rubin-nvl72", 8201, "Rubin NVL72: the 2026 reference for dense AI compute"),
    ("07-infrastructure-compute", "02-meta-nvidia-ncp", 8340, "Meta, official NVIDIA NCP partner with Vera Rubin"),
    ("07-infrastructure-compute", "03-nvidia-huggingface", 8406, "Nvidia acquires Hugging Face"),
    ("07-infrastructure-compute", "04-malaysia-export", 8489, "1.3 million chips in Malaysia: the export-control dossier"),
    ("07-infrastructure-compute", "05-amd-helios", 8555, "AMD Advancing AI: Instinct MI400 and the Helios racks"),
    ("07-infrastructure-compute", "06-compute-deals", 8670, "The compute deals: Anthropic, OpenAI and the Helios ecosystem"),
    ("07-infrastructure-compute", "07-amd-trillion-day", 8776, "AMD crosses $1 trillion in market capitalisation"),
    ("07-infrastructure-compute", "08-intel-18a", 8844, "Intel: Panther Lake, Clearwater Forest and the 18A battle"),
    ("07-infrastructure-compute", "09-aws-agentcore", 8958, "AWS: AgentCore, WorkSpaces and agentic infrastructure"),
    ("07-infrastructure-compute", "10-bedrock-graviton5", 9067, "Bedrock and Graviton5: model catalog and in-house chip"),
    ("07-infrastructure-compute", "11-xai-colossus-v2", 9166, "xAI Colossus v2: 550,000 GPUs"),
    # 08-deals-funding
    ("08-deals-funding", "01-intro-spacex-xai", 9246, "The SpaceX–xAI mega-merger: $1,250 billion in stock"),
    ("08-deals-funding", "02-anthropic-ipo-season", 9322, "Anthropic hypergrowth and the 2026 IPO season"),
    ("08-deals-funding", "03-mistral-zai-raises", 9431, "Mistral and Z.ai: European record and Chinese fundraising"),
    ("08-deals-funding", "04-stability-stripe", 9520, "Stability AI raise and Stripe acquires OpenRouter"),
    ("08-deals-funding", "05-nvidia-huggingface", 9609, "Nvidia acquires Hugging Face: the promise of openness"),
    ("08-deals-funding", "06-meta-scale-manus", 9662, "Meta–Scale AI and Beijing's block of the Manus acquisition"),
    ("08-deals-funding", "07-schwarz-cohere-aleph", 9747, "Schwarz–Cohere and the Cohere + Aleph Alpha merger"),
    ("08-deals-funding", "08-compute-deals-summary", 9843, "The structuring compute deals and the deals summary table"),
    # 09-ai-safety
    ("09-ai-safety", "01-incident-facts-technical", 9935, "Containment incident: facts and technical mechanics"),
    ("09-ai-safety", "02-detection-firstorder", 10022, "Detection, disclosure, and why it is first-order"),
    ("09-ai-safety", "03-astra-critical-effects", 10070, "Direct effects and GPT-6 Astra, first 'Critical' model"),
    ("09-ai-safety", "04-advisory-labs-channels", 10122, "Distillation advisory, the six labs, channels and legal scope"),
    ("09-ai-safety", "05-legal-california", 10219, "California: kill-switch study and regulatory architecture"),
    ("09-ai-safety", "06-gemini-anthropic-frameworks", 10264, "Gemini incident, Anthropic frameworks and coordination antitrust"),
    ("09-ai-safety", "07-coordination-comparison-evalsafety", 10381, "Incident comparison and evaluation safety"),
    ("09-ai-safety", "08-benchmarks-distillation-apis", 10437, "Benchmarks, distillation doctrine and API defence"),
    ("09-ai-safety", "09-killswitch-verifiers-governance", 10510, "Kill switch, certified verifiers and governance layers"),
    ("09-ai-safety", "10-timeline-victim-attribution", 10590, "Safety timeline, Hugging Face collateral and attribution"),
    ("09-ai-safety", "11-scenarios-glossary-blindspots", 10660, "Scenarios 2027, reasoned glossary and blind spots"),
    ("09-ai-safety", "12-polecon-openweights-human", 10744, "Political economy of trust, open weights and the human factor"),
    ("09-ai-safety", "13-synthesis", 10823, "AI safety synthesis: what 2026 changes"),
    # 10-regulation-geopolitics
    ("10-regulation-geopolitics", "01-eu-digital-omnibus", 10872, "EU: the Digital Omnibus, article by article"),
    ("10-regulation-geopolitics", "02-us-state-patchwork", 10996, "US: the state patchwork, law by law"),
    ("10-regulation-geopolitics", "03-federal-preemption", 11132, "Federal preemption: the blocked path"),
    ("10-regulation-geopolitics", "04-g20-carolina", 11244, "G20 and the 'Carolina Principles'"),
    ("10-regulation-geopolitics", "05-us-china-dialogue", 11358, "The US–China dialogue on AI safety"),
    ("10-regulation-geopolitics", "06-export-sovereignty", 11505, "Export controls and technological sovereignty"),
    ("10-regulation-geopolitics", "07-us-frontier-debate", 11642, "The US debate: moratorium, mandatory benchmark and legislative proposals"),
    # 11-consumer-agents-research
    ("11-consumer-agents-research", "01-personal-agents", 11780, "Personal agents: comparison table and agent-by-agent analysis"),
    ("11-consumer-agents-research", "02-agent-comparison", 11966, "Personal agents: positioning, business models, maturity"),
    ("11-consumer-agents-research", "03-voice-multimodal", 12045, "Voice and multimodal: the agentic interface race"),
    ("11-consumer-agents-research", "04-health-youth", 12136, "Health and youth: two sensitive verticals"),
    ("11-consumer-agents-research", "05-distribution", 12188, "Distribution: scale as a strategic weapon"),
    ("11-consumer-agents-research", "06-paper2agent", 12243, "Paper2Agent: from publication to working agent (Nature)"),
    ("11-consumer-agents-research", "07-scientisttwo", 12324, "ScientistTwo: the autonomous discovery cycle (arXiv)"),
    ("11-consumer-agents-research", "08-alma-memory", 12401, "ALMA: meta-learning of agent memories (arXiv)"),
    ("11-consumer-agents-research", "09-concept-steering-breakthroughs", 12464, "Concept steering and applied breakthroughs (proteins, mathematics)"),
    ("11-consumer-agents-research", "10-synthesis", 12569, "Consumer agents and research synthesis: the year the agent left the chat"),
    # 12-reference
    ("12-reference", "01-models-table", 12622, "Appendix 1: complete models table"),
    ("12-reference", "02-glossary", 12772, "Appendix 2: glossary of terms"),
    ("12-reference", "03-actors-index", 12951, "Appendix 3: index of actors"),
    ("12-reference", "04-method-cautions", 13047, "Appendix 4: methodological cautions"),
]

FOLDER_DOMAIN = {
    "00-front-matter": "front-matter",
    "01-timeline-2026": "timeline",
    "02-openai": "openai",
    "03-anthropic": "anthropic",
    "04-google-meta": "google-meta",
    "05-xai-microsoft": "xai-microsoft",
    "06-labs-china-europe": "labs-china-europe",
    "07-infrastructure-compute": "infrastructure-compute",
    "08-deals-funding": "deals-funding",
    "09-ai-safety": "ai-safety",
    "10-regulation-geopolitics": "regulation-geopolitics",
    "11-consumer-agents-research": "consumer-agents-research",
    "12-reference": "reference",
}

FOLDER_TASK = {
    "00-front-matter": "reference",
    "01-timeline-2026": "chronology",
    "02-openai": "actor-profile",
    "03-anthropic": "actor-profile",
    "04-google-meta": "actor-profile",
    "05-xai-microsoft": "actor-profile",
    "06-labs-china-europe": "actor-profile",
    "07-infrastructure-compute": "infrastructure",
    "08-deals-funding": "funding-deals",
    "09-ai-safety": "ai-safety",
    "10-regulation-geopolitics": "regulation",
    "11-consumer-agents-research": "product-research",
    "12-reference": "reference",
}

TASK_OVERRIDES = [
    ("containment", "safety-incident"),
    ("incident", "safety-incident"),
    ("advisory", "distillation"),
    ("distill", "distillation"),
    ("benchmarks", "evaluation"),
    ("research-science", "research"),
    ("research", "research"),
    ("paper2agent", "research"),
    ("scientisttwo", "research"),
    ("alma", "research"),
    ("steering", "research"),
    ("breakthroughs", "research"),
    ("funding", "funding-deals"),
    ("ipo", "funding-deals"),
    ("raises", "funding-deals"),
    ("merger", "funding-deals"),
    ("stripe", "funding-deals"),
    ("scale-manus", "funding-deals"),
    ("schwarz", "funding-deals"),
    ("deals", "funding-deals"),
    ("spacex-xai", "funding-deals"),
    ("trillion-day", "finance"),
    ("killswitch", "regulation"),
    ("legal", "regulation"),
    ("preemption", "regulation"),
    ("g20", "regulation"),
    ("eu-", "regulation"),
    ("us-", "regulation"),
    ("export", "regulation"),
    ("ruling", "regulation"),
    ("rubin", "infrastructure"),
    ("nvidia", "infrastructure"),
    ("amd", "infrastructure"),
    ("intel", "infrastructure"),
    ("aws", "infrastructure"),
    ("bedrock", "infrastructure"),
    ("colossus", "infrastructure"),
    ("malaysia", "infrastructure"),
    ("compute", "infrastructure"),
    ("voice", "product"),
    ("personal-agents", "product"),
    ("agent-comparison", "product"),
    ("distribution", "product"),
    ("health", "product"),
    ("model-lineup", "model-release"),
    ("gpt", "model-release"),
    ("grok", "model-release"),
    ("muse", "model-release"),
    ("gemini", "model-release"),
    ("mai", "model-release"),
    ("astra", "model-release"),
    ("deepseek", "model-release"),
    ("qwen", "model-release"),
    ("kimi", "model-release"),
    ("glm", "model-release"),
    ("mistral", "model-release"),
    ("cohere", "model-release"),
    ("aleph", "model-release"),
    ("sakana", "model-release"),
    ("stepfun", "model-release"),
    ("stability", "model-release"),
    ("inflection", "model-release"),
    ("synthesis", "analysis"),
    ("summary", "analysis"),
    ("outlook", "analysis"),
    ("cross-analysis", "analysis"),
    ("master-timeline", "analysis"),
    ("month-narratives", "analysis"),
    ("reference-tables", "reference"),
    ("models-table", "reference"),
    ("glossary", "reference"),
    ("actors-index", "reference"),
    ("method-cautions", "reference"),
]

CANONICAL = {
    "02-openai/03-containment-account": ["containment-incident"],
    "02-openai/04-containment-analysis": ["containment-incident"],
    "02-openai/01-gpt56-stratification": ["gpt56-ga"],
    "02-openai/05-astra-launch": ["gpt6-astra"],
    "02-openai/08-stripe-openrouter": ["stripe-openrouter"],
    "03-anthropic/01-mythos-launch": ["anthropic-mythos-saga"],
    "03-anthropic/02-shutdown-restoration": ["anthropic-mythos-saga"],
    "03-anthropic/05-funding": ["anthropic-funding"],
    "03-anthropic/06-compute-deals": ["anthropic-compute-deals"],
    "03-anthropic/08-research-science": ["anthropic-research"],
    "04-google-meta/04-irregular-incident": ["irregular-incident"],
    "04-google-meta/07-meta-muse": ["meta-muse"],
    "05-xai-microsoft/01-grok-4x-lineage": ["grok-lineup"],
    "05-xai-microsoft/03-xai-compute-finance": ["spacex-xai-merger"],
    "05-xai-microsoft/04-build-mai-agents": ["microsoft-mai"],
    "07-infrastructure-compute/03-nvidia-huggingface": ["nvidia-huggingface"],
    "07-infrastructure-compute/05-amd-helios": ["amd-mi400"],
    "08-deals-funding/04-stability-stripe": ["stripe-openrouter"],
    "08-deals-funding/06-meta-scale-manus": ["meta-scale-manus"],
    "08-deals-funding/07-schwarz-cohere-aleph": ["cohere-aleph"],
    "09-ai-safety/01-incident-facts-technical": ["containment-incident"],
    "09-ai-safety/03-astra-critical-effects": ["gpt6-astra"],
    "09-ai-safety/04-advisory-labs-channels": ["distillation-advisory"],
    "09-ai-safety/06-gemini-anthropic-frameworks": ["irregular-incident"],
}

KNOWN_ACTORS = {
    "OpenAI": [r"OpenAI", r"GPT-\d", r"Astra", r"ChatGPT", r"Altman"],
    "Anthropic": [r"Anthropic", r"Claude", r"Fable 5", r"Mythos", r"Opus \d", r"Sonnet \d", r"Amodei"],
    "Google": [r"Google", r"Gemini", r"DeepMind", r"Kratsios"],
    "Meta": [r"\bMeta\b", r"Muse", r"LLaMA", r"LeCun", r"Scale AI"],
    "Microsoft": [r"Microsoft", r"Copilot", r"\bMAI\b", r"Suleyman", r"Nadella", r"Scout", r"Solara"],
    "xAI": [r"\bxAI\b", r"Grok", r"Colossus", r"\bMusk\b"],
    "SpaceX": [r"SpaceX", r"Starlink"],
    "Nvidia": [r"Nvidia", r"NVIDIA", r"Rubin", r"Blackwell", r"NVLink", r"\bNCP\b"],
    "AMD": [r"\bAMD\b", r"Instinct MI", r"\bMI4\d0\b", r"Helios", r"\bROCm\b", r"EPYC"],
    "Intel": [r"\bIntel\b", r"Panther Lake", r"Clearwater Forest", r"18A"],
    "AWS": [r"\bAWS\b", r"Amazon", r"Bedrock", r"Graviton", r"Trainium", r"AgentCore"],
    "Apple": [r"\bApple\b", r"iOS", r"\bSiri\b", r"iPhone"],
    "DeepSeek": [r"DeepSeek"],
    "Alibaba": [r"Alibaba", r"Qwen"],
    "Moonshot": [r"Moonshot", r"Kimi"],
    "Z.ai": [r"Z\.ai", r"Zhipu", r"\bGLM\b"],
    "Mistral": [r"Mistral", r"Leanstral", r"Robostral", r"\bVibe\b", r"Koyeb"],
    "Cohere": [r"Cohere", r"Command A\+", r"Aleph Alpha", r"Schwarz", r"StackIT"],
    "Stability AI": [r"Stability AI", r"Stable Audio"],
    "Inflection AI": [r"Inflection", r"Pi Journeys"],
    "Sakana": [r"Sakana", r"Fugu"],
    "StepFun": [r"StepFun", r"Step 5"],
    "Ant": [r"inclusionAI", r"LLaDA"],
    "Huawei": [r"Huawei", r"Ascend"],
    "Qualcomm": [r"Qualcomm"],
    "Samsung": [r"Samsung"],
    "Stripe": [r"Stripe"],
    "OpenRouter": [r"OpenRouter"],
    "Hugging Face": [r"Hugging Face"],
    "JFrog": [r"JFrog", r"Artifactory"],
    "ExploitGym": [r"ExploitGym"],
    "Irregular": [r"\bIrregular\b"],
    "Glasswing": [r"Glasswing"],
    "CISA": [r"CISA", r"\bNSA\b", r"\bFBI\b"],
    "EU": [r"European Union", r"Digital Omnibus", r"\bEU\b"],
    "California": [r"California", r"Newsom", r"\bSB 813\b", r"\bAB 1405\b", r"N-5-26", r"N-9-26"],
    "United States": [r"Trump", r"\bUS\b", r"United States", r"Bessent", r"Sanders", r"Casar", r"\bDOJ\b", r"Congress", r"Senate"],
    "China": [r"Beijing", r"Chinese", r"\bChina\b", r"He Lifeng", r"\bXi\b"],
    "G20": [r"\bG20\b", r"Carolina Principles", r"Doral"],
    "Malaysia": [r"Malaysia"],
    "Kevin Buzzard": [r"Buzzard"],
    "James Zou": [r"James Zou", r"Paper2Agent"],
}

KNOWN_TERMS = [
    "containment", "sandbox escape", "sandbox", "zero-day", "benchmark", "benchmarks",
    "distillation", "export control", "export controls", "sovereignty", "kill switch",
    "open weights", "open-weight", "open source", "personal agent", "personal agents",
    "cloud agent", "full-duplex", "voice", "transcription", "multimodal", "reasoning",
    "pricing", "price war", "deflation", "run-rate", "valuation", "ipo", "merger",
    "acquisition", "funding", "series g", "series h", "forward-commitment",
    "compute", "datacenter", "data centre", "gpu", "gpus", "asic", "accelerator",
    "quantization", "inference", "training", "context window", "moE", "parameters",
    "safeguards", "preparedness framework", "rsp", "asl", "alignment", "model welfare",
    "regulation", "preemption", "moratorium", "advisory", "antitrust",
    "cybersecurity", "cyber", "advisory", "disclosure", "attribution", "liability",
    "memory", "research", "protein", "fermat", "lean", "formalization", "agent",
    "agents", "agentic", "mcp", "tool use", "distribution", "consumer",
    "nvidia", "amd", "intel", "aws", "bedrock", "rubin", "graviton", "helios",
    "chatgpt", "claude", "gemini", "grok", "muse", "copilot", "scout", "kimi",
    "deepseek", "qwen", "glm", "mistral", "cohere", "sakana", "fugu",
    "lawsuit", "settlement", "copyright", "license", "licenses", "watermarking",
    "incident", "containment", "jailbreak", "refusals", "exploit",
    "gpt-5.6", "gpt-6", "gpt-live", "astra", "sol", "terra", "luna",
    "fable 5", "mythos 5", "opus 4", "opus 5", "sonnet 5",
    "gemini 3.8", "gemini 4", "grok 4", "muse spark", "mai",
]


def norm_ws(s: str) -> str:
    return re.sub(r"\s+", " ", s).strip()


def compute_starts(rows):
    starts = []
    for folder, slug, heading, _title in rows:
        line = LINES[heading - 1]
        if heading > 1 and LINES[heading - 2].startswith("<a id="):
            starts.append(heading - 1)
        else:
            starts.append(heading)
    return starts


def anchor_of(start_line):
    raw = LINES[start_line - 1]
    m = re.match(r'<a id="([^"]+)">', raw.strip())
    return m.group(1) if m else ""


def first_heading(text):
    for line in text.splitlines():
        m = re.match(r"^#{1,4} +(.*)$", line)
        if m:
            return m.group(1).strip()
    return ""


MONTHS = {
    "january": "01", "february": "02", "march": "03", "april": "04", "may": "05",
    "june": "06", "july": "07", "august": "08", "september": "09", "october": "10",
    "november": "11", "december": "12",
}


def extract_dates(text):
    found = set()
    for m in re.finditer(r"\b([A-Z][a-z]+) (\d{1,2}), (\d{4})\b", text):
        mon = MONTHS.get(m.group(1).lower())
        if mon:
            found.add(f"{m.group(3)}-{mon}-{int(m.group(2)):02d}")
    for m in re.finditer(r"\b(\d{1,2})/(\d{1,2})/(\d{4})\b", text):
        found.add(f"{m.group(3)}-{int(m.group(2)):02d}-{int(m.group(1)):02d}")
    for m in re.finditer(r"\b(\d{4})-(\d{2})-(\d{2})\b", text):
        found.add(m.group(0))
    for m in re.finditer(r"\b([A-Z][a-z]+) (\d{4})\b", text):
        mon = MONTHS.get(m.group(1).lower())
        if mon:
            found.add(f"{m.group(2)}-{mon}")
    return sorted(found)


def extract_actors(text):
    out = []
    for label, pats in KNOWN_ACTORS.items():
        for p in pats:
            if re.search(p, text):
                out.append(label)
                break
    return sorted(out)


def extract_keywords(text, title):
    low = text.lower()
    kws = []
    for term in KNOWN_TERMS:
        pat = r"(?<![a-z0-9])" + re.escape(term) + r"(?![a-z0-9])"
        if re.search(pat, low):
            kws.append(term)
    base = norm_ws(title).lower()
    kws = sorted(set(kws), key=lambda t: (t not in base, t))
    return kws[:12]


def task_for(folder, slug):
    key = slug
    for sub, task in TASK_OVERRIDES:
        if sub in key:
            return task
    return FOLDER_TASK[folder]


def role_for(folder):
    if folder == "01-timeline-2026":
        return "timeline"
    if folder == "12-reference":
        return "appendix"
    if folder == "00-front-matter":
        return "reference"
    return "deep-dive"


def ystr(s: str) -> str:
    return json.dumps(s, ensure_ascii=False)


def ylist(items):
    return "[" + ", ".join(ystr(x) for x in items) + "]"


def main():
    global LINES
    if not SOURCE.exists():
        sys.exit(f"missing source: {SOURCE}")
    LINES = SOURCE.read_text(encoding="utf-8").splitlines(keepends=True)
    n = len(LINES)

    starts = compute_starts(CHUNKS)
    if starts[0] != 1:
        sys.exit("first chunk must start at line 1")
    for i in range(len(starts)):
        end = (starts[i + 1] - 1) if i + 1 < len(starts) else n
        if end < starts[i]:
            sys.exit(f"overlap/empty at chunk {CHUNKS[i]}")
        if i + 1 < len(starts) and starts[i + 1] != end + 1:
            sys.exit(f"gap at chunk {CHUNKS[i]}")

    # clean generated content (preserve hand-written README.md / NOTES.md)
    keep = {"README.md", "NOTES.md"}
    if OUT.exists():
        for p in sorted(OUT.rglob("*"), reverse=True):
            if p.is_file() and p.name not in keep:
                p.unlink()
            elif p.is_dir():
                p.rmdir()
    OUT.mkdir(parents=True, exist_ok=True)

    manifest = []
    anchor_index = {}
    canonical_index = {}
    folders = {}

    for i, (folder, slug, _heading, title) in enumerate(CHUNKS):
        start = starts[i]
        end = (starts[i + 1] - 1) if i + 1 < len(CHUNKS) else n
        body = "".join(LINES[start - 1:end])
        anchor = anchor_of(start)

        domain = FOLDER_DOMAIN[folder]
        role = role_for(folder)
        task = task_for(folder, slug)
        actors = extract_actors(body)
        dates = extract_dates(body)
        keywords = extract_keywords(body, title)
        canon = CANONICAL.get(f"{folder}/{slug}", [])
        cid = f"{CORPUS_SLUG}/{folder}/{slug}"
        digest = hashlib.sha256(body.encode("utf-8")).hexdigest()

        header = ["---"]
        header.append(f"id: {cid}")
        header.append(f"title: {ystr(title)}")
        header.append(f"domain: {domain}")
        header.append(f"role: {role}")
        header.append(f"task: {task}")
        header.append(f"actors: {ylist(actors)}")
        header.append(f"dates: {ylist(dates)}")
        header.append(f"keywords: {ylist(keywords)}")
        header.append(f"source: docs/RAG/{SOURCE.name}")
        header.append(f"source_anchor: {ystr('#' + anchor if anchor else '')}")
        header.append(f"source_lines: [{start}, {end}]")
        if canon:
            header.append(f"canonical_for: {ylist(canon)}")
        header.append(f"sha256: {digest}")
        header.append("---")

        first_content = body.lstrip()
        parts = ["\n".join(header), ""]
        if not first_content.startswith("# "):
            parts.append(f"# {title}")
            parts.append("")
        parts.append(body)
        content = "\n".join(parts)
        if not content.endswith("\n"):
            content += "\n"

        fdir = OUT / folder
        fdir.mkdir(parents=True, exist_ok=True)
        path = fdir / f"{slug}.md"
        path.write_text(content, encoding="utf-8")

        entry = {
            "id": cid,
            "path": f"{CORPUS_SLUG}/{folder}/{slug}.md",
            "title": title,
            "domain": domain,
            "role": role,
            "task": task,
            "actors": actors,
            "dates": dates,
            "keywords": keywords,
            "source": f"docs/RAG/{SOURCE.name}",
            "source_anchor": anchor,
            "source_lines": [start, end],
            "canonical_for": canon,
            "words": len(body.split()),
            "bytes": len(body.encode("utf-8")),
            "sha256": digest,
        }
        manifest.append(entry)
        for a in re.findall(r'<a id="([^"]+)">', body):
            anchor_index[a] = entry["path"]
        for c in canon:
            canonical_index.setdefault(c, []).append(entry["path"])
        folders.setdefault(folder, []).append(entry)

    # ---- verification: exact cover + reconstruction -------------------------
    recon = "".join("".join(LINES[s - 1:((starts[j + 1] - 1) if j + 1 < len(starts) else n)])
                    for j, s in enumerate(starts))
    source_text = "".join(LINES)
    assert recon == source_text, "reconstruction mismatch"
    assert len(manifest) == len(CHUNKS)

    anchors_in_source = re.findall(r'<a id="([^"]+)">', source_text)
    assert len(anchors_in_source) == len(anchor_index), (
        f"anchor count mismatch: source={len(anchors_in_source)} index={len(anchor_index)}")
    assert len(set(anchors_in_source)) == len(anchors_in_source), "duplicate anchors in source"

    # ---- manifest -----------------------------------------------------------
    corpus_manifest = {
        "corpus": CORPUS_SLUG,
        "title": CORPUS_TITLE,
        "source": f"docs/RAG/{SOURCE.name}",
        "source_lines": n,
        "source_sha256": hashlib.sha256(source_text.encode("utf-8")).hexdigest(),
        "chunk_count": len(manifest),
        "generated_by": "RAG/_tools/build_rag.py",
        "chunks": manifest,
    }
    (OUT / "manifest.json").write_text(
        json.dumps(corpus_manifest, ensure_ascii=False, indent=2) + "\n", encoding="utf-8")

    # ---- INDEX --------------------------------------------------------------
    def rel(path):
        return path.split(f"{CORPUS_SLUG}/", 1)[-1]

    total_words = sum(e["words"] for e in manifest)
    lines = []
    lines.append(f"# INDEX — {CORPUS_TITLE}")
    lines.append("")
    lines.append(f"Corpus `{CORPUS_SLUG}` · **{len(manifest)} fichiers** · "
                 f"{sum(e['source_lines'][1] - e['source_lines'][0] + 1 for e in manifest)} lignes source · "
                 f"~{total_words} mots · partition exacte de "
                 f"`docs/RAG/{SOURCE.name}`.")
    lines.append("")
    lines.append("## Mode d'emploi")
    lines.append("")
    lines.append("1. Filtrer dans `manifest.json` (ou les tableaux ci-dessous) sur "
                 "`domain`, `task`, `actors`, `dates` ou `keywords`.")
    lines.append("2. Ouvrir 1 à 3 fichiers ciblés ; chaque fichier est une unité "
                 "thématique auto-suffisante avec un en-tête YAML.")
    lines.append("3. Pour un événement répété dans plusieurs sections, préférer le "
                 "fichier marqué `canonical_for` (voir la table Événements canoniques).")
    lines.append("")
    lines.append("## Domaines (dossiers → fichiers)")
    lines.append("")
    for folder in folders:
        lines.append(f"### `{folder}/`")
        lines.append("")
        lines.append("| # | fichier | lignes source | rôle | tâche |")
        lines.append("|---|---|---|---|---|")
        for j, e in enumerate(folders[folder], 1):
            lines.append(f"| {j:02d} | [{e['title']}]({rel(e['path'])}) | "
                         f"{e['source_lines'][0]}–{e['source_lines'][1]} | {e['role']} | {e['task']} |")
        lines.append("")
    lines.append("## Par tâche")
    lines.append("")
    by_task = {}
    for e in manifest:
        by_task.setdefault(e["task"], []).append(e)
    for task in sorted(by_task):
        files = ", ".join(f"[{e['title']}]({rel(e['path'])})" for e in by_task[task])
        lines.append(f"- **{task}** — {files}")
    lines.append("")
    lines.append("## Par acteur")
    lines.append("")
    by_actor = {}
    for e in manifest:
        for a in e["actors"]:
            by_actor.setdefault(a, []).append(e["path"])
    for actor in sorted(by_actor):
        lines.append(f"- **{actor}** ({len(by_actor[actor])}) — " +
                     ", ".join(f"[{rel(p)}]({rel(p)})" for p in by_actor[actor]))
    lines.append("")
    lines.append("## Par date")
    lines.append("")
    by_date = {}
    for e in manifest:
        for d in e["dates"]:
            by_date.setdefault(d, []).append(rel(e["path"]))
    for d in sorted(by_date):
        lines.append(f"- **{d}** — " + ", ".join(f"[{p}]({p})" for p in by_date[d]))
    lines.append("")
    lines.append("## Événements canoniques (`canonical_for`)")
    lines.append("")
    lines.append("Un même événement est narré plusieurs fois (chronologie, profil, fond). "
                 "La version de fond est marquée canonique ici ; les autres occurrences "
                 "restent présentes et sont taguées `role: timeline`.")
    lines.append("")
    for event in sorted(canonical_index):
        lines.append(f"- **{event}** — " + ", ".join(
            f"[{p}]({p})" for p in canonical_index[event]))
    lines.append("")
    lines.append("## Ancres source → fichier")
    lines.append("")
    lines.append("`#sNN-M` (liens de la table des matières source) → fichier cible.")
    lines.append("")
    lines.append("| ancre | fichier |")
    lines.append("|---|---|")
    for a in sorted(anchor_index, key=lambda x: (len(x), x)):
        lines.append(f"| `#{a}` | [{anchor_index[a]}]({anchor_index[a]}) |")
    lines.append("")
    lines.append("## Carte de couverture (lignes source)")
    lines.append("")
    lines.append("| plage | fichier |")
    lines.append("|---|---|")
    for e in manifest:
        lines.append(f"| {e['source_lines'][0]}–{e['source_lines'][1]} | {e['path']} |")
    lines.append("")
    (OUT / "INDEX.md").write_text("\n".join(lines) + "\n", encoding="utf-8")

    # ---- global RAG index + manifest ---------------------------------------
    RAG_ROOT.mkdir(parents=True, exist_ok=True)
    global_lines = []
    global_lines.append("# INDEX — RAG")
    global_lines.append("")
    global_lines.append("Corpus RAG de référence pour Cetas. Chaque corpus est une "
                        "partition exacte de sa source, avec index et manifest.")
    global_lines.append("")
    global_lines.append("| corpus | titre | fichiers | source | index |")
    global_lines.append("|---|---|---|---|---|")
    global_lines.append(f"| `{CORPUS_SLUG}` | {CORPUS_TITLE} | {len(manifest)} | "
                        f"`docs/RAG/{SOURCE.name}` | [INDEX]({CORPUS_SLUG}/INDEX.md) |")
    global_lines.append("")
    global_lines.append("Voir aussi [README](README.md) et `manifest.json`.")
    global_lines.append("")
    (RAG_ROOT / "INDEX.md").write_text("\n".join(global_lines) + "\n", encoding="utf-8")
    global_manifest = {
        "corpora": [{
            "corpus": CORPUS_SLUG,
            "title": CORPUS_TITLE,
            "source": f"docs/RAG/{SOURCE.name}",
            "index": f"{CORPUS_SLUG}/INDEX.md",
            "manifest": f"{CORPUS_SLUG}/manifest.json",
            "chunk_count": len(manifest),
        }],
    }
    (RAG_ROOT / "manifest.json").write_text(
        json.dumps(global_manifest, ensure_ascii=False, indent=2) + "\n", encoding="utf-8")

    print(f"OK  {len(manifest)} chunks  |  source {n} lines  "
          f"|  {len(anchor_index)} anchors  |  {total_words} words")
    print(f"    source_sha256 = {corpus_manifest['source_sha256']}")


if __name__ == "__main__":
    LINES: list[str] = []
    main()
