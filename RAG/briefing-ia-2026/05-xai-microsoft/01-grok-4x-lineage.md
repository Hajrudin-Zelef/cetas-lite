---
id: briefing-ia-2026/05-xai-microsoft/01-grok-4x-lineage
title: "xAI: the Grok 4.5→4.8 lineage"
domain: xai-microsoft
role: deep-dive
task: model-release
actors: ["AWS", "Anthropic", "Microsoft", "OpenAI", "xAI"]
dates: ["2026-07", "2026-07-08", "2026-08-12", "2026-09-13", "2026-09-21"]
keywords: ["grok", "grok 4", "agentic", "agents", "aws", "bedrock", "benchmark", "benchmarks", "compute", "context window", "copilot", "distillation"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s05"
source_lines: [5789, 5934]
canonical_for: ["grok-lineup"]
sha256: 974c185c21e55a1fb691ce622d48e7f3231e4b8c873091fe0d1157d8a38359a1
---

# xAI: the Grok 4.5→4.8 lineage

<a id="s05"></a>
## 5. xAI + Microsoft in depth

<a id="s05-2"></a>
### Section introduction: two trajectories converging

The first half of 2026 in artificial intelligence reads largely through two names: xAI and Microsoft.
The first is Elon Musk's creature, born to challenge OpenAI, and burning capital at a pace that even Silicon Valley watches with a raised eyebrow.
The second is OpenAI's former privileged partner, which this year decided to stand on its own two feet when it comes to models.
What makes 2026 singular is that these two trajectories, which started out in opposite directions, end up crossing: Microsoft integrates xAI's Grok models into Copilot, xAI's first incursion into the Microsoft ecosystem despite the personal and industrial rivalry between Elon Musk and Microsoft.
On one side, xAI stacks iterations of Grok 4 — 4.5, 4.6, 4.7, then the thunderous announcement of 4.8 — with a strategy of post-training "upgrades" rather than new foundation models.
On the other, Microsoft deploys the MAI family: seven models announced as trained "from scratch, no distillation," covering reasoning, code, image, transcription, and voice.
Between the two, a 550,000-GPU supercomputer, a record IPO, a $1.25 trillion merger, an always-on assistant, a chip-to-cloud hardware platform, and a philosophical code of conduct dividing the industry.
This section unwinds each of these threads, model by model, date by date, relying exclusively on the verified facts of the dossier.
Every price, every date, every figure cited below comes from that verified base: nothing is extrapolated, nothing is casually rounded.
The analysis, on the other hand, is free: that is where this dossier adds its value, by connecting the facts to one another and exposing what they mean for what comes next.

<a id="s05-3"></a>
### XAI — Grok 4.5 (July 8): the post-training upgrade that changes the game

On July 8, 2026, xAI updates Grok to version 4.5.
This is not a new foundation model: it is a post-training upgrade applied on top of the V9 base.
The distinction is crucial and deserves to be understood, because it structures xAI's entire 2026 strategy.
A new foundation model requires a full pre-training run, months of compute on tens of thousands of GPUs, and a cost counted in tens of millions of dollars.
A post-training upgrade, by contrast, takes an existing model and improves it through fine-tuning, reinforcement, and behavior-optimization techniques — much faster, much cheaper, with targeted gains.
xAI is therefore betting on rapid iteration rather than costly rupture: the same V9 foundation, but refined behavior, notably for code and agents.
The context window is 500,000 tokens — 500K ctx — which places Grok 4.5 in the category of very-long-memory models, capable of ingesting entire codebases or long documents in a single prompt.
Pricing is listed at $2 per million input tokens and $6 per million output tokens — $2/$6.
That is aggressive pricing for a model presented as Opus-class in programming and agentic use.
The "Opus-class" label is not trivial: it explicitly designates Anthropic's family of high-end models, the reference of the moment for code and agents.
xAI is thus claiming that Grok 4.5 plays in the same league as Opus for these uses, at a fraction of the usual flagship price positioning.
Another structural fact: the model was co-developed with Cursor.
Cursor is the AI-assisted code editor that has won over developers, and co-development means Grok 4.5 was fine-tuned on real coding workflows, with usage feedback from a production tool.
That is validation through use, not just through benchmarks: the model is shaped for the edit, test, and debug loops that developers live through every day.
For developers and enterprises, the message is twofold: a serious agentic model, trained with a field partner, at a price that makes intensive use conceivable.
For competitors, the message is more brutal: xAI can now iterate at a pace that classic pre-training cycles do not allow.
The post-training upgrade strategy theoretically allows several qualitative leaps per year on the same foundation, where a classic lab produces only one or two.
The question of durability remains: a post-training upgrade eventually plateaus, and sooner or later a new foundation will be needed.
But in July 2026, xAI is not there yet: Grok 4.5 is presented as a show of force for the iterative method.
The choice of date — July 8, in the middle of summer, far from the big conferences — also suggests confidence: no need for a keynote to get people talking about a model that sells on performance.
Finally, this $2/$6 will become the pricing signature of the entire Grok 4 lineage: neither 4.6 nor 4.7 will move it, as we will see.

<a id="s05-4"></a>
### Grok 4.5 vs 4.6 vs 4.7: a comparative anatomy of three upgrades

Grok 4.6 ships on August 12, 2026, one month and four days after 4.5.
Same logic as its predecessor: a post-training upgrade on the V9 base, not a new foundation model.
The cadence is therefore monthly, which is remarkable for models of this class and confirms xAI's iterative strategy.
What distinguishes 4.6 from 4.5 is not so much the method as the distribution: Grok 4.6 becomes available on AWS Bedrock starting August 19, one week after its release.
Bedrock, Amazon's managed-model platform, exposes Grok cross-region at the same $2/$6 pricing.
Arriving on Bedrock is an enterprise signal: it is the gateway to CIOs, regulated architectures, multi-region deployments with AWS guarantees.
A model available on Bedrock is no longer just an API product for developers: it is an infrastructure component that large enterprises can buy the way they buy compute.
In parallel, in August, Grok 4.5 and 4.6 arrive on GitHub Copilot.
This is xAI's first breakthrough into the Microsoft development universe: Copilot users can now choose Grok models.
The fact that both versions coexist on Copilot in August suggests a smooth transition, with users able to compare or stay on the version they prefer.
Grok 4.7, for its part, ships on September 21, 2026, after five delays.
Five delays: that is the most telling fact about this version, and it deserves a long pause (see the dedicated sub-section below).
On pricing, 4.7 changes nothing: $2/$6, prices unchanged.
Keeping the price constant across three successive versions is a strong strategic choice.
In an industry where every capability gain usually comes with a price hike or price segmentation, xAI freezes its entry ticket.
The implicit message: performance goes up, the price does not move, and the value gap with competitors widens with each iteration.
Let us now compare the three versions on what the verified facts allow us to say.
Method: identical for all three — post-training upgrade on the V9 base.
Context window: 500K ctx documented for 4.5; the facts specify no change for 4.6 and 4.7, and the anti-fabrication rule forbids assuming one.
Price: $2/$6 for all three, no exceptions.
Cursor co-development: documented for 4.5; not mentioned for 4.6 and 4.7.
Distribution: 4.6 stands out via Bedrock (August 19, cross-region); 4.5 and 4.6 share GitHub Copilot in August; 4.7 inherits this distribution with no newly documented fact.
Positioning: "Opus-class" in coding/agents for 4.5; the facts do not explicitly renew the label for 4.6 and 4.7, which does not prevent noting that the lineage claims the same ambition.
Release: 4.5 on July 8 with no documented drama; 4.6 on August 12 likewise; 4.7 on September 21 after five delays and a public downgrade by Musk himself.
This comparative table reveals a two-stage strategy: first prove the method (4.5, co-developed with Cursor, positioned Opus-class), then industrialize distribution (4.6, Bedrock + Copilot), then pay the price of ambition (4.7, chained delays).
The regularity of price and method contrasts with the irregularity of the schedule: one month between 4.5 and 4.6, then more than a month and five delays for 4.7.
It is a sign that post-training iteration, fast when all goes well, remains subject to the vagaries of validation and quality.
A lab that promises a monthly rhythm exposes itself to having to choose between hitting the date and hitting quality.
For 4.7, xAI visibly chose — five times — to delay rather than ship a model below expectations.
Whether the market remembers quality prudence or schedule instability remains to be seen.

<a id="s05-5"></a>
### Grok 4.7 (September 21): chronicle of a chaotic launch

On September 21, 2026, Grok 4.7 finally ships.
The word "finally" is not rhetorical: the version was delayed five times before seeing the light of day.
Five delays for a post-training upgrade is a signal that must be decoded carefully.
In the model industry, an isolated delay is banal: internal validation reveals a problem, you fix it, you slip a week.
Two delays start raising questions; five draw a pattern.
Either the model had persistent quality problems the team could not stabilize, or the stated ambition for this version exceeded what the V9 base could deliver through simple post-training.
The two hypotheses are not mutually exclusive, and the verified facts do not allow deciding between them: the dossier's rule forbids speculating on the technical cause.
What the facts do say, however, is the public dimension of the episode: Elon Musk himself publicly downgraded Grok 4.7 before its release.
A founder who lowers expectations on his own product before launch is a rare and calculated gesture.
Normally, a lab's communication consists of inflating expectations until D-Day, then managing disappointment afterward.
Here, Musk inverts the sequence: he announces in advance that 4.7 will fall short of what was hoped.
Several readings are possible, and they illuminate crisis management à la Musk.
First reading, the most charitable: radical honesty as a trust strategy — better low expectations exceeded than a promise unkept.
Second reading: preparing the ground for 4.8, announced eight days earlier (September 13), with its 2.5 trillion parameters and its in-house C++ stack.
By presenting 4.7 as a modest step, Musk shifts attention to the real event to come: 4.8.
Third reading: constraint — after five delays, something had to be said, and publicly acknowledging the model's limits avoided a speculative frenzy followed by disillusionment.
Whichever reading, the 4.7 episode marks a break in tone in xAI's communication.
Until then, the Grok 4 lineage advanced through shows of force: 4.5 co-developed with Cursor and positioned Opus-class, 4.6 deployed on Bedrock and Copilot.
With 4.7, xAI discovers defensive communication — and does so with the means at hand, i.e., its founder's brutal frankness.
On the commercial front, the price stays at $2/$6, unchanged.
That is consistent with the public downgrade: you cannot both say a model falls short of expectations and raise its price.
Keeping the tariff is therefore also an implicit compensating gesture toward users.
On the distribution front, 4.7 follows in continuity: the lineage has been on GitHub Copilot since August, on Bedrock since August 19, and nothing in the facts indicates a withdrawal.
The 4.7 episode finally raises a fundamental question about the post-training upgrade strategy.
If even a seasoned fine-tuner like xAI must delay an upgrade on a mastered base (V9) five times, then rapid iteration has its limits.
Each post-training cycle is a risk: you can degrade what worked while trying to improve what did not.
Musk's "public downgrade" can then be read as the admission that this cycle produced a modest gain — an admission few labs would make.
4.7 must also be placed in its calendar window: released September 21, eight days after the Grok 4.8 announcement (September 13).
That proximity is not trivial: 4.7 arrives while attention is already fixed on 4.8 and its 2.5 trillion parameters.
One may see a scheduling clumsiness, or on the contrary fine management: 4.7 ships discreetly, without media pressure, while 4.8 captures the light.
Either way, September 21, 2026 will go down in xAI's annals as the day the lab learned to publicly manage imperfection.
That is no small thing: in an industry obsessed with ever-ascending benchmark curves, admitting a plateau is an act of maturity — or of forced lucidity.

<a id="s05-6"></a>
### Grok 4.8 (announced September 13): the 2.5-trillion-parameter bet

On September 13, 2026, xAI announces Grok 4.8.
Two pieces of information, and two only, are verified at this stage: the model will have 2.5 trillion parameters — 2.5T — and will rest on an in-house C++ stack.
Each of these two facts deserves development, because together they sketch a major architectural bet.
Let us start with the 2.5 trillion parameters.
It is an order of magnitude that crushes everything made public in the industry to date.
To fix ideas without inventing a quantified comparison: the largest publicly documented dense models count in the hundreds of billions of parameters, and the most ambitious MoE (mixture of experts) architectures showed around a trillion.
2.5T is therefore a substantial leap beyond the public state of the art — a leap that immediately raises the training question.
Training 2.5 trillion parameters demands colossal compute infrastructure, and that is where Colossus v2 enters the scene with its 550,000 GPUs, presented as the world's largest AI supercomputer.
The link is direct: you do not announce 2.5T without the supercomputer to go with it, and you do not build a 550,000-GPU supercomputer without a model worthy of the investment.
The 4.8 announcement thus gives, retrospectively, the industrial justification for Colossus v2.
Second fact: the in-house C++ stack.
This may be the most revealing of the two facts, because it speaks of method as much as size.
Nearly the entire industry relies on Python stacks (PyTorch, JAX) for training and inference, with optimization layers in C++/CUDA underneath.
Announcing an in-house C++ stack is asserting that xAI is rewriting the plumbing itself, from the ground up.
The possible motivations are legible even without speculating on technical details: total performance control, elimination of generalist-framework overhead, fine optimization for Colossus v2's specific hardware, and independence from architectural choices imposed by dominant frameworks.
It is also a recruiting and culture signal: xAI positions itself as a systems-engineering lab, not just ML researchers.
The downside is well known: an in-house stack means maintenance debt, a learning curve for new arrivals, and a risk of subtle bugs where proven frameworks have already taken the hits.
But for a lab targeting 2.5T parameters, the math may justify it: at that scale, each percent of infrastructure efficiency gained translates into millions of dollars saved and weeks of schedule recovered.
Note what the announcement does not say: no release date, no price, no context window, no benchmarks.
It is an announcement of intent and architecture, not a product sheet.
Eight days later Grok 4.7 will ship — modest, delayed five times, downgraded by Musk.
The contrast between the two announcements is striking and probably deliberate: 4.7 manages the present with humility, 4.8 sells the future with audacity.
The September 13 (4.8 announcement) then September 21 (4.7 release) sequence can be read as choreography: first make people dream of 2.5T, then deliver the intermediate step without pressure.
The central question remains: will 4.8 be a new foundation model or a new upgrade?
The facts do not say, and it would be unwise to conclude.
But 2.5 trillion parameters on an in-house C++ stack does not look like a simple post-training of V9: the scale and the stack rewrite suggest a foundation effort.
If so, xAI would be returning to the heavy method after three light upgrades — proof that rapid iteration and costly rupture do not exclude each other, but succeed one another in a cycle.
Finally, 4.8 raises the price question: the entire 4.5–4.7 lineage is at $2/$6, but a 2.5T-parameter model cannot durably be served at the same tariff without massive subsidy.
xAI already burns $1.5 billion per month; serving 2.5T at 4.7's price would be a financial bet as much as a technical one.
The September 13 announcement is therefore also an implicit promise made to investors: the supercomputer, the in-house stack, and the scale must converge toward a viable economics.

