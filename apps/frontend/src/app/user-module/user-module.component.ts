import { Component, OnInit, OnDestroy } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Location } from '@angular/common';
import { FormControl, FormGroup, ReactiveFormsModule, Validators  } from '@angular/forms';
import { MatFormFieldModule  } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { Subscription } from 'rxjs';
import { UserService } from '../user/user.service';
import { User } from '../user/user'; // Assuming you have a User model

@Component({
  selector: 'user-module',
  templateUrl: './user-module.component.html',
  styleUrls: ['./user-module.component.scss'],
  imports: [MatFormFieldModule, MatInputModule, ReactiveFormsModule ],
  providers: [UserService],
  standalone: true
})
export class UserModuleComponent implements OnInit, OnDestroy {
  userId: string | null;
  subscription: Subscription;
  userForm = new FormGroup({
    user_name: new FormControl('', [Validators.required, Validators.minLength(3), Validators.maxLength(50)]),
    last_name: new FormControl('', [Validators.required, Validators.minLength(3), Validators.maxLength(255)]),
    first_name: new FormControl('', [Validators.required, Validators.minLength(3), Validators.maxLength(255)]),
    email: new FormControl('', [Validators.required, Validators.minLength(3), Validators.maxLength(50)]),
    user_status: new FormControl('',[Validators.required, Validators.maxLength(1), Validators.pattern(/^[IAT]$/)]),
    department: new FormControl('')
  })

  constructor(
    private route: ActivatedRoute,
    private userService: UserService,
    private location: Location) {
    this.userId = null;
    this.subscription = new Subscription();
  }

  ngOnInit(): void {
    this.route.paramMap.subscribe(params => {
      this.userId = params.get('id');
      if (this.userId) {
        this.fetchUser(this.userId);
      }
    });
  }

  fetchUser(id: string): void {
    this.subscription = this.userService.getUser(id).subscribe(
      (user: User) => {
        this.userForm.patchValue({
          user_name: user.user_name,
          last_name: user.last_name,
          first_name: user.first_name,
          email: user.email,
          user_status: user.user_status,
          department: user.department
        });
      }
    );
  }

  ngOnDestroy(): void {
    if (this.subscription) {
      this.subscription.unsubscribe();
    }
  }

  onSubmit(): void {
    const user_name = this.userForm.get("user_name")?.value ?? "";
    const first_name = this.userForm.get("first_name")?.value ?? "";
    const last_name = this.userForm.get("last_name")?.value ?? "";
    const user_status = this.userForm.get("user_status")?.value ?? "";
    const department = this.userForm.get("department")?.value ?? "";
    const email = this.userForm.get("email")?.value ?? "";


    const user = new User(
      user_name,
      first_name,
      last_name,
      email,
      user_status,
      department
    );

    if (this.userId) {
      this.subscription = this.userService.updateUser(this.userId, user)
        .subscribe(
          (user: User) => {
            this.userForm.patchValue({
              user_name: user.user_name,
              last_name: user.last_name,
              first_name: user.first_name,
              email: user.email,
              user_status: user.user_status,
              department: user.department
            });
            this.location.replaceState(`/user/${user.id}`)
            this.userId = user.id || ''
          }
        )
    } else {
      this.subscription = this.userService.createUser(user)
        .subscribe(
          (user: User) => {
            this.userForm.patchValue({
              user_name: user.user_name,
              last_name: user.last_name,
              first_name: user.first_name,
              email: user.email,
              user_status: user.user_status,
              department: user.department
            });
            this.location.replaceState(`/user/${user.id}`)
            this.userId = user.id || ''
          }
        )
    }
  }
}
