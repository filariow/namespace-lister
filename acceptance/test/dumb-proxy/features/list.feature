Feature: List Namespaces

  Scenario: ServiceAccount list namespace
    Given ServiceAccount has access to "10" namespaces
    Given 10 tenant namespaces exist
    Then the ServiceAccount can retrieve only the namespaces they have access to

  Scenario: ClusterRoleBindings are ignored
    Given the ServiceAccount has Cluster-scoped get permission on namespaces
    Given 10 tenant namespaces exist
    Then the ServiceAccount retrieves no namespaces
