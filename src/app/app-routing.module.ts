import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule, Routes } from '@angular/router';
import { AboutPageComponent } from './pages/about-page/about-page.component';

const routes: Routes = [
  { path: '', pathMatch: 'full', loadChildren: './modules/home/home.module#HomeModule' },
  { path: 'blog', loadChildren: './modules/blog/blog.module#BlogModule' },
  { path: 'about', component: AboutPageComponent },
  { path: 'admin', pathMatch: 'full', loadChildren: './modules/admin/admin.module#AdminModule' },
];
@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
