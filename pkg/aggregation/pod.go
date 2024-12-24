package aggregation

import "regexp"

var podAggregationRegexps = map[int]*regexp.Regexp{}
var podAggregationLabelValues = map[int]string{}
var podAggregations = map[string]string{
	"The node had condition.*":                                      "The node had condition",
	"Nameserver limits were exceeded.*":                             "Nameserver limits were exceeded",
	"Successfully pulled image .*":                                  "Successfully pulled image",
	"Successfully assigned .*":                                      "Successfully assigned",
	"Pulling image .*":                                              "Pulling image",
	"Container image .* already present on machine$":                "Container image already present on machine",
	"Failed to pull image .*":                                       "Failed to pull image",
	".*Failed to apply default image tag.*":                         "Failed to apply default image tag",
	".*Failed to inspect image.*":                                   "Failed to inspect image",
	"Container image .* is not present with pull policy of Never.*": "Container image is not present with pull policy of Never",
	".*error determining status: .*":                                "Error determining status",
	"Back-off restarting failed container .* in pod .*":             "Back-off restarting failed container in pod",
	"Readiness probe (failed|errored).*":                            "Readiness probe failed",
	"Startup probe (failed|errored).*":                              "Startup probe failed",
	"Liveness probe (failed|errored).*":                             "Liveness probe failed",
	"Readiness probe warning.*":                                     "Readiness probe warning",
	"Startup probe warning.*":                                       "Startup probe warning",
	"Liveness probe warning.*":                                      "Liveness probe warning",
	"Back-off pulling image.*":                                      "Back-off pulling image",
	"Back-off restarting failed container .*":                       "Back-off restarting failed container",
	"(Created|Started) container .*":                                "Created or started container",
	"\\d+/\\d+ nodes are available: .* Insufficient cpu.*":          "Incufficient cpu",
	"\\d+/\\d+ nodes are available: \\d+ node\\(s\\) didn't match Pod's node affinity/selector.*": "Nodes did not match node affinity/selector",
	"Stopping container .*":                                                     "Stopping container",
	"MountVolume.SetUp failed for volume .*":                                    "Volume setup is failed",
	"\\d+/\\d+ nodes are available: \\d+ node\\(s\\) had untolerated taint.*":   "Nodes had untolerated taint",
	"\\d+/\\d+ nodes are available: persistentvolumeclaim .* not found.*":       "PersistentVolumeClaim is not found",
	"\\d+/\\d+ nodes are available: persistentvolumeclaim .* is being deleted":  "PersistentVolumeClaim is being deleted",
	"Internal PreCreateContainer hook failed: .*":                               "Internal PreCreateContainer hook failed",
	"Internal PreStartContainer hook failed: .*":                                "Internal PreStartContainer hook failed",
	"Container .* exceeded its local ephemeral storage limit .*":                "Container exceeded its local ephemeral storage limit",
	"Pod ephemeral local storage usage exceeds the total limit of containers.*": "Pod ephemeral local storage usage exceeds the total limit of containers",
	".*ephemeral volume.*":                                                      "Ephemeral volume binding is failed",
	"Usage of EmptyDir volume.* exceeds the limit.*":                            "Usage of EmptyDir volume exceeds the limit",
	"Error: .*": "Error occurred",
	"Init container .* failed startup probe.*":                                                            "Init container failed startup probe",
	"Init container .* failed liveness probe.*":                                                           "Init container failed liveness probe",
	"Init container is in .* state, try killing it before restart":                                        "Try killing init container before restart",
	"Container is in .* state, try killing it before restart":                                             "Try killing container before restart",
	"Container .* resize requires restart":                                                                "Container resize requires restart",
	"[Ee]rror killing pod: .*":                                                                            "Error killing pod",
	"[Uu]nable to ensure pod container exists: .*":                                                        "Unable to ensure pod container exists",
	"error making pod data directories: .*":                                                               "Error making pod data directories",
	"network is not ready.*":                                                                              "Network is not ready",
	"Error validating pod .* from %s due to duplicate pod name .*, ignoring.*":                            "Error validating pod due to duplicate pod name, ignoring",
	"Pod .* matches multiple PodDisruptionBudgets.*":                                                      "Pod matches multiple PodDisruptionBudgets",
	"PodResourceClaim.*":                                                                                  "Failed resource claim creation",
	".*request to HTTPS lifecycle hook.* got HTTP response, retry with HTTP succeeded":                    "Request to HTTPS lifecycle hook got HTTP response, retry with HTTP succeeded",
	"Search Line limits were exceeded, some search paths have been omitted, the applied search line is.*": "Search Line limits were exceeded, some search paths have been omitted",
	"skip schedule deleting pod.*":                                                                        "Skip schedule deleting pod",
	"Volume is already used by pod(s).*":                                                                  "Volume is already used by pod(s)",
	"AttachVolume.Attach succeeded.* for volume.*":                                                        "AttachVolume.Attach succeeded for volume",
	"AttachVolume.Attach failed for volume.*":                                                             "AttachVolume.Attach failed for volume",
	"Unable to retrieve some image pull secrets.*":                                                        "Unable to retrieve some image pull secrets",
	"Kubelet detected an IPv.* node IP .*, but the cloud provider selected an IPv.* node IP.*":            "Kubelet detected wrong IP family",
	".*pod has unbound immediate PersistentVolumeClaims":                                                  "Pod has unbound immediate PersistentVolumeClaims",
	"created SCC ranges for .* namespace.*":                                                               "created SCC ranges for namespace",
	"Multi-Attach error for volume .* Volume is already used.*":                                           "Multi-Attach error for volume: volume is already used",
	"Multi-Attach error for volume .* Volume is already exclusively attached to one node":                 "Multi-Attach error for volume: volume is already exclusively attached to one node",
	"Add eth\\d+ .* from .*":                                                                              "Add eth network",
}

var podDisruptionBudgetAggregationRegexps = map[int]*regexp.Regexp{}
var podDisruptionBudgetAggregationLabelValues = map[int]string{}
var podDisruptionBudgetAggregations = map[string]string{
	"Failed to get pods.*":                                                "Failed to get pods",
	"Failed to calculate the number of expected pods.*":                   "Failed to calculate the number of expected pods",
	"Pods selected by this PodDisruptionBudget (selector: .*) were found": "Found unmanaged pods associated with this PDB",
}
