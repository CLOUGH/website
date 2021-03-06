import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { FormsModule } from '@angular/forms';
import { BsDropdownModule, ModalModule,TooltipModule, BsDatepickerModule,TimepickerModule   } from 'ngx-bootstrap';
import { DragDropModule } from '@angular/cdk/drag-drop';

import { AdminRoutingModule } from './admin-routing.module';
import { AdminLeftSideNavComponent } from './components/admin-left-side-nav/admin-left-side-nav.component';
import { AdminLayoutComponent } from './layouts/admin-layout/admin-layout.component';
import { DashboardPageComponent } from './pages/dashboard-page/dashboard-page.component';
import { PageListingsPageComponent } from './pages/page-listings-page/page-listings-page.component';
import { PostsListingsPageComponent } from './pages/posts-listings-page/posts-listings-page.component';
import { ProjectsListingsPageComponent } from './pages/projects-listings-page/projects-listings-page.component';
import { PageCreatePageComponent } from './pages/page-create-page/page-create-page.component';
import { AdminTopNavbarComponent } from './components/admin-top-navbar/admin-top-navbar.component';
import { AdminPageHeadingComponent } from './components/admin-page-heading/admin-page-heading.component';
import { PageEditPageComponent } from './pages/page-edit-page/page-edit-page.component';
import { SharedModule } from 'src/app/shared/shared.module';
import { PageFormComponent } from './components/page-form/page-form.component';
import { PageLayoutEditorComponent } from './components/page-layout-editor/page-layout-editor.component';


@NgModule({
  declarations: [
    AdminLayoutComponent,
    AdminLeftSideNavComponent,
    DashboardPageComponent,
    PageListingsPageComponent,
    PostsListingsPageComponent,
    ProjectsListingsPageComponent,
    PageEditPageComponent,
    PageCreatePageComponent,
    AdminTopNavbarComponent,
    AdminPageHeadingComponent,
    PageFormComponent,
    PageLayoutEditorComponent
  ], imports: [
    CommonModule,
    AdminRoutingModule,
    FontAwesomeModule,
    FormsModule,
    SharedModule,
    BsDropdownModule.forRoot(),
    DragDropModule,
    TooltipModule.forRoot(),
    ModalModule.forRoot(),
    BsDatepickerModule.forRoot(),
    TimepickerModule.forRoot()
  ]
})
export class AdminModule { }
