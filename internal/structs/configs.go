package structs

// iac-ansible options
// setup_dependencies: true
// create_cluster: true
// install_argo: true
// install_argo_apps: true
// remove_kubeadmin: true
//
//iac-ansible variables
// Image version should match argo-containers/argocd-repo-server-sops/Dockerfile

// type Argo_Image struct {
// 	argocd_image_version string `yaml:"argocd_image_version"`
// }

type Env struct {
	env string `yaml:"platform"`
}

type clusters struct {
	envs []string
}

type Gitops_Vars struct {
	region               string `yaml:"region"`
	branch_tag           string `yaml:"branch"`
	organization         string `yaml:"organization"`
	github_url           string `yaml:"github_url"`
	gitops_repo_url      string `yaml:"gitops_repo_url"`
	argocd_image_version string `yaml:"argocd_image_version"`
}

// # AWS SecretsManager vars
// # Secret Locations for nested secrets
type Ocp_secrets struct {
	ocp_secrets                string `yaml:"ocp_secrets"`
	cw_secrets                 string `yaml:"cw_secrets"`
	ocp_install_secrets        string `yaml:"ocp_install_secrets"`
	credentials_mode           string `yaml:"credentials_mode"`
	pull_secret                string `yaml:"pull_secret"`
	ocp_install_aws_access_key string `yaml:"ocp_install_aws_access_key"`
	ocp_install_aws_secret_key string `yaml:"ocp_install_aws_secret_key"`
}

// # Day 1
// # gitops-operator
type Day_One_Vars struct {
	gitops_enabled          bool   `yaml:"gitops_enabled"`
	argocd_pgp_private_key  string `yaml:"argocd_pgp_private_key"`
	git_user                string `yaml:"git_user"`
	git_token               string `yaml:"git_token"`
	gitops_namespace        string `yaml:"gitops_namespace"`
	gitops_admin_passphrase string `yaml:"gitops_admin_passphrase"`
}

// # infra-nodes
type InfraNode struct {
	infra_nodes_enabled bool
	infra               struct {
		ec2_type bool
		ebs      struct {
			encrypted  bool
			iops       int
			volumesize int
			volumetype string
		}
		rootvolume struct {
			iops           int
			size           int
			rootVolumeType string `yaml:"type"`
		}
	}
}

// master-nodes
type MasterNode struct {
	master struct {
		ec2_type   string
		rootvolume struct {
			iops     int
			size     int
			nodeType string
		}
		replicas int
	}
}

// storage-nodes
type StorageNode struct {
	storage_node_enabled bool
	storage              struct {
		ec2_type string
		ebs      struct {
			encrypted  bool
			iops       int
			volumesize int
			volumetype string
		}
		rootvolume struct {
			iops           int
			size           int
			rootVolumeType string
		}
		replicas int
	}
}

// worker-nodes
type WorkerNode struct {
	worker_node_enabled bool
	worker              struct {
		ec2_type string
		ebs      struct {
			encrypted  bool
			iops       int
			volumesize int
			volumetype string
		}
		rootvolume struct {
			iops           int
			size           int
			rootVolumeType string
		}
		replicas int
	}
}

type Nodes struct {
	node string
	role string
}

// machine-health-checks
type HealthChecks struct {
	machine_health_checks_enabled bool
	machinehealthchecks           []Nodes
}
type allowedRegistryDomains struct {
	domainName string
	insecure   bool
}

// manual-remediations
type Redmediations struct {
	manual_remediations_enabled bool
	allowedRegistriesForImport  struct {
		domains           []allowedRegistryDomains
		allowedRegistries []string
	}
}

// Openshift Data Foundation
//Note: Dependant on storage_node_enabled
// type DataFoundations struct {
// 	odf_enabled bool
// }

// // Disable Sample Operators
// type SampleOperators struct {
// 	disable_sample_operators bool
// }

// //Day 2
// // Ploigos Trusted Software Supply Chain
// type tssc struct {
// 	tssc_enabled bool
// }

// // amq-broker
// type amq struct {
// 	amq_broker_enabled bool
// }

// AAP Secrets
type aap struct {
	aap_enabled  bool
	aap_host     string
	aap_username string
	aap_password string
}

// Certs
type certs struct {
	custom_certs_enabled bool
	api_cert             string
	api_cert_key         string
	console_cert         string
	console_cert_key     string
	ingress_cert         string
	ingress_cert_key     string
	ingress_cert_ca      string
	ca_cert              string
}

// cluster-autoscaling
type autoscaling struct {
	cluster_autoscaling_enabled bool
	pod_priority_threshold      int
	resource_limits             struct {
		max_nodes_total int
		cores           struct {
			min int
			max int
		}
		memory struct {
			min int
			max int
		}
	}
	scaledown struct {
		enabled             bool
		delay_after_add     string
		delay_after_delete  string
		delay_after_failure string
		unneeded_time       string
	}
}

//container-security-operator
// container_security_operator: true

// # custom-console
// custom_console: true

// # disable-self-provisioner
// disable_self_provisioner: true

// # file-integrity-operator
// file_integrity_operator: true

type repo struct {
	repo string
}

// gatekeeper-operator
type gatekeeper struct {
	gatekeeper_operator bool
	gk_img_tag          string
	allowed_git_repos   []repo
}

// Cloud Watch Logging
type cloudwatch struct {
	cw_log_forwarding_enabled bool
	cloudwatch_aws_access_key string
	cloudwatch_aws_secret_key string
}

// Oauth Secrets
type oauth struct {
	oauth struct {
		groupsync_enabled     bool
		group_membership_attr string
		basedn                string
		user_name_attr        string
		user_uid_attr         string
		sync_schedule         string
		ldap_enabled          bool
		binddn                string
		ldapurl               string
		ldap_password         string
		radiant_ldab_enabled  bool
		radiant_binddn        string
		radiant_url           string
		radiant_password      string
		okta_oath_enabled     bool
		okta_client_id        string
		okta_client_secret    string
		okta_issuer           string
	}
}

// # openshift-compliance-operator
// openshift_compliance_operator: true

// Cluster Logging
type cluster_logging struct {
	openshift_logging struct {
		enabled           bool
		logging_namespace string
		fluentd_loglevel  string
		// How long to keep application logs
		application_logs_max_age string
		// How long to keep infrastructure logs
		infra_logs_max_age string
		// How long to keep audit logs
		audit_logs_max_age string
		// Number of elasticsearch instances
		elasticsearch_node_count int
		// Storage class for elasticsearch persistent volumes
		elasticsearch_storage_class string
		// Size of each elasticsearch volume
		elasticsearch_storage_size string
		// Number of kibana replicas
		kibana_replicas int
		// Curator log pruning schedule
		curator_schedule string
		// Method for installation plan approval
		install_plan_approval string
	}
}

//Openshift Monitoring
//openshift_monitoring_enabled: true

// ACM Vars
type acm struct {
	acm_enabled  bool
	mgmt_cluster bool
	cluster_set  string
	acm_channel  string
}

// New Relic Secrets
type new_relic struct {
	new_relic_enabled          bool
	new_relic_license_key      string
	new_relic_pixie_deploy_key string
	new_relic_pixie_api_key    string
}

// Docker Registry Secrets
// artifactory_pull_secret: "{{ ocp_secrets.artifactory_pull_secret }}"
// svc_acct: "{{ lookup('aws_secret', 'svc-acct', region=vars['region']) }}"
// svc_acct_name: "{{ svc_acct.account_name }}"
// svc_acct_password: "{{ svc_acct.account_password }}"

// Pager Duty Secrets
type pager_duty struct {
	pager_duty_enabled       bool
	pagerduty_events_key     string
	pagerduty_prometheus_key string
}

// ACS Secrets
type acs struct {
	acs_enabled         bool
	acs_central_cluster bool
	acs_init_bundle     string
}

// Fluentbit Splunk Log Forwarding
type fluent_bit struct {
	fluentbit_enabled bool
	splunk_hec_token  string
	splunk_channel_id string
	cloud_account     string
	fluent_bit        struct {
		namespace      string
		cm_name        string
		secret_name    string
		splunk_hec_url string
		image          struct {
			repository string
			tag        string
			pullPolicy string
		}
		rbac struct {
			nodeAccess bool
		}
		podSecurityPolicy struct {
			create bool
		}
		openShift struct {
			enabled                    bool
			securityContextConstraints struct {
				create bool
			}
		}
		securityContext struct {
			capabilities struct {
				drop                   []string
				readOnlyRootFilesystem bool
				runAsNonRoot           bool
				runAsUser              bool
				privileged             bool
			}
		}
	}
}

// # OADP
// # GLOBAL- region, organization, cluster_name
// oadp_enabled: true

// # Code Ready Workspace
// code_ready_workspace_enabled: false

// # Serverless
// serverless_enabled: false
