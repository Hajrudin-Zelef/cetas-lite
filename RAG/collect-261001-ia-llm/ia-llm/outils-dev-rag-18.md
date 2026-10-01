---
id: collect-261001-ia-llm/ia-llm/outils-dev-rag-18
title: "Outils dev + ingénierie RAG (chunk & corpus)"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["embedding", "reasoning"]
source: docs/RAG/collect-261001-ia-llm/outils_dev_rag.md
source_anchor: ""
source_lines: [3351, 3558]
sha256: 1423b466e817aea668b5e445b0f436bf68b2fec97f7bb966c8fb43ca50265df6
---

# Outils dev + ingénierie RAG (chunk & corpus)

```python
FEWSHOT_RAG = """\
Exemple 1 :
Contexte :
[1] (source: huawei_cli) La commande display interface brief affiche
l'état résumé des interfaces, y compris l'état physique et protocolaire.
[2] (source: guides) ...
Question : Quelle commande affiche l'état des interfaces ?
Réponse : La commande `display interface brief` affiche l'état résumé des
interfaces, incluant l'état physique et protocolaire [1].
Sources : [1] huawei_cli — display interface

Exemple 2 :
Contexte :
[1] (source: guides) L'AP361 consomme 8,8 W via PoE 802.3af.
Question : Quelle est la tension d'alimentation d'un onduleur triphasé ?
Réponse : Je n'ai pas trouvé cette information dans les documents fournis.
"""
```

**Sélection dynamique** (pratique 2026 vérifiée) : avec une banque
d'exemples, choisir par similarité sémantique les 2-3 exemples les plus
proches de la question courante — plus efficace que 5 exemples fixes.

## 115. Chain-of-thought : principe (vérifié)

**Chain-of-Thought (CoT)** : demander au modèle de **raisonner étape par
étape avant de répondre**. Le raisonnement intermédiaire améliore la
précision sur les problèmes multi-étapes (diagnostic, calculs,
procédures).

Le déclencheur **zero-shot CoT** (sans exemples) :
```text
Résous le problème étape par étape.
```
ou l'équivalent anglais « Let's think step by step » (vérifié : fonctionne
dans les deux langues).

**Quand l'utiliser dans ton RAG :**
- Diagnostic (« mon AP ne s'associe pas : quelle séquence de vérification ? »).
- Comparaisons multi-critères (« VFI vs VI pour 40 kVA : quel choix ? »).
- Synthèse de plusieurs chunks contradictoires.

**Quand NE PAS l'utiliser :**
- Question factuelle simple (« quelle commande ? ») → gaspillage de tokens
  (vérifié : CoT = 3 à 6x plus de tokens de sortie).
- Réponse temps-réel critique → latence.

## 116. Variantes de CoT (bref, vérifié 2026)

| Variante | Principe | Usage |
|---|---|---|
| **Zero-shot CoT** | Ajouter « raisonne étape par étape » | raisonnement simple, sans exemples |
| **Few-shot CoT** | Exemples AVEC traces de raisonnement | problèmes multi-étapes complexes |
| **Self-consistency** | Plusieurs chemins de raisonnement, vote majoritaire | décisions à fort enjeu |
| **Tree-of-thoughts** | Explorer plusieurs branches, élaguer | planification complexe |
| **Least-to-most** | Décomposer en sous-problèmes, résoudre séquentiellement | problèmes composés |

**Self-consistency en pratique** : échantillonner 3-5 réponses avec une
température plus haute, prendre la réponse majoritaire. Coût : 5x les
tokens — réservé aux cas où une erreur coûte cher (dimensionnement,
procédure de consignation).

**Note 2026 (vérifié)** : avec les modèles « à raisonnement intégré », le
CoT explicite apporte un gain marginal (+2-3 %) pour +20-80 % de latence.
Pour ton RAG sur modèle standard, le CoT reste pertinent sur les
questions de diagnostic.

```python
# Self-consistency simplifiée (pseudo-code)
reponses = [llm(question, temperature=0.7) for _ in range(5)]
reponse_finale = max(set(reponses), key=reponses.count)  # vote majoritaire
```

## 117. Coûts en tokens des techniques de raisonnement (vérifié 2026)

Ordres de grandeur (source : guide de production 2026) :

```text
Q&A factuel simple :      ~50  tokens de sortie
Zero-shot CoT :           ~150-300 tokens  (3-6x)
CoT structuré :           ~300-500 tokens  (6-10x)
Self-consistency (5x) :   ~750-1500 tokens (15-30x)
```

Règle : n'utiliser le raisonnement que si le gain de précision justifie
le coût. Pour ton RAG personnel, le CoT par défaut sur TOUTES les
questions est un gaspillage — réserve-le aux diagnostics.

## 118. Prompts RAG : template avec contexte récupéré + citations

Le template canonique (vérifié dans plusieurs implémentations 2026) :

```python
RAG_PROMPT = """\
Réponds à la question en te basant UNIQUEMENT sur le contexte ci-dessous.
Inclus des citations avec [1], [2], etc.

Si tu ne peux pas répondre à partir du contexte, dis exactement :
« Je n'ai pas trouvé cette information dans les documents fournis. »

Contexte :
{context}

Question : {question}

Instructions :
1. N'utilise que les informations du contexte.
2. Cite tes sources au format [1], [2].
3. En cas d'incertitude, exprime-la explicitement.

Réponse (avec citations) :"""
```

**Construction du bloc contexte** — les passages numérotés donnent au
modèle quelque chose de concret à citer (vérifié) :

```python
def build_context(chunks: list[dict]) -> str:
    """chunks = résultats de match_chunks (§67) : {content, source, section}."""
    blocks = []
    for i, c in enumerate(chunks, start=1):
        src = c.get("source", "?")
        section = c.get("section", "")
        blocks.append(f"[{i}] (source: {src} — {section})\n{c['content']}")
    return "\n\n".join(blocks)
```

> **Peu de chunks, mais bons** : 3-5 chunks bien classés + consigne de
> grounding > 20 chunks noyés. C'est le correctif n°1 des hallucinations
> RAG (vérifié).

## 119. Gestion du « je ne sais pas » : l'abstention qui sauve

**Le pire échec RAG** : une réponse inventée avec assurance quand le
contexte est vide — elle ressemble exactement à une bonne réponse.
Contre-mesures en 3 couches (toutes vérifiées 2026) :

```python
# COUCHE 1 — garde-fou AVANT l'appel LLM (le moins cher) :
hits = search(question_embedding, k=10)
if not hits or max(h["similarite"] for h in hits) < SEUIL:
    return "Je n'ai pas trouvé cette information dans les documents fournis."
# → tu ne paies même pas l'appel, zéro risque d'hallucination.

# COUCHE 2 — consigne explicite dans le prompt (§118) :
# « Si tu ne peux pas répondre à partir du contexte, dis exactement : … »

# COUCHE 3 — phrase d'abstention SANCTIONNÉE :
# Le modèle a le droit de dire « je ne sais pas » — un RAG qui dit
# parfois « je ne sais pas » fonctionne ; un RAG qui ne le dit jamais
# hallucine (vérifié).
```

**Formulation exacte et stable** : impose UNE phrase d'abstention unique
(« Je n'ai pas trouvé cette information dans les documents fournis. »)
pour pouvoir la détecter en aval (logs, métriques de couverture).

## 120. Trois templates RAG complets (à copier)

**A. Strict (défaut — documentation technique) :**
```text
[System] Tu es un assistant documentation réseau. Tu réponds UNIQUEMENT
d'après le CONTEXTE. Chaque affirmation est citée [N]. Les chiffres sont
recopiés à l'identique. Si l'info manque : « Je n'ai pas trouvé cette
information dans les documents fournis. » Réponds en français.

[User]
CONTEXTE :
{context}

QUESTION : {question}
```

**B. Hybride (contexte + connaissances générales balisées) :**
```text
[System] Réponds d'abord d'après le CONTEXTE (cité [N]). Tu peux compléter
avec tes connaissances générales UNIQUEMENT pour le contexte pédagogique,
en le signalant explicitement par « (connaissances générales) ». Pour tout
fait technique précis (syntaxe, valeur, version) : contexte ou abstention.

[User]
CONTEXTE :
{context}

QUESTION : {question}
```
> À utiliser avec prudence : le balisage est une convention, pas une
> garantie. Le template A reste le défaut pour la doc technique.

**C. Structuré JSON (pour chaînage logiciel) :**
```python
# Schéma Pydantic vérifié (pattern 2026) :
from pydantic import BaseModel, Field

class RAGResponse(BaseModel):
    answer: str = Field(description="Réponse basée sur le contexte")
    confidence: float = Field(ge=0, le=1, description="Confiance 0-1")
    sources: list[str] = Field(description="IDs des documents utilisés")
    reasoning: str = Field(description="Bref raisonnement")

# Appel : llm.with_structured_output(RAGResponse)
# → réponse typée, parsable, validée. Toujours gérer l'échec de parsing
# (try/except + repli vers extraction non structurée).
```

## 121. Anti-patterns : les erreurs classiques (vérifié 2026)

