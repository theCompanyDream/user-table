import { Component, OnInit } from '@angular/core';

// User Stuff
import {MatTableModule} from '@angular/material/table';
import { UserService } from '../user/user.service';
import {User} from "../user/user"

@Component({
  selector: 'user-table',
  templateUrl: './user-table.component.html',
  styleUrls: ['./user-table.component.scss'],
  imports: [MatTableModule],
  providers: [UserService],
  standalone: true
})
export class UserTableComponent implements OnInit {
  users: User[]
  search: string
  pageNumber: number
  limit: number

  constructor(private userService: UserService) {
    this.users = []
    this.search = ""
    this.limit = 10
    this.pageNumber = 1
  }

  ngOnInit(): void {
    this.userService.getUsers(this.search, this.limit, this.pageNumber)
      .subscribe(payload => this.users = payload)
  }

}
