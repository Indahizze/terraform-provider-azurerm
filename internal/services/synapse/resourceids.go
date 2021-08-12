package synapse

//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=FirewallRule -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/firewallRules/firewallRule1
// RoleAssignment cannot be generated at this time
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SparkPool -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/bigDataPools/bigDataPool1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SqlPool -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SqlPoolSecurityAlertPolicy -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1/securityAlertPolicies/Default
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Workspace -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=WorkspaceSecurityAlertPolicy -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/securityAlertPolicies/Default
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=WorkspaceVulnerabilityAssessment -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/vulnerabilityAssessments/default
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=ManagedPrivateEndpoint -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/managedVirtualNetworks/default/managedPrivateEndpoints/endpoint1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=PrivateLinkHub -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/privateLinkHubs/privateLinkHub1
