Feature: User registration
  As a new user
  I want to create an account
  So that I can save and access my trips

  Scenario: Successful registration with valid data
    Given I am on the registration page
    When I fill in my name, a valid email and a secure password
    And I submit the registration form
    Then my account should be created
    And I should see a success message

  Scenario: Registration fails with already used email
    Given there is an existing user with email "user@example.com"
    And I am on the registration page
    When I try to register with email "user@example.com"
    Then I should see an error message about the email already being in use
