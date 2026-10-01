---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1159625-getting-permission-error-from-unifi-network-application-when-c-f27f78a9-2
title: "questions-1159625-getting-permission-error-from-unifi-network-application-when-c-f27f78a9"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1159625-getting-permission-error-from-unifi-network-application-when-c-f27f78a9.md
source_anchor: ""
source_lines: [115, 338]
sha256: c3eae4ba93ef101a689f950d791bdfceb019d45b91f1150c8b29b84500c8c1ea
---

# questions-1159625-getting-permission-error-from-unifi-network-application-when-c-f27f78a9

              env:
              - name: MONGO_INITDB_ROOT_PASSWORD
                valueFrom:
                  secretKeyRef:
                    name: mongodb-secret
                    key: password
              volumeMounts:
              - name: mongo-init-script
                mountPath: /mongo-init-script
            volumes:
            - name: mongo-init-script
              configMap:
                name: mongodb-init-script
The section works fine.
The output from kubectl logs mongodb-init-job-f5f54 is as follows:
MongoDB shell version v4.4.18
connecting to: mongodb://mongodb-0.mongodb-svc.default.svc.cluster.local:27017/?authSource=admin&compressors=disabled&gssapiServiceName=mongodb
Implicit session: session { "id" : UUID("39ad2f51-07a9-4e67-8c5b-68f6fff92872") }
MongoDB server version: 4.4.18
Successfully added user: {
        "user" : "unifi",
        "roles" : [
                {
                        "role" : "dbOwner",
                        "db" : "unifi-db"
                }
        ]
}
Successfully added user: {
        "user" : "unifi",
        "roles" : [
                {
                        "role" : "dbOwner",
                        "db" : "unifi-db_stat"
                }
        ]
}
Which means I could login as admin and create a MongoDB user called unifi with password unifi that has access to the database unifi-db.
I also know that it is accessible via DNS by setting the database host to mongodb-0.mongodb-svc.default.svc.cluster.local using port 27017.
The reason why I chose to install version 4.4.18 of MongoDB is because even though Unifi Network Application supports up to version 7.0, it is not the same case with MongoDB support for the Raspberry Pi boards.
MongoDB can only run up til version 4.4.18 on Raspberry Pi.
Next up is getting Unifi Network Application to run on my cluster server.
I am using the following tasks to do this:
---
- name: Define storage space for Unifi
  kubernetes.core.k8s:
    state: present
    definition:
      apiVersion: v1
      kind: PersistentVolumeClaim
      metadata:
        name: unifi-cluster-pvc
        namespace: default
      spec:
        accessModes:
          - ReadWriteOnce
        storageClassName: longhorn
        resources:
          requests:
            storage: 5Gi
- name: Add unifi deployment
  kubernetes.core.k8s:
    state: present
    definition:
      apiVersion: apps/v1
      kind: Deployment
      metadata:
        name: unifi
        namespace: default
        labels:
          app: unifi
      spec:
        replicas: 1
        selector:
          matchLabels:
            app: unifi
        template:
          metadata:
            labels:
              app: unifi
          spec:
            volumes:
            - name: unifi-config
              persistentVolumeClaim:
                claimName: unifi-cluster-pvc
            containers:
            - name: unifi
              image: lscr.io/linuxserver/unifi-network-application:latest
              ports:
                - containerPort: 3478
                  protocol: UDP
                - containerPort: 10001
                  protocol: UDP
                - containerPort: 5514
                  protocol: UDP
                - containerPort: 8080
                - containerPort: 8443
                - containerPort: 8843
                - containerPort: 8880
                - containerPort: 6789
              volumeMounts:
              - name: unifi-config
                mountPath: /config
              env:
                - name: PUID
                  value: "1000"
                - name: GUID
                  value: "1000"
                - name: MONGO_USER
                  value: "unifi"
                - name: MONGO_PASS
                  value: "unifi"
                - name: MONGO_HOST
                  value: "mongodb-0.mongodb-svc.default.svc.cluster.local"
                - name: MONGO_PORT
                  value: "27017"
                - name: MONGO_DBNAME
                  value: "unifi-db"
                - name: MONGO_TLS
                  value: "false"
- name: Define unifi service ports
  kubernetes.core.k8s:
    state: present
    definition:
      apiVersion: v1
      kind: Service
      metadata:
        name: unifi
        namespace: default
      spec:
        type: LoadBalancer
        selector:
          app: unifi
        ports:
          - name: "8080"
            port: 8080
            targetPort: 8080
          - name: "8443"
            port: 8443
            targetPort: 8443
          - name: "8843"
            port: 8843
            targetPort: 8843
          - name: "8880"
            port: 8880
            targetPort: 8880
          - name: "6789"
            port: 6789
            targetPort: 6789
          - name: "3478"
            port: 3478
            protocol: UDP
            targetPort: 3478
          - name: "10001"
            port: 10001
            protocol: UDP
            targetPort: 10001
          - name: "5514"
            port: 5514
            protocol: UDP
            targetPort: 5514
- name: Add unifi ingress
  kubernetes.core.k8s:
    state: present
    definition:
      apiVersion: networking.k8s.io/v1
      kind: Ingress
      metadata:
        name: unifi
        namespace: default
        annotations:
          nginx.ingress.kubernetes.io/backend-protocol: HTTPS
          cert-manager.io/cluster-issuer: letsencrypt-staging
          cert-manager.io/acme-challenge-type: dns01
      spec:
        ingressClassName: nginx
        tls:
        - hosts:
          - unifi.example.com
          secretName: unifi-tls
        rules:
        - host: 'unifi.example.com'
          http:
            paths:
            - path: /
              pathType: Prefix
              backend:
                service:
                  name: unifi
                  port:
                    number: 8443
Running the command kubectl get pods gives me the following output:
NAME                     READY   STATUS      RESTARTS   AGE
mongodb-0                1/1     Running     0          64m
unifi-5dcdbfb6d9-8wbbp   1/1     Running     0          64m
mongodb-init-job-f5f54   0/1     Completed   2          64m
And I can login to the unifi-db database with the command: kubectl exec mongodb-0 -it -- mongo -u unifi -p unifi unifi-db
The command: show collections
Gives me amongst others the following:
account
admin
alarm
crashlog
dashboard
device
...
So far so good.
However I cannot get the website for Unifi Network Application up and running.
The output from the command kubectl log unifi-5dcdbfb6d9-8wbbp gives amongst others this message:
Error creating bean with name 'statDbService' 
defined in com.ubnt.service.DatabaseSpringContext: 
Command failed with error 13 (Unauthorized): 'not 
authorized on unifi-db_stat to execute command { 
listCollections: 1, cursor: {}, nameOnly: true, $db: 
"unifi-db_stat", lsid: { id: UUID("99b4d07f-b3f3-49c9-9979-dbdd27445881") 
} }' on server mongodb-0.mongodb-svc.default.svc.cluster.local:27017. 
The full response is {"ok": 0.0, "errmsg": "not authorized on unifi-db_stat 
to execute command { listCollections: 1, cursor: {}, nameOnly: true, $db: 
\"unifi-db_stat\", lsid: { id: UUID(\"99b4d07f-b3f3-49c9-9979-dbdd27445881\") 
} }", "code": 13, "codeName": "Unauthorized"}
How do I resolve this issue?
Reading the docs for how role assignment works in MongoDB, I can read that setting the role of a user to dbOwner is basically an all-access pass, since it combines the roles of readWrite, dbAdmin and userAdmin into one role.
One of the permissions that you get is listCollections, which is referenced in the errorcode above.
So what is going on?
