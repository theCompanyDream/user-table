import { Component, OnInit } from '@angular/core';
import { RouterModule } from '@angular/router';
import { MatButtonModule } from '@angular/material/button';
import { MatInputModule } from '@angular/material/input';
import { MatTableDataSource } from '@angular/material/table';
import { MatTableModule } from '@angular/material/table';
import { Subscription } from 'rxjs';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatPaginatorModule } from '@angular/material/paginator';
import { UserService } from '../user/user.service';
import { User } from "../user/user";

@Component({
  selector: 'user-table',
  templateUrl: './user-table.component.html',
  styleUrls: ['./user-table.component.scss'],
  imports: [MatTableModule, MatButtonModule, RouterModule, MatPaginatorModule, MatFormFieldModule, MatInputModule, ReactiveFormsModule, FormsModule],
  providers: [UserService],
  standalone: true
})
export class UserTableComponent implements OnInit {
  users = new MatTableDataSource<User>();
  search: string;
  pageNumber: number;
  limit: number;
  length: number;
  subscription: Subscription;
  displayedColumns: string[] = ['user_name', 'email', 'first_name', 'last_name', 'user_status', 'department', 'id'];

  constructor(private userService: UserService) {
    this.subscription = new Subscription();
    this.pageNumber = 1;
    this.search = "";
    this.limit = 10;
    this.length = 0;
  }

  ngOnInit(): void {
    // Initial load
    this.getQuery(null);
  }

  getQuery(formValues: any): void {
    if (formValues != null ) {
      this.limit = formValues.pageSize
      this.pageNumber = formValues.pageIndex
    }
    this.subscription = this.userService.getUsers(this.search, this.limit, this.pageNumber)
      .subscribe(
        (payload) => {
          this.users.data = payload.users;
          this.length = payload.length;
        }
      );
  }

  searchQuery(event: any): void {
    this.subscription = this.userService.getUsers(this.search, this.limit, this.pageNumber)
      .subscribe(
        (payload) => {
          this.users.data = payload.users;
          this.length = payload.length;
        }
      );
  }

  deleteUser(userId: string): void {
    this.userService.deleteUser(userId).subscribe({
      next: (response) => {
        // Remove the item from the array upon successful deletion
        const newData = this.users.data.filter(user => user.id !== userId);
        this.users.data = newData
        this.length--
      },
      error: (error) => {
        // Handle error
        console.error('Error deleting item', error);
      }
    })
  }
}
