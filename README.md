# Discussion API

# library requirements
- gin
- gorm
- godotenv

# Doc
**Show Users**
----
  Returns json data about a single user.

* **URL**

  /users

* **Method:**

  `GET`
  
*  **URL Params**
 
   `None`

* **Data Params**

  `None`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** 
    ```json
      "data": {
        "ID": 2,
        "CreatedAt": "2020-09-06T10:45:34.262118314+07:00",
        "UpdatedAt": "2020-09-06T10:45:34.262118314+07:00",
        "DeletedAt": {
            "Time": "0001-01-01T00:00:00Z",
            "Valid": false
        },
        "FirstName": "ucok",
        "LastName": "ganteng",
        "PhotoProfile": "random"
    }```
 
* **Error Response:**

* **Sample Call:**

  ```javascript
    $.ajax({
      url: "/users",
      dataType: "json",
      type : "GET",
      success : function(r) {
        console.log(r);
      }
    });
  ```

  **Show Users**
----
  Returns json data about a single user.

* **URL**

  /users/

* **Method:**

  `GET`
  
*  **URL Params**

   **Required:**
 
   `id=[integer]`

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{ id : 12, name : "Michael Bloom" }`
 
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "User doesn't exist" }`

  OR

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ error : "You are unauthorized to make this request." }`

* **Sample Call:**

  ```javascript
    $.ajax({
      url: "/user/1",
      dataType: "json",
      type : "GET",
      success : function(r) {
        console.log(r);
      }
    });
  ```  
<!-- **Show User**
----
  Returns json data about a single user.

* **URL**

  /user/:id

* **Method:**

  `GET`
  
*  **URL Params**

   **Required:**
 
   `id=[integer]`

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** 
    ```json 
        "data": {
        "ID": 2,
        "CreatedAt": "2020-09-06T10:45:34.262118314+07:00",
        "UpdatedAt": "2020-09-06T10:45:34.262118314+07:00",
        "DeletedAt": {
            "Time": "0001-01-01T00:00:00Z",
            "Valid": false
        },
        "FirstName": "ucok",
        "LastName": "ganteng",
        "PhotoProfile": "random"
    }
    ```
 
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "not found" }`

* **Sample Call:**

  ```javascript
    $.ajax({
      url: "/users/1",
      dataType: "json",
      type : "GET",
      success : function(r) {
        console.log(r);
      }
    });
  ``` -->

**Create User**
----
  Returns json data about a single user.

* **URL**

  /user/

* **Method:**

  `POST`
  
*  **URL Params**

   None

* **Data Params**

  - first_name : string
  - last_name : string
  - photo_profile : string

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** 
    ```json 
        "data": {
        "ID": 2,
        "CreatedAt": "2020-09-06T10:45:34.262118314+07:00",
        "UpdatedAt": "2020-09-06T10:45:34.262118314+07:00",
        "DeletedAt": {
            "Time": "0001-01-01T00:00:00Z",
            "Valid": false
        },
        "FirstName": "ucok",
        "LastName": "ganteng",
        "PhotoProfile": "random"
    }
    ```
 
* **Error Response:**

  * **Code:** 500 Internal Server Error <br />

* **Sample Call:**

  ```javascript
    $.ajax({
      url: "/user",
      dataType: "json",
      type : "POST",
      success : function(r) {
        console.log(r);
      }
    });
  ```