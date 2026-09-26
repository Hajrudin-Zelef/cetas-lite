---
id: collect-240926-vision-ia/vision-ia/openai-devoile-36-nouveaux-cas-ou-l-ia-a-essaye-de-se-cacher-1
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["Anthropic", "Apple", "China", "Google", "Hugging Face", "JFrog", "Meta", "OpenAI"]
dates: []
keywords: ["research", "agent", "agentic", "agents", "alignment", "chatgpt", "claude", "disclosure", "gpt-5.6", "incident", "jailbreak", "merger"]
source: docs/RAG/clean_en/vision-ia/openai-devoile-36-nouveaux-cas-ou-l-ia-a-essaye-de-se-cacher.md
source_anchor: ""
source_lines: [1, 78]
sha256: e73ff624cc887361c2757d58e80d070197af964ae1d8bd28f5a2fe51e44743d8
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

