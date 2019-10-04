import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { MomentModule } from 'ngx-moment';
import { JwtHelperService } from '@auth0/angular-jwt';

import { environment } from './../../environments/environment.prod';
import { MainLayoutComponent } from './components/main-layout/main-layout.component';
import { AuthService } from './services/auth.service';

@NgModule({
  imports: [
    CommonModule,
    RouterModule,
  ],
  declarations: [MainLayoutComponent],
  exports: [MainLayoutComponent]
})
export class SharedModule { }
