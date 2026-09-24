---
id: collect-240926-korben/korben/odysseus-l-ia-auto-habergae-de-pewdiepie-enfin-rapide-sur-mac-korben
title: "Odysseus - PewDiePie's self-hosted AI finally fast on Mac"
domain: korben
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Apple", "OpenAI", "OpenRouter"]
dates: []
keywords: ["agent", "agents", "claude", "gpu", "license", "llama", "llama.cpp", "parameters", "qwen", "research"]
source: docs/RAG/clean_en/korben/odysseus-l-ia-auto-habergae-de-pewdiepie-enfin-rapide-sur-mac-korben.md
source_anchor: ""
source_lines: [1, 92]
sha256: 1428e451d0d6ed19e88f4b3494ac9fa882b2499114b14c7c5e696c278fa9791e
---

# Odysseus - PewDiePie's self-hosted AI finally fast on Mac

<!-- source: https://korben.info/odysseus-ia-locale-mac.html -->

# Odysseus - PewDiePie's self-hosted AI finally fast on Mac

## Key takeaways AI-generated summary

1. Odysseus, PewDiePie's self-hosted AI workspace, brings together chat, agents, web search, email, calendar, notes, documents, and image generation locally, with a Compare mode to test two models side by side blindly.
2. On Mac, Docker cannot access the Metal GPU and models run on CPU, so you need to install Odysseus natively via the start-macos.sh script and connect Ollama to run models like Qwen3-30B-A3B without slowdowns.
3. Chat and the Markdown editor with history work well, but the Cookbook, agents, calendar, and factual search remain too fragile: the Deep Research feature invented a video, a brand name, and statistics with no basis.

**Odysseus** is the AI workspace that **PewDiePie** dropped on GitHub at the end of May. Chat, agents, web search, email, calendar, notes, and documents, all brought together in a gorgeous self-hosted cockpit where models and history can stay on YOUR machine. You also get a Markdown editor with version history, built-in image generation, and a Compare mode to pit two models side by side blindly (excluding third-party services tied to web search, email, and calendar, obviously...).

I wanted to install it on my Mac, but I still ran into a little trap that almost turned the experience into a failure... So here's how to run it properly on Mac + a few ideas of what you'll be able to do with this thing.

## Before you start

Odysseus moves fast. The `main` branch is stable and tested, but the `dev` branch can break at any moment. Favor the stable branch by cloning it explicitly:

```
git clone https://github.com/odysseus-dev/odysseus.git
cd odysseus
git checkout main
```
## Step 1: Forget Docker on Mac

On Apple Silicon, Docker cannot access the Metal GPU. Models run on CPU, and it lags like crazy. So on Mac, you're forced to install it natively.

## Step 2: Install Odysseus in 2 commands

Open a terminal and run this:

```
git clone https://github.com/odysseus-dev/odysseus.git
cd odysseus
./start-macos.sh
```
The script then installs the dependencies and starts the server. On my Mac, AirPlay was squatting on port 7000, so the interface opened automatically at `http://127.0.0.1:7860`. Then, on first launch, you need to set a new password for the admin account.

If you're accessing Odysseus from another machine (VPS or remote server), open an SSH tunnel rather than exposing the port directly to the internet:

```
ssh -L 7860:127.0.0.1:7860 [email protected]
```
Then just access `http://127.0.0.1:7860` locally. It's safer than opening the port to the public.

## Step 3: A fast model, like Qwen

For something quick without becoming stupid, I chose **Qwen3-30B-A3B** for my tests. This model only activates part of its parameters with each response, which fits well with the unified RAM of a big Mac.

The first pretty cool feature in this tool is the **Cookbook**, which lets you download models suited to your config. It correctly recognized my M4 Max and its 128 GB but then crashed like crap on llama.cpp with this invalid `--flash-attn auto` option. After removing the parameter, the server stayed missing in action while the interface still showed it as active. Anyway, I gave up on llama.cpp.

*The Cookbook correctly recognizes the M4 Max and suggests suitable models. The launch, though, crashed.*

The other option that presented itself to me was Ollama (I showed you how to install it here). You grab the model you're interested in with it and start the server open on the local network:

```
ollama pull qwen3:30b-a3b
OLLAMA_HOST=0.0.0.0:11434 ollama serve
```
In Odysseus's settings, then add `http://localhost:11434/v1`. And boom, it works!! Version 1.0.2 now also offers MLX in the Cookbook, but Ollama remains the path that really held up during my tests. You might have better luck than me.

*Manually adding the Ollama endpoint was the most reliable path.*

**Bonus:** If you don't have a GPU or prefer reliability, you can also connect an external API like **OpenAI** or **OpenRouter**. In the settings, just add your API key and Odysseus will switch to those cloud models. It's less free than local, but it works without hassle.

## 3 use cases

So, this tool does a lot of things and you can really use it for whatever you want, but for my part, I've isolated three use cases compatible with my work.

First scenario, the morning news roundup. I give it the links received during the night, it sorts them by topic, spots duplicates, and suggests three short angles. I keep the sources in a local document, with objections and points to verify. And as you saw in my screenshots, the chat already knows how to brainstorm since it gave me three usable angles in a few seconds...

*Three article angles in 29 seconds, with a versioned draft open on the right.*

Second scenario, one folder per article. I store the brief, primary links, quotes, screenshots, and versions of the text there. The Markdown editor grabs the title from the first H1 and keeps the history. With two models configured, Compare mode lets you pit two intros or two outlines side by side without sending the draft to a provider.

Third scenario, the local managing editor. It would monitor a dedicated security alerts inbox, prepare the list of urgent topics, and set reminders in the calendar. On paper, this is exactly the kind of tedious chore that an agent should be able to absorb, but in practice, there's no way I'd let it perform automatic actions like sending emails or managing my calendar. For me, it's not safe enough yet.

I asked the agent to find a document and summarize it. It spent nearly 1,800 tokens guessing how to call `manage_documents`, without ever launching the search tool. On its end, the calendar ended with a `Could not extract JSON`, then the web search found the right GitHub repository, but the model chose a fake one.

The icing on the cake is the Deep Research feature. After 4 minutes of work, it generated 13 pages from 6 retained sources for me. The tool gave the right conclusion to the document, but then completely made up a video, a brand number, and several statistics. So to use it in a journalistic context, without full fact-checking, that won't be possible.

*The report claims the site is empty when it contains a complete guide. A nice hallucination, presented as evidence.*

Well, I'm glad I tested it, but Odysseus still looks like a great cockpit with several buttons not yet connected to the engine. The chat via Ollama and documents are usable, but the Cookbook, the agents, the calendar, and factual search remain too fragile. But be careful, don't think I'm criticizing PewDiePie's work, because the project is still very young, so all these small flaws are normal. Odysseus's potential is enormous, but for now, I'm not ready to entrust it with managing my sources, or drafting an article without losing hours verifying everything.

There you go, the code is on GitHub under the AGPL license. I'll let you discover it!

Thanks to Remouk for sharing, and to finish, I'll leave you with this video from Renaud Dékode who also tested Odysseus:

## Comments

starfix!in Surfshark doesn't make you invMorganein Discord guesses your agests3rv1in The Ray-Ban Display arrive eponponin Openpilot - The NHTSA passesfabienin Claude Code makes you choose
