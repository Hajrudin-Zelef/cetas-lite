---
id: collect-261001-rattrapage/rattrapage/grafana-guide-9
title: "Guide Grafana — Dashboards, visualisation et alerting"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-27"]
keywords: ["arr", "valuation"]
source: docs/RAG/collect-261001-rattrapage/grafana_guide.md
source_anchor: ""
source_lines: [1684, 1891]
sha256: 299cdc5e4de0991014cb3ba8da44eb5f7c9f1cbf7b4c07b580550e09a7b2b34c
---

# Guide Grafana — Dashboards, visualisation et alerting

| Expression | Rôle | Exemple |
|---|---|---|
| **Reduce** | Agrège chaque série en un nombre | `Last`, `Mean`, `Max`, `Min`, `Sum`, `Count` |
| **Math** | Calculs et comparaisons | `$B > 85`, `$A + $B`, `$B / $C * 100` |
| **Resample** | Rééchantillonne sur une fenêtre | Moyenne sur 10 min pour lisser |
| **Classic condition** | Ancien modèle simple | `WHEN avg() OF A IS ABOVE 5` |
| **Threshold** | Variante visuelle du seuil | Équivalent Math lisible |

Exemples commentés :

```
# Alerte si le p95 de latence dépasse 2s pendant 5 min (avec lissage)
A : histogram_quantile(0.95, sum by (le) (rate(http_request_duration_seconds_bucket[5m])))
B : Resample — Input A, Window 10m, Downsampler Mean, Upsampler Fill
C : Reduce — Input B, Function Last
D : Math — $C > 2

# Alerte si AUCUNE donnée reçue depuis 15 min (mort du collecteur)
A : up{job="node"}
B : Reduce — Input A, Function Last
C : Math — $B < 1        # up vaut 0 ou 1 ; absent = NoData (géré séparément)

# Comparer deux requêtes : erreurs en % du trafic
A : sum(rate(http_requests_total{code=~"5.."}[5m]))
B : sum(rate(http_requests_total[5m]))
C : Math — ($A / $B) * 100 > 5
```

> **Piège :** dans une expression Math, `$B` désigne la **valeur réduite** de
> B. Si B renvoie plusieurs séries, la comparaison s'applique série par série
> et chaque série devient une **instance d'alerte** distincte (avec ses
> labels). C'est voulu : une alerte par disque plein, pas une alerte globale.

---

## 45. Contact points : e-mail, Slack, webhook, Telegram

**Alerting → Contact points → New contact point.** Un contact point peut
contenir **plusieurs intégrations** (e-mail + Slack + webhook).

**E-mail :**

```
Name: equipe-systemes-mail
To: systemes@mondomaine.fr
Subject: [{{ .Status | toUpper }}] {{ .GroupLabels.SortedPairs.Values | join " " }}
```
Nécessite le SMTP configuré (section 15). Testez avec le bouton **Test**.

**Slack :**

```
Webhook URL: https://hooks.slack.com/services/<A_COMPLETER>
Channel: #alertes-systemes
Title: {{ template "slack.title" . }}
Text: {{ template "slack.text" . }}
```
(Créez le webhook dans les paramètres de votre workspace Slack : Apps →
Incoming WebHooks.)

**Webhook générique** (vers votre ITSM, GLPI, script maison…) :

```
URL: https://itsm.mondomaine.fr/api/grafana-alerts
HTTP Method: POST
Authorization Header: Bearer <A_COMPLETER>
```
Grafana envoie un JSON avec `alerts[]` (status, labels, annotations,
`startsAt`, `endsAt`, `values`). Documentez le format dans votre wiki.

**Telegram :**

```
Bot Token: <A_COMPLETER>   (via @BotFather)
Chat ID: <A_COMPLETER>     (via @userinfobot ou l'API getUpdates)
```

> **Règle d'or :** chaque contact point critique doit être **testé de bout en
> bout** (bouton Test + vraie règle de test) avant la mise en production.
> Une alerte qui part dans le vide est pire que pas d'alerte.

---

## 46. Notification policies : routage, regroupement, répétition

**Alerting → Notification policies.** C'est un **arbre** : la policy racine
(`default`) puis des sous-policies qui matchent sur les labels.

Exemple d'arbre pour une équipe systèmes :

```
default: contact=defaut-mail, group_by=[alertname], group_wait=30s,
         group_interval=5m, repeat_interval=4h
├── matcher severity=critical → contact=astreinte-telephone
│   (group_wait=0s : notification immédiate)
├── matcher equipe=systemes,severity=warning → contact=systemes-slack
│   (group_by=[alertname, instance], repeat_interval=12h)
└── matcher equipe=reseau → contact=reseau-mail
```

Sémantique des temporisations :

| Paramètre | Rôle | Valeur conseillée |
|---|---|---|
| `group_wait` | Attente avant le 1er envoi (pour grouper les alertes qui arrivent ensemble) | `30s` (0s si critique) |
| `group_interval` | Délai entre deux envois pour le même groupe | `5m` |
| `repeat_interval` | Rappel tant que l'alerte est active | `4h` (critique), `12-24h` (warning) |

**Bonnes pratiques :**

- `group_by: [alertname]` au minimum : une notification « 12 disques pleins »
  plutôt que 12 e-mails.
- Ne routez vers l'astreinte téléphonique/SMS **que** `severity=critical`.
- Faites relire l'arbre à quelqu'un d'autre : une erreur de matcher et les
  alertes critiques partent dans le vide (voir section 74).

---

## 47. Silences et mute timings (plages de maintenance)

**Silence** (ponctuel) : **Alerting → Silences → New silence.**

```
Matchers: instance=srv-web-03, alertname=Disque plein imminent
Starts: 2026-09-27 02:00   Ends: 2026-09-27 06:00
Comment: "Migration disque — ticket CHG-1234"
```

- Un silence n'arrête pas l'évaluation : l'alerte passe en Firing mais la
  **notification est supprimée**. L'historique reste visible.
- **Ne créez jamais de silence permanent** « en attendant » : mettez une date
  de fin et un ticket de référence dans le commentaire.

**Mute timing** (récurrent) : **Alerting → Notification policies → Mute timings.**

```
Name: week-end-non-critique
Time intervals:
  - weekdays: [saturday, sunday]
    times: [{start_time: "00:00", end_time: "23:59"}]
```
Puis associez-le à une policy (`mute_time_intervals`). Cas d'usage : pas de
`warning` le week-end, mais les `critical` passent toujours (ne mutez jamais
les critiques sans validation écrite du responsable).

> **Audit :** la liste des silences actifs est visible dans Alerting →
> Silences. En revue d'exploitation hebdo, vérifiez qu'aucun silence
> « temporaire » ne traîne depuis 3 semaines.

---

## 48. Bonnes pratiques d'alerting (ce qui évite la fatigue d'alerte)

La fatigue d'alerte tue l'alerting : trop de bruit → on ignore tout → on rate
la vraie panne. Règles d'hygiène :

1. **Alertez sur des symptômes, pas sur des causes** : « le site ne répond
   pas » (blackbox) avant « CPU à 92 % ».
2. **Chaque alerte doit être actionnable** : si personne ne sait quoi faire
   quand elle sonne, c'est un dashboard, pas une alerte. Ajoutez un lien
   runbook dans l'annotation `description`.
3. **`for` adapté** : 5-10 min pour les warnings, 0-2 min pour les critiques
   vitales. Un `for` à 0 partout = bruit.
4. **Seuils avec marge** : alerte warning à 85 %, critique à 95 % — pas
   l'inverse, pas 99 % (trop tard).
5. **Pas d'alerte sur des métriques qui n'existent pas encore** : testez la
   requête dans Explore d'abord.
6. **Nommez explicitement** : `Disque plein imminent` > `Alerte 12`.
7. **Labels homogènes** : `severity` (critical/warning/info) et `equipe`
   sur **toutes** les règles — c'est la clé du routage.
8. **Limitez le nombre de règles** : 20 règles pertinentes > 200 règles
   copiées d'un template générique jamais triées.
9. **Revoyez trimestriellement** : quelles alertes ont sonné ? Lesquelles
   étaient des faux positifs ? Ajustez ou supprimez.
10. **Testez la chaîne complète** 1 fois par mois : déclenchez une alerte de
    test et vérifiez la réception (e-mail/Slack/téléphone).

---

## 49. Provisioning as code : le principe

Le **provisioning** permet de définir datasources, dashboards, contact points
et règles d'alerte dans des **fichiers YAML/JSON versionnés dans Git**,
appliqués automatiquement au démarrage de Grafana.

Dossier : `/etc/grafana/provisioning/` :

```
provisioning/
├── dashboards/
│   └── dashboards.yaml        # où trouver les JSON de dashboards
├── datasources/
│   └── datasources.yaml       # les datasources
└── alerting/
    ├── contact-points.yaml
    ├── policies.yaml
    └── rules.yaml
```

Avantages :

- **Reproductibilité** : un nouveau serveur Grafana se reconstruit en
  minutes depuis Git.
- **Revue** : chaque changement de dashboard passe par une merge request.
- **Sauvegarde** : Git **est** une partie de la sauvegarde (mais pas la
  seule : la base interne contient aussi utilisateurs, historique, versions —
  voir sections 62-64).

