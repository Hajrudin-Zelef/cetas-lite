---
id: collect-261001-rattrapage/rattrapage/netbox-guide-12
title: "NetBox — Guide complet : source de vérité IPAM / DCIM"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "memory"]
source: docs/RAG/collect-261001-rattrapage/netbox_guide.md
source_anchor: ""
source_lines: [2309, 2520]
sha256: fd731c9a33530ea19fba80256a41fd939765fce73665aca817a0c594b0f6e9b1
---

# NetBox — Guide complet : source de vérité IPAM / DCIM

```bash
# Voir quel Host arrive réellement à Django
tail -20 /var/log/netbox/gunicorn-error.log
# Corriger configuration.py :
ALLOWED_HOSTS = ['netbox.lan-entreprise.fr', '10.10.0.15']
# puis :
python -m py_compile /opt/netbox/netbox/netbox/configuration.py && \
systemctl restart netbox
```

> 💡 En LAB uniquement, `ALLOWED_HOSTS = ['*']` dépanne, mais ne laissez JAMAIS
> cela en production (vulnérabilité Host header attack).

---

## 71. Erreur n°2 : page blanche / 500 après modification de `configuration.py`

**Symptômes** : erreur 500 sur toutes les pages après un changement de config.

**Cause n°1** : erreur de syntaxe Python dans `configuration.py` (virgule manquante,
guillemet non fermé). **Cause n°2** : service non redémarré après modification.

```bash
# 1. Vérifier la syntaxe AVANT tout :
sudo -iu netbox python3 -m py_compile /opt/netbox/netbox/netbox/configuration.py \
  && echo "syntaxe OK"

# 2. Lancer le check Django (détecte les paramètres invalides) :
sudo -iu netbox /opt/netbox/venv/bin/python \
  /opt/netbox/netbox/netbox/manage.py check

# 3. Lire l'erreur exacte :
tail -30 /var/log/netbox/gunicorn-error.log

# 4. Redémarrer PROPREMENT (les deux services) :
systemctl restart netbox netbox-rq
```

> 💡 Prenez l'habitude du rituel « py_compile → check → restart » : il élimine 90 %
> des erreurs 500 liées à la configuration.

---

## 72. Erreur n°3 : « could not connect to server » — PostgreSQL injoignable

**Symptômes** : 500 au démarrage, logs gunicorn : `connection refused`, `password
authentication failed for user "netbox"`.

**Arbre de décision** :

```bash
# PostgreSQL tourne-t-il ?
systemctl status postgresql --no-pager
pg_lsclusters

# Le mot de passe de configuration.py correspond-il ?
PGPASSWORD='...' psql -h localhost -U netbox -d netbox -c 'SELECT 1;'
# -> si KO ici, le problème est côté BDD (mot de passe, droits), pas côté NetBox.

# L'utilisateur a-t-il les droits sur la base ?
sudo -u postgres psql -c '\l' | grep netbox
sudo -u postgres psql -c "ALTER DATABASE netbox OWNER TO netbox;"

# Cas classique après restauration : objets appartenant à postgres au lieu de netbox
# -> refaire le pg_restore en tant que netbox, ou corriger avec :
# sudo -u postgres psql -d netbox -c "REASSIGN OWNED BY postgres TO netbox;"
```

> ⚠️ Après un `pg_restore` fait en `root`/`postgres`, pensez à `chown`-er les objets :
> sinon NetBox lit mais ne peut pas écrire (erreurs 500 à la première modification).

---

## 73. Erreur n°4 : Redis injoignable — cache et files KO

**Symptômes** : lenteurs extrêmes, webhooks jamais envoyés, scripts bloqués « en cours »,
erreurs `Error 111 connecting to localhost:6379. Connection refused.`

```bash
# Redis tourne-t-il ?
systemctl status redis-server --no-pager
redis-cli ping   # PONG attendu

# Si requirepass activé mais pas renseigné dans configuration.py :
redis-cli -a 'CHANGEZ_MOI_Redis_Exemple_2026!' ping

# Vérifier ce que NetBox attend :
grep -A6 "'tasks'" /opt/netbox/netbox/netbox/configuration.py

# Files d'attente bloquées ?
redis-cli -n 0 llen rq:queue:default
redis-cli -n 0 llen rq:queue:high
```

**Remède** : alignez la configuration des deux côtés (mot de passe identique ou
absent des deux côtés), `systemctl restart redis-server netbox netbox-rq`.

---

## 74. Erreur n°5 : `collectstatic` oublié — interface sans CSS

**Symptômes** : après installation ou montée de version, l'interface s'affiche « nue »
(HTML brut, sans mise en forme), ou des icônes/images manquent.

**Cause** : les fichiers statiques n'ont pas été (re)générés, ou nginx ne les sert pas.

```bash
# 1. Régénérer :
sudo -iu netbox /opt/netbox/venv/bin/python \
  /opt/netbox/netbox/netbox/manage.py collectstatic --no-input

# 2. Vérifier le alias nginx :
grep -A2 'location /static/' /etc/nginx/sites-enabled/netbox
# doit pointer vers /opt/netbox/netbox/netbox/static/

# 3. Droits :
ls -ld /opt/netbox/netbox/netbox/static
# lisible par www-data (nginx) : chmod -R a+rX si besoin

nginx -t && systemctl reload nginx
```

> 💡 Après CHAQUE montée de version : `collectstatic` est obligatoire (les assets
> changent). Ajoutez-le à votre checklist de mise à jour (section 68, étape 6).

---

## 75. Erreur n°6 : conflits d'IP et doublons refusés

**Symptômes** : « Cette adresse IP existe déjà » / « Ce préfixe chevauche un préfixe
existant » lors d'une création.

**Ce n'est pas un bug, c'est la protection** : NetBox refuse les doublons dans la
même VRF. Conduite à tenir :

1. Cherchez l'objet existant (recherche globale) : il a peut-être été créé par un
   collègue avec un nom différent.
2. S'il s'agit d'une erreur de saisie (mauvais masque, mauvaise VRF), corrigez
   l'objet existant plutôt que d'en créer un second.
3. Si le chevauchement est légitime (deux usages distincts du même plan), utilisez
   des **VRF** différentes (section 31) : c'est exactement à ça qu'elles servent.
4. Pour les imports CSV massifs : dédupliquez AVANT l'import
   (`sort -u`, ou script Python), sinon l'import s'arrête à la première collision.

---

## 76. Erreur n°7 : token API rejeté (401/403)

**Symptômes** : `curl` renvoie `{"detail":"Invalid token."}` (401) ou
`{"detail":"You do not have permission..."}` (403).

| Code | Cause | Remède |
|---|---|---|
| 401 | token faux, expiré ou révoqué | régénérez-le dans le profil utilisateur |
| 401 | en-tête mal formé | format exact : `Authorization: Token <token>` (pas `Bearer`) |
| 403 | token en lecture seule + tentative d'écriture | cochez « Write enabled » ou utilisez un token dédié |
| 403 | permissions insuffisantes du compte | vérifiez les groupes/permissions (section 43) |

```bash
# Tester un token (doit renvoyer 200 + JSON) :
curl -s -o /dev/null -w "%{http_code}\n" \
  -H "Authorization: Token $NETBOX_TOKEN" \
  https://netbox.lan-entreprise.fr/api/users/me/
```

> 💡 Créez des tokens **par usage** (`svc-ansible-lecture`, `svc-zabbix`, `script-dhcp`)
> avec le périmètre minimal : en cas de fuite, on révoque UN token sans tout casser.

---

## 77. Erreur n°8 : migrations en échec lors de l'installation / upgrade

**Symptômes** : `python netbox/manage.py migrate` s'interrompt avec une traceback
(`django.db.utils...`, `duplicate key`, `relation already exists`).

Conduite à tenir :

```bash
# 1. NE RELANCEZ PAS en boucle : lisez l'erreur complète
sudo -iu netbox /opt/netbox/venv/bin/python \
  /opt/netbox/netbox/netbox/manage.py migrate 2>&1 | tail -40

# 2. Cas fréquent : migration déjà partiellement appliquée
#    -> revenir au snapshot pré-upgrade (section 68, étape 0), COMPRENDRE, puis rejouer.

# 3. Vérifier l'état des migrations :
sudo -iu netbox /opt/netbox/venv/bin/python \
  /opt/netbox/netbox/netbox/manage.py showmigrations | grep -v '\[X\]'
```

**Prévention** : upgrade toujours testé en LAB sur une copie de la BDD de production
(dump restauré). Une migration qui échoue en LAB échouera en production : c'est le
moment de lire les release notes, pas le jour J à 3h du matin.

---

## 78. Erreur n°9 : lenteurs — l'interface rame, l'API timeout

**Symptômes** : pages > 5 s, timeouts gunicorn (120 s), exports CSV qui n'aboutissent pas.

Pistes, dans l'ordre :

1. **Workers gunicorn sous-dimensionnés** : `(2 × vCPU) + 1` minimum (section 13).
   Surveillez `htop` pendant une heure de pointe.
2. **PostgreSQL** : `shared_buffers` trop bas, disque saturé (`df -h`), ou absence
   d'index après un gros import → `VACUUM ANALYZE;` (utilisateur postgres).
3. **Requêtes API sans pagination** : `?limit=1000` au lieu de tout charger ;
   côté scripts, itérez sur `nb.ipam.prefixes.all()` (pynetbox gère la pagination).
4. **Redis saturé** : `redis-cli info memory` → `used_memory` proche de `maxmemory` ?
   Augmentez la RAM ou la limite.
5. **Debug toolbar / DEBUG=True** : vérifiez `DEBUG = False` en production.

