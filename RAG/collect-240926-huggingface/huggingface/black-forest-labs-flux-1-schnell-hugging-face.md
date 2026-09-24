---
id: collect-240926-huggingface/huggingface/black-forest-labs-flux-1-schnell-hugging-face
title: "black-forest-labs-flux-1-schnell-hugging-face"
domain: huggingface
role: reference
task: reference
actors: []
dates: []
keywords: ["acquisition", "apache", "diffusion", "distillation", "exploit", "gpu", "inference", "regulation"]
source: docs/RAG/clean_en/huggingface/black-forest-labs-flux-1-schnell-hugging-face.md
source_anchor: ""
source_lines: [1, 63]
sha256: 6ca47b79cfed5c98b9a108067f7c5ef111d5196a275f1bb00fc2cc820fc0227f
---

# black-forest-labs-flux-1-schnell-hugging-face

<!-- source: https://huggingface.co/black-forest-labs/FLUX.1-schnell -->

`FLUX.1 [schnell]` is a 12 billion parameter rectified flow transformer capable of generating images from text descriptions.
For more information, please read our blog post.

1. Cutting-edge output quality and competitive prompt following, matching the performance of closed source alternatives.
2. Trained using latent adversarial diffusion distillation, `FLUX.1 [schnell]` can generate high-quality images in only 1 to 4 steps.
3. Released under the `apache-2.0` licence, the model can be used for personal, scientific, and commercial purposes.

We provide a reference implementation of `FLUX.1 [schnell]`, as well as sampling code, in a dedicated github repository.
Developers and creatives looking to build on top of `FLUX.1 [schnell]` are encouraged to use this as a starting point.

The FLUX.1 models are also available via API from the following sources

- bfl.ml (currently `FLUX.1 [pro]` )
- replicate.com
- fal.ai
- mystic.ai

`FLUX.1 [schnell]` is also available in Comfy UI for local inference with a node-based workflow.

To use `FLUX.1 [schnell]` with the 🧨 diffusers python library, first install or upgrade diffusers

```
pip install -U diffusers
```
Then you can use `FluxPipeline` to run the model

```
import torch
from diffusers import FluxPipeline
pipe = FluxPipeline.from_pretrained("black-forest-labs/FLUX.1-schnell", torch_dtype=torch.bfloat16)
pipe.enable_model_cpu_offload() #save some VRAM by offloading the model to CPU. Remove this if you have enough GPU power
prompt = "A cat holding a sign that says hello world"
image = pipe(
    prompt,
    guidance_scale=0.0,
    num_inference_steps=4,
    max_sequence_length=256,
    generator=torch.Generator("cpu").manual_seed(0)
).images[0]
image.save("flux-schnell.png")
```
To learn more check out the diffusers documentation

- This model is not intended or able to provide factual information.
- As a statistical model this checkpoint might amplify existing societal biases.
- The model may fail to generate output that matches the prompts.
- Prompt following is heavily influenced by the prompting-style.

The model and its derivatives may not be used

- In any way that violates any applicable national, federal, state, local or international law or regulation.
- For the purpose of exploiting, harming or attempting to exploit or harm minors in any way; including but not limited to the solicitation, creation, acquisition, or dissemination of child exploitative content.
- To generate or disseminate verifiably false information and/or content with the purpose of harming others.
- To generate or disseminate personal identifiable information that can be used to harm an individual.
- To harass, abuse, threaten, stalk, or bully individuals or groups of individuals.
- To create non-consensual nudity or illegal pornographic content.
- For fully automated decision making that adversely impacts an individual's legal rights or otherwise creates or modifies a binding, enforceable obligation.
- Generating or facilitating large-scale disinformation campaigns.

- Downloads last month
- 520,328
