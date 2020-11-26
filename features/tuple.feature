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

  Scenario: Negating a tuple
    Given a ← tuple(1.0, -2.0, 3.0, -4.0)
    Then -a = tuple(-1.0, 2.0, -3.0, 4.0)

  Scenario: multiply a tuple by a scalar
    Given a ← tuple(1.0, -2.0, 3.0, -4.0)
    Then a * 3.5 = tuple(3.5, -7.0, 10.5, -14.0)

  Scenario: multiply a tuple by a fraction
    Given a ← tuple(1.0, -2.0, 3.0, -4.0)
    Then a * 0.5 = tuple(0.5, -1.0, 1.5, -2.0)

  Scenario: dividing a tuple by a scalar
    Given a ← tuple(1.0, -2.0, 3.0, -4.0)
    Then a / 2.0 = tuple(0.5, -1.0, 1.5, -2.0)

  Scenario: computing the magnitude of vector(1, 0, 0)
    Given v ← vector(1.0, 0.0, 0.0)
    Then magnitude(v) = 1.0


  Scenario: computing the magnitude of vector(1, 2, 3)
    Given v ← vector(1.0, 2.0, 3.0)
    Then magnitude(v) = √14

  Scenario: computing the magnitude of vector(-1, -2, -3)
    Given v ← vector(-1.0, -2.0, -3.0)
    Then magnitude(v) = √14

    Scenario: Normalizing vector(4,0,0) gives (1,0,0)
      Given v ← vector(4.0, 0.0, 0.0)
      Then normalize(v) = vector(1.0, 0.0, 0.0)

  Scenario: Normalizing vector(1,2,3)
    Given v ← vector(1.0, 2.0, 3.0)
    Then normalize(v) = vector(0.26726, 0.53452, 0.80178)

  Scenario: Normalizing vector(1,2,3)
    Given v ← vector(1.0, 2.0, 3.0)
    Then norm ← normalize(v)
    Then magnitude(norm) = 1.0

  Scenario: The dot product of two tuples
    Given a ← vector(1.0, 2.0, 3.0)
    And b ← vector(2.0, 3.0, 4.0)
    Then dot(a, b) = 20.0

  Scenario: The cross product of two tuples
    Given a ← vector(1.0, 2.0, 3.0)
    And b ← vector(2.0, 3.0, 4.0)
    Then cross(a, b) = vector(-1.0, 2.0, -1.0)
    And cross(b, a) = vector(1.0, -2.0, 1.0)

  Scenario: Colors are (red, green, blue) tuples
    Given c ← color(-0.5, 0.4, 1.7)
    Then c.red = -0.5
    And c.green = 0.4
    And c.blue = 1.7

  Scenario: Adding colors
    Given c1 ← color(0.9, 0.6, 0.75)
    And c2 ← color(0.7, 0.1, 0.25)
    Then c1 + c2 = color(1.6, 0.7, 1.0)

  Scenario: subtracting colors
    Given c1 ← color(0.9, 0.6, 0.75)
    And c2 ← color(0.7, 0.1, 0.25)
    Then c1 - c2 = color(0.2, 0.5, 0.5)

  Scenario: multiply a color by a scalar
    Given c1 ← color(0.2, 0.3, 0.4)
    Then c1 * 2.0 = color(0.4, 0.6, 0.8)

  Scenario: multiply colors
    Given c1 ← color(1.0, 0.2, 0.4)
    And c2 ← color(0.9, 1.0, 0.1)
    Then c1 * c2 = color(0.9, 0.2, 0.04)

  Scenario: Reflecting a vector approaching at 45°
    Given v ← vector(1, -1, 0)
    And n ← vector(0, 1, 0)
    When r ← reflect(v, n)
    Then r = vector(1, 1, 0)

  Scenario: Reflecting a vector off a slanted surface
    Given v ← vector(0, -1, 0)
    And n ← vector(√2/2, √2/2, 0)
    When r ← reflect(v, n)
    Then r = vector(1, 0, 0)
