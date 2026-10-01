---
id: collect-261001-rattrapage/rattrapage/glpi-guide-11
title: "Guide GLPI ultra-complet — Ticketing, gestion de parc et pilotage SAV"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "apache", "memory"]
source: docs/RAG/collect-261001-rattrapage/glpi_guide.md
source_anchor: ""
source_lines: [1740, 1953]
sha256: ab846400a1f98a679d21497359c756b39e852621028945bb68d11a6118dcda51
---

# Guide GLPI ultra-complet — Ticketing, gestion de parc et pilotage SAV

- [ ] Connexion avec un compte admin.
- [ ] Un ticket créé la veille de la sauvegarde est présent avec ses suivis.
- [ ] Une pièce jointe s'ouvre.
- [ ] Les compteurs copieurs affichent les dernières valeurs.
- [ ] Le cron tourne (`front/cron.php` manuel OK).
- [ ] Envoi d'un courriel de test OK.

> **Exercice annuel :** restauration complète sur VM vierge, chronométrée.
> Objectif : RTO (durée max d'interruption) et RPO (perte de données max)
> **écrits et validés** par la direction.

---

# PARTIE X — MISE À JOUR, PLUGINS, API

## 50. Mise à jour de GLPI : méthode sans casse

### Règle d'or

**Jamais de mise à jour sans sauvegarde complète + testée** (partie IX),
et **jamais directement en production** : d'abord sur une **copie de test**.

### Procédure (ex. 10.0.x → 10.0.y, ou 10.x → 10.(x+1))

```bash
# 1. Sauvegarde complète (section 48) — VÉRIFIER le dump
# 2. Cloner en environnement de test (VM ou dossier + base _test)
# 3. Sur le test : déployer la nouvelle version
cd /tmp && wget https://github.com/glpi-project/glpi/releases/download/10.0.XX/glpi-10.0.XX.tgz
tar xzf glpi-10.0.XX.tgz
sudo rsync -a --delete glpi/ /var/www/glpi/ \
  --exclude=config --exclude=files --exclude=marketplace
sudo chown -R www-data:www-data /var/www/glpi
# 4. Lancer la mise à jour : https://glpi-test/ → assistant « Mise à jour »
#    ou en CLI : sudo -u www-data php /var/www/glpi/bin/console db:update --no-interaction
# 5. Mettre à jour les PLUGINS vers leurs versions compatibles 10.x
# 6. Tests : connexion, ticket complet, recherche, impression PDF, API, cron
# 7. Si OK : reproduire en production (fenêtre de maintenance annoncée)
```

### Points de vigilance

- Vérifiez la **matrice de compatibilité des plugins** avant chaque montée
  de version (un plugin incompatible = page blanche, section 54).
- Les montées **majeures** (ex. 9.5 → 10) exigent souvent PHP plus récent :
  mettez à jour PHP **avant** GLPI.
- Après mise à jour : videz le cache (`Administration > Maintenance > Vider les
  caches`) et relancez le cron manuellement une fois.

---

## 51. Plugins : lesquels installer, lesquels éviter

Installation : `Configuration > Plugins` (marketplace intégré en 10.x) ou dépôt
manuel dans `marketplace/`.

### Plugins utiles pour un SAV (à valider vs votre 10.0.x exacte)

| Plugin | Usage |
|---|---|
| **Champs personnalisés** (Fields) | Compteurs copieurs, codes erreur sur tickets/fiches |
| **Comportements** (Behaviors) | Champs obligatoires, formats |
| **Impression PDF** | Tickets/fiches en PDF pour les clients |
| **Outils d'inventaire / Agent** | Selon version, inventaire natif |
| **Satisfaction avancée** | Sondages enrichis |
| **Tag** | Étiquettes libres (ex. « sous-garantie ») |
| **Carbon** (thème sombre) | Confort techniciens |

### Règles

1. **Un plugin = un besoin réel.** Chaque plugin = surface d'attaque + risque
   à chaque mise à jour.
2. Vérifiez : compatible 10.x ? maintenu (commit < 6 mois) ? avis communauté ?
3. Testez sur l'environnement de **test** d'abord.
4. Avant chaque mise à jour GLPI : mettez à jour les plugins **en premier**.
5. Désinstallez (pas seulement désactivez) les plugins abandonnés.

---

## 52. API REST : authentification et premiers appels

L'API REST (`https://glpi.entreprise.lan/apirest.php`) permet de créer des
tickets, remonter des compteurs, interroger le parc depuis des scripts.

### Activation

1. `Configuration > Générale > API` : activer l'API REST.
2. Créez un utilisateur dédié `svc-api` (profil restreint, ex. Technicien
   sur les entités utiles) — **jamais** le compte admin.
3. `Préférences` de `svc-api` : générer un **token API** (et un token d'app
   si besoin).

### Session

```bash
GLPI="https://glpi.entreprise.lan/apirest.php"
APP_TOKEN="TOKEN_APPLICATION"      # optionnel mais recommandé
USER_TOKEN="TOKEN_UTILISATEUR_svc-api"

# 1. Initier la session
SESSION=$(curl -s -X GET "$GLPI/initSession?user_token=$USER_TOKEN" \
  -H "App-Token: $APP_TOKEN" | python3 -c "import sys,json;print(json.load(sys.stdin)['session_token'])")

# 2. Appel authentifié (ex. : mon profil)
curl -s -H "Session-Token: $SESSION" -H "App-Token: $APP_TOKEN" \
  "$GLPI/getMyProfiles" | python3 -m json.tool

# 3. Tuer la session
curl -s -X GET -H "Session-Token: $SESSION" -H "App-Token: $APP_TOKEN" \
  "$GLPI/killSession"
```

> **Sécurité :** tokens = mots de passe. Stockez-les en coffre, rotation annuelle,
> profil API au moindre privilège, HTTPS obligatoire.

---

## 53. API REST : exemples concrets (tickets, parc, compteurs)

### Créer un ticket d'intervention copieur

```bash
curl -s -X POST -H "Session-Token: $SESSION" -H "App-Token: $APP_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "input": {
      "name": "Bourrage récurrent bac 2 — Mairie, bureau 204",
      "content": "Code J-0511 affiché. Compteur N&B : 212458.",
      "itilcategories_id": 12,
      "urgency": 3, "impact": 2,
      "type": 1,
      "_users_id_requester": 45,
      "entities_id": 3
    }
  }' "$GLPI/Ticket/"
```

### Rechercher les copieurs d'un client

```bash
curl -s -G -H "Session-Token: $SESSION" -H "App-Token: $APP_TOKEN" \
  --data-urlencode "criteria[0][field]=80" \
  --data-urlencode "criteria[0][searchtype]=equals" \
  --data-urlencode "criteria[0][value]=3" \
  "$GLPI/search/Printer" | python3 -m json.tool
```

### Mettre à jour les compteurs (champs personnalisés)

```bash
# Exemple : compteur N&B = champ personnalisé id 101 sur l'imprimante id 27
curl -s -X PUT -H "Session-Token: $SESSION" -H "App-Token: $APP_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"input": {"id": 27, "plugin_fields_compteurnb": 213102}}' \
  "$GLPI/Printer/27"
```

### Script Python : relevé mensuel automatique (SNMP → GLPI)

```python
#!/usr/bin/env python3
"""Relève les compteurs via SNMP et les pousse dans GLPI (API REST)."""
import requests
COPIERS = [  # (ip, communauté, id_fiche_glpi)
    ("192.168.10.21", "COMMUNAUTE_RO", 27),
    ("192.168.10.22", "COMMUNAUTE_RO", 31),
]
# 1. snmpget iso.3.6.1.2.1.43.10.2.1.4.1.1  (prtMarkerLifeCount)
#    via subprocess/pysnmp -> compteur total
# 2. session API (cf. section 52)
# 3. PUT /Printer/<id> avec les champs personnalisés compteurs
# 4. log + alerte si écart > 30 % vs mois précédent
print("Squelette : complétez avec pysnmp + vos IDs de champs personnalisés.")
```

> Ce script, lancé par cron le 1er du mois, **supprime la tournée de relevés
> manuels** sur tout le parc réseau. Vérifiez toujours un échantillon à la main
> les 3 premiers mois.

---

# PARTIE XI — DÉPANNAGE

## 54. Dépannage — page blanche / erreur 500

**Symptôme :** page blanche ou « HTTP 500 » après mise à jour, installation
de plugin, ou modification PHP.

### Diagnostic express

```bash
# 1. Logs applicatifs GLPI (le plus parlant)
tail -n 100 /var/lib/glpi/files/_log/php-errors.log
tail -n 100 /var/lib/glpi/files/_log/error.log
# 2. Log Apache / PHP-FPM
sudo tail -n 100 /var/log/apache2/glpi_error.log
sudo tail -n 100 /var/log/php8.2-fpm.log
# 3. Tester la syntaxe PHP du fichier suspect
php -l /var/www/glpi/marketplace/monplugin/hook.php
```

### Causes fréquentes et remèdes

| Cause | Indice dans les logs | Remède |
|---|---|---|
| Plugin incompatible | `Fatal error ... marketplace/xxx` | Renommer le dossier du plugin en `.off` |
| Extension PHP manquante | `Call to undefined function` | `apt install php-xxx` + restart |
| `memory_limit` trop bas | `Allowed memory size exhausted` | passer à 256M/512M |
| Droits fichiers | `Permission denied ... files/` | `chown -R www-data:www-data` |
| `display_errors=Off` | page blanche sans log web | lire les logs fichiers ci-dessus |
| Cache corrompu après MAJ | erreurs diverses | `Administration > Maintenance > Vider les caches` (ou `rm -rf files/_cache/*`) |

### Mode debug ciblé

