import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { UserTableComponent } from './user-table/user-table.component'
import { UserModuleComponent } from './user-module/user-module.component';


const routes: Routes = [
  { path: "", component: UserTableComponent },
  { path: "user/:id", component: UserModuleComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
