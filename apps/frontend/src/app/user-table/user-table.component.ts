import { Component, OnInit, OnDestroy } from '@angular/core';
import { RouterModule } from '@angular/router';

import {MatButtonModule} from '@angular/material/button';
import { MatInputModule } from '@angular/material/input';
import { MatTableDataSource } from '@angular/material/table';
import { MatTableModule } from '@angular/material/table';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { MatFormFieldModule  } from '@angular/material/form-field';
import {MatPaginatorModule} from '@angular/material/paginator';
import { Subscription } from 'rxjs';
import { UserService } from '../user/user.service';
import { User } from "../user/user";

@Component({
  selector: 'user-table',
  templateUrl: './user-table.component.html',
  styleUrls: ['./user-table.component.scss'],
  imports: [MatTableModule, MatButtonModule, RouterModule, MatPaginatorModule, MatFormFieldModule, MatInputModule, ReactiveFormsModule],
  providers: [UserService],
  standalone: true
})
export class UserTableComponent implements OnInit, OnDestroy {
  users: MatTableDataSource<User>;
  userForm = new FormGroup({
    search: new FormControl(''),
    pageNumber: new FormControl(1),
    limit: new FormControl(10)
  })
  displayedColumns: string[] = ['user_name', 'email', 'first_name', 'last_name', 'user_status', 'department', 'id'];
  subscription: Subscription;

  constructor(private userService: UserService) {
    this.users = new MatTableDataSource<User>();
    this.subscription  = new Subscription();
  }

  ngOnInit(): void {
    this.getQuery(null)
  }

  getQuery(event: any): void {
    console.log("I'm running")
    const search = this.userForm.get('search')?.value || '';
    const limit = this.userForm.get('limit')?.value || 20;
    const pageNumber = this.userForm.get('pageNumber')?.value || 1;

    this.subscription = this.userService.getUsers(search, limit, pageNumber)
      .subscribe(
        (payload) => {
          this.users.data = payload;
        }
      );
  }

  ngOnDestroy(): void {
    if (this.subscription) {
      this.subscription.unsubscribe();
    }
  }

  deleteUser(userId: string): void {
    this.userService.deleteUser(userId).subscribe({
      next: (response) => {
        // Remove the item from the array upon successful deletion
        const newData = this.users.data.filter(user => user.id !== userId);
        this.users.data = newData
      },
      error: (error) => {
        // Handle error
        console.error('Error deleting item', error);
      }
    })
  }
}
