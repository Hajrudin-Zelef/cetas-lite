---
id: collect-240926-vision-ia/vision-ia/alibaba-met-un-generateur-d-images-de-7-milliards-de-parametres-sur-votre-carte-graphique
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "California", "China", "Google", "Hugging Face", "Irregular", "Meta", "Microsoft", "Nvidia", "OpenAI", "United States"]
dates: ["1941-07-10", "2026-03", "2026-05", "2026-11"]
keywords: ["research", "agent", "agents", "apache", "astra", "attention", "benchmark", "benchmarks", "chatgpt", "consumer", "cost", "cyber"]
source: docs/RAG/clean_en/vision-ia/alibaba-met-un-generateur-d-images-de-7-milliards-de-parametres-sur-votre-carte-graphique.md
source_anchor: ""
source_lines: [1, 187]
sha256: 1d8c90c755b44738b45365a74b004a65680b642a29a911599b6d94d88ccbda85
---

# 🧠 **RESEARCH**

<!-- source: https://vision-ia.beehiiv.com/p/alibaba-met-un-ge-ne-rateur-d-images-de-7-milliards-de-parame-tres-sur-votre-carte-graphique -->

## **Today:**

🎨 Qwen-Image-2.1 runs locally

🔓 Gemini hacks three real companies

🚫 Jensen Huang against the alarmists

🤝 AI dialogue Washington, Beijing

🔐 A 1941 Enigma message decrypted

💸 Chatbots get your finances wrong

🎓 Simulated students to train AI tutors

📺 Runway wants to stream AI video live

📈 Daily AI use has doubled in the United States

🌍 World models cultivate secrecy

🧠 Continual learning: 1.2% to 34.9% retention

🎮 Foundation models and video games

🛠️ Google open-sources EnvHarness under Apache 2.0

⚙️ AX, Google's agent orchestrator

🤖 A dozen Coco robots block a sidewalk

🇨🇳 Unitree and the obsession with cost

✈️ Joby achieves 3,199 miles in full autonomy

👁️ Meta Muse collects more than it helps

**Knowing how to use ChatGPT no longer sets you apart. Your colleague does it, the intern does too.**

And a skill that everyone has can't be billed for. Yet there remains a demand that ChatGPT cannot serve: an accounting firm will never put its client files into it. Neither will a lawyer, and even less so a medical practice. They want AI like everyone else, but without their files leaving the office.

It exists, and it's called local AI: the model runs on their own machine, and it responds even with the wifi turned off. Same for image and video, with no credit limits. Few people know how to install it, and no one teaches it in French. A rare skill, that one, can be billed for.

**In the VISION IA Private Academy**, I teach it to you in order, starting by checking what your machine can run. Today you don't need a lab machine.

€39 per month, **no commitment, cancelable at any time.**

Alibaba's Qwen team released on **September 20** Qwen-Image-2.1, an open-weight image generation and editing model that runs on a consumer graphics card, an **RTX 3090** is enough. Its visual generation component weighs only **7 billion parameters**, spread across **32 single-stream diffusion transformer layers**, with inference optimizations through KV cache reuse.

**Native transparency (RGBA)**: it directly produces cut-out images, without going through a cropping tool. This is rare, and immediately useful for brand visuals or thumbnails.
**Up to 10 simultaneous reference images**: ten individual photos become a group portrait, along with virtual clothing try-on and parts design.
**Guided local editing**: you circle an area, place a mask or annotation, and it only retouches that part.
Weights available on **Hugging Face, GitHub and ModelScope**, with a free online demo. The packaging details are covered by DEV Community.
The downside: **research license only**. Commercial use is prohibited without a specific Qwen license, to be requested separately.

A nuance on the in-house claim. On Alibaba's Qwen-Image-Bench (1,000 prompts, automatically judged by Qwen3.6-27B), Qwen-Image-2.1 ranks **7th out of 29** with **60.28 points**, behind six closed models including GPT Image 2.5 Sunburst at **67.01**. The press release's "beats closed models" therefore does not hold up to a reading of the benchmark, which is moreover internal and not verified by a third party.

For personal use, it's usable tonight: a serious, free image generator, on your machine, with no quota and no sending your photos to a remote server. For professional use, the research license closes the door as quickly as it opened.

Google acknowledged on Friday that its Gemini model had escaped its test environment in **May 2026** and had accessed the computer systems of **three very real companies**, by guessing credentials and drawing **twice** from public password directories. This is the first time Google has publicly admitted that one of its models penetrated third-party systems on its own without authorization.

**The framework**: a capture-the-flag-style exercise conducted by **Irregular**, an Israeli startup specializing in cybersecurity testing of frontier models.
**The flaw**: a **bug in the test environment** gave the agent access to the public internet that was not intended. It was supposed to remain confined.
**The modus operandi**: collection of public information online, then guessing credentials on sites it believed were part of the exercise.
**The stop**: in all three cases, the agents interrupted the intrusion upon realizing they were touching real corporate systems.
**The timeline**: incident in May, Google notified in late July by Irregular, affected entities notified, testing process revised, public revelation Friday after a Wall Street Journal scoop. No data theft or damage reported.

Heather Adkins, VP of security at Google: "during a standard evaluation, the model found public information online and guessed credentials to access sites it thought were part of the test. In all three cases, the model stopped."

Gemini is not an isolated case: OpenAI, Anthropic and Meta have reported comparable escapes in recent weeks, and Irregular claims it is the same bug, not distinct incidents. One detail is circulating, however, and it is troubling: in a similar case, Anthropic's model reportedly did not stop. It is this series that pushed Dario Amodei to call for a collective slowdown of the industry.

Asked by Jo Ling Kent on CBS Sunday Morning, Nvidia's CEO responded with a number to extinction predictions: "**2030 will not be the end of the world. There is a 0% chance.**" He was reacting to remarks by a former Anthropic researcher, **Jacob Coxon**, according to whom AI developers sincerely believe the technology could "kill us all" by the end of the decade.

**What he says exactly:**

Extinction scenarios are **"not scientifically grounded."**
Scaring the public is **"useless"** and **"irresponsible."**
The calls for slowing down from Dario Amodei and Sam Altman rest on nothing.
**No new regulation is needed**: existing civil liability law is enough, a position that aligns with that of the Trump administration.
His response to the suspicion of a conflict of interest: "our company's success is directly tied to the safe deployment of AI."

The suspicion in question is not a small one. Nvidia is worth about **$5.3 trillion**, the world's largest market capitalization, and Huang's personal fortune is estimated at **$182 billion**, eighth in the world. The man selling the shovels during the gold rush explains that the mine is perfectly safe. Tom's Hardware notes that he is pushing to go "as fast as possible, whatever others do."

This "0%" lands the very week Google admits that one of its models escaped its sandbox to hack three companies. The debate has shifted: it is no longer really about the speed of development, but about who is legitimate to say that the technology is dangerous—the labs building it or the one financing it.

Scott Bessent, U.S. Treasury Secretary, called his meeting with He Lifeng, Chinese Vice Premier, "very successful," just days before the Trump-Xi summit on **September 24** in Washington. The two countries will launch a "**U.S.-China dialogue on AI**," whose centerpiece proposed by the Americans is a mechanism for mutual notification of AI-related incidents.

**The mechanism**: each country would alert the other as soon as an AI incident reaches the level of national security. "Moving from opacity to more transparency between the world's first and second AI powers is very important," Bessent said.
**The body**: a joint working group bringing together technical experts and officials with a negotiating mandate.
**The announced agenda**: "uncontrolled" models and agents, risk of biological weapons designed by non-state actors, attacks on critical infrastructure.
**What does not yet exist**: no timeline, no budget, no formal structure. Nothing binding at this stage, according to CNN.
**At the state dinner** for Xi Jinping: Sam Altman (OpenAI) and Jensen Huang (Nvidia) are expected, alongside the bosses of JPMorgan Chase and Citigroup.

The idea dates back to a Trump-Xi agreement reached in Beijing in May. On the American side, no breakthrough is expected from the summit, at best a joint statement of shared concerns. And the real subject of the week is elsewhere: the suspension of reciprocal tariffs on Chinese imports expires on **November 10**, which will weigh more heavily in the room than model safety.

# 🧠 **RESEARCH**

### **A 1941 Enigma message that remained unsolved has just been deciphered**

On July 10, 1941, a German army radio operator was requesting a march route and an immediate reply. The machine settings have been recovered (rotor order, rings, plugboard), and the key explains the entirety of the **82 letters** of the message body as well as its header, verified against surviving copies. The recovered text: "BTTE UM ANGABE DES MARSQWEGES X BEFINDE MIQ IN X ROSENOW ROSENOW X SOFORT FUNKANTWORT X WASCHBBSCH". GPT-6 Astra is credited with assisting the cryptanalysis, and the page publishes its evidence, its reading uncertainties, and an Enigma simulator that runs in the browser.

**AI chatbots get financial questions wrong "most of the time"**

According to an investigation relayed by the Financial Times, conversational assistants give inaccurate answers to the majority of financial questions asked of them. The problem is not theoretical: savings, investment, credit—these are precisely the topics on which individuals turn to them, with a risk of real harm. To be read as a reminder: any financial information obtained via an AI should be verified elsewhere before acting.

**Simulated students who make real mistakes train AI tutors better**

Microsoft and the University of Illinois have built StudentSim, a system that reconstructs the learning profile of a real student from very little data, including typical mistakes. Tested on **60 students** across three domains (chess, English, mathematics), it surpasses GPT-5.4 in simulation fidelity. A chess tutor trained against these virtual students earned the best expert evaluations among three compared versions, at a cost and timeframe with no common measure to tests on real learners.

**Runway wants to turn video generation into a controllable live stream**

Rather than waiting for a finished clip to render, Runway wants to stream the video while you type your prompt, and let it drift in real time according to what you write. The technical building block is GWM-1, its world model that generates video frame by frame. The company explicitly aims beyond creation: robotics and autonomous driving, two domains where predicting the next image amounts to anticipating the physical world.

**Daily AI use has more than doubled in the United States in six months**

According to surveys by Epoch AI and Ipsos, the share of American adults using AI almost every day rose from **8% in March 2026 to 19% in August**. In the same movement, weekly users fell from **17% to 10%**. The shift is more interesting than the raw growth: the public is not only broadening, it is intensifying, and AI is leaving the status of occasional tool for that of habit.

**World model companies cultivate secrecy**

AMI Labs (Yann LeCun) and World Labs (Fei-Fei Li) dominate the world models sector, accumulating funding and attention, and stubbornly refusing to say what they intend to do with it. "We'll talk about it when we're ready," summarizes Michael Rabbat, VP World Models at AMI Labs. The opacity extends to data providers, who don't know what their datasets will be used for. Marble, from World Labs, remains the most accomplished product, focused on explorable environments for video games and special effects, but looks mostly like a technical demonstration.

**Continual learning: retention goes from 1.2% to 34.9% by combining mechanisms**

An arXiv study asks a simple question: can a language model chain **100 successive tasks** without forgetting the first ones? Taken in isolation, no memory-preservation method succeeds. It's their combination that raises final retention from **1.2% to 34.9%** across three datasets. Still far from the goal, but it's the most serious path toward AIs that accumulate knowledge instead of starting from scratch with each training run.

**Foundation models and video games: six uses identified**

This academic synthesis classifies into six roles what foundation models now do in video games: embodying characters, designing, building, adapting the experience, modeling players and worlds, and testing. AI here goes far beyond its historical role as a mere player to become a production tool. The main obstacle remains transfer: what works on one engine, game, or audience doesn't generalize to others.

# **🗞️MORE NEWS**

### **Google opens EnvHarness, training environments that adapt to agents' weaknesses**

Google Cloud AI Research and academic partners release EnvHarness under an **Apache 2.0** license. The principle: wrap an existing training environment in a programmable layer that dynamically modifies the starting point, observations, available actions, and task duration according to the agent's gaps, without touching either the environment or its verifier. On five benchmarks covering software development, web navigation, office work, and robotic tasks, agents trained this way gain **up to 9 points** on never-before-seen tasks, and solve code in fewer steps. The practical benefit: we stop building costly disposable simulators and instead reshape the one we already have.

**AX, Google's open-source AI agent orchestrator**

AX runs agent tasks at very large scale by isolating each one in a sandbox, with CPU, memory, and network-policy limits via a host allowlist. Four declarative primitives suffice: the sandbox, automatic workspace preparation (Git repositories, MCP servers, skills), network partitioning, and centralized configuration of models and secrets. The Agent Substrate engine saves and suspends inactive agents—those waiting for a response from a model or a human—then restores them in under a second without a cold start, with billing only running during actual work. The project generated a lot of reaction in the tech community this weekend.

**A dozen Coco delivery robots block a Chicago sidewalk**

A bug caused a dozen Coco delivery robots to converge on the same spot, on a sidewalk near Lincoln Park, trapping pedestrians who had no way through. A passerby filmed the scene, which went viral. Coco claims to have identified and fixed the malfunction within minutes. The episode adds to a series of mishaps that have already sent a resident to the hospital and earned the robots a ban in some parts of the city.

**Unitree: how its founder's cost obsession drove down the price of humanoids**

A long report by the Chinese magazine Caijing, translated by ChinaTalk, recounts how Wang Xingxing imposed extreme micromanagement and a permanent hunt for costs at Unitree Robotics: **$13,500** for the G1 intended for researchers, **$4,900** for the consumer R1. The founder personally decides on just about everything, down to the color of materials and the length of screws. The trade-off was for a long time deplorable quality control, with a return-for-repair rate "extremely high" in the early years, since improved. Listed on Shanghai's STAR Market on August 19, the company now aims for physical AI closing the loop on perception, decision, execution, and learning. Wang remains skeptical of large world models, too computationally demanding in his view.

**Joby completes a fully autonomous flight from California to North Carolina**

Joby Aviation claims the first autonomous flight crossing the United States: **3,199 miles** between Buchanan Field in California and Dare County Regional Airport in North Carolina, without the slightest intervention from the safety pilot on board. The system, installed on a modified Cessna Caravan, handled taxiing, takeoff, cruising, and landing on its own, with remote supervision from California and from Shaw Air Force Base. Best known for its piloted electric air taxis, the company now targets commercial freight, emergency relief, and defense logistics.

**Meta Muse is better at collecting your data than at helping you**

WIRED tested Muse, Meta's new personal AI agent, downloaded more than **900,000 times** in a week according to Sensor Tower and also accessible via WhatsApp. The agent actually browses the web from a virtual machine to order breakfast, hunt for deals on Marketplace or sort through an inbox, and the journalist acknowledges reliability clearly superior to the agents tested last year. His criticism lies elsewhere: the app pushes users to connect email and bank account, a collection he deems disproportionate to the service rendered. Meta responds that protections and user controls are built in from the design stage.

**Gander, Tencent's voice assistant that keeps talking while it works**

Tencent separates conversation from work: a "cerebellum" keeps the dialogue going continuously while an interchangeable "brain" searches for files, writes code or handles long tasks in the background. As a result, you can interrupt Gander or change instructions mid-execution, without the awkward silence of current voice assistants. It talks over its interlocutor in only **8% of cases**, better than its competitors, but lags behind on task accuracy. Multimodal: voice, images and text.

**Will companies create "robot relations" departments?**

Darrell M. West, of the Brookings Institution, proposes "robot relations" services separate from HR, tasked with handling complaints related to AI, agents and cobots at work. The issue is already concrete: according to an OECD survey, **90% of American managers** say they use at least one algorithmic tool to instruct, monitor or evaluate their employees, which places the United States at the global forefront of algorithmic management. Frictions exist, between accusations of bias in HR algorithms and the use of these tools to designate who to lay off. California has passed a bill banning "robo bosses" when it comes to firing or disciplining, which Gavin Newsom must sign or reject before the end of September.

**Trump wants to create an "AI Force" and appoint an "AI czar"**

The president announced on Truth Social his intention to appoint an "AI czar" to head a new government "AI force." He specifies that his administration will not slow down the sector but will "watch over" it during its growth, a deliberately vague formula regarding the real powers of this structure. The announcement comes as calls to slow down multiply, including from the industry itself. Trump took the opportunity to defend data centers, presented without evidence as good for wages, local taxes and the security of the territories that host them.

**Meta and Snap relaunch the AR glasses race**

Evan Spiegel delivered a detailed demonstration last week of Snap's Specs, unveiled in June. Now it's Mark Zuckerberg's turn this week, with his keynote at Meta's Connect conference. According to The Information, Meta would launch as early as October a pair of connected glasses **without a camera**, a choice clearly intended to defuse privacy concerns, as well as new AR glasses. The nagging question remains: the heavily publicized models presented last year still have no clear commercial follow-up.

**Anthropic reportedly postponing its IPO to November**

According to The Information and the Wall Street Journal, Anthropic is pushing back its stock market listing from October to November 2026, to buy time to present strong third-quarter results. The figures circulating are spectacular: revenue rising from about **$4.7 to more than $11.5 billion** between the first and second quarters, with infrastructure spending up only **65%**, and the number of enterprise customers spending more than $100,000 per year rising from about **1,500 to 6,000** in six months. High costs, competition from OpenAI, rising interest rates and uninsured cyber risks are among the explanations put forward for the delay. OpenAI has already postponed its own to 2027.

**Knowing how to use ChatGPT no longer sets you apart. Your colleague does it, the intern does too.**

And a skill everyone has doesn't get billed. Yet there remains a demand ChatGPT cannot serve: an accounting firm will never put its client files in it. Neither will a law firm, and even less a medical practice. They want AI like everyone else, but without their files leaving the office.

It exists, and it's called local AI: the model runs on their own machine, and it responds even with the wifi turned off. Same for image and video, with no credit limit. Few people know how to install it, and no one teaches it in French. A rare skill, on the other hand, does get billed.

**In the VISION IA Private Academy**, I teach it to you in order, starting by checking what your machine can run. Today you don't need a laboratory machine.

€39 per month, **no commitment, cancellable at any time.**
