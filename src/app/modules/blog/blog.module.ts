import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { BlogRoutingModule } from './blog-routing.module';
import { BlogListingsPageComponent } from './components/blog-listings-page/blog-listings-page.component';
import { BlogDetailPageComponent } from './components/blog-detail-page/blog-detail-page.component';
import { SharedModule } from 'src/app/shared/shared.module';


@NgModule({
  declarations: [BlogListingsPageComponent, BlogDetailPageComponent],
  imports: [
    SharedModule,
    CommonModule,
    BlogRoutingModule
  ]
})
export class BlogModule { }
