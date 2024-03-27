import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { MatFormFieldModule  } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';


import { UserService } from '../user/user.service';

@Component({
  selector: 'user-module',
  templateUrl: './user-module.component.html',
  styleUrls: ['./user-module.component.scss'],
  imports: [MatFormFieldModule, MatInputModule, ReactiveFormsModule ],
  providers: [UserService],
  standalone: true
})
export class UserModuleComponent implements OnInit {
  userId: string | null;
  userForm = new FormGroup({
    user_name: new FormControl(''),
    last_name: new FormControl(''),
    email: new FormControl(''),
    user_status: new FormControl(''),
    department: new FormControl('')
  })

  constructor(private route: ActivatedRoute, userService: UserService) {
    this.userId = null
  }

  ngOnInit(): void {
    this.route.paramMap.subscribe(params => {
      this.userId = params.get('id');
    });
  }

  onSubmit(): void {

  }
}
