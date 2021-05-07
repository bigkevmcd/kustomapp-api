Feature: Application initialisation
  As a Project Lead
  In order to start deploying our application via GitOps
  I need to create a new directory structure for deploying my environments

  Scenario: Repository initialisation
    Given a temporary directory in the environment
    When I run successfully "./kapp init --output KAPP_TEMP --env dev --env staging"
    Then I should get the message "app initialised"
    And a tree of files should be generated
      """
      bases/kustomization.yaml
      dev/bases/kustomization.yaml
      staging/bases/kustomization.yaml
      """

  Scenario: Adding a target to an environment
    Given a temporary directory in the environment
    When I run successfully "./kapp init --output KAPP_TEMP --env dev --env staging"
    And I run successfully "./kapp target add --dir KAPP_TEMP --env dev --name eu-west-2"
    Then I should get the message "target added"
    And a tree of files should be generated
      """
      bases/kustomization.yaml
      dev/bases/kustomization.yaml
      staging/bases/kustomization.yaml
      staging/eu-west-2/kustomization.yaml
      """
