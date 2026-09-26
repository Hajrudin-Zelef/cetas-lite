---
id: collect-240926-huggingface/huggingface/qwen-qwen3-vl-embedding-8b-hugging-face-3
title: "Load the model"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "SGLang", "vLLM"]
dates: []
keywords: ["embedding", "embeddings", "multimodal", "qwen", "reranker", "sglang", "training", "vllm"]
source: docs/RAG/clean_en/huggingface/qwen-qwen3-vl-embedding-8b-hugging-face.md
source_anchor: ""
source_lines: [207, 403]
sha256: 5139ee39fd54b888bb5e59bf7db0184c6638d252eadffd7bee51539fa53780fe
---

# Load the model

    return conversation
def prepare_vllm_inputs(input_dict: Dict[str, Any], llm, instruction: str = "Represent the user's input.") -> Dict[str, Any]:
    text = input_dict.get('text')
    image = input_dict.get('image')
    
    conversation = format_input_to_conversation(input_dict, instruction)
    
    prompt_text = llm.llm_engine.tokenizer.apply_chat_template(
        conversation, 
        tokenize=False, 
        add_generation_prompt=True
    )
    
    multi_modal_data = None
    if image:
        if isinstance(image, str):
            if image.startswith(('http', 'https', 'oss')):
                try:
                    image_obj = fetch_image(image)
                    multi_modal_data = {"image": image_obj}
                except Exception as e:
                    print(f"Warning: Failed to fetch image {image}: {e}")
            else:
                abs_image_path = os.path.abspath(image)
                if os.path.exists(abs_image_path):
                    from PIL import Image
                    image_obj = Image.open(abs_image_path)
                    multi_modal_data = {"image": image_obj}
                else:
                    print(f"Warning: Image file not found: {abs_image_path}")
        else:
            multi_modal_data = {"image": image}
    
    result = {
        "prompt": prompt_text,
        "multi_modal_data": multi_modal_data
    }
    return result
def main():
    parser = argparse.ArgumentParser(description="Offline Similarity Check with vLLM")
    parser.add_argument("--model-path", type=str, default="models/Qwen3-VL-Embedding-8B", help="Path to the model")
    parser.add_argument("--dtype", type=str, default="bfloat16", help="Data type (e.g., bfloat16)")
    args = parser.parse_args()
    print(f"Loading model from {args.model_path}...")
    
    engine_args = EngineArgs(
        model=args.model_path,
        runner="pooling",
        dtype=args.dtype,
        trust_remote_code=True,
    )
    
    llm = LLM(**vars(engine_args))
    
    all_inputs = queries + documents
    vllm_inputs = [prepare_vllm_inputs(inp, llm) for inp in all_inputs]
    
    
    outputs = llm.embed(vllm_inputs)
    
    embeddings_list = []
    for i, output in enumerate(outputs):
        emb = output.outputs.embedding
        embeddings_list.append(emb)
        print(f"Input {i} embedding shape: {len(emb)}")
    
    embeddings = np.array(embeddings_list)
    print(f"\nEmbeddings shape: {embeddings.shape}")
    
    num_queries = len(queries)
    query_embeddings = embeddings[:num_queries]
    doc_embeddings = embeddings[num_queries:]
    
    similarity_scores = query_embeddings @ doc_embeddings.T
    
    print("\nSimilarity Scores:")
    print(similarity_scores.tolist())
    
if __name__ == "__main__":
    main()
```
```
import argparse
import numpy as np
import torch
import os
from typing import List, Dict, Any
from sglang.srt.entrypoints.engine import Engine
# Define a list of query texts
queries = [
    {"text": "A woman playing with her dog on a beach at sunset."},
    {"text": "Pet owner training dog outdoors near water."},
    {"text": "Woman surfing on waves during a sunny day."},
    {"text": "City skyline view from a high-rise building at night."}
]
# Define a list of document texts and images
documents = [
    {"text": "A woman shares a joyful moment with her golden retriever on a sun-drenched beach at sunset, as the dog offers its paw in a heartwarming display of companionship and trust."},
    {"image": "https://qianwen-res.oss-cn-beijing.aliyuncs.com/Qwen-VL/assets/demo.jpeg"},
    {"text": "A woman shares a joyful moment with her golden retriever on a sun-drenched beach at sunset, as the dog offers its paw in a heartwarming display of companionship and trust.", "image": "https://qianwen-res.oss-cn-beijing.aliyuncs.com/Qwen-VL/assets/demo.jpeg"}
]
def format_input_to_conversation(input_dict: Dict[str, Any], instruction: str = "Represent the user's input.") -> List[Dict]:
    content = []
    
    text = input_dict.get('text')
    image = input_dict.get('image')
    if image:
        image_content = None
        if isinstance(image, str):
            if image.startswith(('http', 'oss')):
                image_content = image
            else:
                abs_image_path = os.path.abspath(image)
                image_content = 'file://' + abs_image_path
        else:
            image_content = image
        if image_content:
            content.append({
                'type': 'image', 'image': image_content,
            })
    if text:
        content.append({'type': 'text', 'text': text})
    if not content:
        content.append({'type': 'text', 'text': ""})
    conversation = [
        {"role": "system", "content": [{"type": "text", "text": instruction}]},
        {"role": "user", "content": content}
    ]
    return conversation
def convert_to_sglang_format(input_dict: Dict[str, Any], engine: Engine, instruction: str = "Represent the user's input.") -> Dict[str, Any]:
    conversation = format_input_to_conversation(input_dict, instruction)
    
    text_for_api = engine.tokenizer_manager.tokenizer.apply_chat_template(
        conversation, 
        tokenize=False, 
        add_generation_prompt=True
    )
    result = {"text": text_for_api}
    
    image = input_dict.get('image')
    if image and isinstance(image, str):
        result["image"] = image
        
        
    return result
def main():
    parser = argparse.ArgumentParser(description="Offline Similarity Check with SGLang")
    parser.add_argument("--model-path", type=str, default="models/Qwen3-VL-Embedding-8B", help="Path to the model")
    parser.add_argument("--dtype", type=str, default="bfloat16", help="Data type (e.g., bfloat16)")
    args = parser.parse_args()
    print(f"Loading model from {args.model_path}...")
    
    engine = Engine(
        model_path=args.model_path,
        is_embedding=True,
        dtype=args.dtype,
        trust_remote_code=True,
    )
    inputs = queries + documents
    sglang_inputs = [convert_to_sglang_format(inp, engine) for inp in inputs]
    print(sglang_inputs[:])
    print(f"sglang_inputs: {sglang_inputs}")
    print(f"Processing {len(sglang_inputs)} inputs...")
    prompts = [inp['text'] for inp in sglang_inputs]
    images = [inp.get('image') for inp in sglang_inputs]
    results = engine.encode(prompts, image_data=images)
    
    embeddings_list = []
    for res in results:
        embeddings_list.append(res['embedding'])
            
    embeddings = np.array(embeddings_list)
    print(f"Embeddings shape: {embeddings.shape}")
    num_queries = len(queries)
    query_embeddings = embeddings[:num_queries]
    doc_embeddings = embeddings[num_queries:]
    
    similarity_scores = (query_embeddings @ doc_embeddings.T)
    print("\nSimilarity Scores:")
    print(similarity_scores.tolist())
if __name__ == "__main__":
    main()
```
For more usage examples, please visit our GitHub repository.

If you find our work helpful, feel free to give us a cite.

```
@article{qwen3vlembedding,
  title={Qwen3-VL-Embedding and Qwen3-VL-Reranker: A Unified Framework for State-of-the-Art Multimodal Retrieval and Ranking},
  author={Li, Mingxin and Zhang, Yanzhao and Long, Dingkun and Chen Keqin and Song, Sibo and Bai, Shuai and Yang, Zhibo and Xie, Pengjun and Yang, An and Liu, Dayiheng and Zhou, Jingren and Lin, Junyang},
  journal={arXiv preprint arXiv:2601.04720},
  year={2026}
}
```
- Downloads last month
- 1,215,903
