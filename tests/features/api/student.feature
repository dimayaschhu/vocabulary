Feature: Це апішка для менеджерів
  Сервер запускаєтся на 8080 порті
  Даним інструментом менеджери додають учнів та слова

  Scenario: успішне створення учня
    When I send "POST" request to "/student/create" with JSON body:
      """json
          {"name":"test student","lesson": 1,"chatId": 123}
      """
    Then method should return status code 200 and JSON response:
    """json
      {"name":  "test student"}
    """
    Then I see next records in "students" table:
      | _id   | name         | lesson | chatId |
      | <_id> | test student | 1      | 123    |

  Scenario: дублікат при створенні учня
    Then The next fixtures exist in "students" table:
      | _id   | name         | lesson | chatId |
      | <_id> | test student | 1      | 123    |
    When I send "POST" request to "/student/create" with JSON body:
      """json
          {"name":"test student","lesson": 1,"chatId": 123}
      """
    Then method should return status code 400 and JSON response:
    """json
      {"error":  "exist student.Name: test student"}
    """

#  Scenario: не коректні дані при створення учня