# Go language folder structure for web app/server
[Youtube Link](https://www.youtube.com/watch?v=JYxTIqOWrqI&ab_channel=TheCodingLab)  
  
1. You can place the readme texts here, like setup steps of the project, or any other information about the project that you think would be helpful for other deverlopers/user
2. Create ".env" file, where we will keep all environment related constants, like database username, password, secret key of jwt if any, etc
3. Create the "main.go" file in root, this will serve as the project/web-server entry point
4. Create a "Makefile", this will contain some basic commands to run the project
5. Create these folders:
    1. **Controller**: contains controller that accepts a request and calls a particular service to process it
    2. **Services**: contains service that has the logic for doing all the process like manipulating DB and giving back the desired result
    3. **Models**: contains the struct for storing data in DB or fetching data from DB
    4. **DB**: Keep your DB related operations here. These can be used in services to fetch/update/insert/delete data from DB. *Place DB connection file here as well*
    5. **Routes**: Keep all your routes here and pass the name of controller
    6. **Config**: Keep all your configurations here like fetching variables from .env file, or even db config
    7. **Constants**: Keep all constants here, can also categorize them in separate files.
    8. **Utils**: Keep commonly used functions here, like capitalizing first letter, etc