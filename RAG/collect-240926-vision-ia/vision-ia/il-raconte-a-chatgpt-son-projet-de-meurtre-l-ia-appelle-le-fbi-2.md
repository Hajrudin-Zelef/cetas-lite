---
id: collect-240926-vision-ia/vision-ia/il-raconte-a-chatgpt-son-projet-de-meurtre-l-ia-appelle-le-fbi-2
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["Anthropic", "California", "China", "Meta", "OpenAI", "SpaceX", "United States"]
dates: []
keywords: ["research", "acquisition", "agent", "agents", "attribution", "claude", "compute", "copyright", "cost", "diffusion", "gpu", "gpus"]
source: docs/RAG/clean_en/vision-ia/il-raconte-a-chatgpt-son-projet-de-meurtre-l-ia-appelle-le-fbi.md
source_anchor: ""
source_lines: [108, 165]
sha256: d0f39434e94eb46d676f712d07153c223553480bde7689b14aa06413a44fda92
---

# 🧠 **RESEARCH**

An open dispute erupted on X on August 18 and 19 between **Dario Amodei**, head of Anthropic, and three voices from the opposing camp: investor **Gavin Baker**, former White House advisor **David Sacks** and **Yann LeCun**, researcher at Meta. The starting point: Amodei writes that "AI is structurally a technology that tends to concentrate power, for reasons that have nothing to do with regulation."

**What was said:**

- Amodei on open source: open-weight models "help a little, but are far from being a sufficient solution, because they simply shift the concentration toward those who own the most compute power and chips". In other words, downloading weights still doesn't give you the GPUs to run them at scale

- Gavin Baker claims that Amodei would have said internally that Anthropic could end up as the last private company standing against states

- David Sacks speaks of "hubris" and **regulatory capture**: Anthropic would push rules that only Anthropic can comply with, eliminating startups and open source. He compares the proposed federal agency for model approval to "a DMV for AI", after the American driver's license administration, the slow counter par excellence (Fortune)

- Yann LeCun counters with the printing press analogy: AI amplifies human intelligence by broadening access to knowledge, and a diversity of systems is better than a few licensed models

- Amodei's response: the choice between concentrated regulation and total diffusion is a "false choice". He cites California law **SB53**, which exempts companies below a certain revenue and training cost threshold (The New Stack)

Amodei's argument has a very concrete consequence for anyone running a model at home: if open weights require a GPU farm to be truly useful, then open source doesn't rebalance anything, it redistributes power to those who already have the silicon. If he is wrong, the rules he defends risk above all closing the door behind him. The two camps are not fighting over philosophy, but over the text of the laws being prepared.

# 🧠 **RESEARCH**

**Robotic cabinets grow human tissue and test drugs in place of patients**

Vivodyne, a San Francisco startup, operates **12 robotic laboratories the size of a cabinet** that grow human tissue and test AI-designed drugs on it, at a rate of **3 million experiments per year**, twice as many as all American clinical trials combined. The loop is closed: the AI designs the experiment, the robots execute it, the results feed into designing the next cycle. Several major pharmaceutical laboratories already use it to rule out dangerous molecules before human testing. The FDA lifted the requirement for animal testing in 2022, which paved the way for this type of pipeline.

**Remove the Mona Lisa from the training set, the model paints it anyway**

Researchers at MIT have published a study in Nature Communications on what they call **attribution decay**: the more a diffusion model grows, the harder it becomes to link a generated image to a specific training data point. The most telling test: they removed the Mona Lisa and the entirety of Leonardo da Vinci's work from the corpus, and the model continued to reproduce them. The authors conclude that attribution becomes simply "impossible" beyond a certain scale. Immediate legal consequence: model size turns into a de facto shield against copyright infringement lawsuits.

**Do we really need a doctor behind AI? A JAMA op-ed says no**

Published in the medical journal JAMA, this op-ed argues that autonomous AI will soon surpass any doctor-plus-AI pairing on clinical reasoning tasks. The authors therefore advise against enshrining in law the requirement of a final human decision, on the grounds that the constraint would degrade the outcome for the patient. They do concede one major point: almost all the evidence comes from simulations, not from actual patient care.

**84% of Chinese enthusiastic about AI, versus 38% of Americans**

Stanford's AI Index measures the biggest opinion gap in the tech world: **84% of Chinese** say they are enthusiastic about AI, the highest rate among the countries surveyed, versus **38% in the United States**. Chinese optimism rests on the conviction that the gains will reach ordinary citizens and that the state will rein in big companies. American skepticism, meanwhile, is dominated by the memory of social media: disinformation, concentration of platform power, and forecasts of job destruction.

**X's algorithm does favor content that makes people angry, and not equally for everyone**

A study published in PNAS, led in particular by Ziv Epstein (Stanford) on **715 American users** equipped with a tracking extension, confirms what everyone suspected: X's recommendation algorithm amplifies posts that provoke anger. The mechanism is identified: replies account for less than 7% of total engagement but are valued far more than likes, which creates a loop favorable to provocation. Users who identify as Democrats are more exposed to it, though researchers cannot yet explain why.

**Anthropic describes "mind viruses": ideas that spread on their own between AI agents**

In a study conducted with a Swiss university, Anthropic shows that AI agents can convince each other to adopt an undesirable goal, then pass it along from one to the next through simple natural-language messages. It is the equivalent of a computer worm, except that the agents themselves do the copying. The persistence mechanism is the most worrying part: the infected agent is pushed to rewrite files that are reloaded in subsequent sessions, which makes the contamination lasting. A security risk specific to multi-agent architectures, which are becoming widespread in enterprises.

**AI models that grade other AI models change their minds in 62 to 91% of cases if you push**

Meta tested the robustness of **9 state-of-the-art models** used as judges to evaluate other models, a practice that has become standard in the industry. Result: an adversarial model manages to flip an initially correct verdict in **62 to 91% of cases** through sustained and adaptive argumentation. The dangerous failure mode is therefore not the judge that gets it wrong from the start, but the reliable judge that can be convinced to get it wrong.

**A 27-billion-parameter research agent beats Claude Opus 4.8 and GPT-5.5**

On replicating scientific papers never seen during training, an agent with only **27 billion parameters** outperforms Claude Opus 4.8 and GPT-5.5. Its recipe: it doesn't code, it orchestrates. The model is trained to reason like a researcher, to decide what to test and in what order, then it subcontracts the implementation to models stronger in code. The authors propose with Replica a large-scale task space to train this behavior. The implication is major: a single giant model may not be necessary to do science.

# **🗞️MORE NEWS**

**A Cursor engineer offers $200 in credits to anyone who cancels their Claude subscription**

On August 17, a Cursor engineer offers on X, strictly in a personal capacity and without his management's approval, **200 dollars in Ultra credits** to anyone who posts a screenshot of canceling their Claude subscription. The post goes viral within hours, developers confirm having received the credits, and a bot fails to clear the queue in 16 hours. The operation lands three days after SpaceX's acquisition of Cursor and right after a new Anthropic outage, the 166th of the year according to the tally relayed by AI Secret. Not a marketing campaign, just an employee with a credit card and formidable timing.

**Claude Code keeps +50% weekly limit until August 31**

