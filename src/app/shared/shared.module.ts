import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { MomentModule } from 'ngx-moment';

import { MainLayoutComponent } from './layouts/main-layout/main-layout.component';

@NgModule({
  imports: [
    CommonModule,
    RouterModule,
    FontAwesomeModule
  ],
  declarations: [MainLayoutComponent],
  exports: [MainLayoutComponent]
})
export class SharedModule { }
