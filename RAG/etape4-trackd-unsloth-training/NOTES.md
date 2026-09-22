# NOTES — corpus `etape4-trackd-unsloth-training`

## Génération

Mode auto (`RAG/_tools/build_rag.py`) avec `first_is_content: true` (H1 unique = contenu),
cibles **45–90 lignes**. Dossier `00-unsloth-training` (`folder_name`), fichiers par titre H2.

Source sans ancre HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

« Step 4 — Track D : Unsloth + outillage d'entraînement / fine-tuning (2026) ».

- Unsloth (société, financement, produits, gains vitesse/VRAM annoncés vs indépendants,
  nouveautés 2026, licence) ;
- Axolotl (v0.19.0, timeline 2026) ;
- Llama-Factory, TorchTune (PyTorch/Meta), TRL (Hugging Face), écosystème PEFT/LoRA ;
- frameworks distribués : DeepSpeed, FSDP2, Megatron-LM ;
- extras d'inférence : TensorRT-LLM, ONNX Runtime GenAI, llama-server, ExLlamaV3 + TabbyAPI ;
- financement / M&A de l'outillage, journal d'incertitudes, métadonnées de collecte.

Labels de provenance (`[official]`, `[vendor-reported]`, `[independent]`, `[secondary]`,
`[unverified]`) conservés verbatim.

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
