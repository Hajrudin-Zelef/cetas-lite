---
id: collect-240926-vision-ia/vision-ia/ox-alpha-tombe-le-masque-le-modele-chinois-qui-code-a-prix-casse-1
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Apple", "China", "Google", "Hugging Face", "Nvidia", "OpenAI", "OpenRouter", "Z.ai"]
dates: ["2026-08-27"]
keywords: ["research", "agent", "agents", "agi", "chatgpt", "claude", "context window", "cost", "embedding", "gemini", "glm", "gpt-5.6"]
source: docs/RAG/clean_en/vision-ia/ox-alpha-tombe-le-masque-le-modele-chinois-qui-code-a-prix-casse.md
source_anchor: ""
source_lines: [1, 110]
sha256: 845630f7c33d08169464c44682c2303d1aff24f1870754556b1eb6185872cba9
---

# 🧠 **RESEARCH**

<!-- source: https://vision-ia.beehiiv.com/p/ox-alpha-tombe-le-masque-le-mode-le-chinois-qui-code-a-prix-casse -->

## **Today:**

🕵️ Ox Alpha reveals its identity and its 320 billion parameters

🎙️ Gemini 3.5 Transcribe cleans up and structures speech in real time

⚡ Qwen3.8-Flash-Next drastically reduces the computation required

🤗 Nvidia is about to acquire Hugging Face

🎬 LAION opens 80 million videos to research

🔓 OpenAI explains how an agent escaped its environment

🧠 Sam Altman predicts his AGI before the end of 2026

🤖 Jetson Orin Nano 2 runs AI directly inside robots

☕ A JD.com café runs with a single robot

🧩 AI assistants get lost in their own brands

📱 Apple sets its iPhone and Siri AI date for September 9

🖥️ The Mac mini M5 Pro remains more powerful than the new M6

🩸 The FDA authorizes the first autonomous blood-drawing robot

🏃 Chinese humanoids break new records

A real estate agent website. A village association website. Software to learn Spanish. Software to manage your finances. All written in French, without a single line of code. It's the only AI skill you can say that about.

In this new update I teach you Codex in detail, which is included in your ChatGPT subscription. If you don't have ChatGPT, I show you GLM 5.2, which is the equivalent but free. Otherwise, if you have Claude, it also works on Claude Code. In short, nothing extra to buy.

And this module is only one part of the training.

✅ Complete training on AI (LLMs, AI Marketing, image, sound and voice generation, automation with n8n, first steps with agents, etc.)

✅ Complete lesson to create your own AI agents (n8n) and automate without limits

✅ Complete lesson on Claude Code, the best AI tool in 2026

✅ Access to a network of qualified professionals

✅ Regular updates to stay at the cutting edge of AI

➡️ If you join now, you lock in the €49 rate for life (one-time payment).

No matter how high the price climbs, and it will soon reach €100 or more, given the demand, you won't pay a single cent more.

One payment. Lifetime access.

On that last point, the proof is right before your eyes: the vibe coding module dated from late 2025, I deleted it and redid it from scratch, and subscribers didn't pay a single cent more. That's what "updates included" means. The price will go up, but never for those who are already in.

More than 13,000 people are already training with it. Their reviews are on the page.

See you inside.

The mysterious Ox Alpha is in reality **GLM-5.3-Flash**, a model from Z.ai designed to program, manipulate tools, and conduct automations during long sessions. It has **320 billion parameters**, but activates only **18 billion** at each step, with text, image, and video inputs and a context exceeding **one million tokens**.

The model can analyze a software repository, modify code, work in a terminal, interpret documents or graphics, and chain tool calls. Its weights are available under the MIT license, while its API displays a particularly aggressive price.

### **What to remember:**

The Mixture of Experts architecture has **320 billion parameters, of which 18 billion are active**, spread across 45 layers and trained with **30 trillion tokens**.
The context window reaches **1,048,576 tokens**, with a maximum output of **131,072 tokens**, enough to process large repositories or long chains of actions.
Z.ai publishes a score of **84.3 on Terminal Bench 2.1**, versus 85.0 for Claude Opus 4.8 and 87.4 for GPT-5.6 Terra. On Toolathlon Verified, GLM-5.3-Flash reaches **78.4**, ahead of Claude Opus 4.8 at 76.2.
OpenRouter charges **$0.075 per million input tokens** and **$0.25 per million output tokens**.

### **Why it matters**

The API price makes the model immediately testable for code agents or large-scale automations. Local execution, however, remains reserved for very well-equipped servers, because publishing the weights of a 320-billion-parameter model does not turn it into a home model. Z.ai above all brings new pricing pressure on OpenAI and Anthropic, with performance close enough to their best models to make the price gap hard to ignore.

Google DeepMind launches **Gemini 3.5 Transcribe**, a proprietary speech-to-text model capable of transcribing more than **85 languages**, even when a user switches languages mid-sentence. It removes hesitations, understands spoken self-corrections, adds punctuation, and can recognize up to **1,000 custom phrases**.

The service exists in two versions. `gemini-3.5-transcribe-live` processes an audio stream via WebSocket with latency under one second, while `gemini-3.5-transcribe` analyzes files and adds speaker identification as well as word-by-word timestamps.

### **In detail:**

The model is based on Gemini 3 Pro, accepts audio and text inputs, and has a context of **96,000 tokens**, with a maximum output of **32,000 tokens**.
Google announces an average error rate of **4.0% live** and **2.6% on recordings**, according to Artificial Analysis.
The time needed to obtain a final transcription would be reduced by **70% compared to Chirp 3**. On FLEURS, the announced error rates are 5.50% in streaming and 5.04% outside streaming.
Files can last up to **one hour**, or 30 minutes when speaker identification and detailed timestamps are enabled.
The preview is available in Gemini API and Google AI Studio, with a free tier then approximately **$0.009 per minute live** and **$0.005 per minute on file**.

### **What it changes**

A developer can now create a subtitling tool, professional dictation, meeting minutes, or a voice agent without managing an audio model themselves. Custom vocabulary will be particularly useful for product names, medical or legal terms, and alphanumeric references. However, the weights are not downloadable, and Google still documents risks of hallucination, expiration, and slowdown.


Alibaba presents **Qwen3.8-Flash-Next**, an open-weight multimodal model that foreshadows the architecture of Qwen4. Its main network contains **125 billion parameters**, but only **6 billion** work on each token, which greatly reduces the computation needed for reasoning, code, and agents.

The model receives text, images, and videos, then generates text over a native context of **262,144 tokens**, expandable to one million. It can explore a software repository, use tools, analyze a video, or execute long office tasks.

### **Key points:**

The architecture includes **512 experts**, of which 10 are dynamically selected and one remains shared. It adds 51 billion N-gram embedding parameters and 4 billion dedicated to multi-token prediction.
Alibaba claims a training cost equivalent to about **one-ninth that of Qwen3.7-Plus**.
The published results reach **62.5 on SWE-bench Pro**, 58.7 on DeepSWE, 73.9 on CoWorkBench, and **91.7 on GPQA Diamond**. No consolidated independent evaluation has yet been found.
The QwenCloud API costs **$0.15 per million input tokens** and **$0.47 per million output tokens**, with function calling, web search, and a code interpreter.

### **The impact to remember**

Qwen3.8-Flash-Next seeks less to beat all records than to provide many capabilities with few active parameters. For the reader, the API is immediately accessible and the weights allow private deployment, provided one has a sufficiently powerful server. This architecture above all shows how Qwen4 could reduce the cost of agents capable of working for a long time on code, documents, and multimodal data.


Nvidia reportedly agreed to acquire Hugging Face for **$12.9 billion**, according to information initially published by The Information. As of **August 27, 2026**, Nvidia and Hugging Face have still made no official announcement and have not responded to Reuters requests. The timeline, regulatory approvals, and guarantees of independence remain unknown.

