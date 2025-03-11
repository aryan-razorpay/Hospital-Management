# Hospital Management System

## 📌 Project Overview

The **Hospital Management System** is a RESTful API built using **Golang (Gin framework), MySQL, and GORM**. It allows hospitals to manage doctors and patients efficiently, including adding, updating, and retrieving records from a database.

---

## 🛠 Tech Stack

- **Golang** (Gin Framework)
- **MySQL** (Database)
- **GORM** (ORM for Golang)

---

## ⚡ Features

✅ Add and retrieve doctors and patients\
✅ Assign doctors to patients\
✅ Update contact details\
✅ Fetch all patients assigned to a doctor\
✅ RESTful API architecture

---

## 📂 Folder Structure

```
hospital-management/
│── cmd/                 # Entry point of the application
│   └── main.go          # Main application file
│
├── config/              # Configuration settings
│   ├── main.go        # Database connection
│
├── models/              # Database models
│   ├── doctor.go        # Doctor schema
│   ├── patient.go       # Patient schema
│
├── methods/        # Data access layer
│   ├── doctor.methods.go   # Doctor DB operations
│   ├── patient.emthods.go  # Patient DB operations
│
├── mutex/            # Mutexing
│   ├── main.go
│
├── controllers/         # API handlers
│   ├── doctor.controller.go
│   ├── patient.controller.go
│
├── routes/              # API route definitions
│   ├── routes.go
│
├── utils/               # Helper functions
│   ├── response.go      # Standardized API responses
│   ├── validators.go    # Input validation
│
├── .gitignore
│
├── go.mod               # Go module dependencies
├── go.sum              # Checksum for dependencies
│
└── Readme.md 

```

---

## 🚀 Getting Started

### 1️⃣ Clone the repository

```sh
git clone https://github.com/aryan-razorpay/Hospital-Management.git
cd Hospital-Management
```

### 2️⃣ Install dependencies

```sh
go mod tidy
```

### 3️⃣ Set up MySQL database

1. Create a database in MySQL:
   ```sql
   CREATE DATABASE hospital_db;
   ```
2. Update **config/config.go** with your MySQL credentials:
   ```go
   dsn := "root:your_password@tcp(localhost:3306)/hospital_db?charset=utf8mb4&parseTime=True&loc=Local"
   ```

### 4️⃣ Run the application

```sh
go run cmd/main.go
```

Server runs at [**http://localhost:8080**](http://localhost:8080) 🚀

---

## 📌 API Endpoints

### 🔹 **Doctor Routes**

- **Create Doctor** → `POST /doctor/`
- **Get Doctor by ID** → `GET /doctor/:id`
- **Update Doctor Contact** → `PATCH /doctor/:id`

### 🔹 **Patient Routes**

- **Create Patient** → `POST /patient/`
- **Get Patient by ID** → `GET /patient/:id`
- **Update Patient Details** → `PATCH /patient/:id`
- **Get Patients by Doctor ID** → `GET /fetchPatientByDoctorId/:doctorId`

---


## 💡 Contribution Guidelines

1. Fork the repository 🍴
2. Create a feature branch (`git checkout -b feature-branch`)
3. Commit changes (`git commit -m 'Added new feature'`)
4. Push to the branch (`git push origin feature-branch`)
5. Submit a Pull Request ✅

---

**Developed by **[**Aryan Panchal**](https://github.com/aryan-razorpay)** 🚀**

