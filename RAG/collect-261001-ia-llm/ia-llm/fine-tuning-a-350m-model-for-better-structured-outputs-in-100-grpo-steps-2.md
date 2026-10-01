---
id: collect-261001-ia-llm/ia-llm/fine-tuning-a-350m-model-for-better-structured-outputs-in-100-grpo-steps-2
title: "Fine-tuning a 350M Model for Better Structured Outputs in 100 GRPO Steps"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba"]
dates: []
keywords: ["fine-tuning", "benchmark", "gguf", "gpu", "latency", "llama", "llama.cpp", "lora", "training"]
source: docs/RAG/collect-261001-ia-llm/fine-tuning-a-350m-model-for-better-structured-outputs-in-100-grpo-steps.md
source_anchor: ""
source_lines: [145, 274]
sha256: 59b6c611447abcb24d64435dd1c259c73f5beefcab1a7a16e9d018601e0b18c1
---

# Fine-tuning a 350M Model for Better Structured Outputs in 100 GRPO Steps

```
from trl import GRPOConfig
training_args = GRPOConfig(
    output_dir="./outputs/lfm25-350m-nemotron-schema-grpo",
    learning_rate=5e-5,
    max_steps=100,
    warmup_steps=10,
    num_generations=8,              # completions sampled per prompt group
    per_device_train_batch_size=4,
    gradient_accumulation_steps=8,  # 4 prompt groups per optimizer step
    steps_per_generation=2,
    max_completion_length=1024,     # room for nested JSON
    mask_truncated_completions=False,
    temperature=1.1,                # hotter sampling keeps groups varied
    beta=0.01,                      # KL penalty toward the reference model
    reward_weights=[1.0, 0.5, 2.0], # json_format, field_count, schema_validation
    logging_steps=1,
    save_steps=100,
)
```
As you can see in the notebook, over the run, all three reward components climb, the KL from the reference model lifts off zero after warmup, and the truncated-completion fraction stays near zero.

Finally, we merge the LoRA adapter back into the base weights and save it as a single self-contained checkpoint, ready to convert to GGUF for serving:

```
MERGED_DIR = f"{training_args.output_dir}-merged"
merged_model = trainer.model.merge_and_unload()
merged_model.save_pretrained(MERGED_DIR)
tokenizer.save_pretrained(MERGED_DIR)
```
After GRPO fine-tuning, we rerun the IFStruct evaluation. For this, we need to convert the merged model checkpoint into a BF16 GGUF. The converter script ships with the llama.cpp source, so we clone the repo once and install the converter's `gguf` package.

```
git clone --depth 1 https://github.com/ggml-org/llama.cpp
pip install ./llama.cpp/gguf-py
mkdir -p models
python llama.cpp/convert_hf_to_gguf.py \
  PATH_TO_YOUR_MERGED_MODEL \
  --outfile ./models/lfm25-350m-grpo-bf16.gguf \
  --outtype bf16
```
Then we serve the merged model with the following command:

```
llama-server \
  -m ./models/lfm25-350m-grpo-bf16.gguf \
  --alias lfm25-350m-grpo-structured-output \
  -c 32768 \
  -np 4 \
  -ngl 99 \
  --host 127.0.0.1 \
  --port 8081
```
Then, we will run the full IFStruct evaluation again with the fine-tuned model:

```
uv run ifstruct-eval \
  --model lfm25-350m-grpo-structured-output \
  --base-url http://localhost:8081/v1 \
  --api-key dummy \
  --dataset data/test.jsonl \
  --results-file results/lfm25-350m-grpo.json \
  --n-threads 4 \
  --max-tokens 2048 \
  -v
```
```
============================================================
Model: lfm25-350m-grpo-structured-output
============================================================
Overall: 594/2000 passed (29.7%)
Average latency: 1518ms
By format:
  JSON: 319/1000 passed (31.9%)
  YAML: 275/1000 passed (27.5%)
By top-level structure:
  Wrapper key 300/1011 passed (29.7%)
  Bare list   294/989 passed (29.7%)
By entity type:
  test__camera_review                 5/83 passed (6.0%)
  test__clinical_trial                31/104 passed (29.8%)
  test__conference_schedule           11/87 passed (12.6%)
  test__escaping__bug_report_batch    32/89 passed (36.0%)
  test__escaping__config_snippet_audit 24/85 passed (28.2%)
  test__escaping__customer_email_thread 9/73 passed (12.3%)
  test__escaping__dialogue_sample     17/95 passed (17.9%)
  test__escaping__interview_transcript_segment 13/80 passed (16.2%)
  test__escaping__log_parser_examples 33/72 passed (45.8%)
  test__escaping__pr_discussion       26/87 passed (29.9%)
  test__escaping__repro_steps_batch   23/73 passed (31.5%)
  test__escaping__screenplay_scene    34/92 passed (37.0%)
  test__escaping__short_story_chapter 24/84 passed (28.6%)
  test__escaping__support_ticket_batch 36/73 passed (49.3%)
  test__escaping__terminal_session_notes 23/70 passed (32.9%)
  test__event_ticket_booking          62/107 passed (57.9%)
  test__gpu_review                    7/94 passed (7.4%)
  test__invoice                       36/86 passed (41.9%)
  test__job_posting                   33/85 passed (38.8%)
  test__real_estate_listing           32/82 passed (39.0%)
  test__recipe                        7/70 passed (10.0%)
  test__rental_car_booking            37/79 passed (46.8%)
  test__scientific_experiment         14/69 passed (20.3%)
  test__travel_itinerary              25/81 passed (30.9%)
Common errors:
  7331x required field missing
  890x wrong item count
  555x type mismatch
  102x expected bare list, got wrapper
   62x extraneous field 'metadata.tone'
   55x 6 is greater than maximum 5
   49x extraneous field 'speaker_labels'
   47x extraneous field 'tone'
   44x 'cups' not in allowed values ['mg', 'g', 'kg', 'oz', 'lb', 'ml', 'l', 'cl', 'dl'
   44x extraneous field 'notes'
```
Comparing the two runs on the identical serving stack:

| IFStruct group | base | GRPO-tuned | Δ | 
|---|---|---|---|
| **Overall** | 22.6% | **29.7%** | **+7.1** | 
| JSON | 18.0% | 31.9% | +13.9 | 
| YAML | 27.2% | 27.5% | +0.3 | 
| Wrapper key | 28.5% | 29.7% | +1.2 | 
| Bare list | 16.6% | 29.7% | +13.1 | 

The gains land exactly where the training aimed: the JSON pass rate rises by nearly 14 points (18.0% → 31.9%), while YAML stays mostly the same. While this is still below the Qwen3.5-2B score of 33.15%, it shows that even light task-specific fine-tuning can bring a small model close to a larger one.

A short GRPO run with about 500 samples and 100 steps can lift a small 350M parameter model from 22.6% to 29.7% on IFStruct. The takeaway is that a cheap, task-specific reward signal can make a small model substantially more reliable about *form*, closing much of the gap to models several times its size.

To reproduce or extend this work, see the original IFStruct v1.0 blog post, the Liquid4All/ifstruct benchmark repo, and the LiquidAI/ifstruct-v1.0 dataset.
