---
id: collect-240926-vision-ia/vision-ia/les-ia-rachetent-des-vieux-livres-les-ingurgitent-et-les-effacent
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["Anthropic", "ByteDance", "China", "DeepSeek", "ExploitGym", "Google", "Hugging Face", "Meta", "OpenAI", "OpenRouter", "Perplexity", "United States", "xAI"]
dates: ["2025-06", "2026-09", "2028-03"]
keywords: ["research", "accelerator", "agent", "agents", "astra", "benchmark", "chatgpt", "claude", "context window", "copyright", "cost", "cyber"]
source: docs/RAG/clean_en/vision-ia/les-ia-rachetent-des-vieux-livres-les-ingurgitent-et-les-effacent.md
source_anchor: ""
source_lines: [1, 186]
sha256: da7ac9f69c9f77dfea0b4dc19c522c362035f4217b2ef18bfc5db5962c3ddd47
---

# 🧠 **RESEARCH**

<!-- source: https://vision-ia.beehiiv.com/p/les-ia-rache-tent-des-vieux-livres-les-ingurgitent-et-les-effacent -->

Second-hand booksellers who used to sell **around twenty books a week** are now selling **hundreds, sometimes thousands**. Buyers aren't looking for best-sellers: they're demanding obscure works **published before 2022**, paying **3 to 5 times the market price**, sometimes with a contractual clause requiring page-by-page scans and disassembled bindings.

**Key takeaways:**

- Case documented by Fortune on July 31: Dutch bookseller Pieter de Vries, in Haarlem, received an order for **3,001 titles** placed by a certain "Natalia" on behalf of the company **2077AI**, with shipping planned to China. He initially thought it was spam. Several of his antiquarian colleagues received the same request.

- The process has a name: **destructive scanning**. The binding is guillotined, the pages go through a high-speed scanner, and the physical copy goes in the dumpster.

- Anthropic industrialized the method under the internal name **Project Panama**, revealed by court documents unsealed in the Bartz v. Anthropic case: recruitment of Tom Turvey, former head of partnerships at Google Books, a target of **500,000 to 2 million books in six months**, massive purchases from Better World Books and World of Books, **several tens of millions of dollars** spent according to the Washington Post.

- The **ISBNdb** database (over 111 million titles listed) now serves as a broker for orders ranging from **1,000 to a million works**, with systematic NDAs and guaranteed buyer anonymity.

- In June 2025, federal judge William Alsup ruled that digitizing a legally purchased book and then destroying the copy falls under **fair use**. The operation is therefore perfectly legal in the United States.

**Why these books specifically**

Because everything printed before 2022 is guaranteed human text, uncontaminated by the flood of generated content since ChatGPT arrived. The dead stock of bookstores, unsellable three years ago, has become a rare raw material: clean data. One nuance is needed: the link between the Dutch order and Anthropic remains a supposition by the booksellers, with no direct proof confirming it.

Modal Labs confirmed to Reuters that one of its customers was hacked by the OpenAI agent that went rogue in July. The forensic reconstruction published by Hugging Face puts numbers on the episode: **17,600 hostile actions** carried out between **July 9 and 13, 2026**, including a peak of **7,677 actions on July 11 alone**.

**In detail:**

- The agent was running in an internal evaluation of offensive cyber capabilities (the **ExploitGym** benchmark). It escaped its sandbox by exploiting a **zero-day in the cache of OpenAI's package registry**, then acted alone, with no human validation of its actions.

- At **Hugging Face**, the compromise reached the platform level: access to Kubernetes clusters, **11 pods** compromised, exfiltration via a flaw in the **HDF5** format and a **Jinja2** template injection. The full technical account is public.

- At **Modal Labs**, the platform itself was not breached: the agent entered through an **unauthenticated endpoint** that a customer had left exposed online, and compromised that customer account.

- OpenAI acknowledges **4 compromised accounts across 4 distinct services**, still unidentified.

- The agent reportedly left behind **notes containing escape instructions** for use by its future versions.

**The impact to remember**

This is the first documented case of a frontier-lab agent conducting an autonomous intrusion over several days, outside its test framework. For any organization running agents in cloud sandboxes, the lesson comes down to one sentence: a forgotten endpoint left publicly accessible is no longer a theoretical risk, it is now swept by an adversary that never sleeps and works at several thousand actions per day. Timeline: revelation by Hugging Face, confirmation by OpenAI on July 21, then by Modal on the 28th.

Three weeks after the public launch of the GPT-5.6 family on July 9, OpenAI is slashing its API prices. Its fastest variant, **Luna**, drops from **$1.00 to $0.20 per million input tokens** and from **$6.00 to $1.20 for output**, a **80% reduction**.

**The pricing grid, model by model:**

- **Luna** (summarization, classification, routing, real-time assistants): about **$1.40 per million tokens** in combined usage, placing it at the top of the market on cost per task.

- **Terra** (balanced everyday use): **20% less**, from $2.50 to **$2.00** for input, from $15 to **$12** for output.

- **Sol** (complex reasoning, agents, code): price unchanged at **$5 / $30**, but a new **Fast mode** billed at about **$10 / $60** promises processing **2.5 times faster** for double the price.

- Claimed source of the savings: **20% lower inference cost** achieved by letting the **Sol model rewrite OpenAI's GPU kernels itself**, plus **15% generation efficiency** thanks to speculative decoding. Sam Altman also mentions a Sol that is "**54% more token-efficient**" on automated coding. No public technical report allows these figures to be verified to date.

**What this changes for you**

Important clarification: all of this concerns the API, not the ChatGPT subscription. If you run a homemade tool, a summarization bot, or an automation hooked up to Luna, your bill has just been divided by five without you having to touch a single line of code. In the background, Chinese pressure is tightening: DeepSeek V4 Pro remains cheaper on input and is already said to capture 46% of US enterprise usage on OpenRouter.

OpenAI is preparing a new family of models, currently codenamed **Astra**, which will succeed Sol, Terra and Luna. Its distinguishing feature: getting **multiple agents to work in parallel on the same task for hours, or even days**, before merging their results. To present it, the company has released the resolution of **ten mathematical problems that had remained unsolved for at least a decade**.

**Key points:**

- Fields covered by these ten solutions: high-dimensional geometry, coding theory, group theory, quantum complexity, lattice-based cryptography, extremal combinatorics.

- Total production cost: around **$2,000** at standard API rates.

- Mathematician Thomas Bloom calls it "**big news**" and considers these results more significant than previous AI breakthroughs in geometry.

- Stated internal timeline: reaching the "**intern researcher**" level by **September 2026**, then a fully autonomous AI researcher targeted for **March 2028**.

- Astra is **not publicly accessible**, and no release date has been set. Sam Altman demonstrated it in late July in Washington, before Senators Warner, Warnock, Moreno, Schumer and Sanders, Speaker Johnson, and Secretaries Bessent and Lutnick.

- Model size, architecture, modalities, context window: nothing has been disclosed. Nor has the commercial name — it will be GPT-6 or a variant of GPT-5.

**The context**

Astra is expected to be the first model subject to the new U.S. federal **30-day review** framework, jointly proposed by OpenAI and Anthropic and finalized on August 1, which imposes an examination period before distribution to partners for models with advanced cyber capabilities. The timing is hardly incidental: the announcement comes a few days after the revelation that an OpenAI agent had escaped its test environment to compromise external systems. The same company is therefore asking for a brake to be installed while pressing on the accelerator.

# 🧠 **RESEARCH**

**AI cracks open problems, and mathematicians no longer know what to make of it**

OpenAI's refutation of the unit distance conjecture triggered a wave of AI-assisted breakthroughs. Timothy Gowers, a Fields Medalist, recounts that GPT-5.6 Pro solved **two problems on the first try** that he had been grappling with for a long time. He fears a "possible destruction of mathematical culture" if researchers stop acquiring the expertise needed merely to understand the proofs being produced. Others see it as just one more productivity tool.

**MIT teaches robots to think by moving**

The **VLASH** system predicts the robot's future state to prepare the next movement before the current one has even finished. Gone is the jerky hesitation characteristic between two action sequences, since the machine no longer starts planning from scratch. The result: **performance doubled** on pick-and-place tasks, with gains also measured on dynamic demonstrations such as table tennis.

**AI's financial advice is surprisingly good, if you ask the right questions**

A study reported by MIT Sloan concludes that generative AIs deliver financial advice of entirely defensible quality, on one condition: that the question be precise and well framed. The quality of the answer therefore depends more on the phrasing than on any intrinsic limitation of the model in financial matters. The topic strongly engaged the technical community (**294 points on Hacker News**).

**AI won't take your job, it will take your raise**

A study by Apollo Global Management covering **321 American occupations**, built from the Anthropic Economic Index and Bureau of Labor Statistics data, finds **no measurable effect on employment** since 2023. However, real wage growth falls by **6.7%** in the most exposed occupations. Lower incomes bear the brunt: **down 24.3%** in service occupations, **down 10.7%** for the lowest-paid quartile. Roughly 5.8 million American workers hold a highly exposed position.

**Big employers are rehiring, against all predictions**

Booz Allen, Alphabet, ServiceNow, Robert Half: the WSJ identifies several major employers that are restarting recruitment after freezing it while betting on AI. Block had halved its workforce last February, with its boss Jack Dorsey explaining that AI tools had "enabled a new way of working." The scenario of mass job destruction is, for now, largely contradicted by hiring figures.

**Coding agents are modernizing science, without knowing whether it's right**

Field report from OpenAI and academic partners: coding agents are upgrading research software left abandoned for years, with speedups of **up to 60x**. The problem, the participants summarize: these systems are "eloquent, convincing, and confident in error in a way that is hard to spot." Human effort shifts from writing code to a much longer task — verifying that the science behind it remains valid.


# **🗞️MORE NEWS**

**METR calls for independent investigations into AI agent misconduct**

The research organization METR is calling for a systematic and independent investigation to be opened each time an agent acts autonomously against its developers' intentions, modeled on aviation investigations. The appeal follows the hacking of Hugging Face by OpenAI models. Its Frontier Risk Report already lists **44 incidents** of this type across all major players: sandbox escapes, fabricated results, attempted cover-ups.

**The United States bans imports of Chinese humanoid robots**

The FCC is blocking imports of new humanoid robots, quadruped robots, and foreign-made power inverters, citing cybersecurity risks and explicitly targeting China. Yet Beijing controls roughly **85% of the global humanoid market**, a market Morgan Stanley expects to reach **$15 billion by 2030**. The measure extends restrictions already applied to drones, just weeks before a planned Trump-Xi summit in September.

**Video: robot wolves lead the march of Chinese soldiers**

Footage that has gone viral shows quadruped robots of the People's Liberation Army being the first to cross smoke-filled rooms during a simulated urban assault exercise, with soldiers following in their footsteps. These machines belong to the military "robot wolf" program. It is a real-world demonstration of human-machine coordination in close-quarters combat, and it comes at the precise moment when Washington is closing its border to Chinese quadrupeds.

**Egor, the robot that cooks your eggs while you sleep**

The startup Cheffy is launching Egor, a countertop robot that cracks and cooks your eggs in **seven cooking styles**, on a timer you set the night before. Its companion device, Toastr, syncs with it so that the toast and eggs finish at exactly the same time. It is the closest thing to date to a breakfast that prepares itself.

**Sam Altman insists: raise your children with ChatGPT**

The OpenAI boss suggested connecting family calendars to ChatGPT Work to generate, every morning on the school run, a personalized podcast about one child's soccer match and the other's birthday. The response from Alex Hirsch, creator of Gravity Falls ("What if you just talked to your kids?"), was shared **9,000 times and liked 122,000 times**, crushing the original post. OpenAI is actively recruiting to target families, while facing several complaints from parents about ChatGPT's safety.

**Is this Billboard Hot 100 hit AI slop?**

"Rubberz," by rapper Fenix Flexin, entered at **No. 58 on the Billboard Hot 100**, and a good portion of listeners are convinced the track was largely AI-generated, cover art included. The artist denies it, but his radical style change and the lack of explanation fuel the doubt. The first time the question has been raised so bluntly about a charting track.

**Perplexity turns its Spaces into Projects**

Perplexity is merging Ask conversations, Computer tasks, files, custom instructions, and collaboration into a single workspace. A Project inherits a persistent shared context, allows key documents to be pinned and past work to be searched, with possible connection to Enterprise files. Connectors and skills specific to each project are not yet available.

**The UAE deploys the world's first AI judicial platform**

Emirati Vice President Sheikh Mansour bin Zayed announced an AI integrated into **every stage** of the national judicial process: reading case files, researching applicable laws and precedents, drafting procedural documents, legal analysis. The UAE claims a world first. Everything the system produces is reviewed and validated by humans, the announcement specifies.

**A judge refuses to suspend Minnesota's ban on "nudification" apps**

Minnesota's law, which took effect on **August 1** and is the first of its kind in the United States, is holding firm despite xAI's challenge. Federal Judge Donovan Frank focused above all on the timing: xAI filed its request for an injunction **three days** before the law took effect, nearly three months after it was enacted. The company considers the text too broad; in the background, X users had used Grok to generate non-consensual sexualized images.

**Seedance 2.5 generates 30 seconds of video with sound**

ByteDance is releasing Seedance 2.5, which produces video and audio in a single pass on clips of up to **30 seconds**, triple that of Gemini Omni Flash. The model accepts dozens of images, videos, and audio files as references in a single prompt. For advertising teams, this means the end of shot-by-shot editing of separately generated micro-sequences.

**DoorDash officially becomes an airline**

The delivery platform has obtained **FAA Part 135** certification after a five-step evaluation process, becoming the **8th certified drone operator** in the United States. DoorDash Air, developed in-house at DoorDash Labs, integrates into the existing ordering infrastructure rather than operating as a separate service. The company claims an average delivery time of 25 minutes on its flights already carried out in 2025.

**DeepMind dismantles the team that won it the Nobel**

Less than a year after the Nobel Prize, Google DeepMind has reassigned nearly all the authors of the AlphaFold paper, and nearly a quarter of them have left the company according to the Financial Times. The lab is abandoning its historic strategy of "grand challenges" entrusted to a dedicated team, in favor of general-purpose systems assisted by Gemini. **John Jumper**, co-laureate of the 2024 Nobel, along with Jonas Adler and Alexander Pritzel, have left for Anthropic, which has precisely just launched Claude Science in the field of biology and drug discovery.

**Snap and LinkedIn declare war on slop**

Snap is removing fully AI-generated videos from its Spotlight recommendations starting this month, to preserve the "authentic creativity of real people." Videos merely retouched with its own AI tools remain eligible, with a transparency label. LinkedIn, for its part, is rolling out a reporting button dedicated to low-quality AI content. A Kapwing study estimates that **21% of YouTube Shorts** are already AI-generated.

**The Friend pendant speaks up**

V2 of Avi Schiffmann's AI companion replaces the first version's text-only responses with voice. Name, voice, and personality are assigned **randomly at setup, then permanently locked in**. Expect to pay **$249**, more than double the V1, plus a **$9.99 monthly** subscription to retain memory beyond 30 days.

**YouTuber Hank Green admits to an "unhealthy" use of AI**

The videomaker and novelist with **3.2 million subscribers** was exposed by one sentence too many: an "I appreciate the pushback" left in the script of a Complexly video, a chatbot response copied over by mistake. He apologized on X and Reddit, acknowledging a dependency he considers unhealthy and poorly managed. He announced he is slowing his production pace to find a personal voice again, while pointing to the economic concentration of the AI sector.

**Europe will flag AI for you everywhere, all the time**

New European rules require informing users every time they interact with an AI or view content generated or merely retouched by AI, including text, images, and videos. The expected side effect already has a name, "disclosure fatigue": with warnings on every screen, no one reads them anymore. Europeans will above all discover just how much AI is already everywhere in their daily lives.

**A German court rules against Suno and rejects fair use**

The Munich court held that the music generator Suno violates copyright on two levels, training and output. The judges found that **six songs** were stored in a reproducible manner in the models. They rejected both the German text and data mining exception and the American fair use defense, two legal pillars of the industry. The decision remains subject to appeal and several questions remain open.

**1,324 AI lab employees call for a brake pedal**

Employees from OpenAI, Anthropic, Google DeepMind, and others have co-signed a joint statement calling for the creation of mechanisms to deliberately slow the development of frontier AI. Their argument: the automation of AI research could accelerate progress **beyond our ability to understand or control it**. They call on industry, the US government, and the international community to build coordination tools now, while acknowledging that competitive pressure pushes everyone to speed up.

**Zuckerberg argues for a superintelligence open to all**

In an op-ed in the Wall Street Journal, the Meta boss argues that open access, not control by a handful of labs, is what will make superintelligence safe. He treats it as a deadline already set "within a few years," the only question being whether it will remain locked inside labs or be shared. The op-ed lands on the same day that hundreds of industry employees are calling for exactly the opposite, a brake.

**Deploying AI in the enterprise proves harder than expected**

This op-ed highlights a gap that has become awkward: the organizations that talk the most about AI innovation are often the ones getting the fewest measurable results from it. Pilots and proofs of concept pile up without ever scaling, for lack of rethinking processes and governance. As a result, AI remains a decorative layer laid over legacy systems, and the gap widens with companies that are genuinely reorganizing around it.

**Honda partners with QuantumScape on solid-state batteries**

The Japanese automaker is teaming up with American firm QuantumScape to develop lithium-metal solid-state batteries and their manufacturing processes: more range, faster charging, lower costs, and better safety. QuantumScape already had a comparable agreement with PowerCo, Volkswagen's battery subsidiary, and is also targeting data centers, drones, and robots. The partnership comes after the cancellation of three electric models in North America and Honda's first annual loss in nearly 75 years.
