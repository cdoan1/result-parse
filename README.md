# simple-results-parser

```
Render ginkgo results...
Successfully Opened results.xml
✅ management-ingress testing should have the expected running ingress pod in namespace: open-cluster-management (ingress/g0)
✅ management-ingress testing should allow the user to login to web console (ingress/g0)
➖ mongodb testing should have the expected running mongodb pod in namespace: open-cluster-management
✅ mongodb testing should create mongodb test pod (mongodb/g0)
✅ mongodb testing should connect to mongodb server (mongodb/g0)
✅ mongodb testing should insert data to mongodb server (mongodb/g0)
✅ mongodb testing should found data from mongodb server (mongodb/g0)
✅ mongodb testing should remove data from mongodb server (mongodb/g0)
✅ mongodb testing should delete mongodb test pod (mongodb/g0)
✅ Given a hub API should have the expected deployments in open-cluster-management namespace (install/g0)
✅ Given a hub API should have the expected deployments in hive namespace (install/g0)
➖ Given a hub API should be able to provision an AWS cluster using hive (cluster/g5/hive/aws)
➖ Given a hub API should be able to provision a GCP cluster using hive (cluster/g5/hive/gcp)
➖ Given a hub API should be able to provision an Azure cluster using hive (cluster/g5/hive/azure)
✅ Given a hub cluster web console when the user is already authenticated should allow adding multiple baremetal assets (cluster/g0/hive/baremetal/ui/asset)
✅ Given a hub cluster web console when the user is already authenticated should allow adding a provider connection record for baremetal resource (cluster/g0/hive/baremetal/ui/pc)
✅ Given a hub cluster web console when the user is already authenticated should allow provision OCP on a baremetal resource (cluster/g0/hive/baremetal/ui/provision)
➖ Given a hub cluster web console when the user is already authenticated should allow provision OCP on a baremetal resource (cluster/g6/hive/baremetal/ui/provision)
✅ Given a hub cluster web console when the user is already authenticated should have a menu (nav/g0/menu)
✅ Given a hub cluster web console when the user is already authenticated should have have toolbar (nav/g0/toolbar)
✅ Given a hub cluster web console when the user is already authenticated should allow the user to view and create Cloud Connections (cluster/g0)
➖ Manual import cluster Given a list of clusters to import (cluster/g0/manual-import-service-resources)
✅ Given a hub cluster web console should allow the user to login to web console (ui/g0)
✅ Given a hub cluster web console when the user is already authenticated should allow the user to view the Overview (ui/g0)
✅ Given a hub cluster web console when the user is already authenticated should allow the user to view the Cluster Overview (cluster/g0)
✅ Given a hub cluster web console when the user is already authenticated should allow the user to import the local cluster (cluster/g0)
✅ Given a hub cluster web console when the user is already authenticated should allow the user to view the Search Engine (search/g0)
✅ Given a hub cluster web console when the user is already authenticated should allow the user to view the Visual Web Terminal (vwt/g0)
✅ Given a hub cluster web console when the user is already authenticated should allow the user to view Applications (apps/g0)
✅ Given a hub cluster web console when the user is already authenticated should allow the user to view Policies (grc/g0)
✅ Given a hub cluster web console when the user is already authenticated should have a button to create new Policies (grc/g0)
✅ Given a hub cluster web console when the user is already authenticated should be able to create a simple policy (grc/g0)
✅ Given a hub cluster web console when the user is already authenticated should allow the user to view the Topology (topology/g0)
🍺  Summary: Tests:27/6 Failures:0 Errors:0 Time:264.776
🍺  NAME         STATUS   ROLES    AGE   VERSION   INTERNAL-IP       EXTERNAL-IP   OS-IMAGE                                                       KERNEL-VERSION                CONTAINER-RUNTIME
master-0-0   Ready    master   18d   v1.17.1   192.168.123.127   <none>        Red Hat Enterprise Linux CoreOS 44.81.202004130830-0 (Ootpa)   4.18.0-147.8.1.el8_1.x86_64   cri-o://1.17.3-1.rhaos4.4.el8
master-0-1   Ready    master   18d   v1.17.1   192.168.123.139   <none>        Red Hat Enterprise Linux CoreOS 44.81.202004130830-0 (Ootpa)   4.18.0-147.8.1.el8_1.x86_64   cri-o://1.17.3-1.rhaos4.4.el8
master-0-2   Ready    master   18d   v1.17.1   192.168.123.100   <none>        Red Hat Enterprise Linux CoreOS 44.81.202004130830-0 (Ootpa)   4.18.0-147.8.1.el8_1.x86_64   cri-o://1.17.3-1.rhaos4.4.el8
worker-0-0   Ready    worker   18d   v1.17.1   192.168.123.102   <none>        Red Hat Enterprise Linux CoreOS 44.81.202004130830-0 (Ootpa)   4.18.0-147.8.1.el8_1.x86_64   cri-o://1.17.3-1.rhaos4.4.el8
worker-0-1   Ready    worker   18d   v1.17.1   192.168.123.131   <none>        Red Hat Enterprise Linux CoreOS 44.81.202004130830-0 (Ootpa)   4.18.0-147.8.1.el8_1.x86_64   cri-o://1.17.3-1.rhaos4.4.el8

🍺  multiclusterhub version: '1.0.0'
```
