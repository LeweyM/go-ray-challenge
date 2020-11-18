Feature: tuples
  Scenario: a tuple with w=1.0 is a point
    Given a ← tuple(4.3, -4.2, 3.1, 1.0)
    Then a = point(4.3, -4.2, 3.1)

  Scenario: a tuple with w=0.0 is a vector
    Given a ← tuple(4.3, -4.5, 2.1, 0.0)
    Then a = vector(4.3, -4.5, 2.1)

  Scenario: point() creates tuples with w=1
    Given a ← point(4.3, -4.2, 3.1)
    Then a = tuple(4.3, -4.2, 3.1, 1.0)

  Scenario: vector() creates tuples with w=0
    Given a ← vector(4.3, -4.2, 3.1)
    Then a = tuple(4.3, -4.2, 3.1, 0.0)

  Scenario: Adding two tuples
    Given a1 ← tuple(3.0, -2.0, 5.0, 1.0)
    And   a2 ← tuple(-2.0, 3.0, 1.0, 0.0)
    Then a1 + a2 = tuple(1.0, 1.0, 6.0, 1.0)

  Scenario: subtracting two points
    Given p1 ← point(3.0, 2.0, 1.0)
    And   p2 ← point(5.0, 6.0, 7.0)
    Then p1 - p2 = vector(-2.0, -4.0, -6.0)

  Scenario: subtracting a vector from a point
    Given p ← point(3.0, 2.0, 1.0)
    And   v ← vector(5.0, 6.0, 7.0)
    Then p - v = point(-2.0, -4.0, -6.0)

  Scenario: subtracting two vectors
    Given v1 ← vector(3.0, 2.0, 1.0)
    And   v2 ← vector(5.0, 6.0, 7.0)
    Then v1 - v2 = vector(-2.0, -4.0, -6.0)
