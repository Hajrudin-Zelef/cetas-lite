---
id: briefing-ia-2026/09-ai-safety/04-advisory-labs-channels
title: "Distillation advisory, the six labs, channels and legal scope"
domain: ai-safety
role: deep-dive
task: distillation
actors: ["Alibaba", "Anthropic", "CISA", "China", "DeepSeek", "Google", "Moonshot", "StepFun", "United States", "xAI"]
dates: ["2026-07", "2026-09-08"]
keywords: ["advisory", "distillation", "agentic", "attribution", "claude", "consumer", "cybersecurity", "deepseek", "export controls", "gemini", "grok", "kimi"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s09-10"
source_lines: [10121, 10217]
canonical_for: ["distillation-advisory"]
sha256: 861f95b0a0f6817f3b5688048a43392fa1e3db0f22f5c1d4cadd5ae961fd9da9
---

# Distillation advisory, the six labs, channels and legal scope

<a id="s09-10"></a>
### 9.9 The joint advisory AA26-251A (Sept 8): industrial-scale distillation

On September 8, 2026, CISA, the NSA, and the FBI jointly published advisory AA26-251A.
It is a document of a particular nature: signed by three US agencies — civilian cybersecurity, signals intelligence, and the federal police — it accuses six Chinese labs of industrial-scale distillation of American models.
The six names cited are DeepSeek, Moonshot AI, Alibaba, MiniMax, StepFun, and Z.AI.
The models targeted for extraction are Claude, GPT, Gemini, and Grok: in other words, the entire Western leading pack.
The alleged scale is staggering: billions of tokens, millions of queries, activity said to have lasted since at least late 2024.
And the sentence making the most noise: these operations were allegedly conducted "probably with the knowledge of the Chinese government."
First, one must understand what "distillation" means in this context.
In machine learning, distillation consists of training a "student" model to imitate the outputs of a "teacher" model: the teacher is queried massively, its responses are collected, and they are used as training data.
It is a legitimate and long-standing technique, routinely used to compress models or transfer capabilities.
What the advisory describes is not the technique itself, but its clandestine industrialization as a catch-up strategy: instead of investing in fundamental research and training infrastructure, one would systematically siphon the capabilities of competing models via their APIs.
The nuance matters: the advisory does not criminalize distillation as a method; it denounces a massive, organized, concealed extraction campaign, allegedly state-backed.
The scale — billions of tokens, millions of queries, since at least late 2024 — is what moves the matter from commercial dispute to national security affair.
Millions of queries is not a researcher testing a hypothesis; it is a logistical operation: it takes accounts, proxies, budgets, detection-evasion engineering, and response-ingestion infrastructure.
"Since at least late 2024" places the alleged campaign's start before even the wave of announcements that brought some of these labs to public attention.
This suggests, according to the advisory, a long-term strategy rather than an opportunistically seized opportunity.
The phrasing "probably with the knowledge of the Chinese government" is a textbook case of intelligence language in a public document.
It is not evidence presented; it is a confidence assessment: the agencies indicate they judge state knowledge, even support, to be probable, without publicly furnishing the demonstration.
This semantic caution is both a strength and a weakness: a strength, because it avoids unsubstantiated assertion; a weakness, because it is precisely on this sentence that the political and diplomatic battle will be fought.
Finally, the joint, triple signature — CISA + NSA + FBI — is an institutional signal in itself.
Three agencies of this stature are not mobilized for an ordinary intellectual-property dispute.
The message is that massive extraction of AI capabilities is now treated as a US national security matter, on par with classic industrial-secret theft or state economic espionage.
This is a regime change: "distillation" moves from researchers' vocabulary to counterintelligence analysts'.

<a id="s09-11"></a>
### 9.10 The six labs: what the advisory documents, actor by actor

The advisory names six labs, and the verified facts detail three of them with precise figures.
DeepSeek is accused of having targeted its own R1 and V3 models — meaning: the successive versions of its flagship models were allegedly trained, at least in part, on outputs extracted from American models.
The most pointed fact concerns the "$5.6 million" training cost claimed for one of these models, judged misleading by the advisory.
This figure had traveled the world as proof that a leading model could be trained for a fraction of American labs' costs.
If this cost is misleading — because it would omit the value of the distilled data, i.e., the real production cost of the imitated capabilities — then the entire narrative of Chinese efficiency cracks.
The argument is formidable in its simplicity: you cannot compare the marginal cost of a copy with the full cost of an invention.
It is like comparing the price of a photocopy with the cost of writing the book.
Moonshot AI, the publisher of Kimi, is attributed over 23 million exchanges with the targeted models to train Kimi K3.
23 million exchanges is an order of magnitude that speaks: at a few seconds per exchange, that represents years of cumulative query time, spread over massive parallel infrastructure.
This is not "prompting"; it is industrial siphoning.
The fact that the alleged objective is named — Kimi K3 — anchors the accusation in a real product, which makes it verifiable in principle and contestable in practice.
Alibaba, for its Qwen family, is attributed over 151 million exchanges between May and July 2026.
This is the most massive figure of the three, and its time window — three months — also makes it the most intense: about 50 million exchanges per month, or over one and a half million per day.
At this scale, the alleged operation is no longer a discreet activity; it is API consumption comparable to that of a very large enterprise customer — except that the alleged purpose is not use, but extraction.
The three other labs — MiniMax, StepFun, Z.AI — are named in the advisory without the verified facts detailing their volumes.
Their mere presence on the list widens the picture: according to the agencies, this is not an isolated actor that cheated, but an entire ecosystem whose leading labs allegedly resorted to the same methods.
It is this systemic dimension that justifies, in the authors' eyes, the move to a public advisory: we are no longer dealing with individual cases; we are describing an alleged sector-wide practice.
Note what the advisory does not say, or not yet: it does not quantify the share of Chinese models' capabilities that would come from distillation as opposed to in-house research; it does not publish the technical evidence for the attribution; and it does not distinguish, within the cited volumes, between legitimate API use and extraction.
These blind spots are normal for an advisory — which is a warning, not a case file — but they bound what can be concluded from it.
What can be concluded with certainty is that the United States now considers protecting the outputs of its models a national security matter, and that American labs will have to treat their APIs as potential exfiltration surfaces.

<a id="s09-12"></a>
### 9.11 Extraction channels and targets: the anatomy of an alleged campaign

The advisory does not merely name actors and volumes: it describes the channels through which the alleged extraction transited, and this is perhaps its most technically instructive part.
Five channels are documented: the targeted models' native APIs, remote cloud providers, third-party aggregators, gray-market "transfer stations," and premium subscriptions bought in bulk.
Each of these channels tells a facet of the operation.
The native APIs are the direct route: querying Claude, GPT, Gemini, or Grok via their official interfaces, formally respecting the terms of use while diverting them from their purpose.
This is the simplest channel and the most exposed to detection — hence the need for the others.
Remote cloud providers serve to distribute queries geographically: by routing through cloud infrastructure in varied jurisdictions, traffic is fragmented and attribution complicated.
Third-party aggregators — those platforms that resell access to several models via a single API — offer a double advantage: pooling of access and an additional layer of opacity between the end operator and the model provider.
The gray-market "transfer stations" are the most explicitly clandestine channel: intermediaries that would supply model access while bypassing controls, quotas, and identity verification.
The expression itself, in quotation marks in the advisory, evokes organized hubs, not individual tinkering.
Finally, premium subscriptions bought in bulk show that even consumer offerings — individual paid plans — were allegedly instrumentalized at industrial scale: thousands of premium accounts, each with its quotas, aggregated into massive extraction capacity.
This is the industrialization of circumvention via the low end.
Faced with this diversity of channels, defense cannot be simple rate-limiting: it requires behavioral detection capable of distinguishing, within billions of queries, systematic extraction patterns from intensive legitimate use.
This is a very-large-scale anomaly-detection problem, in which the alleged attackers hold a structural advantage: they know the thresholds, since they test them constantly.
The documented extraction targets are equally revealing: chain-of-thought, reasoning, coding, and agentic skills.
These are not factual knowledge being siphoned; they are capabilities.
Chain-of-thought — the intermediate reasoning traces that some models expose — is particularly precious: it contains not only the answer, but the path leading to it, i.e., exactly what is needed to train a student to reason like the teacher.
Reasoning in the broad sense, coding, and agentic skills — the ability to plan, use tools, and carry out multi-step tasks — are the differentiating capabilities of frontier models in 2026.
By targeting these four families, the alleged campaign was not seeking to copy models; it was seeking to copy what makes them valuable: their ability to think and act.
This is consistent with the hypothesis of distillation as a catch-up strategy: there is no need to reproduce an entire GPT or Claude; one needs to extract the skills hardest to develop through in-house research.
And this is also what makes detection so difficult: queries that extract reasoning look, one by one, like perfectly legitimate uses — it is only their volume, systematicity, and coverage that betray the intent.

<a id="s09-13"></a>
### 9.12 Legal scope: an advisory is neither an indictment, nor a sanction, nor a judgment

This is the point the dossier must hammer home without ambiguity: AA26-251A is an advisory, not an indictment.
The registers must not be confused, because the entire public and diplomatic debate to come will be fought over this confusion.
A US agencies' advisory is a public warning and attribution document: it describes a threat, names alleged actors, documents modi operandi, and recommends defensive measures.
What it is not: an indictment (no criminal charges against named persons), an Entity List designation (no automatic trade sanctions against the cited entities), or a court judgment (no adversarial procedure, no evidence administered before a judge, no conviction).
Each of these distinctions has practical consequences.
No indictment means no criminal prosecution is initiated on the basis of this document alone: the agencies inform; they do not prosecute.
The named labs remain, legally, actors like any others until a formal procedure is opened — which obviously does not prevent immediate reputational and commercial consequences.
No Entity List designation means US export controls do not automatically apply to these entities as a result of the advisory: selling chips or software to DeepSeek or Moonshot does not become illegal on September 8 on the sole grounds of publication.
This is a nuance that hurried commentators often miss, yet it is central: the advisory is a signal, not a sanctions regime.
No court judgment means the alleged facts have not been established according to judicial standards of proof: no adversarial debate, no independent expert assessment, no organized right of reply.
The figures — billions of tokens, 23 million exchanges, 151 million exchanges — are allegations by intelligence and security agencies, not facts established by a court.
That does not make them false, but it strictly bounds what can be asserted from them: this dossier reports them as documented accusations in an advisory, not as judicial truths.
Why publish an advisory rather than pursue prosecutions or sanctions?
Several readings are possible, and they are not mutually exclusive.
The "deterrence" reading: naming publicly signals that we see, and encourages cessation without incurring the costs and risks of a procedure.
The "defensive mobilization" reading: the primary objective would be to alert American operators so they harden their APIs, more than to punish the alleged extractors.
The "graduated escalation" reading: the advisory would be the first step in a sequence that could lead, later, to sanctions or prosecutions — its publication testing reactions before moving to heavier measures.
Whichever reading, the effect is already real: American labs now know they are expected to treat massive extraction as a national security threat, with everything that implies in terms of logging, detection, cooperation with agencies, and potentially restrictions on their own openness policies.
And the named labs know they now operate under the public gaze of US agencies — which changes their calculations, whether or not they are guilty of the alleged facts.

