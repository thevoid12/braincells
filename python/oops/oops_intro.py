# - Everything in Python is an object. An object has a state and behaviors. To create an object, you define a class first. 
#     And then, from the class, you can create one or more objects. The objects are instances of a class.
# - Classes are callable.

print("hello world")
class Person:
    pass # pass acts as a placeholder

# Creating an object of the class
obj = Person()
print(obj)
obj.name="void" # eventhoigh we didn't define name in the class, we can still assign it
# This is because Python allows dynamic attributes
# to be added to objects at runtime.
print(obj.name)

class Person1:
    def __init__(self, name, age):
        self.name = name
        self.age = age
# The __init__ method is a constructor that initializes the object's attributes
# It is called when an object of the class is created.
# self is not a keyword—it's just a convention, but it's strongly recommended to use it.
# It's the first parameter of instance methods in a class.
# It allows you to refer to the current object.
# here name and age are instance attributes ie cannot be accessed without an object of the class. scope is just within the object.

obj2=Person1("void",24)
print(obj2.name)
print(obj2.age)
# The attributes name and age are now defined in the class
# and can be accessed through the object.

# obj3=Person1()
# print(obj3) # This will print the object reference
# print(obj3.name) # This will raise an error because name is not defined

class Person:
    count = 0 # this is a class attribute, shared by all instances of the class
    # Class attributes are defined directly within the class body and the value 
    # are shared across all instances. so ob3 can also access count even obj4 can also access it and so on kinda like global variable within the class scope.
    def __init__(self, name, age):
        self.name = name
        self.age = age

    def greet(self): # this is a method of the class
        return f"Hello, my name is {self.name} and I am {self.age} years old."
# Methods are functions defined inside a class that operate on the object's attributes.

obj3 = Person("Alice", 30)
print(obj3.greet())


class Person:
    count = 0

    def __init__(self, name):
        self.name = name
        Person.count += 1

    @classmethod
    def get_count(cls):
        # Person("void")
        return cls.count
# @classmethod is a method that:
# Belongs to the class, not the instance.
# Takes cls as the first argument (instead of self), which refers to the class itself.
# Can be called on the class itself or an instance, but it does not access or modify instance-specific data.
Person("Alice")
Person("Bob")
print(Person.get_count())  # Output: 2

# static methods are similar to class methods, but they do not take cls or self as the first argument.
# They are used when you want to define a method that does not need to access or modify
# class or instance-specific data. They are defined using the @staticmethod decorator.
# Behaves like a regular function, but is grouped logically inside a class.
# Use it when the method:
# Is related to the class logically, but
# Doesn’t need access to the class (cls) or instance (self).
class TemperatureConverter:
    @staticmethod
    def celsius_to_fahrenheit(c):
        return 9 * c / 5 + 32

    @staticmethod
    def fahrenheit_to_celsius(f):
        return 5 * (f - 32) / 9
    
print(TemperatureConverter.fahrenheit_to_celsius(100))

# inheritence

class Person:
    count = 0
    def __init__(self, name, age):
        self.name = name
        self.age = age

    def greet(self): 
        return f"Hello, my name is {self.name} and I am {self.age} years old."
    
class Employee(Person):
    def __init__(self, name,age, job):
        super().__init__(name,age) # super() is used to call the parent class's constructor. When using super(), you don’t pass self explicitly. Python handles that for you.
        self.job = job
    
    def greet(self):
        return super().greet() +" I work as a {self.job}."

# Inheritance allows a class (child class) to inherit attributes and methods from another class (parent class).
Emp= Employee("void",24,"Software Engineer")
print(Emp.greet())