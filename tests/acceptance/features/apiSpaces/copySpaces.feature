@api @skipOnOcV10
Feature: copy file
  As a user
  I want to be able to copy files
  So that I can manage my files

  Background:
    Given these users have been created with default attributes and without skeleton files:
      | username |
      | Alice    |
      | Brian    |

  Scenario Outline: Copying a file within a same space project with role manager and editor
    Given the administrator has given "Alice" the role "Space Admin" using the settings api
    And user "Alice" has created a space "Project" with the default quota using the GraphApi
    And user "Alice" has created a folder "newfolder" in space "Project"
    And user "Alice" has uploaded a file inside space "Project" with content "some content" to "insideSpace.txt"
    And user "Alice" has shared a space "Project" to user "Brian" with role "<role>"
    When user "Brian" copies file "insideSpace.txt" to "newfolder/insideSpace.txt" in space "Project" using the WebDAV API
    And the HTTP status code should be "201"
    And for user "Alice" the content of the file "newfolder/insideSpace.txt" of the space "Project" should be "some content"
    Examples:
      | role    |
      | manager |
      | editor  |


  Scenario: Copying a file within a same space project with role viewer
    Given the administrator has given "Alice" the role "Space Admin" using the settings api
    And user "Alice" has created a space "Project" with the default quota using the GraphApi
    And user "Alice" has created a folder "newfolder" in space "Project"
    And user "Alice" has uploaded a file inside space "Project" with content "some content" to "insideSpace.txt"
    And user "Alice" has shared a space "Project" to user "Brian" with role "viewer"
    When user "Brian" copies file "insideSpace.txt" to "newfolder/insideSpace.txt" in space "Project" using the WebDAV API
    And the HTTP status code should be "403"


  Scenario Outline: User copy a file from a space project with different role to a space project with manager role
    Given the administrator has given "Brian" the role "Space Admin" using the settings api
    And user "Brian" has created a space "Project1" with the default quota using the GraphApi
    And user "Brian" has created a space "Project2" with the default quota using the GraphApi
    And user "Brian" has uploaded a file inside space "Project1" with content "Project1 content" to "project1.txt"
    And user "Brian" has shared a space "Project2" to user "Alice" with role "manager"
    And user "Brian" has shared a space "Project1" to user "Alice" with role "<role>"
    When user "Alice" copies file "project1.txt" from space "Project1" to "project1.txt" in space "Project2" using the WebDAV API
    And the HTTP status code should be "201"
    And for user "Alice" the content of the file "project1.txt" of the space "Project2" should be "Project1 content"
    Examples:
      | role    |
      | manager |
      | editor  |
      | viewer  |


  Scenario Outline: User copy a file from a space project with different role to a space project with editor role
    Given the administrator has given "Brian" the role "Space Admin" using the settings api
    And user "Brian" has created a space "Project1" with the default quota using the GraphApi
    And user "Brian" has created a space "Project2" with the default quota using the GraphApi
    And user "Brian" has uploaded a file inside space "Project1" with content "Project1 content" to "project1.txt"
    And user "Brian" has shared a space "Project2" to user "Alice" with role "editor"
    And user "Brian" has shared a space "Project1" to user "Alice" with role "<role>"
    When user "Alice" copies file "project1.txt" from space "Project1" to "project1.txt" in space "Project2" using the WebDAV API
    And the HTTP status code should be "201"
    And for user "Alice" the content of the file "project1.txt" of the space "Project2" should be "Project1 content"
    Examples:
      | role    |
      | manager |
      | editor  |
      | viewer  |

  Scenario Outline: User copy a file from a space project with different role to a space project with editor role
    Given the administrator has given "Brian" the role "Space Admin" using the settings api
    And user "Brian" has created a space "Project1" with the default quota using the GraphApi
    And user "Brian" has created a space "Project2" with the default quota using the GraphApi
    And user "Brian" has uploaded a file inside space "Project1" with content "Project1 content" to "project1.txt"
    And user "Brian" has shared a space "Project2" to user "Alice" with role "viewer"
    And user "Brian" has shared a space "Project1" to user "Alice" with role "<role>"
    When user "Alice" copies file "project1.txt" from space "Project1" to "project1.txt" in space "Project2" using the WebDAV API
    And the HTTP status code should be "403"
    Examples:
      | role    |
      | manager |
      | editor  |


  Scenario Outline: User copy a file from space project to personal with different role
    Given the administrator has given "Brian" the role "Space Admin" using the settings api
    And user "Brian" has created a space "Project" with the default quota using the GraphApi
    And user "Brian" has uploaded a file inside space "Project" with content "Project content" to "project.txt"
    And user "Brian" has shared a space "Project" to user "Alice" with role "<role>"
    When user "Alice" copies file "project.txt" from space "Project" to "project.txt" inside personal using the WebDAV API
    And the HTTP status code should be "201"
    And as "Alice" file "/project.txt" should exist
    And the content of file "project.txt" for user "Alice" should be "Project content"
    Examples:
      | role    |
      | manager |
      | editor  |
      | viewer  |
