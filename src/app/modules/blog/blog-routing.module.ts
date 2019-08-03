import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { BlogDetailPageComponent } from './pages/blog-detail-page/blog-detail-page.component';
import { BlogListingsPageComponent } from './pages/blog-listings-page/blog-listings-page.component';


const routes: Routes = [
  { path: '', pathMatch: 'full', component: BlogListingsPageComponent },
  { path: ':id', component: BlogDetailPageComponent }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class BlogRoutingModule { }