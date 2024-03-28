import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

// Material Ui
import {MatMenuModule} from '@angular/material/menu';
import { HttpClientModule } from '@angular/common/http';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import {MatIconModule} from '@angular/material/icon';
import {MatToolbarModule} from '@angular/material/toolbar';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';

@NgModule({
  declarations: [
    AppComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    NoopAnimationsModule,
    MatButtonModule,
    MatMenuModule,
    HttpClientModule,
    MatInputModule,
    MatFormFieldModule,
    MatToolbarModule,
    MatIconModule
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
