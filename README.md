# rcs-ocm-addons

This repository is complementary to the [`rcs-ocm-deployer`](https://github.com/dana-team/rcs-ocm-deployer) and serves as a repo to host all the `OCM addons` for the project.

## Status AddOn

By applying this add-on to the Hub Cluster, the `Capp` status will automatically be synced back from Managed/Spoke Clusters to the Hub cluster. It has two components:
- `agent` - deployed on the Managed/Spoke clusters; responsible for syncing the `Capp` status between the Managed Cluster and the Hub Cluster.
- `manager` - deployed on the Hub Cluster; responsible for deploying the `agent` on the Managed/Spoke clusters.

## Score AddOn

Based on the [resource-usage-collect-addon](https://github.com/open-cluster-management-io/addon-contrib/tree/main/resource-usage-collect-addon), this add-on implements an `AddonPlacementScore` which gives a score to each Managed Cluster. The score is in the range of [-100, 100] and is computed from collecting resource usage information (CPU Request and Memory Request).

## Getting Started

Check the guide on the `rcs-ocm-deployer` [repo](https://github.com/mzeevi/rcs-ocm-deployer/tree/main?tab=readme-ov-file#deploy-the-add-ons) for information about deploying the `AddOns` on an `OCM` cluster.

Deploy using the Helm chart located at `charts/rcs-ocm-addons`:

```bash
$ helm upgrade --install rcs-ocm-addons --namespace open-cluster-management --create-namespace oci://ghcr.io/dana-team/helm-charts/rcs-ocm-addons --version <release>
```

### Build your own image

```bash
$ make docker-build docker-push IMG=<registry>/rcs-ocm-addons:<tag>
```
