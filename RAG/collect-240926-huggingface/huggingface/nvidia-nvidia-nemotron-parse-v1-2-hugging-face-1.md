---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-parse-v1-2-hugging-face-1
title: "Load model and processor"
domain: huggingface
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["agentic", "attention", "decode", "inference", "nvidia", "training", "vllm"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-parse-v1-2-hugging-face.md
source_anchor: ""
source_lines: [1, 148]
sha256: 4659052ba3c94218a7b879f40da4059213dd513474f114d0073c828237096af3
---

# Load model and processor

<!-- source: https://huggingface.co/nvidia/NVIDIA-Nemotron-Parse-v1.2 -->

NVIDIA Nemotron Parse v1.2 is designed to understand document semantics and extract text and tables elements with spatial grounding. Given an image, NVIDIA Nemotron Parse v1.2 produces structured annotations, including formatted text, bounding-boxes and the corresponding semantic classes, ordered according to the document's reading flow. It overcomes the shortcomings of traditional OCR technologies that struggle with complex document layouts with structural variability, and helps transform unstructured documents into actionable and machine-usable representations. This has several downstream benefits such as increasing the availability of training-data for Large Language Models (LLMs), improving the accuracy of extractor, curator, retriever and AI agentic applications, and enhancing document understanding pipelines.

This model is ready for commercial use. 

Compared to NVIDIA-Nemotron-Parse v1.1 the following changes are made in v1.2:

NVIDIA-Nemotron-Parse v1.2 extends the prompt interface with a fourth prompt token category — **Text-in-picture prompts** — controlling whether the model extracts text from embedded images (e.g. photographs, figures):

- `<predict_text_in_pic>` : extract text found inside pictures.
- `<predict_no_text_in_pic>` : skip text extraction from pictures, reducing generation time by avoiding unnecessary tokens.

The full v1.2 prompt now consists of four prompt tokens (one from each category):

`</s><s><predict_bbox><predict_classes><output_markdown><predict_no_text_in_pic>`

**Important:** NVIDIA-Nemotron-Parse v1.1 used a three-token prompt (e.g. `</s><s><predict_bbox><predict_classes><output_markdown>`). In v1.2, the fourth token is required. The prompt is not validated at runtime — the model will still generate output with the old three-token prompt — but because all v1.2 training data included four prompt tokens, generation quality will be significantly degraded without it.


Output elements now follow the document's natural reading order across all semantic classes, including Footnotes, Page-Footers, Tables, Pictures, and Captions.

Note that input image size in v1.2 changed from 2048x1648 to 2048x1664 compared to v1.1

You can use a public image *nvcr.io/nvidia/pytorch:25.03-py3* with the following library versions installed on top:

```
pip install accelerate==1.12.0
pip install albumentations==2.0.8
pip install transformers==5.6.1
pip install timm==1.0.22
```
```
import torch
from PIL import Image, ImageDraw
from transformers import AutoModel, AutoProcessor, AutoTokenizer, AutoConfig, AutoImageProcessor, GenerationConfig
from postprocessing import extract_classes_bboxes, transform_bbox_to_original, postprocess_text
# Load model and processor
model_path = "nvidia/NVIDIA-Nemotron-Parse-v1.2"  # Or use a local path
device = "cuda:0"
model = AutoModel.from_pretrained(
    model_path,
    trust_remote_code=True,
    torch_dtype=torch.bfloat16
).to(device).eval()
tokenizer = AutoTokenizer.from_pretrained(model_path)
processor = AutoProcessor.from_pretrained(model_path, trust_remote_code=True)
# Load image
image = Image.open("path/to/your/image.jpg")
task_prompt = "</s><s><predict_bbox><predict_classes><output_markdown><predict_no_text_in_pic>"
# task_prompt = "</s><s><predict_bbox><predict_classes><output_markdown><predict_text_in_pic>"
# Process image
inputs = processor(images=[image], text=task_prompt, return_tensors="pt", add_special_tokens=False).to(device)
generation_config = GenerationConfig.from_pretrained(model_path, trust_remote_code=True)
# Generate text
outputs = model.generate(**inputs,  generation_config=generation_config)
# Decode the generated text
generated_text = processor.batch_decode(outputs, skip_special_tokens=True)[0]
```
```
from PIL import Image, ImageDraw
from postprocessing import extract_classes_bboxes, transform_bbox_to_original, postprocess_text
classes, bboxes, texts = extract_classes_bboxes(generated_text)
bboxes = [transform_bbox_to_original(bbox, image.width, image.height) for bbox in bboxes]
# Specify output formats for postprocessing
table_format = 'latex' # latex | HTML | markdown | json | csv
text_format = 'markdown' # markdown | plain
blank_text_in_figures = False # remove text inside 'Picture' class
texts = [postprocess_text(text, cls = cls, table_format=table_format, text_format=text_format, blank_text_in_figures=blank_text_in_figures) for text, cls in zip(texts, classes)]
for cl, bb, txt in zip(classes, bboxes, texts):
    print(cl, ': ', txt)
draw = ImageDraw.Draw(image)
for bbox in bboxes:
  draw.rectangle((bbox[0], bbox[1], bbox[2], bbox[3]), outline="red")
```
Supported values for `table_format`: `'latex'` | `'HTML'` | `'markdown'` | `'json'` | `'csv'`

Nemotron-Parse-v1.2 is available in vllm main and can be found in vllm/vllm-openai:v0.14.1 docker image.

Note: when running on A100/A10 we recommend running vllm serve with *--attention-backend=TRITON_ATTN*

You will need to install the following dependencies on top, and then follow the VLLM Inference example below:

```
pip install albumentations timm open_clip_torch
```
```
from vllm import LLM, SamplingParams
from PIL import Image
def main():
    sampling_params = SamplingParams(
        temperature=0,
        top_k=1,
        repetition_penalty=1.1,
        max_tokens=9000,
        skip_special_tokens=False,
    )
    
    llm = LLM(
        model="nvidia/NVIDIA-Nemotron-Parse-v1.2",
        max_num_seqs=64,
        limit_mm_per_prompt={"image": 1},
        dtype="bfloat16",
        trust_remote_code=True,
    )
    
    image = Image.open("<YOUR-IMAGE-PATH>")
    
    prompts = [
        {  # Implicit prompt
            "prompt": "</s><s><predict_bbox><predict_classes><output_markdown><predict_no_text_in_pic>",
            "multi_modal_data": {
                "image": image
            },
        },
        {  # Explicit encoder/decoder prompt
            "encoder_prompt": {
                "prompt": "",
                "multi_modal_data": {
                    "image": image
                },
            },
            "decoder_prompt": "</s><s><predict_bbox><predict_classes><output_markdown><predict_no_text_in_pic>",
        },
    ]
    
    outputs = llm.generate(prompts, sampling_params)
    
    for output in outputs:
        prompt = output.prompt
        generated_text = output.outputs[0].text
        print(f"Decoder prompt: {prompt!r}, Generated text: {generated_text!r}")
if __name__ == "__main__":
    main()
```
Alternatively, you can start a vllm server as:

```
vllm serve nvidia/NVIDIA-Nemotron-Parse-v1.2 \
    --dtype bfloat16 \
    --max-num-seqs 8 \
    --limit-mm-per-prompt '{"image": 1}' \
    --trust-remote-code \
    --port 8000 \
    --chat-template chat_template.jinja
```
with *chat_template.jinja* provided in this repository. Then, you can run inference as:

