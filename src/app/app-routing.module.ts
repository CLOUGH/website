import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule, Routes } from '@angular/router';
import { MaintenancePageComponent } from './maintenance-page/maintenance-page.component';

const routes: Routes = [
  { path: 'maintenace', component: MaintenancePageComponent },
  { path: '', pathMatch: 'full', loadChildren: './modules/home/home.module#HomeModule' },
  { path: 'auth', loadChildren: './modules/auth/auth.module#AuthModule' },
  { path: 'blog', loadChildren: './modules/blog/blog.module#BlogModule' },
];
@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [ RouterModule ]
})
export class AppRoutingModule { }
