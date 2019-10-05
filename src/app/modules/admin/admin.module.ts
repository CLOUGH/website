import { PageEditPageComponent } from './../pages/admin/components/page-edit-page/page-edit-page.component';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';

import { AdminRoutingModule } from './admin-routing.module';
import { AdminLeftSideNavComponent } from './layouts/admin-left-side-nav/admin-left-side-nav.component';
import { AdminLayoutComponent } from './layouts/admin-layout/admin-layout.component';
import { DashboardPageComponent } from './components/dashboard-page/dashboard-page.component';
import { PageListingsPageComponent } from './components/page-listings-page/page-listings-page.component';
import { PostsListingsPageComponent } from './components/posts-listings-page/posts-listings-page.component';
import { ProjectsListingsPageComponent } from './components/projects-listings-page/projects-listings-page.component';


@NgModule({
  declarations: [
    AdminLayoutComponent,
    AdminLeftSideNavComponent,
    DashboardPageComponent,
    PageListingsPageComponent,
    PostsListingsPageComponent,
    ProjectsListingsPageComponent,
    PageEditPageComponent
  ], imports: [
    CommonModule,
    AdminRoutingModule,
    FontAwesomeModule
  ]
})
export class AdminModule { }
