---
id: collect-261001-ia-llm/ia-llm/prompt-injection-securiser-son-llm-en-13-etapes-2026-4
title: "prompt-injection-securiser-son-llm-en-13-etapes-2026"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["cyber", "incident", "open source"]
source: docs/RAG/collect-261001-ia-llm/prompt-injection-securiser-son-llm-en-13-etapes-2026.md
source_anchor: ""
source_lines: [260, 339]
sha256: d791f867a17481ab0684d9c8d089eef79706c216721eb235ce7b81c9eb0465d8
---

# prompt-injection-securiser-son-llm-en-13-etapes-2026

Ajoutez un scanner de secrets (gitleaks ou trufflehog, par exemple) directement dans votre pipeline CI/CD pour bloquer toute clé API, jeton ou mot de passe qui se glisserait accidentellement dans un commit. Ce contrôle coûte quelques secondes d’exécution à chaque pull request et évite l’un des scénarios de fuite les plus fréquents sur les projets IA, celui d’une clé de test oubliée dans un fichier de configuration versionné puis republié en open source par erreur.

## Étape 10 : mettre en place la supervision et l’alerting en production

Une fois les journaux de l’étape 6 centralisés dans votre SIEM, définissez des règles de corrélation spécifiques aux applications IA. Un taux anormal de détections du pré-filtre sur une même session, une succession de tentatives de contournement du prompt système, ou un pic de refus du filtre de sortie constituent des signaux qui méritent une alerte immédiate à l’équipe sécurité, pas seulement une ligne de log noyée parmi des milliers d’autres événements.

Le contexte français rend cette vigilance d’autant plus nécessaire que l’ENISA observe une activité soutenue de groupes de rançongiciels comme Qilin, Hunters International et CL0P contre des plateformes SaaS riches en données, une catégorie où entrent la majorité des applications métier dotées d’un assistant IA. Une supervision réactive réduit le temps entre la première tentative d’intrusion et sa détection, un facteur qui conditionne directement l’ampleur des dégâts en cas d’incident réel.

## Étape 11 : documenter la conformité à l’AI Act européen

Depuis le 2 août 2026, les fournisseurs et déployeurs de systèmes d’IA à haut risque doivent démontrer, documents à l’appui, la mise en œuvre d’une gestion des risques, d’une gouvernance des données robuste, d’une documentation technique détaillée, d’une journalisation complète, d’une supervision humaine effective et de garanties de robustesse et de cybersécurité. Concrètement, votre dossier de conformité doit inclure la cartographie de l’étape 1, les résultats de vos scans de red teaming de l’étape 8, vos journaux d’audit de l’étape 6 et la liste des CVE corrigées identifiées à l’étape 7.

La Commission européenne a précisé qu’elle exerce désormais pleinement ses pouvoirs de supervision et de sanction sur les modèles à usage général présentant des risques cyber systémiques. Les organisations qui traitent la documentation de conformité comme une simple formalité administrative, sans lien avec les mesures techniques réellement déployées, s’exposent à un décalage risqué entre ce qu’elles déclarent et ce qu’elles opèrent réellement en production.

- Registre des systèmes IA classés à haut risque, avec justification de la classification retenue
- Cartographie des flux de données et des sources de contenu tiers (issue de l’étape 1)
- Historique des scans de red teaming et des taux de succès des attaques testées (issu de l’étape 8)
- Journal d’audit conservé sur une durée définie par votre politique de rétention (issu de l’étape 6)
- Liste des dépendances et de leur statut de correctif face aux CVE connues (issue de l’étape 7)
- Procédure de supervision humaine décrivant qui peut interrompre ou corriger une décision du système

Conservez ces éléments dans un espace centralisé accessible à l’équipe conformité, pas uniquement dans les outils internes de l’équipe technique. Un contrôle réglementaire se déroule rarement au rythme des sprints de développement, et la capacité à produire ce dossier en quelques heures plutôt qu’en plusieurs semaines fait souvent la différence lors d’une inspection.

## Étape 12 : tester la résilience face aux abus de volume et au déni de service

Un modèle de langage reste une ressource de calcul coûteuse, et un attaquant qui échoue à provoquer une injection peut se rabattre sur une saturation pure et simple de vos appels API pour générer des coûts ou indisponibilités. Mettez en place une limitation de débit par session et par utilisateur, avec des seuils différenciés entre trafic authentifié et anonyme, et surveillez les schémas de requêtes répétitives qui cherchent à épuiser votre budget de tokens plutôt qu’à obtenir une réponse utile.

```
from fastapi import FastAPI, HTTPException, Request
from collections import defaultdict
from time import time
app = FastAPI()
request_log: dict[str, list[float]] = defaultdict(list)
RATE_LIMIT = 20         # requêtes
WINDOW_SECONDS = 60     # par minute
@app.middleware("http")
async def rate_limiter(request: Request, call_next):
    client_id = request.headers.get("x-session-id", request.client.host)
    now = time()
    request_log[client_id] = [t for t in request_log[client_id] if now - t < WINDOW_SECONDS]
    if len(request_log[client_id]) >= RATE_LIMIT:
        raise HTTPException(status_code=429, detail="Trop de requêtes, réessayez dans une minute")
    request_log[client_id].append(now)
    return await call_next(request)
```
## Étape 13 : projet complet, assembler toutes les couches dans une application FastAPI

Voici l’assemblage final qui combine les douze étapes précédentes dans une seule application de démonstration fonctionnelle. Ce squelette constitue une base de départ réaliste pour un projet interne, pas un produit fini, mais il illustre l’articulation complète des couches défensives.

```
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
app = FastAPI(title="Assistant IA durci")
class ChatRequest(BaseModel):
    session_id: str
    message: str
@app.post("/chat")
async def chat_endpoint(payload: ChatRequest):
    # 1. Pré-filtre (étape 2)
    scan = prefilter_prompt(payload.message)
    if scan.is_suspect:
        log_interaction(payload.session_id, scan, [], {"safe": True, "pii_detected": {}, "suspect_links": []})
        raise HTTPException(status_code=400, detail="Requête rejetée par le contrôle de sécurité")
    # 2. Isolation du prompt système (étape 3)
    final_prompt = build_prompt(scan.normalized_input)
    # 3. Appel du modèle (fournisseur au choix, non détaillé ici)
    model_output, tool_calls = call_llm_with_tools(final_prompt)
    # 4. Filtre de sortie (étape 4)
    post_check = postfilter_output(model_output)
    if not post_check["safe"]:
        log_interaction(payload.session_id, scan, tool_calls, post_check)
        raise HTTPException(status_code=502, detail="Réponse bloquée par le contrôle de sortie")
    # 5. Journalisation complète (étape 6)
    log_interaction(payload.session_id, scan, tool_calls, post_check)
    return {"response": model_output}
```
La fonction call_llm_with_tools n’est volontairement pas détaillée ici, car son implémentation dépend de votre fournisseur de modèle. Elle doit néanmoins respecter la logique d’allowlist de l’étape 5 pour tout appel d’outil qu’elle déclenche.

## Exemples de sortie attendus une fois le déploiement terminé

