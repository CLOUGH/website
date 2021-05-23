
import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { ProjectsListingsPageComponent } from './pages/projects-listings-page/projects-listings-page.component';
import { PageListingsPageComponent } from './pages/page-listings-page/page-listings-page.component';
import { PageEditPageComponent } from './pages/page-edit-page/page-edit-page.component';
import { AdminLayoutComponent } from './layouts/admin-layout/admin-layout.component';
import { DashboardPageComponent } from './pages/dashboard-page/dashboard-page.component';
import { PostsListingsPageComponent } from './pages/posts-listings-page/posts-listings-page.component';
import { PageCreatePageComponent } from './pages/page-create-page/page-create-page.component';


const routes: Routes = [
  {
    path: '', component: AdminLayoutComponent, children: [
      { path: '', component: DashboardPageComponent },
      { path: 'pages', component: PageListingsPageComponent},
      { path: 'pages/create', component: PageCreatePageComponent},
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
