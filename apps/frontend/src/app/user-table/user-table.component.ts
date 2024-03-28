import { Component, OnInit, OnDestroy } from '@angular/core';
import { RouterModule } from '@angular/router';

import {MatButtonModule} from '@angular/material/button';
import { MatTableDataSource } from '@angular/material/table';
import { MatTableModule } from '@angular/material/table';
import { Subscription } from 'rxjs';
import { UserService } from '../user/user.service';
import { User } from "../user/user";

@Component({
  selector: 'user-table',
  templateUrl: './user-table.component.html',
  styleUrls: ['./user-table.component.scss'],
  imports: [MatTableModule, MatButtonModule, RouterModule],
  providers: [UserService],
  standalone: true
})
export class UserTableComponent implements OnInit, OnDestroy {
  users: MatTableDataSource<User>;
  search: string = "";
  pageNumber: number = 1;
  limit: number = 10;
  displayedColumns: string[] = ['user_name', 'email', 'first_name', 'last_name', 'user_status', 'department', 'id'];
  subscription: Subscription;

  constructor(private userService: UserService) {
    this.users = new MatTableDataSource<User>();
    this.subscription  = new Subscription();
  }

  ngOnInit(): void {
    this.subscription = this.userService.getUsers(this.search, this.limit, this.pageNumber)
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

  deleteUser(id: number): void {
    console.log(id)
    // this.userService.deleteUser(id);
  }
}
