STUDENT MANAGEMENT APPLICATION (Go + Angular)

A simplified implementation of the Go + Angular intern coding assessment. It includes only the requested student form, student table, add/list/delete APIs, JSON responses and basic validation.

No database, login, charts, Docker or PostgreSQL is used. Students are held in a Go slice in memory. Restarting the backend clears the list.

REQUIREMENTS
- Go installed (the provided Go 1.27.x is suitable)
- Node.js 22.16.0 or a compatible version
- npm
- Angular CLI 21.2.x (already installed on the target laptop)

RUN BACKEND
1. Open PowerShell in the backend folder.
2. Run: go run .
3. Keep this terminal open. API runs at http://localhost:8080.

RUN FRONTEND
1. Open a second PowerShell in the frontend folder.
2. Run: npm install
3. Run: npm start
4. Open http://localhost:4200 in your browser.

API ENDPOINTS
- GET    http://localhost:8080/api/students       (retrieve all students)
- POST   http://localhost:8080/api/students       (add a student)
- DELETE http://localhost:8080/api/students/{id}  (delete by ID)

POST JSON example
{
  "name": "Asha Kumar",
  "email": "asha@example.com",
  "age": 21
}
