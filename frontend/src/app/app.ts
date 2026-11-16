import { Component, OnInit, inject } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { Student } from './student.model';
import { StudentService } from './student.service';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [FormsModule],
  templateUrl: './app.html',
  styleUrl: './app.scss',
})
export class App implements OnInit {
  private studentService = inject(StudentService);

  students: Student[] = [];
  name = '';
  email = '';
  age: number | null = null;
  message = '';
  error = '';

  ngOnInit(): void {
    this.loadStudents();
  }

  loadStudents(): void {
    this.studentService.getStudents().subscribe({
      next: (data) => { this.students = data; this.error = ''; },
      error: () => { this.error = 'Could not connect to the Go server. Make sure the backend is running.'; },
    });
  }

  addStudent(): void {
    this.message = '';
    this.error = '';
    if (!this.name.trim() || !this.email.trim() || this.age === null || this.age < 1 || this.age > 120) {
      this.error = 'Enter a name, valid email and age between 1 and 120.';
      return;
    }

    this.studentService.addStudent({ name: this.name.trim(), email: this.email.trim(), age: Number(this.age) }).subscribe({
      next: () => {
        this.name = ''; this.email = ''; this.age = null;
        this.message = 'Student added successfully.';
        this.loadStudents();
      },
      error: () => { this.error = 'Unable to add student. Please check the details and try again.'; },
    });
  }

  deleteStudent(id: number): void {
    if (!confirm('Are you sure you want to delete this student?')) return;
    this.studentService.deleteStudent(id).subscribe({
      next: () => { this.message = 'Student deleted successfully.'; this.loadStudents(); },
      error: () => { this.error = 'Unable to delete student.'; },
    });
  }
}
