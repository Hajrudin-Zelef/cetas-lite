---
id: collect-261001-ia-llm/ia-llm/training-and-finetuning-multimodal-embedding-reranker-models-with-sentence-transformers-1
title: "['text', 'image', 'video', 'message']"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba"]
dates: []
keywords: ["attention", "embedding", "embeddings", "gemini", "memory", "multimodal", "parameters", "qwen", "reranker", "revenue", "training"]
source: docs/RAG/collect-261001-ia-llm/training-and-finetuning-multimodal-embedding-reranker-models-with-sentence-transformers.md
source_anchor: ""
source_lines: [1, 82]
sha256: 2cf3596bdaae8f0a448d05f2f9346a8169b7bae8a70bbe5a191e0fc090f1d532
---

# ['text', 'image', 'video', 'message']

**train or finetune**these multimodal models on your own data.

As a practical example, I'll walk through finetuning `Qwen/Qwen3-VL-Embedding-2B` for Visual Document Retrieval (VDR), the task of retrieving relevant document pages (as images, with charts, tables, and layout intact) for a given text query. The resulting `tomaarsen/Qwen3-VL-Embedding-2B-vdr` demonstrates how much performance you can gain by finetuning on your own domain. On my evaluation data, the finetuned model achieves an NDCG@10 of 0.947 compared to the base model's 0.888, and outperforms all existing VDR models I tested against, including models up to 4x its size.

If you're new to multimodal models in Sentence Transformers, I recommend reading Multimodal Embedding & Reranker Models with Sentence Transformers first. For training text-only embedding, reranker, or sparse embedding models, see the Prior Blogposts section at the end.


General-purpose multimodal embedding models like `Qwen/Qwen3-VL-Embedding-2B` are trained on diverse data to perform well across a wide range of languages and tasks: image-text matching, visual question answering, document understanding, and more. But this generality means the model is rarely the best choice for any specific task.

Consider Visual Document Retrieval: given a text query like "What was the company's Q3 revenue?", the model must find the most relevant document screenshot from a corpus of thousands. This requires understanding document layouts, charts, tables, and text, which is a very different skill from e.g. matching pictures of shoes with product descriptions.

By finetuning on domain-specific data, the model can learn these specialized patterns. In my experiment, finetuning improved NDCG@10 from 0.888 to 0.947, ahead of every recent multimodal model I tested, including ones up to 4x larger.

Training multimodal Sentence Transformer models involves the same components as training text-only models:

1. **Model** : The multimodal model to train or finetune.
2. **Dataset** : The data used for training and evaluation.
3. **Loss Function** : A function that quantifies the model's performance and guides the optimization process.
4. **Training Arguments** (optional): Parameters that influence training performance and tracking/debugging.
5. **Evaluator** (optional): A tool for evaluating the model before, during, or after training.
6. **Trainer** : Brings together the model, dataset, loss function, and other components for training.

The multimodal training pipeline uses the same `SentenceTransformerTrainer` as text-only training. The key difference is that your datasets contain images (or other modalities) alongside text, and the model's processor handles the image preprocessing automatically.

Let's walk through each component, using Visual Document Retrieval (matching text queries to document screenshots) as a running example.

The most common approach is to finetune an existing multimodal embedding model, or to start from a Vision-Language Model (VLM) checkpoint. The `Transformer` module automatically detects supported modalities from the model's processor.

To finetune an existing multimodal embedding model (e.g. one that already has a `modules.json` file), you can pass `processor_kwargs` and `model_kwargs` to control preprocessing and model loading respectively. `processor_kwargs` are passed directly to `AutoProcessor.from_pretrained(...)` (e.g., image resolution bounds: higher `max_pixels` means higher quality but more memory), while `model_kwargs` are passed to the appropriate `AutoModel.from_pretrained(...)` call (e.g., precision, attention implementation):

```
from sentence_transformers import SentenceTransformer
model = SentenceTransformer(
    "Qwen/Qwen3-VL-Embedding-2B",
    model_kwargs={"attn_implementation": "flash_attention_2", "torch_dtype": "bfloat16"},
    processor_kwargs={"min_pixels": 28 * 28, "max_pixels": 600 * 600},
)
```
You can also start from a fresh VLM checkpoint that hasn't been trained for embeddings yet. Sentence Transformers will attempt to recognize the architecture, infer the supported modalities from the processor, and set up the appropriate forward method and pooling. If the automatic detection doesn't work perfectly for a particular model, the configuration in the saved `sentence_bert_config.json` can be edited to adjust modality settings, forward methods, and output handling:

```
from sentence_transformers import SentenceTransformer
model = SentenceTransformer("Qwen/Qwen3-VL-2B")
```
In both cases, the `Transformer` module inspects the processor to determine which modalities are available, and `Pooling` is added automatically if needed. You can verify the supported modalities:

```
print(model.modalities)
# ['text', 'image', 'video', 'message']
print(model.supports("image"))
# True
```
## Alternative: Building multimodal models with Router

Instead of using a single VLM backbone, you can compose separate encoders for different modalities using the `Router` module. This lets you combine any existing encoders and route inputs to the appropriate one based on detected modality:

```
from sentence_transformers import SentenceTransformer
from sentence_transformers.sentence_transformer.modules import Dense, Pooling, Router, Transformer
# Create separate encoders for different modalities
text_encoder = Transformer("sentence-transformers/all-MiniLM-L6-v2")
text_pooling = Pooling(text_encoder.get_embedding_dimension(), pooling_mode="mean")
text_projection = Dense(text_encoder.get_embedding_dimension(), 768)
# SigLIP outputs pooled embeddings directly, so no separate Pooling module is needed
image_encoder = Transformer("google/siglip2-base-patch16-224")
# Route inputs based on modality
router = Router(
    sub_modules={
        "text": [text_encoder, text_pooling, text_projection],
        "image": [image_encoder],
    },
)
model = SentenceTransformer(modules=[router])
```
Since Router-based multimodal models use separate encoders per modality, their embedding spaces are initially unaligned. Training is required to align the spaces for meaningful cross-modal similarity. The `Dense` projection layer shown above helps map embeddings from different encoders into a shared space.


This approach is useful when you want to use lightweight, specialized encoders rather than a large VLM. You can also combine Router-based multimodality with task-based routing (e.g. different encoders for queries vs. documents) using `route_mappings`. See the `Router` documentation for advanced routing scenarios.

For this example, I use the `tomaarsen/llamaindex-vdr-en-train-preprocessed` dataset, a preprocessed English subset of `llamaindex/vdr-multilingual-train`. The source dataset was released alongside the Visual Document Retrieval Goes Multilingual blogpost by LlamaIndex, and consists of ~500k multilingual query-image samples collected from public internet PDFs, with queries synthetically generated using VLMs (gemini-1.5-pro and Qwen2-VL-72B). 
My preprocessed version filters to the 53,512 English samples and resolves 4 of the 16 ID-based hard negatives per sample into actual document screenshot images, so it can be used directly for training without further preprocessing:

