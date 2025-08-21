# random interesting info
- String objects are immutable, which means that once created, their values cannot be changed. The
- by default use double for decimals. if you need to use float you need to add f in the end like 1.234f
- ```java
    String s; // the default value of s is NULL not ""

    ```

- In Java SE 7 and later, any number of underscore characters (_) can appear anywhere between digits in a numerical literal. This feature enables you, for example. to separate groups of digits in numeric literals, which can improve the readability of your code.

    ```java
        long creditCardNumber = 1234_5678_9012_3456L; // L to denote its a long
    ```
- length of the aray 
```java
anArray.length
```
- infinite for loop
```java
// infinite loop
for ( ; ; ) {

    // your code goes here
}
```
- labled break:In Java, a labeled break lets you break out of outer loops, not just the nearest enclosing loop. This is useful when you're inside nested loops and want to exit more than one level at once.
```java
outer:
for (int i = 0; i < 5; i++) {
    for (int j = 0; j < 5; j++) {
        if (i * j > 6) {
            break outer;  // Breaks out of BOTH loops
        }
        System.out.println("i=" + i + ", j=" + j);
    }
}
System.out.println("Finished");
```
- The last branching statement is the yield statement. The yield statement exits from the current switch expression it is in. A yield statement is always followed by an expression that must produce a value. This expression must not be void. The value of this expression is the value produced by the enclosing switch expression.
- In Java, switch statements fall through unless each case ends with break, return, or throw.
```java
int month = 8;
List<String> futureMonths = new ArrayList<>();

switch (month) {
    case 1:  futureMonths.add("January");
    case 2:  futureMonths.add("February");
    case 3:  futureMonths.add("March");
    case 4:  futureMonths.add("April");
    case 5:  futureMonths.add("May");
    case 6:  futureMonths.add("June");
    case 7:  futureMonths.add("July");
    case 8:  futureMonths.add("August");
    case 9:  futureMonths.add("September");
    case 10: futureMonths.add("October");
    case 11: futureMonths.add("November");
    case 12: futureMonths.add("December");
             break;
    default: break;
}
System.out.println(futureMonths);
// op:[August, September, October, November, December]
```
- because of this they brought a new syntax called switch expression (->). if we use it it will allow fall through so no need breaks to manage and It will exit the loop once that case is run
```java
public String convertToLabel(int quarter) {
    String quarterLabel =
        switch (quarter) {
            case 0  -> {
                System.out.println("Q1 - Winter");
                yield "Q1 - Winter";
            }
            default -> "Unknown quarter";
        };
    }
    return quarterLabel;
}
```

- The selector variable of a switch statement can be an object, so this object can be null. You should protect your code from null selector variables, because in this case the switch statement will throw a NullPointerException.
- as per convention class names start with capital letter
- constructors they use the name of the class and have no return type.
- Without static → A feature exists only inside an object you create.
With static → The feature exists once for the whole class, and you can use it without creating an object.