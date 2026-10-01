---
id: collect-261001-ia-llm/ia-llm/ia-generatif-media-5
title: "IA générative : image, vidéo, recherche"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Apple", "DeepSeek", "Google", "Microsoft", "OpenAI", "Perplexity", "xAI"]
dates: ["2026-09-27"]
keywords: ["arr", "attention", "claude", "compute", "deepseek", "gemini", "grok", "incident", "llama", "mcp", "perplexity", "reasoning"]
source: docs/RAG/collect-261001-ia-llm/ia_generatif_media.md
source_anchor: ""
source_lines: [440, 524]
sha256: 19394d56bfe8f497651a09ef496e2b3d3c0f2f329a4d1972f923aeea2b524516
---

# IA générative : image, vidéo, recherche

Points d'attention : `schnell` est ultra-rapide mais moins fidèle que `dev` ; pour FLUX.2 via diffusers, vérifier la classe de pipeline supportée dans la version installée de `diffusers` (le support suit les releases BFL avec quelques semaines de décalage — **à vérifier** au moment où tu lis ceci).

## 26. FLUX : cas d'usage pro (angle Zelef)

1. **Illustrations de documentation technique** : schémas de principe (synoptique onduleur, cheminement câbles), illustrations pour procédures. Générer en local (klein 4B / schnell) = 0 coût marginal, confidentialité totale.
2. **Photos d'ambiance pour supports** : salle serveurs, atelier, armoire électrique — quand tu n'as pas le droit de photographier un site client ou que la photo réelle est inexploitable.
3. **Maquettes avant intervention** : visualiser un réaménagement de local technique (nouvelle baie, nouvel onduleur) avec Kontext en édition d'une photo réelle : « ajoute une deuxième armoire à droite ».
4. **Cohérence de personnage/mascotte** : multi-références FLUX.2 pour garder le même technicien/équipement sur une série de visuels de formation.
5. **Schémas avec texte** : FLUX.2 et Ideogram sont les meilleurs pour le texte dans l'image — utile pour des étiquettes, mais **toujours relire** : une étiquette « 400 V » devenue « 4000 V » dans une doc est un incident.

## 27. FLUX : limites honnêtes

- **Le photoréalisme « parfait »** : Nano Banana Pro et Imagen 4 Ultra restent devant pour la peau, les portraits, la photo produit haut de gamme (tests indépendants juin 2026).
- **Petite typographie** : les longs passages de texte finissent par baver ; pour une affiche avec beaucoup de texte, composer le texte **après** en PAO, pas dans le prompt.
- **Le « look propre » par défaut** : FLUX sort des images très nettes, parfois trop lisses (« AI sheen »). Pour un grain documentaire, le post-traitement ou un prompt explicite (« grain photographique, lumière naturelle imparfaite ») aident.
- **Vitesse en local** : 1-2 min/image en 1024² sur RTX 4090 (FLUX.2 dev, 20 steps). Pour du volume, le cloud ou klein sont préférables.
- **Édition conversationnelle** : Kontext est excellent en édition par instruction, mais le pattern « chat multi-tours » de Nano Banana est plus naturel pour itérer en dialogue.

---

# PARTIE C — PERPLEXITY : la recherche IA

## 28. Perplexity : ce que c'est (« plexplxity » = coquille)

**Perplexity AI** est un **moteur de réponse** (answer engine) : tu poses une question en langage naturel, il cherche sur le web en temps réel et te répond avec **des citations sources cliquables**. La différence avec un chatbot classique : chaque affirmation importante est ancrée à une source que tu peux vérifier. Pour un métier où une valeur fausse coûte cher (dimensionnement, norme, procédure), c'est le point décisif.

État au 27/09/2026 : produit mature, modèle économique **100 % abonnement/API** — la régie publicitaire lancée en octobre 2025 a été **arrêtée en février 2026** (confiance utilisateurs et précision des réponses). Perplexity a aussi signé un partenariat compute avec Microsoft Azure (annoncé janvier 2026, ~750 M$ selon la presse).

## 29. Perplexity : offres et prix (vérifiés le 27/09/2026)

| Offre | Prix | Contenu essentiel |
|-------|------|-------------------|
| **Free** | 0 $ | Recherches basiques illimitées + ~5 recherches Pro/jour, citations |
| **Pro** | **20 $/mois** (200 $/an, ~16,67 $/mois) | Recherches Pro illimitées, ~20 Deep Research/jour (jusqu'à 500/jour selon sources), choix du modèle (GPT-5.x, Claude, Gemini 3.x, Grok, Sonar), upload de fichiers, Spaces, génération d'image/vidéo, **5 $/mois de crédit API inclus** |
| **Max** | **200 $/mois** (2 000 $/an) | Tout Pro + Labs illimités, Deep Research illimité, Perplexity Computer (10 000 crédits/mois), vidéo cinématique, accès prioritaire aux nouveautés |
| **Education Pro** | 10 $/mois | Pro pour étudiants vérifiés |
| **Enterprise Pro** | 40 $/siège/mois (400 $/an) | Admin, SSO, Spaces partagés, sécurité renforcée |
| **Enterprise Max** | 325 $/siège/mois | Limites maximales |
| **Comet** (navigateur) | **Gratuit** depuis octobre 2025 | Navigateur avec assistant intégré (Windows, Mac, Android, iOS) |
| **Comet Plus** | 5 $/mois | Contenus éditeurs premium |

Fonctionnalités notables : **Spaces** (espaces de travail collaboratifs avec tâches planifiées qui rafraîchissent une réponse périodiquement — parfait pour une veille), **Labs** (génère tableurs, rapports, mini-apps), **Model Council** (pose la même question à plusieurs modèles frontières et compare), **Deep Research** (recherche approfondie multi-sources).

## 30. L'API Sonar : oui, elle existe (et c'est le vrai sujet pro)

L'**API Sonar** expose la recherche web temps réel + génération avec citations dans tes propres outils. Caractéristiques vérifiées le 27/09/2026 :

- **Endpoint** : `https://api.perplexity.ai` — **compatible OpenAI** (tu peux réutiliser tes clients/SDK existants en changeant l'URL de base et la clé).
- **Modèles** : `sonar` (léger/rapide), `sonar-pro` (recherche approfondie), `sonar-reasoning`, `sonar-deep-research`, plus une **Search API** (résultats web bruts, sans génération).
- **Pas d'offre gratuite permanente** côté API (des crédits d'essai ~25-50 $ pour les nouveaux comptes selon les sources ; les abonnés Pro reçoivent **5 $/mois** de crédit API).
- **Tarifs indicatifs** (grille publique, vérifiée sept 2026) :
  - `sonar` : **~1 $ / 1 M tokens** (entrée et sortie)
  - `sonar-pro` : **~3 $ / 1 M tokens en entrée, ~15 $ / 1 M en sortie**
  - **Frais de requête** : **5 $ à 14 $ pour 1 000 requêtes** selon le modèle (la Search API brute : 5 $/1 000 recherches)
  - Tokens de citation ~2 $/1M, tokens de raisonnement ~3-8 $/1M selon les docs tierces (**à vérifier** sur la page officielle)
- SDK officiels **Python et TypeScript**, serveur **MCP** (Cursor, VS Code, Claude Desktop), programme startups (6 mois d'Enterprise Pro + 5 000 $ de crédits API selon les sources — **à vérifier**).
- Les modèles Sonar sont **propriétaires** (fine-tunés sur des bases ouvertes type Llama 3.3 / DeepSeek R1 selon les analyses, mais poids et infra de recherche fermés).

Exemple minimal (compatible OpenAI) :

```python
import os
from openai import OpenAI   # client OpenAI standard, réutilisé tel quel

client = OpenAI(
    api_key=os.environ["PERPLEXITY_API_KEY"],
    base_url="https://api.perplexity.ai",
)

resp = client.chat.completions.create(
    model="sonar-pro",
    messages=[{
        "role": "user",
        "content": "Quelle est la tension de floating recommandée pour des batteries "
                   "VRLA AGM d'onduleur en 2026 ? Donne les valeurs par élément et "
                   "cite tes sources.",
    }],
)
print(resp.choices[0].message.content)
# La réponse inclut des citations [1], [2]... : les récupérer et les vérifier,
# surtout pour des valeurs d'ingénierie.
```

## 31. Perplexity : cas d'usage pro (angle Zelef)

