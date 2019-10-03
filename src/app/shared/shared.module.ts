import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MainLayoutComponent } from './components/main-layout/main-layout.component';
import { AdminLayoutComponent } from './components/admin-layout/admin-layout.component';
import { RouterModule } from '@angular/router';
import { MomentModule } from 'ngx-moment';

@NgModule({
  imports: [
    CommonModule,
    RouterModule,
    MomentModule
  ],
  declarations: [MainLayoutComponent, AdminLayoutComponent],
  exports: [MainLayoutComponent, AdminLayoutComponent],
})
export class SharedModule { }
