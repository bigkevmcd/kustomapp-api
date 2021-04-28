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
      """
