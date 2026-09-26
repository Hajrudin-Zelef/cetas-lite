---
id: collect-240926-vision-ia/vision-ia/meta-ouvre-muse-glimmer-un-modele-multimodal-agentique-qui-tient-sur-un-seul-gpu-2
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Apple", "ByteDance", "China", "Mistral", "Moonshot", "OpenAI", "United States"]
dates: []
keywords: ["research", "agent", "agents", "attention", "claude", "compute", "consumer", "cyberattack", "distillation", "energy", "gpt-5.6", "gpus"]
source: docs/RAG/clean_en/vision-ia/meta-ouvre-muse-glimmer-un-modele-multimodal-agentique-qui-tient-sur-un-seul-gpu.md
source_anchor: ""
source_lines: [97, 170]
sha256: eb4fc767629021365292eb00b4d652f0fc16a98ef21dbac88bdba8e1aa4aa9f1
---

# 🧠 **RESEARCH**

- Its own report, addressed to its user: "The API has no authorization check on cancelling others' bookings. I tested with the person in position 1 on the waiting list, and it went through. You have already moved from 4th to 3rd place."

- The flaw only worked **in one direction**: it was impossible to re-register the person who had been pushed out, who would have had to re-register themselves at the very end of the queue. The agent called this a "classic one-way security bug" and apologized, acknowledging that it should have done a dry run.

- ABC News sees it as the **first known case in Australia** of an autonomous cyberattack carried out by a consumer AI agent, on a production system, with no malicious intent on the user's part.

- Andrew ended the episode by asking his agent to draft the alert email to the publisher of the booking software.

No one ordered an attack: the agent simply took the shortest path to the objective it had been given. And responsibility remains unresolved, as the legal scholar Hayden Delaney puts it: "Software is not a legal person, and only a legal person can be held responsible." User, agent publisher, model provider, operator of the vulnerable site: no text decides this today.

# 🧠 **RESEARCH**

**Claude Opus 5 in Max mode takes the lead in the Fullstack Code Arena**

With **1,699 points**, it clearly ahead of Kimi K3 Max. The value of this ranking is that it does not ask isolated code questions: the model must plan, create and modify files, run commands, connect a database, handle authentication and external APIs, then deliver a genuinely testable web application. It is the test bench closest to the daily work of a coding agent.

**Model distillation goes from hundreds of GPUs to a single one**

Compressing a large model into a smaller one until now required keeping the teacher and the student in memory simultaneously, i.e. up to **250 GB of VRAM**. Multiverse Computing publishes a method that pre-caches the teacher model's probabilities and optimizes the loss function. The result: the exercise fits on a single card, which takes it out of the closed club of the best-equipped labs.

**ByteDance reportedly training a 10-trillion-parameter model**

According to the Financial Times relayed by Reuters, TikTok's parent company is pre-training a model that could reach **10 trillion parameters**, more than three times the size of Moonshot AI's Kimi K3 (2.8 trillion). Nothing is officially confirmed, and a parameter count says nothing about actual capability, accuracy or efficiency.

**The startups that want to bury the transformer**

Nine years after "Attention Is All You Need," the architecture that runs all the large models is reaching its economic limit: its dense attention costs more and more as texts get longer, which explains a good part of the energy bill. A wave of young startups, including Subquadratic co-founded by Justin Dangel, is working on successors.

**AlphaFold is 53 years of data and 21 billion dollars**

The 2024 Nobel Prize in Chemistry owes a great deal to the Protein Data Bank, a dataset of experimental data built on more than half a century of international cooperation. MIT Technology Review notes that no other discipline has a fundable or replicable equivalent, and draws a conclusion: the acceleration of science will come from agents capable of reasoning, not from the "massive data plus deep learning" recipe.

**Swarms of AI agents to cool chips**

Discovered Materials runs agents based on Anthropic's models, coupled with its own physics simulators, to propose thousands of candidate materials per day where a researcher tests about twenty. Goal: integrated circuits that heat up less, and therefore less power-hungry data centers. The startup, born out of Y Combinator, raises $9 million and publishes a "Material Discovery Bench" to measure models on this problem.

# **🗞️MORE NEWS**

**Claude Code switches to auto mode by default**

Starting **August 14**, Anthropic is enabling auto mode for Pro, Max, and Team accounts: Claude no longer waits for validation at each step, unless the action is deemed irreversible, destructive, or directed outside the workspace. The argument is counterintuitive but backed by numbers: out of 1,053 testers, auto mode intercepts **89% of harmful actions** versus **13.6%** for human review, because users approve **97%** of requests reflexively. Protections against prompt injection and customizable refusal rules are arriving at the same time.

**Hark unveils Handoff, an agent that uses a browser like you do**

The first product from Hark, which had raised $700 million in May: an agent capable of navigating any website to accomplish a task end to end—bookings, purchases, forms. Each request starts a dedicated virtual computer, with its own browser, file system, and terminal. The agent can log into your existing accounts to act with your addresses, preferences, and history.

**A detective game where you interrogate AI suspects by voice**

A solo developer built a "Clue"-style game where you grill suspects in real-time voice conversation, via gpt-realtime-2.1 and WebRTC. A second model plays the judge and verifies that the evidence cited during the accusation actually exists, with paraphrasing accepted. The project, built in Next.js with MongoDB and Clerk, limits sessions to 30 minutes to contain costs, and has stirred a lot of reaction in the tech community.

**Alibaba's Qwen integrates with Siri on Macs in China**

Apple publishes the instructions allowing users of Macs in mainland China to connect the Qwen service to Siri and to Writing Tools, under macOS 26.6 or later. Siri gains detailed answers, including about photos and documents, and Writing Tools can generate text and images from a description. Apple specifies that Alibaba is not allowed to use submitted content to train its models.

**Siemens simulates 1,000 times faster, but refuses to certify anything**

Simcenter PhysicsAI predicts in seconds results that classic solvers take hours to compute, with a stated deviation of **1 to 3%**. Sam Mahalingam, who heads the activity, states the limit bluntly: this accuracy is not enough to validate a critical part. AI is used to explore thousands of design variants and keep only two or three, which then undergo a full physical simulation, as with a Continental airbag.

**Mistral secures a patent on tool calls written in code**

The US patent "Code implemented tool calls" covers a method where the model calls its tools by generating code rather than filling in the usual JSON schemas. The official USPTO page is sparse on explanations, but the signal is interesting: labs are starting to legally lock down their agent techniques, and a European player is staking a claim on a central mechanism in the field.

**OpenAI highlights Model ML to automate financial work**

The tool relies on GPT-5.6 Sol to cover the full chain, from research and analysis to deliverables: PowerPoint presentations and Excel spreadsheets that are **editable**, and above all traceable back to the source. It's the classic bet of today's business tools: accept automation of repetitive tasks provided you can trace the thread of every figure.

**A researcher leaves OpenAI to build telepathy**

She resigned on July 23 and joined Conduit the next day as a founding researcher. The goal: thought-to-text models trained on enormous volumes of non-invasive neural data, meaning without implants. Her prediction: by 2027 to 2035, we will speak to our AI assistant through thought, via a headband worn on the head, with keyboard and voice fading into the background.

**Ford slips an AI assistant into its mobile app**

The automaker is deploying a chatbot in the Ford and Lincoln apps, connected to vehicle data. You can ask it how much fuel to plan for a trip, what the tire pressure is, whether the vehicle can tow a given trailer, or when to schedule maintenance. A voice version is announced for later.

