import { ProgressbarModule } from 'ngx-bootstrap';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { PageRoutingModule } from './page-routing.module';
import { PageDetailComponent } from './page-detail/page-detail.component';
import { PostModule } from './../post/post.module';
import { ProjectModule } from '../project/project.module';
import { SharedModule } from '../../shared/shared.module';


@NgModule({
  declarations: [PageDetailComponent],
  imports: [
    CommonModule,
    PageRoutingModule,
    SharedModule,
    ProjectModule,
    PostModule,
    ProgressbarModule.forRoot()
  ]
})
export class PageModule { }
