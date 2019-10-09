import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule, Routes } from '@angular/router';

const routes: Routes = [
  { path: 'login', loadChildren: './modules/login/login.module#LoginModule'},
  {
    path: 'admin', children: [
      {path: '', loadChildren: './modules/admin/admin.module#AdminModule'},
    ]
  },
  { path: 'error', loadChildren: './modules/error/error.module#ErrorModule' },
  { path: '', loadChildren: './modules/page/page.module#PageModule' },
];
@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
