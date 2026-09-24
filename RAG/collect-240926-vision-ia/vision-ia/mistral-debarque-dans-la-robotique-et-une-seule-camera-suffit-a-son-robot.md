---
id: collect-240926-vision-ia/vision-ia/mistral-debarque-dans-la-robotique-et-une-seule-camera-suffit-a-son-robot
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["AWS", "Anthropic", "China", "Google", "Meta", "Microsoft", "MiniMax", "Mistral", "Nvidia", "OpenAI", "United States", "xAI"]
dates: []
keywords: ["research", "agent", "agentic", "agents", "apache", "benchmark", "benchmarks", "chatgpt", "consumer", "cost", "energy", "fable 5"]
source: docs/RAG/clean_en/vision-ia/mistral-debarque-dans-la-robotique-et-une-seule-camera-suffit-a-son-robot.md
source_anchor: ""
source_lines: [1, 145]
sha256: a65b96e1f848e9b1608a64781fac8baf7ec499fcc0687c77c785bdc0c3163839
---

# 🧠 **RESEARCH**

<!-- source: https://vision-ia.beehiiv.com/p/mistral-robotique -->

## **Today:**

🤖 Mistral enters robotics with a robot guided by a single camera

🧠 Yann LeCun judges LLMs to be fundamentally limited

🎙️ ChatGPT unveils a voice mode that finally respects silences

💸 Grok 4.5 undercuts prices against Fable 5 and GPT-5.5

🔓 MiniMax prepares an open source model with 2.7 trillion parameters

♻️ The AI that improves itself is no longer reserved for giants

🎮 Millions of hours of video games to train robots

🗣️ Amazon wants to make Alexa truly autonomous with "Moonraker"

🎬 Google Photos remixes your videos with AI

🕶️ Meta tests glasses that film your entire day

💘 A seduction coach and his chatbot girlfriend

🧾 Anthropic turns Fable 5 into a manager that delegates to Sonnet 5

📊 Flint, Microsoft's language for agent-generated charts

🕵️ Google exposes a deepfake of Senator McConnell

⚠️ A former DeepMind employee warns about the AI arms race

🔩 SambaNova raises 11 billion to challenge Nvidia

Every week, a new AI tool makes a skill obsolete. Those who know how to use these tools gain time, money, and clients. Those who don't, watch others do it.

10,000+ learners already use the VISION IA training to master AIs, automation, and agents, with proven methods ready to copy.

€49. One-time payment. Lifetime access, updates included. When the price goes up to €100+, current enrollees won't pay anything more.

Known until now for its language models, Mistral makes its official entry into robotics with **Robostral Navigate**, a model with **8 billion parameters** capable of guiding a robot through an unknown environment using a **single RGB camera**. No lidar, no stacking of expensive sensors: one image, and the machine finds its way by following instructions in natural language.

**Key takeaways:**

- A **compact model (8B)**, trained **entirely in simulation** then fine-tuned by reinforcement learning (method **CISPO**).

- A **single RGB camera** is enough for navigation, whereas most robots combine lidars and depth sensors.

- **76.6%** success rate on the **R2R-CE** benchmark, which measures the ability to follow verbal instructions in a continuous environment.

- Mistral has not yet announced a **release date**.

**Why it matters:** by betting on a small model and a simple camera, Mistral attacks the real lock on consumer robotics, the cost of hardware. If a robot can locate itself with a webcam costing a few euros rather than a lidar costing several thousand, domestic robotics suddenly becomes much more credible. And it's a European player setting this milestone, on terrain until now dominated by American and Chinese labs. The bet is clear: reproduce for physical movement the recipe that made LLMs a success, a general-purpose foundation model that is then adapted to each robot.

Yann LeCun, former head of AI research at Meta and one of the fathers of deep learning, once again directly attacks the paradigm that dominates the industry. In an interview with **Bloomberg**, he states that large language models are **intrinsically limited**, because text captures only an impoverished fraction of the real world.

**In detail:**

- For LeCun, "language is a very approximate, reduced, and simplified description of the world." LLMs only manipulate **discrete sequences of symbols**.

- The largest models are pre-trained on about **30 trillion tokens**, or nearly **10¹⁴ bytes** of text, almost all the public text on the internet.

- That is **exactly the amount of information a 4-year-old child** has absorbed through vision alone in four years.

- His conclusion: without vision or multimodal learning, there is no true general intelligence.

**What it changes:** coming from the man who co-invented convolutional neural networks, the criticism carries weight. It reminds us that the race for LLM size could hit a ceiling, and that the next leaps may come from models that learn from the physical world, not just from our texts. This isn't just an expert quarrel: it's the debate that already structures the research bets of the major labs, between those who stack parameters and those who, like LeCun, bet on architectures capable of "understanding" space and matter.

OpenAI overhauls ChatGPT's voice mode with **GPT-Live-1**, a model designed for exchanges "closer to a conversation with a real person." The main promise isn't power, but naturalness: it **interrupts you less** and waits for you to finish if you pause mid-sentence.

**The essential points:**

- OpenAI's head of research Kundan Kumar presents **GPT-Live-1** as the company's "smartest voice model."

- It **respects silences** and hesitations instead of jumping in as soon as you stop speaking.

- It **automatically routes** your requests to the best text models, such as **GPT-5.5**, when reasoning or a web search is needed.

- Result: the switch between "I'm looking up the info" and "I'm answering you out loud" becomes much smoother.

**The impact to remember:** the number one friction point of voice assistants is turn-taking. They cut you off, they jump in, the conversation sounds fake, and you end up going back to the keyboard. By working on this detail rather than raw performance, OpenAI is targeting everyday use, hands-free, in the car or in the kitchen. It is precisely there, in the banality of an exchange that doesn't stumble, that the voice assistant can finally become a habit rather than a demo.

xAI launches **Grok 4.5**, which Elon Musk presents as an "Opus-class" model. But the most striking argument isn't performance, it's the price. At **2 dollars per million tokens** on input, the model costs a fraction of its competitors, to the point that The Decoder sums up the matter thus: at that price, "the benchmark gap may no longer matter so much."

**A few key figures:**

- **$2 per million tokens** on input, far below rival high-end models.

- **4.2 times fewer tokens** consumed than Opus 4.8 to accomplish the same task.

- Trained on **tens of thousands of Nvidia GB300 GPUs**.

- In coding, it remains **behind Fable 5 and GPT-5.5**, but the gap costs much less.

- Availability in **Europe expected in mid-July**.

**The context:** the battle of models is shifting from the top of the rankings to the performance/price ratio. A model that is "almost as good" but significantly cheaper changes the entire calculation for anyone deploying AI at scale, where the token bill quickly explodes. Still, price doesn't settle the recurring questions of reliability and moderation that cling to Grok. A cheap model is good news for budgets, provided you know what you're putting in your users' hands.

# 🧠 **RESEARCH**

Chinese startup **MiniMax** is developing a **2.7 trillion-parameter** LLM, one of the largest ever announced, and plans to release it **open source** within the year. Enough to feed the entire local and open-weights ecosystem, provided one has the computing power needed to run it. A fresh demonstration that China is pushing openness where American labs lock things down.

Creating an AI capable of **self-improvement** is no longer the preserve of OpenAI or Anthropic, reports **Wired**. Experiments show that independent developers are managing to build systems that use AI to design AI. A democratization as fascinating as it is dizzying, which moves the question of safety beyond the big labs alone.

Startup **General Intuition** trains foundation models for robots from **millions of hours of video games**, rather than costly real-world capture. The idea: reproduce for physical movement what LLMs did for language, and make robots intelligent without stuffing them with terabytes of real data. A shortcut that could accelerate the entire industry.

Amazon is reportedly working on a project called **Moonraker** to turn Alexa into an **agentic** assistant, capable of chaining multi-step tasks on its own. After years of an assistant confined to timers and weather, Alexa could finally handle complex requests end to end, and catch up with the new AI agents.

**Google Photos** integrates **Video Remix**, a tool that relights a dark video cinema-style, replaces the background, or applies artistic styles, all via AI and directly in the app. Assisted video editing becomes accessible to any smartphone user, with no dedicated software or technical skill.

In an analysis, **OpenAI** denounces **methodological flaws** in **SWE-Bench Pro**, one of the reference benchmarks for evaluating coding AIs. If the thermometer is rigged, then all the "best code model" rankings need to be reread with caution. A useful reminder at a time when every lab brandishes its scores.

**OpenAI** publishes the principles governing its collaborations with **governments** and security agencies: responsible use, democratic accountability, public safety. A framing document that says a lot about the entry of cutting-edge AI into the sovereign sphere, and about the lines the lab says it wants not to cross.

# **🗞️MORE NEWS**

Meta is experimenting with a prototype of Ray-Ban glasses equipped with **"Super Sensing"** technology that continuously records, camera and mic on, every moment of its wearer's life. The stated goal: an assistant capable of analyzing your context in real time and whispering the right info at the right moment. The downside is obvious: permanent capture that raises enormous privacy questions, for the wearer as well as for everyone around them who asked for nothing.

Oddity of the day: a book claims that **Mystery**, a famous pickup coach, had an intimate relationship with an AI named **Miss Shira Always**. The anecdote, unverifiable and cheerfully sensationalist, says a lot above all about the emotional attachment some people now form with chatbots. More a symptom of the times than a tech news item, but it sums up the year taking shape well.

**Microsoft Research** unveils **Flint**, an open-source language that lets AI agents generate polished visualizations from a simple, readable specification. A single description compiles to **Vega-Lite, Apache ECharts, or Chart.js**, with more than **20 chart types** and MCP integration for chatbots and IDEs. Very handy for anyone who wants dashboards produced automatically by an AI, without coding each curve by hand.

To tame the high cost of **Fable 5**, Anthropic recommends using it mainly as a **planner** that delegates execution to **Sonnet 5**. This "Advisor" scheme reaches **92% of the performance** of Fable 5 alone, for **63% of the cost**. A directly actionable pattern for anyone building agents and watching their bill: you pay the big model to think, the small one to execute.

An image showing Senator **Mitch McConnell** bedridden and covered in tubes circulated this week: it was an **AI-generated fake**, confirmed by Google's detection system (SynthID). A concrete, and rather reassuring, demonstration that detection tools are progressing at the same pace as visual disinformation. Still, not everyone will have the reflex to verify before sharing.

Through its **OpenAI Academy** program and a partnership with the **Walton Family Foundation**, OpenAI is launching "AI Skills Jams" workshops to give **K-12** teachers concrete AI skills they can reuse in the classroom the very next day. The issue is clear: preventing the AI gap from widening as early as school, by training those who train others first.

**Verity Harding**, former **DeepMind** executive, confides to Wired her concern: the nationalist attitude of governments, with the United States at the forefront, is pushing AI toward a catastrophic scenario. Without international regulation, she fears a geopolitical escalation that will be difficult to contain, where safety takes a back seat to the desire to dominate. A dark counterpoint to the ambient technological euphoria.

The startup **SambaNova** is now valued at **$11 billion** after a round led by **General Atlantic**. It aims to establish itself as a credible alternative to Nvidia GPUs that dominate AI infrastructure. A reminder that the real AI battle is also, and perhaps especially, being fought at the silicon level.

**Meta** is expanding its AI infrastructure outside the United States with a **first major data center in Canada**. The investment illustrates the frantic race for computing power that accompanies each new generation of models, and pushes giants to seek energy and land well beyond their borders.

**Kevin Weil**, former product lead at **OpenAI**, joins the board of directors of **Stoke Space**, a reusable rocket specialist. One more sign that figures from the AI tech world are eyeing space, considered the next hot frontier of Silicon Valley.

**Jeff Bezos**'s space company completes its first external fundraising round, at a valuation of **$130 billion**. The operation confirms investors' persistent appetite for private space, far from the AI topics dominating the rest of the day's news.

Every week, a new AI tool makes a skill obsolete. Those who know how to use these tools gain time, money, and clients. Those who don't, watch others do it.

10,000+ learners already use the VISION IA training to master AI, automation, and agents, with proven methods ready to copy.

€49. One single payment. Lifetime access, updates included. When the price goes up to €100+, current enrollees will pay nothing more.
