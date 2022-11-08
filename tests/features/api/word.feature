Feature: Це апішка для менеджерів
  Сервер запускаєтся на 8080 порті
  Даним інструментом менеджери додають учнів та слова

  Scenario: успішне створення слова
    When I send "POST" request to "/word/create" with JSON body:
      """json
          {"name":"test word","lesson": 1,"translate": "translate"}
      """
    Then method should return status code 200 and JSON response:
    """json
      {"name":  "test word"}
    """
    Then I see next records in "words" table:
      | _id   | name      | lesson | translate |
      | <_id> | test word | 1      | translate |

  Scenario: дублікат при створенні слова
    Then The next fixtures exist in "words" table:
       | name      | translate | lesson |
       | test word | translate | 1      |
    When I send "POST" request to "/word/create" with JSON body:
      """json
          {"name":"test word","lesson": 1,"translate": "translate"}
      """
    Then method should return status code 400 and JSON response:
    """json
      {"error":  "exist word.Name: test word"}
    """

#  Scenario:
#  Scenario: успішне створення слова
#  Scenario: не коректні дані при створення слова
#  Scenario: дублікат при створенні слово
#  Scenario: не коректні дані при створення учня