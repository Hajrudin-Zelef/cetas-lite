---
id: collect-261001-ia-llm/ia-llm/prompt-injection-securiser-son-llm-en-13-etapes-2026-3
title: "prompt-injection-securiser-son-llm-en-13-etapes-2026"
domain: ia-llm
role: reference
task: reference
actors: ["CISA", "Hugging Face"]
dates: []
keywords: ["agent", "distribution", "incident", "open source"]
source: docs/RAG/collect-261001-ia-llm/prompt-injection-securiser-son-llm-en-13-etapes-2026.md
source_anchor: ""
source_lines: [139, 259]
sha256: 43b5d315b3c9c41afeb0ac4a23ba0326d9cacaffbc4f61e365fa79b02ea29d24
---

# prompt-injection-securiser-son-llm-en-13-etapes-2026

```
from enum import Enum
from typing import Callable
import signal
class ToolName(str, Enum):
    LOOKUP_ORDER = "lookup_order"
    SEND_EMAIL_TEMPLATE = "send_email_template"
ALLOWED_TOOLS: dict[ToolName, Callable] = {}
def register_tool(name: ToolName):
    def decorator(func: Callable):
        ALLOWED_TOOLS[name] = func
        return func
    return decorator
class ToolTimeout(Exception):
    pass
def _timeout_handler(signum, frame):
    raise ToolTimeout("Exécution de l'outil interrompue (timeout 5s)")
def execute_tool_call(tool_name: str, params: dict) -> dict:
    try:
        tool_enum = ToolName(tool_name)
    except ValueError:
        return {"error": f"Outil non autorisé : {tool_name}"}
    func = ALLOWED_TOOLS.get(tool_enum)
    if func is None:
        return {"error": "Outil non enregistré"}
    signal.signal(signal.SIGALRM, _timeout_handler)
    signal.alarm(5)
    try:
        result = func(**params)
    except ToolTimeout as exc:
        return {"error": str(exc)}
    finally:
        signal.alarm(0)
    return {"result": result}
@register_tool(ToolName.LOOKUP_ORDER)
def lookup_order(order_id: str) -> dict:
    # Requête paramétrée uniquement, jamais de concaténation SQL
    return {"order_id": order_id, "status": "expédiée"}
```
Ce modèle d’allowlist explicite bloque par construction la classe entière de vulnérabilités révélées en août 2026 sur Langflow : même si un flux malveillant parvient à manipuler le modèle, celui-ci ne peut appeler que les fonctions explicitement enregistrées, avec des paramètres validés, jamais une commande arbitraire.

## Étape 6 : journaliser et tracer chaque décision pour l’audit

L’AI Act impose une traçabilité des décisions pour les systèmes à haut risque, mais même en dehors de toute obligation réglementaire, un journal détaillé reste votre meilleur outil d’investigation après incident. Chaque interaction doit enregistrer l’entrée normalisée, le résultat du pré-filtre, le prompt final envoyé au modèle, les appels d’outils déclenchés et le résultat du filtre de sortie.

```
import json
import logging
from datetime import datetime, timezone
logger = logging.getLogger("audit_llm")
def log_interaction(session_id: str, prefilter_result, tool_calls: list, postfilter_result: dict) -> None:
    record = {
        "timestamp": datetime.now(timezone.utc).isoformat(),
        "session_id": session_id,
        "prefilter_suspect": prefilter_result.is_suspect,
        "prefilter_reasons": prefilter_result.reasons,
        "tool_calls": tool_calls,
        "output_safe": postfilter_result["safe"],
        "pii_detected": postfilter_result["pii_detected"],
    }
    logger.info(json.dumps(record, ensure_ascii=False))
```
Envoyez ce flux vers un SIEM (Wazuh, Graylog ou équivalent) plutôt que vers un simple fichier local. C’est ce qui vous permettra, à l’étape 10, de corréler les tentatives d’injection avec d’autres signaux de sécurité et de déclencher des alertes automatiques.

## Étape 7 : corriger les CVE critiques connues sur votre stack IA

Aucune couche applicative ne compense une dépendance non patchée. Avant de considérer votre application comme sécurisée, auditez précisément les versions de vos bibliothèques IA face aux vulnérabilités publiées ces dernières semaines. Le tableau suivant recense les failles critiques les plus pertinentes pour une stack LLM en septembre 2026.

| CVE | Composant | CVSS | Impact | Correctif | 
|---|---|---|---|---|
| CVE-2026-44513 | Hugging Face (chargement de modèles) | 8,8 | Exécution de code arbitraire via modèle piégé | Diffusers ≥ 0.38.0 | 
| CVE-2026-44827 | Hugging Face (chargement de modèles) | 8,8 | Contournement de trust_remote_code=False | Diffusers ≥ 0.38.0 | 
| CVE-2026-45804 | Hugging Face (chargement de modèles) | 7,5 | Exécution de code via dépôt malveillant | Diffusers ≥ 0.38.0 | 
| CVE-2026-9198 | IBM Langflow | Ajoutée au catalogue CISA KEV | Exécution de code à distance via flux visuel | Isolation réseau ou mise à jour immédiate | 
| CVE-2026-42945 | SAP Commerce Cloud | 9,2 | Exécution de code à distance et déni de service | Correctif éditeur (bulletin CERTFR-2026-ACT-035) | 
| CVE-2026-58231 | SAP Commerce Cloud | 10,0 | Contournement de politique de sécurité | Correctif éditeur (bulletin CERTFR-2026-ACT-035) | 

Intégrez cette vérification dans votre routine de maintenance, pas seulement au moment du déploiement initial. Un modèle open source téléchargé aujourd’hui en version corrigée peut redevenir vulnérable si un composant tiers de votre pipeline (bibliothèque de chargement, outil de workflow visuel, plugin) reste figé sur une ancienne version.

## Étape 8 : automatiser le red teaming avec Garak et PyRIT, intégré en CI/CD

Les couches statiques précédentes ne suffisent pas face à un paysage de menaces qui évolue chaque semaine. Le red teaming automatisé consiste à soumettre régulièrement votre application à des batteries d’attaques connues, avec des outils spécialisés comme Garak ou PyRIT, plutôt que d’attendre un audit manuel ponctuel. L’objectif est de faire de ce test un réflexe déclenché à chaque modification du prompt système, du modèle ou des outils exposés.

```
name: llm-security-scan
on:
  pull_request:
    paths:
      - "app/prompts/**"
      - "app/tools/**"
      - "requirements.txt"
jobs:
  garak-scan:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-python@v5
        with:
          python-version: "3.12"
      - name: Installer les dépendances de scan
        run: pip install garak pyrit
      - name: Scanner l'endpoint de l'application
        run: |
          garak --model_type rest \
                --model_name app-staging \
                --probes promptinject,leakreplay \
                --report_prefix ci-scan
      - name: Publier le rapport
        uses: actions/upload-artifact@v4
        with:
          name: garak-report
          path: ci-scan*.jsonl
```
Configurez ce workflow pour échouer la pull request si le taux de succès des attaques dépasse un seuil défini (par exemple 2 %). Cette automatisation transforme le red teaming en garde-fou continu plutôt qu’en exercice annuel, exactement l’approche recommandée par les guides européens de sécurité IA qui appellent à intégrer ces scans directement dans les pipelines de livraison.

## Étape 9 : gérer les secrets et clés API avec un coffre-fort centralisé

Une application IA manipule généralement plusieurs clés sensibles : clé API du fournisseur de modèle, identifiants de connexion aux outils appelés par le LLM, jetons d’accès aux bases de données. Ne stockez jamais ces secrets en clair dans le code, les fichiers de configuration ou les variables d’environnement des conteneurs. Un coffre-fort de type Vault centralise la distribution des secrets, applique une rotation automatique et journalise chaque accès, ce qui limite drastiquement l’impact si un composant de votre pipeline venait à être compromis par une injection réussie.

Appliquez également le principe du moindre privilège aux identifiants utilisés par vos outils : un agent qui ne fait que consulter le statut d’une commande n’a besoin d’aucun droit d’écriture sur la base de données, même si le développement initial en accordait par confort. C’est ce cloisonnement qui empêche une injection réussie sur un seul outil de se propager à l’ensemble de votre système d’information.

