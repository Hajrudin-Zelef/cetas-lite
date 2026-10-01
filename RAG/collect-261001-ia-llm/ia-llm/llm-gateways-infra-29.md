---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-29
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["Google", "Huawei", "OpenAI", "United States"]
dates: ["2026-09-27"]
keywords: ["attention", "dpo", "embedding", "embeddings", "gemini", "research", "transcription"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [3382, 3468]
sha256: ddece4e81059161e019796dd322a763831d8e6091fa6ef9f6fa727ac2962ecc2
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

- **Pas d'API publique officielle** : tout est web UI. Il existe des libs tierces non officielles (notebooklm-py, open-notebooklm) qui pilotent l'interface — fragiles par construction (un changement de RPC côté Google et ça casse), à proscrire en prod.
- **Zéro contrôle sur le pipeline** : tu ne choisis ni le chunking (section 140), ni le modèle d'embedding (tu es sur text-embedding-3-small, lui utilise l'embedding de Gemini), ni le reranking, ni la stratégie de récupération. Quand la qualité baisse, tu ne peux pas debugger — tu subis.
- **Tes données quittent ton infra** : cloud Google, juridiction US. Pour des docs d'exploitation internes, c'est un non-starter sans accord DPO.
- **Qualité = qualité des sources** : « answers depend on what you upload » — un corpus mal curé donne des réponses mal ancrées, exactement comme ton RAG. L'outil ne remplace pas la curation.
- **Quotas imprévisibles** : le budget de calcul rend les gros batchs (réindexer 500 docs, générer 200 quiz) aléatoires en gratuit.
- **Enfermement** : pas d'export structuré propre du graphe de connaissances ; tes carnets vivent chez Google.

**Doctrine** : NotebookLM = labo d'idées et étalon qualité, jamais une brique de ton pipeline. Ton RAG reste chez toi : tes embeddings, ton chunking, tes modèles, tes coûts maîtrisés (sections 140–141, 167).

### 166.5. Workflow concret : ton premier carnet « docs onduleurs » en 15 minutes

```text
1. notebook.google → Nouveau carnet → nomme-le « Onduleurs – base documentaire »
2. Ajouter des sources : glisse 5–10 PDF (manuels Easy UPS 3S, ton guide onduleurs_ups_guide.md
   exporté en PDF, 2-3 fiches constructeurs). Formats acceptés : PDF, Google Docs/Slides,
   TXT/MD, liens web, vidéos YouTube (transcription), audio.
3. Attends l'indexation (quelques minutes). Vérifie le compteur de sources.
4. Pose tes 5 questions les plus fréquentes du support, ex. :
   « Quelle section décrit la procédure de bypass de maintenance ? »
5. Pour chaque réponse : CLIQUE les citations. Si la citation ne contient pas vraiment
   la réponse → note-le : c'est exactement le test « fidélité » à appliquer à ton RAG.
6. Génère un Audio Overview (2–3 min) : écoute ce que « ton corpus raconté » donne —
   révélateur des trous et redondances du corpus.
```

Ce que tu en tires pour ton RAG : la liste des questions où NotebookLM échoue te dit où TON pipeline doit être meilleur (ou au moins aussi bon) ; les citations te donnent le format de réponse cible ; l'audio te fait « entendre » la qualité de ton corpus.

### 166.6. Les features 2026 qui comptent vraiment pour un tech

- **Exécution de code** : chaque carnet a un environnement Python cloud. Cas d'usage : « à partir de ces 3 rapports de maintenance en PDF, sors-moi un tableau des pannes par mois » — sans exporter les fichiers. Pour toi : prototype d'analyse documentaire avant d'automatiser dans ton pipeline.
- **Data Tables** : extraction structurée multi-sources (« compare les caractéristiques des Easy UPS 3S/3M, Eaton 9PX et Vertiv Liebert à partir de ces fiches ») → tableau exportable. C'est du RAG + structuration en un clic.
- **Deep Research** : le carnet peut chercher ET croiser tes sources avec le web — utile pour une veille (« qu'est-ce qui a changé sur la norme IEC 62040 depuis mon manuel de 2023 ? »), mais attention : dès qu'il sort de tes sources, la garantie d'ancrage tombe.
- **Quiz / flashcards** : génère des QCM depuis tes guides — recyclable pour former ton équipe (tes 72 sections Kyocera → 50 questions d'habilitation, par exemple).

### 166.7. Check-list gouvernance avant d'y mettre un doc d'entreprise

1. **Classification** : public / interne / confidentiel — seul le « public » et l'« interne non sensible » vont dans un compte Google standard. Le confidentiel : Workspace avec garanties contractuelles, ou rien.
2. **Compte dédié** : n'utilise pas ton compte perso — un compte de service/équipe, avec 2FA, dont l'accès est révocable.
3. **Pas de données personnelles** : les rapports d'intervention contiennent des noms, des sites clients — anonymise avant upload (c'est aussi une bonne hygiène pour ton RAG).
4. **Traçabilité** : note quel corpus est dans quel carnet (un simple `CORPUS.md`) — dans 6 mois, tu dois savoir d'où vient une réponse que tu as citée.
5. **Sortie** : tout ce qui est produit dans NotebookLM (notes, rapports) et réutilisé ailleurs doit être re-vérifié — l'outil ne signe pas ses affirmations, les citations si.

### 166.8. NotebookLM vs ton RAG maison : le comparatif honnête

| Critère | NotebookLM (Gemini Notebook) | Ton RAG (pipeline maison) |
|---|---|---|
| Mise en route | 15 min, zéro code | Jours/semaines de dev |
| Citations | Inline, cliquables, extraits — excellent | À construire (ton chantier) |
| Contrôle chunking/embedding | Aucun | Total (sections 140–141) |
| Choix du LLM | Gemini imposé | Ton routeur (LiteLLM) |
| API programmable | Non (libs tierces fragiles) | Oui, natif |
| Données | Cloud Google | Chez toi (on-prem possible) |
| Coût | Gratuit (quotas) | Tokens + infra, prévisible |
| Quotas | Budget de calcul, imprévisible en pic | Tes limites à toi |
| Audit/debug | Boîte noire | Logs, métriques, reproductible |

Lecture : NotebookLM gagne sur le time-to-value et l'UX des citations ; ton RAG gagne sur tout le reste dès que c'est sérieux (prod, confidentialité, automatisation). Utilise le premier pour **apprendre** (UX cible, test de corpus), le second pour **servir**.

### 166.9. Pour ton équipe : former avec les quiz générés

Cas d'usage sous-estimé pour un chef de service : tes guides (Kyocera, onduleurs, Huawei — des milliers de lignes) → carnet NotebookLM par thème → génération de **quiz/flashcards** → 15 min de formation hebdo pour l'équipe, notée. Avantages : questions ancrées dans TES docs (pas de QCM générique), génération en 2 minutes, révision facile quand une procédure change. Limite : relis chaque quiz généré — le modèle peut produire des questions ambiguës ou des réponses discutables sur des points techniques fins (couple de serrage, seuils d'alarme). Le quiz est un brouillon, pas une habilitation.

---

## 167. Token / tokenisation : l'unité qui drive ton coût et ta fenêtre de contexte

**En une phrase :** tout ce que tu paies et tout ce qui tient en mémoire se compte en tokens, pas en mots ni en caractères — et le français coûte plus cher que l'anglais à contenu égal. Cette section te donne les algorithmes, les outils et les vrais chiffres.

### 167.1. Ce qu'est un token

Un token = l'unité atomique que le modèle lit et génère. Ce n'est ni un mot, ni un caractère : c'est un **morceau de mot** appris statistiquement. Exemples mesurés le 27/09/2026 avec `tiktoken` (encodeur `cl100k_base`, celui de GPT-3.5/4) :

```python
import tiktoken
enc = tiktoken.get_encoding("cl100k_base")

enc.encode("Hello world")          # [9906, 1917]  → 2 tokens
enc.encode("Le rapport de maintenance du groupe électrogène doit être archivé avant vendredi.")
# → 18 tokens (81 caractères, soit 4,5 car/token)
enc.encode("The generator maintenance report must be archived before Friday.")
# → 10 tokens (64 caractères, soit 6,4 car/token)
```

Même sens, **18 tokens en français contre 10 en anglais** : à contenu égal, ton corpus français coûte ~1,5 à 1,8x plus cher en tokens que l'équivalent anglais sur un tokenizer entraîné majoritairement sur de l'anglais. C'est mécanique, pas un complot : le vocabulaire contient des mots anglais entiers (« maintenance » = 1 token) mais découpe le français en sous-morceaux.

Vocabulaires mesurés : `cl100k_base` = **100 277** tokens ; `o200k_base` (nouvelle génération OpenAI) = **200 019** tokens. Plus de vocabulaire = mots rares mieux couverts = moins de fragmentation.

