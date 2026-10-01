---
id: collect-261001-rattrapage/rattrapage/grafana-guide-14
title: "Guide Grafana — Dashboards, visualisation et alerting"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/grafana_guide.md
source_anchor: ""
source_lines: [2855, 3052]
sha256: bc8287f73dc8a223c58251df64b7e72c3e05625c53ce26656a4113034e69e3ba
---

# État : sur secteur / sur batterie (Stat avec mapping de valeur)
ups_on_battery{ups="$onduleur"}   # 0 = secteur, 1 = batterie
# Value mappings : 0 → "SECTEUR" (vert), 1 → "BATTERIE" (rouge)

# Autonomie restante estimée en minutes (Gauge, min 0)
ups_battery_runtime_minutes{ups="$onduleur"}
```

**Row « Tendances » (Time series) :**

```promql
# Charge % sur 7 jours (dimensionnement : dérive-t-elle ?)
ups_load_percent{ups="$onduleur"}

# Température batterie (dérive = vieillissement accéléré, cf. loi d'Arrhenius)
ups_battery_temperature{ups="$onduleur"}

# Tension batterie par bloc (écart entre blocs = bloc faible)
ups_battery_voltage{ups="$onduleur"}
```

**Row « Événements » :**

- **State timeline** : `ups_on_battery` (passages sur batterie = coupures
  secteur vues par l'onduleur).
- **Logs** (si syslog de la carte réseau dans Loki) :
  `{job="snmp-ups"} |= "on battery"`.

**Alertes recommandées :**

| Règle | Condition | Severity |
|---|---|---|
| Onduleur sur batterie | `ups_on_battery == 1` for 1m | critical |
| Autonomie faible | `ups_battery_runtime_minutes < 10` | critical |
| Charge élevée | `ups_load_percent > 80` for 15m | warning |
| Batterie à remplacer | `ups_battery_status != "ok"` ou âge > 4 ans | warning |
| Température batterie | `ups_battery_temperature > 30` for 30m | warning |

> **Lien avec l'exploitation :** ce dashboard se combine avec le scénario
> d'arrêt automatique des serveurs (NUT + script de shutdown) : l'alerte
> « autonomie < 10 min » doit arriver **avant** l'arrêt auto, pas après.
> Documentez la procédure dans un panel Text du dashboard.

---

## 72. Kiosque / TV murale (playlists, kiosk mode)

Un écran mural affiche en rotation les dashboards clés, sans interaction.

**1. Organisation et compte dédiés :**

- Créez une organisation `Kiosque`, un utilisateur `kiosque` (rôle **Viewer**).
- Ne partagez que les dossiers nécessaires avec cette orga.
- Alternative sans login : `auth.anonymous` limité à l'orga Kiosque en Viewer
  (section 13) — pratique mais à réserver à un réseau d'affichage isolé.

**2. Playlist :** Dashboards → Playlists → New.

- Ajoutez les dashboards (« Vue d'ensemble », « Réseau », « Onduleurs »),
  par tag ou sélection.
- Intervalle : `1m` par dashboard (assez pour lire, pas assez pour s'ennuyer).
- Cochez **Autofit** si les dashboards ont des tailles différentes.

**3. Kiosk mode :** ouvrez la playlist avec `?kiosk` dans l'URL :

```
https://grafana.mondomaine.fr/playlists/play/<ID>?kiosk
```

- `?kiosk` : masque menus et barre de navigation.
- `?kiosk&autofitpanels` : ajuste les panels à l'écran.
- Pour un navigateur en plein écran au démarrage (Raspberry Pi / mini-PC) :

```bash
chromium-browser --kiosk --incognito \
  "https://grafana.mondomaine.fr/playlists/play/<ID>?kiosk"
```

**4. Bonnes pratiques d'écran mural :**

- Dashboards **dédiés** (pas les dashboards d'investigation) : gros chiffres,
  peu de panels, `Color mode → Background` sur les Stat critiques.
- Refresh auto : `5m` ou `10m` (pas `5s` : ça fatigue et charge les backends).
- **Ne jamais afficher d'identifiants** ni de panel avec des données
  sensibles sur un écran visible par tous.
- Prévoyez le plan B : si Grafana est down, l'écran affiche une erreur —
  c'est aussi un signal (voir section 61).

---

## 73. Dépannage : la datasource ne répond pas

Symptôme : panels en « Error », **Save & test** rouge, ou
`Data source error` dans les logs.

**Arbre de décision (dans l'ordre) :**

1. **Le backend est-il vivant ?** Depuis le **serveur Grafana** :
   ```bash
   curl -s -o /dev/null -w "%{http_code}\n" http://prometheus.mondomaine.fr:9090/-/healthy
   curl -s http://loki.mondomaine.fr:3100/ready
   nc -vz bdd.mondomaine.fr 5432
   ```
   Si ça échoue ici, le problème est réseau/backend, **pas** Grafana.

2. **DNS** : `getent hosts prometheus.mondomaine.fr` — comparez avec ce que
   Grafana résout (un `/etc/hosts` différent ou un DNS menteur arrive plus
   souvent qu'on ne croit).

3. **Pare-feu** : le port est-il ouvert **depuis** le serveur Grafana
   (règles sortantes du FW Grafana + entrantes du FW backend) ?

4. **Identifiants** : testez-les à la main :
   ```bash
   curl -s -u user:<A_COMPLETER> http://prometheus.mondomaine.fr:9090/api/v1/query?query=up
   PGPASSWORD='<A_COMPLETER>' psql -h bdd.mondomaine.fr -U grafana_ro -d metier -c "SELECT 1;"
   ```

5. **URL dans la datasource** : `http` vs `https`, port, path (`/api_jsonrpc.php`
   pour Zabbix). Copiez l'URL exacte qui marche en `curl`.

6. **Logs Grafana** : `grep -i "datasource\|proxy" /var/log/grafana/grafana.log`
   — le message d'erreur exact est souvent explicite
   (`dial tcp: connection refused`, `x509: certificate signed by unknown authority`…).

7. **TLS** : certificat expiré ou CA inconnue → `curl -vk` pour voir, puis
   ajoutez la CA dans la datasource ou corrigez le certificat côté backend.

8. **Timeout** : requêtes lourdes (SQL sur 1 an) → augmentez `queryTimeout`
   dans la datasource.

> **Réflexe :** 8 fois sur 10, c'est le réseau ou les identifiants, pas
> Grafana. Le test `curl`/`nc` **depuis le serveur Grafana** tranche en
> 30 secondes.

---

## 74. Dépannage : les alertes ne partent pas

Symptôme : la règle est en Firing mais personne ne reçoit rien.

**Checklist dans l'ordre :**

1. **La règle s'évalue-t-elle ?** Alerting → Alert rules → état de la règle.
   `Error` ou `NoData` permanent = requête cassée (testez-la dans Explore).
   Vérifiez aussi `grafana_alerting_rule_evaluation_failures_total`.

2. **Le contact point fonctionne-t-il ?** Bouton **Test** sur le contact
   point. Si le test échoue : SMTP (section 15), webhook URL, token Slack —
   corrigez d'abord ça.

3. **Le routage est-il correct ?** Alerting → Notification policies →
   **« View notification routing »** (aperçu du routage pour des labels
   donnés). Erreurs classiques : matcher `severity = Critical` (casse ! les
   labels sont sensibles à la casse) alors que la règle envoie
   `severity="critical"` ; faute de frappe dans le nom du label
   (`equipe` vs `équipe`).

4. **Un silence ou mute timing actif ?** Alerting → Silences : un silence
   oublié mange les notifications en silence (c'est son travail).

5. **`group_wait` / `repeat_interval`** : avec `group_wait: 5m`, la première
   notification arrive 5 min après le Firing — c'est normal, pas un bug.

6. **Notifications en échec côté Grafana :**
   ```promql
   sum(rate(grafana_alerting_notification_failed_total[5m])) > 0
   ```
   + logs : `grep -i "notify\|webhook\|smtp" /var/log/grafana/grafana.log`.

7. **L'alerte est-elle vraiment en Firing ?** `Pending` ≠ notifié (selon
   `for`). Vérifiez l'historique de l'instance (Alerting → instance →
   History).

> **Test de non-régression mensuel :** une règle de test
> (`Test notification — toujours firing` sur une métrique constante) routée
> vers chaque contact point critique, avec un silence permanent… non !
> Sans silence : elle doit notifier à chaque `repeat_interval`. Si le
> téléphone ne sonne plus, vous le saurez.

---

## 75. Dépannage : 10 cas concrets résolus

**Cas 1 — « Data source error: dial tcp: connection refused »**
Cause : Prometheus arrêté ou firewall. Fix : `systemctl status prometheus`
sur le backend, `curl` depuis Grafana (section 73, étape 1).

**Cas 2 — Panels vides après changement d'heure / de fuseau**
Cause : plage de temps relative vs données horodatées en UTC. Fix : vérifier
le fuseau du dashboard (en haut à droite) et celui du backend ; les données
doivent être en UTC côté stockage.

**Cas 3 — `rate()` renvoie des valeurs absurdes**
Cause : `rate()` appliqué à une gauge, ou intervalle < 2× scrape interval.
Fix : n'utiliser `rate()` que sur les compteurs `_total`, avec
`[$__rate_interval]`.

