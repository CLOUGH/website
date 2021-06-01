import { MomentModule } from 'ngx-moment';
import { SharedModule } from '../../shared/shared.module';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ProjectRoutingModule } from './project-routing.module';
import { ProjectPreviewComponent } from './project-preview/project-preview.component';
import { CarouselModule } from 'ngx-bootstrap';


@NgModule({
  declarations: [ProjectPreviewComponent],
  imports: [
    CommonModule,
    ProjectRoutingModule,
    MomentModule,
    SharedModule,
    CarouselModule,
  ],
  exports: [ProjectPreviewComponent]
})
export class ProjectModule { }
