import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CarouselModule, ProgressbarModule } from 'ngx-bootstrap';
import { MomentModule } from 'ngx-moment';

import { HomeRoutingModule } from './home-routing.module';
import { HomePageComponent } from './components/home-page/home-page.component';
import { CoreModule } from 'src/app/core/core.module';
import { PostPreviewComponent } from './components/post-preview/post-preview.component';
import { ProjectPreviewComponent } from './components/project-preview/project-preview.component';

@NgModule({
  declarations: [HomePageComponent, PostPreviewComponent, ProjectPreviewComponent],
  imports: [
    CommonModule,
    HomeRoutingModule,
    CoreModule,
    MomentModule,
    CarouselModule,
    ProgressbarModule.forRoot()
  ]
})
export class HomeModule { }
