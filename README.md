Here's the updated **README** with the new API endpoint that allows you to check how many users have been created:

---

# 🎉 **Auth Microservices** 🚀

Welcome to the **Auth Microservices**! This project provides secure user authentication and authorization for your applications using JWT (JSON Web Tokens). It also provides endpoints for managing user registration, login, token refresh, protected data access, and optionally checking how many users have been created.

## 🛠 **Features**

- **User Registration** (Sign Up)
- **User Login** (JWT Authentication)
- **Token Refresh** (Get a new access token using the refresh token)
- **Protected Data Access** (Access restricted data with JWT)
- **User List** (Optionally check how many users have been created)

---

## 🚀 **API Endpoints**

### 1. **Login**: `/login` (POST)

Authenticate the user with **email** and **password** to obtain an access token.

**Request Example:**

```bash
curl --location 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "abc9@gmail.com",
    "password": "123465"
}'
```

**Response:**

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFiYzlAZ21haWwuY29tIiwiZXhwIjoxNzMxMDA3MjIwLCJ1c2VyX2lkIjoiNWM4NzZkZmQtNmU5OS00Y2Q3LWI2YmEtZDljOTI0MzhkNDUxIn0.XsW8K1SycOmzgCrSPKfCnNsTzOV6uSZh5dJuSNVmSVs"
}
```

---

### 2. **Signup**: `/signup` (POST)

Create a new user account with **email**, **first name**, **last name**, and **password**.

**Request Example:**

```bash
curl --location 'http://localhost:8080/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "abc9@gmail.com",
    "first_name": "abc",
    "last_name": "xyz",
    "password": "123465"
}'
```

**Response:**

```json
{
    "message": "User created successfully"
}
```

---

### 3. **Refresh Token**: `/auth/refresh` (POST)

Use your **refresh token** to get a **new access token**.

**Request Example:**

```bash
curl --location 'http://localhost:8080/auth/refresh' \
--header 'Content-Type: application/json' \
--data '{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFiYzlAZ21haWwuY29tIiwiZXhwIjoxNzMxMDA3MjIwLCJ1c2VyX2lkIjoiNWM4NzZkZmQtNmU5OS00Y2Q3LWI2YmEtZDljOTI0MzhkNDUxIn0.XsW8K1SycOmzgCrSPKfCnNsTzOV6uSZh5dJuSNVmSVs"
}'
```

**Response:**

```json
{
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFiYzlAZ21haWwuY29tIiwiZXhwIjoxNzMxMDA5MTcwLCJ1c2VyX2lkIjoiNWM4NzZkZmQtNmU5OS00Y2Q3LWI2YmEtZDljOTI0MzhkNDUxIn0.L2YWtDnnXglAYnQQPfIoMelXztGXtWsXDVM0UAw5H7M"
}
```

---

### 4. **Protected Data**: `/protected/secure-data` (GET)

Access protected data by passing a valid **JWT** in the **Authorization** header.

**Request Example:**

```bash
curl --location 'http://localhost:8080/protected/secure-data' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFiYzlAZ21haWwuY29tIiwiZXhwIjoxNzMxMDA3MjIwLCJ1c2VyX2lkIjoiNWM4NzZkZmQtNmU5OS00Y2Q3LWI2YmEtZDljOTI0MzhkNDUxIn0.XsW8K1SycOmzgCrSPKfCnNsTzOV6uSZh5dJuSNVmSVs'
```

**Response:**

```json
{
    "email": "abc9@gmail.com",
    "message": "This is protected data",
    "user_id": "5c876dfd-6e99-4cd7-b6ba-d9c92438d451"
}
```

---

### 5. **Users List**: `/users` (GET) *(Optional)*

Optionally check how many users have been created. This endpoint returns a list of users, including their **first name**, **last name**, and **email**.

**Request Example:**

```bash
curl --location 'http://localhost:8080/users'
```

**Response:**

```json
[
    {
        "first_name": "abc",
        "last_name": "xyz",
        "email": "abc1@gmail.com"
    },
    {
        "first_name": "abc",
        "last_name": "xyz",
        "email": "abc2@gmail.com"
    },
    {
        "first_name": "abc",
        "last_name": "xyz",
        "email": "abc3@gmail.com"
    },
    {
        "first_name": "abc",
        "last_name": "xyz",
        "email": "abc4@gmail.com"
    },
    {
        "first_name": "abc",
        "last_name": "xyz",
        "email": "abc5@gmail.com"
    },
    {
        "first_name": "abc",
        "last_name": "xyz",
        "email": "abc6@gmail.com"
    },
    {
        "first_name": "abc",
        "last_name": "xyz",
        "email": "abc7@gmail.com"
    },
    {
        "first_name": "abc",
        "last_name": "xyz",
        "email": "abc8@gmail.com"
    },
    {
        "first_name": "abc",
        "last_name": "xyz",
        "email": "abc9@gmail.com"
    }
]
```

---

## 🔑 **Response Details**

### Protected Data Response

Here’s a breakdown of the **response** from the `/protected/secure-data` endpoint:

```json
{
    "email": "abc9@gmail.com",
    "message": "This is protected data",
    "user_id": "5c876dfd-6e99-4cd7-b6ba-d9c92438d451"
}
```

- **email**: The email address of the authenticated user.
- **message**: A message confirming access to protected data.
- **user_id**: A unique identifier for the user, assigned after successful registration.

---


## 🐳 **Docker Setup**

You can also run the application in a Docker container using **Docker Compose**.

1. **Build and start the services**:

```bash
docker-compose up --build
```

2. **Access the app** at `http://localhost:8080`.

---

## 💬 **Contributing**

We welcome contributions! Feel free to fork the repo, make changes, and open a pull request.

---

## 📜 **License**

This project is licensed under the **MIT License**.

---

**Enjoy building your app with secure authentication!** 🎉🚀

---

This updated README includes the new `/users` endpoint that allows you to check the list of users created in the system. Let me know if you need further modifications!