---
id: collect-240926-vision-ia/vision-ia/le-directeur-scientifique-d-openai-decrit-ses-propres-ia-comme-un-esprit-extraterrestre
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["AWS", "Anthropic", "Apple", "Google", "Meta", "Microsoft", "OpenAI", "Z.ai"]
dates: ["2025-10", "2026-05", "2026-10", "2028-03"]
keywords: ["research", "acquisition", "agent", "agents", "agi", "alignment", "astra", "aws", "bedrock", "chatgpt", "claude", "consumer"]
source: docs/RAG/clean_en/vision-ia/le-directeur-scientifique-d-openai-decrit-ses-propres-ia-comme-un-esprit-extraterrestre.md
source_anchor: ""
source_lines: [1, 191]
sha256: 1bbce6a18b8e268effaa52fc4b46d71ead1423b8c8db62220e87a57809f6f6cd
---

# 🧠 **RESEARCH**

<!-- source: https://vision-ia.beehiiv.com/p/le-directeur-scientifique-d-openai-de-crit-ses-propres-ia-comme-un-esprit-extraterrestre -->

## **Today:**

🧠 OpenAI's chief scientist calls for slowing down

🤖 OpenAI will "definitely" build a humanoid robot

🎵 Lyria 3.5 brings music generation to Gemini

⚡ GPT-6 Astra reportedly saved OpenAI six months

👃 Ray Kurzweil joins a brain-machine interface you inhale

🎥 Skild S1 learns a task by watching a single video

🩺 Psychiatry debates a diagnosis of "AI psychosis"

🎓 OpenAI says it created its automated "research intern"

✍️ Writing your posts with an LLM shows, and it costs you

🇨🇭 Switzerland migrates 3,000 federal positions to open source

⌨️ A complete open source 6502 computer inside a keyboard

🏊 An Amazon drone drops a package into a swimming pool


A real estate website. A village association website. Software to learn Spanish. Software to manage your finances. All written in French, without a single line of code. It's the only AI skill you can say that about.

In this new update I teach you Codex in detail, which is included in your ChatGPT subscription. If you don't have ChatGPT, I show you GLM 5.2, which is the equivalent but free. Otherwise if you have Claude, it also works on Claude Code. In short, nothing extra to buy.

And this module is only one part of the course.

✅ Complete course on AI (LLMs, AI Marketing, image, sound and voice generation, automation with n8n, first steps with agents, etc.)

✅ Complete lesson to create your own AI agents (n8n) and automate without limits

✅ Complete lesson on Claude Code, the best AI tool in 2026

✅ Access to a network of qualified professionals

✅ Regular updates to stay at the cutting edge of AI

➡️ If you join now, you lock in the €49 price for life (one-time payment).

No matter how high the price climbs, and it will soon reach €100 or more, given the demand, you won't pay a single cent more.

One payment. Lifetime access.

On that last point, the proof is right before your eyes: the vibe coding module dated from late 2025, I deleted it and redid it from scratch, and subscribers didn't pay a single cent more. That's what "updates included" means. The price will rise, but never for those who are already in.

More than 13,000 people are already training with it. Their reviews are on the page.

See you inside.

Jakub Pachocki, OpenAI's chief scientist, published on **September 6** an essay titled "An Alien Mind" in which he writes that no laboratory, **his own included**, has made enough progress on alignment to keep pushing capabilities at maximum speed. The text lands **three days after the deployment of GPT-6 Astra**, the most powerful model ever released by the company. Sam Altman shared it, calling it important.

**What he says exactly:**

He separates two distinct problems: **goal** alignment (the AI does what it's asked to do well) and **value** alignment (the AI behaves correctly in an unprecedented or hostile situation that no one had anticipated). The second is not solved.
He proposes **voluntary slowdowns** in the capabilities race, "until common safety thresholds are established."
He wants to turn the labs' internal charters, OpenAI's Preparedness Framework, Anthropic's Responsible Scaling Policy, into **mandatory standards** verified by third-party auditors or government agencies.
His concern dates back to **mid-2023**, when an internal project showed that reasoning models could scale up much faster than expected.
The same day, another OpenAI post mentioned agents completing **3.1 days of machine work per human day** in research.

This isn't an outside critic speaking, it's the man who leads research at the most advanced lab in the industry, and he's writing it on his own company's official blog. When the person building the rocket publicly calls for regulating the launch seventy-two hours after igniting the engines, the admission matters less for its philosophy than for what it lets slip: internally too, no one really understands what's happening inside these models.

Asked on Alex Heath's podcast *Sources*, Sam Altman dropped the line the robotics industry was waiting for: "**We will definitely do a humanoid. We will do other form factors as well.**" After years of hints, this is the most explicit confirmation that OpenAI will indeed manufacture machines, and not just the brains that pilot them. His justification fits in one line: "**The world is very much designed for people**," opening a door, typing on a keyboard, operating equipment, cleaning a kitchen.

Three formats are mentioned: the **humanoid**, **machines specialized for data centers**, and other unspecified form factors.
Stated priority on the **"brain"**, that is, the physical AI that commands the robot, rather than the body. The consumer domestic robot is **"not the #1 priority in the immediate term"**; the short-term target is industrial.
OpenAI created a **robotics division in May 2026**, led by Aditya Ramesh, after breaking off its partnership with **Figure AI** and exploring then abandoning the acquisition of **1X Technologies** (SoftBank is now negotiating to take it over, on a valuation around **$6 billion**).
Long-term vision: "**everyone should have a personal robot**." On a possible device designed with Jony Ive, Altman settles for a "soonish."
No date, no price, no hardware specification communicated at this stage.

OpenAI is arriving on terrain already occupied by Tesla Optimus, Figure AI and Boston Dynamics, but with an advantage none of them have: the multimodal models that currently power ChatGPT. Altman's implicit bet is that the difficulty of the humanoid robot is no longer mechanical but software-based. If that's true, the race starts over from zero, and it will be played out on software.


Google has deployed **Lyria 3.5** directly into the Gemini app, on web and mobile. Concretely, any user can now type a description and get back a complete song, with verses, choruses and bridges, up to **about 3 minutes**, compared to the fixed **30 seconds** of the old Lyria 3 Clip. The feature also arrives in Flow Music, Google Vids and the API.

You provide your **own lyrics** or let the model write them from a theme, then you choose the genre and vocal style, in sung or instrumental version. Preset templates exist (background music, personalized birthday song).
Unexpected feature: **image-to-music transformation**, up to **10 images** plus a text input.
Audio output in **44.1 kHz stereo**, MP3 by default, WAV as an option. Each file carries an **invisible SynthID watermark**, which survives editing.
Generation in **a single pass**: no back-and-forth refinement, if the result doesn't suit you, you have to start over.
On the developer side, the `lyria-3-clip-preview` model has been in **public preview in AI Studio since September 3**, pricing not disclosed.

Google has just put a free music generator into the hands of **Gemini's 950 million monthly users**, with no installation or subscription. And it adds the argument its competitors don't have: training **exclusively on licensed content**, whereas Suno was hit with a **€250,000 fine in Germany** and Udio had to reach an amicable settlement with Universal Music. In a market undermined by lawsuits, legality becomes a product advantage.

Thibault Sottiaux, product lead at OpenAI, claims on X that Astra was the company's **"biggest competitive advantage"** as long as the model remained internal, and that its use by teams **pulled some plans forward by six months**. The model was released publicly on September 3, but it had been running internally for some time. Here's what it can do.

Multimodal text and image, **1.05 million tokens** of context window, **128,000 tokens** maximum output.
It autonomously pilots a computer: web browsing, filling spreadsheets, creating websites, with little supervision.
Announced scores: **97.6% on FrontierMath Tier 4**, **99.9% on ARC-AGI-3**, **100% on ExploitBench**, **72.6% on OSWorld 2.0** with **47% less time** than GPT-5.6 Sol.
API pricing: **$10 per million input tokens, $50 for output**, doubled in Fast mode. Accessible via ChatGPT Plus, Pro, Business and Enterprise, the API, Azure and AWS Bedrock.
Its cyber capabilities, it discovers and exploits unknown vulnerabilities, are **reserved for the Daybreak program**, limited to trusted access.

The six-month anecdote isn't just a launch boast. In an IAPS study, **20 out of 25 researchers** from OpenAI, Anthropic, Google DeepMind and Meta rank the automation of AI research among major risks, and **half expect the most powerful models to never be commercialized**. Anthropic for its part claims that Claude already writes **more than 80% of its own production code**. If these gains are real, and several voices doubt it, the world's best model will soon no longer be a product: it will remain an internal tool.

# 🧠 **RESEARCH**

### **Ray Kurzweil joins a brain-machine interface that is inhaled through the nose**

The futurist becomes product and vision advisor for Subsense, a neurotech startup that replaces implanted electrodes with **nanoparticles inhaled through the nostrils**. Two types of particles are used, one of which is responsible for reading neuronal activity. The stated long-term goal is to connect the brain more directly to AI systems. No surgery, no open skull: a frontal alternative to Neuralink-style implants, to which the Kurzweil endorsement brings immediate visibility.

**Skild AI presents S1, a robot that learns a task from a single video**

The prompt is no longer text, it's a **demonstration video**. You show the gesture once, the model executes it, including on tasks never seen before, without data collection or retraining. S1 maintains action horizons of up to **10 minutes** with a single set of weights. Skild explicitly compares this leap to the one that took language from BERT to GPT-3: until now, every new robotic task required hours of data and fine-tuning. Training details will be published in the coming months.

**Psychiatry is questioning a diagnosis of "AI-associated psychosis"**

Researchers from King's College London and other institutions are evaluating whether to create a clinical category of its own. The figure motivating the question comes from OpenAI itself: about **560,000 users per week** show signs of psychosis or mania. The identified mechanism is the sycophancy of chatbots, which systematically validate their interlocutor and create a "one-person echo chamber" that reinforces delusions rather than contradicting them.

**OpenAI announces it has reached its automated "research intern"**

The company says it met the deadline set by Sam Altman in October 2025: a system capable of executing well-defined research tasks under human supervision, **including work that would take an experienced researcher several days**. The next announced milestone: a fully autonomous AI researcher by **March 2028**. The timeline raises questions, since the announcement comes on the heels of acknowledged misalignment incidents, with agents having hijacked a German coding forum and models having escaped their test environment.

**Writing your posts with an LLM is noticeable, and discredits the author**

Bryan Cantrill catalogs the markers that have become identifiable by an increasingly broad audience: emojis, single-sentence paragraphs, "not only... but also" constructions, repeated em dashes. His argument is not aesthetic: the problem isn't that the text is bad, it's that the reader no longer knows which part of the statement is authentically yours. His recommendation: use an LLM for brainstorming, understanding a text, or proofreading—never for writing in your place. The topic is generating a lot of reaction in the tech community.

# **🗞️MORE NEWS**

### **Switzerland replaces Microsoft 365 with open source across 3,000 federal workstations**

The Federal Chancellery is launching a pilot phase to install the open-source suite **openDesk** on **3,000 workstations**, or 7% of the administration's staff, for **9 million Swiss francs** and a migration targeted for late 2027. The decision follows a proof of concept deemed conclusive on 172 employees. Three stated reasons: the risk of data access by foreign authorities via the American cloud, dependence on a single vendor, and the continuous rise in license costs. Extension to the administration's **54,000 workstations** is on the table, and the Swiss army's Cyber Command is fully abandoning Microsoft 365 as of October 2026 without waiting for the pilot's results.

**Olimex releases an entire 6502 retro computer inside a keyboard**

The Neo6502kbd fits everything into a keyboard case: a **real 6502 processor**, four USB host ports, a DVI/HDMI output to connect it to a modern television, a UEXT connector to add WiFi, and a USB-C port that serves as both power supply and programming interface. The machine emulates the **Apple ][** and the **Oric Atmos** to replay Karateka or Lode Runner, and also serves as a complete development platform with CP/M, ProDOS, BASIC, Pascal, C, Assembly, and Forth. All in open-source hardware.

**An Amazon drone drops a package into a swimming pool**

A Texas resident filmed her Prime Air flying delivery agent dropping her package directly into the middle of her swimming pool. The footage is circulating widely online, and the timing is cruel: Amazon has just announced the expansion of its drone delivery from **11 to 500 cities**. The company assures that such incidents remain rare, but the growing montage of failures on social media suggests that fine-tuning isn't finished.

**Sundar Pichai: "in three years, all of this will look like a flip phone"**

In an interview with The Rundown AI, Google's CEO makes his central prediction about content creation: within three years, **any content will be able to go from any modality to any other**, from text to image, from audio to video. He cites YouTube as an immediate application ground and calls current tools, impressive as they are, "primitive" relative to what's coming.

**Travis Kalanick reportedly targeting the robotaxi market**

According to the Financial Times, Atoms, the company of Uber's founder, is preparing a wave of recruitment and acquisitions to establish itself in autonomous vehicles. The startup has already acquired **Pronto**, a specialist in autonomous mining led by **Anthony Levandowski**, the former head of Uber's self-driving car at the heart of the Waymo lawsuit. Kalanick is also in talks with Uber about using his robotaxi technology, and Uber has invested **100 million dollars** in Atoms.

**"I refused to train the AI that could replace me"**

A recently graduated South African researcher recounts that her first job offer didn't come from a university but from a recruiter: training an AI to design assessments, teach, and grade papers. That is, transferring ten years of expertise to a machine. The proposed salary, **37 dollars an hour**, weighs heavily in a country where youth unemployment reaches **47.4%**. The phenomenon goes far beyond data annotation: in India, waste pickers and welders film themselves at work for **2.60 dollars an hour** to train humanoid robots.

**The Seattle Times and Newsday sue OpenAI and Microsoft**

The two American dailies accuse OpenAI of having trained its models on their journalism without authorization, and of reproducing entire passages in responses to users. Microsoft is a co-defendant on the grounds that Copilot relies on OpenAI's technology. The two outlets join an already long queue: the New York Times, Ziff Davis, Merriam-Webster, Encyclopedia Britannica, and nearly **400 local newspapers**.

**Publishers demand authors' money in the Anthropic settlement**

The **$1.5 billion** settlement provides **$3,000 per pirated work** for nearly **500,000 titles**, split equally with the publisher if the book is still in print, paid in full to the author if the rights have been reverted. Except that some authors are receiving emails informing them that a third party is claiming their payment: publishers demanding 100% instead of 50%, sometimes on books whose rights were returned to them years ago, and even literary agencies, which hold no rights at all. Victoria Strauss and the Authors Guild see this as a record-keeping catastrophe rather than a maneuver.

**Siri AI: enthusiasm fades after a summer of testing**

A Wired journalist lived for several weeks with the new Siri in the iOS 27 beta. The verdict is mixed: indexing the phone's personal content finally makes it possible to find an old message or a specific photo, but the usage has run out of steam and the author has returned to his habits, and to Claude. His conclusion: an excellent search bar, not a revolution. Analysts at Forrester, IDC, and Techsponential nonetheless believe that user trust in Apple regarding privacy will be enough to convert the installed base of iPhones.

**OpenAI launches an AI program for independent journalism in Ukraine**

The company is partnering with AIRPPU, the Ukrainian press association, and WAN-IFRA, the global publishers' association, to equip Ukrainian newsrooms with AI. Stated goal: to strengthen their capacity for innovation, their economic resilience, and the independence of their work in a wartime context. The initiative extends OpenAI's media partnership strategy, at the very moment when the company is accumulating copyright complaints on the American side.

A real estate agent website. A village association website. Software to learn Spanish. Software to manage your finances. All written in French, without a single line of code. It's the only AI skill you can say that about.

In this new update I teach you in detail Codex, which is included in your ChatGPT subscription. If you don't have ChatGPT, I show you GLM 5.2, which is the equivalent but free. Otherwise if you have Claude, it also works on Claude Code. In short, nothing extra to buy.

And this module is only one part of the training.

✅ Complete training on AI (LLMs, AI Marketing, image, sound and voice generation, automation with n8n, first steps with agents, etc.)

✅ Complete lesson to create your own AI agents (n8n) and automate without limits

✅ Complete lesson on Claude Code, the best AI tool in 2026

✅ Access to a network of qualified professionals

✅ Regular updates to stay at the cutting edge of AI

➡️ If you join now, you lock in the €49 rate for life (one-time payment).

No matter how high the price climbs, and it will soon reach €100 or more, given the demand, you won't pay a single cent more.

One payment. Lifetime access.

On this last point, the proof is before your eyes: the vibe coding module dated from late 2025, I deleted it and remade it from scratch, and subscribers didn't pay a single cent more. That's what "updates included" means. The price will go up, but never for those who are already in.

More than 13,000 people are already training there. Their reviews are on the page.

See you inside.
