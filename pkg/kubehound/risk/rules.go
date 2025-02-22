package risk

var CriticalRoleMap = map[string]bool{
	"admin":                       true,
	"cluster-admin":               true,
	"edit":                        true,
	"kindnet":                     true,
	"kubeadm:get-nodes":           true,
	"local-path-provisioner-role": true,
	"system:aggregate-to-admin":   true,
	"system:aggregate-to-edit":    true,
	"system:aggregate-to-view":    true,
	"system:certificates.k8s.io:certificatesigningrequests:nodeclient":     true,
	"system:certificates.k8s.io:certificatesigningrequests:selfnodeclient": true,
	"system:certificates.k8s.io:kube-apiserver-client-approver":            true,
	"system:certificates.k8s.io:kube-apiserver-client-kubelet-approver":    true,
	"system:certificates.k8s.io:kubelet-serving-approver":                  true,
	"system:certificates.k8s.io:legacy-unknown-approver":                   true,
	"system:controller:attachdetach-controller":                            true,
	"system:controller:bootstrap-signer":                                   true,
	"system:controller:certificate-controller":                             true,
	"system:controller:clusterrole-aggregation-controller":                 true,
	"system:controller:cronjob-controller":                                 true,
	"system:controller:daemon-set-controller":                              true,
	"system:controller:deployment-controller":                              true,
	"system:controller:disruption-controller":                              true,
	"system:controller:endpoint-controller":                                true,
	"system:controller:endpointslice-controller":                           true,
	"system:controller:endpointslicemirroring-controller":                  true,
	"system:controller:ephemeral-volume-controller":                        true,
	"system:controller:expand-controller":                                  true,
	"system:controller:generic-garbage-collector":                          true,
	"system:controller:horizontal-pod-autoscaler":                          true,
	"system:controller:job-controller":                                     true,
	"system:controller:namespace-controller":                               true,
	"system:controller:node-controller":                                    true,
	"system:controller:persistent-volume-binder":                           true,
	"system:controller:pod-garbage-collector":                              true,
	"system:controller:pv-protection-controller":                           true,
	"system:controller:pvc-protection-controller":                          true,
	"system:controller:replicaset-controller":                              true,
	"system:controller:replication-controller":                             true,
	"system:controller:resourcequota-controller":                           true,
	"system:controller:root-ca-cert-publisher":                             true,
	"system:controller:route-controller":                                   true,
	"system:controller:service-account-controller":                         true,
	"system:controller:service-controller":                                 true,
	"system:controller:statefulset-controller":                             true,
	"system:controller:token-cleaner":                                      true,
	"system:controller:ttl-after-finished-controller":                      true,
	"system:controller:ttl-controller":                                     true,
	"system:coredns":                                                       true,
	"system:heapster":                                                      true,
	"system:kube-aggregator":                                               true,
	"system:kube-controller-manager":                                       true,
	"system:kube-dns":                                                      true,
	"system:kube-scheduler":                                                true,
	"system:kubelet-api-admin":                                             true,
	"system:monitoring":                                                    true,
	"system:node-bootstrapper":                                             true,
	"system:node-problem-detector":                                         true,
	"system:node-proxier":                                                  true,
	"system:persistent-volume-provisioner":                                 true,
	"system:public-info-viewer ":                                           true,
	"system:service-account-issuer-discovery":                              true,
	"system:volume-scheduler":                                              true,
	"view":                                                                 true,
}

// TODO identities map
