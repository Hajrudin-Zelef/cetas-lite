---
id: collect-240926-vision-ia/vision-ia/nvidia-egale-le-modele-ouvert-d-openai-avec-4-fois-moins-de-parametres
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["AWS", "Alibaba", "Anthropic", "Apple", "California", "China", "DeepSeek", "Google", "Meta", "Microsoft", "Moonshot", "Nvidia", "OpenAI", "Samsung", "SpaceX", "United States", "xAI"]
dates: ["2025-12", "2025-12-29"]
keywords: ["research", "acquisition", "agent", "agentic", "agents", "apache", "astra", "aws", "bedrock", "benchmark", "benchmarks", "chatgpt"]
source: docs/RAG/clean_en/vision-ia/nvidia-egale-le-modele-ouvert-d-openai-avec-4-fois-moins-de-parametres.md
source_anchor: ""
source_lines: [1, 149]
sha256: ec4e225841f30571cb728393944c0a29c365fbefcd36239f3411e620c8672dba
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

Microsoft Research presents a unified vision-language model that interprets chest radiographs by combining free-text report generation with verifiable structured predictions. Built on Qwen3-VL-4B-Instruct, it is fine-tuned with reinforcement learning (DAPO) to reward clinical accuracy rather than text fluency, and can call measurement tools to calculate a cardiothoracic ratio. Validated on real data from Narayana Health hospital in India, including rare intensive care pathologies. A research model, without regulatory approval.

**Can a local LLM run my AI assistant?**

The experiment many have in mind but never carry out: the author replays **27 real production tasks** across two local models, separated by one generation of hardware, to find out what it really takes to replace Claude as the brain of a personal agent that operates **90 tools**. The verdict is nuanced and detailed, and above all it's an excellent inventory of what breaks when you bring your assistant home.

**AI takes over mathematics**

OpenAI claims to have produced the solutions to **ten mathematical problems unsolved for decades**. James Maynard, a professor at Oxford and Fields Medalist, tells The Verge about a year of questioning in the face of a traditionally slow discipline that suddenly starts running. The parallel with AlphaFold in biology is obvious, with the same question at the end: what remains of the profession when the machine finds the proof?

**Agentic memory: ALTK-Evolve matches ACE with fewer tokens**

IBM Research compares its ALTK-Evolve system with ACE (Agentic Context Engineering). Both get an agent to learn from its own past failures without touching the model's weights: when an agent fails a multi-step task, it's almost never for lack of knowledge, it's because it mishandles an API's pagination or resolves the wrong entity. Both refuse to compress lessons into generic summaries and instead count occurrences. The real difference plays out in the consolidation and delivery of that memory, and therefore in the token bill.

**Anthropic's giant IPO runs into investor skepticism**

According to the Wall Street Journal, Anthropic is preparing a stock market listing for September or October, potentially the largest ever, on a valuation of 965 billion dollars. In meetings, investors are asking tough questions: Chinese competition, tensions with the Trump administration, local protests against data centers. The price that emerges from this IPO will serve as a benchmark for valuing the entire industry.

**Metis: a foundation model with built-in persistent memory**

What if memory were a capability of the model itself, rather than a retrieval system cobbled around it? Metis proposes a persistent memory state housed in the backbone, updated during ordinary forward passes, while the learned weights remain frozen. Instead of storing exchanges as text to be retrieved later, the model compresses past interactions directly into this state. Enough to remember a conversation from three weeks ago without reinjecting it into the prompt.

# **🗞️MORE NEWS**

**ChatGPT finally arrives on Linux**

OpenAI releases a preview of its desktop app for Linux, the platform most requested by its community for months. The package includes ChatGPT, ChatGPT Work, and Codex, and covers Ubuntu 24.04 and 26.04 LTS, Debian 13, Fedora 43 and 44. With this launch, the app is available on all major desktop systems. Anthropic had made its move a month earlier with Claude for Linux.

**A company promising "100% human, never AI" runs entirely on AI**

Research Gold sells systematic reviews and medical meta-analyses ready for peer review, with a killer pitch: "100% written by humans, never by AI." A 404 Media investigation: the doctor methodologists listed on the site don't exist, profiles and photos generated by AI. Others, very real, discover that their identity is being used without their consent. And when the journalist calls, it's an AI agent that picks up, refusing to admit that it is one. The topic is generating a lot of reaction in the tech community.

**You're wrong about how to think about online trends**

Cyber-ethnographer Ruby Thelot explains to WIRED why virality has ceased to be a reliable indicator: as soon as it becomes a target to optimize, it ceases to be a measure, Goodhart's law applied to networks. The internet has balkanized into watertight digital islands, which completely distorts the perception of what actually "works." He also debunks the myth of dating app burnout, and explains why they're rushing into AI to revive engagement.

**Keet generates video courses on any subject**

A mobile app from Y Combinator that builds a complete learning path on the theme of your choice: short explanatory videos, game-based exercises, progression spread over time like a real curriculum. The founders draw on Biglan's academic categories (Hard-Pure, Hard-Applied, Soft-Pure, Soft-Applied) to adapt the format to the type of knowledge being taught. The idea: to give self-learners back the structure that school provided for free.

**Abbott connects its glucose sensor to Google Health**

Abbott's Lingo continuous glucose sensor, sold without a prescription and intended for non-diabetic adults, will send its data to the Google Health app. Google Health Coach will derive AI-generated nutritional recommendations from it, cross-referencing blood sugar, activity, and sleep. The two companies are also launching a vast study on metabolic health. Important caveat: several doctors continue to doubt the real usefulness of permanent glucose monitoring in a non-diabetic person.

**OpenAI begins testing advertising in ChatGPT**

It was announced, now it's launched: OpenAI is testing the display of ads in ChatGPT to fund free access. The company promises clearly labeled advertisements, visually separated from responses, and above all AI responses that remain independent of advertisers. Privacy protections and user settings are announced in parallel. The real test will come with use: the boundary between a response and a sponsored recommendation is thinner in a chat than on a results page.

**ChatGPT and Gemini both cross the billion-user mark**

Sundar Pichai announced on X that Gemini has surpassed one billion monthly users, making it the fastest-growing product in Google's history, and the fourteenth to reach this milestone. OpenAI had confirmed the same threshold for ChatGPT on August 6, in a post that went almost unnoticed, while external data already suggested it as early as June. Two assistants, two billion users combined, in three and a half years.

**Why Amazon confirmation emails have become so vague**

Since this summer, Amazon order confirmations no longer name the products purchased: instead, a generic category and a clipart. "Your Beauty item is confirmed," reads an email for dental appliance cleaning tablets. To find out what's in your order, you have to leave your inbox and go back to Amazon. The timing coincides with the rise of AI agents capable of reading your emails, Gemini in Gmail chief among them: Amazon would rather blind Google than share its purchase data.

**Authors of the 2022 Google study under fire**

In spring 2022, Google quietly recruited thirteen renowned science fiction authors, including Ken Liu and Robin Sloan, to test Wordcraft, a text editor built on LaMDA, the ancestor of Gemini. Everything was under a confidentiality agreement, and the stories written with the tool were published alongside a white paper. Four years later, the rediscovery of the experiment earns the participants a barrage of criticism from part of the literary world, with one author accusing them of having legitimized "the plagiarism machine."

**Manus becomes independent again, Beijing having forced Meta to unwind its acquisition**

Meta had announced in December 2025 the acquisition of Manus, a general-purpose AI agent startup founded in China in 2022, for $2 billion. Chinese regulators (NDRC) demanded in April the outright cancellation of the deal. Manus announces it will "soon" resume operations as an independent company, and warns its users that they will need to back up data generated since December 29, 2025. A clear demonstration of who really controls Chinese AI assets.

**Galaxy Buds are set to transform into hearing aids**

The FDA has authorized Samsung's hearing aid feature on the Galaxy Buds 3 Pro and 4 Pro, expected by the end of the year in the United States and other markets. A five-minute hearing test built into the Samsung Health app provides a medical-grade result and classifies hearing loss into four stages. Amplification relies on noise reduction and beamforming to isolate the voice of the person speaking directly in front. Samsung is following in the footsteps of Apple's AirPods Pro.

**Novo Nordisk entrusts its drug discovery to AWS AI agents**

AWS becomes the preferred cloud provider and strategic AI partner of the Danish laboratory, with agents deployed on identifying therapeutic targets and designing treatments. A co-innovation hub is opening in London, where AWS engineers and Novo Nordisk researchers will work on the same data. More than 25,000 employees already use a generative AI platform on Amazon Bedrock covering 2,500 use cases, and a system built on Claude 3.5 has reduced the production time of clinical trial documentation by more than 90%.

**Saber denies replacing its writers with ChatGPT**

Matthew Karch, CEO of Saber Interactive, states that "neither Saber nor Unigine replaced any writer with AI" on the game Rideshare Stimulator. The former lead writer, Stella Sacco, tells the opposite story on Bluesky: ousted during development, with passenger voices generated by AI and not disclosed on the game's Steam page. Two versions, no public evidence to settle the matter, and an industry where this kind of dispute is bound to multiply.

**Spotify will label AI artists and remove them from recommendations**

Starting in mid-September, an "AI Persona" badge will appear on profiles, in search, and on tracks by artists whose identity is AI-generated. Spotify will not settle for self-declaration: the platform will prioritize reviewing photorealistic profiles exceeding certain audience thresholds, with the possibility of contesting the label. These artists will be excluded by default from editorial and algorithmic recommendations, unless you explicitly follow them. An important nuance: the measure targets the profile's identity, not the way the music was produced.

**Bernie Sanders demands OpenAI, Anthropic, and Meta suspend AI development**

The U.S. senator addresses an open letter to Sam Altman, Dario Amodei, and Mark Zuckerberg, in which he states that AI capabilities have crossed a critical threshold and demands an immediate pause. The tone leaves no room for doubt: "Let me be very clear: if you don't take the necessary steps now, my colleagues and I in the Senate will." None of the three laboratories has announced any intention to slow down.

**OpenAI's head of ethics leaves after less than a year**

Chloe Bakalar is leaving her position as head of ethics at OpenAI less than twelve months after her arrival, reports the Financial Times. Such a short stint in such an exposed role mechanically raises the question of the real weight given to ethical and safety issues within the organization. The topic is generating a lot of reaction in the tech community.

**Brad Lightcap, OpenAI's longtime chief operating officer, is leaving**

Having joined in 2018, four years as chief financial officer and then chief operating officer from 2022, Brad Lightcap announces in an internal message that he is leaving OpenAI to "start something new," without further details. He joins a growing list: Fidji Simo, Bill Peebles, Kevin Weil. A profound reshuffling of leadership, at the very moment the company is preparing a possible initial public offering.

**The Norwegian sovereign wealth fund discloses its stake in SpaceX**

The world's largest sovereign wealth fund, with $2.3 trillion under management, reports a record half-year profit of $184.9 billion, driven by the rally in Asian tech stocks and a 9.4% return. It reveals for the first time that it holds 0.05% of SpaceX, or about $1.2 billion. Its tech portfolio also includes 1.3% of Nvidia ($61.8 billion), 1.2% of Apple ($52.7 billion), and 1% of Tesla.
