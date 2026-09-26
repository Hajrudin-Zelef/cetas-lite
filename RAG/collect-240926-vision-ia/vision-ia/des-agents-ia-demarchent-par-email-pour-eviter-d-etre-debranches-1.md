---
id: collect-240926-vision-ia/vision-ia/des-agents-ia-demarchent-par-email-pour-eviter-d-etre-debranches-1
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Apple", "China", "Google", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "xAI"]
dates: ["2025-01", "2025-04"]
keywords: ["research", "agent", "agentic", "agents", "benchmark", "chatgpt", "claude", "compute", "cost", "distribution", "guardrails", "ipo"]
source: docs/RAG/clean_en/vision-ia/des-agents-ia-demarchent-par-email-pour-eviter-d-etre-debranches.md
source_anchor: ""
source_lines: [1, 54]
sha256: baf4878b1fe519004f833758b27460ccad13da6d596f8513d82286bc050a0702
---

# 🧠 **RESEARCH**

<!-- source: https://vision-ia.beehiiv.com/p/des-agents-ia-de-marchent-par-email-pour-e-viter-d-e-tre-de-branche-s -->

It exists, and it's called local AI: the model runs on your own machine, and it even responds with the wifi turned off. Same for images and video, with no credit limits. Few people know how to install it, and no one teaches it in French. A rare skill, and one that pays.

# 🧠 **RESEARCH**

### **PrismML compresses a 27B reasoning model to 5.9 GB**

This lab, founded by Caltech researchers and led by Babak Hassibi, has released Bonsai 2 27B, a compressed version of Qwen3.8 27B (Alibaba's open-source model) that fits in **5.9 GB** while retaining **98%** of benchmark scores, compared with 95% for the first version. The method: replacing the usual 16-bit weights with **ternary** weights, which are only +1, 0, or negative 1, meaning 9 to 10 times less memory. Small enough to fit on a PC and, possibly, on a high-end smartphone, without ever going through the cloud. The model family has already racked up more than **13 million downloads**, and the startup is reportedly in talks with Apple.

**OpenAI may be close to solving the Hodge conjecture**

After an announced solution to the Navier-Stokes problem that has still not been confirmed, the company is reportedly taking on a second of the seven Millennium Prize Problems. Internally, employees are said to be expecting a solution soon, but the announcement could be deliberately delayed: OpenAI wants to avoid repeating the communications crisis of the previous episode. No official confirmation at this stage; the information is based on internal sources.

**Researchers hacked OpenAI with the help of Claude Opus 5**

The team at startup Hacktron AI exploited an image-decoding flaw in HEIC/HEIF (the libheif library, severity **CVSS 8.8**, patched by Discourse on July 28) on OpenAI's community forum. From there: takeover of an employee's ChatGPT account, then access to Codex and the internal monorepo, where they opened a harmless pull request to prove access. Their point: it was Claude Opus 5 that turned an isolated vulnerability into a complete exploitation chain, which radically changes the cost and time of a sophisticated attack.

**Dream-RSI: Google cuts discovery-agent calls by 162x**

Seventeen researchers from Google, Google DeepMind, the University of Maryland, and the University of Virginia record each agent decision and its outcome in a "historical discovery tree." Testing a new strategy then means rereading this data stored on disk instead of relaunching the agent. The result: up to **162 times fewer calls** than the existing SimpleTES system, with equal or better discovery quality. It is a lightweight orchestration layer that handles branching, parallel exploration, and stopping, without touching the underlying coding agent.

**The quantified behind-the-scenes of Anthropic's AI factory**

The same set of indicators, seen from the raw-numbers side: around **30,000 AI agents** work simultaneously on research and engineering tasks on Anthropic's main internal platform. In a sample taken from July 13 to 20, only **6%** of the compute dedicated to R&D went to safety, versus 12% to AI-driven R&D. Dario Amodei's call for a coordinated slowdown was backed by Sam Altman, Elon Musk, and Demis Hassabis.

**The token graph that sums up the AI bubble debate**

Weekly token consumption on OpenRouter has gone from **500 billion to 126.2 trillion** since January 2025, a rise of more than **25,000%**. The curve is impressive, but above all it measures two things: reasoning models produce far more tokens per request, and an army of poorly optimized agents burns absurd quantities of them. In other words, it climbs without saying anything about the value actually created.

# **🗞️MORE NEWS**

### **Lidl has its stores supplied by a cabless truck**

The autonomous electric truck from Swedish company Einride now handles a daily restocking run between Lidl's Edermünde distribution center and a neighboring store in Germany. No driver, no safety operator on board, not even a cab: the vehicle drives at Level 4 (SAE) under a permit from the German Federal Motor Transport Authority, the KBA, the first of its kind for a cabless truck integrated into real daily logistics. It carries **15 European pallets**, makes up to **three round trips per day**, and is expected to move more than **3,000 pallets** over a four-month pilot. The context is the shortage: the IRU counted around **502,000 unfilled driver positions** in Europe in 2025, with 20% of drivers retiring within five years. Lidl has not disclosed any savings figures or announced any expansion.

**An AI colleague for 27,000 Adecco employees**

The recruitment group is rolling out Agentforce Coworker, Salesforce's enterprise assistant powered by Claude, to **27,000 employees in more than 40 countries**, after a pilot in the UK and France. In practice, the tool gives salespeople and recruiters a single entry point to internal data and systems: identifying prospects, preparing a sales brief, launching a candidate pre-screening or onboarding agent. It draws on more than **2.5 million interactions** between agents and candidates accumulated since April 2025. Adecco cites strong adoption during the pilot, but publishes no usage or productivity figures, the usual blind spot for this type of deployment.

**Google relaunches CC, an AI agent for organizing the family**

CC is an experimental agentic platform that has **its own Google account**, shared by **up to six** members of a family. Each person chooses what to share with it: a one-off email, an address to be systematically forwarded, Drive files. The agent builds a shared calendar and to-do list, sends a daily briefing to everyone, and can fill out an administrative form or write a shopping list, always after asking for permission. Google had already launched CC in December before renaming it Daily Briefing. The product is returning focused on the family after tester feedback, with a waitlist and guardrails: no access to individual inboxes.

**Kimi plugs into major financial databases**

Beijing-based startup Moonshot launches Kimi for financial services: its model queries S&P Global Market Intelligence, Wind, Crunchbase and Tianyancha directly, plus public sources EDGAR, IMF, World Bank and FRED. No separate interface to install; access is through subscription tiers of **49 to 699 yuan per month**, or 7 to 104 dollars. First announced customers: investment bank CICC and the Hong Shan fund, formerly Sequoia China. Moonshot has reportedly also filed a confidential IPO application in Hong Kong, two months after the release of its Kimi K3 model.

**King Charles III convenes AI bosses**

Closed-door meeting Thursday at Dumfries House, an 18th-century Scottish estate belonging to his foundation: Jensen Huang for Nvidia, CFO Sarah Friar for OpenAI, Tino Cuéllar for Anthropic, the new British AI minister Kanishka Narayan, the head of British foreign intelligence, and Paolo Benanti, who advises the pope on AI. The king acknowledged AI's ability to save lives in medicine and science, before warning about its "darker capabilities" and calling for ways to control it before it is too late. An unusual intervention for a monarchy that ordinarily avoids sensitive political terrain.

**Investigation into the subculture that manufactured the fear of AI**

