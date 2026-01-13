Feature: User login
  As a registered user
  I want to log in with my email and password
  So that I can securely access my trips

  Scenario: Successful login with valid credentials
    Given I am on the login page
    And there is a user with email "user@example.com" and a valid password
    When I enter "user@example.com" and the correct password
    And I submit the login form
    Then I should be logged in
    And I should be redirected to the authenticated home screen

  Scenario: Login fails with invalid password
    Given I am on the login page
    And there is a user with email "user@example.com"
    When I enter "user@example.com" and a wrong password
    And I submit the login form
    Then I should see a generic error message about invalid credentials
    And I should remain on the login page
