#!/usr/bin/env python3
"""Build the RAG corpora.

Each corpus is a contiguous line-range partition of a source Markdown file,
cut on heading boundaries. One file per chunk with a YAML metadata header,
plus INDEX.md and manifest.json. Deterministic and idempotent; refuses to
finish unless the partition is an exact cover of the source.

Sources: docs/RAG/*.md
Output:  RAG/<corpus-slug>/
"""

from __future__ import annotations

import hashlib
import json
import re
import sys
import unicodedata
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]
RAG_ROOT = ROOT / "RAG"

# --------------------------------------------------------------------------
# corpus: briefing-ia-2026  (do not touch: committed, must stay byte-identical)
# --------------------------------------------------------------------------

IA_CHUNKS = [
    ("00-front-matter", "00-title", 1, "Reference Dossier — AI News 2026"),
    ("00-front-matter", "01-table-of-contents", 7, "Detailed table of contents (source)"),
    ("00-front-matter", "02-alphabetical-index", 285, "Alphabetical index (source)"),
    ("00-front-matter", "03-thematic-keywords", 381, "Thematic keywords (source)"),
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
    ("08-deals-funding", "01-intro-spacex-xai", 9246, "The SpaceX–xAI mega-merger: $1,250 billion in stock"),
    ("08-deals-funding", "02-anthropic-ipo-season", 9322, "Anthropic hypergrowth and the 2026 IPO season"),
    ("08-deals-funding", "03-mistral-zai-raises", 9431, "Mistral and Z.ai: European record and Chinese fundraising"),
    ("08-deals-funding", "04-stability-stripe", 9520, "Stability AI raise and Stripe acquires OpenRouter"),
    ("08-deals-funding", "05-nvidia-huggingface", 9609, "Nvidia acquires Hugging Face: the promise of openness"),
    ("08-deals-funding", "06-meta-scale-manus", 9662, "Meta–Scale AI and Beijing's block of the Manus acquisition"),
    ("08-deals-funding", "07-schwarz-cohere-aleph", 9747, "Schwarz–Cohere and the Cohere + Aleph Alpha merger"),
    ("08-deals-funding", "08-compute-deals-summary", 9843, "The structuring compute deals and the deals summary table"),
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
    ("10-regulation-geopolitics", "01-eu-digital-omnibus", 10872, "EU: the Digital Omnibus, article by article"),
    ("10-regulation-geopolitics", "02-us-state-patchwork", 10996, "US: the state patchwork, law by law"),
    ("10-regulation-geopolitics", "03-federal-preemption", 11132, "Federal preemption: the blocked path"),
    ("10-regulation-geopolitics", "04-g20-carolina", 11244, "G20 and the 'Carolina Principles'"),
    ("10-regulation-geopolitics", "05-us-china-dialogue", 11358, "The US–China dialogue on AI safety"),
    ("10-regulation-geopolitics", "06-export-sovereignty", 11505, "Export controls and technological sovereignty"),
    ("10-regulation-geopolitics", "07-us-frontier-debate", 11642, "The US debate: moratorium, mandatory benchmark and legislative proposals"),
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
    ("12-reference", "01-models-table", 12622, "Appendix 1: complete models table"),
    ("12-reference", "02-glossary", 12772, "Appendix 2: glossary of terms"),
    ("12-reference", "03-actors-index", 12951, "Appendix 3: index of actors"),
    ("12-reference", "04-method-cautions", 13047, "Appendix 4: methodological cautions"),
]

IA_FOLDER_DOMAIN = {
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

IA_FOLDER_TASK = {
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

IA_TASK_OVERRIDES = [
    ("containment", "safety-incident"), ("incident", "safety-incident"),
    ("advisory", "distillation"), ("distill", "distillation"),
    ("benchmarks", "evaluation"), ("research-science", "research"), ("research", "research"),
    ("paper2agent", "research"), ("scientisttwo", "research"), ("alma", "research"),
    ("steering", "research"), ("breakthroughs", "research"),
    ("funding", "funding-deals"), ("ipo", "funding-deals"), ("raises", "funding-deals"),
    ("merger", "funding-deals"), ("stripe", "funding-deals"), ("scale-manus", "funding-deals"),
    ("schwarz", "funding-deals"), ("deals", "funding-deals"), ("spacex-xai", "funding-deals"),
    ("trillion-day", "finance"), ("killswitch", "regulation"), ("legal", "regulation"),
    ("preemption", "regulation"), ("g20", "regulation"), ("eu-", "regulation"),
    ("us-", "regulation"), ("export", "regulation"), ("ruling", "regulation"),
    ("rubin", "infrastructure"), ("nvidia", "infrastructure"), ("amd", "infrastructure"),
    ("intel", "infrastructure"), ("aws", "infrastructure"), ("bedrock", "infrastructure"),
    ("colossus", "infrastructure"), ("malaysia", "infrastructure"), ("compute", "infrastructure"),
    ("voice", "product"), ("personal-agents", "product"), ("agent-comparison", "product"),
    ("distribution", "product"), ("health", "product"),
    ("model-lineup", "model-release"), ("gpt", "model-release"), ("grok", "model-release"),
    ("muse", "model-release"), ("gemini", "model-release"), ("mai", "model-release"),
    ("astra", "model-release"), ("deepseek", "model-release"), ("qwen", "model-release"),
    ("kimi", "model-release"), ("glm", "model-release"), ("mistral", "model-release"),
    ("cohere", "model-release"), ("aleph", "model-release"), ("sakana", "model-release"),
    ("stepfun", "model-release"), ("stability", "model-release"), ("inflection", "model-release"),
    ("synthesis", "analysis"), ("summary", "analysis"), ("outlook", "analysis"),
    ("cross-analysis", "analysis"), ("master-timeline", "analysis"), ("month-narratives", "analysis"),
    ("reference-tables", "reference"), ("models-table", "reference"), ("glossary", "reference"),
    ("actors-index", "reference"), ("method-cautions", "reference"),
]

IA_CANONICAL = {
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

IA_ACTORS = {
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

IA_TERMS = [
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

# --------------------------------------------------------------------------
# corpus: briefing-general-tech-2026
# --------------------------------------------------------------------------

GT_CHUNKS = [
    ("00-front-matter", "00-title", 1, "General Tech News 2026 — Hardware, Infrastructure & Consumer Tech"),
    ("00-front-matter", "01-table-of-contents", 11, "Front matter — table of contents"),
    ("00-front-matter", "02-alphabetical-index", 173, "Alphabetical index (source)"),
    ("00-front-matter", "03-thematic-keywords", 323, "Thematic keywords (source)"),
    ("00-front-matter", "04-dossier-front-matter", 384, "Reference Dossier — front matter and detailed table of contents"),
    ("01-chronology-2026", "00-introduction-methodology", 417, "Introduction, purpose and methodology"),
    ("01-chronology-2026", "01-february-2026", 609, "February 2026"),
    ("01-chronology-2026", "02-march-2026", 704, "March 2026"),
    ("01-chronology-2026", "03-april-2026", 775, "April 2026"),
    ("01-chronology-2026", "04-may-2026", 831, "May 2026"),
    ("01-chronology-2026", "05-june-2026", 874, "June 2026"),
    ("01-chronology-2026", "06-july-2026", 960, "July 2026"),
    ("01-chronology-2026", "07-august-2026", 1033, "August 2026"),
    ("01-chronology-2026", "08-september-2026", 1106, "September 2026"),
    ("02-gpus-accelerators", "00-chapter-introduction", 1333, "GPUs and AI accelerators — introduction"),
    ("02-gpus-accelerators", "01-nvidia-vera-rubin", 1443, "Nvidia: Vera Rubin in full production"),
    ("02-gpus-accelerators", "02-rubin-nvl72-specs", 1510, "Rubin NVL72: architecture and specifications"),
    ("02-gpus-accelerators", "03-mlperf-inference-v61", 1583, "MLPerf Inference v6.1: Rubin's first benchmark outing"),
    ("02-gpus-accelerators", "04-huang-agi-moment", 1669, "Jensen Huang and the 'AGI has arrived' moment"),
    ("02-gpus-accelerators", "05-amd-advancing-ai-helios", 1726, "AMD: Advancing AI and the Helios ramp"),
    ("02-gpus-accelerators", "06-mi455x-memory", 1798, "MI455X and the memory play"),
    ("02-gpus-accelerators", "07-amd-deals", 1871, "AMD's deals: Anthropic, OpenAI, Meta"),
    ("02-gpus-accelerators", "08-amd-trillion", 1944, "AMD crosses $1 trillion"),
    ("02-gpus-accelerators", "09-intel-panther-lake", 2006, "Intel: Core Ultra X9 and the Panther Lake generation"),
    ("02-gpus-accelerators", "10-intel-clearwater-forest", 2062, "Intel: Xeon 6+ Clearwater Forest and the datacenter fightback"),
    ("02-gpus-accelerators", "11-intel-crescent-island", 2128, "Intel Crescent Island: the LPDDR5X inference bet"),
    ("02-gpus-accelerators", "12-huawei-ascend-roadmap", 2204, "Huawei: Ascend roadmap at Huawei Connect 2026"),
    ("02-gpus-accelerators", "13-ascend-960dt-960pr", 2264, "Ascend 960DT vs 960PR: training and inference SKUs"),
    ("02-gpus-accelerators", "14-atlas-960-superpod", 2322, "Atlas 960 SuperPoD: scale, NPO optics and the 2.3x/2.5x claims"),
    ("02-gpus-accelerators", "15-per-chip-gap", 2396, "The per-chip gap: ~2 years behind Blackwell, ~10x under Rubin"),
    ("02-gpus-accelerators", "16-huawei-china-first-pytorch", 2459, "Huawei's China-first strategy and PyTorch support"),
    ("02-gpus-accelerators", "17-qualcomm-ai200", 2521, "Qualcomm AI200: rack-scale inference"),
    ("02-gpus-accelerators", "18-fujitsu-monaka", 2588, "Fujitsu MONAKA: sovereign inference from Japan"),
    ("02-gpus-accelerators", "19-competitive-landscape", 2644, "Competitive landscape: who stands where"),
    ("03-servers-datacenters", "00-chapter-introduction", 2883, "Servers and datacenters — introduction"),
    ("03-servers-datacenters", "01-idc-server-market", 2966, "The server market at +52%: IDC Q2 2026"),
    ("03-servers-datacenters", "02-gpu-accelerated-value", 3077, "GPU-accelerated systems dominate value"),
    ("03-servers-datacenters", "03-memory-crisis-ramageddon", 3192, "The memory crisis: 'RAMageddon'"),
    ("03-servers-datacenters", "04-dram-nand-prices", 3338, "DRAM and NAND: prices, quotes, horizons"),
    ("03-servers-datacenters", "05-hyperscaler-capex", 3500, "Hyperscaler capex: ~$725–730 billion in 2026"),
    ("03-servers-datacenters", "06-custom-asics", 3632, "Custom ASICs gain ground"),
    ("03-servers-datacenters", "07-apple-pcc-m5", 3764, "Apple: Private Cloud Compute moves to M5"),
    ("03-servers-datacenters", "08-m8-ultra-rumor", 3855, "The M8 Ultra enterprise server rumor"),
    ("03-servers-datacenters", "09-power-800v-dc", 3959, "Power delivery: the 800V DC transition"),
    ("03-servers-datacenters", "10-ualink", 4083, "UALink: the open coalition against NVLink"),
    ("04-training-inference-quantization", "00-chapter-introduction", 4230, "Training vs inference architectures and quantization — introduction"),
    ("04-training-inference-quantization", "01-specialization", 4303, "The great specialization: training vs inference"),
    ("04-training-inference-quantization", "02-mlperf-v61-illustration", 4411, "MLPerf v6.1 as an illustration of the shift"),
    ("04-training-inference-quantization", "03-nvidia-dynamo", 4503, "Nvidia Dynamo and disaggregated serving"),
    ("04-training-inference-quantization", "04-training-racks", 4592, "Training racks: dense, HBM, high power"),
    ("04-training-inference-quantization", "05-inference-racks", 4669, "Inference racks: disaggregated, LPDDR, edge"),
    ("04-training-inference-quantization", "06-cost-per-token", 4718, "Cost per token as the decisive metric"),
    ("04-training-inference-quantization", "07-quantization-economics", 4795, "Quantization: the economics (McKinsey, June 2026)"),
    ("04-training-inference-quantization", "08-fp8", 4875, "FP8: the safe production standard"),
    ("04-training-inference-quantization", "09-int4-nvfp4-mxfp4", 4943, "INT4, NVFP4, MXFP4: the aggressive standard"),
    ("04-training-inference-quantization", "10-awq-vs-gptq", 5034, "AWQ vs GPTQ, Marlin kernels"),
    ("04-training-inference-quantization", "11-bitnet-158", 5129, "BitNet and the 1.58-bit frontier"),
    ("04-training-inference-quantization", "12-prismml-bonsai", 5211, "PrismML Bonsai: 27B models in a few gigabytes"),
    ("04-training-inference-quantization", "13-turboquant", 5296, "TurboQuant: 3-bit KV cache"),
    ("04-training-inference-quantization", "14-runtimes", 5402, "Runtimes: vLLM, TensorRT-LLM, SGLang, llama.cpp"),
    ("04-training-inference-quantization", "15-what-quantization-changes", 5455, "What quantization changes for the industry"),
    ("05-networking-optics", "00-chapter-introduction", 5588, "Networking and optics — introduction"),
    ("05-networking-optics", "01-fscom-portfolio", 5606, "FS.com: the AI networking portfolio"),
    ("05-networking-optics", "02-16t-osfp-verified-nvidia", 5688, "1.6T OSFP and the 'verified on NVIDIA' claim"),
    ("05-networking-optics", "03-d7070-muxponder", 5774, "D7070 800G muxponder"),
    ("05-networking-optics", "04-wifi7-ampcon", 5838, "Wi-Fi 7 campus and AmpCon"),
    ("05-networking-optics", "05-fscom-financials", 5892, "FS.com H1 2026 financials"),
    ("05-networking-optics", "06-16t-ethernet-shipping", 5946, "1.6T Ethernet is shipping now"),
    ("05-networking-optics", "07-oif-1600zr", 6024, "OIF 1600ZR: coherent 1.6T over a single wavelength"),
    ("05-networking-optics", "08-ecoc-2026-malaga", 6091, "ECOC 2026 in Malaga"),
    ("05-networking-optics", "09-linear-pluggable-optics", 6168, "Linear Pluggable Optics (LPO)"),
    ("05-networking-optics", "10-macom-448g", 6251, "MACOM: 448G/lane toward 3.2T"),
    ("05-networking-optics", "11-credo-zeroflap", 6324, "Credo ZeroFlap 1.6T"),
    ("05-networking-optics", "12-telxius-nokia-800g", 6415, "Telxius deploys Nokia 800G coherent pluggables"),
    ("05-networking-optics", "13-chapter-synthesis", 6487, "Chapter synthesis: 2026, the year the fabric became the product"),
    ("06-consumer-devices", "00-chapter-front-matter", 6642, "Chapter 7 — Consumer Devices: front matter and table of contents"),
    ("06-consumer-devices", "01-iphone-duo", 6697, "iPhone Duo: Apple's first foldable"),
    ("06-consumer-devices", "02-surprise-and-shine", 6754, "The 'Surprise and shine' event"),
    ("06-consumer-devices", "03-duo-specs", 6819, "Duo specifications in detail"),
    ("06-consumer-devices", "04-nano-texture-crease", 6941, "Nano-texture and the crease question"),
    ("06-consumer-devices", "05-pricing-forecasts", 7058, "Pricing, availability and market forecasts"),
    ("06-consumer-devices", "06-samsung-switchers", 7137, "Samsung's offensive: the 'Switchers' campaign"),
    ("06-consumer-devices", "07-galaxy-z-fold-8", 7231, "Galaxy Z Fold 8: lighter, thinner, IP48"),
    ("06-consumer-devices", "08-foldable-market-shares", 7331, "Foldable market shares: projections"),
    ("06-consumer-devices", "09-googlebook", 7443, "Googlebook: Google returns to AI laptops"),
    ("06-consumer-devices", "10-aluminium-codename", 7526, "'Aluminium': codename, not a commercial name"),
    ("06-consumer-devices", "11-googlebook-lineup", 7587, "Googlebook lineup, pricing and availability"),
    ("06-consumer-devices", "12-ai-pcs", 7695, "AI PCs: 55% of shipments, but units fall"),
    ("06-consumer-devices", "13-consumer-price-ramageddon", 7779, "Consumer price pressure and 'RAMageddon'"),
    ("07-agi-governance", "00-chapter-introduction", 7889, "AGI debate and governance — introduction"),
    ("07-agi-governance", "01-huang-agi-arrived", 7955, "Huang: 'AGI has arrived'"),
    ("07-agi-governance", "02-brockman-agi-era", 8042, "Brockman: 'the AGI era'"),
    ("07-agi-governance", "03-deepmind-institute", 8133, "The DeepMind Institute"),
    ("07-agi-governance", "04-amodei-pace-frontier", 8239, "Amodei: 'We Must Pace the Frontier'"),
    ("07-agi-governance", "05-openai-standards", 8330, "OpenAI: 'Building standards for the next phase of AI'"),
    ("07-agi-governance", "06-recursive-self-improvement", 8446, "Recursive self-improvement and the Hugging Face precedent"),
    ("07-agi-governance", "07-22-country-declaration", 8548, "The 22-country declaration"),
    ("07-agi-governance", "08-us-china-bessent-he", 8663, "US-China: Bessent–He Lifeng and the incident line"),
    ("07-agi-governance", "09-altman-unsc", 8794, "Altman at the UN Security Council (planned 23/09)"),
    ("07-agi-governance", "10-trump-xi-summit", 8881, "Trump–Xi summit (planned 24/09)"),
    ("07-agi-governance", "11-un-scientific-panel", 8953, "The UN Independent International Scientific Panel on AI"),
    ("07-agi-governance", "12-definitions-caution-camp", 9049, "Definitions, timelines and the caution camp"),
    ("08-market-analysis", "00-chapter-introduction", 9136, "Market, deals and cross-cutting analysis — introduction"),
    ("08-market-analysis", "01-deals-recap", 9159, "Deals and funding recap table"),
    ("08-market-analysis", "02-memory-binding-constraint", 9243, "Memory as the binding constraint"),
    ("08-market-analysis", "03-us-china-stack", 9404, "US-China stack competition"),
    ("08-market-analysis", "04-agentic-infrastructure", 9505, "The agentic infrastructure shift"),
    ("08-market-analysis", "05-what-to-watch", 9601, "What to watch next"),
    ("09-appendices", "00-chapter-introduction", 9671, "Appendices — introduction"),
    ("09-appendices", "01-glossary", 9676, "Glossary"),
    ("09-appendices", "02-actors-index", 9821, "Actors index"),
    ("09-appendices", "03-methodological-notes", 9950, "Methodological notes and sensitive points"),
]

GT_FOLDER_DOMAIN = {
    "00-front-matter": "front-matter",
    "01-chronology-2026": "chronology",
    "02-gpus-accelerators": "gpus-accelerators",
    "03-servers-datacenters": "servers-datacenters",
    "04-training-inference-quantization": "training-inference-quantization",
    "05-networking-optics": "networking-optics",
    "06-consumer-devices": "consumer-devices",
    "07-agi-governance": "agi-governance",
    "08-market-analysis": "market-analysis",
    "09-appendices": "appendices",
}

GT_FOLDER_TASK = {
    "00-front-matter": "reference",
    "01-chronology-2026": "chronology",
    "02-gpus-accelerators": "hardware",
    "03-servers-datacenters": "infrastructure",
    "04-training-inference-quantization": "architecture",
    "05-networking-optics": "networking",
    "06-consumer-devices": "consumer",
    "07-agi-governance": "governance",
    "08-market-analysis": "analysis",
    "09-appendices": "reference",
}

GT_TASK_OVERRIDES = [
    ("ramageddon", "memory-crisis"), ("memory", "memory-crisis"), ("dram", "memory-crisis"),
    ("nand", "memory-crisis"), ("hbm", "memory-crisis"),
    ("capex", "funding-deals"), ("deals", "funding-deals"), ("trillion", "finance"),
    ("ualink", "interconnect"), ("nvlink", "interconnect"),
    ("quantization", "quantization"), ("fp8", "quantization"), ("int4", "quantization"),
    ("awq", "quantization"), ("bitnet", "quantization"), ("bonsai", "quantization"),
    ("turboquant", "quantization"), ("training-racks", "architecture"), ("inference-racks", "architecture"),
    ("dynamo", "architecture"), ("specialization", "architecture"),
    ("mlperf", "benchmark"), ("cost-per-token", "economics"),
    ("nvidia", "hardware"), ("rubin", "hardware"), ("amd", "hardware"), ("intel", "hardware"),
    ("huawei", "hardware"), ("ascend", "hardware"), ("qualcomm", "hardware"), ("fujitsu", "hardware"),
    ("gpus", "hardware"), ("gpu-accelerated", "hardware"),
    ("agi", "governance"), ("governance", "governance"), ("declaration", "governance"),
    ("amodei", "governance"), ("brockman", "governance"), ("huang", "governance"),
    ("deepmind", "governance"), ("openai-standards", "governance"), ("self-improvement", "governance"),
    ("unsc", "governance"), ("trump", "governance"), ("un-scientific", "governance"),
    ("definitions", "governance"), ("altman", "governance"),
    ("us-china", "geopolitics"), ("export", "geopolitics"),
    ("iphone", "consumer"), ("galaxy", "consumer"), ("foldable", "consumer"),
    ("googlebook", "consumer"), ("ai-pcs", "consumer"), ("pricing", "consumer"),
    ("aluminium", "consumer"), ("surprise", "consumer"),
    ("module", "infrastructure"), ("server", "infrastructure"), ("power-800v", "infrastructure"),
    ("custom-asics", "infrastructure"), ("apple-pcc", "infrastructure"), ("m8-ultra", "infrastructure"),
    ("idc", "market"), ("what-to-watch", "analysis"), ("us-china-stack", "geopolitics"),
    ("agentic-infrastructure", "architecture"),
    ("ethernet", "networking"), ("optics", "networking"), ("oif", "networking"),
    ("ecoc", "networking"), ("muxponder", "networking"), ("macom", "networking"),
    ("credo", "networking"), ("telxius", "networking"), ("fscom", "networking"),
    ("wifi7", "networking"), ("synthesis", "analysis"),
    ("glossary", "reference"), ("actors-index", "reference"), ("methodological", "reference"),
]

GT_CANONICAL = {
    "02-gpus-accelerators/01-nvidia-vera-rubin": ["nvidia-vera-rubin"],
    "02-gpus-accelerators/02-rubin-nvl72-specs": ["nvidia-vera-rubin"],
    "02-gpus-accelerators/03-mlperf-inference-v61": ["mlperf-v61"],
    "02-gpus-accelerators/04-huang-agi-moment": ["agi-arrival"],
    "02-gpus-accelerators/05-amd-advancing-ai-helios": ["amd-helios"],
    "02-gpus-accelerators/06-mi455x-memory": ["amd-mi455x"],
    "02-gpus-accelerators/07-amd-deals": ["amd-deals"],
    "02-gpus-accelerators/08-amd-trillion": ["amd-trillion"],
    "02-gpus-accelerators/09-intel-panther-lake": ["intel-panther-lake"],
    "02-gpus-accelerators/10-intel-clearwater-forest": ["intel-xeon6"],
    "02-gpus-accelerators/11-intel-crescent-island": ["intel-crescent-island"],
    "02-gpus-accelerators/12-huawei-ascend-roadmap": ["huawei-ascend"],
    "02-gpus-accelerators/13-ascend-960dt-960pr": ["huawei-ascend"],
    "02-gpus-accelerators/14-atlas-960-superpod": ["huawei-ascend"],
    "02-gpus-accelerators/15-per-chip-gap": ["huawei-ascend"],
    "02-gpus-accelerators/16-huawei-china-first-pytorch": ["huawei-ascend"],
    "02-gpus-accelerators/17-qualcomm-ai200": ["qualcomm-ai200"],
    "02-gpus-accelerators/18-fujitsu-monaka": ["fujitsu-monaka"],
    "02-gpus-accelerators/19-competitive-landscape": ["gpu-landscape"],
    "03-servers-datacenters/03-memory-crisis-ramageddon": ["memory-crisis"],
    "03-servers-datacenters/04-dram-nand-prices": ["memory-crisis"],
    "03-servers-datacenters/05-hyperscaler-capex": ["hyperscaler-capex"],
    "03-servers-datacenters/06-custom-asics": ["custom-asics"],
    "03-servers-datacenters/07-apple-pcc-m5": ["apple-pcc"],
    "03-servers-datacenters/09-power-800v-dc": ["800v-dc"],
    "03-servers-datacenters/10-ualink": ["ualink"],
    "04-training-inference-quantization/01-specialization": ["training-inference-split"],
    "04-training-inference-quantization/03-nvidia-dynamo": ["disaggregated-serving"],
    "04-training-inference-quantization/06-cost-per-token": ["cost-per-token"],
    "04-training-inference-quantization/07-quantization-economics": ["quantization"],
    "04-training-inference-quantization/13-turboquant": ["quantization"],
    "05-networking-optics/06-16t-ethernet-shipping": ["16t-ethernet"],
    "05-networking-optics/07-oif-1600zr": ["oif-1600zr"],
    "05-networking-optics/09-linear-pluggable-optics": ["lpo-cpo-npo"],
    "05-networking-optics/10-macom-448g": ["448g-lane"],
    "05-networking-optics/11-credo-zeroflap": ["credo-zeroflap"],
    "06-consumer-devices/01-iphone-duo": ["iphone-duo"],
    "06-consumer-devices/07-galaxy-z-fold-8": ["galaxy-z-fold-8"],
    "06-consumer-devices/09-googlebook": ["googlebook"],
    "06-consumer-devices/13-consumer-price-ramageddon": ["memory-crisis"],
    "07-agi-governance/01-huang-agi-arrived": ["agi-arrival"],
    "07-agi-governance/03-deepmind-institute": ["deepmind-institute"],
    "07-agi-governance/04-amodei-pace-frontier": ["amodei-pace-frontier"],
    "07-agi-governance/05-openai-standards": ["openai-standards"],
    "07-agi-governance/07-22-country-declaration": ["22-country-declaration"],
    "07-agi-governance/08-us-china-bessent-he": ["us-china-ai-dialogue"],
    "08-market-analysis/02-memory-binding-constraint": ["memory-crisis"],
    "08-market-analysis/03-us-china-stack": ["us-china-ai-dialogue"],
    "08-market-analysis/04-agentic-infrastructure": ["agentic-infrastructure"],
}

GT_ACTORS = {
    "Nvidia": [r"Nvidia", r"NVIDIA", r"Rubin", r"Blackwell", r"NVLink", r"GB200"],
    "AMD": [r"\bAMD\b", r"Instinct MI", r"\bMI\d{3}", r"Helios", r"\bROCm\b", r"EPYC", r"Versal"],
    "Intel": [r"\bIntel\b", r"Panther Lake", r"Clearwater Forest", r"Crescent Island", r"18A", r"Xeon", r"Darkmont", r"Xe3P"],
    "Huawei": [r"Huawei", r"Ascend", r"Atlas \d", r"SuperPoD", r"UnifiedBus", r"Lingqu", r"TorchNPU"],
    "Qualcomm": [r"Qualcomm", r"AI2\d0", r"HUMAIN", r"Cloud AI"],
    "Fujitsu": [r"Fujitsu", r"MONAKA"],
    "Apple": [r"\bApple\b", r"iPhone", r"\bMac\b", r"\bM5\b", r"\bM8\b", r"Private Cloud Compute", r"\bPCC\b"],
    "Samsung": [r"Samsung", r"Galaxy"],
    "Google": [r"Google", r"Gemini", r"\bTPU\b", r"DeepMind", r"Android", r"ChromeOS", r"Googlebook", r"Aluminium"],
    "Meta": [r"\bMeta\b", r"LLaMA", r"Muse", r"MTIA"],
    "Microsoft": [r"Microsoft", r"Maia", r"Azure", r"Copilot"],
    "Amazon": [r"Amazon", r"\bAWS\b", r"Trainium", r"Graviton", r"AgentCore", r"Annapurna"],
    "Broadcom": [r"Broadcom", r"Tomahawk", r"MTIA"],
    "Micron": [r"Micron"],
    "SK Hynix": [r"SK Hynix", r"SK Group", r"Chey"],
    "TSMC": [r"TSMC"],
    "OpenAI": [r"OpenAI", r"ChatGPT", r"Altman", r"Brockman"],
    "Anthropic": [r"Anthropic", r"Claude", r"Amodei"],
    "xAI": [r"\bxAI\b", r"Grok", r"\bMusk\b"],
    "FS.com": [r"FS\.com"],
    "Credo": [r"Credo", r"ZeroFlap"],
    "MACOM": [r"MACOM"],
    "Nokia": [r"Nokia"],
    "Telxius": [r"Telxius"],
    "UALink": [r"UALink"],
    "OIF": [r"\bOIF\b", r"1600ZR"],
    "OCP": [r"Open Compute", r"\bOCP\b"],
    "MLCommons": [r"MLCommons", r"MLPerf"],
    "IDC": [r"\bIDC\b"],
    "Gartner": [r"Gartner"],
    "TrendForce": [r"TrendForce"],
    "ASUS": [r"ASUS"],
    "Lenovo": [r"Lenovo"],
    "Acer": [r"Acer"],
    "Dell": [r"Dell"],
    "HP": [r"\bHP\b"],
    "CoreWeave": [r"CoreWeave"],
    "Crusoe": [r"Crusoe"],
    "Nebius": [r"Nebius"],
    "Positron": [r"Positron"],
    "Groq": [r"Groq"],
    "Cerebras": [r"Cerebras"],
    "Taalas": [r"Taalas"],
    "PrismML": [r"PrismML", r"Bonsai"],
    "BitNet": [r"BitNet"],
    "UN": [r"\bUN\b", r"United Nations", r"Guterres"],
    "China": [r"Beijing", r"Chinese", r"\bChina\b", r"Xi Jinping", r"He Lifeng"],
    "United States": [r"Trump", r"Bessent", r"United States"],
}

GT_TERMS = [
    "hbm", "hbm3", "hbm4", "dram", "nand", "ramageddon", "lpddr5x", "lpddr",
    "gpu", "gpus", "accelerator", "asic", "tpu", "trainium", "mtia", "maia",
    "nvlink", "ualink", "npo", "cpo", "lpo", "muxponder", "coherent optics", "optics",
    "ethernet", "serdes", "dsp", "wavelength", "dci",
    "kv cache", "prefill", "decode", "disaggregated", "quantization",
    "fp8", "fp4", "int4", "nvfp4", "mxfp4", "awq", "gptq", "bitnet", "turboquant",
    "tokens per dollar", "cost per token", "tokens-as-a-service", "capex", "hyperscaler",
    "custom silicon", "chiplet", "3nm", "2nm", "18a", "cowos", "packaging",
    "liquid cooling", "800v dc", "power delivery", "rack-scale", "superpod",
    "ascend", "vera rubin", "blackwell", "helios", "mi455x", "mi400", "panther lake",
    "clearwater forest", "crescent island", "monaka", "ai200",
    "foldable", "googlebook", "ai pc", "nano-texture", "crease",
    "inference", "training", "benchmark", "mlperf", "mlperf v6.1",
    "agi", "recursive self-improvement", "governance", "export controls", "sovereignty",
    "open weights", "distillation", "agentic", "disaggregated serving",
    "vllm", "tensorrt", "sglang", "llama.cpp", "pruning", "perplexity",
    "ipo", "merger", "acquisition", "funding", "valuation", "warrants", "run-rate",
    "gpt-6", "gemini", "grok", "claude", "muse", "llama", "qwen",
    "foundry", "wafer", "dram price", "memory shortage",
]

KB_EXTRA_ACTORS = {
    "Hugging Face": [r"Hugging ?Face"], "Groq": [r"Groq"], "Cerebras": [r"Cerebras"],
    "vLLM": [r"vLLM"], "SGLang": [r"SGLang"], "TensorRT-LLM": [r"TensorRT"],
    "Unsloth": [r"Unsloth"], "MiniMax": [r"MiniMax"], "LongCat": [r"LongCat"],
    "Meituan": [r"Meituan"], "Xiaomi": [r"Xiaomi"], "ByteDance": [r"ByteDance"],
    "Falcon": [r"Falcon"], "TII": [r"\bTII\b"], "Perplexity": [r"Perplexity"],
    "Poolside": [r"Poolside"], "CoreWeave": [r"CoreWeave"], "Crusoe": [r"Crusoe"],
    "Nebius": [r"Nebius"], "Lambda": [r"Lambda"], "Together AI": [r"Together AI"],
    "Fireworks AI": [r"Fireworks"], "Baseten": [r"Baseten"], "Nscale": [r"Nscale"],
    "IREN": [r"IREN"], "Oracle": [r"Oracle"], "Broadcom": [r"Broadcom"],
    "TSMC": [r"TSMC"], "Applied Digital": [r"Applied Digital"],
    "Fluidstack": [r"Fluidstack"], "StepFun": [r"StepFun"], "Baidu": [r"Baidu"],
}
KB_ACTORS = {**IA_ACTORS, **KB_EXTRA_ACTORS}

KB_TERMS = sorted(set(IA_TERMS + GT_TERMS + [
    "moe", "mixture of experts", "kv cache", "inference engine", "speculative decoding",
    "attention", "context window", "rotary", "gqa", "flash attention", "paged attention",
    "vllm", "sglang", "tensorrt-llm", "llama.cpp", "gguf", "awq", "gptq", "safetensors",
    "quantization", "fine-tuning", "lora", "qlora", "rlhf", "dpo", "pretraining",
    "embedding", "embeddings", "reranker", "multimodal", "omni", "video generation",
    "text-to-video", "text-to-image", "diffusion", "robotics", "humanoid",
    "neocloud", "funding round", "series a", "series b", "series c", "valuation",
    "revenue", "arr", "backlog", "market cap", "safety incident", "jailbreak",
    "data breach", "cyberattack", "guardrails", "regulation", "governance",
    "model context protocol", "mcp", "tool calling", "jailbreak", "alignment",
    "open weights", "license", "apache", "mit license", "benchmark", "leaderboard",
    "pricing", "cost", "latency", "throughput", "tokens per second", "energy",
]))


def slugify(s):
    s = re.sub(r"[§#]", " ", s)
    s = unicodedata.normalize("NFKD", s).encode("ascii", "ignore").decode("ascii")
    s = re.sub(r"[^A-Za-z0-9]+", "-", s)
    return s.strip("-").lower()[:60] or "part"


def _strip_num(s):
    return re.sub(r"^\s*(?:§\s*)?\d+[\.\)]?\s*", "", s).strip() or s.strip()


_TASK_MAP = [
    ("compute", "funding-deals"), ("capital", "funding-deals"), ("deal", "funding-deals"),
    ("funding", "funding-deals"), ("startup", "funding-deals"), ("ipo", "finance"),
    ("public market", "finance"), ("econom", "finance"), ("invest", "funding-deals"),
    ("safety incident", "ai-safety"), ("safety", "ai-safety"), ("breach", "ai-safety"),
    ("governance", "regulation"), ("regulat", "regulation"), ("policy", "regulation"),
    ("gpu", "hardware"), ("hardware", "hardware"), ("chip", "hardware"), ("robotic", "robotics"),
    ("inference engine", "architecture"), ("kv cache", "architecture"), ("moe", "architecture"),
    ("training", "training"), ("quantiz", "quantization"), ("format", "quantization"),
    ("quantization", "quantization"), ("multimodal", "multimodal"), ("video", "multimodal"),
    ("agent", "agents"), ("mcp", "agents"), ("hugging face", "platform"),
    ("frontier", "model-release"), ("open-weight", "model-release"), ("model inventory", "reference"),
    ("pricing", "pricing"), ("benchmark", "benchmark"), ("license", "licenses"),
    ("index", "reference"), ("keyword", "reference"), ("contradiction", "reference"),
    ("consolidation", "reference"), ("dedup", "reference"),
    ("deepseek", "actor-profile"), ("qwen", "actor-profile"), ("glm", "actor-profile"),
    ("kimi", "actor-profile"), ("minimax", "actor-profile"), ("xiaomi", "actor-profile"),
    ("openai", "actor-profile"), ("anthropic", "actor-profile"), ("google", "actor-profile"),
    ("meta", "actor-profile"), ("xai", "actor-profile"), ("mistral", "actor-profile"),
    ("nvidia", "actor-profile"), ("cohere", "actor-profile"), ("other chinese", "actor-profile"),
    ("coding", "model-release"), ("reasoning", "model-release"), ("small model", "model-release"),
]


def task_for_section(stitle, title):
    low = (stitle + " " + title).lower()
    for k, t in _TASK_MAP:
        if k in low:
            return t
    return "reference"


def _blank_split(lines, a, b, maxl):
    blanks = [i for i in range(a, b + 1) if lines[i - 1].strip() == ""]
    out = []
    cur = a
    while b - cur + 1 > maxl:
        target = cur + maxl - 1
        cand = [p for p in blanks if (cur + maxl // 2) < p <= target]
        cut = cand[-1] if cand else target
        out.append((cur, cut))
        cur = cut + 1
    out.append((cur, b))
    return out


def _split_piece(lines, a, b, heads, maxl):
    if b - a + 1 <= maxl:
        return [(a, b)]
    h3 = [h[0] for h in heads if h[1] == 3 and a < h[0] <= b]
    if h3:
        bounds = [a] + h3 + [b + 1]
        out = []
        for j in range(len(bounds) - 1):
            x, y = bounds[j], bounds[j + 1] - 1
            out.extend(_blank_split(lines, x, y, maxl) if y - x + 1 > maxl else [(x, y)])
        return out
    return _blank_split(lines, a, b, maxl)


def _merge_small(pieces, minl, maxl):
    merged = []
    i = 0
    while i < len(pieces):
        a, b = pieces[i]
        i += 1
        while i < len(pieces) and (b - a + 1) < minl and (pieces[i][1] - a + 1) <= maxl:
            b = pieces[i][1]
            i += 1
        merged.append((a, b))
    return merged


def auto_items(cfg, lines, n):
    maxl = cfg.get("max_lines", 250)
    minl = cfg.get("min_lines", 80)
    heads = []
    for i, ln in enumerate(lines, 1):
        m = re.match(r"^(#{1,3}) +(.*)$", ln)
        if m:
            heads.append((i, len(m.group(1)), m.group(2).strip()))
    h1 = [h for h in heads if h[1] == 1]
    if not h1:
        sys.exit(f"[{cfg['slug']}] auto: no H1 heading")
    # Un titre H1 sans aucun texte avant le titre H1 suivant est un separateur
    # (ex. « # PART 1 — vLLM » suivi d'un autre H1) : on le rattache au bloc
    # suivant pour ne pas produire un chunk d'une ou deux lignes.
    blocks = []
    for k, h in enumerate(h1):
        bs = h[0]
        be = h1[k + 1][0] - 1 if k + 1 < len(h1) else n
        body = [l for l in lines[bs - 1:be]
                if l.strip() and not re.match(r"^#{1,3} ", l)]
        blocks.append([bs, be, h[2], not body])
    sections = []
    k = 0
    while k < len(blocks):
        bs, be, stitle, empty = blocks[k]
        j = k + 1
        while j < len(blocks) and blocks[j][3]:
            j += 1
        if empty and j < len(blocks):
            sections.append((bs, blocks[j][1], stitle))
            k = j + 1
        else:
            sections.append((bs, be, stitle))
            k += 1
    items = []
    for idx, (bs, be, stitle) in enumerate(sections):
        h2 = [x[0] for x in heads if x[1] == 2 and bs <= x[0] <= be]
        bounds = [bs] + h2 + [be + 1]
        pieces = []
        for j in range(len(bounds) - 1):
            pieces.extend(_split_piece(lines, bounds[j], bounds[j + 1] - 1, heads, maxl))
        merged = _merge_small(pieces, minl, maxl)
        cleaned = []
        for (a, b) in merged:
            if cleaned and (b - a + 1) < minl and (b - cleaned[-1][0] + 1) <= maxl + 30:
                cleaned[-1] = (cleaned[-1][0], b)
            else:
                cleaned.append((a, b))
        merged = cleaned
        front = idx == 0 and not cfg.get("first_is_content")
        if front:
            fbase = "00-front-matter"
        elif idx == 0 and cfg.get("folder_name"):
            fbase = f"00-{slugify(cfg['folder_name'])}"
        else:
            fbase = f"{idx:02d}-{slugify(_strip_num(stitle))}"
        is_annex = "annex" in stitle.lower()
        seen = {}
        for k, (a, b) in enumerate(merged):
            hd = next((t for (l, lv, t) in heads if a <= l <= b), "")
            is_overview = a <= bs <= b
            if is_overview and (not hd or slugify(_strip_num(hd)) == slugify(_strip_num(stitle))):
                fname = "overview"
                title = stitle
            elif hd:
                fname = slugify(hd)
                title = hd
            else:
                fname = f"part-{k + 1}"
                title = f"{stitle} (part {k + 1})"
            seen[fname] = seen.get(fname, 0) + 1
            if seen[fname] > 1:
                fname = f"{fname}-{seen[fname]}"
            role = "reference" if front else ("appendix" if is_annex else "deep-dive")
            dom = "front-matter" if front else ("appendix" if is_annex else slugify(_strip_num(stitle)))
            items.append({
                "folder": fbase, "slug": fname, "title": title, "section": stitle,
                "start": a, "end": b, "anchor": "", "domain": dom, "role": role,
                "task": task_for_section(stitle, title),
            })
    return items


# --------------------------------------------------------------------------
# generic machinery
# --------------------------------------------------------------------------

MONTHS = {
    "january": "01", "february": "02", "march": "03", "april": "04", "may": "05",
    "june": "06", "july": "07", "august": "08", "september": "09", "october": "10",
    "november": "11", "december": "12",
}


def norm_ws(s: str) -> str:
    return re.sub(r"\s+", " ", s).strip()


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


def extract_actors(text, actors):
    out = []
    for label, pats in actors.items():
        for p in pats:
            if re.search(p, text):
                out.append(label)
                break
    return sorted(out)


def extract_keywords(text, title, terms):
    low = text.lower()
    kws = []
    for term in terms:
        pat = r"(?<![a-z0-9])" + re.escape(term) + r"(?![a-z0-9])"
        if re.search(pat, low):
            kws.append(term)
    base = norm_ws(title).lower()
    kws = sorted(set(kws), key=lambda t: (t not in base, t))
    return kws[:12]


def parse_anchors(text: str):
    return re.findall(r'<a id="([^"]+)">', text)


class Corpus:
    def __init__(self, cfg, lines, n):
        self.cfg = cfg
        self.lines = lines
        self.n = n
        self.chunks = cfg["chunks"]
        self.starts = self._compute_starts()

    def _compute_starts(self):
        starts = []
        for folder, slug, heading, _title in self.chunks:
            if heading <= 1:
                starts.append(1)
                continue
            a = self.lines[heading - 2]           # line heading-1
            if a.startswith("<a id="):
                starts.append(heading - 1)
                continue
            b = heading - 2                        # line heading-2
            if b >= 1 and self.lines[b - 1].startswith("<a id=") and self.lines[heading - 2].strip() == "":
                starts.append(b)
                continue
            starts.append(heading)
        return starts

    def end_of(self, i):
        return (self.starts[i + 1] - 1) if i + 1 < len(self.starts) else self.n

    def anchor_of(self, start_line):
        raw = self.lines[start_line - 1]
        m = re.match(r'<a id="([^"]+)">', raw.strip())
        return m.group(1) if m else ""

    def role_for(self, folder, slug):
        if folder in ("01-timeline-2026", "01-chronology-2026"):
            return "timeline"
        if folder == "00-front-matter":
            return "reference"
        if folder in ("12-reference", "09-appendices"):
            return "appendix"
        return "deep-dive"

    def task_for(self, folder, slug):
        for sub, task in self.cfg["task_overrides"]:
            if sub in slug:
                return task
        return self.cfg["folder_task"][folder]

    def extract_dates(self, text):
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

    def extract_actors(self, text):
        out = []
        for label, pats in self.cfg["actors"].items():
            for p in pats:
                if re.search(p, text):
                    out.append(label)
                    break
        return sorted(out)

    def extract_keywords(self, text, title):
        low = text.lower()
        kws = []
        for term in self.cfg["terms"]:
            pat = r"(?<![a-z0-9])" + re.escape(term) + r"(?![a-z0-9])"
            if re.search(pat, low):
                kws.append(term)
        base = norm_ws(title).lower()
        kws = sorted(set(kws), key=lambda t: (t not in base, t))
        return kws[:12]

    def first_heading(self, text):
        for line in text.splitlines():
            m = re.match(r"^#{1,4} +(.*)$", line)
            if m:
                return m.group(1).strip()
        return ""


def ystr(s):
    return json.dumps(s, ensure_ascii=False)


def ylist(items):
    return "[" + ", ".join(ystr(x) for x in items) + "]"


def explicit_items(cfg, lines, n):
    cs = Corpus(cfg, lines, n)
    if cs.starts[0] != 1:
        sys.exit(f"[{cfg['slug']}] first chunk must start at line 1")
    items = []
    for i, (folder, slug, _heading, title) in enumerate(cs.chunks):
        start, end = cs.starts[i], cs.end_of(i)
        if end < start or (i + 1 < len(cs.starts) and cs.starts[i + 1] != end + 1):
            sys.exit(f"[{cfg['slug']}] gap/overlap at {folder}/{slug}")
        items.append({
            "folder": folder, "slug": slug, "title": title, "section": "",
            "start": start, "end": end, "anchor": cs.anchor_of(start),
            "domain": cfg["folder_domain"][folder],
            "role": cs.role_for(folder, slug), "task": cs.task_for(folder, slug),
        })
    return items


def build_corpus(cfg, lines, n):
    auto = cfg.get("mode") == "auto"
    items = auto_items(cfg, lines, n) if auto else explicit_items(cfg, lines, n)

    if items[0]["start"] != 1:
        sys.exit(f"[{cfg['slug']}] first chunk must start at line 1")
    for i, it in enumerate(items):
        if it["end"] < it["start"]:
            sys.exit(f"[{cfg['slug']}] empty chunk {it['slug']}")
        if i + 1 < len(items) and it["end"] + 1 != items[i + 1]["start"]:
            sys.exit(f"[{cfg['slug']}] gap/overlap before {items[i + 1]['slug']}")
    if items[-1]["end"] != n:
        sys.exit(f"[{cfg['slug']}] last chunk must end at {n}")

    out = RAG_ROOT / cfg["slug"]
    out.mkdir(parents=True, exist_ok=True)
    keep = {"README.md", "NOTES.md"}
    for p in sorted(out.rglob("*"), reverse=True):
        if p.is_file() and p.name not in keep:
            p.unlink()
        elif p.is_dir():
            p.rmdir()

    manifest, anchor_index, canonical_index, folders = [], {}, {}, {}
    for it in items:
        folder, slug, title = it["folder"], it["slug"], it["title"]
        start, end, anchor = it["start"], it["end"], it.get("anchor", "")
        domain, role, task = it["domain"], it["role"], it["task"]
        section = it.get("section", "")
        body = "".join(lines[start - 1:end])
        actors = extract_actors(body, cfg["actors"])
        dates = extract_dates(body)
        keywords = extract_keywords(body, title, cfg["terms"])
        canon = cfg.get("canonical", {}).get(f"{folder}/{slug}", [])
        cid = f"{cfg['slug']}/{folder}/{slug}"
        digest = hashlib.sha256(body.encode("utf-8")).hexdigest()

        header = ["---", f"id: {cid}", f"title: {ystr(title)}", f"domain: {domain}",
                  f"role: {role}", f"task: {task}", f"actors: {ylist(actors)}",
                  f"dates: {ylist(dates)}", f"keywords: {ylist(keywords)}",
                  f"source: {cfg['source']}", f"source_anchor: {ystr('#' + anchor if anchor else '')}",
                  f"source_lines: [{start}, {end}]"]
        if auto and section:
            header.append(f"section: {ystr(section)}")
        if cfg.get("delta_of"):
            header.append(f"delta_of: {cfg['delta_of']}")
        if canon:
            header.append(f"canonical_for: {ylist(canon)}")
        header.append(f"sha256: {digest}")
        header.append("---")

        parts = ["\n".join(header), ""]
        if not body.lstrip().startswith("# "):
            parts += [f"# {title}", ""]
        parts.append(body)
        content = "\n".join(parts)
        if not content.endswith("\n"):
            content += "\n"

        fdir = out / folder
        fdir.mkdir(parents=True, exist_ok=True)
        (fdir / f"{slug}.md").write_text(content, encoding="utf-8")

        entry = {
            "id": cid, "path": f"{cfg['slug']}/{folder}/{slug}.md", "title": title,
            "domain": domain, "role": role, "task": task, "actors": actors, "dates": dates,
            "keywords": keywords, "source": cfg["source"], "source_anchor": anchor,
            "source_lines": [start, end], "canonical_for": canon,
            "words": len(body.split()), "bytes": len(body.encode("utf-8")), "sha256": digest,
        }
        if auto and section:
            entry["section"] = section
        if cfg.get("delta_of"):
            entry["delta_of"] = cfg["delta_of"]
        manifest.append(entry)
        for a in parse_anchors(body):
            anchor_index[a] = entry["path"]
        for c in canon:
            canonical_index.setdefault(c, []).append(entry["path"])
        folders.setdefault(folder, []).append(entry)

    total_words = sum(e["words"] for e in manifest)
    lines_out = ["# INDEX — " + cfg["title"], "",
                 f"Corpus `{cfg['slug']}` · **{len(manifest)} fichiers** · "
                 f"{sum(e['source_lines'][1] - e['source_lines'][0] + 1 for e in manifest)} lignes source · "
                 f"~{total_words} mots · partition exacte de `{cfg['source']}`.", ""]
    if cfg.get("delta_of"):
        lines_out += [
            f"> **Corpus delta** — volume de faits nouveaux ou corrigés, à lire "
            f"*relativement* à [`{cfg['delta_of']}`](../{cfg['delta_of']}/INDEX.md). "
            f"Le fond (contexte, historique, définition d'un événement) reste dans la KB "
            f"principale ; ce corpus ne porte que le delta. Chaque fichier est tagué "
            f"`delta_of: {cfg['delta_of']}` afin que la recherche sache que ces chunks "
            f"complètent (et ne remplacent pas) la source canonique.", ""]
    if cfg.get("relationship") == "base":
        deltas = [k["slug"] for k in CORPORA if k.get("delta_of") == cfg["slug"]]
        if deltas:
            lines_out += ["**Corpus liés (deltas) :** " +
                          ", ".join(f"[`{d}`](../{d}/INDEX.md)" for d in deltas) + ".", ""]
    lines_out += ["## Mode d'emploi", "",
                 "1. Filtrer dans `manifest.json` (ou les tableaux ci-dessous) sur "
                 "`domain`, `task`, `actors`, `dates` ou `keywords`.",
                 "2. Ouvrir 1 à 3 fichiers ciblés ; chaque fichier est une unité "
                 "thématique auto-suffisante avec un en-tête YAML.",
                 "3. Pour un événement répété dans plusieurs sections, préférer le "
                 "fichier marqué `canonical_for` (voir la table Événements canoniques).",
                 "", "## Domaines (dossiers → fichiers)", ""]
    for folder in folders:
        lines_out += [f"### `{folder}/`", "", "| # | fichier | lignes source | rôle | tâche |",
                      "|---|---|---|---|---|"]
        for j, e in enumerate(folders[folder], 1):
            rel = e["path"].split(f"{cfg['slug']}/", 1)[-1]
            lines_out.append(f"| {j:02d} | [{e['title']}]({rel}) | "
                             f"{e['source_lines'][0]}–{e['source_lines'][1]} | {e['role']} | {e['task']} |")
        lines_out.append("")
    lines_out += ["## Par tâche", ""]
    by_task = {}
    for e in manifest:
        by_task.setdefault(e["task"], []).append(e)
    for task in sorted(by_task):
        files = ", ".join(f"[{e['title']}]({e['path'].split('/', 1)[-1]})" for e in by_task[task])
        lines_out.append(f"- **{task}** — {files}")
    lines_out += ["", "## Par acteur", ""]
    by_actor = {}
    for e in manifest:
        for a in e["actors"]:
            by_actor.setdefault(a, []).append(e["path"])
    for actor in sorted(by_actor):
        def r(p):
            return p.split(f"{cfg['slug']}/", 1)[-1]
        lines_out.append(f"- **{actor}** ({len(by_actor[actor])}) — " +
                         ", ".join(f"[{r(p)}]({r(p)})" for p in by_actor[actor]))
    lines_out += ["", "## Par date", ""]
    by_date = {}
    for e in manifest:
        for d in e["dates"]:
            by_date.setdefault(d, []).append(e["path"].split(f"{cfg['slug']}/", 1)[-1])
    for d in sorted(by_date):
        lines_out.append(f"- **{d}** — " + ", ".join(f"[{p}]({p})" for p in by_date[d]))
    if canonical_index:
        lines_out += ["", "## Événements canoniques (`canonical_for`)", "",
                      "Un même événement est narré plusieurs fois (chronologie, profil, fond). "
                      "La version de fond est marquée canonique ici ; les autres occurrences "
                      "restent présentes et sont taguées `role: timeline`.", ""]
        for event in sorted(canonical_index):
            lines_out.append(f"- **{event}** — " + ", ".join(
                f"[{p}]({p})" for p in canonical_index[event]))
    if anchor_index:
        lines_out += ["", "## Ancres source → fichier", "",
                      f"`{cfg['anchor_label']}` (liens de la table des matières source) → fichier cible.", "",
                      "| ancre | fichier |", "|---|---|"]
        for a in sorted(anchor_index, key=lambda x: (len(x), x)):
            lines_out.append(f"| `#{a}` | [{anchor_index[a]}]({anchor_index[a]}) |")
    lines_out += ["", "## Carte de couverture (lignes source)", "", "| plage | fichier |", "|---|---|"]
    for e in manifest:
        lines_out.append(f"| {e['source_lines'][0]}–{e['source_lines'][1]} | {e['path']} |")
    lines_out.append("")
    (out / "INDEX.md").write_text("\n".join(lines_out) + "\n", encoding="utf-8")

    corpus_manifest = {
        "corpus": cfg["slug"], "title": cfg["title"], "source": cfg["source"],
        "source_lines": n, "source_sha256": hashlib.sha256("".join(lines).encode("utf-8")).hexdigest(),
        "chunk_count": len(manifest), "generated_by": "RAG/_tools/build_rag.py",
    }
    if cfg.get("relationship"):
        corpus_manifest["relationship"] = cfg["relationship"]
    if cfg.get("delta_of"):
        corpus_manifest["delta_of"] = cfg["delta_of"]
    corpus_manifest["chunks"] = manifest
    (out / "manifest.json").write_text(
        json.dumps(corpus_manifest, ensure_ascii=False, indent=2) + "\n", encoding="utf-8")

    return corpus_manifest


def verify_corpus(cfg, lines, n, man):
    ordered = sorted(man["chunks"], key=lambda c: c["source_lines"][0])
    assert ordered[0]["source_lines"][0] == 1, "cover must start at 1"
    assert ordered[-1]["source_lines"][1] == n, "cover must end at n"
    cursor = 1
    for c in ordered:
        a, b = c["source_lines"]
        assert a == cursor and b >= a, f"gap/overlap at {c['path']}"
        cursor = b + 1
    recon = "".join("".join(lines[c["source_lines"][0] - 1:c["source_lines"][1]]) for c in ordered)
    assert recon == "".join(lines), "reconstruction mismatch"
    src = "".join(lines)
    sa = parse_anchors(src)
    assert len(sa) == len(set(sa)), "duplicate source anchors"
    found = {}
    for c in ordered:
        body = "".join(lines[c["source_lines"][0] - 1:c["source_lines"][1]])
        assert hashlib.sha256(body.encode("utf-8")).hexdigest() == c["sha256"]
        for a in parse_anchors(body):
            found[a] = found.get(a, 0) + 1
    assert set(found) == set(sa), "anchors missing from partition"
    assert all(v == 1 for v in found.values()), "anchor duplicated"
    return len(sa)


CORPORA = [
    {
        "slug": "briefing-ia-2026",
        "title": "AI News 2026 — Reference Dossier",
        "source": "docs/RAG/briefing-ia-2026-en.md",
        "chunks": IA_CHUNKS, "folder_domain": IA_FOLDER_DOMAIN, "folder_task": IA_FOLDER_TASK,
        "task_overrides": IA_TASK_OVERRIDES, "canonical": IA_CANONICAL,
        "actors": IA_ACTORS, "terms": IA_TERMS, "anchor_label": "#sNN-M",
    },
    {
        "slug": "briefing-general-tech-2026",
        "title": "General Tech News 2026 — Hardware, Infrastructure & Consumer Tech",
        "source": "docs/RAG/briefing-general-tech-2026-en.md",
        "chunks": GT_CHUNKS, "folder_domain": GT_FOLDER_DOMAIN, "folder_task": GT_FOLDER_TASK,
        "task_overrides": GT_TASK_OVERRIDES, "canonical": GT_CANONICAL,
        "actors": GT_ACTORS, "terms": GT_TERMS, "anchor_label": "#gNN-M",
    },
    {
        "slug": "ai-industry-kb-2026",
        "title": "AI Industry Knowledge Base 2026",
        "source": "docs/RAG/ai-industry-knowledge-base-2026.md",
        "mode": "auto", "max_lines": 170, "min_lines": 70,
        "actors": KB_ACTORS, "terms": KB_TERMS, "anchor_label": "(aucune ancre source)",
        "relationship": "base",
    },
    {
        "slug": "ai-industry-kb-2026-wave6",
        "title": "AI Industry Knowledge Base 2026 — Wave 6 Consolidation",
        "source": "docs/RAG/ai-industry-knowledge-base-2026-wave6.md",
        "mode": "auto", "max_lines": 170, "min_lines": 70,
        "actors": KB_ACTORS, "terms": KB_TERMS, "anchor_label": "(aucune ancre source)",
        "relationship": "delta", "delta_of": "ai-industry-kb-2026",
    },
    {
        "slug": "frontier-models-2026",
        "title": "Frontier AI Models 2026 — Vague 1 (EN)",
        "source": "docs/RAG/Grands titres IA modèlesEN.md",
        "mode": "auto", "max_lines": 110, "min_lines": 50,
        "actors": KB_ACTORS, "terms": KB_TERMS, "anchor_label": "(aucune ancre source)",
    },
    {
        "slug": "labs-grok-platforms-2026",
        "title": "Labs, Grok, Tools & Platforms 2026 — Vague 2 (EN)",
        "source": "docs/RAG/Labos, Grok, outils & plateformes_EN.md",
        "mode": "auto", "max_lines": 110, "min_lines": 50,
        "actors": KB_ACTORS, "terms": KB_TERMS, "anchor_label": "(aucune ancre source)",
    },
    {
        "slug": "open-local-models-2026",
        "title": "Open / Local AI Models 2026 (EN)",
        "source": "docs/RAG/Modèles IA open  locauxEN.md",
        "mode": "auto", "max_lines": 110, "min_lines": 50,
        "actors": KB_ACTORS, "terms": KB_TERMS, "anchor_label": "(aucune ancre source)",
    },
    {
        "slug": "tools-platforms-2026",
        "title": "AI Tools & Platforms 2026 (Step 2)",
        "source": "docs/RAG/Outils & plateformes IAEN.md",
        "mode": "auto", "max_lines": 110, "min_lines": 50, "first_is_content": True,
        "actors": KB_ACTORS, "terms": KB_TERMS, "anchor_label": "(aucune ancre source)",
    },
    {
        "slug": "etape4-trackd-unsloth-training",
        "title": "Step 4 — Track D : Unsloth & outillage d'entraînement/fine-tuning 2026",
        "source": "docs/RAG/etape4_trackD_unsloth_training.md",
        "mode": "auto", "max_lines": 90, "min_lines": 45, "first_is_content": True, "folder_name": "unsloth-training",
        "actors": KB_ACTORS, "terms": KB_TERMS, "anchor_label": "(aucune ancre source)",
    },
    {
        "slug": "etape5-tracka-nvidia",
        "title": "Step 5 — Track A : Nvidia (2026)",
        "source": "docs/RAG/etape5_trackA_nvidia.md",
        "mode": "auto", "max_lines": 90, "min_lines": 45, "first_is_content": True, "folder_name": "nvidia",
        "actors": KB_ACTORS, "terms": KB_TERMS, "anchor_label": "(aucune ancre source)",
    },
    {
        "slug": "etape5-trackc-huawei-intel",
        "title": "Step 5 — Track C : Huawei & Intel (2026)",
        "source": "docs/RAG/etape5_trackC_huawei_intel.md",
        "mode": "auto", "max_lines": 90, "min_lines": 45, "first_is_content": True, "folder_name": "huawei-intel",
        "actors": KB_ACTORS, "terms": KB_TERMS, "anchor_label": "(aucune ancre source)",
    },
    {
        "slug": "etape4-trackc-cuda-rocm-pytorch",
        "title": "Step 4 — Track C : CUDA / ROCm / PyTorch (2026)",
        "source": "docs/RAG/etape4_trackC_cuda_rocm_pytorch.md",
        "mode": "auto", "max_lines": 90, "min_lines": 45, "first_is_content": True, "folder_name": "cuda-rocm-pytorch",
        "actors": KB_ACTORS, "terms": KB_TERMS, "anchor_label": "(aucune ancre source)",
    },
    {
        "slug": "etape5-trackb-amd",
        "title": "Step 5 — Track B : AMD (2026)",
        "source": "docs/RAG/etape5_trackB_amd.md",
        "mode": "auto", "max_lines": 90, "min_lines": 45, "first_is_content": True, "folder_name": "amd",
        "actors": KB_ACTORS, "terms": KB_TERMS, "anchor_label": "(aucune ancre source)",
    },
    {
        "slug": "etape4-tracka-vllm-sglang",
        "title": "Step 4 — Track A : vLLM + SGLang (2026)",
        "source": "docs/RAG/etape4_trackA_vllm_sglang.md",
        "mode": "auto", "max_lines": 90, "min_lines": 45,
        "actors": KB_ACTORS, "terms": KB_TERMS, "anchor_label": "(aucune ancre source)",
    },
    {
        "slug": "etape4-trackb-local-inference",
        "title": "Step 4 — Track B : pile d'inférence locale (llama.cpp, Ollama, LM Studio)",
        "source": "docs/RAG/etape4_trackB_local_inference.md",
        "mode": "auto", "max_lines": 90, "min_lines": 45, "first_is_content": True, "folder_name": "local-inference",
        "actors": KB_ACTORS, "terms": KB_TERMS, "anchor_label": "(aucune ancre source)",
    },
    {
        "slug": "etape5-trackd-servers",
        "title": "Step 5 — Track D : serveurs IA, marché et réseau datacenter (2026)",
        "source": "docs/RAG/etape5_trackD_servers.md",
        "mode": "auto", "max_lines": 90, "min_lines": 45,
        "actors": KB_ACTORS, "terms": KB_TERMS, "anchor_label": "(aucune ancre source)",
    },
    {
        "slug": "labs-hyperscalers-2026",
        "title": "Step 3 — Labs & Hyperscalers (2026)",
        "source": "docs/RAG/Labos  hyperscalersEN.md",
        "mode": "auto", "max_lines": 90, "min_lines": 45, "first_is_content": True, "folder_name": "labs-hyperscalers",
        "actors": KB_ACTORS, "terms": KB_TERMS, "anchor_label": "(aucune ancre source)",
    },
]


def main():
    RAG_ROOT.mkdir(parents=True, exist_ok=True)
    summaries = []
    for cfg in CORPORA:
        src = ROOT / cfg["source"]
        if not src.exists():
            sys.exit(f"missing source: {src}")
        lines = src.read_text(encoding="utf-8").splitlines(keepends=True)
        n = len(lines)
        man = build_corpus(cfg, lines, n)
        anchors = verify_corpus(cfg, lines, n, man)
        summaries.append(man)
        print(f"OK  {cfg['slug']}: {man['chunk_count']} chunks | {n} lines | "
              f"{anchors} anchors | {sum(c['words'] for c in man['chunks'])} words")
        print(f"    source_sha256 = {man['source_sha256']}")

    gl = ["# INDEX — RAG", "",
          "Corpus RAG de référence pour Cetas. Chaque corpus est une "
          "partition exacte de sa source, avec index et manifest.", "",
          "Un corpus `delta` ne contient que les faits nouveaux/corrigés d'une source "
          "déjà couverte par un corpus `base` ; il ne la remplace pas. `delta_of` "
          "indique la base visée.", "",
          "| corpus | titre | relation | fichiers | source | index |",
          "|---|---|---|---|---|---|"]
    for m in summaries:
        rel = m.get("relationship", "")
        if rel == "delta" and m.get("delta_of"):
            rel = f"delta de `{m['delta_of']}`"
        gl.append(f"| `{m['corpus']}` | {m['title']} | {rel} | {m['chunk_count']} | "
                  f"`{m['source']}` | [INDEX]({m['corpus']}/INDEX.md) |")
    gl += ["", "Voir aussi [README](README.md) et `manifest.json`."]
    (RAG_ROOT / "INDEX.md").write_text("\n".join(gl) + "\n", encoding="utf-8")
    (RAG_ROOT / "manifest.json").write_text(json.dumps({
        "corpora": [{
            "corpus": m["corpus"], "title": m["title"], "source": m["source"],
            **({"relationship": m["relationship"]} if m.get("relationship") else {}),
            **({"delta_of": m["delta_of"]} if m.get("delta_of") else {}),
            "index": f"{m['corpus']}/INDEX.md", "manifest": f"{m['corpus']}/manifest.json",
            "chunk_count": m["chunk_count"]} for m in summaries],
    }, ensure_ascii=False, indent=2) + "\n", encoding="utf-8")


if __name__ == "__main__":
    main()
