# PCS-API

### API error code

Here are the different error codes that the API can return

| Code erreur | Description                                                                                                                                                    |
|-------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------|
| 1           | The password must be greater than 8 and less than 128 characters long, at least have an uppercase letter, a lowercase letter, a number and a special character |
| 2           | Invalid email                                                                                                                                                  |
| 3           | Bad typeUser                                                                                                                                                   |
| 4           | Missing content userDTO                                                                                                                                        |
| 5           | Email already exists                                                                                                                                           |
| 6           | Phone already exists                                                                                                                                           |
| 7           | Wrong login or password                                                                                                                                        |
| 8           | invalid token                                                                                                                                                  |
| 9           | Unauthorized chat access                                                                                                                                       |
| 10          | userID uuids in chatDTO are invalid                                                                                                                            |
| 11          | Invalid chat creation                                                                                                                                          |
| 12          | Invalid message in chatDTO                                                                                                                                     |
| 13          | Invalid message creation                                                                                                                                       |


### Application error code

Here are the various error codes that may occur in the event of an error in the program

| Code erreur | Description                                  |
|-------------|----------------------------------------------|
| 1           | Error opening config.env file                |
| 2           | Error when trying to connect to the database |
| 3           | Error convert str to int (env)               |
| 5           | invalid token key                            |