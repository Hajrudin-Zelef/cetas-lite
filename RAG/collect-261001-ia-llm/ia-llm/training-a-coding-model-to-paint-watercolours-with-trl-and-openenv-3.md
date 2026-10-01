---
id: collect-261001-ia-llm/ia-llm/training-a-coding-model-to-paint-watercolours-with-trl-and-openenv-3
title: "training-a-coding-model-to-paint-watercolours-with-trl-and-openenv"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba"]
dates: []
keywords: ["training", "benchmark", "cost", "gpu", "inference", "qwen"]
source: docs/RAG/collect-261001-ia-llm/training-a-coding-model-to-paint-watercolours-with-trl-and-openenv.md
source_anchor: ""
source_lines: [182, 279]
sha256: 9a0af3a0c75deeb1f976616de7c1e98a41b7ca998052763323f4fa3c390a9ce8
---

# training-a-coding-model-to-paint-watercolours-with-trl-and-openenv

**What the pairwise judge changes is the top.** In `hps-only`, paintings got more
reliable without getting better. The quality of the good ones added just +0.03 to the
group mean, and once HPSv3 saw petals around a centre and a stem, it stopped asking for
more pigment. With the judge on, the other half of the story appears. *Better* here means closer to
the pool, so closer to what I rated as "good" or something I liked more. That added +0.12 in `judge-led` and +0.16
in `hps-led`, the best of each step rose too, and
paint coverage doubled in both runs (0.11 to 0.23, and 0.13 to 0.30) where `hps-only`
barely moved it. With a reference left to beat, a good painting can still get better,
and the model starts being rewarded for using more pigment.

One more finding. The model ignores an explicit instruction, and it is right to. The
system prompt asks for fifteen to thirty filled shapes. If we look at the real mean, it is between 7 and 9, and
`n_shapes` barely correlates with reward in any run (+0.000, −0.14, +0.07). The policy
is not rewarded for obeying that sentence, so it does not obey it.

There is also a ceiling on the `hps-only` route. If every rollout matched its good
ones, that run's mean would sit at 0.771. Whether more steps would break it is
an open question.

The paintings also show something that the tables miss. Within each run, they all look similar. As training advances, the rewards inside each group get closer together, and the median paintings in the opening video look like takes of the same flower. That is GRPO doing what it is designed to do with a pool built from one subject. The pool decides what counts as variety, the same way it decides what counts as quality. If the reward only pays for matching one flower, the model learns to paint that one flower. More diverse output would need a more diverse pool, and building one is more curation work. Surya's newer compositions are an example of this. Alex Yango's animal paintings are the same recipe with different choices in the pool. This is the biggest difference between an aesthetic reward and a maths grader. Behind the number there is a very human job, deciding what belongs in the reward set. Jason Liu's essay on taste says the general version in one line. AI shifted the bottleneck from making to noticing.

Surya closes his blog with some of his favourites. Instead of picking mine, below is a wall with the 178 paintings the reward scored highest across the two judge runs, the same number the reference pool holds, in no particular order. Open it and pick your own.

To choose, you looked at many and kept a few, and that is exactly the job that built the reward of this project. Every painting of every run, with its sketch and its reward, is in the rollouts datasets, and browsable in this gallery.

Since the reward was partly based on my taste, it is fair to close with my verdict as a viewer. To my
eye, `judge-led` is the run that ends up the most diverse and the most artistically
interesting. `hps-led` paints convincing watercolours, but its best ones share a soft,
wet-on-wet look that is almost a style of its own. `hps-only` converges the hardest, and
most of its paintings settle on the same colours. You can judge for yourself in the
gallery, which has every painting of
every run, sortable by step and by reward.

This project is mostly infra. A run needs a trainer, two Spaces, an inference router and a websocket to stay healthy for hours straight, and every piece that fails quietly turns into a wrong number somewhere else. Half the work is checking that the number you read matches what actually happened.

**Failures of the infrastructure were entering the reward as zeros.** A render that
timed out or a scorer that did not answer scored the same as a bad painting, 0.0 inside
the group. Across all my runs that was about 1.5% of rollouts, and in the worst run it
reached 5.2%. That trains the model on noise, so those paths now
return `None` and the rollout is excluded from the group.

**I also found a bug in OpenEnv, and sent the fix upstream.** The client keeps one
persistent websocket, and a socket closed by the far end stayed cached, so every later
call failed even though the environment was healthy. It cost me two half-finished runs
to find it. The fix is submitted
upstream, and the runs launched with
it have been running clean since.

**The reward of a step depends on which references it drew.** The pairwise judge
samples four references per step, so every step faces a different set of rivals, and
some draws are simply harder. GRPO itself is mostly safe, because advantages are
computed inside the group and a hard draw moves the whole group together. The curve I
was reading was not safe, and some of what looked like a bad step was just a hard draw.
The image above is one example. Step 12 scored half a point below step 11 mostly because
it drew the hardest references of the run, while the paintings themselves look close.

Rounded numbers, and only for the runs that finished.

| piece | what it needs | 
|---|---|
| trainer | 1 H200. **18 hours** for 60 steps, about**34** for 110 | 
| HPSv3 | an `a100-large` Space, up for the whole run | 
| the environment | a `cpu-upgrade` Space, which renders comfortably in time | 
| the pairwise judge | Inference Providers quota for `Qwen/Qwen3-VL-30B-A3B-Instruct` | 
| the pool, one-off | openly licensed photos from iNaturalist, Inference Providers quota for the four generators, and rating is your own hours | 

A step is eight rollouts and takes fifteen to eighteen minutes, of which **70 to 80% is
rendering**. A single render takes 69 to 96 seconds against a 90 second deadline. Part of that is
expected, the Space has no GPU, so Chromium renders the WEBGL canvas in software and
p5.brush's bleeds and textures are heavy pixel work. Even so, I expected it to be
faster, and I have not found the full cause.

A scorer can cost more than the training that uses it: HPSv3 has to be up for the whole run, so pause the Space, or set its sleep timer, when the run ends.

Everything runs on HF Jobs, with the environment as a Docker Space and metrics in trackio.

The rule of this project was to reproduce the recipe with every resource open, not to improve it, so a list of untried ideas piled up along the way. These are the ones I would actually try, in order of how much evidence there is.

**Multi-step, and letting the model see what it paints.** This is the first thing I
would try. The original blog trains single turn, so I trained single turn, and in this
setup the model paints with its eyes closed. No image ever goes in, and the only
feedback it gets is one number. The evidence that a feedback loop works is the pool
itself. The reference paintings came from models iterating three rounds under a vision
critic, and the later rounds are better. The material that defines the reward was made
with a loop the policy never gets.

**Smaller models.** There is evidence that 35B is more than needed. In my side
experiments a 4B already wrote valid sketches that passed the gate. If a 4B can learn
this, the cost of the experiment drops by an order of magnitude.

Other ideas on the list are SFT on the pool sources before starting RL, rewarding pigment
explicitly, moving the judge's reference mix from easy to hard as the run advances until
only `love` remains, widening the ten-method allowlist for more visual range (my
attempts on this crashed more sketches and broke the watercolour look), and checking how consistent the pairwise judge really is by
scoring the same image twice.

And the method is not specific to flowers. Alex Yango painted animals with the same mechanism, and Brendan Hogan trained canvas animations against a pool of hand-rated clips. I had also played with something similar before using Simon Willison's pelican benchmark, where code is rendered to an image and scored.

