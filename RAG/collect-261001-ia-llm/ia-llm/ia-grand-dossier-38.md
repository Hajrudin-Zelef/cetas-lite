---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-38
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Cerebras", "DeepSeek", "Groq", "Hugging Face", "OpenAI", "OpenRouter", "vLLM"]
dates: []
keywords: ["claude", "deepseek", "embeddings", "fable 5", "gguf", "gpt-5.6", "gpt-6", "gpu", "luna", "opus 5", "sandbox", "sonnet 5"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [2805, 2913]
sha256: e85b43e94eeb62275fec4e383b59148641db8791edf37d3455af11c4fa8642d6
---

# IA — Le grand dossier

**Protocole d'évaluation** (sans lui, le « local » est une croyance) : 50–100 questions réelles avec réponses attendues ; mesurer **précision@k**, **taux de citation correcte**, **hallucinations** ; comparer local (Qwen3-32B) vs cloud (Sonnet/Opus) sur le même jeu. Décision chiffrée, pas idéologique.

### 7.2. Classification / extraction : le bulk silencieux

Trier des tickets, extraire des références de pièces depuis des PDF, normaliser des logs : un **8B local** à 100+ tok/s fait ça en boucle pour le prix de l'électricité. Pattern :

```python
# classify.py — Ollama local, sortie JSON stricte
import json, requests

def classify(ticket: str) -> dict:
    r = requests.post("http://localhost:11434/v1/chat/completions", json={
        "model": "qwen3:8b",
        "temperature": 0.0,
        "response_format": {"type": "json_object"},   # JSON garanti (Ollama >= 0.3)
        "messages": [
            {"role": "system", "content":
             'Classe le ticket. Réponds UNIQUEMENT en JSON : {"categorie": "electrique|reseau|clim|autre", "urgence": 1-5, "resume": "..."}'},
            {"role": "user", "content": ticket}],
    }, timeout=60)
    return json.loads(r.json()["choices"][0]["message"]["content"])
```

À 5 000 tickets/jour × ~300 tokens : **~0 € en local** vs ~30–150 $/jour en API frontier (selon modèle). Le seuil de rentabilité du GPU se calcule en semaines.

### 7.3. Code : l'assistant local du sysadmin

- **Autocomplétion** : Qwen2.5-Coder-7B/14B ou StarCoder2 via Ollama + extension (Continue, Codeium local) : correct pour bash/Python/PowerShell courants.
- **Revue** : Qwen3-32B en local pour relire un script avant prod (« trouve les bugs, propose le diff »).
- **Limite honnête** : pour du refactoring multi-fichiers complexe, les frontier (Sonnet 5, GPT-5.6 Terra, Devstral 2) restent devant — d'où le **routage par tâche** : local pour le quotidien, cloud pour le dur (via LiteLLM, une variable d'env à changer).

### 7.4. Supervision intelligente (ton terrain)

- **Résumé d'alertes** : agréger 200 alertes Zabbix en 5 lignes actionnables (8B local, cron toutes les 5 min).
- **Analyse de logs** : détection d'anomalies sur logs onduleurs/switches (patterns + LLM pour le langage naturel des messages).
- **Runbooks augmentés** : RAG local branché sur tes procédures internes → l'astreinte interroge en langage naturel (« procédure de bypass EATON 93PM ») au lieu de chercher le PDF.
- **Garde-fou** : jamais d'action automatique destructive sans validation humaine ; le LLM **propose**, l'humain **valide** (principe d'exploitation non négociable).

---

## 8. Comparatif local vs cloud : le tableau de décision

### 8.1. Le match en un tableau

| Critère | **Local** (ton GPU) | **Cloud API** (providers) | Gagnant si… |
|---|---|---|---|
| Coût à faible volume (<1M tokens/jour) | Amortissement GPU : ~5–15 $/jour | ~0,10–5 $/jour | **Cloud** (pas d'investissement) |
| Coût à fort volume (>50M tokens/jour) | ~2–5 $/jour d'élec + amortissement | 10–500 $/jour (selon modèle) | **Local** (10–100× moins cher) |
| Coût marginal | ~0 (électricité) | Chaque token se paie | **Local** dès que le GPU tourne |
| Latence interactive | 50–190 tok/s (8B, 4090/5090) ; 12–28 tok/s (70B, Mac) | 190–3 000 tok/s (DeepInfra→Cerebras) ; TTFT <1 s (Groq) | **Cloud** (sauf 8B local bien réglé) |
| Débit batch | Limité par ton GPU | Quasi infini (batch API -50 %) | **Cloud** |
| Confidentialité | Totale (air-gap possible) | Contractuelle (DPA, zero-retention payante) | **Local** (données sensibles) |
| Disponibilité | La tienne (panne = toi) | 99,9 % SLA (hyperscalers) ; variable (petits hosts) | **Cloud** (sauf site isolé → local) |
| Qualité max | 32–70B ouverts (≈ 70–85 % d'un frontier) | Frontier (GPT-6, Fable 5, Opus 5.5) | **Cloud** |
| Reproductibilité | Totale (poids épinglés) | Le provider change les poids sans prévenir | **Local** (audit, certif) |
| Maintenance | Drivers, VRAM, mises à jour, monitoring : **toi** | **Eux** (tu paies la marge) | **Cloud** (si pas d'équipe infra — mais tu en es une) |
| Démarrage | Achat + install (jours) | 5 minutes, une clé API | **Cloud** |
| Verrouillage | Aucun (poids ouverts) | Modèle délisté, prix changé, API dépréciée | **Local** |

### 8.2. Le calcul de seuil (la formule qui décide)

```
Coût_cloud_mensuel  = tokens_mois/1e6 × prix_1M_moyen
Coût_local_mensuel  = (prix_GPU / 36) + elec_mensuelle + ton_temps_h × taux_horaire
Seuil de bascule    : tokens_mois > Coût_local_mensuel / prix_1M_moyen × 1e6
```

**Exemple chiffré** (RTX 4090 à 2 000 $, 36 mois → 55 $/mois ; élec 100 W moyens → ~15 $/mois ; temps d'exploitation 4 h/mois valorisées 50 $/h → 200 $/mois ; total ~270 $/mois) :
- vs GPT-5.6 Luna (0,20 $/1M in) : seuil ≈ **1,35 Md tokens/mois** — le cloud gagne presque toujours à ce prix ;
- vs Claude Sonnet 5 (3,00 $/1M in) : seuil ≈ **90M tokens/mois** (~3M/jour) — le local devient rentable pour un RAG d'équipe actif ;
- vs usage « embeddings + 8B » (coût cloud ~0,05 $/1M) : le local ne se justifie que par la **confidentialité**, pas par le prix.

**Conclusion** : le local ne bat pas le cloud « pas cher » (Luna, DeepSeek, DeepInfra) sur le seul prix — il le bat sur **confidentialité + reproductibilité + coût marginal nul à fort volume**. D'où l'architecture hybride (§4.8, §5.4) : **local par défaut pour le sensible et le bulk, cloud pour la qualité et les pics**.

### 8.3. Arbre de décision (à afficher dans la salle d'exploitation)

```
Les données peuvent-elles quitter le site ? (contrat, secret, appel d'offres)
├─ NON  → LOCAL (Ollama/vLLM sur site, poids ouverts, air-gap si besoin)
└─ OUI  → Quel volume ?
          ├─ < 1M tokens/jour  → CLOUD direct (OpenRouter ou provider, pas de gateway)
          ├─ 1–50M/jour        → CLOUD + LiteLLM (fallbacks, budgets, routage coût)
          └─ > 50M/jour        → HYBRIDE : local (bulk, embeddings, triage)
                                 + cloud frontier (qualité, pics) via LiteLLM
              La qualité d'un 32B local suffit-elle ? (eval §7.1)
              ├─ OUI → local d'abord, escalade cloud sur confiance basse
              └─ NON → cloud d'abord, local pour le pré/post-traitement
```

---


---

## 11. Sécurité et conformité de l'IA locale

Le local n'est pas « sûr par magie » : c'est un serveur comme un autre, avec des risques spécifiques.

### 11.1. Chaîne de confiance des poids

| Risque | Réalité | Contre-mesure |
|---|---|---|
| Poids malveillants | Un GGUF peut embarquer du code (désérialisation, templates) | Télécharger uniquement depuis les **comptes officiels** Hugging Face (organisations vérifiées) ; vérifier les hash SHA quand publiés |
| Modèle « backdooré » | Comportement caché (exfiltration via tool calling) | Tester en sandbox réseau fermé avant prod ; auditer les templates de chat |
| Licence piégée | Variante « MIT » qui n'en est pas une sur un fork | Vérifier la licence sur le **dépôt d'origine**, pas sur le fork |
| Fuite par les logs | Prompts sensibles dans les logs du serveur | Désactiver le log du contenu ; chiffrer les volumes ; rétention courte |

### 11.2. Conformité (UE/FR)

