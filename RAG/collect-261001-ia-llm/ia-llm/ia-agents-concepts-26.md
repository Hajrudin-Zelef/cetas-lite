---
id: collect-261001-ia-llm/ia-llm/ia-agents-concepts-26
title: "Concepts : agents IA, agentic, autonomie"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "OpenAI", "OpenRouter"]
dates: ["2026-09-27"]
keywords: ["agent", "agents", "benchmarks", "claude", "fine-tuning", "gpt-5.6", "multimodal", "pricing", "valuation"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_concepts.md
source_anchor: ""
source_lines: [3564, 3713]
sha256: 3b8ca3546247c470572541665ad281ba055ea958f0549cc40ce9f782138bf803
---

# Concepts : agents IA, agentic, autonomie

- Endpoint : `POST https://api.typesafe.ai/v1/systemone`, auth
  `Authorization: Bearer <API_KEY>` (clé depuis `console.typesafe.ai`).
  Corps : `state` + `model` (`"jev-latest"`) + `questions` (dict clé → type
  + `instructions` + `criteria`).
- Exemple commenté (repris du quickstart officiel, adapté au routage
  d'agent) :
```bash
# Pas de credential réel ici : export TYPESAFE_API_KEY="..." au préalable.
curl -X POST https://api.typesafe.ai/v1/systemone \
  -H "Authorization: Bearer $TYPESAFE_API_KEY" \
  -H "Content-Type: application/json" \
  -d @- << 'EOF'
{
  "state": "Agent: je dois redémarrer le service web sur srv-web-03. Plan : systemctl restart nginx.",
  "model": "jev-latest",
  "questions": {
    "risk_level": {
      "type": "choice",
      "instructions": "Quel est le niveau de risque de cette action planifiée par l'agent ?",
      "criteria": {
        "low": "Lecture seule ou action réversible, impact local",
        "medium": "Action avec effet de bord, réversible avec effort",
        "high": "Action destructive ou irréversible, impact production"
      }
    },
    "is_safe_to_execute": {
      "type": "noul",
      "instructions": "L'action planifiée peut être exécutée sans validation humaine"
    },
    "blast_radius": {
      "type": "score",
      "instructions": "Ampleur de l'impact si l'action tourne mal",
      "criteria": ["Aucun impact", "Un service", "Plusieurs services", "Production entière"]
    }
  }
}
EOF
```
- Réponse (forme officielle) :
```json
{
  "model": "jev-1.13.0",
  "answers": {
    "risk_level": {
      "type": "choice",
      "choice": "medium",
      "confidence": 0.78,
      "probabilities": { "low": 0.12, "medium": 0.78, "high": 0.10 }
    },
    "is_safe_to_execute": { "type": "noul", "noul": 0.34 },
    "blast_radius": {
      "type": "score", "score": 1.0, "confidence": 0.9,
      "legend": { "0": "Aucun impact", "1": "Un service", "2": "Plusieurs services", "3": "Production entière" },
      "probabilities": { "0": 0.05, "1": 0.85, "2": 0.08, "3": 0.02 }
    }
  },
  "usage": { "input_tokens": 392, "output_tokens": 65 }
}
```
- SDK Python officiel : `pip install typesafe-sdk` (Python ≥ 3.10), client
  `TypeSafeClient()` qui lit `TYPESAFE_API_KEY` dans l'environnement,
  classes `Choice`, `Score`, `Noul`, méthode `client.system_one(state=...,
  questions={...})`. Skill agent officielle pour Claude Code :
  `npx skills add typesafe-ai/skills --skill typesafe-ai` (ou plugin
  `typesafe@typesafe-ai`).
- **Routes tierces vérifiées** : **Vercel AI Gateway** (`typesafe-ai/jev` —
  lancement le plus vite adopté de l'histoire du gateway : ~13 % des paid
  teams en 24 h, 2× la famille GPT-5.6), **Cloudflare Workers AI**
  (`typesafe/jev`, appelable dans un Worker), **OpenRouter**
  (`typesafe/jev-1.13` et `typesafe/jev-latest` — via un endpoint décisions
  dédié `/api/alpha/decisions`, pas le chat completions habituel),
  **LangChain** (intégration `TypeSafeClassifier`, questions exposées comme
  Runnable), **Langfuse**, Netlify AI Gateway, MotherDuck (`prompt_jev()` en
  SQL).

### 137.5. Pricing et positionnement vs un LLM classique

- Chiffres vérifiés (fournisseur) : **0,042 $ / million de tokens input,
  output gratuit** (pas de texte généré → rien à facturer), latence
  **70–500 ms**, contexte **64k** (des notes communautaires détaillent 32k
  pour le state + la plus longue question — à vérifier dans la doc),
  **1 200 requêtes/min** et 250k tokens/s (limites communautaires, dynamiques).
- Comparaison pour **la même tâche de décision** (ex. router un ticket) :

|  | Jev | LLM frontier classique (même tâche) |
|---|---|---|
| Coût input | 0,042 $/M tokens | plusieurs $/M tokens (ex. GPT-5 Nano : 0,05 $/M) |
| Coût output | **0** (pas de génération) | facturé plein pot (raisonnement + JSON) |
| Latence | 70–500 ms | 3–329 s (mesures TypeSafe sur LLM frontier) |
| Coût par cas (eval TypeSafe) | 0,0004 $/cas | 0,176 $/cas (classe Opus) — chiffres **fournisseur**, non reproduits indépendamment |
| Fiabilité du format | garantie par construction | parsing + retries de JSON malformé |
| Ce qu'il ne fait pas | rédiger, raisonner, appeler des outils | tout le reste (c'est son métier) |

- Repère indépendant : une étude de calibration a fait tourner **4 621
  évaluations pour ~0,06 $** — à ce prix, tester un seuil de décision coûte
  moins cher que le café. L'économie réelle n'est pas « Jev vs LLM », c'est
  **« N décisions par jour × (coût LLM − coût Jev) »** : le gain explose sur
  les volumes (triage d'alertes, routage de tickets).

### 137.6. Cas d'usage concrets pour agents

- **Routage (quel outil appeler)** : l'agent produit un plan (« je dois
  redémarrer nginx sur srv-web-03 »), Jev classe l'intention en `choice`
  parmi les outils disponibles + `noul` « l'outil choisi couvre-t-il la
  demande ? ». Le LLM ne sert plus qu'aux cas ambigus.
- **Scoring (prioriser des tickets)** : `score` sur une rubrique
  sévérité/urgence définie par toi, `noul` « ce ticket est-il actionnable
  sans humain ? ». Tri automatique de files à fort volume.
- **Classification** : intent, sentiment, catégorie de document — le cas
  d'école du quickstart officiel (département billing/technical/sales +
  frustration + urgence, **3 questions en 1 appel**).
- **Garde-fous (le cas le plus pertinent pour ce guide)** : avant chaque
  action à effet de bord d'un agent autonome, une question `noul` « cette
  action est-elle sûre ? » + `choice` du niveau de risque. Pattern :
  `confidence ≥ seuil → exécuter`, `en dessous → escalader vers le LLM`,
  `risque high → humain`. C'est exactement la « couche de décision rapide
  devant un LLM » : **Jev décide vite et pas cher, le LLM n'intervient que
  sur l'ambigu**.
- **Évaluation / calibration de pipelines** : scorer des sorties de LLM à
  la chaîne (pertinence, toxicité, conformité à une spec) pour quelques
  centimes.

### 137.7. Limites honnêtes

- **Texte uniquement** : pas d'image, pas d'audio, pas de vidéo. Pour un
  agent multimodal, la couche vision reste un LLM classique.
- **Il ne voit que le `state` que tu assembles** : pas d'outils, pas de
  navigation, pas de lecture de fichier à deux sauts. La qualité de la
  décision dépend entièrement de la qualité du state — garbage in, décision
  calibrée sur du garbage.
- **Pas un raisonneur** : les évaluations indépendantes convergent — Jev
  marche mieux sur des questions **petites et décomposées** ; une question
  qui demande un raisonnement étendu, c'est le travail du LLM, pas le sien.
- **Écosystème jeune** : pas de self-host / VPC / on-prem (API hébergée
  uniquement), **pas de fine-tuning** (on le pilote via state + criteria
  uniquement), données d'entraînement maison non publiées, benchmarks
  **fournisseur uniquement** (193,6× plus rapide / 444,6× moins cher :
  chiffres TypeSafe, à prendre avec des pincettes), signups directs en pause
  au 27/09/2026.
- **Anglais > autres langues** : le français fonctionne mais la doc
  recommande de tester la calibration dans la langue cible avant de mettre
  des seuils en production.
- **Critique de fond** (Simon Willison) : c'est un retour vers des systèmes
  plus « boîte noire » — on met du texte, on récupère un flottant, et un
  classement (ex. « bonne ville ? » : Cupertino en haut, East Palo Alto en
  bas) est **difficile à interroger**. Pour un garde-fou critique, garde
  toujours une voie d'escalade humaine.

### 137.8. Ce que ça change dans une architecture d'agent

