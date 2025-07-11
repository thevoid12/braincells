# django

### additional resources for indepth understanding
- url dispatcher:
https://docs.djangoproject.com/en/5.2/topics/http/urls/

### install
uv venv .venv
source .venv/bin/activate
uv pip install django

## to activate virutal environment
source .venv/bin/activate

## to deactivate
deactivate

- in django we have manage.py which is similar to our main exe in golang. you can run a bunch of terminal commands on manage.py like creating a package etc t
- 
```py 
django-admin startproject mysite ./ 
```
 this gives us base template
- 
  ```py 
  python manage.py startapp polls 
  ```
  we are creating a app called polls
- django works on dry principle- do not repeat yourself
- 
    ```py
    python manage.py migrate 
    ``` 
The migrate command looks at the INSTALLED_APPS setting and creates any necessary database tables according to the database settings in your mysite/settings.py file and the database migrations shipped with the app
- Django apps are “pluggable”: You can use an app in multiple projects, and you can distribute apps, because they don’t have to be tied to a given Django installation.
- 
  ```py
   python manage.py makemigrations polls 
  ```
  creates the migration file using the model in polls.(link the model in polls into settings.py)
-
    ```py
    python manage.py sqlmigrate polls 0001
    ```
here polls is the app name 0001 is the migration craeted by makemigration. The sqlmigrate command doesn’t actually run the migration on your database - instead, it prints it to the screen so that you can see what SQL Django thinks is required. It’s useful for checking what Django is going to do or if you have database administrators who require SQL scripts for changes.

-
    ```py
    python manage.py check
    ```
this checks for any problems in your project without making migrations or touching the database.
-
    ```py
    python manage.py migrate
    ```
 this The migrate command takes all the migrations that haven’t been applied (Django tracks which ones are applied using a special table in your database called django_migrations) and runs them against your database - essentially, synchronizing the changes you made to your models with the schema in the database.

- so this is a 3 step process
  - Change your models (in models.py).

  -  Run python manage.py makemigrations to create migrations for those changes
    - Run python manage.py migrate to apply those changes to the database.
- QuerySet (filer/exclude) objects are lazy – the act of creating a QuerySet doesn’t involve any database activity. You can stack filters together all day long, and Django won’t actually run the query until the QuerySet is evaluated.
  - field__lookuptype=value. 
- views are the handlers. All Django wants is that HttpResponse. Or an exception in the views. we can do whatever we want with views.
-  A URLconf maps URL patterns to views.
-  to store html templates store it in a folder called templates which django will look at it by default for html templates. eg if the app name is polls, create a folder inside polls called template and again polls and then index.html ie polls/templates/polls/index.html. this is the convention
-  

 