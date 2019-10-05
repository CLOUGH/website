import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { ProjectsListingsPageComponent } from './components/projects-listings-page/projects-listings-page.component';
import { PageListingsPageComponent } from './components/page-listings-page/page-listings-page.component';
import { PageEditPageComponent } from './../pages/admin/components/page-edit-page/page-edit-page.component';
import { AdminLayoutComponent } from './layouts/admin-layout/admin-layout.component';
import { DashboardPageComponent } from './components/dashboard-page/dashboard-page.component';
import { PostsListingsPageComponent } from './components/posts-listings-page/posts-listings-page.component';


const routes: Routes = [
  {
    path: '', component: AdminLayoutComponent, children: [
      { path: '', component: DashboardPageComponent },
      { path: 'pages', component: PageListingsPageComponent},
      { path: 'pages/:id', component: PageEditPageComponent},
      { path: 'posts', component: PostsListingsPageComponent},
      { path: 'projects', component: ProjectsListingsPageComponent}
    ]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class AdminRoutingModule { }
