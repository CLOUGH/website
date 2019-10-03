import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { MomentModule } from 'ngx-moment';

import { environment } from './../../environments/environment.prod';
import { MainLayoutComponent } from './components/main-layout/main-layout.component';
import { AdminLayoutComponent } from './components/admin-layout/admin-layout.component';

@NgModule({
  imports: [
    CommonModule,
    RouterModule,
  ],
  declarations: [MainLayoutComponent, AdminLayoutComponent],
  exports: [MainLayoutComponent, AdminLayoutComponent],
})
export class SharedModule { }
