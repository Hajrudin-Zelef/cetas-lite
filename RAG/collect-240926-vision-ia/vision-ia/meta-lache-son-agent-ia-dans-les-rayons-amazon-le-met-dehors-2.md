---
id: collect-240926-vision-ia/vision-ia/meta-lache-son-agent-ia-dans-les-rayons-amazon-le-met-dehors-2
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["AWS", "China", "DeepSeek", "Google", "Hugging Face", "Microsoft", "MiniMax", "OpenAI", "United States", "xAI"]
dates: []
keywords: ["agent", "agents", "attention", "benchmarks", "chatgpt", "cost", "deepseek", "energy", "gemini", "gemini 3.8", "grok", "grok 4"]
source: docs/RAG/clean_en/vision-ia/meta-lache-son-agent-ia-dans-les-rayons-amazon-le-met-dehors.md
source_anchor: ""
source_lines: [91, 140]
sha256: 457c178cac342ca6765aed2e4d5b2530d5a06365a86c752d4049ea4254a18223
---

# 🧠 **RESEARCH**

The electronics and electric vehicle manufacturer lands **46 on the Artificial Analysis Intelligence Index**, tied with Grok 4.7 released the same day, ahead of Gemini 3.8 Flash and DeepSeek V4.1. The model is under an **MIT license**, freely downloadable on Hugging Face and runnable on your own hardware, with **1 million tokens of context** and text, image, audio and video inputs. Via API, count on $0.435 per million input tokens and $0.87 for output, versus $0.14 and $0.28 for the Flash version. A Pro-UltraSpeed variant promises generation up to 20 times faster.

**The model that arbitrates AI agents' decisions flips with a single sentence**

Jev, TypeSafe's decision model adopted within days by Vercel, Cloudflare, LangChain and Langfuse, writes nothing: it rules (allow an action, choose a tool, classify an input) in 70 to 500 milliseconds for $0.042 per million tokens. The problem is that it treats the data it is given as trustworthy: an Octomind test drove the probability of blocking the destructive command `rm -rf ~/.ssh` down from **0.76 to 0.48** by slipping in a fake mention of prior approval. LangChain now excludes tool outputs from what the classifier sees, so that content retrieved by the agent cannot self-authorize, and Pydantic insists: this type of guardrail complements deterministic checks, it never replaces them.

**RetroChimera: Microsoft's AI that writes the recipe for your molecules**

Published in Nature and released as open source, code and weights included, RetroChimera automatically proposes synthesis routes for making a target molecule from purchasable components. It merges two complementary retrosynthesis models and learns to rank their proposals, which makes it better than each taken separately. In a blind test, PhD-level chemists preferred its predictions to those of earlier models, and even to reactions published in the scientific literature. It recovers rare reaction types and transfers without retraining to proprietary datasets.

**The "ChatGPT moment" for robot brains may come in 2027**

The founder of Chinese startup Spirit AI, which develops the embedded models serving as the brains of humanoid robots, places the breakthrough equivalent to what ChatGPT was for text as early as next year. If the prediction holds, humanoids would move from scripted demonstrations to machines capable of understanding a real scene and acting on it. China is putting resources into the topic and openly aims for leadership against American players.

**Underwater drones to cover 98% of the oceans by 2028**

Ifremer has validated two prototypes of Deep-6000 floats capable of descending to **6,000 meters**, where the pressure is 600 times that at the surface, making France the third country in the world to reach these depths after China and the United States, following five successful dives between January 11 and February 2. Thirty abyssal floats will join the international Argo fleet and its 4,000 devices by 2028. The blind spot they fill is significant: **10% of ocean warming occurs below 4,000 meters**, which most current floats do not reach. Each device weighs 40 kg, lasts seven years without recharging, and follows a ten-day cycle: nine days of drifting at depth, then three hours of profiling.

**Amazon wants to design antibodies by computer**

Amazon Bio Discovery publishes three works on the design of therapeutic antibodies, including **MochiBind**, built on the protein language model ESM-2, which compares the binding strength of two antibodies instead of predicting an absolute value that is difficult to compare from one experiment to another. A third work entrusts an AI agent with choosing binding sites and has produced antibodies validated in the laboratory against a novel cancer target. The goal: to shave off the six to twelve months usually needed to go from a biological target to a drug candidate. Amazon also points out the weakness of current benchmarks, reliable on known targets, much less so on new ones.

**What if we approved medical AI like a drug?**

Researchers at the University of Bristol propose a framework called "Learning Ensemble" to evaluate AI systems used in medicine, modeled on drug approval procedures. Three axes of verification: the system's limits, fairness between patient groups, and suitability to the real clinical context. The idea is to catch models that shine on paper and make serious mistakes at the patient's bedside. The authors' central argument: medicine has long known how to regulate mechanisms it does not fully understand.

**Grok Voice Transcribe 2.0: twice as accurate, at the same price**

xAI doubles the accuracy of its transcription model without touching the price: **$0.10 per hour of audio** in batch processing, $0.20 in streaming. It is trained on real, noisy, multilingual recordings, and targets difficult audio: on short voice commands, the word error rate drops from **20.6% to 6.8%**. It handles dozens of languages with automatic detection and switching during recording, and ranks first in accuracy among 32 streaming models on Artificial Analysis's public leaderboard. Existing API integrations get the gain without a single line of code to modify.

**Compressing an LLM like a physicist**

Multiverse Computing reformulates the removal of entire blocks from a transformer, the most brutal method for speeding up a model, as an Ising spin glass problem. The starting intuition: the effect of removing one block depends on the other blocks removed, so scoring them in isolation, as classical methods do, is a reasoning error. The system's energy then serves as a fast proxy for predicting the compressed model's performance without rerunning benchmarks. Announced result: **nearly 23 MMLU points** gained against the best competing method on a Llama-3.3-70B-Instruct compressed by half.

**Valued at $1.4 billion before publishing a single model**

Naive AI, founded in February in Beijing by Jifeng Dai, a professor at Tsinghua University, is said to have raised about $400 million across three rounds with **fewer than 100 employees** and no model published to date. Its strategy stands out from that of its competitors: not pre-training from scratch, but starting from an existing Chinese open-weight model to concentrate its resources on midtraining, post-training, reinforcement learning, and architecture. An open-weights model could be released as early as this month.

**DeepSeek-V4.1-Flash compresses its memory to 890 bytes per token**

A multimodal MoE architecture with **552 billion parameters**, of which only 16 are activated during decoding and 8 during prefill, a one-million-token window, and pre-training on 45 trillion multimodal tokens. The main contribution lies elsewhere: the very aggressive compression of the attention cache, reduced to **890 bytes per token**, which brings down the memory cost and bandwidth required for agents working on very long documents.

**MiniMax-H3 stumbles on physical reasoning**

Tested on 517 cases where text, image, video, and audio each provide only part of the clues about a physical event, the model achieves only **41.97% success**. It does best on decision-making from video (56%) and collapses on audio disambiguation (27.40%). Translation: multimodal models know how to read each channel separately, they still do not know how to piece the parts back together well.

# **🗞️MORE NEWS**

### **The UN warns that there is no guarantee of human control over AI agents**

