---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1159625-getting-permission-error-from-unifi-network-application-when-c-f27f78a9-1
title: "questions-1159625-getting-permission-error-from-unifi-network-application-when-c-f27f78a9"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1159625-getting-permission-error-from-unifi-network-application-when-c-f27f78a9.md
source_anchor: ""
source_lines: [1, 114]
sha256: ebda631a8d1afd693a3fe1e124159bb8c95837902697373e56f0af5654f1b5e9
---

# questions-1159625-getting-permission-error-from-unifi-network-application-when-c-f27f78a9

I running K3S Kubernetes on a cluster server consisting of a mixture of Raspberry 4 and Raspberry 5 nodes.
I want to install Unifi Network Application on the cluster server and have come pretty far in configuring along with MetalLB, Longhorn and Cert-Manager by using Ansible.
I am using linuxserver.io's Docker image for my cluster server and it requires me to install a seperate pod containing MongoDB.
Searching online for how to install MongoDB gave me the following Ansible file:
- name: Generate MongoDB secrets
  kubernetes.core.k8s:
    state: present
    definition:
      apiVersion: v1
      kind: Secret
      metadata:
        name: mongodb-secret
        namespace: default
      type: Opaque
      data:
        password: VmVyeVNlY3JldFBhc3NvcmQ= # "VerySecretPassword" encoded in Base64
- name: Create Unifi DB init script
  kubernetes.core.k8s:
    state: present
    definition:
      apiVersion: v1
      kind: ConfigMap
      metadata:
        name: mongodb-init-script
        namespace: default
      data:
        init-script.js: |-
          db.getSiblingDB("unifi-db").createUser({user: "unifi", pwd: "unifi", roles: [{role: "dbOwner", db: "unifi-db"}]});
          db.getSiblingDB("unifi-db_stat").createUser({user: "unifi", pwd: "unifi", roles: [{role: "dbOwner", db: "unifi-db_stat"}]});
- name: Create MongoDB service
  kubernetes.core.k8s:
    state: present
    definition:
      apiVersion: v1
      kind: Service
      metadata:
        name: mongodb-svc
        namespace: default
        labels: 
          app: mongodb
      spec:
        type: ClusterIP
        ports:
        - protocol: TCP
          port: 27017
          targetPort: 27017
        selector:
          app: mongodb
- name: Create MongoDB statefull set
  kubernetes.core.k8s:
    state: present
    definition:
      apiVersion: apps/v1
      kind: StatefulSet
      metadata:
        name: mongodb
        namespace: default
        labels:
          app: mongodb
      spec:
        serviceName: mongodb-svc
        replicas: 1
        selector:
          matchLabels:
            app: mongodb
        template:
          metadata:
            labels:
              app: mongodb
          spec:
            containers:
            - name: mongodb
              image: mongo:4.4.18 # Do NOT set it to 'latest'. I'll explain below.
              ports:
              - containerPort: 27017
              env:
              - name: MONGO_INITDB_ROOT_USERNAME
                value: admin
              - name: MONGO_INITDB_ROOT_PASSWORD
                valueFrom:
                  secretKeyRef:
                    name: mongodb-secret
                    key: password
              volumeMounts:
                - name: mongodb-data
                  mountPath: /data/db
        volumeClaimTemplates:
        - metadata:
            name: mongodb-data
          spec:
            accessModes: ["ReadWriteOnce"]
            resources:
              requests:
                storage: 1Gi
- name: Initialize MongoDB
  kubernetes.core.k8s:
    state: present
    definition:
      apiVersion: batch/v1
      kind: Job
      metadata:
        name: mongodb-init-job
        namespace: default
      spec:
        template:
          metadata:
            name: mongodb-init-pod
          spec:
            restartPolicy: OnFailure
            containers:
            - name: mongodb-init-container
              image: mongo:4.4.18 # Again: Do NOT set it to 'latest'.
              command: [ "mongo", "--host", "mongodb-0.mongodb-svc.default.svc.cluster.local", "--authenticationDatabase", "admin", "--username", "admin", "--password", "$(MONGO_INITDB_ROOT_PASSWORD)", "/mongo-init-script/init-script.js" ]
              
