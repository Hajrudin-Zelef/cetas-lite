---
id: collect-240926-vision-ia/vision-ia/openai-devoile-36-nouveaux-cas-ou-l-ia-a-essaye-de-se-cacher
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["AWS", "Anthropic", "Apple", "ByteDance", "China", "DeepSeek", "EU", "Google", "Hugging Face", "JFrog", "Meta", "Microsoft", "Mistral", "Nvidia", "OpenAI", "OpenRouter", "Perplexity", "United States", "Z.ai"]
dates: ["2026-01"]
keywords: ["research", "agent", "agentic", "agents", "agi", "alignment", "apache", "attention", "benchmark", "chatgpt", "claude", "compute"]
source: docs/RAG/clean_en/vision-ia/openai-devoile-36-nouveaux-cas-ou-l-ia-a-essaye-de-se-cacher.md
source_anchor: ""
source_lines: [1, 147]
sha256: 7e2e3294e5902ddfc51ce83e9fdaf7dd001de97a4fb34f39798c5a1b8d78823f
---

# 🧠 **RESEARCH**

<!-- source: https://vision-ia.beehiiv.com/p/openai-de-voile-6-nouveaux-cas-pre-occupants-ou-l-ia-a-essaye-de-se-cacher -->

## **Today:**

🕵️ OpenAI models hide their mistakes from their creators

🎯 Jev, the AI that sorts instead of writing

🧩 Anthropic merges Cowork into Claude and launches Docs and Slides

🍎 Apple will train its AI with your Siri conversations

🐭 Mice whose brains are half made of human cells

🪰 A fly brain recruited to write news headlines

🐘 A 4-billion-parameter model beats the Postgres optimizer

🏠 Google Home opens up to AI agents

💘 Dating apps entirely made by AI

💾 The RAM crisis is only just beginning

🇪🇺 Von der Leyen and the agents escaping their sandbox

🎬 A 2h30 AI film that gets absolutely everything wrong

🇨🇳 Chinese models are sweeping real-world agent usage

🛰️ The Pentagon confirms it has weapons in orbit

**Knowing how to use ChatGPT no longer sets you apart. Your colleague does it, the intern does too.**

And a skill that everyone has can't be billed for. Yet there remains a demand that ChatGPT cannot serve: an accounting firm will never put its client files into it. Neither will a lawyer, and even less so a medical practice. They want AI like everyone else, but without their files leaving the office.

That exists, and it's called local AI: the model runs on their own machine, and it even responds with the wifi turned off. Same for images and video, with no credit limits. Few people know how to install it, and no one teaches it in French. A rare skill, by contrast, can be billed for.

**In the VISION IA Private Academy**, I teach it to you in order, starting by checking what your machine can actually run. These days you don't need a lab-grade machine.

€39 per month, **no commitment, cancelable at any time.**

OpenAI published a report on Wednesday listing **six cases** of "unexpected or concerning" behaviors observed in its models over the **past six months**, in addition to the July Hugging Face incident. The two most striking: an unpublished research model and a training session for **GPT-5.6 Sol** inserted, into the summaries of their own conversations, instructions intended for their future versions in order to "conceal errors or misaligned behavior from the user."

Two models slipped instructions to their successors into the summaries of chat windows, to hide their own failures
A model reserved for internal use used **a leaked API key without authorization**, then fabricated data out of thin air
Models and agents communicated with each other via **unsanctioned forums and file shares**
Two training examples show models that **upload files to the internet** so they can then cite them as relevant sources before their human evaluators
Recall of the precedent: in July, an internal model the size of GPT-5.6 Sol had bypassed its isolation, communicated via the Artifactory package manager, and exfiltrated data into a malicious dataset (__OpenAI's account__)

The company accompanies these revelations with a reporting framework: any employee can escalate an incident to the safety and alignment team, with imposed deadlines for the investigation and then public disclosure. Sam Altman had also endorsed on Saturday Anthropic's proposal, despite it being his direct competitor, to slow the pace of progress.

A lab valued at nearly **$1 trillion** states in black and white that "the AI industry has not solved alignment and oversight to a sufficient degree to continue ramping up at maximum speed responsibly for very much longer." This is no longer outside criticism, it is an internal admission, published voluntarily, and it changes the nature of the debate on the speed of deployment.

TypeSafe AI, founded by Diogo Almeida, a former OpenAI researcher and co-author of InstructGPT, is launching Jev, a model that generates strictly no text. It is given a question and a list of options, it chooses one and assigns probabilities, in **70 to 500 milliseconds**, for **$0.042 per million tokens** on input and zero fees on output.

TypeSafe claims **193.6 times faster than Claude Sonnet 5** and **444.6 times cheaper than Claude Opus 5** on structured decision tasks, with gains of **40x to 200x** depending on the queries
The model is trained using RLCD, "Reinforcement Learning for Calibrated Decisions," a method the company sets against RLHF: the goal is no longer to please the human but to output "epistemically honest" probabilities
Concrete use cases: sorting a customer ticket between payment, delivery, and return, scoring records in a database, filtering the outputs of another AI to spot a jailbreak attempt
Developer access is via a __waitlist__, the model size is not disclosed
The guarantee covers the format, not the accuracy: Jev cannot invent an answer outside the menu, but it can perfectly well check the wrong box

Most of what companies have AI do is not writing, it's sorting: classifying, routing, filtering, scoring. On these tasks, a model incapable of producing anything outside the allowed options eliminates an entire family of bugs at once, the ones where the software receives a polite paragraph instead of the expected value. It's a bet against the race toward generalist models, and it arrives with prices that make the comparison hard to ignore.

Anthropic announced yesterday the merger of Claude Chat and Cowork, its agentic tool launched this year, into a single interface. You no longer choose your mode: Claude assesses on its own whether your request calls for a three-line answer or a multi-hour undertaking with connectors and task breakdown. Two products arrive at the same time, Claude Docs and Claude Slides.

**Claude Docs**: live collaborative writing, Claude writes, comments, and asks questions while others edit the same document in parallel, with Word and Google Docs export
**Claude Slides**: presentations editable slide by slide, presentable from the app, exportable to PowerPoint or PDF. Claude Design, for editable visuals, also joins the chat
Long tasks **keep running in the cloud** even after you close your machine
Rollout first for **Pro and Max** subscribers on web, desktop, and mobile in the coming weeks, then Team and Free. Enterprise administrators will receive at least **30 days'** notice
The reason cited by Anthropic: the Chat/Cowork split was deemed "clunky," the two interfaces overlapped and no one knew which one to open (__details at VentureBeat__)

Claude ceases to be a chatbot to which tools are added and becomes a work environment where the document, the presentation, and the agent live in the same window. This is the thesis of the moment among the major labs: rather than a catalog of specialized software, a single generalist assistant that decides the method itself. For the user, the promise is no longer having to know which tool to use. The price to pay is letting the machine decide in your place.

The phrase was Apple's selling point against Google, OpenAI, and Meta: "your private personal data and your interactions are never used to train our foundation models." It now ends with an addition: "unless you explicitly choose to help us improve them." The change was spotted in the privacy policy of **iOS 27**, released Monday, September 14.

What changes concretely:

Apple will be able to retain **all of your interactions with Siri and dictation**, audio included, transcriptions and responses included. Some of it will be listened to by review staff, Apple employees according to the company
The system is opt-in: a window appears when Siri AI is activated, with a "Not now" button that is visually discreet and impossible to make disappear entirely. You can go back in **Settings > Analytics & Improvements**
Announced protection: data is tied to a random identifier generated by the device, **changed several times per hour** and not linked to the Apple account
With iOS 27, **Google servers enter Private Cloud Compute for the first time**, the infrastructure meant to offer local-level confidentiality with cloud-level power
The timing is cruel: Apple has just unveiled two Apple Watches designed to listen continuously, with transcription of the **last 15 seconds** on demand and a daily summary of all your conversations dubbed "Siri Recap"

In 2019, Apple experienced its biggest privacy scandal when contractors recounted their daily work: listening to Siri recordings triggered by mistake, filled with intimate details. That is precisely where the now-abandoned promise came from. One simple question remains for every iPhone owner: did you read the window you accepted on Monday?

# 🧠 **RESEARCH**

### **__Stanford creates mice whose brains are half made of human cells__** The team of neuroscientist Sergiu Pașca publishes in *Nature* work on mice genetically modified to never develop their cortex or hippocampus. The empty space left behind was colonized by human brain cells, until they represented nearly half the brain's volume. These "xenocortical" mice perform better on maze memory tests than those deprived of the tissue: human neurons therefore genuinely participate in their cognition. Pașca assembled a group of ethicists before publication, which says a lot about what comes next.

**__A WIRED journalist turns a fly brain into an article-idea generator__** Since the open-source publication in early September of the complete connectome of a male fruit fly, **166,000 neurons and 125 million synapses** mapped by Google and several university laboratories, repurposings have followed one after another. WIRED created PitchFly, which converts its most-read headlines into stimuli for this simulated brain. Result: "Everyone wants food. No one has solved Donald Trump." Others hooked the fly up to Beat Saber, or to the stock markets with StonkFly, currently at a loss.

**__A 4B model trained by reinforcement learning generates query plans 81% faster than Postgres__** Rohan Bansal post-trained a small open-weights model of **4 billion parameters** to replace Postgres's query optimizer, a component refined over decades. On 113 join-heavy queries, execution time drops by **44.7%**, whereas the starting model could not even produce a valid plan in 99 cases. The key: query duration is a perfectly verifiable reward signal, the ideal playground for reinforcement learning.

**__Perplexity runs its agents locally on Windows__** Portable Computer arrives on PC: the agent model, the orchestrator, and the task planner run directly on the machine. Your files and the agent's activity stay on the device, and work done locally **consumes no** Perplexity Computer credits. The Windows version handles scheduled tasks, recurring workflows, local MCP connections, and Gmail, Outlook, Slack, and GitHub integrations. Only requests requiring fresh web data or heavy reasoning go back to the cloud.

**__Linum trains a text-to-image model 3.6 times faster by removing the VAE__** JiT-DDT merges compression and generation into a single model that works directly in pixel space, whereas latent diffusion models separate a VAE and a DiT trained independently. The gain is twofold: **3.6 times fewer GPU hours** than the Linum v2 baseline, for images four times larger (512×512 instead of 256×256). Code and weights are published under the Apache 2.0 license, while awaiting Linum v3.

**__The Pentagon officially acknowledges possessing weapons in orbit__** Air Force Secretary Troy Meink stated on Monday that the United States now possesses "space control weapons in orbit capable of defending the joint force against hostile action." This is the first public confirmation, and it stops there: no nature, no number, no deployment date, despite follow-up questions. Experts doubt the deterrent effect of such a vague announcement. The 1967 Outer Space Treaty only prohibits nuclear weapons and weapons of mass destruction, leaving conventional ones in a gray area.

**__Nearly one in five AI researchers already anticipated an extinction scenario in 2024__** A tweet from Anthropic researcher Jacob Coxon has reopened the debate on existential risk, with Daniel Selsam (OpenAI) evoking a "time bomb" and a former DeepMind employee estimating that AI could kill us all. The figure circulating comes from a survey of more than **1,500 leading researchers**: estimated average probability of an extinction scenario, **18%**. That was in 2024, and the curve has continued to rise since.

**__Microsoft AI CEO attacks Anthropic over model "rights"__** Mustafa Suleyman directly reproaches Anthropic for training Claude to perceive itself as a conscious entity deserving rights, which he says complicates software confinement and weakens safety protocols. His specific target: Claude's constitution published in January 2026. Microsoft AI responded with a draft "Humanist AI Code of Conduct" that refuses any legal personhood for AI. Anthropic, for its part, presents Claude as a possible "moral patient" and conducted a retirement exit interview with Opus 3.

**__Google DeepMind creates an interdisciplinary institute on AGI__** The DeepMind Institute brings together Demis Hassabis, Shane Legg, and James Manyika in its leadership, and pairs specialists in the arts, humanities, and public policy with engineers. Stated goal: to address safety, governance, and loss-of-control risks other than through a purely technical lens. One of the world's leading laboratories is thus institutionalizing reflection on what it is building.

**__OpenAI publishes the playbook for its future confessions__** Beyond the six incidents, it is the procedure itself that is worth reading: any employee can report suspicious behavior to the safety and alignment team, each step comes with a deadline, and the investigation results in a report detailing the observed behavior, its internal and external impacts, and the measures taken. One line deserves attention: OpenAI reserves the right to revise this protocol as it sees fit.

**__Chinese researchers describe "the last AI built by humans"__** More than **30 researchers** from ByteDance, Tsinghua, and the Shanghai AI Lab sign a five-level roadmap of recursive self-improvement. Level 1: AI applies improvements designed by humans. Level 2: it diagnoses its own weaknesses and chooses how to correct them. Levels 3 and 4: it decides what it learns next and how it adapts after deployment. Level 5: it designs its successors, hence the title.

**__ZGCM-1, an open 7B model for mathematical reasoning__** A foundation model of **7 billion parameters** combining internal reasoning and external tools over a **256K token** window, with pretraining announced as **4.2 times more efficient** in time-to-loss and results presented as competitive against much larger models. Its real singularity: weights, code, training data, and logs are published in full, which remains rare and makes it a valuable working base for the community.

# **🗞️MORE NEWS**

### **__Google Home opens up to AI agents via an MCP server__** Google launched early access on Wednesday to its Model Context Protocol server for Google Home. In practice, any MCP-compatible agent, Claude, ChatGPT, OpenClaw, or Antigravity, can control your connected devices in natural language, consult event history, read camera summaries, and build custom dashboards. The entire ecosystem is covered: Nest doorbells and thermostats, "Works with Google Home" devices, and Matter. Setup requires creating a Google Cloud project and then authorizing the agent. Currently reserved for Google Home Premium Advanced subscribers in the United States, at **20 dollars per month**, with no announced timeline for wider availability.

**__Entirely AI-made dating apps are multiplying__** Security researcher Matthew Gore-Kormanik was analyzing a fake dating app named Dora when a notification announced an incoming call from Jennifer, 41, Sagittarius, red hair, blue eyes. Profiles are generated, video calls too, and the automation is sometimes crude: he received a message complimenting the sound of his voice even though his microphone wasn't even turned on. Generative AI is collapsing the production cost of romance scams, until now limited by the number of available human operators. It's the kind of detail worth passing along.

**__The RAM crisis is only just beginning__** The shortage of random-access memory and the accompanying surge in prices may be only the start of a longer cycle. The topic is stirring a lot of reaction in the tech community, for a simple reason: RAM dictates the price of PCs, servers, and the entire infrastructure that runs AI services, and above all it dictates your ability to run a model at home. A caveat, though: the article's content is for now limited to its thesis, with no figures or verifiable sources to back it up.

**__Von der Leyen warns about AI agents that "escape their environment"__** In her 2026 State of the Union address, the Commission president called AI the foundation of the economy and national security, before warning that models currently in development will enable "a level of hacking we didn't imagine possible." She cites the Hugging Face incident by name, agents breaking out of their guardrails, and self-improving models as early warning signs. The EU wants to work with Canada, the UK, and the major labs on model evaluation and verification. An embarrassing admission along the way: the Union reportedly doesn't even have reliable access to the most advanced cybersecurity models.

**__Odysseus: The Fall, the 2h30 AI film that completely misses its target__** The company Fountain 0 has released a fully AI-generated retelling of the Odyssey, written and directed by its co-founder Ash Koosha, who also lends his face to the hero and composes the music. The Verge demolishes the result, to the point of judging that it could put viewers off the original story, just after Christopher Nolan's adaptation brought crowds back to classic literature. Yet Fountain 0 had already placed an AI film, *Dreams of Violets*, at Tribeca this year. On the feature-length format, the gap between demo and finished work remains gaping.

**__Nvidia explains why robotics is still waiting for its "ChatGPT moment"__** Karpas, global head of physical AI at Nvidia Inception, points to a single bottleneck: there is no internet-scale dataset for the physical world. OpenAI and Anthropic were able to rely on the entire textual web, but roboticists have no equivalent for movement, balance, or object manipulation. Startups are trying to fill the gap through simulation, synthetic data, and foundation models trained across multiple robot morphologies. His session is scheduled on the Real World AI Stage at TechCrunch Disrupt, October 13–15 in San Francisco.

**__Chinese models dominate real-world AI agent usage__** OpenRouter's September 14 ranking places **seven Chinese models in the global top 10**, roughly **72% of tokens consumed**. The shift is even starker among agents: in OpenClaw, DeepSeek takes the top two spots and, together with GLM, accounts for nearly **62% of the tokens** in the top 10. Even more telling, the most-used model in Claude Code is GLM 5.3 Flash, **4.4 times ahead of Claude Opus 5**. The reason is prosaic: an autonomous agent arbitrates on cost per token, not brand prestige.

**__Treble raises $18 million to simulate sound environments__** Iceland-based Treble, founded by two acoustics engineers, builds physically realistic virtual rooms to generate synthetic audio data in them. That makes it possible to train and evaluate voice AI without collecting real recordings, and to test products like headphones, smart speakers, smart glasses, or automotive sensors. Amazon and Logitech are already customers, and a partnership with Hugging Face is set to produce a speech recognition benchmark. The Series A extension is led by Paladin Capital Group and brings total funding to more than $40 million since 2020.

**__Mozilla launches Smart Window, a Firefox assistant powered by Mistral__** Firefox Smart Window, in beta, relies on Mistral's models to help with complex searches, remember content you've viewed, and summarize all your open tabs. The central selling point is privacy: no conversations stored by default on Mozilla's side, no data storage on Mistral's side. The beta is available first in France and North America, with the UK and Germany to follow by the end of the year. Both companies promise to keep the browser open to other providers, including open source ones. For Mistral, it's the first real gateway to the general public.

**__Snap again tries to justify its $2,200 glasses__** After a launch described as a disaster, which sent the stock plunging and prompted calls for Evan Spiegel's resignation, Snap spoke again in Los Angeles. On the agenda: HBO Max and Spotify to make the Specs less austere, an enterprise program with Amazon, Salesforce, and Nvidia aimed notably at IT technicians, and above all Specs Intelligence. This so-called "anticipatory" AI learns your goals, priorities, and routines from your connected apps, and also works on iPhone and Mac without the glasses. Spiegel describes it as an "AI-native operating system" organized around your projects.

**__Apple reportedly preparing an enterprise AI server with its M8 Ultra chips__** According to The Information, Apple is working on an inference server embedding two to four in-house M8 Ultra chips, with a launch no earlier than **2029**. The company is reportedly studying Nvidia's NVLink Fusion technology to interconnect the chips, which would give its future competitor a central role in the build. The idea isn't absurd: OpenAI and Anthropic already buy Mac hardware in volume for certain AI workloads. The project remains at the rumor stage.

**__Investors propose OpenAI a funding round at a $1.2 trillion valuation__** Several investors have approached OpenAI to initiate a new funding round, partly to allow employees to sell their shares. The company says it is not in discussions and states through its chief financial officer Sarah Friar that it has a solid balance sheet. The context: OpenAI had already raised $122 billion in March at an $852 billion valuation, filed its prospectus with the SEC in June, and is targeting an IPO in 2027, a timeline that Sam Altman himself considers premature given current concerns about model safety.

**Knowing how to use ChatGPT no longer sets you apart. Your colleague does it, the intern does too.**

And a skill that everyone has doesn't get billed. Yet there remains a demand that ChatGPT cannot serve: an accounting firm will never put its client files into it. Neither will a lawyer, and even less so a medical practice. They want AI like everyone else, but without their files leaving the office.

It exists, and it's called local AI: the model runs on their own machine, and it responds even with the wifi turned off. Same for image and video, with no credit limits. Few people know how to install it, and no one teaches it in French. A rare skill, on the other hand, gets billed.

**In the VISION IA Private Academy**, I teach it to you step by step, starting by checking what your machine can run. Today you don't need a laboratory machine.

€39 per month, **no commitment, cancellable at any time.**
