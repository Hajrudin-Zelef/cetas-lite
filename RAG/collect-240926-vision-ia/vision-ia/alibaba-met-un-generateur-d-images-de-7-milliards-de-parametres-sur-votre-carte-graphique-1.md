---
id: collect-240926-vision-ia/vision-ia/alibaba-met-un-generateur-d-images-de-7-milliards-de-parametres-sur-votre-carte-graphique-1
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "China", "Google", "Hugging Face", "Irregular", "Meta", "Nvidia", "OpenAI", "United States"]
dates: ["2026-05"]
keywords: ["research", "agent", "agents", "apache", "benchmark", "chatgpt", "consumer", "cost", "cybersecurity", "diffusion", "gemini", "incident"]
source: docs/RAG/clean_en/vision-ia/alibaba-met-un-generateur-d-images-de-7-milliards-de-parametres-sur-votre-carte-graphique.md
source_anchor: ""
source_lines: [1, 90]
sha256: 94ab2f0866e7a1279b4bfa64498772d867204488ab9483ee8cdcc79001131306
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

