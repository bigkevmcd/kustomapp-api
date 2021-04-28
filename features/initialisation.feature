Feature: Application initialisation
  As a Project Lead
  In order to start deploying our application via GitOps
  I need to create a new directory structure for deploying my environments

  Scenario: Repository initialisation
    Given a temporary directory in the environment
    When I run successfully "./kapp init --path $KAPP_TEMP --environments dev,staging"
    Then I should get the message "app initialised"
    And a tree of files should be generated
      """
├── bases
│   ├── kustomization.yaml
├── dev
│   ├── bases
│   │   └── kustomization.yaml
└── staging
    └── bases
        └── kustomization.yaml
      """
