Banking-API

I am creating a simple api that allows the user to deposit and take out money from their bank account. I am doing this to practise my Golang and RESTful API skills.

Basic goals:

1. Have a basic working GET endpoint that shows user the total amount of money in their bank account
2. Have a basic working POST endpoint that will deposit money into the user's bank account
3. Have a basic working PUT endpoint that allows user to take out money from their bank account


Basic Stretch goals:

1. Validation on the how the sum of money is being entered in ie £100.00 rather than 100
2. Abstraction of the API so the logic isn't all sitting in main
3. Some helpers func 


Advanced Stretch Goals:

1. Set up database so data is persisted
2. Host database and repo on docker, using dockerfile 
3. Think about security requirements and try to counter attacks ie SQL injections


Extra Advanced Stretch Goals:

Have a user login system with a username and password which allows user to access their account. This may require JWTs.



