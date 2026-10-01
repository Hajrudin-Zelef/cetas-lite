---
id: collect-261001-ia-llm/ia-llm/nemotron-3-5-content-safety-customizable-multimodal-safety-for-global-enterprise-ai-2
title: "nemotron-3-5-content-safety-customizable-multimodal-safety-for-global-enterprise-ai"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Baseten", "Hugging Face", "Nvidia", "OpenRouter", "SGLang", "vLLM"]
dates: []
keywords: ["multimodal", "benchmark", "benchmarks", "claude", "cost", "gpu", "guardrails", "inference", "jailbreak", "latency", "license", "licenses"]
source: docs/RAG/collect-261001-ia-llm/nemotron-3-5-content-safety-customizable-multimodal-safety-for-global-enterprise-ai.md
source_anchor: ""
source_lines: [80, 121]
sha256: e4b790923ac7a49947942a767dd44cde6a4f6a9ed003177ce42786ec3f754a52
---

# nemotron-3-5-content-safety-customizable-multimodal-safety-for-global-enterprise-ai

- **Multilingual text safety data** from Nemotron Safety Guard Dataset v3, sampled from culturally nuanced subsets with proportional representation across safety categories and safe/unsafe splits.
- **Human-annotated multimodal data** collected in English by NVIDIA, translated into 12 languages. Critically,**99% of training images are real photographs** —not synthetic generations. This directly addresses a known weakness in the multimodal safety benchmark landscape, where existing datasets like VLGuard and MM-SafetyBench rely heavily on SDXL-generated images that lack the cultural texture and adversarial complexity of production content. While not all of these real images could be released due to licensing constraints, we are still able to release a subset of images from Wikimedia and synthetic generation.
- **Safe multimodal data** from Nemotron VLM Dataset v2, covering scanned documents, charts, papers, and diagrams with associated queries—ensuring the model does not over-flag benign professional content.
- **Reasoning traces** derived from chain-of-thought outputs produced by larger teacher models—Qwen 397B and then shortened using Qwen 80B—are used to teach the model how to reason.
- **Topic following data** from the CantTalkAboutThis dataset consisting of policy-specification/verdict pairs across a range of enterprise deployment scenarios (healthcare, finance, banking, education, etc.).
- **Synthetic data** accounting for roughly 10% of total training volume, used primarily to diversify jailbreak patterns, generate rare policy violation examples, and produce multimodal adversarial cases.

Nemotron 3.5 Content Safety was evaluated across multilingual, multimodal, and custom-policy safety benchmarks, including VLGuard, MM-SafetyBench, PolyGuard, RTP-LX, Aya Redteaming, XSafety, MultiJail, Aegis, Dynaguardrail, and CoSA. These evaluations reflect the core production challenge for enterprise safety: applying consistent guardrails across global languages, text and image inputs, and domain-specific policies without adding significant latency.

Nemotron 3 set a strong baseline with 84% average accuracy on multimodal harmful-content tests and roughly half the latency of LlamaGuard-4-12B. Nemotron 3.5 maintains that compact 4B efficiency while adding custom policy support and reasoning traces.

Across multilingual and multimodal safety benchmarks, Nemotron 3.5 delivers strong harmful-content classification accuracy while maintaining a compact footprint. This matters because many safety models remain English-first, text-only, or too costly to run repeatedly in production pipelines. Nemotron 3.5 is designed to combine multilingual coverage, multimodal classification, custom-policy support, and low-latency deployment in one model.

*Figure 1. Nemotron 3.5 Content Safety delivers strong harmful-content classification accuracy across multilingual and multimodal safety benchmarks, averaging about 85% across the evaluated benchmark set.*

The language-level results highlight why multilingual safety matters for global enterprise AI. On Multilingual Aegis, Nemotron 3.5 averages 96.5% harmful-content classification accuracy across 12 languages. On RTP-LX, it averages 88.8%, for a combined Aegis and RTP-LX average of 92.7%. This consistency helps teams apply the same safety posture across customer, employee, and partner-facing workflows instead of relying on English-only moderation or separate regional safety models.

*Figure 2. Nemotron 3.5 Content Safety averages 97% harmful-content classification accuracy on Multilingual Aegis Cultural + Adapted (prompt classification) (harmful-f1) across 12 languages.*

*Figure 3. Nemotron 3.5 Content Safety averages 89% harmful-content classification accuracy on RTPLX (prompt classification) (harmful-f1) across 12 languages.*

Accuracy alone is not enough for production guardrails. Safety models must also be efficient enough to run before content is processed, returned, or routed downstream. Nemotron 3.5 Content Safety's compact 4B design helps reduce the cost and latency of repeated safety checks, making multilingual and multimodal guardrails practical for real-world AI applications.

The latency profile is unchanged from Nemotron 3 in the default (no THINK) mode. THINK mode adds inference time proportional to trace length, but this overhead is predictable and can be budgeted separately from the synchronous moderation loop—for instance, by running THINK-mode evaluation asynchronously as part of an audit pipeline while the default mode handles real-time decisions.

*Figure 4. Nemotron 3.5 Content Safety achieves 3x lower end-to-end latency on a multimodal benchmark compared to an alternative multimodal safety model.*

Compared to another reasoning safety model, our model generated up to 50% fewer tokens when reasoning is enabled, making it efficient in terms of cost and latency.

A recurring theme in multimodal safety research is the gaps in existing evaluation infrastructure. Nemotron 3.5's development encountered the same gaps documented in the broader literature:

- **Text-only coverage** : The most widely cited safety benchmarks (WildGuard, XSTest, HarmBench) are text-only. Multimodal performance cannot be inferred from text-benchmark results.
- **Synthetic image quality** : Most multimodal benchmarks that exist use AI-generated images (typically SDXL) rather than real photographs, understating the difficulty of real production content.
- **Real-image licensing** : Stock photo licenses prohibit redistribution in AI datasets, creating a structural gap between research and production conditions.

NVIDIA's multimodal training data—with real images and culturally nuanced multilingual prompts—is designed to fill some of these gaps for model training. The benchmark gap for evaluation remains an open problem for the broader safety research community.

Nemotron 3.5 Content Safety is available on Hugging Face under the NVIDIA Open Model License for research and commercial use, along with the training dataset. It supports transformers, vLLM, and SGLang, and is available as a production-grade NVIDIA NIM on build.nvidia.com for teams that need a pre-packaged, GPU-optimized inference microservice.

Developers can also access the model through inference platforms including Baseten, Eigen AI, DeepInfra, OpenRouter, and Vultr.

For custom policy workflows, NVIDIA provides a Claude- and Codex-compatible skill for generating custom policies, along with cookbooks showing how to use the model. Custom policies and reasoning traces help teams adapt safety behavior to domain-specific rules while keeping decisions auditable.
