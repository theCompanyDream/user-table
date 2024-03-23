import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

import { UserService } from '../user/user.service';

@Component({
  selector: 'user-module',
  templateUrl: './user-module.component.html',
  styleUrls: ['./user-module.component.scss'],
  providers: [UserService]
})
export class UserModuleComponent implements OnInit {
  userId: string | null;

  constructor(private route: ActivatedRoute) {
    this.userId = null
  }

  ngOnInit(): void {
    this.route.paramMap.subscribe(params => {
      this.userId = params.get('id');
    });
  }
}
