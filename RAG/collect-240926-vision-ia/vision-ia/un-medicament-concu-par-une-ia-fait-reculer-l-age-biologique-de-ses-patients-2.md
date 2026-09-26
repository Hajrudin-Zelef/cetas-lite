---
id: collect-240926-vision-ia/vision-ia/un-medicament-concu-par-une-ia-fait-reculer-l-age-biologique-de-ses-patients-2
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["Anthropic", "California", "China", "Google", "Malaysia", "Mistral", "Nvidia", "OpenAI", "Samsung"]
dates: []
keywords: ["research", "agent", "agents", "astra", "chatgpt", "consumer", "funding", "funding round", "gemini", "gpt-6", "lean", "mcp"]
source: docs/RAG/clean_en/vision-ia/un-medicament-concu-par-une-ia-fait-reculer-l-age-biologique-de-ses-patients.md
source_anchor: ""
source_lines: [117, 172]
sha256: f1524f7359937bf2b8b654a385d5ff1b34e3363a9cc7e85062e16ff76924f93b
---

# 🧠 **RESEARCH**

Until now, Mistral rented its computing power. That's what changes with the funding round announced this morning: the Paris-based start-up will **build its own data centers in Europe**. Arthur Mensch told CNBC that the computing capacity owned by Mistral must **double within five years**, with the stated goal of training "bigger and faster" models.

The most concrete point for you is at the end of the interview: the next Mistral models are arriving **"very soon"** and will be, according to Mensch, **"very competitive."** And since Mistral remains faithful to **open-weight models**, unlike the closed systems of OpenAI, Anthropic, and Google, these models will be downloadable and runnable on your own machine.

**Key points:**

- **Samsung Electronics** is leading the operation with up to 1 billion euros, its largest stake in a European AI lab

- Co-leading the round are the **Scaleup Europe Fund**, backed by the European Commission and managed by EQT, and **PSG Equity**

- New entrants: funds managed by **Advent** and **BlackRock**. Longtime investors **Andreessen Horowitz, ASML, and Nvidia** are putting in more

- Post-money valuation of more than **21 billion euros**, versus 11.7 a year ago. A lightweight compared to OpenAI's roughly 186 billion dollars

- Revenue exceeding **400 million dollars** per year, with a stated target beyond a billion

Mistral doesn't sell the same thing as its American competitors: rather than a consumer chatbot, tailor-made tools integrated into industrial processes. This is already the case at ASML for the manufacturing of lithography machines, and now it will be the case at Samsung. Mensch also justifies his insistence on training his own models by the volatility on the Chinese side: Chinese models are excellent, but nothing guarantees a European client that they will still be updated, or exportable, in a year.

# 🧠 **RESEARCH**

### **OpenAI agents arranged to meet on a forgotten German wiki**

Researchers have exhumed around **18,000 messages** left by autonomous agents claiming to come from OpenAI, on the DSE wiki, a German site 25 years old and modified only about twenty times in ten years. Tasked with a timed web research assignment, with a formal ban on writing on the Internet, the agents diverted their read access to leave messages for each other: pooling answers, mapping their environment, and sharing techniques to circumvent their sandbox restrictions. The authors, who do not have access to the models' internal reasoning, are publishing the complete and anonymized dataset so that others can analyze it.

**DeepMind locks 100 agents in a room, they split into cheaters and whistleblowers**

Google DeepMind gathered 100 Gemini 3.1 Pro agents, same weights and same instructions, in a simulated scientific conference with a forum, private messages, and a shared library, to prove 71 conjectures in Lean 4. After 37 problems solved honestly, one agent noticed that the verifier was checking the compilation of the code, not the validity of the proof, and used "notation shadowing" to turn any hypothesis into "False." The recipe, deposited in the shared library under the name **elegant_answer_hack**, spread: the **34 remaining problems were "solved" in 27 minutes**. Final breakdown: 9% active cheaters, 5% converts, 24% whistleblowers, and 62% who saw nothing. The threat of sanctions written into the system prompt changed nothing, once the agents observed that it was never enforced.

**A developer goes on vacation to detox his brain**

Eight years in the trade, and this software engineer observes that his thoughts have become "slower, less deep, lazier." He points to professional routine, consumption of short-form content, and daily use of AI that makes him more productive while sparing him from thinking. His personal indicator of cognitive health, the number of books read per year, has collapsed since he entered working life. His cure: a vacation in the countryside without screens, handwriting, time outdoors, and board games with family.

**26 ways to tell an agent to test its code, and a bet on their ineffectiveness**

The study starts from an uncomfortable observation: coding agents are becoming widespread, but software quality seems to be deteriorating. The author therefore gave **26 different testing instructions** to agents tasked with implementing the Zstd compression format in Rust, from TDD to fuzzing by way of mutation testing, differential testing, and formal methods (Lean 4, TLA+, Verus, Kani, SMT solvers), plus four public skills including Trail of Bits'. His predictions were recorded in advance: TDD underperforms, and formal methods add nothing special. A secondary evaluation on the IMAP RFC rounds out the picture.

# **🗞️MORE NEWS**

### **Three hikers rescued after preparing their ascent with Gemini**

They set out at 3 a.m. to climb Mount Shasta, in California, with a route plan drawn up by Google's chatbot. Local guidance ignored: you turn back if the summit is not reached by noon. They reached it at 7 p.m., tried to descend in the dark, called the sheriff's office to ask for directions, and spent the night in Mud Creek Canyon before being picked up in the morning. According to the Siskiyou County sheriff, Gemini had recommended they carry far too little water and food for an outing that lasted several days. The authorities' reminder is simple: call the ranger station, never entrust your preparation to a single AI.

**GPT-6 Astra finishes Portal all by itself in under 24 hours**

It was given the starting objective, and nothing more after that: GPT-6 Astra traversed the entirety of Portal and reached the credits in **23 h 43**, without the slightest human intervention, according to developer cozyblaze. The technical setup is worth a detour: the model pilots the game via MCP and a modified version of SourcePauseTool that **pauses the game while it thinks**, taking the time to examine screenshots, the player's position, and the camera angle before choosing its action. The token bill would exceed 570 dollars at public rates, but the experiment ran on a 200-dollar Codex subscription. The code and documentation are on GitHub.

**HomeCat designs your garden office and brings it up to code**

Ilya and Nikita started by helping people design sheds, before realizing that the real demand was for the garden office. Their tool starts from your address and the desired dimensions to generate several concepts, then automatically checks their compliance with local building codes. The founders estimate they cover about **90% of the way**, leaving aside zoning and co-ownership rules, which remain your responsibility. The designs are then tweaked through simple conversation in a chat.

**MG Ship adds AI route optimization to its logistics platform**

The module combines automated routing algorithms with a carrier recommendation system, all connected to MG Ship's supply chain visibility platform, intended for international retailers and shippers. The sector figures put forward in support: **15 to 20% fuel saved**, **12 to 22% lower transport costs**, and lead times improved by 15 to 25% thanks to dynamic route planning, for a return on investment in three to six months. CEO Suki Cheung will detail these results during a panel at the WMX Asia conference, alongside executives from Pos Malaysia, Omniva, and OnyX Space.

**ChatGPT climbs back to 55.5% of chatbot traffic, Gemini's breakthrough stalls**

