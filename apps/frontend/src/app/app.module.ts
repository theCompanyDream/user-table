import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

// Material Ui
import { MatSlideToggleModule } from '@angular/material/slide-toggle';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';

// User Stuff
import { UserModuleComponent } from './user-module/user-module.component';
import { UserService } from './user/user.service';
import { UserTableComponent } from './user-table/user-table.component'

@NgModule({
  declarations: [
    AppComponent,
    UserModuleComponent,
    UserTableComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    NoopAnimationsModule,
    MatSlideToggleModule
  ],
  providers: [UserService],
  bootstrap: [AppComponent]
})
export class AppModule { }
