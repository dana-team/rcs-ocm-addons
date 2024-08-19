# rcs-ocm-addons

![Version: 0.0.0](https://img.shields.io/badge/Version-0.0.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: latest](https://img.shields.io/badge/AppVersion-latest-informational?style=flat-square)

A Helm chart for the rcs-score-addon

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` | Node affinity rules for scheduling pods. Allows you to specify advanced node selection constraints. |
| deploymentConfig | object | `{"agentInstallNamespace":"open-cluster-management-agent-addon","customizedVariables":[{"name":"MAX_CPU_COUNT","value":"1"},{"name":"MIN_CPU_COUNT","value":"0"},{"name":"MAX_MEMORY_BYTES","value":"104857"},{"name":"MIN_MEMORY_BYTES","value":"0"}],"name":"rcs-score-deploy-config","namespace":"open-cluster-management-hub"}` | Configurations for the AddonDeploymentConfig object |
| deploymentConfig.agentInstallNamespace | string | `"open-cluster-management-agent-addon"` | Namespace where the agent addon is installed |
| deploymentConfig.customizedVariables | list | `[{"name":"MAX_CPU_COUNT","value":"1"},{"name":"MIN_CPU_COUNT","value":"0"},{"name":"MAX_MEMORY_BYTES","value":"104857"},{"name":"MIN_MEMORY_BYTES","value":"0"}]` | Customzied variables for the addon |
| deploymentConfig.name | string | `"rcs-score-deploy-config"` | Name of the AddonDeploymentConfig |
| deploymentConfig.namespace | string | `"open-cluster-management-hub"` | Namespace of the AddonDeploymentConfig |
| fullnameOverride | string | `""` |  |
| image.pullPolicy | string | `"IfNotPresent"` | The pull policy for the image. |
| image.repository | string | `"ghcr.io/dana-team/rcs-ocm-addons"` | The repository of the manager container image. |
| image.tag | string | `""` | The tag of the manager container image. |
| installStrategy.score.placements[0].configs[0].group | string | `"addon.open-cluster-management.io"` |  |
| installStrategy.score.placements[0].configs[0].name | string | `"rcs-score-deploy-config"` |  |
| installStrategy.score.placements[0].configs[0].namespace | string | `"open-cluster-management-hu"` |  |
| installStrategy.score.placements[0].configs[0].resource | string | `"addondeploymentconfigs"` |  |
| installStrategy.score.placements[0].name | string | `"all-clusters"` |  |
| installStrategy.score.placements[0].namespace | string | `"test"` |  |
| installStrategy.score.placements[0].rolloutStrategy.type | string | `"All"` |  |
| installStrategy.score.type | string | `"Placements"` |  |
| installStrategy.status.placements[0].name | string | `"all-clusters"` |  |
| installStrategy.status.placements[0].namespace | string | `"test"` |  |
| installStrategy.status.placements[0].rolloutStrategy.type | string | `"All"` |  |
| installStrategy.status.type | string | `"Placements"` |  |
| manager | object | `{"resources":{"limits":{"cpu":"500m","memory":"128Mi"},"requests":{"cpu":"10m","memory":"64Mi"}},"score":{"args":["manager"],"command":["/score-addon"]},"securityContext":{"allowPrivilegeEscalation":false,"capabilities":{"drop":["ALL"]}},"status":{"args":["manager"],"command":["/status-addon"]}}` | Configuration for the manager container. |
| manager.resources | object | `{"limits":{"cpu":"500m","memory":"128Mi"},"requests":{"cpu":"10m","memory":"64Mi"}}` | Resource requests and limits for the manager container. |
| manager.score.args | list | `["manager"]` | Command-line arguments passed to the score-addon manager container. |
| manager.score.command | list | `["/score-addon"]` | Command-line commands passed to the score-addon manager container. |
| manager.securityContext | object | `{"allowPrivilegeEscalation":false,"capabilities":{"drop":["ALL"]}}` | Security settings for the manager container. |
| manager.status.args | list | `["manager"]` | Command-line arguments passed to the status-adon manager container. |
| manager.status.command | list | `["/status-addon"]` | Command-line commands passed to the status-addon manager container. |
| nameOverride | string | `""` |  |
| nodeSelector | object | `{}` | Node selector for scheduling pods. Allows you to specify node labels for pod assignment. |
| replicaCount | int | `1` | The number of replicas for the deployment. |
| tolerations | list | `[]` | Node tolerations for scheduling pods. Allows the pods to be scheduled on nodes with matching taints. |

