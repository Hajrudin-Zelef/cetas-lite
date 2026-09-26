---
id: collect-240926-vision-ia/vision-ia/nvidia-egale-le-modele-ouvert-d-openai-avec-4-fois-moins-de-parametres-1
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "California", "China", "DeepSeek", "Google", "Meta", "Microsoft", "Moonshot", "OpenAI", "xAI"]
dates: []
keywords: ["research", "acquisition", "agent", "agents", "apache", "astra", "benchmarks", "claude", "deepseek", "distillation", "gemini", "gpt-5.6"]
source: docs/RAG/clean_en/vision-ia/nvidia-egale-le-modele-ouvert-d-openai-avec-4-fois-moins-de-parametres.md
source_anchor: ""
source_lines: [1, 54]
sha256: 4ab1312093f292389d14c1bdb46bf2e8fb7958ade1667c377cd5cf43e55a7409
---

# 🧠 **RESEARCH**

<!-- source: https://vision-ia.beehiiv.com/p/nvidia-e-gale-le-mode-le-ouvert-d-openai-avec-4-fois-moins-de-parame-tres -->

Sundar Pichai confirmed this morning that **Koray Kavukcuoglu** is taking the helm of Google DeepMind with the title of Senior Vice President, reporting directly to the CEO of Alphabet. Demis Hassabis, co-founder and historic boss, is stepping down from operational leadership to become president of DeepMind and Chief Scientist of Alphabet. Behind the polished press release, it's a lab under tension that is changing hands.

Kavukcuoglu is no parachuted-in outsider: he arrived in 2012, two years before the Google acquisition, a computer science PhD from NYU under the supervision of Yann LeCun, he contributed to DQN, AlphaGo and WaveNet. He now oversees the Gemini models, frontier research, the Gemini app and developer platforms.

Google hasn't released a frontier model since **Gemini 3.1 Pro in February**, while Anthropic (Mythos) and OpenAI (GPT-5.6) have taken the lead
**Gemini 3.5 Pro missed three release dates**: June, mid-July, then August, and Gemini 3.6 Flash is judged behind OpenAI, Anthropic, xAI, Meta and several Chinese labs on intelligence benchmarks
The talent drain is spectacular: Jeff Dean is leaving after **27 years** at Google to start his own company, Noam Shazeer (co-lead of Gemini) moved to OpenAI in June, and Nobel laureate John Jumper left for Anthropic along with Jonas Adler and Alexander Pritzel, three of the architects of AlphaFold
Internal climate degraded according to Fortune: 60-hour weeks, low morale, and a union movement launched in May
An internal team dubbed **Code Strike** was assembled to catch up on code, a field where Morningstar judges Anthropic and OpenAI to be "miles ahead"

Audience isn't the problem: the Gemini app just crossed one billion monthly users and the Gemma models have accumulated more than 900 million downloads. The problem is the technological frontier, and analysts read this appointment as a refocusing: fewer academic projects and world models dear to Hassabis, more LLMs and developer tools. For you, nothing to test today, but the next Gemini will tell whether the pivot worked.

Alibaba has released **Qwen-MM-Plugins**, an open source repository under Apache-2.0 license that adds multimodal capabilities to the code agents you already use. The idea is clever: rather than forcing you to adopt Qwen Code, Qwen turns its vision and audio into a **tool layer reusable by any agent**.

Once installed, your agent gains functions it discovers and chains on its own: reading an image or a video, doing OCR, locating an object in a photo, segmenting, transcribing audio, cropping. You point to a file with `@` and ask in plain English for whatever you want.

**8 separately installable modules**: core (images, videos, documents, 3D files), api (vision, OCR, grounding, transcription, segmentation via the Qwen VL and Omni models), search (web search, page extraction, reverse image search), video-memory (hierarchical memory for long videos), video-edit (image, video, audio generation and editing), blender (3D modeling and rendering), freecad (parametric CAD, STEP, STL, FEM formats) and edu-agent
**Six compatible agents**: Claude Code, Codex, Qoder, OpenClaw, Qwen Code and Gemini CLI, with a shared configuration file
A single script does everything: `curl -fsSL https://raw.githubusercontent.com/QwenLM/Qwen-MM-Plugins/main/install.sh | bash` installs, configures, verifies and uninstalls
**The core module requires no API key.** The others require a DashScope key, and Serper for web search
Free to install, you only pay for the third-party API calls you trigger

This is the kind of building block that changes the nature of a tool without changing the tool. A code agent that could only read text can now open a scanned invoice, describe a video, output a Blender render or manipulate a CAD plan. And while OpenAI and Anthropic sell closed multimodal models, Qwen is freely distributing what it takes to make their own agents see.

Since August 11, **Meta Glasses are banned in all courts in England and Wales**. His Majesty's Courts & Tribunals Service, which manages all criminal, civil and family jurisdictions, collects them at the entrance and returns them at the exit. A delicious detail: smartphones, however, remain tolerated, as long as they don't film.

The reason is obvious once stated. Filming without authorization in a court is contempt of court, punishable by prosecution. But you don't discreetly stick a phone on your face, whereas a pair of glasses with an integrated camera allows you to record an entire hearing without anyone noticing.

**7 million pairs** sold in 2025, with a 12-megapixel camera, video up to 3K at 30 frames per second, starting at **299 dollars**
The State of New York beat London to it: ban on connected glasses (Meta, Xreal, RayNeo) in its **1,240 courts since July 20**, to prevent secret identification of jurors and witnesses
The flaw that worries institutions: the LED meant to signal recording can be neutralized by a 60-dollar mod, documented by 404 Media
The trigger dates back to February: during a trial in California, a judge had to threaten Mark Zuckerberg's entourage with contempt of court to get him to remove his glasses
A London court had already dealt with connected glasses, but for a completely different use: a man is accused of having had his answers fed to him during cross-examination via the built-in speakers

The rejection goes far beyond the courts: the Wetherspoons pub chain, restaurants, theaters, concert halls and British conventions are also banning them, after a wave of videos of people filmed without their consent and sometimes identified online. Meta declined to comment. When a product gets confiscated at the entrance to the courts of an entire country, it's no longer an image problem, it's a design problem.

# 🧠 **RESEARCH**

**AI is no longer artificial: lab-grown brains to replace LLMs**

Biologists are manufacturing human brain organoids from stem cells: clusters of a few million living neurons that emit real brain waves after eight months of culture. At UC San Diego, these clusters guide robots through mazes; at Johns Hopkins, they serve as the basis for biocomputing; a Melbourne startup has them playing Pong and Doom. The bet: a living intelligence rather than an artificial one, as a head-on alternative to large language models.

**AMIE, Google's medical AI, conducts real-time video consultations**

Google Research and Google DeepMind are moving AMIE from text chat to video calls. Built on Gemini and Project Astra with a multi-agent architecture, the system interprets visual and audio signals, guides a remote physical examination and reasons about the diagnosis in real time. In a randomized study with patient actors and general practitioners, AMIE scores well on history-taking, diagnostic accuracy and communication quality, and the simulated patients preferred video to text. It remains a research system, not clinically deployed.

**A new technique reveals the hidden thoughts of AI models**

Researchers from the University of Tübingen, Max Planck and MATS have found how to extract, via a simple API, the internal reasoning that frontier models are supposed to keep to themselves. "Every major provider we tested shares this vulnerability," summarizes Alexander Panfilov: the flaw made it possible to recover passwords and API keys slipped into the reasoning, and it has since been fixed. An explosive side effect: the reasoning of Kimi K3, from Moonshot AI, looks strangely like that of Claude Opus 4.8 and GPT-5.6, unlike DeepSeek. A hint of distillation, not proof.

**CARE-X, Microsoft's model that reads chest X-rays**

