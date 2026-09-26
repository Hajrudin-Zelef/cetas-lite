---
id: collect-240926-vision-ia/vision-ia/mistral-debarque-dans-la-robotique-et-une-seule-camera-suffit-a-son-robot-1
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["AWS", "Anthropic", "China", "Google", "Meta", "Microsoft", "MiniMax", "Mistral", "Nvidia", "OpenAI", "xAI"]
dates: []
keywords: ["research", "agent", "agents", "benchmark", "chatgpt", "consumer", "cost", "fable 5", "gpt-live", "gpus", "grok", "grok 4"]
source: docs/RAG/clean_en/vision-ia/mistral-debarque-dans-la-robotique-et-une-seule-camera-suffit-a-son-robot.md
source_anchor: ""
source_lines: [1, 108]
sha256: e9d1c624acc5b65232524963c8dd8856613181fc433f4c86a09a9a768d642572
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

